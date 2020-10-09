// Copyright (c) 2020 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

package atomic

import (
	"sync/atomic"
	"unsafe"
)

// Float32 represents an float32.
type Float32 struct {
	v uint32
}

// NewFloat32 returns a new Float32.
func NewFloat32(val float32) *Float32 {
	addr := &Float32{}
	addr.Store(val)
	return addr
}

// Swap atomically stores new into *addr and returns the previous *addr value.
func (addr *Float32) Swap(new float32) (old float32) {
	var v = atomic.SwapUint32(&addr.v, *(*uint32)(unsafe.Pointer(&new)))
	return *(*float32)(unsafe.Pointer(&v))
}

// CompareAndSwap executes the compare-and-swap operation for an float32 value.
func (addr *Float32) CompareAndSwap(old, new float32) (swapped bool) {
	return atomic.CompareAndSwapUint32(&addr.v, *(*uint32)(unsafe.Pointer(&old)), *(*uint32)(unsafe.Pointer(&new)))
}

// Add atomically adds delta to *addr and returns the new value.
func (addr *Float32) Add(delta float32) (new float32) {
	for {
		old := addr.Load()
		new = old + delta
		if addr.CompareAndSwap(old, new) {
			return
		}
	}
}

// Load atomically loads *addr.
func (addr *Float32) Load() (val float32) {
	var v = atomic.LoadUint32(&addr.v)
	return *(*float32)(unsafe.Pointer(&v))
}

// Store atomically stores val into *addr.
func (addr *Float32) Store(val float32) {
	atomic.StoreUint32(&addr.v, *(*uint32)(unsafe.Pointer(&val)))
}
