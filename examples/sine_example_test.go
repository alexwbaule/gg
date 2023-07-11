// Copyright ©2022 The gg Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package examples_test

import (
	"log"
	"math"
	"testing"

	"github.com/alexwbaule/gg"
)

func ExampleSine() {
	const W = 1200
	const H = 60
	dc := gg.NewContext(W, H)
	// dc.SetHexColor("#FFFFFF")
	// dc.Clear()
	dc.ScaleAbout(0.95, 0.75, W/2, H/2)
	for i := 0; i < W; i++ {
		a := float64(i) * 2 * math.Pi / W * 8
		x := float64(i)
		y := (math.Sin(a) + 1) / 2 * H
		dc.LineTo(x, y)
	}
	dc.ClosePath()
	dc.SetHexColor("#3E606F")
	dc.FillPreserve()
	dc.SetHexColor("#19344180")
	dc.SetLineWidth(8)
	dc.Stroke()

	err := dc.SavePNG("testdata/sine.png")
	if err != nil {
		log.Fatalf("could not save to file: %+v", err)
	}
}

func TestSine(t *testing.T) {
	chkimg(ExampleSine, t, "sine.png")
}
