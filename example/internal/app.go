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

package internal

import "fmt"

type AppConfig struct {
    Directory string `desc:"Directory to browse" default:"."`
    Server ServerConfig // support multi-level config          
}

type ServerConfig struct {
    Url string `desc:"Server url" default:"localhost"`
    Port int `desc:"Username" default:"8080"`
}

type app struct {
	Config AppConfig
}

func NewApp(cfg AppConfig) *app {
	return &app{
		Config: cfg,
	}
}

func (a *app) Run() {
	fmt.Printf("Directory: %s\nUrl: %s:%d\n", a.Config.Directory, a.Config.Server.Url, a.Config.Server.Port)
}