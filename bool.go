// Copyright (c) 2020 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

package atomic

import (
	"sync/atomic"
)

// Bool represents an bool.
type Bool struct {
	v      uint32
	seting uint32
}

// NewBool returns a new Bool.
func NewBool(val bool) *Bool {
	addr := &Bool{}
	addr.Store(val)
	return addr
}

// Swap atomically stores new into *addr and returns the previous *addr value.
func (addr *Bool) Swap(new bool) (old bool) {
	return SwapBool(addr, new)
}

// CompareAndSwap executes the compare-and-swap operation for an bool value.
func (addr *Bool) CompareAndSwap(old, new bool) (swapped bool) {
	return CompareAndSwapBool(addr, old, new)
}

// Add atomically adds delta to *addr and returns the new value.
func (addr *Bool) Add(delta bool) (new bool) {
	return AddBool(addr, delta)
}

// Load atomically loads *addr.
func (addr *Bool) Load() (val bool) {
	return LoadBool(addr)
}

// Store atomically stores val into *addr.
func (addr *Bool) Store(val bool) {
	StoreBool(addr, val)
}

// SwapBool atomically stores new into *addr and returns the previous *addr value.
func SwapBool(addr *Bool, new bool) (old bool) {
	for {
		if !atomic.CompareAndSwapUint32(&addr.seting, 0, 1) {
			continue
		}
		old = LoadBool(addr)
		StoreBool(addr, new)
		atomic.StoreUint32(&addr.seting, 0)
		return
	}
}

// CompareAndSwapBool executes the compare-and-swap operation for an bool value.
func CompareAndSwapBool(addr *Bool, old, new bool) (swapped bool) {
	for {
		if !atomic.CompareAndSwapUint32(&addr.seting, 0, 1) {
			continue
		}
		if LoadBool(addr) == old {
			StoreBool(addr, new)
			atomic.StoreUint32(&addr.seting, 0)
			return true
		}
		atomic.StoreUint32(&addr.seting, 0)
		return false
	}
}

// AddBool atomically adds delta to *addr and returns the new value.
func AddBool(addr *Bool, delta bool) (new bool) {
	for {
		if !atomic.CompareAndSwapUint32(&addr.seting, 0, 1) {
			continue
		}
		new = LoadBool(addr) && delta
		StoreBool(addr, new)
		atomic.StoreUint32(&addr.seting, 0)
		return
	}
}

// LoadBool atomically loads *addr.
func LoadBool(addr *Bool) (val bool) {
	if atomic.LoadUint32(&addr.v) == 1 {
		return true
	}
	return false
}

// StoreBool atomically stores val into *addr.
func StoreBool(addr *Bool, val bool) {
	if val {
		atomic.StoreUint32(&addr.v, 1)
	} else {
		atomic.StoreUint32(&addr.v, 0)
	}
}
