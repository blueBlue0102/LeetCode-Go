# 75. Sort Colors

<https://leetcode.com/problems/sort-colors/>

這是一個 Dutch National Flag Algorithm 的題目  
題目已限縮數字侷限在 `[0, 1, 2]`  
設計三個指標，`left`, `right`, `curr`  
遇到 `0` 時將 `nums[curr]` 與 `nums[left]` 進行 swap；  
遇到 `2` 時將 `nums[curr]` 與 `nums[right]` 進行 swap；

如此便能以時間複雜度 $O(n)$ 的方式將 `nums` 中的數字進行排列

## Takeaway


