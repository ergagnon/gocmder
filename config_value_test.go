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
	"reflect"
	"testing"

	"github.com/stretchr/testify/suite"
)

type configValueTestSuite struct {
	suite.Suite
}

func (s *configValueTestSuite) TestCreateConfigValues() {
	cfgs := createConfigValues(configTest{})

	s.Equal(8, len(cfgs))

	s.ElementsMatch(
		[]configValue{
			{
				name:         "foo",
				kind:         reflect.String,
				desc:         "foo",
				defaultValue: any("foo"),
				isHidden:     false,
				isRequired:   true,
			},
			{
				name:         "bar",
				kind:         reflect.String,
				desc:         "bar",
				defaultValue: any("bar"),
				isHidden:     true,
				isRequired:   false,
			},
			{
				name:         "sc.foostring",
				kind:         reflect.String,
				desc:         "foostring",
				defaultValue: any("foostring"),
				isHidden:     false,
				isRequired:   true,
			},
			{
				name:         "sc.barstring",
				kind:         reflect.String,
				desc:         "barstring",
				defaultValue: any("barstring"),
				isHidden:     true,
				isRequired:   false,
			},
			{
				name:         "sc.ic.fooint",
				kind:         reflect.Int,
				desc:         "fooint",
				defaultValue: any(1),
				isHidden:     false,
				isRequired:   true,
			},
			{
				name:         "sc.ic.barint",
				kind:         reflect.Int,
				desc:         "barint",
				defaultValue: any(2),
				isHidden:     true,
				isRequired:   false,
			},
			{
				name:         "sc.ic.bc.foobool",
				kind:         reflect.Bool,
				desc:         "foobool",
				defaultValue: any(true),
				isHidden:     false,
				isRequired:   true,
			},
			{
				name:         "sc.ic.bc.barbool",
				kind:         reflect.Bool,
				desc:         "barbool",
				defaultValue: any(false),
				isHidden:     true,
				isRequired:   false,
			},
		},
		cfgs,
	)
}

func TestConfigValueTestSuite(t *testing.T) {
	suite.Run(t, new(configValueTestSuite))
}

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
