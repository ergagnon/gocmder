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
