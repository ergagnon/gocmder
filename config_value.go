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
	"strconv"
	"strings"
)

const (
	descKey         = "desc"
	defaultValueKey = "default"
	isHiddenKey     = "hidden"
	isRequiredKey   = "required"
)

type configValue struct {
	name         string
	kind         reflect.Kind
	desc         string
	defaultValue any
	isHidden     bool
	isRequired   bool
}

func newConfigValue(name string, sf reflect.StructField) configValue {
	var defaultValue any = sf.Tag.Get(defaultValueKey)

	kind := sf.Type.Kind()

	if kind == reflect.Int {
		defaultValue, _ = strconv.Atoi(defaultValue.(string))
	}

	if kind == reflect.Bool {
		defaultValue, _ = strconv.ParseBool(defaultValue.(string))
	}

	isHidden, _ := strconv.ParseBool(sf.Tag.Get(isHiddenKey))
	isRequired, _ := strconv.ParseBool(sf.Tag.Get(isRequiredKey))

	return configValue{
		name:         name,
		kind:         kind,
		desc:         sf.Tag.Get(descKey),
		defaultValue: defaultValue,
		isHidden:     isHidden,
		isRequired:   isRequired,
	}
}

func createConfigValues(cfg any) []configValue {
	configValues := make([]configValue, 0)
	recursivelyExtractConfigValues(cfg, "", &configValues)
	return configValues
}

func recursivelyExtractConfigValues(cfg any, prefix string, cfgValues *[]configValue) {
	cfgType := reflect.TypeOf(cfg)

	if cfgType.Name() == "StructField" {
		sf := cfg.(reflect.StructField)
		cfgType = sf.Type
	}

	for _, vf := range reflect.VisibleFields(cfgType) {
		name := strings.ToLower(vf.Name)

		if vf.Type.Kind() == reflect.Struct {
			recursivelyExtractConfigValues(vf, prefix+name+".", cfgValues)
		} else {
			*cfgValues = append(*cfgValues, newConfigValue(prefix+name, vf))
		}
	}
}
