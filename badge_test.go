package badge

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"

	. "github.com/essentialkaos/check"
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

func (s *BadgeSuite) NewGeneratorFromReader(c *C) {
	b, err := os.ReadFile("Verdana.ttf")
	if err != nil {
		c.Fatal(err.Error())
	}

	_, err = NewGeneratorFromReader(bytes.NewReader(b), 0)

	if err != nil {
		c.Fatal(err.Error())
	}
}

func (s *BadgeSuite) TestPlastic(c *C) {
	srcBadge, err := ioutil.ReadFile("testdata/plastic.svg")

	if err != nil {
		c.Fatal(err.Error())
	}

	ourBadge := s.generator.GeneratePlastic("label", "message", COLOR_RED)

	c.Assert(string(ourBadge), Equals, string(srcBadge))
}

func (s *BadgeSuite) TestFlat(c *C) {
	srcBadge, err := ioutil.ReadFile("testdata/flat.svg")

	if err != nil {
		c.Fatal(err.Error())
	}

	ourBadge := s.generator.GenerateFlat("label", "message", COLOR_RED)

	c.Assert(string(ourBadge), Equals, string(srcBadge))
}

func (s *BadgeSuite) TestFlatSquare(c *C) {
	srcBadge, err := ioutil.ReadFile("testdata/square.svg")

	if err != nil {
		c.Fatal(err.Error())
	}

	ourBadge := s.generator.GenerateFlatSquare("label", "message", COLOR_RED)

	c.Assert(string(ourBadge), Equals, string(srcBadge))
}

func (s *BadgeSuite) TestPlasticSimple(c *C) {
	srcBadge, err := ioutil.ReadFile("testdata/plastic_simple.svg")

	if err != nil {
		c.Fatal(err.Error())
	}

	ourBadge := s.generator.GeneratePlastic("", "message", COLOR_RED)
	c.Assert(string(ourBadge), Equals, string(srcBadge))

	ourBadge = s.generator.GeneratePlasticSimple("message", COLOR_RED)
	c.Assert(string(ourBadge), Equals, string(srcBadge))
}

func (s *BadgeSuite) TestFlatSimple(c *C) {
	srcBadge, err := ioutil.ReadFile("testdata/flat_simple.svg")

	if err != nil {
		c.Fatal(err.Error())
	}

	ourBadge := s.generator.GenerateFlat("", "message", COLOR_RED)
	c.Assert(string(ourBadge), Equals, string(srcBadge))

	ourBadge = s.generator.GenerateFlatSimple("message", COLOR_RED)
	c.Assert(string(ourBadge), Equals, string(srcBadge))
}

func (s *BadgeSuite) TestFlatSquareSimple(c *C) {
	srcBadge, err := ioutil.ReadFile("testdata/square_simple.svg")

	if err != nil {
		c.Fatal(err.Error())
	}

	ourBadge := s.generator.GenerateFlatSquare("", "message", COLOR_RED)
	c.Assert(string(ourBadge), Equals, string(srcBadge))

	ourBadge = s.generator.GenerateFlatSquareSimple("message", COLOR_RED)
	c.Assert(string(ourBadge), Equals, string(srcBadge))
}

func (s *BadgeSuite) TestBlackAndWhite(c *C) {
	bb, err := ioutil.ReadFile("testdata/black.svg")

	if err != nil {
		c.Fatal(err.Error())
	}

	wb, err := ioutil.ReadFile("testdata/white.svg")

	if err != nil {
		c.Fatal(err.Error())
	}

	ourBadge := s.generator.GenerateFlatSimple("message", "#FFFFFF")
	c.Assert(string(ourBadge), Equals, string(wb))

	ourBadge = s.generator.GenerateFlatSimple("message", "#000000")
	c.Assert(string(ourBadge), Equals, string(bb))
}

func (s *BadgeSuite) TestAux(c *C) {
	c.Assert(parseColor("000000"), Equals, int64(0x000000))
	c.Assert(parseColor("#000000"), Equals, int64(0x000000))
	c.Assert(parseColor("#4c1"), Equals, int64(0x44cc11))

	c.Assert(formatColor(0x000000), Equals, "#000")
	c.Assert(formatColor(0xFCA1B4), Equals, "#fca1b4")

	c.Assert(calcLumColor(0.7), Equals, 0.4479884124418833)
	c.Assert(calcLumColor(0.01), Equals, 0.0007739938080495357)
}
