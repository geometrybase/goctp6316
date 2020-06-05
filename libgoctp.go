// +build linux,cgo

// Copyright 2012 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goctp

/*
#cgo linux LDFLAGS: -fPIC -L${SRCDIR}/20190305_tradeapi_linux64 -Wl,-rpath,${SRCDIR}/api/v6.3.16_20190305_tradeapi_linux64 -lthostmduserapi6316 -lthosttraderapi6316 -lstdc++
#cgo linux CPPFLAGS: -fPIC -I${SRCDIR}/20190305_tradeapi_linux64
*/
import "C"
