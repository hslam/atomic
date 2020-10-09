// Copyright (c) 2020 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

package atomic

import (
	"sync/atomic"
)

// Uint8 represents an uint8.
type Uint8 struct {
	v uint32
}

// NewUint8 returns a new Uint8.
func NewUint8(val uint8) *Uint8 {
	addr := &Uint8{}
	addr.Store(val)
	return addr
}

// Swap atomically stores new into *addr and returns the previous *addr value.
func (addr *Uint8) Swap(new uint8) (old uint8) {
	var v = atomic.SwapUint32(&addr.v, uint32(new))
	return uint8(v)
}

// CompareAndSwap executes the compare-and-swap operation for an uint8 value.
func (addr *Uint8) CompareAndSwap(old, new uint8) (swapped bool) {
	return atomic.CompareAndSwapUint32(&addr.v, uint32(old), uint32(new))
}

// Add atomically adds delta to *addr and returns the new value.
func (addr *Uint8) Add(delta uint8) (new uint8) {
	for {
		old := addr.Load()
		new = old + delta
		if addr.CompareAndSwap(old, new) {
			return
		}
	}
}

// Load atomically loads *addr.
func (addr *Uint8) Load() (val uint8) {
	var v = atomic.LoadUint32(&addr.v)
	return uint8(v)
}

// Store atomically stores val into *addr.
func (addr *Uint8) Store(val uint8) {
	atomic.StoreUint32(&addr.v, uint32(val))
}
