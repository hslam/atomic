// Copyright (c) 2020 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

package atomic

import (
	"sync/atomic"
)

// Uint16 represents an uint16.
type Uint16 struct {
	v      uint32
	seting uint32
}

// NewUint16 returns a new Uint16.
func NewUint16(val uint16) *Uint16 {
	addr := &Uint16{}
	addr.Store(val)
	return addr
}

// Swap atomically stores new into *addr and returns the previous *addr value.
func (addr *Uint16) Swap(new uint16) (old uint16) {
	return SwapUint16(addr, new)
}

// CompareAndSwap executes the compare-and-swap operation for an uint16 value.
func (addr *Uint16) CompareAndSwap(old, new uint16) (swapped bool) {
	return CompareAndSwapUint16(addr, old, new)
}

// Add atomically adds delta to *addr and returns the new value.
func (addr *Uint16) Add(delta uint16) (new uint16) {
	return AddUint16(addr, delta)
}

// Load atomically loads *addr.
func (addr *Uint16) Load() (val uint16) {
	return LoadUint16(addr)
}

// Store atomically stores val into *addr.
func (addr *Uint16) Store(val uint16) {
	StoreUint16(addr, val)
}

// SwapUint16 atomically stores new into *addr and returns the previous *addr value.
func SwapUint16(addr *Uint16, new uint16) (old uint16) {
	for {
		if !atomic.CompareAndSwapUint32(&addr.seting, 0, 1) {
			continue
		}
		old = LoadUint16(addr)
		StoreUint16(addr, new)
		atomic.StoreUint32(&addr.seting, 0)
		return
	}
}

// CompareAndSwapUint16 executes the compare-and-swap operation for an uint16 value.
func CompareAndSwapUint16(addr *Uint16, old, new uint16) (swapped bool) {
	for {
		if !atomic.CompareAndSwapUint32(&addr.seting, 0, 1) {
			continue
		}
		if LoadUint16(addr) == old {
			StoreUint16(addr, new)
			atomic.StoreUint32(&addr.seting, 0)
			return true
		}
		atomic.StoreUint32(&addr.seting, 0)
		return false
	}
}

// AddUint16 atomically adds delta to *addr and returns the new value.
func AddUint16(addr *Uint16, delta uint16) (new uint16) {
	for {
		if !atomic.CompareAndSwapUint32(&addr.seting, 0, 1) {
			continue
		}
		new = LoadUint16(addr) + delta
		StoreUint16(addr, new)
		atomic.StoreUint32(&addr.seting, 0)
		return
	}
}

// LoadUint16 atomically loads *addr.
func LoadUint16(addr *Uint16) (val uint16) {
	var v = atomic.LoadUint32(&addr.v)
	return uint16(v)
}

// StoreUint16 atomically stores val into *addr.
func StoreUint16(addr *Uint16, val uint16) {
	atomic.StoreUint32(&addr.v, uint32(val))
}
