package chilru

import (
	"testing"
)

func TestLru_Add(t *testing.T) {
	lru := New[string, int](2)
	ok := lru.Add("a", 1)
	if !ok {
		t.Fail()
		return
	}
	ok = lru.Add("b", 2)
	if !ok {
		t.Fail()
		return
	}
	ok = lru.Add("b", 3)
	if ok {
		t.Fail()
		return
	}

	v, ok := lru.Get("a")
	if !ok && v != 1 {
		t.Fail()
		return
	}

	v, ok = lru.Get("b")
	if !ok && v != 2 {
		t.Fail()
		return
	}
}

func TestLru_Lru(t *testing.T) {
	lru := New[string, int](2)
	ok := lru.Add("a", 1)
	if !ok {
		t.Fail()
		return
	}
	ok = lru.Add("b", 2)
	if !ok {
		t.Fail()
		return
	}
	ok = lru.Add("c", 3)
	if !ok {
		t.Fail()
		return
	}

	lru.Range(func(k, v any) bool {
		t.Log("key", k)
		t.Log("value", v)
		return true
	})

}

func TestLru_Set(t *testing.T) {
	lru := New[string, int](2)
	lru.Set("a", 1)

	lru.Set("b", 2)

	lru.Set("b", 3)

	lru.Range(func(k, v any) bool {
		t.Log("key", k)
		t.Log("value", v)
		return true
	})
}

func BenchmarkLru_Add(b *testing.B) {
	lru := New[string, int](10240)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lru.Add("a", i)
	}
}
