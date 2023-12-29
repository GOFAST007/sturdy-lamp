// Copyright 2021 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package netplan

import (
	jc "github.com/juju/testing/checkers"
	gc "gopkg.in/check.v1"
	goyaml "gopkg.in/yaml.v2"
)

type NetplanConfigurationMergeSuite struct {
}

var _ = gc.Suite(&NetplanConfigurationMergeSuite{})

func (s *NetplanConfigurationMergeSuite) TestMergeNetplanConfigsWithEmptyBaseFile(c *gc.C) {
	base := ``

	src := `
network:
  version: 314
  renderer: whateverd
  ethernets:
    enp5s0:
      dhcp4: false
`[1:]

	s.assertMergeResult(c, base, src, src)
}

func (s *NetplanConfigurationMergeSuite) TestMergeNetplanConfigsWithEmptySourceFile(c *gc.C) {
	base := `
network:
  renderer: networkd
  version: 2
  ethernets:
    enp5s0:
      dhcp4: true
      optional: true
`[1:]

	src := ``

	s.assertMergeResult(c, base, src, base)
}

func (s *NetplanConfigurationMergeSuite) TestMergeNetplanConfigsScalarValuOverwrite(c *gc.C) {
	base := `
network:
  renderer: networkd
  version: 2
  ethernets:
    enp5s0:
      dhcp4: true
      optional: true
`[1:]

	src := `
network:
  version: 314        # Overwrite numeric value
  renderer: whateverd # Overwrite string value
  ethernets:
    enp5s0:
      dhcp4: false    # Overwrite bool value
`[1:]

	exp := `
network:
  version: 314
  renderer: whateverd
  ethernets:
    enp5s0:
      dhcp4: false
      optional: true
`[1:]

	s.assertMergeResult(c, base, src, exp)
}

func (s *NetplanConfigurationMergeSuite) TestMergeNetplanConfigsMergeMapValues(c *gc.C) {
	base := `
network:
  renderer: networkd
  version: 2
  ethernets:
    enp5s0:
      dhcp4: true
      optional: true
      nameservers:
        addresses:
        - 1.1.1.1
        - 8.8.8.8
`[1:]

	src := `
network:
  ethernets:
    newnic0:            # Add a new key to the map
      dhcp4: true
    enp5s0:
      dhcp4: false      # Overwrite value of existing key
      optional: true
      nameservers:
        search: [foo.local, baz.local] # Add new key to nested, existing map
`[1:]

	exp := `
network:
  renderer: networkd
  version: 2
  ethernets:
    enp5s0:
      dhcp4: false
      optional: true
      nameservers:
        addresses:
        - 1.1.1.1
        - 8.8.8.8
        search:
        - foo.local
        - baz.local
    newnic0:
      dhcp4: true
`[1:]

	s.assertMergeResult(c, base, src, exp)
}

func (s *NetplanConfigurationMergeSuite) TestMergeNetplanConfigsSliceConcatenation(c *gc.C) {
	base := `
network:
  ethernets:
    enp5s0:
      nameservers:
        addresses:
        - 1.1.1.1
        - 8.8.8.8
`[1:]

	src := `
network:
  ethernets:
    enp5s0:
      nameservers:
        addresses:
        - 8.8.8.8     # Value already exists in base sequence and should be skipped
        - 42.42.42.42 # Value does not exist in base sequence and should be added
        - 1.1.1.1     # Value already exists in base sequence and should be skipped
`[1:]

	exp := `
network:
  ethernets:
    enp5s0:
      nameservers:
        addresses:
        - 1.1.1.1
        - 8.8.8.8
        - 42.42.42.42
`[1:]

	s.assertMergeResult(c, base, src, exp)
}

func (s *NetplanConfigurationMergeSuite) TestMergeNetplanConfigsErrorsWhenMergingMaps(c *gc.C) {
	base := `
network:
  ethernets:
    enp5s0:
      nameservers:
        addresses:
        - 1.1.1.1
        - 8.8.8.8
`[1:]

	src := `
network:
  ethernets: "this is not the map you are looking for"
`[1:]

	s.assertMergeError(c, base, src, `merging configuration key "network": merging configuration key "ethernets": configuration values have different types \(destination: map.*, src: string\)`)
}

func (s *NetplanConfigurationMergeSuite) TestMergeNetplanConfigsErrorsWhenMergingSlices(c *gc.C) {
	base := `
network:
  ethernets:
    enp5s0:
      nameservers:
        addresses:
        - 1.1.1.1
        - 8.8.8.8
`[1:]

	src := `
network:
  ethernets:
    enp5s0:
      nameservers:
        addresses: "this is not the slice you are looking for"
`[1:]

	s.assertMergeError(c, base, src, `merging configuration key "network": merging configuration key "ethernets": merging configuration key "enp5s0": merging configuration key "nameservers": merging configuration key "addresses": configuration values have different types \(destination: \[\].*, src: string\)`)
}

func (s *NetplanConfigurationMergeSuite) assertMergeResult(c *gc.C, base, src, exp string) {
	var (
		baseMap, srcMap, expMap map[interface{}]interface{}
	)

	c.Assert(goyaml.Unmarshal([]byte(base), &baseMap), jc.ErrorIsNil)
	c.Assert(goyaml.Unmarshal([]byte(src), &srcMap), jc.ErrorIsNil)
	c.Assert(goyaml.Unmarshal([]byte(exp), &expMap), jc.ErrorIsNil)

	mergeRes, err := mergeNetplanConfigs(baseMap, srcMap)
	c.Assert(err, jc.ErrorIsNil)

	mergeResMap, ok := mergeRes.(map[interface{}]interface{})
	c.Assert(ok, jc.IsTrue, gc.Commentf("expected merge result to be a map[interface{}]interface{}; got %T", mergeRes))
	c.Assert(mergeResMap, gc.DeepEquals, expMap)
}

func (s *NetplanConfigurationMergeSuite) assertMergeError(c *gc.C, base, src, expErr string) {
	var (
		baseMap, srcMap map[interface{}]interface{}
	)

	c.Assert(goyaml.Unmarshal([]byte(base), &baseMap), jc.ErrorIsNil)
	c.Assert(goyaml.Unmarshal([]byte(src), &srcMap), jc.ErrorIsNil)

	_, err := mergeNetplanConfigs(baseMap, srcMap)
	c.Assert(err, gc.ErrorMatches, expErr)
}
