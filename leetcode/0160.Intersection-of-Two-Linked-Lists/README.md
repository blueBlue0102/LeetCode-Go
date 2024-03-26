# 160. Intersection of Two Linked Lists

<https://leetcode.com/problems/intersection-of-two-linked-lists/>

暴力解可以使用 hash table  
但這樣的時間複雜度是 $O(M + N)$，空間複雜度是 $O(M)$

而透過雙指針的話  
雖然時間複雜度仍然是 $O(M + N)$，但空間複雜度可以縮減成 $O(1)$

在 headA 和 headB 各放置一個指針  
假設 A 和 B 之間有交集  
且 headA 到交集點之間有 a 個節點；headB 到交集點之間有 b 個節點  
交集點到尾端間有 c 個節點

$(a + c) + b = (b + c) + a$  
$(a + c) + b$ 的意思是 headA 的指針從 A 的頭走到尾之後，再從 headB 開始出發  
$(b + c) + a$ 的意思是 headB 的指針從 B 的頭走到尾之後，再從 headA 開始出發  
兩個指針同時進行，當兩個指針都開始從另一個 List 開始重頭走時，接下來各指針的每一步，都有可能是交集節點  
假如任一指針已經走到另一個 List 的尾端但仍找不到交集時，便是兩個 List 就是沒有交集  
（兩個指針會同時在另一個 List 的最後指向 null）

## Takeaway

- 雙指針
