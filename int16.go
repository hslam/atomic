// Copyright (c) 2020 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

package atomic

import (
	"sync/atomic"
)

// Int16 represents an int16.
type Int16 struct {
	v      uint32
	seting uint32
}

// NewInt16 returns a new Int16.
func NewInt16(val int16) *Int16 {
	addr := &Int16{}
	addr.Store(val)
	return addr
}

// Swap atomically stores new into *addr and returns the previous *addr value.
func (addr *Int16) Swap(new int16) (old int16) {
	for {
		if !atomic.CompareAndSwapUint32(&addr.seting, 0, 1) {
			continue
		}
		old = addr.Load()
		addr.Store(new)
		atomic.StoreUint32(&addr.seting, 0)
		return
	}
}

// CompareAndSwap executes the compare-and-swap operation for an int16 value.
func (addr *Int16) CompareAndSwap(old, new int16) (swapped bool) {
	for {
		if !atomic.CompareAndSwapUint32(&addr.seting, 0, 1) {
			continue
		}
		if addr.Load() == old {
			addr.Store(new)
			atomic.StoreUint32(&addr.seting, 0)
			return true
		}
		atomic.StoreUint32(&addr.seting, 0)
		return false
	}
}

// Add atomically adds delta to *addr and returns the new value.
func (addr *Int16) Add(delta int16) (new int16) {
	for {
		if !atomic.CompareAndSwapUint32(&addr.seting, 0, 1) {
			continue
		}
		new = addr.Load() + delta
		addr.Store(new)
		atomic.StoreUint32(&addr.seting, 0)
		return
	}
}

// Load atomically loads *addr.
func (addr *Int16) Load() (val int16) {
	var v = atomic.LoadUint32(&addr.v)
	return int16(v)
}

// Store atomically stores val into *addr.
func (addr *Int16) Store(val int16) {
	atomic.StoreUint32(&addr.v, uint32(val))
}
