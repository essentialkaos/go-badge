// Package badge provides methods for generating SVG badges
package badge

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"

	"github.com/golang/freetype/truetype"

	"golang.org/x/image/font"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// VERSION is current package version
const VERSION = "1.3.0"

const (
	COLOR_BLUE        = "#007ec6"
	COLOR_BRIGHTGREEN = "#4c1"
	COLOR_GREEN       = "#97ca00"
	COLOR_GREY        = "#555"
	COLOR_LIGHTGREY   = "#9f9f9f"
	COLOR_ORANGE      = "#fe7d37"
	COLOR_RED         = "#e05d44"
	COLOR_YELLOW      = "#dfb317"
	COLOR_YELLOWGREEN = "#a4a61d"

	COLOR_SUCCESS       = "#4c1"
	COLOR_IMPORTANT     = "#fe7d37"
	COLOR_CRITICAL      = "#e05d44"
	COLOR_INFORMATIONAL = "#007ec6"
	COLOR_INACTIVE      = "#9f9f9f"
)

const (
	DEFAULT_OFFSET  = 9 // default font offset
	DEFAULT_SPACING = 0 // default letter spacing
)

// ////////////////////////////////////////////////////////////////////////////////// //

const _TEMPLATE_PLASTIC = `<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="{WIDTH}" height="18" role="img" aria-label="{LABEL}: {MESSAGE}"><title>{LABEL}: {MESSAGE}</title><linearGradient id="s" x2="0" y2="100%"><stop offset="0"  stop-color="#fff" stop-opacity=".7"/><stop offset=".1" stop-color="#aaa" stop-opacity=".1"/><stop offset=".9" stop-color="#000" stop-opacity=".3"/><stop offset="1"  stop-color="#000" stop-opacity=".5"/></linearGradient><clipPath id="r"><rect width="{WIDTH}" height="18" rx="4" fill="#fff"/></clipPath><g clip-path="url(#r)"><rect width="{LABEL_WIDTH}" height="18" fill="#555"/><rect x="{LABEL_WIDTH}" width="{MESSAGE_WIDTH}" height="18" fill="{COLOR}"/><rect width="{WIDTH}" height="18" fill="url(#s)"/></g><g fill="#fff" text-anchor="middle" font-family="{FONT},Geneva,DejaVu Sans,sans-serif" text-rendering="geometricPrecision" font-size="{FONT_SIZE}"><text aria-hidden="true" x="{LABEL_X}" y="140" fill="#010101" fill-opacity=".3" transform="scale(.1)" textLength="{LABEL_LENGTH}">{LABEL}</text><text x="{LABEL_X}" y="130" transform="scale(.1)" fill="#fff" textLength="{LABEL_LENGTH}">{LABEL}</text><text aria-hidden="true" x="{MESSAGE_X}" y="140" fill="{MESSAGE_SHADOW}" fill-opacity=".3" transform="scale(.1)" textLength="{MESSAGE_LENGTH}">{MESSAGE}</text><text x="{MESSAGE_X}" y="130" transform="scale(.1)" fill="{MESSAGE_COLOR}" textLength="{MESSAGE_LENGTH}">{MESSAGE}</text></g></svg>`

const _TEMPLATE_FLAT = `<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="{WIDTH}" height="20" role="img" aria-label="{LABEL}: {MESSAGE}"><title>{LABEL}: {MESSAGE}</title><linearGradient id="s" x2="0" y2="100%"><stop offset="0" stop-color="#bbb" stop-opacity=".1"/><stop offset="1" stop-opacity=".1"/></linearGradient><clipPath id="r"><rect width="{WIDTH}" height="20" rx="3" fill="#fff"/></clipPath><g clip-path="url(#r)"><rect width="{LABEL_WIDTH}" height="20" fill="#555"/><rect x="{LABEL_WIDTH}" width="{MESSAGE_WIDTH}" height="20" fill="{COLOR}"/><rect width="{WIDTH}" height="20" fill="url(#s)"/></g><g fill="#fff" text-anchor="middle" font-family="{FONT},Geneva,DejaVu Sans,sans-serif" text-rendering="geometricPrecision" font-size="{FONT_SIZE}"><text aria-hidden="true" x="{LABEL_X}" y="150" fill="#010101" fill-opacity=".3" transform="scale(.1)" textLength="{LABEL_LENGTH}">{LABEL}</text><text x="{LABEL_X}" y="140" transform="scale(.1)" fill="#fff" textLength="{LABEL_LENGTH}">{LABEL}</text><text aria-hidden="true" x="{MESSAGE_X}" y="150" fill="{MESSAGE_SHADOW}" fill-opacity=".3" transform="scale(.1)" textLength="{MESSAGE_LENGTH}">{MESSAGE}</text><text x="{MESSAGE_X}" y="140" transform="scale(.1)" fill="{MESSAGE_COLOR}" textLength="{MESSAGE_LENGTH}">{MESSAGE}</text></g></svg>`

const _TEMPLATE_FLAT_SQUARE = `<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="{WIDTH}" height="20" role="img" aria-label="{LABEL}: {MESSAGE}"><title>{LABEL}: {MESSAGE}</title><g shape-rendering="crispEdges"><rect width="{LABEL_WIDTH}" height="20" fill="#555"/><rect x="{LABEL_WIDTH}" width="{MESSAGE_WIDTH}" height="20" fill="{COLOR}"/></g><g fill="#fff" text-anchor="middle" font-family="{FONT},Geneva,DejaVu Sans,sans-serif" text-rendering="geometricPrecision" font-size="{FONT_SIZE}"><text x="{LABEL_X}" y="140" transform="scale(.1)" fill="#fff" textLength="{LABEL_LENGTH}">{LABEL}</text><text x="{MESSAGE_X}" y="140" transform="scale(.1)" fill="{MESSAGE_COLOR}" textLength="{MESSAGE_LENGTH}">{MESSAGE}</text></g></svg>`

const _TEMPLATE_FLAT_SIMPLE = `<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="{WIDTH}" height="20" role="img" aria-label="{MESSAGE}"><title>{MESSAGE}</title><linearGradient id="s" x2="0" y2="100%"><stop offset="0" stop-color="#bbb" stop-opacity=".1"/><stop offset="1" stop-opacity=".1"/></linearGradient><clipPath id="r"><rect width="{WIDTH}" height="20" rx="3" fill="#fff"/></clipPath><g clip-path="url(#r)"><rect width="0" height="20" fill="{COLOR}"/><rect x="0" width="{WIDTH}" height="20" fill="{COLOR}"/><rect width="{WIDTH}" height="20" fill="url(#s)"/></g><g fill="#fff" text-anchor="middle" font-family="{FONT},Geneva,DejaVu Sans,sans-serif" text-rendering="geometricPrecision" font-size="{FONT_SIZE}"><text aria-hidden="true" x="{MESSAGE_X}" y="150" fill="{MESSAGE_SHADOW}" fill-opacity=".3" transform="scale(.1)" textLength="{MESSAGE_LENGTH}">{MESSAGE}</text><text x="{MESSAGE_X}" y="140" transform="scale(.1)" fill="{MESSAGE_COLOR}" textLength="{MESSAGE_LENGTH}">{MESSAGE}</text></g></svg>`

const _TEMPLATE_PLASTIC_SIMPLE = `<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="{WIDTH}" height="18" role="img" aria-label="{MESSAGE}"><title>TESTTEST</title><linearGradient id="s" x2="0" y2="100%"><stop offset="0" stop-color="#fff" stop-opacity=".7"/><stop offset=".1" stop-color="#aaa" stop-opacity=".1"/><stop offset=".9" stop-color="#000" stop-opacity=".3"/><stop offset="1" stop-color="#000" stop-opacity=".5"/></linearGradient><clipPath id="r"><rect width="{WIDTH}" height="18" rx="4" fill="#fff"/></clipPath><g clip-path="url(#r)"><rect width="0" height="18" fill="{COLOR}"/><rect x="0" width="{WIDTH}" height="18" fill="{COLOR}"/><rect width="{WIDTH}" height="18" fill="url(#s)"/></g><g fill="#fff" text-anchor="middle" font-family="{FONT},Geneva,DejaVu Sans,sans-serif" text-rendering="geometricPrecision" font-size="{FONT_SIZE}"><text aria-hidden="true" x="{MESSAGE_X}" y="140" fill="{MESSAGE_SHADOW}" fill-opacity=".3" transform="scale(.1)" textLength="{MESSAGE_LENGTH}">{MESSAGE}</text><text x="{MESSAGE_X}" y="130" transform="scale(.1)" fill="{MESSAGE_COLOR}" textLength="{MESSAGE_LENGTH}">{MESSAGE}</text></g></svg>`

const _TEMPLATE_FLAT_SQUARE_SIMPLE = `<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="{WIDTH}" height="20" role="img" aria-label="{MESSAGE}"><title>{MESSAGE}</title><g shape-rendering="crispEdges"><rect width="0" height="20" fill="{COLOR}"/><rect x="0" width="{WIDTH}" height="20" fill="{COLOR}"/></g><g fill="#fff" text-anchor="middle" font-family="{FONT},Geneva,DejaVu Sans,sans-serif" text-rendering="geometricPrecision" font-size="{FONT_SIZE}"><text x="{MESSAGE_X}" y="140" transform="scale(.1)" fill="{MESSAGE_COLOR}" textLength="{MESSAGE_LENGTH}">{MESSAGE}</text></g></svg>`

// ////////////////////////////////////////////////////////////////////////////////// //

// Generator is badge generator
type Generator struct {
	Offset  int     // Text offset
	Spacing float64 // Letter spacing

	fontSize int
	fontName string
	drawer   *font.Drawer
}

// ////////////////////////////////////////////////////////////////////////////////// //

// NewGenerator creates new badge generator with given font
func NewGenerator(fontFile string, fontSize int) (*Generator, error) {
	fontData, err := ioutil.ReadFile(fontFile)

	if err != nil {
		return nil, err
	}

	fontTTF, err := truetype.Parse(fontData)

	if err != nil {
		return nil, err
	}

	return &Generator{
		Offset:  DEFAULT_OFFSET,
		Spacing: DEFAULT_SPACING,

		fontSize: fontSize,

		fontName: fontTTF.Name(truetype.NameIDFontFullName),
		drawer: &font.Drawer{
			Face: truetype.NewFace(fontTTF, &truetype.Options{
				Size:    float64(fontSize),
				DPI:     72,
				Hinting: font.HintingFull,
			}),
		},
	}, nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

// GeneratePlastic generates SVG badge in plastic style
func (g *Generator) GeneratePlastic(label, message, color string) []byte {
	if label == "" {
		return g.generateBadgeSimple(_TEMPLATE_PLASTIC_SIMPLE, message, color)
	}

	return g.generateBadge(_TEMPLATE_PLASTIC, label, message, color)
}

// GenerateFlat generates SVG badge in flat style
func (g *Generator) GenerateFlat(label, message, color string) []byte {
	if label == "" {
		return g.generateBadgeSimple(_TEMPLATE_FLAT_SIMPLE, message, color)
	}

	return g.generateBadge(_TEMPLATE_FLAT, label, message, color)
}

// GenerateFlatSquare generates SVG badge in flat-square style
func (g *Generator) GenerateFlatSquare(label, message, color string) []byte {
	if label == "" {
		return g.generateBadgeSimple(_TEMPLATE_FLAT_SQUARE_SIMPLE, message, color)
	}

	return g.generateBadge(_TEMPLATE_FLAT_SQUARE, label, message, color)
}

// GeneratePlasticSimple generates SVG simple badge in plastic style
func (g *Generator) GeneratePlasticSimple(message, color string) []byte {
	return g.generateBadgeSimple(_TEMPLATE_PLASTIC_SIMPLE, message, color)
}

// GenerateFlatSimple generates SVG simple badge in flat style
func (g *Generator) GenerateFlatSimple(message, color string) []byte {
	return g.generateBadgeSimple(_TEMPLATE_FLAT_SIMPLE, message, color)
}

// GenerateFlatSquareSimple generates SVG simple badge in flat-square style
func (g *Generator) GenerateFlatSquareSimple(message, color string) []byte {
	return g.generateBadgeSimple(_TEMPLATE_FLAT_SQUARE_SIMPLE, message, color)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// generateBadge generates badge with given template
func (g *Generator) generateBadge(template, label, message, color string) []byte {
	c := parseColor(color)

	gF := float64(g.Offset)
	lW := float64(g.drawer.MeasureString(label)>>6) + gF
	mW := float64(g.drawer.MeasureString(message)>>6) + gF
	fW := lW + mW
	lX := (lW/2 + 1) * 10
	mX := (lW + (mW / 2) - 1) * 10
	lL := (lW - gF) * (10.0 + g.Spacing - 0.5)
	mL := (mW - gF) * (10.0 + g.Spacing - 0.5)
	fS := g.fontSize * 10

	mC, mS := getMessageColors(c)

	badge := strings.ReplaceAll(template, "{LABEL}", label)
	badge = strings.ReplaceAll(badge, "{MESSAGE}", message)
	badge = strings.ReplaceAll(badge, "{COLOR}", formatColor(c))
	badge = strings.ReplaceAll(badge, "{WIDTH}", formatFloat(fW))
	badge = strings.ReplaceAll(badge, "{LABEL_WIDTH}", formatFloat(lW))
	badge = strings.ReplaceAll(badge, "{MESSAGE_WIDTH}", formatFloat(mW))
	badge = strings.ReplaceAll(badge, "{LABEL_X}", formatFloat(lX))
	badge = strings.ReplaceAll(badge, "{MESSAGE_X}", formatFloat(mX))
	badge = strings.ReplaceAll(badge, "{LABEL_LENGTH}", formatFloat(lL))
	badge = strings.ReplaceAll(badge, "{MESSAGE_LENGTH}", formatFloat(mL))
	badge = strings.ReplaceAll(badge, "{MESSAGE_COLOR}", mC)
	badge = strings.ReplaceAll(badge, "{MESSAGE_SHADOW}", mS)
	badge = strings.ReplaceAll(badge, "{FONT}", g.fontName)
	badge = strings.ReplaceAll(badge, "{FONT_SIZE}", strconv.Itoa(fS))

	return []byte(badge)
}

// generateBadgeSimple generates badge with given template
func (g *Generator) generateBadgeSimple(template, message, color string) []byte {
	c := parseColor(color)

	gF := float64(g.Offset)
	fW := float64(g.drawer.MeasureString(message)>>6) + gF
	mX := (fW / 2) * 10
	mL := (fW - gF) * (10.0 + g.Spacing)
	fS := g.fontSize * 10

	mC, mS := getMessageColors(c)

	badge := strings.ReplaceAll(template, "{MESSAGE}", message)
	badge = strings.ReplaceAll(badge, "{COLOR}", formatColor(c))
	badge = strings.ReplaceAll(badge, "{WIDTH}", formatFloat(fW))
	badge = strings.ReplaceAll(badge, "{MESSAGE_X}", formatFloat(mX))
	badge = strings.ReplaceAll(badge, "{MESSAGE_LENGTH}", formatFloat(mL))
	badge = strings.ReplaceAll(badge, "{MESSAGE_COLOR}", mC)
	badge = strings.ReplaceAll(badge, "{MESSAGE_SHADOW}", mS)
	badge = strings.ReplaceAll(badge, "{FONT}", g.fontName)
	badge = strings.ReplaceAll(badge, "{FONT_SIZE}", strconv.Itoa(fS))

	return []byte(badge)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// parseColor parses hex color
func parseColor(c string) int64 {
	if strings.HasPrefix(c, "#") {
		c = strings.TrimLeft(c, "#")
	}

	// Shorthand hex color
	if len(c) == 3 {
		c = strings.Repeat(string(c[0]), 2) +
			strings.Repeat(string(c[1]), 2) +
			strings.Repeat(string(c[2]), 2)
	}

	i, _ := strconv.ParseInt(c, 16, 32)

	return i
}

// formatColor formats color
func formatColor(c int64) string {
	k := fmt.Sprintf("%06x", c)

	if k[0] == k[1] && k[2] == k[3] && k[4] == k[5] {
		k = k[0:1] + k[2:3] + k[4:5]
	}

	return "#" + k
}

// formatFloat formats float values
func formatFloat(v float64) string {
	return strconv.FormatFloat(v, 'f', 0, 64)
}

// getMessageColors returns message text and shadow colors based on color of badge
func getMessageColors(c int64) (string, string) {
	if c == 0 || calcLuminance(c) < 0.65 {
		return "#fff", "#010101"
	}

	return "#333", "#ccc"
}

// calcLuminance calculates relative luminance
func calcLuminance(color int64) float64 {
	r := calcLumColor(float64(color>>16&0xFF) / 255)
	g := calcLumColor(float64(color>>8&0xFF) / 255)
	b := calcLumColor(float64(color&0xFF) / 255)

	return 0.2126*r + 0.7152*g + 0.0722*b
}

// calcLumColor calculates luminance for one color
func calcLumColor(c float64) float64 {
	if c <= 0.03928 {
		return c / 12.92
	}

	return math.Pow(((c + 0.055) / 1.055), 2.4)
}
