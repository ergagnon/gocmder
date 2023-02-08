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

type OnFinalizeFunc func(cfg any) error

func NewCmder(cfg any, onFinalize OnFinalizeFunc, opts ...CmderOption) *Cmder {
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
		RunE:    c.run(),
	}

	cobra.OnFinalize(func() {
		onFinalize(c.cfg)
	})

	c.init(createConfigItems(cfg))

	return c
}

func (c *Cmder) Cobra() *cobra.Command {
	return c.cobra
}

func (c *Cmder) Viper() *viper.Viper {
	return c.viper
}

func (c *Cmder) Execute() error {
	return c.cobra.Execute()
}

func (c *Cmder) init(items []configItem) error {
	
	for _, item := range items {
		if err := c.addCliFlag(item); err != nil {
			return err
		}
		
		if err := c.setDefaultConfigValue(item); err != nil {
			return err
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
	
	switch item.kind {
	case reflect.String:
		c.cobra.Flags().String(item.name, item.defaultValue.(string), item.desc)
	case reflect.Bool:
		c.cobra.Flags().Bool(item.name, item.defaultValue.(bool), item.desc)
	case reflect.Int:
		c.cobra.Flags().Int(item.name, item.defaultValue.(int), item.desc)
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
	default:
		return fmt.Errorf("unsupported type %s", item.kind)
	}

	return nil
}

func (c *Cmder) connectViperAndCobra(item configItem) error {
	if item.isHidden {
		return nil
	}
	
	if err := c.viper.BindPFlag(item.name, c.cobra.Flags().Lookup(toFlagName(item.name))); err != nil {
		return err
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

func (c *Cmder) run() func(cmd *cobra.Command, _ []string) error {
	return func(cmd *cobra.Command, _ []string) error {

		if err := c.viper.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
				return err
			}
		}

		if err := c.viper.Unmarshal(c.cfg); err != nil {
			return err
		}

		return nil 
	}
}
