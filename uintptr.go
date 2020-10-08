// Copyright (c) 2020 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

package atomic

import (
	"sync/atomic"
)

// Uintptr represents an uintptr.
type Uintptr struct {
	v uintptr
}

// NewUintptr returns a new Uintptr.
func NewUintptr(val uintptr) *Uintptr {
	addr := &Uintptr{}
	addr.Store(val)
	return addr
}

// Swap atomically stores new into *addr and returns the previous *addr value.
func (addr *Uintptr) Swap(new uintptr) (old uintptr) {
	return atomic.SwapUintptr(&addr.v, new)
}

// CompareAndSwap executes the compare-and-swap operation for an uintptr value.
func (addr *Uintptr) CompareAndSwap(old, new uintptr) (swapped bool) {
	return atomic.CompareAndSwapUintptr(&addr.v, old, new)
}

// Add atomically adds delta to *addr and returns the new value.
func (addr *Uintptr) Add(delta uintptr) (new uintptr) {
	return atomic.AddUintptr(&addr.v, delta)
}

// Load atomically loads *addr.
func (addr *Uintptr) Load() (val uintptr) {
	return atomic.LoadUintptr(&addr.v)
}

// Store atomically stores val into *addr.
func (addr *Uintptr) Store(val uintptr) {
	atomic.StoreUintptr(&addr.v, val)
}
