package app

import (
	"github.com/timeredbull/commandmocker"
	"github.com/timeredbull/tsuru/api/bind"
	"github.com/timeredbull/tsuru/repository"
	. "launchpad.net/gocheck"
)

func (s *S) TestCommand(c *C) {
	var err error
	s.tmpdir, err = commandmocker.Add("juju", "Linux")
	c.Assert(err, IsNil)
	defer commandmocker.Remove(s.tmpdir)
	u := Unit{Type: "django", Name: "myUnit", Machine: 1}
	output, err := u.Command("uname")
	c.Assert(err, IsNil)
	c.Assert(string(output), Equals, "Linux")
}

func (s *S) TestCommandShouldAcceptMultipleParams(c *C) {
	dir, err := commandmocker.Add("juju", "$*")
	c.Assert(err, IsNil)
	defer commandmocker.Remove(dir)
	u := Unit{Type: "django", Name: "myUnit", Machine: 1}
	out, err := u.Command("uname", "-a")
	c.Assert(string(out), Matches, ".* uname -a")
}

func (s *S) TestExecuteHook(c *C) {
	appUnit := Unit{Type: "django", Name: "myUnit"}
	_, err := appUnit.ExecuteHook("requirements")
	c.Assert(err, IsNil)
}

func (s *S) TestUnitShouldBeARepositoryUnit(c *C) {
	var unit repository.Unit
	c.Assert(&Unit{}, Implements, &unit)
}

func (s *S) TestUnitShouldBeABinderUnit(c *C) {
	var unit bind.Unit
	c.Assert(&Unit{}, Implements, &unit)
}
