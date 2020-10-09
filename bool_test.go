// Copyright (c) 2020 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

package atomic

import (
	"sync"
	"testing"
)

func TestBool(t *testing.T) {
	addr := NewBool(false)
	if addr.Load() != false {
		t.Error(addr.Load())
	}

	addr.Store(true)
	if addr.Load() != true {
		t.Error(addr.Load())
	}
	var delta = true
	if addr.Add(delta) != true {
		t.Error(addr.Load())
	}
	if addr.Load() != true {
		t.Error(addr.Load())
	}
	var new = false
	if addr.Swap(new) != true {
		t.Error(addr.Load())
	}
	var old = new
	new = true
	if !addr.CompareAndSwap(old, new) {
		t.Error(addr.Load())
	}
	if addr.CompareAndSwap(old, new) {
		t.Error(addr.Load())
	}
}

func TestAddBool(t *testing.T) {
	addr := NewBool(false)
	var wg sync.WaitGroup
	for i := 0; i < 8192; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			addr.Add(false)
		}()
	}
	wg.Wait()
}

func TestCompareAndSwapBool(t *testing.T) {
	addr := NewBool(false)
	var wg sync.WaitGroup
	for i := 0; i < 8192; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			addr.CompareAndSwap(false, true)
		}()
	}
	wg.Wait()
}

func TestSwapBool(t *testing.T) {
	addr := NewBool(false)
	var wg sync.WaitGroup
	for i := 0; i < 8192; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			addr.Swap(false)
		}()
	}
	wg.Wait()
}

func BenchmarkSwapBool(b *testing.B) {
	addr := NewBool(false)
	for i := 0; i < b.N; i++ {
		addr.Swap(false)
	}
}

func BenchmarkCompareAndSwapBool(b *testing.B) {
	addr := NewBool(false)
	for i := 0; i < b.N; i++ {
		addr.CompareAndSwap(false, true)
	}
}

func BenchmarkAddBool(b *testing.B) {
	addr := NewBool(false)
	for i := 0; i < b.N; i++ {
		addr.Add(false)
	}
}

func BenchmarkStoreBool(b *testing.B) {
	addr := NewBool(false)
	for i := 0; i < b.N; i++ {
		addr.Store(false)
	}
}

func BenchmarkLoadBool(b *testing.B) {
	addr := NewBool(false)
	for i := 0; i < b.N; i++ {
		addr.Load()
	}
}
