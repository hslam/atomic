// Copyright (c) 2020 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

package atomic

import (
	"sync"
	"testing"
)

func TestString(t *testing.T) {
	var val = "Hello World"
	addr := NewString(val)
	if addr.Load() != val {
		t.Error(addr.Load())
	}
	addr.Store(val[:5])
	if addr.Load() != val[:5] {
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

	addr = &String{}
	if addr.Load() != "" {
		t.Error(addr.Load())
	}
}

func TestAddString(t *testing.T) {
	addr := NewString("")
	var wg sync.WaitGroup
	for i := 0; i < 8192; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			addr.Add("")
		}()
	}
	wg.Wait()
}

func TestCompareAndSwapString(t *testing.T) {
	addr := NewString("")
	var wg sync.WaitGroup
	for i := 0; i < 8192; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			addr.CompareAndSwap("", "")
		}()
	}
	wg.Wait()
}

func TestSwapString(t *testing.T) {
	addr := NewString("")
	var wg sync.WaitGroup
	for i := 0; i < 8192; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			addr.Swap("")
		}()
	}
	wg.Wait()
}

func BenchmarkSwapString(b *testing.B) {
	addr := NewString("")
	for i := 0; i < b.N; i++ {
		addr.Swap("")
	}
}

func BenchmarkCompareAndSwapString(b *testing.B) {
	addr := NewString("")
	for i := 0; i < b.N; i++ {
		addr.CompareAndSwap("", "")
	}
}

func BenchmarkAddString(b *testing.B) {
	addr := NewString("")
	for i := 0; i < b.N; i++ {
		addr.Add("")
	}
}

func BenchmarkStoreString(b *testing.B) {
	addr := NewString("")
	for i := 0; i < b.N; i++ {
		addr.Store("")
	}
}

func BenchmarkLoadString(b *testing.B) {
	addr := NewString("")
	for i := 0; i < b.N; i++ {
		addr.Load()
	}
}
