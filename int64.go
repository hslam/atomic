// Copyright (c) 2020 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

package atomic

import (
	"sync/atomic"
)

// Int64 represents an int64.
type Int64 struct {
	v int64
}

// NewInt64 returns a new Int64.
func NewInt64(val int64) *Int64 {
	addr := &Int64{}
	addr.Store(val)
	return addr
}

// Swap atomically stores new into *addr and returns the previous *addr value.
func (addr *Int64) Swap(new int64) (old int64) {
	return SwapInt64(addr, new)
}

// CompareAndSwap executes the compare-and-swap operation for an int32 value.
func (addr *Int64) CompareAndSwap(old, new int64) (swapped bool) {
	return CompareAndSwapInt64(addr, old, new)
}

// Add atomically adds delta to *addr and returns the new value.
func (addr *Int64) Add(delta int64) (new int64) {
	return AddInt64(addr, delta)
}

// Load atomically loads *addr.
func (addr *Int64) Load() (val int64) {
	return LoadInt64(addr)
}

// Store atomically stores val into *addr.
func (addr *Int64) Store(val int64) {
	StoreInt64(addr, val)
}

// SwapInt64 atomically stores new into *addr and returns the previous *addr value.
func SwapInt64(addr *Int64, new int64) (old int64) {
	return atomic.SwapInt64(&addr.v, new)
}

// CompareAndSwapInt64 executes the compare-and-swap operation for an int32 value.
func CompareAndSwapInt64(addr *Int64, old, new int64) (swapped bool) {
	return atomic.CompareAndSwapInt64(&addr.v, old, new)
}

// AddInt64 atomically adds delta to *addr and returns the new value.
func AddInt64(addr *Int64, delta int64) (new int64) {
	return atomic.AddInt64(&addr.v, delta)
}

// LoadInt64 atomically loads *addr.
func LoadInt64(addr *Int64) (val int64) {
	return atomic.LoadInt64(&addr.v)
}

// StoreInt64 atomically stores val into *addr.
func StoreInt64(addr *Int64, val int64) {
	atomic.StoreInt64(&addr.v, val)
}
