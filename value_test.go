// Copyright (c) 2020 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

package atomic

import (
	"sync"
	"testing"
)

func TestValue(t *testing.T) {
	var val = "Hello World"
	var equalFunc EqualFunc = func(old, load interface{}) (equal bool) {
		return old == load
	}
	var addFunc AddFunc = func(old, delta interface{}) (new interface{}) {
		return old.(string) + delta.(string)
	}
	addr := NewValue(val, equalFunc, addFunc)
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
	var old = addr.Load()
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
	var equalFunc EqualFunc = func(old, load interface{}) (equal bool) {
		return old == load
	}
	var addFunc AddFunc = func(old, delta interface{}) (new interface{}) {
		return old.(string) + delta.(string)
	}
	addr := NewValue("", equalFunc, nil)
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

func testCompareAndSwapValue(t *testing.T) {
	var equalFunc EqualFunc = func(old, load interface{}) (equal bool) {
		return old == load
	}
	addr := &Value{EqualFunc: equalFunc}
	var wg sync.WaitGroup
	for i := 0; i < 64; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			addr.compareAndSwap("", "")
		}()
	}
	wg.Wait()
}

func TestCompareAndSwapValue(t *testing.T) {
	for i := 0; i < 16382; i++ {
		testCompareAndSwapValue(t)
	}
	addr := &Value{}
	func() {
		defer func() {
			if err := recover(); err == nil {
				t.Error("should panic")
			}
		}()
		addr.CompareAndSwap("", "")
	}()
	var equalFunc EqualFunc = func(old, load interface{}) (equal bool) {
		return old == load
	}
	addr.EqualFunc = equalFunc
	func() {
		defer func() {
			if err := recover(); err == nil {
				t.Error("should panic")
			}
		}()
		addr.compareAndSwap("", nil)
	}()
	addr.compareAndSwap("", "")
	func() {
		defer func() {
			if err := recover(); err == nil {
				t.Error("should panic")
			}
		}()
		addr.compareAndSwap(nil, "")
	}()
	func() {
		defer func() {
			if err := recover(); err == nil {
				t.Error("should panic")
			}
		}()
		addr.compareAndSwap(1, "")
	}()
	func() {
		defer func() {
			if err := recover(); err == nil {
				t.Error("should panic")
			}
		}()
		addr.compareAndSwap("", 1)
	}()
}

func TestSwapValue(t *testing.T) {
	var equalFunc EqualFunc = func(old, load interface{}) (equal bool) {
		return old == load
	}
	addr := NewValue("", equalFunc, nil)
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
	var equalFunc EqualFunc = func(old, load interface{}) (equal bool) {
		return old == load
	}
	addr := NewValue("", equalFunc, nil)
	for i := 0; i < b.N; i++ {
		addr.Swap("")
	}
}

func BenchmarkCompareAndSwapValue(b *testing.B) {
	var equalFunc EqualFunc = func(old, load interface{}) (equal bool) {
		return old == load
	}
	addr := NewValue("", equalFunc, nil)
	for i := 0; i < b.N; i++ {
		addr.CompareAndSwap("", "")
	}
}

func BenchmarkAddValue(b *testing.B) {
	var equalFunc EqualFunc = func(old, load interface{}) (equal bool) {
		return old == load
	}
	var addFunc AddFunc = func(old, delta interface{}) (new interface{}) {
		return old.(string) + delta.(string)
	}
	addr := NewValue("", equalFunc, addFunc)
	for i := 0; i < b.N; i++ {
		addr.Add("")
	}
}

func BenchmarkStoreValue(b *testing.B) {
	addr := NewValue("", nil, nil)
	for i := 0; i < b.N; i++ {
		addr.Store("")
	}
}

func BenchmarkLoadValue(b *testing.B) {
	addr := NewValue("", nil, nil)
	for i := 0; i < b.N; i++ {
		addr.Load()
	}
}
