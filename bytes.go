// Copyright (c) 2020 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

package atomic

import (
	"bytes"
	"sync/atomic"
)

// Bytes represents an []byte.
type Bytes struct {
	v       *atomic.Value
	seting  uint32
	initing uint32
}

// NewBytes returns a new Bytes.
func NewBytes(val []byte) *Bytes {
	addr := &Bytes{v: &atomic.Value{}}
	addr.Store(val)
	return addr
}

// Swap atomically stores new into *addr and returns the previous *addr value.
func (addr *Bytes) Swap(new []byte) (old []byte) {
	return SwapBytes(addr, new)
}

// CompareAndSwap executes the compare-and-swap operation for an []byte value.
func (addr *Bytes) CompareAndSwap(old, new []byte) (swapped bool) {
	return CompareAndSwapBytes(addr, old, new)
}

// Add atomically adds delta to *addr and returns the new value.
func (addr *Bytes) Add(delta []byte) (new []byte) {
	return AddBytes(addr, delta)
}

// Load atomically loads *addr.
func (addr *Bytes) Load() (val []byte) {
	return LoadBytes(addr)
}

// Store atomically stores val into *addr.
func (addr *Bytes) Store(val []byte) {
	StoreBytes(addr, val)
}

// SwapBytes atomically stores new into *addr and returns the previous *addr value.
func SwapBytes(addr *Bytes, new []byte) (old []byte) {
	for {
		if !atomic.CompareAndSwapUint32(&addr.seting, 0, 1) {
			continue
		}
		old = LoadBytes(addr)
		StoreBytes(addr, new)
		atomic.StoreUint32(&addr.seting, 0)
		return
	}
}

// CompareAndSwapBytes executes the compare-and-swap operation for an []byte value.
func CompareAndSwapBytes(addr *Bytes, old, new []byte) (swapped bool) {
	for {
		if !atomic.CompareAndSwapUint32(&addr.seting, 0, 1) {
			continue
		}
		if bytes.Equal(LoadBytes(addr), old) {
			StoreBytes(addr, new)
			atomic.StoreUint32(&addr.seting, 0)
			return true
		}
		atomic.StoreUint32(&addr.seting, 0)
		return false
	}
}

// AddBytes atomically adds delta to *addr and returns the new value.
func AddBytes(addr *Bytes, delta []byte) (new []byte) {
	for {
		if !atomic.CompareAndSwapUint32(&addr.seting, 0, 1) {
			continue
		}
		new = append(LoadBytes(addr), delta...)
		StoreBytes(addr, new)
		atomic.StoreUint32(&addr.seting, 0)
		return
	}
}

// LoadBytes atomically loads *addr.
func LoadBytes(addr *Bytes) (val []byte) {
	if addr.v == nil {
		return nil
	}
	var ok bool
	if val, ok = addr.v.Load().([]byte); ok {
		return val
	}
	return nil
}

// StoreBytes atomically stores val into *addr.
func StoreBytes(addr *Bytes, val []byte) {
	if addr.v == nil {
		initBytes(addr)
	}
	addr.v.Store(val)
}

func initBytes(addr *Bytes) {
	for {
		if !atomic.CompareAndSwapUint32(&addr.initing, 0, 1) {
			continue
		}
		if addr.v == nil {
			addr.v = &atomic.Value{}
		}
		atomic.StoreUint32(&addr.initing, 0)
		break
	}
}
