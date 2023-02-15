// Copyright © 2023 Eric Gagnon <github.com/ergagnon>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gocmder

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/suite"
)

type cmderTestSuite struct {
	suite.Suite
	buf bytes.Buffer
}

func (s *cmderTestSuite) TestNewCmder() {
	cmder, err := NewCmder(rootConfig{}, func(cfg any) {})

	s.NoError(err)

	cmder.Cobra().SetArgs([]string{"--help"})
	cmder.Cobra().SetOutput(&s.buf)

	err = cmder.Execute()

	s.NoError(err)
}

func (s *cmderTestSuite) TestNewCmderWithoutRequiredArgs() {
	cmder, err := NewCmder(rootConfig{}, func(cfg any) {})

	s.NoError(err)

	cmder.Cobra().SetArgs([]string{})
	cmder.Cobra().SetOutput(&s.buf)

	err = cmder.Execute()

	s.EqualError(err, "required flag(s) \"foo\" not set")
}

func (s *cmderTestSuite) TestNewCmderWithInvalidArgs() {
	cmder, err := NewCmder(rootConfig{}, func(cfg any) {})

	s.NoError(err)

	cmder.Cobra().SetArgs([]string{"--foo", "foo", "--bar", "bar"})
	cmder.Cobra().SetOutput(&s.buf)

	err = cmder.Execute()

	s.EqualError(err, "invalid argument \"bar\" for \"--bar\" flag: strconv.ParseInt: parsing \"bar\": invalid syntax")
}

func (s *cmderTestSuite) TestNewCmderValidArgs() {
	cmder, err := NewCmder(rootConfig{}, func(cfg any) {
		c := cfg.(rootConfig)
		s.Equal("im a foo", c.Foo)
		s.Equal(2, c.Bar)
		s.Equal(float32(1.2), c.Child.Decimal)
		s.Equal(true, c.Child.Boolean)
		s.Equal("hide and seek", c.Child.Hidden)
	})

	s.NoError(err)

	cmder.Cobra().SetArgs([]string{"--foo", "im a foo"})
	cmder.Cobra().SetOutput(&s.buf)

	err = cmder.Execute()

	s.NoError(err)
}

func (s *cmderTestSuite) TestNewCmderWithVersionOptions() {
	cmder, err := NewCmder(rootConfig{}, func(cfg any) {}, WithVersion("1.2.3"))

	s.NoError(err)

	cmder.cobra.SetArgs([]string{"--version"})
	cmder.cobra.SetOutput(&s.buf)

	err = cmder.Execute()
	s.NoError(err)

	version := s.buf.Bytes()
	s.Equal("version 1.2.3\n", string(version))
}

func (s *cmderTestSuite) TestNewCmderWithEnvVariable_OverrideDefaultValue() {
	onfinalizeCalled := false
	cmder, err := NewCmder(rootConfig{}, func(cfg any) {
		c := cfg.(rootConfig)
		s.Equal("im a foo", c.Foo)
		s.Equal(3, c.Bar)
		s.Equal(float32(3.14), c.Child.Decimal)
		s.Equal(false, c.Child.Boolean)
		s.Equal("i found you", c.Child.Hidden)
		onfinalizeCalled = true
	}, WithPrefix("TEST"))

	s.NoError(err)

	s.T().Setenv("TEST_CHILD_HIDDEN", "i found you")
	s.T().Setenv("TEST_CHILD_DECIMAL", "3.14")
	s.T().Setenv("TEST_BAR", "3")
	s.T().Setenv("TEST_CHILD_BOOLEAN", "false")

	cmder.cobra.SetArgs([]string{"--foo", "im a foo"})
	cmder.cobra.SetOutput(&s.buf)

	err = cmder.Execute()
	s.NoError(err)

	s.True(onfinalizeCalled)
}

func (s *cmderTestSuite) TestNewCmderWithConfigFile() {
	fs := afero.NewMemMapFs()

	dir, err := os.Getwd()

	s.NoError(err)

	err = fs.Mkdir(dir, 0755)
	s.NoError(err)

	file, err := fs.Create(filepath.Join(dir, "config.yaml"))
	s.NoError(err)

	_, err = file.Write([]byte(`
foo: im a foo
bar: 3
child:
  decimal: 3.14
  boolean: false
  hidden: i found you
`))

	s.NoError(err)

	onfinalizeCalled := false
	cmder, err := NewCmder(rootConfig{}, func(cfg any) {
		c := cfg.(rootConfig)
		s.Equal("im a foo", c.Foo)
		s.Equal(3, c.Bar)
		s.Equal(float32(3.14), c.Child.Decimal)
		s.Equal(false, c.Child.Boolean)
		s.Equal("i found you", c.Child.Hidden)
		onfinalizeCalled = true
	}, WithFS(fs), WithConfigFile(filepath.Join(dir, "config.yaml")))

	s.NoError(err)

	cmder.cobra.SetArgs([]string{"--foo", "im a foo"})
	cmder.cobra.SetOutput(&s.buf)
	

	err = cmder.Execute()

	s.NoError(err)
	s.True(onfinalizeCalled)
}

func TestCmderTestSuite(t *testing.T) {
	suite.Run(t, new(cmderTestSuite))
}

func (s *cmderTestSuite) TearDownTest() {
	s.buf.Reset()
}

type rootConfig struct {
	Foo   string `desc:"foo" required:"true"`
	Bar   int    `desc:"bar" default:"2"`
	Child childConfig
}

type childConfig struct {
	Decimal float32 `desc:"decimal" default:"1.2"`
	Boolean bool    `desc:"boolean" default:"true"`
	Hidden  string  `desc:"hidden" default:"hide and seek" hidden:"true"`
}
