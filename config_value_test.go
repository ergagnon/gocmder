// Copyright © 2023 Eric Gagnon <github.com/ergagnon>

// This program is free software: you can redistribute it and/or modify it
// under the terms of the GNU General Public License as published by the
// Free Software Foundation, either version 3 of the License, or (at your option)
// any later version.

// This program is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS
// FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

// You should have received a copy of the GNU General Public License along with
// this program. If not, see <https://www.gnu.org/licenses/>.

package gocmder

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type configTest struct {
	foo string `desc:"foo" default:"foo" required:"true" hidden:"false"`
	bar string `desc:"bar" default:"bar" required:"false" hidden:"true"`
	sc  stringConfigTest
}

type stringConfigTest struct {
	foostring string `desc:"foostring" default:"foostring" required:"true" hidden:"false"`
	barstring string `desc:"barstring" default:"barstring" required:"false" hidden:"true"`
	ic        intConfigTest
}

type intConfigTest struct {
	fooint int `desc:"fooint" default:"1" required:"true" hidden:"false"`
	barint int `desc:"barint" default:"2" required:"false" hidden:"true"`
	bc     boolConfigTest
}

type boolConfigTest struct {
	foobool bool `desc:"foobool" default:"true" required:"true" hidden:"false"`
	barbool bool `desc:"barbool" default:"false" required:"false" hidden:"true"`
}

type configValueTestSuite struct {
	suite.Suite
}

func (s *configValueTestSuite) TestConfigValue() {
	cfgs := make([]configValue, 0)
	recursivelyExtractConfigValues(configTest{}, "", &cfgs)

	s.Equal(8, len(cfgs))
}

func TestConfigValueTestSuite(t *testing.T) {
	suite.Run(t, new(configValueTestSuite))
}
