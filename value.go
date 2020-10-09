package atomic

import (
	"sync/atomic"
	"unsafe"
)

// AddFunc is a add function.
type AddFunc func(old, delta interface{}) (new interface{})

// Value provides an atomic load and store of a consistently typed value.
// The zero value for a Value returns nil from Load.
// Once Store has been called, a Value must not be copied.
//
// A Value must not be copied after first use.
type Value struct {
	fast    bool
	v       atomic.Value
	seting  uint32
	AddFunc AddFunc
}

// ifaceWords is interface{} internal representation.
type ifaceWords struct {
	typ  unsafe.Pointer
	data unsafe.Pointer
}

// NewValue returns a new Value.
func NewValue(val interface{}, addFunc AddFunc) *Value {
	addr := &Value{AddFunc: addFunc}
	addr.Store(val)
	return addr
}

// NewFastValue returns a new FastValue.
func NewFastValue(val interface{}, addFunc AddFunc) *Value {
	addr := &Value{AddFunc: addFunc, fast: true}
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
	if v.fast {
		return v.fastCompareAndSwap(old, new)
	}
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

func (v *Value) fastCompareAndSwap(old, new interface{}) (swapped bool) {
	if new == nil {
		panic("github.com/hslam/atomic: new is nil")
	}
	vp := (*ifaceWords)(unsafe.Pointer(&v.v))
	np := (*ifaceWords)(unsafe.Pointer(&new))
	for {
		typ := LoadPointer(&vp.typ)
		if typ == nil {
			// Attempt to start first store.
			// Disable preemption so that other goroutines can use
			// active spin wait to wait for completion; and so that
			// GC does not see the fake type accidentally.
			if !CompareAndSwapPointer(&vp.typ, nil, unsafe.Pointer(^uintptr(0))) {
				return false
			}
			// Complete first store.
			StorePointer(&vp.data, np.data)
			StorePointer(&vp.typ, np.typ)
			return
		}
		if uintptr(typ) == ^uintptr(0) {
			// First store in progress. Wait.
			// Since we disable preemption around the first store,
			// we can wait with active spinning.
			return false
		}
		if old == nil {
			panic("github.com/hslam/atomic: old is nil")
		}
		// First store completed. Check type and overwrite data.
		op := (*ifaceWords)(unsafe.Pointer(&old))
		if typ != op.typ {
			panic("github.com/hslam/atomic: old is inconsistently typed value")
		}
		if typ != np.typ {
			panic("github.com/hslam/atomic: new is inconsistently typed value")
		}
		return atomic.CompareAndSwapPointer(&vp.data, op.data, np.data)
	}
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
