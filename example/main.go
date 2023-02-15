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

package main

import (
	"os"

	"github.com/ergagnon/gocmder"
	"github.com/ergagnon/gocmder/example/internal"
)

func main() {
	cli, err := gocmder.NewCmder(internal.AppConfig{}, func(cfg any) {
		app := internal.NewApp(cfg.(internal.AppConfig))
		app.Run()		
	})

	if err != nil {
		os.Exit(1)
	}

	if err := cli.Execute(); err != nil {
		os.Exit(1)
	}
}