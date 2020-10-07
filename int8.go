// Copyright (c) 2020 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

package atomic

import (
	"sync/atomic"
)

// Int8 represents an int8.
type Int8 struct {
	v      uint32
	seting uint32
}

// NewInt8 returns a new Int8.
func NewInt8(val int8) *Int8 {
	addr := &Int8{}
	addr.Store(val)
	return addr
}

// Swap atomically stores new into *addr and returns the previous *addr value.
func (addr *Int8) Swap(new int8) (old int8) {
	return SwapInt8(addr, new)
}

// CompareAndSwap executes the compare-and-swap operation for an int16 value.
func (addr *Int8) CompareAndSwap(old, new int8) (swapped bool) {
	return CompareAndSwapInt8(addr, old, new)
}

// Add atomically adds delta to *addr and returns the new value.
func (addr *Int8) Add(delta int8) (new int8) {
	return AddInt8(addr, delta)
}

// Load atomically loads *addr.
func (addr *Int8) Load() (val int8) {
	return LoadInt8(addr)
}

// Store atomically stores val into *addr.
func (addr *Int8) Store(val int8) {
	StoreInt8(addr, val)
}

// SwapInt8 atomically stores new into *addr and returns the previous *addr value.
func SwapInt8(addr *Int8, new int8) (old int8) {
	for {
		if !atomic.CompareAndSwapUint32(&addr.seting, 0, 1) {
			continue
		}
		old = LoadInt8(addr)
		StoreInt8(addr, new)
		atomic.StoreUint32(&addr.seting, 0)
		return
	}
}

// CompareAndSwapInt8 executes the compare-and-swap operation for an int16 value.
func CompareAndSwapInt8(addr *Int8, old, new int8) (swapped bool) {
	for {
		if !atomic.CompareAndSwapUint32(&addr.seting, 0, 1) {
			continue
		}
		if LoadInt8(addr) == old {
			StoreInt8(addr, new)
			atomic.StoreUint32(&addr.seting, 0)
			return true
		}
		atomic.StoreUint32(&addr.seting, 0)
		return false
	}
}

// AddInt8 atomically adds delta to *addr and returns the new value.
func AddInt8(addr *Int8, delta int8) (new int8) {
	for {
		if !atomic.CompareAndSwapUint32(&addr.seting, 0, 1) {
			continue
		}
		new = LoadInt8(addr) + delta
		StoreInt8(addr, new)
		atomic.StoreUint32(&addr.seting, 0)
		return
	}
}

// LoadInt8 atomically loads *addr.
func LoadInt8(addr *Int8) (val int8) {
	var v = atomic.LoadUint32(&addr.v)
	return int8(v)
}

// StoreInt8 atomically stores val into *addr.
func StoreInt8(addr *Int8, val int8) {
	atomic.StoreUint32(&addr.v, uint32(val))
}
