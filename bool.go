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

// CompareAndSwap executes the compare-and-swap operation for an bool value.
func (addr *Bool) CompareAndSwap(old, new bool) (swapped bool) {
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
func (addr *Bool) Add(delta bool) (new bool) {
	for {
		if !atomic.CompareAndSwapUint32(&addr.seting, 0, 1) {
			continue
		}
		new = addr.Load() && delta
		addr.Store(new)
		atomic.StoreUint32(&addr.seting, 0)
		return
	}
}

// Load atomically loads *addr.
func (addr *Bool) Load() (val bool) {
	if atomic.LoadUint32(&addr.v) == 1 {
		return true
	}
	return false
}

// Store atomically stores val into *addr.
func (addr *Bool) Store(val bool) {
	if val {
		atomic.StoreUint32(&addr.v, 1)
	} else {
		atomic.StoreUint32(&addr.v, 0)
	}
}
