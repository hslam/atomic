package atomic

import "sync/atomic"

// AddFunc is a add function.
type AddFunc func(old, delta interface{}) (new interface{})

// Value provides an atomic load and store of a consistently typed value.
// The zero value for a Value returns nil from Load.
// Once Store has been called, a Value must not be copied.
//
// A Value must not be copied after first use.
type Value struct {
	v       atomic.Value
	seting  uint32
	AddFunc AddFunc
}

// NewValue returns a new Value.
func NewValue(val interface{}, addFunc AddFunc) *Value {
	addr := &Value{AddFunc: addFunc}
	addr.Store(val)
	return addr
}

// Swap atomically stores new into *addr and returns the previous *addr value.
func (v *Value) Swap(new interface{}) (old interface{}) {
	for {
		if !atomic.CompareAndSwapUint32(&v.seting, 0, 1) {
			continue
		}
		old = v.Load()
		v.Store(new)
		atomic.StoreUint32(&v.seting, 0)
		return
	}
}

// CompareAndSwap executes the compare-and-swap operation for an interface{} value.
func (v *Value) CompareAndSwap(old, new interface{}) (swapped bool) {
	for {
		if !atomic.CompareAndSwapUint32(&v.seting, 0, 1) {
			continue
		}
		if v.Load() == old {
			v.Store(new)
			atomic.StoreUint32(&v.seting, 0)
			return true
		}
		atomic.StoreUint32(&v.seting, 0)
		return false
	}
}

// Add atomically adds delta to *addr and returns the new value.
func (v *Value) Add(delta interface{}) (new interface{}) {
	if v.AddFunc == nil {
		panic("AddFunc is nil")
	}
	for {
		if !atomic.CompareAndSwapUint32(&v.seting, 0, 1) {
			continue
		}
		new = v.AddFunc(v.Load(), delta)
		v.Store(new)
		atomic.StoreUint32(&v.seting, 0)
		return
	}
}

// Load returns the value set by the most recent Store.
// It returns nil if there has been no call to Store for this Value.
func (v *Value) Load() (x interface{}) {
	return v.v.Load()
}

// Store sets the value of the Value to x.
// All calls to Store for a given Value must use values of the same concrete type.
// Store of an inconsistent type panics, as does Store(nil).
func (v *Value) Store(x interface{}) {
	v.v.Store(x)
}
