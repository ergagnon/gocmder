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
	"reflect"
	"testing"

	"github.com/stretchr/testify/suite"
)

type configItemTestSuite struct {
	suite.Suite
}

func (s *configItemTestSuite) TestCreateConfigItems() {
	cfgs := createConfigItems(configTest{})

	s.Equal(8, len(cfgs))

	s.ElementsMatch(
		[]configItem{
			{
				name:            "foo",
				kind:            reflect.String,
				desc:            "foo",
				defaultValue:    any("foo"),
				hasDefaultValue: true,
				isHidden:        false,
				isRequired:      true,
			},
			{
				name:            "bar",
				kind:            reflect.String,
				desc:            "bar",
				defaultValue:    any("bar"),
				hasDefaultValue: true,
				isHidden:        true,
				isRequired:      false,
			},
			{
				name:            "sc.foostring",
				kind:            reflect.String,
				desc:            "foostring",
				defaultValue:    any("foostring"),
				hasDefaultValue: true,
				isHidden:        false,
				isRequired:      true,
			},
			{
				name:            "sc.barstring",
				kind:            reflect.String,
				desc:            "barstring",
				defaultValue:    any(""),
				hasDefaultValue: false,
				isHidden:        true,
				isRequired:      false,
			},
			{
				name:            "sc.ic.fooint",
				kind:            reflect.Int,
				desc:            "fooint",
				defaultValue:    any(1),
				hasDefaultValue: true,
				isHidden:        false,
				isRequired:      true,
			},
			{
				name:            "sc.ic.barint",
				kind:            reflect.Int,
				desc:            "barint",
				defaultValue:    any(0),
				hasDefaultValue: false,
				isHidden:        true,
				isRequired:      false,
			},
			{
				name:            "sc.ic.bc.foobool",
				kind:            reflect.Bool,
				desc:            "foobool",
				defaultValue:    any(true),
				hasDefaultValue: true,
				isHidden:        false,
				isRequired:      true,
			},
			{
				name:            "sc.ic.bc.barbool",
				kind:            reflect.Bool,
				desc:            "barbool",
				defaultValue:    any(false),
				hasDefaultValue: false,
				isHidden:        true,
				isRequired:      false,
			},
		},
		cfgs,
	)
}

func (s *configItemTestSuite) TestCreateConfigItemsWithFloat() {
	cfgs := createConfigItems(floatConfigTest{})
	s.Equal(1, len(cfgs))

	item := cfgs[0]

	s.Equal("foofloat", item.name)
	s.Equal(reflect.Float32, item.kind)
	s.Equal("foofloat", item.desc)
	s.False(item.isRequired)
	s.False(item.isHidden)
	s.InDelta(float32(1.23), item.defaultValue, 0.0001)
	s.True(item.hasDefaultValue)
}

func TestConfigItemTestSuite(t *testing.T) {
	suite.Run(t, new(configItemTestSuite))
}

type floatConfigTest struct {
	foofloat float32 `desc:"foofloat" default:"1.23"`
}

type configTest struct {
	foo string `desc:"foo" default:"foo" required:"true" hidden:"false"`
	bar string `desc:"bar" default:"bar" required:"false" hidden:"true"`
	sc  stringConfigTest
}

type stringConfigTest struct {
	foostring string `desc:"foostring" default:"foostring" required:"true" hidden:"false"`
	barstring string `desc:"barstring" required:"false" hidden:"true"`
	ic        intConfigTest
}

type intConfigTest struct {
	fooint int `desc:"fooint" default:"1" required:"true" hidden:"false"`
	barint int `desc:"barint" required:"false" hidden:"true"`
	bc     boolConfigTest
}

type boolConfigTest struct {
	foobool bool `desc:"foobool" default:"true" required:"true" hidden:"false"`
	barbool bool `desc:"barbool" required:"false" hidden:"true"`
}
