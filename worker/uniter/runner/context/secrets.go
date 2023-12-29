// Copyright 2022 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package context

import (
	"github.com/juju/errors"
	"github.com/juju/loggo"

	"github.com/juju/juju/api/agent/uniter"
	"github.com/juju/juju/core/secrets"
)

// secretsChangeRecorder records the intent tp update, remove or
// change secret access permission during a hook execution.
type secretsChangeRecorder struct {
	logger loggo.Logger

	pendingCreates     map[string]uniter.SecretCreateArg
	pendingUpdates     map[string]uniter.SecretUpdateArg
	pendingDeletes     map[string]uniter.SecretDeleteArg
	pendingGrants      map[string]uniter.SecretGrantRevokeArgs
	pendingRevokes     map[string]uniter.SecretGrantRevokeArgs
	pendingTrackLatest map[string]bool
}

func newSecretsChangeRecorder(logger loggo.Logger) *secretsChangeRecorder {
	return &secretsChangeRecorder{
		logger:             logger,
		pendingCreates:     make(map[string]uniter.SecretCreateArg),
		pendingUpdates:     make(map[string]uniter.SecretUpdateArg),
		pendingDeletes:     make(map[string]uniter.SecretDeleteArg),
		pendingGrants:      make(map[string]uniter.SecretGrantRevokeArgs),
		pendingRevokes:     make(map[string]uniter.SecretGrantRevokeArgs),
		pendingTrackLatest: make(map[string]bool),
	}
}

func (s *secretsChangeRecorder) haveContentUpdates() bool {
	return len(s.pendingCreates) > 0 || len(s.pendingUpdates) > 0 ||
		len(s.pendingDeletes) > 0
}

func (s *secretsChangeRecorder) create(arg uniter.SecretCreateArg) error {
	delete(s.pendingDeletes, arg.URI.ID)
	for _, c := range s.pendingCreates {
		if c.Label != nil && arg.Label != nil && *c.Label == *arg.Label {
			return errors.AlreadyExistsf("secret with label %q", *arg.Label)
		}
	}
	s.pendingCreates[arg.URI.ID] = arg
	return nil
}

func (s *secretsChangeRecorder) update(arg uniter.SecretUpdateArg) {
	delete(s.pendingDeletes, arg.URI.ID)
	if c, ok := s.pendingCreates[arg.URI.ID]; ok {
		if arg.Label != nil {
			c.Label = arg.Label
		}
		if arg.Description != nil {
			c.Description = arg.Description
		}
		if !arg.Value.IsEmpty() {
			c.Value = arg.Value
		}
		if arg.RotatePolicy != nil {
			c.RotatePolicy = arg.RotatePolicy
		}
		if arg.ExpireTime != nil {
			c.ExpireTime = arg.ExpireTime
		}
		s.pendingCreates[arg.URI.ID] = c
		return
	}
	s.pendingUpdates[arg.URI.ID] = arg
}

func (s *secretsChangeRecorder) remove(uri *secrets.URI, revision *int) {
	delete(s.pendingCreates, uri.ID)
	delete(s.pendingUpdates, uri.ID)
	delete(s.pendingGrants, uri.ID)
	delete(s.pendingRevokes, uri.ID)
	delete(s.pendingTrackLatest, uri.ID)
	s.pendingDeletes[uri.ID] = uniter.SecretDeleteArg{URI: uri, Revision: revision}
}

func (s *secretsChangeRecorder) grant(arg uniter.SecretGrantRevokeArgs) {
	s.pendingGrants[arg.URI.ID] = arg
}

func (s *secretsChangeRecorder) revoke(arg uniter.SecretGrantRevokeArgs) {
	s.pendingRevokes[arg.URI.ID] = arg
}
