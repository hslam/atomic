// Copyright (c) 2020 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

package atomic

import (
	"sync/atomic"
	"unsafe"
)

// Float64 represents an float64.
type Float64 struct {
	v      uint64
	seting uint32
}

// NewFloat64 returns a new Float64.
func NewFloat64(val float64) *Float64 {
	addr := &Float64{}
	addr.Store(val)
	return addr
}

// Swap atomically stores new into *addr and returns the previous *addr value.
func (addr *Float64) Swap(new float64) (old float64) {
	return SwapFloat64(addr, new)
}

// CompareAndSwap executes the compare-and-swap operation for an float64 value.
func (addr *Float64) CompareAndSwap(old, new float64) (swapped bool) {
	return CompareAndSwapFloat64(addr, old, new)
}

// Add atomically adds delta to *addr and returns the new value.
func (addr *Float64) Add(delta float64) (new float64) {
	return AddFloat64(addr, delta)
}

// Load atomically loads *addr.
func (addr *Float64) Load() (val float64) {
	return LoadFloat64(addr)
}

// Store atomically stores val into *addr.
func (addr *Float64) Store(val float64) {
	StoreFloat64(addr, val)
}

// SwapFloat64 atomically stores new into *addr and returns the previous *addr value.
func SwapFloat64(addr *Float64, new float64) (old float64) {
	for {
		if !atomic.CompareAndSwapUint32(&addr.seting, 0, 1) {
			continue
		}
		old = LoadFloat64(addr)
		StoreFloat64(addr, new)
		atomic.StoreUint32(&addr.seting, 0)
		return
	}
}

// CompareAndSwapFloat64 executes the compare-and-swap operation for an float64 value.
func CompareAndSwapFloat64(addr *Float64, old, new float64) (swapped bool) {
	for {
		if !atomic.CompareAndSwapUint32(&addr.seting, 0, 1) {
			continue
		}
		if LoadFloat64(addr) == old {
			StoreFloat64(addr, new)
			atomic.StoreUint32(&addr.seting, 0)
			return true
		}
		atomic.StoreUint32(&addr.seting, 0)
		return false
	}
}

// AddFloat64 atomically adds delta to *addr and returns the new value.
func AddFloat64(addr *Float64, delta float64) (new float64) {
	for {
		if !atomic.CompareAndSwapUint32(&addr.seting, 0, 1) {
			continue
		}
		new = LoadFloat64(addr) + delta
		StoreFloat64(addr, new)
		atomic.StoreUint32(&addr.seting, 0)
		return
	}
}

// LoadFloat64 atomically loads *addr.
func LoadFloat64(addr *Float64) (val float64) {
	var v = atomic.LoadUint64(&addr.v)
	return *(*float64)(unsafe.Pointer(&v))
}

// StoreFloat64 atomically stores val into *addr.
func StoreFloat64(addr *Float64, val float64) {
	atomic.StoreUint64(&addr.v, *(*uint64)(unsafe.Pointer(&val)))
}
