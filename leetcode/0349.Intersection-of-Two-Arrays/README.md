# 349. Intersection of Two Arrays

<https://leetcode.com/problems/intersection-of-two-arrays/>

只需要對其中一個 array 建立 hash set 之後，再來輪巡另一個 array，每當該數字於 hast set 中出現便放入 result  
當放入 result 時，須更新刪除原來的 hash set，以免重覆的數字被放入

## Takeaway

無
