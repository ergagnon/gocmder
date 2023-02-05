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
	"strconv"
	"strings"
)

const (
	descKey     = "desc"
	defaultValueKey = "default"
	isHiddenKey     = "hidden"
	isRequiredKey   = "required"
)

type configValue struct {
	name 	  string
	kind reflect.Kind
	desc 	  string
	defaultValue any
	isHidden bool
	isRequired bool
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
		name: name,
		kind: kind,
		desc: sf.Tag.Get(descKey),	
		defaultValue: defaultValue,
		isHidden: isHidden,
		isRequired: isRequired,
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
			recursivelyExtractConfigValues(vf, prefix + name + ".", cfgValues)
		} else {
			*cfgValues = append(*cfgValues, newConfigValue(prefix + name, vf))
		}
	}	
}