// Package badge provides methods for generating SVG badges
package badge

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2021 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/golang/freetype/truetype"

	"golang.org/x/image/font"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const _TEMPLATE_PLASTIC = `<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="{WIDTH}" height="18" role="img" aria-label="{LABEL}: {MESSAGE}"><title>{LABEL}: {MESSAGE}</title><linearGradient id="s" x2="0" y2="100%"><stop offset="0"  stop-color="#fff" stop-opacity=".7"/><stop offset=".1" stop-color="#aaa" stop-opacity=".1"/><stop offset=".9" stop-color="#000" stop-opacity=".3"/><stop offset="1"  stop-color="#000" stop-opacity=".5"/></linearGradient><clipPath id="r"><rect width="{WIDTH}" height="18" rx="4" fill="#fff"/></clipPath><g clip-path="url(#r)"><rect width="{LABEL_WIDTH}" height="18" fill="#555"/><rect x="{LABEL_WIDTH}" width="{MESSAGE_WIDTH}" height="18" fill="{COLOR}"/><rect width="{WIDTH}" height="18" fill="url(#s)"/></g><g fill="#fff" text-anchor="middle" font-family="Verdana,Geneva,DejaVu Sans,sans-serif" text-rendering="geometricPrecision" font-size="110"><text aria-hidden="true" x="{LABEL_X}" y="140" fill="#010101" fill-opacity=".3" transform="scale(.1)" textLength="{LABEL_LENGTH}">{LABEL}</text><text x="{LABEL_X}" y="130" transform="scale(.1)" fill="#fff" textLength="{LABEL_LENGTH}">{LABEL}</text><text aria-hidden="true" x="{MESSAGE_X}" y="140" fill="#010101" fill-opacity=".3" transform="scale(.1)" textLength="{MESSAGE_LENGTH}">{MESSAGE}</text><text x="{MESSAGE_X}" y="130" transform="scale(.1)" fill="#fff" textLength="{MESSAGE_LENGTH}">{MESSAGE}</text></g></svg>`

const _TEMPLATE_FLAT = `<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="{WIDTH}" height="20" role="img" aria-label="{LABEL}: {MESSAGE}"><title>{LABEL}: {MESSAGE}</title><linearGradient id="s" x2="0" y2="100%"><stop offset="0" stop-color="#bbb" stop-opacity=".1"/><stop offset="1" stop-opacity=".1"/></linearGradient><clipPath id="r"><rect width="{WIDTH}" height="20" rx="3" fill="#fff"/></clipPath><g clip-path="url(#r)"><rect width="{LABEL_WIDTH}" height="20" fill="#555"/><rect x="{LABEL_WIDTH}" width="{MESSAGE_WIDTH}" height="20" fill="{COLOR}"/><rect width="{WIDTH}" height="20" fill="url(#s)"/></g><g fill="#fff" text-anchor="middle" font-family="Verdana,Geneva,DejaVu Sans,sans-serif" text-rendering="geometricPrecision" font-size="110"><text aria-hidden="true" x="{LABEL_X}" y="150" fill="#010101" fill-opacity=".3" transform="scale(.1)" textLength="{LABEL_LENGTH}">{LABEL}</text><text x="{LABEL_X}" y="140" transform="scale(.1)" fill="#fff" textLength="{LABEL_LENGTH}">{LABEL}</text><text aria-hidden="true" x="{MESSAGE_X}" y="150" fill="#010101" fill-opacity=".3" transform="scale(.1)" textLength="{MESSAGE_LENGTH}">{MESSAGE}</text><text x="{MESSAGE_X}" y="140" transform="scale(.1)" fill="#fff" textLength="{MESSAGE_LENGTH}">{MESSAGE}</text></g></svg>`

const _TEMPLATE_FLAT_SQUARE = `<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="{WIDTH}" height="20" role="img" aria-label="{LABEL}: {MESSAGE}"><title>{LABEL}: {MESSAGE}</title><g shape-rendering="crispEdges"><rect width="{LABEL_WIDTH}" height="20" fill="#555"/><rect x="{LABEL_WIDTH}" width="{MESSAGE_WIDTH}" height="20" fill="{COLOR}"/></g><g fill="#fff" text-anchor="middle" font-family="Verdana,Geneva,DejaVu Sans,sans-serif" text-rendering="geometricPrecision" font-size="110"><text x="{LABEL_X}" y="140" transform="scale(.1)" fill="#fff" textLength="{LABEL_LENGTH}">{LABEL}</text><text x="{MESSAGE_X}" y="140" transform="scale(.1)" fill="#fff" textLength="{MESSAGE_LENGTH}">{MESSAGE}</text></g></svg>`

// ////////////////////////////////////////////////////////////////////////////////// //

type Generator struct {
	Offset int

	drawer *font.Drawer
}

// ////////////////////////////////////////////////////////////////////////////////// //

func NewGenerator(fontFile string) (*Generator, error) {
	fontData, err := ioutil.ReadFile(fontFile)

	if err != nil {
		return nil, err
	}

	fontTTF, err := truetype.Parse(fontData)

	if err != nil {
		return nil, err
	}

	return &Generator{
		Offset: 9,
		drawer: &font.Drawer{
			Face: truetype.NewFace(fontTTF, &truetype.Options{
				Size:    11,
				DPI:     72,
				Hinting: font.HintingFull,
			}),
		},
	}, nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

// GeneratePlastic generates SVG badge in plastic style
func (g *Generator) GeneratePlastic(label, message, color string) string {
	return g.generateBadge(_TEMPLATE_PLASTIC, label, message, color)
}

// GenerateFlat generates SVG badge in flat style
func (g *Generator) GenerateFlat(label, message, color string) string {
	return g.generateBadge(_TEMPLATE_FLAT, label, message, color)
}

// GenerateFlatSquare generates SVG badge in flat-square style
func (g *Generator) GenerateFlatSquare(label, message, color string) string {
	return g.generateBadge(_TEMPLATE_FLAT_SQUARE, label, message, color)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// generateBadge generates badge with given template
func (g *Generator) generateBadge(template, label, message, color string) string {
	if !strings.HasPrefix(color, "#") {
		color = "#" + color
	}

	lW := int(g.drawer.MeasureString(label)>>6) + g.Offset
	mW := int(g.drawer.MeasureString(message)>>6) + g.Offset
	fW := lW + mW
	lX := ((lW/2 + 1) * 10) + 5
	mX := ((lW + (mW / 2) - 1) * 10) + 5
	lL := (lW - 10) * 10
	mL := (mW - 10) * 10

	badge := strings.ReplaceAll(template, "{LABEL}", label)
	badge = strings.ReplaceAll(badge, "{MESSAGE}", message)
	badge = strings.ReplaceAll(badge, "{COLOR}", color)
	badge = strings.ReplaceAll(badge, "{WIDTH}", strconv.Itoa(fW))
	badge = strings.ReplaceAll(badge, "{LABEL_WIDTH}", strconv.Itoa(lW))
	badge = strings.ReplaceAll(badge, "{MESSAGE_WIDTH}", strconv.Itoa(mW))
	badge = strings.ReplaceAll(badge, "{LABEL_X}", strconv.Itoa(lX))
	badge = strings.ReplaceAll(badge, "{MESSAGE_X}", strconv.Itoa(mX))
	badge = strings.ReplaceAll(badge, "{LABEL_LENGTH}", strconv.Itoa(lL))
	badge = strings.ReplaceAll(badge, "{MESSAGE_LENGTH}", strconv.Itoa(mL))

	return badge
}
