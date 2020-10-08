// Copyright (c) 2020 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

// Package atomic provides low-level atomic memory primitives
// useful for implementing synchronization algorithms.
package atomic

import (
	"sync/atomic"
	"unsafe"
)

// SwapInt32 atomically stores new into *addr and returns the previous *addr value.
func SwapInt32(addr *int32, new int32) (old int32) {
	return atomic.SwapInt32(addr, new)
}

// SwapInt64 atomically stores new into *addr and returns the previous *addr value.
func SwapInt64(addr *int64, new int64) (old int64) {
	return atomic.SwapInt64(addr, new)
}

// SwapUint32 atomically stores new into *addr and returns the previous *addr value.
func SwapUint32(addr *uint32, new uint32) (old uint32) {
	return atomic.SwapUint32(addr, new)
}

// SwapUint64 atomically stores new into *addr and returns the previous *addr value.
func SwapUint64(addr *uint64, new uint64) (old uint64) {
	return atomic.SwapUint64(addr, new)
}

// SwapUintptr atomically stores new into *addr and returns the previous *addr value.
func SwapUintptr(addr *uintptr, new uintptr) (old uintptr) {
	return atomic.SwapUintptr(addr, new)
}

// SwapPointer atomically stores new into *addr and returns the previous *addr value.
func SwapPointer(addr *unsafe.Pointer, new unsafe.Pointer) (old unsafe.Pointer) {
	return atomic.SwapPointer(addr, new)
}

// CompareAndSwapInt32 executes the compare-and-swap operation for an int32 value.
func CompareAndSwapInt32(addr *int32, old, new int32) (swapped bool) {
	return atomic.CompareAndSwapInt32(addr, old, new)
}

// CompareAndSwapInt64 executes the compare-and-swap operation for an int64 value.
func CompareAndSwapInt64(addr *int64, old, new int64) (swapped bool) {
	return atomic.CompareAndSwapInt64(addr, old, new)
}

// CompareAndSwapUint32 executes the compare-and-swap operation for a uint32 value.
func CompareAndSwapUint32(addr *uint32, old, new uint32) (swapped bool) {
	return atomic.CompareAndSwapUint32(addr, old, new)
}

// CompareAndSwapUint64 executes the compare-and-swap operation for a uint64 value.
func CompareAndSwapUint64(addr *uint64, old, new uint64) (swapped bool) {
	return atomic.CompareAndSwapUint64(addr, old, new)
}

// CompareAndSwapUintptr executes the compare-and-swap operation for a uintptr value.
func CompareAndSwapUintptr(addr *uintptr, old, new uintptr) (swapped bool) {
	return atomic.CompareAndSwapUintptr(addr, old, new)
}

// CompareAndSwapPointer executes the compare-and-swap operation for a unsafe.Pointer value.
func CompareAndSwapPointer(addr *unsafe.Pointer, old, new unsafe.Pointer) (swapped bool) {
	return atomic.CompareAndSwapPointer(addr, old, new)
}

// AddInt32 atomically adds delta to *addr and returns the new value.
func AddInt32(addr *int32, delta int32) (new int32) {
	return atomic.AddInt32(addr, delta)
}

// AddUint32 atomically adds delta to *addr and returns the new value.
// To subtract a signed positive constant value c from x, do AddUint32(&x, ^uint32(c-1)).
// In particular, to decrement x, do AddUint32(&x, ^uint32(0)).
func AddUint32(addr *uint32, delta uint32) (new uint32) {
	return atomic.AddUint32(addr, delta)
}

// AddInt64 atomically adds delta to *addr and returns the new value.
func AddInt64(addr *int64, delta int64) (new int64) {
	return atomic.AddInt64(addr, delta)
}

// AddUint64 atomically adds delta to *addr and returns the new value.
// To subtract a signed positive constant value c from x, do AddUint64(&x, ^uint64(c-1)).
// In particular, to decrement x, do AddUint64(&x, ^uint64(0)).
func AddUint64(addr *uint64, delta uint64) (new uint64) {
	return atomic.AddUint64(addr, delta)
}

// AddUintptr atomically adds delta to *addr and returns the new value.
func AddUintptr(addr *uintptr, delta uintptr) (new uintptr) {
	return atomic.AddUintptr(addr, delta)
}

// LoadInt32 atomically loads *addr.
func LoadInt32(addr *int32) (val int32) {
	return atomic.LoadInt32(addr)
}

// LoadInt64 atomically loads *addr.
func LoadInt64(addr *int64) (val int64) {
	return atomic.LoadInt64(addr)
}

// LoadUint32 atomically loads *addr.
func LoadUint32(addr *uint32) (val uint32) {
	return atomic.LoadUint32(addr)
}

// LoadUint64 atomically loads *addr.
func LoadUint64(addr *uint64) (val uint64) {
	return atomic.LoadUint64(addr)
}

// LoadUintptr atomically loads *addr.
func LoadUintptr(addr *uintptr) (val uintptr) {
	return atomic.LoadUintptr(addr)
}

// LoadPointer atomically loads *addr.
func LoadPointer(addr *unsafe.Pointer) (val unsafe.Pointer) {
	return atomic.LoadPointer(addr)
}

// StoreInt32 atomically stores val into *addr.
func StoreInt32(addr *int32, val int32) {
	atomic.StoreInt32(addr, val)
}

// StoreInt64 atomically stores val into *addr.
func StoreInt64(addr *int64, val int64) {
	atomic.StoreInt64(addr, val)
}

// StoreUint32 atomically stores val into *addr.
func StoreUint32(addr *uint32, val uint32) {
	atomic.StoreUint32(addr, val)
}

// StoreUint64 atomically stores val into *addr.
func StoreUint64(addr *uint64, val uint64) {
	atomic.StoreUint64(addr, val)
}

// StoreUintptr atomically stores val into *addr.
func StoreUintptr(addr *uintptr, val uintptr) {
	atomic.StoreUintptr(addr, val)
}

// StorePointer atomically stores val into *addr.
func StorePointer(addr *unsafe.Pointer, val unsafe.Pointer) {
	atomic.StorePointer(addr, val)
}
