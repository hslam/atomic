// Copyright (c) 2020 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

package atomic

import (
	"sync/atomic"
)

// Int16 represents an int16.
type Int16 struct {
	v      uint32
	seting uint32
}

// NewInt16 returns a new Int16.
func NewInt16(val int16) *Int16 {
	addr := &Int16{}
	addr.Store(val)
	return addr
}

// Swap atomically stores new into *addr and returns the previous *addr value.
func (addr *Int16) Swap(new int16) (old int16) {
	return SwapInt16(addr, new)
}

// CompareAndSwap executes the compare-and-swap operation for an int16 value.
func (addr *Int16) CompareAndSwap(old, new int16) (swapped bool) {
	return CompareAndSwapInt16(addr, old, new)
}

// Add atomically adds delta to *addr and returns the new value.
func (addr *Int16) Add(delta int16) (new int16) {
	return AddInt16(addr, delta)
}

// Load atomically loads *addr.
func (addr *Int16) Load() (val int16) {
	return LoadInt16(addr)
}

// Store atomically stores val into *addr.
func (addr *Int16) Store(val int16) {
	StoreInt16(addr, val)
}

// SwapInt16 atomically stores new into *addr and returns the previous *addr value.
func SwapInt16(addr *Int16, new int16) (old int16) {
	for {
		if !atomic.CompareAndSwapUint32(&addr.seting, 0, 1) {
			continue
		}
		old = LoadInt16(addr)
		StoreInt16(addr, new)
		atomic.StoreUint32(&addr.seting, 0)
		return
	}
}

// CompareAndSwapInt16 executes the compare-and-swap operation for an int16 value.
func CompareAndSwapInt16(addr *Int16, old, new int16) (swapped bool) {
	for {
		if !atomic.CompareAndSwapUint32(&addr.seting, 0, 1) {
			continue
		}
		if LoadInt16(addr) == old {
			StoreInt16(addr, new)
			atomic.StoreUint32(&addr.seting, 0)
			return true
		}
		atomic.StoreUint32(&addr.seting, 0)
		return false
	}
}

// AddInt16 atomically adds delta to *addr and returns the new value.
func AddInt16(addr *Int16, delta int16) (new int16) {
	for {
		if !atomic.CompareAndSwapUint32(&addr.seting, 0, 1) {
			continue
		}
		new = LoadInt16(addr) + delta
		StoreInt16(addr, new)
		atomic.StoreUint32(&addr.seting, 0)
		return
	}
}

// LoadInt16 atomically loads *addr.
func LoadInt16(addr *Int16) (val int16) {
	var v = atomic.LoadUint32(&addr.v)
	return int16(v)
}

// StoreInt16 atomically stores val into *addr.
func StoreInt16(addr *Int16, val int16) {
	atomic.StoreUint32(&addr.v, uint32(val))
}
