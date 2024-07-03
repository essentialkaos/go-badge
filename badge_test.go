package badge

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"os"
	"testing"

	. "github.com/essentialkaos/check"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type BadgeSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

type ErrReader struct{}

func (e *ErrReader) Read(b []byte) (int, error) {
	return 0, fmt.Errorf("ERROR")
}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&BadgeSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //
func (s *BadgeSuite) SetUpTest(c *C) {
	fd, err := os.Open("Verdana.ttf")

	if err != nil {
		c.Fatal(err.Error())
	}

	defer fd.Close()
}

func (s *BadgeSuite) TestConstructors(c *C) {
	fd, err := os.Open("Verdana.ttf")

	if err != nil {
		c.Fatal(err.Error())
	}

	defer fd.Close()

	_, err = NewGenerator("Verdana.ttf", 11)
	c.Assert(err, IsNil)
}

func (s *BadgeSuite) TestErrors(c *C) {
	_, err := NewGenerator("unknown.ttf", 0)
	c.Assert(err, NotNil)

	_, err = NewGenerator("badge.go", 0)
	c.Assert(err, NotNil)

	errReader := &ErrReader{}

	_, err = NewGeneratorFromReader(errReader, 0)
	c.Assert(err, NotNil)

	_, err = NewGeneratorFromBytes(nil, 0)
	c.Assert(err, NotNil)
}

func (s *BadgeSuite) TestPlastic(c *C) {
	srcBadge, err := os.ReadFile("testdata/plastic.svg")

	if err != nil {
		c.Fatal(err.Error())
	}

	gen, err := NewGenerator("Verdana.ttf", 11)
	c.Assert(err, IsNil)

	ourBadge := gen.GeneratePlastic("label", "message", COLOR_RED)

	c.Assert(string(ourBadge), Equals, string(srcBadge))
}

func (s *BadgeSuite) TestFlat(c *C) {
	srcBadge, err := os.ReadFile("testdata/flat.svg")

	if err != nil {
		c.Fatal(err.Error())
	}

	gen, err := NewGenerator("Verdana.ttf", 11)
	c.Assert(err, IsNil)

	ourBadge := gen.GenerateFlat("label", "message", COLOR_RED)

	c.Assert(string(ourBadge), Equals, string(srcBadge))
}

func (s *BadgeSuite) TestFlatSquare(c *C) {
	srcBadge, err := os.ReadFile("testdata/square.svg")

	if err != nil {
		c.Fatal(err.Error())
	}

	gen, err := NewGenerator("Verdana.ttf", 11)
	c.Assert(err, IsNil)

	ourBadge := gen.GenerateFlatSquare("label", "message", COLOR_RED)

	c.Assert(string(ourBadge), Equals, string(srcBadge))
}

func (s *BadgeSuite) TestPlasticSimple(c *C) {
	srcBadge, err := os.ReadFile("testdata/plastic_simple.svg")

	if err != nil {
		c.Fatal(err.Error())
	}

	gen, err := NewGenerator("Verdana.ttf", 11)
	c.Assert(err, IsNil)

	ourBadge := gen.GeneratePlastic("", "message", COLOR_RED)
	c.Assert(string(ourBadge), Equals, string(srcBadge))

	ourBadge = gen.GeneratePlasticSimple("message", COLOR_RED)
	c.Assert(string(ourBadge), Equals, string(srcBadge))
}

func (s *BadgeSuite) TestFlatSimple(c *C) {
	srcBadge, err := os.ReadFile("testdata/flat_simple.svg")

	if err != nil {
		c.Fatal(err.Error())
	}

	gen, err := NewGenerator("Verdana.ttf", 11)
	c.Assert(err, IsNil)

	ourBadge := gen.GenerateFlat("", "message", COLOR_RED)
	c.Assert(string(ourBadge), Equals, string(srcBadge))

	ourBadge = gen.GenerateFlatSimple("message", COLOR_RED)
	c.Assert(string(ourBadge), Equals, string(srcBadge))
}

func (s *BadgeSuite) TestFlatSquareSimple(c *C) {
	srcBadge, err := os.ReadFile("testdata/square_simple.svg")

	if err != nil {
		c.Fatal(err.Error())
	}

	gen, err := NewGenerator("Verdana.ttf", 11)
	c.Assert(err, IsNil)

	ourBadge := gen.GenerateFlatSquare("", "message", COLOR_RED)
	c.Assert(string(ourBadge), Equals, string(srcBadge))

	ourBadge = gen.GenerateFlatSquareSimple("message", COLOR_RED)
	c.Assert(string(ourBadge), Equals, string(srcBadge))
}

func (s *BadgeSuite) TestBlackAndWhite(c *C) {
	bb, err := os.ReadFile("testdata/black.svg")

	if err != nil {
		c.Fatal(err.Error())
	}

	wb, err := os.ReadFile("testdata/white.svg")

	if err != nil {
		c.Fatal(err.Error())
	}

	gen, err := NewGenerator("Verdana.ttf", 11)
	c.Assert(err, IsNil)

	ourBadge := gen.GenerateFlatSimple("message", "#FFFFFF")
	c.Assert(string(ourBadge), Equals, string(wb))

	ourBadge = gen.GenerateFlatSimple("message", "#000000")
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
