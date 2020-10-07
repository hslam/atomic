// Copyright (c) 2020 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

package atomic

import (
	"sync/atomic"
	"unsafe"
)

// Float32 represents an float32.
type Float32 struct {
	v      uint32
	seting uint32
}

// NewFloat32 returns a new Float32.
func NewFloat32(val float32) *Float32 {
	addr := &Float32{}
	addr.Store(val)
	return addr
}

// Swap atomically stores new into *addr and returns the previous *addr value.
func (addr *Float32) Swap(new float32) (old float32) {
	return SwapFloat32(addr, new)
}

// CompareAndSwap executes the compare-and-swap operation for an float32 value.
func (addr *Float32) CompareAndSwap(old, new float32) (swapped bool) {
	return CompareAndSwapFloat32(addr, old, new)
}

// Add atomically adds delta to *addr and returns the new value.
func (addr *Float32) Add(delta float32) (new float32) {
	return AddFloat32(addr, delta)
}

// Load atomically loads *addr.
func (addr *Float32) Load() (val float32) {
	return LoadFloat32(addr)
}

// Store atomically stores val into *addr.
func (addr *Float32) Store(val float32) {
	StoreFloat32(addr, val)
}

// SwapFloat32 atomically stores new into *addr and returns the previous *addr value.
func SwapFloat32(addr *Float32, new float32) (old float32) {
	for {
		if !atomic.CompareAndSwapUint32(&addr.seting, 0, 1) {
			continue
		}
		old = LoadFloat32(addr)
		StoreFloat32(addr, new)
		atomic.StoreUint32(&addr.seting, 0)
		return
	}
}

// CompareAndSwapFloat32 executes the compare-and-swap operation for an float32 value.
func CompareAndSwapFloat32(addr *Float32, old, new float32) (swapped bool) {
	for {
		if !atomic.CompareAndSwapUint32(&addr.seting, 0, 1) {
			continue
		}
		if LoadFloat32(addr) == old {
			StoreFloat32(addr, new)
			atomic.StoreUint32(&addr.seting, 0)
			return true
		}
		atomic.StoreUint32(&addr.seting, 0)
		return false
	}
}

// AddFloat32 atomically adds delta to *addr and returns the new value.
func AddFloat32(addr *Float32, delta float32) (new float32) {
	for {
		if !atomic.CompareAndSwapUint32(&addr.seting, 0, 1) {
			continue
		}
		new = LoadFloat32(addr) + delta
		StoreFloat32(addr, new)
		atomic.StoreUint32(&addr.seting, 0)
		return
	}
}

// LoadFloat32 atomically loads *addr.
func LoadFloat32(addr *Float32) (val float32) {
	var v = atomic.LoadUint32(&addr.v)
	return *(*float32)(unsafe.Pointer(&v))
}

// StoreFloat32 atomically stores val into *addr.
func StoreFloat32(addr *Float32, val float32) {
	atomic.StoreUint32(&addr.v, *(*uint32)(unsafe.Pointer(&val)))
}
