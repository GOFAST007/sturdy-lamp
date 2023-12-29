// Copyright 2018 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package mongo

import (
	"os"
	"os/exec"

	"github.com/juju/errors"
)

// SearchTools represents the OS functionality we need to find the correct MongoDB executable.
// The mock for this (used in testing) is automatically generated by 'go generate' from the following line
//
//go:generate go run go.uber.org/mock/mockgen -package mongo -destination searchtoolsmock_test.go github.com/juju/juju/mongo SearchTools
type SearchTools interface {
	// GetCommandOutput execs the given command, and returns the CombinedOutput, or any error that occurred.
	GetCommandOutput(name string, arg ...string) (string, error)

	// Exists just returns if a given path is available (eg os.Stat() has a value)
	Exists(string) bool
}

// MongodFinder searches expected paths to find a version of Mongo and determine what version it is.
type MongodFinder struct {
	search SearchTools
}

// NewMongodFinder returns a type that will help search for mongod, using normal OS tools.
func NewMongodFinder() *MongodFinder {
	return &MongodFinder{
		search: OSSearchTools{},
	}
}

// InstalledAt tries to find an installed mongo snap.
func (m *MongodFinder) InstalledAt() (string, error) {
	if m.search.Exists(JujuDbSnapMongodPath) {
		return JujuDbSnapMongodPath, nil
	}
	return "", errors.Errorf("juju-db snap not installed, no mongo at %s", JujuDbSnapMongodPath)
}

type OSSearchTools struct{}

func (OSSearchTools) Exists(name string) bool {
	_, err := os.Stat(name)
	return err == nil
}

func (OSSearchTools) GetCommandOutput(name string, arg ...string) (string, error) {
	cmd := exec.Command(name, arg...)
	output, err := cmd.CombinedOutput()
	return string(output), errors.Trace(err)
}