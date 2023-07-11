// Copyright ©2023 The gg Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package examples_test

import (
	"os"
	"path/filepath"
	"testing"

	"gonum.org/v1/plot/cmpimg"
)

func chkimg(fn func(), t *testing.T, fname string) {
	t.Helper()
	cmpimg.CheckPlot(fn, t, fname)
	if !t.Failed() {
		_ = os.Remove(filepath.Join("testdata", fname))
	}
}
