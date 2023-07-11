// Copyright ©2022 The gg Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package examples_test

import (
	"image/color"
	"log"
	"testing"

	"git.sr.ht/~sbinet/gg"
)

func ExampleGradientLinear() {
	dc := gg.NewContext(500, 400)

	grad := gg.NewLinearGradient(20, 320, 400, 20)
	grad.AddColorStop(0, color.RGBA{0, 255, 0, 255})
	grad.AddColorStop(1, color.RGBA{0, 0, 255, 255})
	grad.AddColorStop(0.5, color.RGBA{255, 0, 0, 255})

	dc.SetColor(color.White)
	dc.DrawRectangle(20, 20, 400-20, 300)
	dc.Stroke()

	dc.SetStrokeStyle(grad)
	dc.SetLineWidth(4)
	dc.MoveTo(10, 10)
	dc.LineTo(410, 10)
	dc.LineTo(410, 100)
	dc.LineTo(10, 100)
	dc.ClosePath()
	dc.Stroke()

	dc.SetFillStyle(grad)
	dc.MoveTo(10, 120)
	dc.LineTo(410, 120)
	dc.LineTo(410, 300)
	dc.LineTo(10, 300)
	dc.ClosePath()
	dc.Fill()

	err := dc.SavePNG("testdata/gradient-linear.png")
	if err != nil {
		log.Fatalf("could not save to file: %+v", err)
	}

}

func TestGradientLinear(t *testing.T) {
	chkimg(ExampleGradientLinear, t, "gradient-linear.png")
}
