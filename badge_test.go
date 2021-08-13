package badge

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2021 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"io/ioutil"
	"testing"

	. "pkg.re/essentialkaos/check.v1"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type BadgeSuite struct {
	generator *Generator
}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&BadgeSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //
func (s *BadgeSuite) SetUpTest(c *C) {
	var err error

	s.generator, err = NewGenerator("Verdana.ttf", 11)

	if err != nil {
		c.Fatal(err.Error())
	}
}

func (s *BadgeSuite) TestErrors(c *C) {
	_, err := NewGenerator("unknown.ttf", 0)

	c.Assert(err, NotNil)

	_, err = NewGenerator("badge.go", 0)

	c.Assert(err, NotNil)
}

func (s *BadgeSuite) TestPlastic(c *C) {
	srcBadge, err := ioutil.ReadFile("testdata/plastic.svg")

	if err != nil {
		c.Fatal(err.Error())
	}

	ourBadge := s.generator.GeneratePlastic("test1", "good", "ff69b4")

	c.Assert(ourBadge, DeepEquals, srcBadge)
}

func (s *BadgeSuite) TestFlat(c *C) {
	srcBadge, err := ioutil.ReadFile("testdata/flat.svg")

	if err != nil {
		c.Fatal(err.Error())
	}

	ourBadge := s.generator.GenerateFlat("test1", "good", "ff69b4")

	c.Assert(ourBadge, DeepEquals, srcBadge)
}

func (s *BadgeSuite) TestFlatSquare(c *C) {
	srcBadge, err := ioutil.ReadFile("testdata/flat-square.svg")

	if err != nil {
		c.Fatal(err.Error())
	}

	ourBadge := s.generator.GenerateFlatSquare("test1", "good", "ff69b4")

	c.Assert(ourBadge, DeepEquals, srcBadge)
}
