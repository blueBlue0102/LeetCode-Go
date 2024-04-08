package leetcode

import "testing"

// 以下測試由 ChatGPT 生成後修改

func TestMyHashMap_Put(t *testing.T) {
	// 初始化 MyHashMap
	myMap := Constructor()

	// 測試 Put 方法
	myMap.Put(1, 10)
	myMap.Put(2, 20)
	myMap.Put(3, 30)
	myMap.Put(1+BUCKETS_SIZE, 40)
	myMap.Put(2+BUCKETS_SIZE, 50)
	myMap.Put(3+BUCKETS_SIZE, 60)

	// 驗證 Put 方法是否正確
	if myMap.buckets[HashFunction(1)] == nil || myMap.buckets[HashFunction(1)].Val != 10 {
		t.Errorf("Error: Expected value 10 for key 1, got %d", myMap.buckets[HashFunction(1)].Val)
	}
	if myMap.buckets[HashFunction(2)] == nil || myMap.buckets[HashFunction(2)].Val != 20 {
		t.Errorf("Error: Expected value 20 for key 2, got %d", myMap.buckets[HashFunction(2)].Val)
	}
	if myMap.buckets[HashFunction(3)] == nil || myMap.buckets[HashFunction(3)].Val != 30 {
		t.Errorf("Error: Expected value 30 for key 3, got %d", myMap.buckets[HashFunction(3)].Val)
	}
	if myMap.buckets[HashFunction(1+BUCKETS_SIZE)] == nil || myMap.buckets[HashFunction(1+BUCKETS_SIZE)].Next == nil || myMap.buckets[HashFunction(1+BUCKETS_SIZE)].Next.Val != 40 {
		t.Errorf("Error: Expected value 40 for key 1+BUCKETS_SIZE, got %d", myMap.buckets[HashFunction(1+BUCKETS_SIZE)].Next.Val)
	}
	if myMap.buckets[HashFunction(2+BUCKETS_SIZE)] == nil || myMap.buckets[HashFunction(2+BUCKETS_SIZE)].Next == nil || myMap.buckets[HashFunction(2+BUCKETS_SIZE)].Next.Val != 50 {
		t.Errorf("Error: Expected value 50 for key 2+BUCKETS_SIZE, got %d", myMap.buckets[HashFunction(2+BUCKETS_SIZE)].Next.Val)
	}
	if myMap.buckets[HashFunction(3+BUCKETS_SIZE)] == nil || myMap.buckets[HashFunction(3+BUCKETS_SIZE)].Next == nil || myMap.buckets[HashFunction(3+BUCKETS_SIZE)].Next.Val != 60 {
		t.Errorf("Error: Expected value 60 for key 3+BUCKETS_SIZE, got %d", myMap.buckets[HashFunction(3+BUCKETS_SIZE)].Next.Val)
	}
}

func TestMyHashMap_Get(t *testing.T) {
	// 初始化 MyHashMap
	myMap := Constructor()

	// 測試 Get 方法
	myMap.Put(1, 10)
	myMap.Put(2, 20)
	myMap.Put(3, 30)
	myMap.Put(1+BUCKETS_SIZE, 40)
	myMap.Put(2+BUCKETS_SIZE, 50)
	myMap.Put(3+BUCKETS_SIZE, 60)

	// 驗證 Get 方法是否正確
	if myMap.Get(1) != 10 {
		t.Errorf("Error: Expected value 10 for key 1, got %d", myMap.Get(1))
	}
	if myMap.Get(4) != -1 {
		t.Errorf("Error: Expected value -1 for key 4, got %d", myMap.Get(4))
	}
	if myMap.Get(1+BUCKETS_SIZE) != 40 {
		t.Errorf("Error: Expected value 40 for key 1+BUCKETS_SIZE, got %d", myMap.Get(1+BUCKETS_SIZE))
	}
}

func TestMyHashMap_Remove(t *testing.T) {
	// 初始化 MyHashMap
	myMap := Constructor()

	// 測試 Remove 方法
	myMap.Put(1, 10)
	myMap.Put(2, 20)
	myMap.Put(3, 30)

	myMap.Remove(2)

	// 驗證 Remove 方法是否正確
	if myMap.buckets[HashFunction(2)] != nil {
		t.Errorf("Error: Expected value nil for key 2 after removal, got %d", myMap.buckets[HashFunction(2)].Val)
	}
}
