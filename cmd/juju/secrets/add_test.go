// Copyright 2023 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package secrets_test

import (
	"os"
	"path/filepath"

	"github.com/juju/cmd/v3/cmdtesting"
	jujutesting "github.com/juju/testing"
	jc "github.com/juju/testing/checkers"
	"go.uber.org/mock/gomock"
	gc "gopkg.in/check.v1"

	"github.com/juju/juju/cmd/juju/secrets"
	"github.com/juju/juju/cmd/juju/secrets/mocks"
	coresecrets "github.com/juju/juju/core/secrets"
	"github.com/juju/juju/jujuclient"
)

type addSuite struct {
	jujutesting.IsolationSuite
	store      *jujuclient.MemStore
	secretsAPI *mocks.MockAddSecretsAPI
}

var _ = gc.Suite(&addSuite{})

func (s *addSuite) SetUpTest(c *gc.C) {
	s.IsolationSuite.SetUpTest(c)
	store := jujuclient.NewMemStore()
	store.Controllers["mycontroller"] = jujuclient.ControllerDetails{}
	store.CurrentControllerName = "mycontroller"
	s.store = store
}

func (s *addSuite) setup(c *gc.C) *gomock.Controller {
	ctrl := gomock.NewController(c)
	s.secretsAPI = mocks.NewMockAddSecretsAPI(ctrl)
	return ctrl
}

func (s *addSuite) TestAddDataFromArg(c *gc.C) {
	defer s.setup(c).Finish()

	uri := coresecrets.NewURI()
	s.secretsAPI.EXPECT().CreateSecret("my-secret", "this is a secret.", map[string]string{"foo": "YmFy"}).Return(uri.String(), nil)
	s.secretsAPI.EXPECT().Close().Return(nil)

	ctx, err := cmdtesting.RunCommand(c, secrets.NewAddCommandForTest(s.store, s.secretsAPI), "my-secret", "foo=bar", "--info", "this is a secret.")
	c.Assert(err, jc.ErrorIsNil)
	out := cmdtesting.Stdout(ctx)
	c.Assert(out, gc.Equals, uri.String()+"\n")
}

func (s *addSuite) TestAddDataFromFile(c *gc.C) {
	defer s.setup(c).Finish()

	uri := coresecrets.NewURI()
	s.secretsAPI.EXPECT().CreateSecret("my-secret", "this is a secret.", map[string]string{"foo": "YmFy"}).Return(uri.String(), nil)
	s.secretsAPI.EXPECT().Close().Return(nil)

	dir := c.MkDir()
	path := filepath.Join(dir, "data.txt")
	data := `
foo: bar
    `
	err := os.WriteFile(path, []byte(data), 0644)
	c.Assert(err, jc.ErrorIsNil)

	ctx, err := cmdtesting.RunCommand(c, secrets.NewAddCommandForTest(s.store, s.secretsAPI), "my-secret", "--file", path, "--info", "this is a secret.")
	c.Assert(err, jc.ErrorIsNil)
	out := cmdtesting.Stdout(ctx)
	c.Assert(out, gc.Equals, uri.String()+"\n")
}

func (s *addSuite) TestAddEmptyData(c *gc.C) {
	defer s.setup(c).Finish()

	_, err := cmdtesting.RunCommand(c, secrets.NewAddCommandForTest(s.store, s.secretsAPI), "my-secret", "--info", "this is a secret.")
	c.Assert(err, gc.ErrorMatches, `missing secret value or filename`)
}
