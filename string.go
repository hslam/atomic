// Copyright (c) 2020 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

package atomic

// String represents an string.
type String struct {
	v Value
}

// NewString returns a new String.
func NewString(val string) *String {
	addr := &String{}
	addr.Store(val)
	return addr
}

// Swap atomically stores new into *addr and returns the previous *addr value.
func (addr *String) Swap(new string) (old string) {
	for {
		load := addr.v.Load()
		if addr.v.compareAndSwap(load, new) {
			return load.(string)
		}
	}
}

// CompareAndSwap executes the compare-and-swap operation for an string value.
func (addr *String) CompareAndSwap(old, new string) (swapped bool) {
	load := addr.v.Load()
	if old != load {
		return false
	}
	return addr.v.compareAndSwap(load, new)
}

// Add atomically adds delta to *addr and returns the new value.
func (addr *String) Add(delta string) (new string) {
	for {
		old := addr.v.Load()
		new = old.(string) + delta
		if addr.v.compareAndSwap(old, new) {
			return
		}
	}
}

// Load atomically loads *addr.
func (addr *String) Load() (val string) {
	v := addr.v.Load()
	if v == nil {
		return ""
	}
	return v.(string)
}

// Store atomically stores val into *addr.
func (addr *String) Store(val string) {
	addr.v.Store(val)
}
