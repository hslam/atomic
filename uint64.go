// Copyright (c) 2020 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

package atomic

import (
	"sync/atomic"
)

// Uint64 represents an uint64.
type Uint64 struct {
	v uint64
}

// NewUint64 returns a new Uint64.
func NewUint64(val uint64) *Uint64 {
	addr := &Uint64{}
	addr.Store(val)
	return addr
}

// Swap atomically stores new into *addr and returns the previous *addr value.
func (addr *Uint64) Swap(new uint64) (old uint64) {
	return SwapUint64(addr, new)
}

// CompareAndSwap executes the compare-and-swap operation for an uint64 value.
func (addr *Uint64) CompareAndSwap(old, new uint64) (swapped bool) {
	return CompareAndSwapUint64(addr, old, new)
}

// Add atomically adds delta to *addr and returns the new value.
func (addr *Uint64) Add(delta uint64) (new uint64) {
	return AddUint64(addr, delta)
}

// Load atomically loads *addr.
func (addr *Uint64) Load() (val uint64) {
	return LoadUint64(addr)
}

// Store atomically stores val into *addr.
func (addr *Uint64) Store(val uint64) {
	StoreUint64(addr, val)
}

// SwapUint64 atomically stores new into *addr and returns the previous *addr value.
func SwapUint64(addr *Uint64, new uint64) (old uint64) {
	return atomic.SwapUint64(&addr.v, new)
}

// CompareAndSwapUint64 executes the compare-and-swap operation for an uint64 value.
func CompareAndSwapUint64(addr *Uint64, old, new uint64) (swapped bool) {
	return atomic.CompareAndSwapUint64(&addr.v, old, new)
}

// AddUint64 atomically adds delta to *addr and returns the new value.
func AddUint64(addr *Uint64, delta uint64) (new uint64) {
	return atomic.AddUint64(&addr.v, delta)
}

// LoadUint64 atomically loads *addr.
func LoadUint64(addr *Uint64) (val uint64) {
	return atomic.LoadUint64(&addr.v)
}

// StoreUint64 atomically stores val into *addr.
func StoreUint64(addr *Uint64, val uint64) {
	atomic.StoreUint64(&addr.v, val)
}
