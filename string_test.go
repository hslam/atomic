package atomic

import (
	"testing"
)

func TestString(t *testing.T) {
	var val = "Hello World"
	str := NewString(val)
	if str.Load() != val {
		t.Error(str.Load())
	}
	val = "Hello"
	str.Store(val)
	if str.Load() != val {
		t.Error(str.Load())
	}
	var delta = " World"
	if str.Add(delta) != val+delta {
		t.Error(str.Load())
	}
	if str.Load() != val+delta {
		t.Error(str.Load())
	}
	var new = "Foo"
	if str.Swap(new) != val+delta {
		t.Error(str.Load())
	}
	var old = new
	new = "Bar"
	if !str.CompareAndSwap(old, new) {
		t.Error(str.Load())
	}
}
