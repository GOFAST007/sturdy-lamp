// Copyright 2023 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package secretspruner

import (
	"strconv"
	"strings"

	"github.com/juju/collections/set"
	"github.com/juju/errors"
	"github.com/juju/worker/v3"
	"github.com/juju/worker/v3/catacomb"

	coresecrets "github.com/juju/juju/core/secrets"
	"github.com/juju/juju/core/watcher"
)

// logger is here to stop the desire of creating a package level logger.
// Don't do this, instead use the one passed as manifold config.
type logger interface{}

var _ logger = struct{}{}

// Logger represents the methods used by the worker to log information.
type Logger interface {
	Debugf(string, ...interface{})
	Warningf(string, ...interface{})
	Infof(string, ...interface{})
}

// SecretsFacade instances provide a set of API for the worker to deal with secret prune.
type SecretsFacade interface {
	WatchRevisionsToPrune() (watcher.StringsWatcher, error)
	DeleteRevisions(uri *coresecrets.URI, revisions ...int) error
}

// Config defines the operation of the Worker.
type Config struct {
	SecretsFacade
	Logger Logger
}

// Validate returns an error if config cannot drive the Worker.
func (config Config) Validate() error {
	if config.SecretsFacade == nil {
		return errors.NotValidf("nil SecretsFacade")
	}
	if config.Logger == nil {
		return errors.NotValidf("nil Logger")
	}
	return nil
}

// NewWorker returns a secretspruner Worker backed by config, or an error.
func NewWorker(config Config) (worker.Worker, error) {
	if err := config.Validate(); err != nil {
		return nil, errors.Trace(err)
	}

	w := &Worker{config: config}
	err := catacomb.Invoke(catacomb.Plan{
		Site: &w.catacomb,
		Work: w.loop,
	})
	return w, errors.Trace(err)
}

// Worker prunes the user supplied secret revisions.
type Worker struct {
	catacomb catacomb.Catacomb
	config   Config
}

// Kill is defined on worker.Worker.
func (w *Worker) Kill() {
	w.catacomb.Kill(nil)
}

// Wait is part of the worker.Worker interface.
func (w *Worker) Wait() error {
	return w.catacomb.Wait()
}

func (w *Worker) loop() (err error) {
	watcher, err := w.config.SecretsFacade.WatchRevisionsToPrune()
	if err != nil {
		return errors.Trace(err)
	}
	if err := w.catacomb.Add(watcher); err != nil {
		return errors.Trace(err)
	}

	for {
		select {
		case <-w.catacomb.Dying():
			return errors.Trace(w.catacomb.ErrDying())
		// TODO: watch for secret's auto-prune config changes.
		// then delete any obsolete revisions.
		case changes, ok := <-watcher.Changes():
			if !ok {
				return errors.New("secret prune changed watch closed")
			}
			w.config.Logger.Debugf("got user supplied secret revisions to prune")

			if len(changes) == 0 {
				w.config.Logger.Debugf("no secret revisions to prune")
				continue
			}
			revisions, err := w.processChanges(changes...)
			if err != nil {
				return errors.Trace(err)
			}
			for uriStr, revs := range revisions {
				w.config.Logger.Debugf("pruning secret revisions %q: %v", uriStr, revs.SortedValues())
				uri, err := coresecrets.ParseURI(uriStr)
				if err != nil {
					return errors.Trace(err)
				}
				if err := w.config.SecretsFacade.DeleteRevisions(uri, revs.SortedValues()...); err != nil {
					return errors.Trace(err)
				}
			}
		}
	}
}

func (w *Worker) processChanges(changes ...string) (map[string]set.Ints, error) {
	out := make(map[string]set.Ints)
	for _, revInfo := range changes {
		w.config.Logger.Warningf("revInfo %q, out %#v", revInfo, out)
		parts := strings.Split(revInfo, "/")
		uri := parts[0]
		if len(parts) < 2 {
			// This should never happen.
			w.config.Logger.Debugf("secret %q has been removed, no need to prune", revInfo)
			continue
		}
		rev, err := strconv.Atoi(parts[1])
		if err != nil {
			// This should never happen.
			return nil, errors.NotValidf("secret revision %q for %q", parts[1], uri)
		}
		if _, ok := out[uri]; !ok {
			out[uri] = set.NewInts()
		}
		out[uri].Add(rev)
	}
	return out, nil
}
