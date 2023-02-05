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
	"testing"

	"github.com/stretchr/testify/suite"
)

type optionsTestSuite struct {
	suite.Suite
}

func (s *optionsTestSuite) TestWithVersion() {
	cmd := Cmder{}
	WithVersion("v1.0.0")(&cmd)

	s.Equal("v1.0.0", cmd.version)
}

func (s *optionsTestSuite) TestWithLongDescription() {
	cmd := Cmder{}
	WithLongDesc("long description")(&cmd)

	s.Equal("long description", cmd.longDesc)
}

func (s *optionsTestSuite) TestWithShortDescription() {
	cmd := Cmder{}
	WithShortDesc("short description")(&cmd)

	s.Equal("short description", cmd.shortDesc)
}

func TestOptionsTestSuite(t *testing.T) {
	suite.Run(t, new(optionsTestSuite))
}
