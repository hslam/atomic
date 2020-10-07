// Copyright (c) 2020 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

package atomic

import (
	"sync/atomic"
)

// String represents an string.
type String struct {
	v      Value
	seting uint32
}

// NewString returns a new String.
func NewString(val string) *String {
	addr := &String{}
	addr.Store(val)
	return addr
}

// Swap atomically stores new into *addr and returns the previous *addr value.
func (addr *String) Swap(new string) (old string) {
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

// CompareAndSwap executes the compare-and-swap operation for an string value.
func (addr *String) CompareAndSwap(old, new string) (swapped bool) {
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
func (addr *String) Add(delta string) (new string) {
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
func (addr *String) Load() (val string) {
	v := addr.v.Load()
	if v == nil {
		return ""
	}
	return v.(string)
}

// Store atomically stores val into *addr.
func (addr *String) Store(val string) {
	addr.v.Store(val)
}
