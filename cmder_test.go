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
	"testing"

	"github.com/stretchr/testify/suite"
)

type cmderTestSuite struct {
	suite.Suite
}

func (s *cmderTestSuite) TestNewCmder() {
	cmder, err := NewCmder(rootConfig{}, func(cfg any) error {
		return nil
	})

	s.NoError(err)

	cmder.cobra.SetArgs([]string{"--help"})

	err = cmder.Execute()

	s.NoError(err)
}

func (s *cmderTestSuite) TestNewCmderWithoutRequiredArgs() {
	cmder, err := NewCmder(rootConfig{}, func(cfg any) error {
		return nil
	})

	s.NoError(err)

	cmder.cobra.SetArgs([]string{})

	err = cmder.Execute()

	s.EqualError(err, "required flag(s) \"foo\" not set")
}

func (s *cmderTestSuite) TestNewCmderWithInvalidArgs() {
	cmder, err := NewCmder(rootConfig{}, func(cfg any) error {
		return nil
	})

	s.NoError(err)

	cmder.cobra.SetArgs([]string{"--foo", "foo", "--bar", "bar"})
	err = cmder.Execute()

	s.EqualError(err, "invalid argument \"bar\" for \"--bar\" flag: strconv.ParseInt: parsing \"bar\": invalid syntax")
}

func (s *cmderTestSuite) TestNewCmderValidArgs() {
	cmder, err := NewCmder(rootConfig{}, func(cfg any) error {
		c := cfg.(rootConfig)
		s.Equal("im a foo", c.Foo)
		s.Equal(2, c.Bar)
		s.Equal(float32(1.2), c.Child.Decimal)
		s.Equal(true, c.Child.Boolean)
		s.Equal("hide and seek", c.Child.Hidden)

		return nil
	})

	s.NoError(err)

	cmder.cobra.SetArgs([]string{"--foo", "im a foo"})

	err = cmder.Execute()

	s.NoError(err)
}

func TestCmderTestSuite(t *testing.T) {
	suite.Run(t, new(cmderTestSuite))
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
