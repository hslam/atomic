// Copyright (c) 2020 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

package atomic

import (
	"bytes"
	"sync/atomic"
)

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
		old := addr.v.Load()
		if addr.v.fastCompareAndSwap(old, new) {
			return old.([]byte)
		}
	}
}

// CompareAndSwap executes the compare-and-swap operation for an []byte value.
func (addr *Bytes) CompareAndSwap(old, new []byte) (swapped bool) {
	for {
		if !atomic.CompareAndSwapUint32(&addr.v.seting, 0, 1) {
			continue
		}
		if bytes.Equal(addr.Load(), old) {
			addr.Store(new)
			atomic.StoreUint32(&addr.v.seting, 0)
			return true
		}
		atomic.StoreUint32(&addr.v.seting, 0)
		return false
	}
}

// Add atomically adds delta to *addr and returns the new value.
func (addr *Bytes) Add(delta []byte) (new []byte) {
	for {
		old := addr.v.Load()
		new = append(old.([]byte), delta...)
		if addr.v.fastCompareAndSwap(old, new) {
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
