// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build generate

package registry

//go:generate go run github.com/mjgrzybek/syswindows/mkwinsyscall -output zsyscall_windows.go syscall.go
