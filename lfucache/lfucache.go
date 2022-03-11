package lfucache

import (
    "container/list"
    "fmt"
)

// LFU Cache
////////////////////////////////////////////////////////////////////////////////////////////////////

// Public API

type LFUCache struct {
    // Key --> node mapping
    table map[int]*list.Element

    // List of lists
    buckets *list.List

    // Capacity
    maxSize int
}

func Constructor(capacity int) LFUCache {
    return LFUCache{
        make(map[int]*list.Element),
        list.New(),
        capacity}
}

func (lfu *LFUCache) Get(key int) int {
    elem, contains := lfu.table[key]
    if !contains {
        return -1
    }
    rankedUp := lfu.rankUp(elem)
    lfu.table[key] = rankedUp
    return getLRUNode(rankedUp).value
}

func (lfu *LFUCache) Put(key int, value int) {
    if lfu.maxSize == 0 {
        return
    }
    elem, contains := lfu.table[key]
    if contains {
        lfu.table[key] = lfu.rankUp(modify(elem, value))
        return
    }
    if len(lfu.table) == lfu.maxSize {
        lastBucket := lfu.buckets.Back()
        lastLRU := getBucket(lastBucket).lru
        toRemove := lastLRU.Back()
        delete(lfu.table, getLRUNode(toRemove).key)
        if lastLRU.Len() == 1 {
            lfu.buckets.Remove(lastBucket)
        } else {
            lastLRU.Remove(toRemove)
        }
    }
    minRankBucket := lfu.buckets.Back()
    if lfu.buckets.Len() == 0 || getBucket(minRankBucket).rank > 0 {
        minRankBucket = lfu.buckets.PushBack(newBucket(0))
    }
    lfu.table[key] = getBucket(minRankBucket).lru.PushFront(newLRUNode(key, value, minRankBucket))
}

// Private implementation

func (lfu *LFUCache) rankUp(elem *list.Element) *list.Element {
    node := getLRUNode(elem)
    nextBucket := lfu.nextBucket(node.bucket)
    newElem := getBucket(nextBucket).lru.PushFront(newLRUNode(node.key, node.value, nextBucket))
    lru := getBucket(node.bucket).lru
    if lru.Len() == 1 {
        lfu.buckets.Remove(node.bucket)
    } else {
        lru.Remove(elem)
    }
    return newElem
}

func modify(elem *list.Element, newValue int) *list.Element {
    node, lru := getLRUNode(elem), getLRU(elem)
    modified := lru.PushFront(newLRUNode(node.key, newValue, node.bucket))
    lru.Remove(elem)
    return modified
}

func (lfu *LFUCache) nextBucket(bucketNode *list.Element) *list.Element {
    nextFreq, nextBucketNode := getBucket(bucketNode).rank+1, bucketNode.Prev()
    if nextBucketNode != nil && getBucket(nextBucketNode).rank == nextFreq {
        return bucketNode.Prev()
    }
    return lfu.buckets.InsertBefore(newBucket(nextFreq), bucketNode)
}

// Logging

func (lfu *LFUCache) printState() {
    for bucket := lfu.buckets.Front(); bucket != nil; bucket = bucket.Next() {
        fmt.Printf("%d: ", getBucket(bucket).rank)
        last := getBucket(bucket).lru.Back()
        for node := getBucket(bucket).lru.Front(); node != nil; node = node.Next() {
            fmt.Printf("(%d, %d)", getLRUNode(node).key, getLRUNode(node).value)
            if node != last {
                fmt.Printf("-->")
            }
        }
        fmt.Printf("\n")
    }
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// LRU Bucket
////////////////////////////////////////////////////////////////////////////////////////////////////
type bucketNode struct {
    lru  *list.List
    rank uint64
}

func newBucket(rank uint64) bucketNode {
    return bucketNode{
        lru:  list.New(),
        rank: rank}
}

func getBucket(node *list.Element) bucketNode {
    return node.Value.(bucketNode)
}

func getLRU(elem *list.Element) *list.List {
    return getBucket(getLRUNode(elem).bucket).lru
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// LRU Node
////////////////////////////////////////////////////////////////////////////////////////////////////
type lruNode struct {
    key   int
    value int

    // Pointer to bucket
    bucket *list.Element
}

func newLRUNode(key, value int, bucket *list.Element) lruNode {
    return lruNode{
        key:    key,
        value:  value,
        bucket: bucket}
}

func getLRUNode(elem *list.Element) lruNode {
    return elem.Value.(lruNode)
}

////////////////////////////////////////////////////////////////////////////////////////////////////
