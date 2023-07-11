// Copyright ©2022 The gg Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package examples_test

import (
	"image/color"
	"log"
	"testing"

	"github.com/alexwbaule/gg"
	"github.com/go-fonts/anton/antonregular"
)

const (
	W = 1024
	H = 512
)

func ExampleGradientText() {
	dc := gg.NewContext(W, H)

	// draw text
	dc.SetRGB(0, 0, 0)

	err := dc.LoadFontFaceFromBytes(antonregular.TTF, 128)
	if err != nil {
		log.Fatalf("could not load bold font: %+v", err)
	}
	dc.DrawStringAnchored("Gradient Text", W/2, H/2, 0.5, 0.5)

	// get the context as an alpha mask
	mask := dc.AsMask()

	// clear the context
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	// set a gradient
	g := gg.NewLinearGradient(0, 0, W, H)
	g.AddColorStop(0, color.RGBA{255, 0, 0, 255})
	g.AddColorStop(1, color.RGBA{0, 0, 255, 255})
	dc.SetFillStyle(g)

	// using the mask, fill the context with the gradient
	dc.SetMask(mask)
	dc.DrawRectangle(0, 0, W, H)
	dc.Fill()

	err = dc.SavePNG("testdata/gradient-text.png")
	if err != nil {
		log.Fatalf("could not save to file: %+v", err)
	}
}

func TestGradientText(t *testing.T) {
	chkimg(ExampleGradientText, t, "gradient-text.png")
}
