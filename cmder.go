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

func (c *Cmder) run() func(cmd *cobra.Command, _ []string) error {
	return func(cmd *cobra.Command, _ []string) error {
		return cmd.Help()
	}
}
