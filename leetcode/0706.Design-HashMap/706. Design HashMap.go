package leetcode

const BUCKETS_SIZE = 2131

type MyListNode struct {
	OriginalKey int
	Val         int
	Next        *MyListNode
}

type Bucket *MyListNode

type MyHashMap struct {
	buckets []Bucket
}

// 負責將傳入 key 進行 hash 後得出映射位置
func HashFunction(key int) int {
	return key % BUCKETS_SIZE
}

func Constructor() MyHashMap {
	return MyHashMap{buckets: make([]Bucket, BUCKETS_SIZE)}
}

func (myHashMap *MyHashMap) Put(key int, value int) {
	hashKey := HashFunction(key)
	prev := &MyListNode{Next: myHashMap.buckets[hashKey]}

	// 嘗試搜尋是否有相對應 OriginalKey 的節點
	// 若有則更新該節點而非新增節點
	for prev.Next != nil {
		if prev.Next.OriginalKey == key {
			prev.Next.Val = value
			return
		} else {
			prev = prev.Next
		}
	}
	// 沒找到相對應 OriginalKey，所以要新增節點
	prev.Next = &MyListNode{OriginalKey: key, Val: value}
	if myHashMap.buckets[hashKey] == nil {
		myHashMap.buckets[hashKey] = prev.Next
	}
}

// 找不到時回傳 -1
func (myHashMap *MyHashMap) Get(key int) int {
	hashKey := HashFunction(key)
	pointer := myHashMap.buckets[hashKey]

	for pointer != nil {
		if pointer.OriginalKey == key {
			return pointer.Val
		} else {
			pointer = pointer.Next
		}
	}
	return -1
}

func (myHashMap *MyHashMap) Remove(key int) {
	hashKey := HashFunction(key)
	prev := &MyListNode{Next: myHashMap.buckets[hashKey]}

	for prev.Next != nil {
		if prev.Next.OriginalKey == key {
			if prev.Next == myHashMap.buckets[hashKey] {
				myHashMap.buckets[hashKey] = prev.Next.Next
			} else {
				prev.Next = prev.Next.Next
			}
			return
		} else {
			prev = prev.Next
		}
	}
}
