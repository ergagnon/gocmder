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
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Cmder struct {
	cfg any
	cobra *cobra.Command
	viper *viper.Viper
	longDesc string
	shortDesc string
	version string
}

type OnFinalizeFunc func(cfg any) error

func NewCmder(cfg any, onFinalize OnFinalizeFunc, opts ...CmderOption) *Cmder {
	c := &Cmder{
		cfg: cfg,		
		viper: viper.New(),
	}

	for _, opt := range opts {
		opt(c)
	}
	
	c.cobra = &cobra.Command{
		Short: c.shortDesc,
		Long: c.longDesc,
		Version: c.version,
		RunE: c.run(),
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

func (c *Cmder) run() func (cmd *cobra.Command, _ []string) error {
	return func (cmd *cobra.Command, _ []string) error {
		return cmd.Help()
	}
}