// Copyright (c) 2020 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

package atomic

import (
	"testing"
	"unsafe"
)

func TestPointer(t *testing.T) {
	var v string
	var vp = unsafe.Pointer(&v)
	var addr = NewPointer(vp)
	var v1 = "Hello World"
	var vp1 = unsafe.Pointer(&v1)
	addr.Store(vp1)
	if addr.Load() != vp1 {
		t.Log(addr.Load())
	}
	var v2 = "Foo"
	var vp2 = unsafe.Pointer(&v2)
	old := addr.Swap(vp2)
	if old != vp1 {
		t.Log(addr.Load())
	}
	var v3 = "Bar"
	var vp3 = unsafe.Pointer(&v3)
	if !addr.CompareAndSwap(vp2, vp3) {
		t.Log(addr.Load())
	}
	if addr.CompareAndSwap(vp2, vp3) {
		t.Log(addr.Load())
	}
}

func BenchmarkSwapPointer(b *testing.B) {
	var v string
	var vp = unsafe.Pointer(&v)
	var addr = NewPointer(vp)
	var v1 = "Hello World"
	var vp1 = unsafe.Pointer(&v1)
	for i := 0; i < b.N; i++ {
		addr.Swap(vp1)
	}
}

func BenchmarkCompareAndSwapPointer(b *testing.B) {
	var v string
	var vp = unsafe.Pointer(&v)
	var addr = NewPointer(vp)
	var v1 = "Hello World"
	var vp1 = unsafe.Pointer(&v1)
	for i := 0; i < b.N; i++ {
		addr.CompareAndSwap(vp, vp1)
	}
}

func BenchmarkStorePointer(b *testing.B) {
	var v string
	var vp = unsafe.Pointer(&v)
	var addr = NewPointer(vp)
	var v1 = "Hello World"
	var vp1 = unsafe.Pointer(&v1)
	for i := 0; i < b.N; i++ {
		addr.Store(vp1)
	}
}

func BenchmarkLoadPointer(b *testing.B) {
	var v string
	var vp = unsafe.Pointer(&v)
	var addr = NewPointer(vp)
	for i := 0; i < b.N; i++ {
		addr.Load()
	}
}
