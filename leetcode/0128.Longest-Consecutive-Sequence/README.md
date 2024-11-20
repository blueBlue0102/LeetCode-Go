# 128. Longest Consecutive Sequence

<https://leetcode.com/problems/longest-consecutive-sequence/>

給定一個未排序的整數陣列  
找出其中最長的連續數字長度

例如給定 `[100,4,200,1,3,2]`  
裡面最長的連續數列是 `[1, 2, 3, 4]`，所以 output 要輸出 `4`

並且，題目要求在時間複雜度 $O(n)$ 解出來

先 assign 一個 hash table 將所有數字儲存起來  
這個操作的時間複雜度和空間複雜度都是 $O(n)$  

接著開始輪巡這個 table  
每當取得一個數字後，就開始從 table 中連續的尋找是否有再大 1 或再小 1 的數字存在  
只要找到了，就在 table 中註記該數字已被找到  
這樣的操作是在輪巡 table，table 的長度取決於給定數列的大小  
所以時間複雜度是 $O(n)$

總共所花的時間是 $O(n+n)=O(2n)=O(n)$

## Takeaway

- Hash Table
