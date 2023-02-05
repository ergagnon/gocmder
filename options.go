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

type CmderOption func(*Cmder)

// WithShortDesc sets the short description for the command.
// This is used by the "help" command.
func WithShortDesc(short string) CmderOption {
	return func(c *Cmder) {
		c.shortDesc = short
	}
}

// WithLongDesc sets the long description for the command.
// This is used by the "help" command.
func WithLongDesc(long string) CmderOption {
	return func(c *Cmder) {
		c.longDesc = long
	}
}

// WithVersion sets the version string for the command.
// This is used by the "version" command.
func WithVersion(version string) CmderOption {
	return func(c *Cmder) {
		c.version = version
	}
}

// WithPrefix sets the prefix for environment variables.
// For example, if the prefix is "APP", then the environment variable
// "APP_DEBUG" will be used to set the value of the flag "--debug".
func WithPrefix(prefix string) CmderOption {
	return func(c *Cmder) {
		c.viper.SetEnvPrefix(prefix)
	}
}
