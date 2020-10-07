// Copyright (c) 2020 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

package atomic

import (
	"sync/atomic"
)

// Uint8 represents an uint8.
type Uint8 struct {
	v      uint32
	seting uint32
}

// NewUint8 returns a new Uint8.
func NewUint8(val uint8) *Uint8 {
	addr := &Uint8{}
	addr.Store(val)
	return addr
}

// Swap atomically stores new into *addr and returns the previous *addr value.
func (addr *Uint8) Swap(new uint8) (old uint8) {
	return SwapUint8(addr, new)
}

// CompareAndSwap executes the compare-and-swap operation for an uint8 value.
func (addr *Uint8) CompareAndSwap(old, new uint8) (swapped bool) {
	return CompareAndSwapUint8(addr, old, new)
}

// Add atomically adds delta to *addr and returns the new value.
func (addr *Uint8) Add(delta uint8) (new uint8) {
	return AddUint8(addr, delta)
}

// Load atomically loads *addr.
func (addr *Uint8) Load() (val uint8) {
	return LoadUint8(addr)
}

// Store atomically stores val into *addr.
func (addr *Uint8) Store(val uint8) {
	StoreUint8(addr, val)
}

// SwapUint8 atomically stores new into *addr and returns the previous *addr value.
func SwapUint8(addr *Uint8, new uint8) (old uint8) {
	for {
		if !atomic.CompareAndSwapUint32(&addr.seting, 0, 1) {
			continue
		}
		old = LoadUint8(addr)
		StoreUint8(addr, new)
		atomic.StoreUint32(&addr.seting, 0)
		return
	}
}

// CompareAndSwapUint8 executes the compare-and-swap operation for an uint8 value.
func CompareAndSwapUint8(addr *Uint8, old, new uint8) (swapped bool) {
	for {
		if !atomic.CompareAndSwapUint32(&addr.seting, 0, 1) {
			continue
		}
		if LoadUint8(addr) == old {
			StoreUint8(addr, new)
			atomic.StoreUint32(&addr.seting, 0)
			return true
		}
		atomic.StoreUint32(&addr.seting, 0)
		return false
	}
}

// AddUint8 atomically adds delta to *addr and returns the new value.
func AddUint8(addr *Uint8, delta uint8) (new uint8) {
	for {
		if !atomic.CompareAndSwapUint32(&addr.seting, 0, 1) {
			continue
		}
		new = LoadUint8(addr) + delta
		StoreUint8(addr, new)
		atomic.StoreUint32(&addr.seting, 0)
		return
	}
}

// LoadUint8 atomically loads *addr.
func LoadUint8(addr *Uint8) (val uint8) {
	var v = atomic.LoadUint32(&addr.v)
	return uint8(v)
}

// StoreUint8 atomically stores val into *addr.
func StoreUint8(addr *Uint8, val uint8) {
	atomic.StoreUint32(&addr.v, uint32(val))
}
