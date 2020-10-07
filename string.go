package atomic

import (
	"sync/atomic"
)

// String represents an string.
type String struct {
	v       *atomic.Value
	seting  uint32
	initing uint32
}

// NewString returns a new String.
func NewString(val string) *String {
	addr := &String{v: &atomic.Value{}}
	addr.Store(val)
	return addr
}

// Swap atomically stores new into *addr and returns the previous *addr value.
func (addr *String) Swap(new string) (old string) {
	return SwapString(addr, new)
}

// CompareAndSwap executes the compare-and-swap operation for an string value.
func (addr *String) CompareAndSwap(old, new string) (swapped bool) {
	return CompareAndSwapString(addr, old, new)
}

// Add atomically adds delta to *addr and returns the new value.
func (addr *String) Add(delta string) (new string) {
	return AddString(addr, delta)
}

// Load atomically loads *addr.
func (addr *String) Load() (val string) {
	return LoadString(addr)
}

// Store atomically stores val into *addr.
func (addr *String) Store(val string) {
	StoreString(addr, val)
}

// SwapString atomically stores new into *addr and returns the previous *addr value.
func SwapString(addr *String, new string) (old string) {
	for {
		if !atomic.CompareAndSwapUint32(&addr.seting, 0, 1) {
			continue
		}
		old = LoadString(addr)
		StoreString(addr, new)
		atomic.StoreUint32(&addr.seting, 0)
		return
	}
}

// CompareAndSwapString executes the compare-and-swap operation for an string value.
func CompareAndSwapString(addr *String, old, new string) (swapped bool) {
	for {
		if !atomic.CompareAndSwapUint32(&addr.seting, 0, 1) {
			continue
		}
		if LoadString(addr) == old {
			StoreString(addr, new)
			atomic.StoreUint32(&addr.seting, 0)
			return true
		}
		atomic.StoreUint32(&addr.seting, 0)
		return false
	}
}

// AddString atomically adds delta to *addr and returns the new value.
func AddString(addr *String, delta string) (new string) {
	for {
		if !atomic.CompareAndSwapUint32(&addr.seting, 0, 1) {
			continue
		}
		new = LoadString(addr) + delta
		StoreString(addr, new)
		atomic.StoreUint32(&addr.seting, 0)
		return
	}
}

// LoadString atomically loads *addr.
func LoadString(addr *String) (val string) {
	if addr.v == nil {
		return ""
	}
	var ok bool
	if val, ok = addr.v.Load().(string); ok {
		return val
	}
	return ""
}

// StoreString atomically stores val into *addr.
func StoreString(addr *String, val string) {
	if addr.v == nil {
		initString(addr)
	}
	addr.v.Store(val)
}

func initString(addr *String) {
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
