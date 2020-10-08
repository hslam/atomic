// Copyright (c) 2020 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

package atomic

import (
	"testing"
	"unsafe"
)

func TestAtomicInt32(t *testing.T) {
	var v int32
	StoreInt32(&v, 1)
	if LoadInt32(&v) != 1 {
		t.Log(LoadInt32(&v))
	}
	old := SwapInt32(&v, 2)
	if old != 1 {
		t.Log(LoadInt32(&v))
	}
	if AddInt32(&v, 1) != 3 {
		t.Log(LoadInt32(&v))
	}
	if !CompareAndSwapInt32(&v, 3, 4) {
		t.Log(LoadInt32(&v))
	}
	if CompareAndSwapInt32(&v, 3, 4) {
		t.Log(LoadInt32(&v))
	}
}

func TestAtomicInt64(t *testing.T) {
	var v int64
	StoreInt64(&v, 1)
	if LoadInt64(&v) != 1 {
		t.Log(LoadInt64(&v))
	}
	old := SwapInt64(&v, 2)
	if old != 1 {
		t.Log(LoadInt64(&v))
	}
	if AddInt64(&v, 1) != 3 {
		t.Log(LoadInt64(&v))
	}
	if !CompareAndSwapInt64(&v, 3, 4) {
		t.Log(LoadInt64(&v))
	}
	if CompareAndSwapInt64(&v, 3, 4) {
		t.Log(LoadInt64(&v))
	}
}

func TestAtomicUint32(t *testing.T) {
	var v uint32
	StoreUint32(&v, 1)
	if LoadUint32(&v) != 1 {
		t.Log(LoadUint32(&v))
	}
	old := SwapUint32(&v, 2)
	if old != 1 {
		t.Log(LoadUint32(&v))
	}
	if AddUint32(&v, 1) != 3 {
		t.Log(LoadUint32(&v))
	}
	if !CompareAndSwapUint32(&v, 3, 4) {
		t.Log(LoadUint32(&v))
	}
	if CompareAndSwapUint32(&v, 3, 4) {
		t.Log(LoadUint32(&v))
	}
}

func TestAtomicUint64(t *testing.T) {
	var v uint64
	StoreUint64(&v, 1)
	if LoadUint64(&v) != 1 {
		t.Log(LoadUint64(&v))
	}
	old := SwapUint64(&v, 2)
	if old != 1 {
		t.Log(LoadUint64(&v))
	}
	if AddUint64(&v, 1) != 3 {
		t.Log(LoadUint64(&v))
	}
	if !CompareAndSwapUint64(&v, 3, 4) {
		t.Log(LoadUint64(&v))
	}
	if CompareAndSwapUint64(&v, 3, 4) {
		t.Log(LoadUint64(&v))
	}
}

func TestAtomicUintptr(t *testing.T) {
	var v uintptr
	StoreUintptr(&v, 1)
	if LoadUintptr(&v) != 1 {
		t.Log(LoadUintptr(&v))
	}
	old := SwapUintptr(&v, 2)
	if old != 1 {
		t.Log(LoadUintptr(&v))
	}
	if AddUintptr(&v, 1) != 3 {
		t.Log(LoadUintptr(&v))
	}
	if !CompareAndSwapUintptr(&v, 3, 4) {
		t.Log(LoadUintptr(&v))
	}
	if CompareAndSwapUintptr(&v, 3, 4) {
		t.Log(LoadUintptr(&v))
	}
}

func TestAtomicPointer(t *testing.T) {
	var v string
	var vp = unsafe.Pointer(&v)
	var v1 = "Hello World"
	var vp1 = unsafe.Pointer(&v1)
	StorePointer(&vp, vp1)
	if LoadPointer(&vp) != vp1 {
		t.Log(LoadPointer(&vp))
	}
	var v2 = "Foo"
	var vp2 = unsafe.Pointer(&v2)
	old := SwapPointer(&vp, vp2)
	if old != vp1 {
		t.Log(LoadPointer(&vp))
	}
	var v3 = "Bar"
	var vp3 = unsafe.Pointer(&v3)
	if !CompareAndSwapPointer(&vp, vp2, vp3) {
		t.Log(LoadPointer(&vp))
	}
	if CompareAndSwapPointer(&vp, vp2, vp3) {
		t.Log(LoadPointer(&vp))
	}
}
