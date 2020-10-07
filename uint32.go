// Copyright (c) 2020 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

package atomic

import (
	"sync/atomic"
)

// Uint32 represents an uint32.
type Uint32 struct {
	v uint32
}

// NewUint32 returns a new Uint32.
func NewUint32(val uint32) *Uint32 {
	addr := &Uint32{}
	addr.Store(val)
	return addr
}

// Swap atomically stores new into *addr and returns the previous *addr value.
func (addr *Uint32) Swap(new uint32) (old uint32) {
	return atomic.SwapUint32(&addr.v, new)
}

// CompareAndSwap executes the compare-and-swap operation for an uint32 value.
func (addr *Uint32) CompareAndSwap(old, new uint32) (swapped bool) {
	return atomic.CompareAndSwapUint32(&addr.v, old, new)
}

// Add atomically adds delta to *addr and returns the new value.
func (addr *Uint32) Add(delta uint32) (new uint32) {
	return atomic.AddUint32(&addr.v, delta)
}

// Load atomically loads *addr.
func (addr *Uint32) Load() (val uint32) {
	return atomic.LoadUint32(&addr.v)
}

// Store atomically stores val into *addr.
func (addr *Uint32) Store(val uint32) {
	atomic.StoreUint32(&addr.v, val)
}
