// Copyright (c) 2020 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

package atomic

import (
	"sync/atomic"
	"unsafe"
)

// Pointer represents an unsafe.Pointer.
type Pointer struct {
	v unsafe.Pointer
}

// NewPointer returns a new Pointer.
func NewPointer(val unsafe.Pointer) *Pointer {
	addr := &Pointer{}
	addr.Store(val)
	return addr
}

// Swap atomically stores new into *addr and returns the previous *addr value.
func (addr *Pointer) Swap(new unsafe.Pointer) (old unsafe.Pointer) {
	return atomic.SwapPointer(&addr.v, new)
}

// CompareAndSwap executes the compare-and-swap operation for an uint64 value.
func (addr *Pointer) CompareAndSwap(old, new unsafe.Pointer) (swapped bool) {
	return atomic.CompareAndSwapPointer(&addr.v, old, new)
}

// Load atomically loads *addr.
func (addr *Pointer) Load() (val unsafe.Pointer) {
	return atomic.LoadPointer(&addr.v)
}

// Store atomically stores val into *addr.
func (addr *Pointer) Store(val unsafe.Pointer) {
	atomic.StorePointer(&addr.v, val)
}
