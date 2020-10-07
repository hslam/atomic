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
	return atomic.SwapInt32(&addr.v, new)
}

// CompareAndSwap executes the compare-and-swap operation for an int32 value.
func (addr *Int32) CompareAndSwap(old, new int32) (swapped bool) {
	return atomic.CompareAndSwapInt32(&addr.v, old, new)
}

// Add atomically adds delta to *addr and returns the new value.
func (addr *Int32) Add(delta int32) (new int32) {
	return atomic.AddInt32(&addr.v, delta)
}

// Load atomically loads *addr.
func (addr *Int32) Load() (val int32) {
	return atomic.LoadInt32(&addr.v)
}

// Store atomically stores val into *addr.
func (addr *Int32) Store(val int32) {
	atomic.StoreInt32(&addr.v, val)
}
