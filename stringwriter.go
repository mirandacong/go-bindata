// Copyright 2018 The go-bindata Authors. All rights reserved.
// Use of this source code is governed by a CC0 1.0 Universal (CC0 1.0)
// Public Domain Dedication license that can be found in the LICENSE file.

package bindata

import (
	"io"
)

const lowerHex = "0123456789abcdef"

//
// StringWriter define a writer to write content of file.
//
type StringWriter struct {
	io.Writer
	c int
}

func (w *StringWriter) Write(p []byte) (n int, err error) {
	if len(p) == 0 {
		return
	}

	buf := []byte(`\x00`)
	var b byte

	for n, b = range p {
		buf[2] = lowerHex[b/16]
		buf[3] = lowerHex[b%16]

		_, err = w.Writer.Write(buf)
		if err != nil {
			return
		}

		w.c++

		// 28 fits nicely with tab width at 4 and a 120 char line limit
		if w.c%28 == 0 {
			_, err = w.Writer.Write([]byte("\" +\n\t\""))
			if err != nil {
				return
			}
		}
	}

	n++

	return
}
