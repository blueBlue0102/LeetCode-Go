# 142. Linked List Cycle II

<https://leetcode.com/problems/linked-list-cycle-ii/>

暴力解是建立一個 hash table，將遍歷過的節點都記錄下來  
但這樣時間和空間複雜度都是 $O(n)$

最佳解是透過 [Floyd's tortoise and hare](https://en.wikipedia.org/wiki/Cycle_detection#Floyd's_tortoise_and_hare) 演算法

藉由快慢指針，快指針一次走兩步，慢指針一次走一步  
若有環，則相當於快指針將會從慢指針的後面，在每次的 iterate 中會追上一步  
因此若有環，則快慢指針最終會相遇

...  
由於過於複雜，可以直接參考網路上其他人的說明  
<https://github.com/youngyangyang04/leetcode-master/blob/master/problems/0142.%E7%8E%AF%E5%BD%A2%E9%93%BE%E8%A1%A8II.md>

最佳解的時間複雜度仍然是 $O(n)$，但空間複雜度是 $O(1)$

## Takeaway

- Floyd's tortoise and hare 演算法
