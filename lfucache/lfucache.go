package lfucache

import "container/list"

type kvc struct {
    key, val, cnt int
}

func incCnt(elem *list.Element) {
    elem.Value = kvc{elem.Value.(kvc).key, elem.Value.(kvc).val, elem.Value.(kvc).cnt + 1}
}

func chgVal(elem *list.Element, value int) {
    elem.Value = kvc{elem.Value.(kvc).key, value, elem.Value.(kvc).cnt}
}

type LFUCache struct {
    table   map[int]*list.Element
    elems   *list.List
    maxSize int
}

func Constructor(capacity int) LFUCache {
    return LFUCache{
        table:   make(map[int]*list.Element),
        elems:   list.New(),
        maxSize: capacity}
}

func (lfu *LFUCache) Get(key int) int {
    elem, used := lfu.table[key]
    if !used {
        return -1
    }
    incCnt(elem)
    if elem != lfu.elems.Front() && elem.Prev().Value.(kvc).cnt <= elem.Value.(kvc).cnt {
        lfu.elems.MoveBefore(elem, elem.Prev())
    }
    return elem.Value.(kvc).val
}

func (lfu *LFUCache) Put(key int, value int) {
    if lfu.maxSize == 0 {
        return
    }
    elem, used := lfu.table[key]
    if used {
        chgVal(elem, value)
        incCnt(elem)
    } else {
        if lfu.elems.Len() == lfu.maxSize {
            elem = lfu.elems.Back()
            k := elem.Value.(kvc).key
            lfu.elems.Remove(elem)
            delete(lfu.table, k)
        }
        elem = lfu.elems.PushBack(kvc{key, value, 0})
        lfu.table[key] = elem
    }
    if elem != lfu.elems.Front() && elem.Prev().Value.(kvc).cnt <= elem.Value.(kvc).cnt {
        lfu.elems.MoveBefore(elem, elem.Prev())
    } else if elem != lfu.elems.Back() && elem.Prev().Value.(kvc).cnt > elem.Value.(kvc).cnt {
        lfu.elems.MoveAfter(elem, elem.Prev())
    }
}
