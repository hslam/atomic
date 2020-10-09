// Copyright (c) 2020 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

package atomic

import (
	"sync/atomic"
	"unsafe"
)

// Float64 represents an float64.
type Float64 struct {
	v uint64
}

// NewFloat64 returns a new Float64.
func NewFloat64(val float64) *Float64 {
	addr := &Float64{}
	addr.Store(val)
	return addr
}

// Swap atomically stores new into *addr and returns the previous *addr value.
func (addr *Float64) Swap(new float64) (old float64) {
	var v = atomic.SwapUint64(&addr.v, *(*uint64)(unsafe.Pointer(&new)))
	return *(*float64)(unsafe.Pointer(&v))
}

// CompareAndSwap executes the compare-and-swap operation for an float64 value.
func (addr *Float64) CompareAndSwap(old, new float64) (swapped bool) {
	return atomic.CompareAndSwapUint64(&addr.v, *(*uint64)(unsafe.Pointer(&old)), *(*uint64)(unsafe.Pointer(&new)))
}

// Add atomically adds delta to *addr and returns the new value.
func (addr *Float64) Add(delta float64) (new float64) {
	for {
		old := addr.Load()
		new = old + delta
		if addr.CompareAndSwap(old, new) {
			return
		}
	}
}

// Load atomically loads *addr.
func (addr *Float64) Load() (val float64) {
	var v = atomic.LoadUint64(&addr.v)
	return *(*float64)(unsafe.Pointer(&v))
}

// Store atomically stores val into *addr.
func (addr *Float64) Store(val float64) {
	atomic.StoreUint64(&addr.v, *(*uint64)(unsafe.Pointer(&val)))
}
