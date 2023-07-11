// Copyright ©2022 The gg Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package examples_test

import (
	"log"
	"testing"

	"git.sr.ht/~sbinet/gg"
)

func ExampleInvertMask() {
	dc := gg.NewContext(1024, 1024)
	dc.DrawCircle(512, 512, 384)
	dc.Clip()
	dc.InvertMask()
	dc.DrawRectangle(0, 0, 1024, 1024)
	dc.SetRGB(0, 0, 0)
	dc.Fill()

	err := dc.SavePNG("testdata/invert-mask.png")
	if err != nil {
		log.Fatalf("could not save to file: %+v", err)
	}
}

func TestInvertMask(t *testing.T) {
	chkimg(ExampleInvertMask, t, "invert-mask.png")
}
