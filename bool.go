// Copyright (c) 2020 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

package atomic

import (
	"sync/atomic"
)

// Bool represents an bool.
type Bool struct {
	v uint32
}

// NewBool returns a new Bool.
func NewBool(val bool) *Bool {
	addr := &Bool{}
	addr.Store(val)
	return addr
}

// Swap atomically stores new into *addr and returns the previous *addr value.
func (addr *Bool) Swap(new bool) (old bool) {
	var v = atomic.SwapUint32(&addr.v, boolToUint32(new))
	return uint32ToBool(v)
}

// CompareAndSwap executes the compare-and-swap operation for an bool value.
func (addr *Bool) CompareAndSwap(old, new bool) (swapped bool) {
	return atomic.CompareAndSwapUint32(&addr.v, boolToUint32(old), boolToUint32(new))
}

// Add atomically adds delta to *addr and returns the new value.
func (addr *Bool) Add(delta bool) (new bool) {
	for {
		old := addr.Load()
		new = old && delta
		if addr.CompareAndSwap(old, new) {
			return
		}
	}
}

// Load atomically loads *addr.
func (addr *Bool) Load() (val bool) {
	return uint32ToBool(atomic.LoadUint32(&addr.v))
}

// Store atomically stores val into *addr.
func (addr *Bool) Store(val bool) {
	atomic.StoreUint32(&addr.v, boolToUint32(val))
}

func boolToUint32(val bool) uint32 {
	if val {
		return 1
	}
	return 0
}

func uint32ToBool(val uint32) bool {
	if val > 0 {
		return true
	}
	return false
}
