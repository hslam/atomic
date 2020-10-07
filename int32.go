// Copyright (c) 2020 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

package atomic

import (
	"sync/atomic"
)

// Int32 represents an int32.
type Int32 struct {
	v int32
}

// NewInt32 returns a new Int32.
func NewInt32(val int32) *Int32 {
	addr := &Int32{}
	addr.Store(val)
	return addr
}

// Swap atomically stores new into *addr and returns the previous *addr value.
func (addr *Int32) Swap(new int32) (old int32) {
	return SwapInt32(addr, new)
}

// CompareAndSwap executes the compare-and-swap operation for an int32 value.
func (addr *Int32) CompareAndSwap(old, new int32) (swapped bool) {
	return CompareAndSwapInt32(addr, old, new)
}

// Add atomically adds delta to *addr and returns the new value.
func (addr *Int32) Add(delta int32) (new int32) {
	return AddInt32(addr, delta)
}

// Load atomically loads *addr.
func (addr *Int32) Load() (val int32) {
	return LoadInt32(addr)
}

// Store atomically stores val into *addr.
func (addr *Int32) Store(val int32) {
	StoreInt32(addr, val)
}

// SwapInt32 atomically stores new into *addr and returns the previous *addr value.
func SwapInt32(addr *Int32, new int32) (old int32) {
	return atomic.SwapInt32(&addr.v, new)
}

// CompareAndSwapInt32 executes the compare-and-swap operation for an int32 value.
func CompareAndSwapInt32(addr *Int32, old, new int32) (swapped bool) {
	return atomic.CompareAndSwapInt32(&addr.v, old, new)
}

// AddInt32 atomically adds delta to *addr and returns the new value.
func AddInt32(addr *Int32, delta int32) (new int32) {
	return atomic.AddInt32(&addr.v, delta)
}

// LoadInt32 atomically loads *addr.
func LoadInt32(addr *Int32) (val int32) {
	return atomic.LoadInt32(&addr.v)
}

// StoreInt32 atomically stores val into *addr.
func StoreInt32(addr *Int32, val int32) {
	atomic.StoreInt32(&addr.v, val)
}
