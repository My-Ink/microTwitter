package tests

import (
    "Solutions/lfucache"
    "testing"
)

func TestLFUCacheSequential(t *testing.T) {
    t.Run("PutGet", func(t *testing.T) {
        lfu := lfucache.Constructor(1)
        lfu.Put(1, 1)
        AssertEqual(t, lfu.Get(1), 1)
    })
    t.Run("GetFromEmpty", func(t *testing.T) {
        lfu := lfucache.Constructor(100)
        AssertEqual(t, lfu.Get(1), -1)
    })
    t.Run("ZeroCapacity", func(t *testing.T) {
        lfu := lfucache.Constructor(0)
        lfu.Put(1, 1)
        AssertEqual(t, lfu.Get(1), -1)
    })
    t.Run("LeastRecentlyUsed", func(t *testing.T) {
        lfu := lfucache.Constructor(2)
        lfu.Put(1, 1)
        lfu.Put(2, 2)
        lfu.Put(3, 3)
        AssertEqual(t, lfu.Get(2), 2)
        AssertEqual(t, lfu.Get(1), -1)
    })
    t.Run("LeastFrequentlyUsed", func(t *testing.T) {
        lfu := lfucache.Constructor(2)
        lfu.Put(1, 1)
        lfu.Put(2, 2)
        for i := 0; i < 100; i++ {
            _ = lfu.Get(1)
        }
        for i := 0; i < 99; i++ {
            _ = lfu.Get(2)
        }
        lfu.Put(3, 3)
        AssertEqual(t, lfu.Get(2), -1)
        AssertEqual(t, lfu.Get(1), 1)
    })
    t.Run("Test 1", func(t *testing.T) {
        lfu := lfucache.Constructor(3)
        lfu.Put(1, 1)
        lfu.Put(2, 2)
        lfu.Put(3, 3)
        lfu.Put(4, 4)
        AssertEqual(t, lfu.Get(4), 4)
        AssertEqual(t, lfu.Get(3), 3)
        AssertEqual(t, lfu.Get(2), 2)
        AssertEqual(t, lfu.Get(1), -1)
        lfu.Put(5, 5)
        AssertEqual(t, lfu.Get(1), -1)
        AssertEqual(t, lfu.Get(2), 2)
        AssertEqual(t, lfu.Get(3), 3)
        AssertEqual(t, lfu.Get(4), -1)
        AssertEqual(t, lfu.Get(5), 5)
    })
    t.Run("Test 2", func(t *testing.T) {
        lfu := lfucache.Constructor(3)
        lfu.Put(2, 2)
        lfu.Put(1, 1)
        AssertEqual(t, lfu.Get(2), 2)
        AssertEqual(t, lfu.Get(1), 1)
        AssertEqual(t, lfu.Get(2), 2)
        lfu.Put(3, 3)
        lfu.Put(4, 4)
        AssertEqual(t, lfu.Get(3), -1)
        AssertEqual(t, lfu.Get(2), 2)
        AssertEqual(t, lfu.Get(1), 1)
        AssertEqual(t, lfu.Get(4), 4)
    })
}
