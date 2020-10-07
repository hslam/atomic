package atomic

import "sync/atomic"

// Value provides an atomic load and store of a consistently typed value.
// The zero value for a Value returns nil from Load.
// Once Store has been called, a Value must not be copied.
//
// A Value must not be copied after first use.
type Value struct {
	v      *atomic.Value
	inited uint32
}

// NewValue returns a new Value.
func NewValue(val string) *Value {
	addr := &Value{v: &atomic.Value{}}
	addr.Store(val)
	return addr
}

// Load returns the value set by the most recent Store.
// It returns nil if there has been no call to Store for this Value.
func (v *Value) Load() (x interface{}) {
	if v.v == nil {
		return nil
	}
	return v.v.Load()
}

// Store sets the value of the Value to x.
// All calls to Store for a given Value must use values of the same concrete type.
// Store of an inconsistent type panics, as does Store(nil).
func (v *Value) Store(x interface{}) {
	if v.v == nil {
		v.init()
	}
	v.v.Store(x)
}

func (v *Value) init() {
	for {
		if v.v != nil {
			break
		}
		if atomic.CompareAndSwapUint32(&v.inited, 0, 1) {
			v.v = &atomic.Value{}
		}
	}
}
