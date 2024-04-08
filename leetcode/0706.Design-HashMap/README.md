# 706. Design HashMap

<https://leetcode.com/problems/design-hashmap/>

需要先了解 hash map 的底層是如何運作

- hash function  
  決定 key 值要透過什麼演算法進行 hash  
  已知題目的 key 資料型態限定為 int  
  所以 hash 的方式採用對質數取餘數 modulo 的方式  
  至於要用哪個質數，由於題目已限定 key 介於 $0 ～ 10^6$ 之間  
  所以我個人隨便根據質數表取個 `2131`
  這樣即表示，至多會有 `2131` 個 bucket
- collision handling  
  當發生碰撞時，處理機制為何  
  處理方式大致可以分為兩類，Open Addressing 和 Separate Chaining

  Open Addressing: Linear Probing, Quadratic Probing, Double Hashing...  
  Separate Chaining: bucket(linked list)

  決定採用 Separate Chaining 的方式  
  因此 hash map 的本質將會是 []Linked List
- hash map 的空間擴張及減縮  
  就像 Dynamic Sized Arrays 一樣的概念  
  當 hash map 的內容越多時就需要考慮宣告空間的擴充  
  當內容刪除時，所宣告的空間需要釋放  
  但看了 leetcode 裡的討論區，似乎較少討論這件事  
  因此忽略這件事

有了以上這些知識之後便能開始進行實作

## Takeaway

- hash table 的底層運作
