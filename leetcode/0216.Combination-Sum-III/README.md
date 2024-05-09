# 216. Combination Sum III

<https://leetcode.com/problems/combination-sum-iii/>

找出 `k` 個數字加起來等於 `n` 的所有組合  
而且數字只能是 `0~9` 且不能重覆使用

由於要窮舉，所以是 Backtracking 題目  
首先思考樹的模樣  
最終 array 的樣子要從 `[1~9, 2~9, 3~9, ...]` 裡面去找出組合  
而且一旦還有位置需要填滿，但是數字已經大於目標了，此時就不該繼續遍歷樹

## Takeaway

- Backtracking
