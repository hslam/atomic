// Copyright (c) 2020 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

package atomic

import (
	"unsafe"
)

// bytesEqual reports whether a and b
// are the same length and contain the same bytes.
// A nil argument is equivalent to an empty slice.
func bytesEqual(a, b []byte) bool {
	// Neither cmd/compile nor gccgo allocates for these string conversions.
	return *(*string)(unsafe.Pointer(&a)) == *(*string)(unsafe.Pointer(&b))
}

// Bytes represents an []byte.
type Bytes struct {
	v Value
}

// NewBytes returns a new Bytes.
func NewBytes(val []byte) *Bytes {
	addr := &Bytes{}
	addr.Store(val)
	return addr
}

// Swap atomically stores new into *addr and returns the previous *addr value.
func (addr *Bytes) Swap(new []byte) (old []byte) {
	for {
		load := addr.v.Load()
		if addr.v.compareAndSwap(load, new) {
			return load.([]byte)
		}
	}
}

// CompareAndSwap executes the compare-and-swap operation for an []byte value.
func (addr *Bytes) CompareAndSwap(old, new []byte) (swapped bool) {
	load := addr.v.Load()
	if !bytesEqual(old, load.([]byte)) {
		return false
	}
	return addr.v.compareAndSwap(load, new)
}

// Add atomically adds delta to *addr and returns the new value.
func (addr *Bytes) Add(delta []byte) (new []byte) {
	for {
		old := addr.v.Load()
		new = append(old.([]byte), delta...)
		if addr.v.compareAndSwap(old, new) {
			return
		}
	}
}

// Load atomically loads *addr.
func (addr *Bytes) Load() (val []byte) {
	v := addr.v.Load()
	if v == nil {
		return nil
	}
	return v.([]byte)
}

// Store atomically stores val into *addr.
func (addr *Bytes) Store(val []byte) {
	addr.v.Store(val)
}
