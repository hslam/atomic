// Copyright (c) 2020 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

package atomic

import (
	"sync/atomic"
)

// Int8 represents an int8.
type Int8 struct {
	v uint32
}

// NewInt8 returns a new Int8.
func NewInt8(val int8) *Int8 {
	addr := &Int8{}
	addr.Store(val)
	return addr
}

// Swap atomically stores new into *addr and returns the previous *addr value.
func (addr *Int8) Swap(new int8) (old int8) {
	var v = atomic.SwapUint32(&addr.v, uint32(new))
	return int8(v)
}

// CompareAndSwap executes the compare-and-swap operation for an int16 value.
func (addr *Int8) CompareAndSwap(old, new int8) (swapped bool) {
	return atomic.CompareAndSwapUint32(&addr.v, uint32(old), uint32(new))
}

// Add atomically adds delta to *addr and returns the new value.
func (addr *Int8) Add(delta int8) (new int8) {
	for {
		old := addr.Load()
		new = old + delta
		if addr.CompareAndSwap(old, new) {
			return
		}
	}
}

// Load atomically loads *addr.
func (addr *Int8) Load() (val int8) {
	var v = atomic.LoadUint32(&addr.v)
	return int8(v)
}

// Store atomically stores val into *addr.
func (addr *Int8) Store(val int8) {
	atomic.StoreUint32(&addr.v, uint32(val))
}
