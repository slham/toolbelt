package cache

import (
	"testing"
)

func TestNewLRUCache_BadCapacity(t *testing.T) {
	_, err := NewLRUCache(0)
	if err == nil {
		t.Error("LRUCache should not be created with a capacity less than 1.")
	}
}

func TestNewLRUCache_HappyPath(t *testing.T) {
	lruCache, err := NewLRUCache(2)
	if err != nil {
		t.Error(err)
	}

	if lruCache == nil {
		t.Error("LRUCache not instantiated properly.")
	}
}

func TestSet(t *testing.T) {
	lruCache, err := NewLRUCache(2)
	if err != nil {
		t.Error(err)
	}
	lruCache.Set("1", "alpha")
	lruCache.Set("2", "bravo")
	lruCache.Set("3", "charlie")
	lruCache.Set("4", "delta")

	if lruCache.Head.Key != "4" {
		t.Error("Head not being properly assigned.")
	}

	if lruCache.Tail.Key != "3" {
		t.Error("Tail not being properly assigned.")
	}

	if lruCache.Head.Value != "delta" || lruCache.Tail.Value != "charlie" {
		t.Error("Values not being properly assigned.")
	}

}

func TestGet(t *testing.T) {
	lruCache, err := NewLRUCache(2)
	if err != nil {
		t.Error(err)
	}

	lruCache.Set("1", "alpha")

	if lruCache.Get("1") != "alpha" {
		t.Error("Lookup is failing.")
	}

	if lruCache.Get("11") != nil {
		t.Error("Should not have found anything.")
	}
}
