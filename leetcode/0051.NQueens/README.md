# 51. NQueens

<https://leetcode.com/problems/n-queens/>

queen 的攻擊規則是 same row, column and diagonal  
題目的要求是在 `N x N` 的 matrix 中放入 `N` 個 queen

所以每一列都要可以放置一個 queen  
一列一列放置，當發現該列沒有任何空間能放置時即中斷

## Takeaway

- Backtracking
