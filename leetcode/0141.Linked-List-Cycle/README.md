# 141. Linked List Cycle

<https://leetcode.com/problems/linked-list-cycle/>

建立快慢指針，慢指針一次走一步，快指針一次走二步  
若有環，則快指針會繞回慢指針前面，並且一步一步追上慢指針  
若慢指針最後遇到 `null`，則就表示沒有環

## Takeaway

- Linked List
- 快慢指針
