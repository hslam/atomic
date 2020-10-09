// Copyright (c) 2020 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

package atomic

import (
	"sync/atomic"
)

// Uint16 represents an uint16.
type Uint16 struct {
	v uint32
}

// NewUint16 returns a new Uint16.
func NewUint16(val uint16) *Uint16 {
	addr := &Uint16{}
	addr.Store(val)
	return addr
}

// Swap atomically stores new into *addr and returns the previous *addr value.
func (addr *Uint16) Swap(new uint16) (old uint16) {
	for {
		old = addr.Load()
		if addr.CompareAndSwap(old, new) {
			return
		}
	}
}

// CompareAndSwap executes the compare-and-swap operation for an uint16 value.
func (addr *Uint16) CompareAndSwap(old, new uint16) (swapped bool) {
	return atomic.CompareAndSwapUint32(&addr.v, uint32(old), uint32(new))
}

// Add atomically adds delta to *addr and returns the new value.
func (addr *Uint16) Add(delta uint16) (new uint16) {
	for {
		old := addr.Load()
		new = old + delta
		if addr.CompareAndSwap(old, new) {
			return
		}
	}
}

// Load atomically loads *addr.
func (addr *Uint16) Load() (val uint16) {
	var v = atomic.LoadUint32(&addr.v)
	return uint16(v)
}

// Store atomically stores val into *addr.
func (addr *Uint16) Store(val uint16) {
	atomic.StoreUint32(&addr.v, uint32(val))
}
