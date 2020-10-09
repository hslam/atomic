package atomic

import (
	"sync"
	"testing"
)

func TestValue(t *testing.T) {
	var val = "Hello World"
	var addFunc AddFunc = func(old, delta interface{}) (new interface{}) {
		return old.(string) + delta.(string)
	}
	addr := NewValue(val, addFunc)
	if addr.Load() != val {
		t.Error(addr.Load())
	}
	addr.Store(val[:5])
	if addr.Load().(string) != val[:5] {
		t.Error(addr.Load())
	}
	var delta = val[5:]
	if addr.Add(delta) != val {
		t.Error(addr.Load())
	}
	if addr.Load() != val {
		t.Error(addr.Load())
	}
	var new = "Foo"
	if addr.Swap(new) != val {
		t.Error(addr.Load())
	}
	var old = new
	new = "Bar"
	if !addr.CompareAndSwap(old, new) {
		t.Error(addr.Load())
	}
	if addr.CompareAndSwap(old, new) {
		t.Error(addr.Load())
	}

	addr = &Value{}
	if addr.Load() != nil {
		t.Error(addr.Load())
	}
}

func TestAddValue(t *testing.T) {
	var addFunc AddFunc = func(old, delta interface{}) (new interface{}) {
		return old.(string) + delta.(string)
	}
	addr := NewValue("", nil)
	func() {
		defer func() {
			if err := recover(); err == nil {
				t.Error("should panic")
			}
		}()
		addr.Add("")
	}()
	addr.AddFunc = addFunc
	var wg sync.WaitGroup
	for i := 0; i < 16382; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			addr.Add("")
		}()
	}
	wg.Wait()
}

func TestCompareAndSwapValue(t *testing.T) {
	addr := NewValue("", nil)
	var wg sync.WaitGroup
	for i := 0; i < 16382; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			addr.CompareAndSwap("", "")
		}()
	}
	wg.Wait()
}

func TestSwapValue(t *testing.T) {
	addr := NewValue("", nil)
	var wg sync.WaitGroup
	for i := 0; i < 16382; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			addr.Swap("")
		}()
	}
	wg.Wait()
}

func BenchmarkSwapValue(b *testing.B) {
	addr := NewValue("", nil)
	for i := 0; i < b.N; i++ {
		addr.Swap("")
	}
}

func BenchmarkCompareAndSwapValue(b *testing.B) {
	addr := NewValue("", nil)
	for i := 0; i < b.N; i++ {
		addr.CompareAndSwap("", "")
	}
}

func BenchmarkAddValue(b *testing.B) {
	var addFunc AddFunc = func(old, delta interface{}) (new interface{}) {
		return old.(string) + delta.(string)
	}
	addr := NewValue("", addFunc)
	for i := 0; i < b.N; i++ {
		addr.Add("")
	}
}

func BenchmarkStoreValue(b *testing.B) {
	addr := NewValue("", nil)
	for i := 0; i < b.N; i++ {
		addr.Store("")
	}
}

func BenchmarkLoadValue(b *testing.B) {
	addr := NewValue("", nil)
	for i := 0; i < b.N; i++ {
		addr.Load()
	}
}
