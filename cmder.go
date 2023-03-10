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
	"fmt"
	"reflect"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Cmder struct {
	cfg       any
	cobra     *cobra.Command
	viper     *viper.Viper
	longDesc  string
	shortDesc string
	version   string
	envPrefix string
}

type OnFinalizeFunc func(cfg any)

// NewCmder creates a new Cmder instance. It takes a config struct, a callback function
// when the command is finalized and a variadic list of options.
func NewCmder(cfg any, onFinalize OnFinalizeFunc, opts ...CmderOption) (*Cmder, error) {
	c := &Cmder{
		cfg:   cfg,
		viper: viper.New(),
	}

	for _, opt := range opts {
		opt(c)
	}

	c.cobra = &cobra.Command{
		Short:   c.shortDesc,
		Long:    c.longDesc,
		Version: c.version,
		PreRunE: c.preRunE,
		RunE:    c.runE,
	}

	cobra.OnFinalize(func() {
		onFinalize(c.cfg)
	})

	if err := c.init(createConfigItems(cfg)); err != nil {
		return nil, err
	}

	return c, nil
}

// Returns the Cobra instance.
func (c *Cmder) Cobra() *cobra.Command {
	return c.cobra
}

// Returns the Viper instance.
func (c *Cmder) Viper() *viper.Viper {
	return c.viper
}

// Executes the command.
func (c *Cmder) Execute() error {
	return c.cobra.Execute()
}

func (c *Cmder) init(items []configItem) error {
	for _, item := range items {
		if err := c.addCliFlag(item); err != nil {
			return err
		}

		if item.hasDefaultValue {
			if err := c.setDefaultConfigValue(item); err != nil {
				return err
			}
		}

		if err := c.connectViperAndCobra(item); err != nil {
			return err
		}
	}

	return nil
}

func (c *Cmder) addCliFlag(item configItem) error {
	if item.isHidden {
		return nil
	}

	flagName := toFlagName(item.name)

	switch item.kind {
	case reflect.String:
		c.cobra.Flags().String(flagName, item.defaultValue.(string), item.desc)
	case reflect.Bool:
		c.cobra.Flags().Bool(flagName, item.defaultValue.(bool), item.desc)
	case reflect.Int:
		c.cobra.Flags().Int(flagName, item.defaultValue.(int), item.desc)
	case reflect.Float32:
		c.cobra.Flags().Float32(flagName, item.defaultValue.(float32), item.desc)
	default:
		return fmt.Errorf("unsupported type %s", item.kind)
	}

	if item.isRequired {
		c.cobra.MarkFlagRequired(item.name)
	}

	return nil
}

func (c *Cmder) setDefaultConfigValue(item configItem) error {
	switch item.kind {
	case reflect.String:
		c.viper.SetDefault(item.name, item.defaultValue.(string))
	case reflect.Bool:
		c.viper.SetDefault(item.name, item.defaultValue.(bool))
	case reflect.Int:
		c.viper.SetDefault(item.name, item.defaultValue.(int))
	case reflect.Float32:
		c.viper.SetDefault(item.name, item.defaultValue.(float32))
	default:
		return fmt.Errorf("unsupported type %s", item.kind)
	}

	return nil
}

func (c *Cmder) connectViperAndCobra(item configItem) error {
	if !item.isHidden {
		if err := c.viper.BindPFlag(item.name, c.cobra.Flags().Lookup(toFlagName(item.name))); err != nil {
			return err
		}
	}

	if err := c.viper.BindEnv(item.name, toEnvName(c.envPrefix, item.name)); err != nil {
		return err
	}

	return nil
}

func toFlagName(name string) string {
	return strings.ToLower(strings.Replace(name, ".", "-", -1))
}

func toEnvName(prefix, name string) string {
	if prefix != "" {
		name = prefix + "_" + name
	}

	return strings.ToUpper(strings.Replace(name, ".", "_", -1))
}

func (c *Cmder) preRunE(cmd *cobra.Command, _ []string) error {
	if err := c.viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	return nil
}

func (c *Cmder) runE(cmd *cobra.Command, _ []string) error {
	if err := c.viper.Unmarshal(&c.cfg); err != nil {
		return err
	}

	return nil
}
