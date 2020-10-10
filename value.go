package atomic

import (
	"sync/atomic"
	"unsafe"
)

// AddFunc is a add function.
type AddFunc func(old, delta interface{}) (new interface{})

// EqualFunc is a equal function.
type EqualFunc func(old, load interface{}) (equal bool)

// Value provides an atomic load and store of a consistently typed value.
// The zero value for a Value returns nil from Load.
// Once Store has been called, a Value must not be copied.
//
// A Value must not be copied after first use.
type Value struct {
	v         atomic.Value
	EqualFunc EqualFunc
	AddFunc   AddFunc
}

// ifaceWords is interface{} internal representation.
type ifaceWords struct {
	typ  unsafe.Pointer
	data unsafe.Pointer
}

// NewValue returns a new Value.
func NewValue(val interface{}, equalFunc EqualFunc, addFunc AddFunc) *Value {
	addr := &Value{EqualFunc: equalFunc, AddFunc: addFunc}
	addr.Store(val)
	return addr
}

// Swap atomically stores new into *addr and returns the previous *addr value.
func (v *Value) Swap(new interface{}) (old interface{}) {
	for {
		old = v.Load()
		if v.CompareAndSwap(old, new) {
			return
		}
	}
}

// CompareAndSwap executes the compare-and-swap operation for an interface{} value.
func (v *Value) CompareAndSwap(old, new interface{}) (swapped bool) {
	if v.EqualFunc == nil {
		panic("EqualFunc is nil")
	}
	load := v.Load()
	if !v.EqualFunc(old, load) {
		return false
	}
	return v.compareAndSwap(load, new)
}

func (v *Value) compareAndSwap(old, new interface{}) (swapped bool) {
	if new == nil {
		panic("github.com/hslam/atomic: new is nil")
	}
	vp := (*ifaceWords)(unsafe.Pointer(&v.v))
	np := (*ifaceWords)(unsafe.Pointer(&new))
	typ := LoadPointer(&vp.typ)
	if typ == nil {
		// Attempt to start first store.
		if !CompareAndSwapPointer(&vp.typ, nil, unsafe.Pointer(^uintptr(0))) {
			return false
		}
		// Complete first store.
		StorePointer(&vp.data, np.data)
		StorePointer(&vp.typ, np.typ)
		return
	}
	if uintptr(typ) == ^uintptr(0) {
		// First store in progress.
		return false
	}
	if old == nil {
		panic("github.com/hslam/atomic: old is nil")
	}
	// First store completed. Check type.
	op := (*ifaceWords)(unsafe.Pointer(&old))
	if typ != op.typ {
		panic("github.com/hslam/atomic: old is inconsistently typed value")
	}
	if typ != np.typ {
		panic("github.com/hslam/atomic: new is inconsistently typed value")
	}
	return atomic.CompareAndSwapPointer(&vp.data, op.data, np.data)
}

// Add atomically adds delta to *addr and returns the new value.
func (v *Value) Add(delta interface{}) (new interface{}) {
	if v.AddFunc == nil {
		panic("AddFunc is nil")
	}
	for {
		old := v.Load()
		new = v.AddFunc(old, delta)
		if v.CompareAndSwap(old, new) {
			return
		}
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
