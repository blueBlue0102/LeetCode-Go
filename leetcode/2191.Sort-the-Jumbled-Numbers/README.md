# 2191. Sort the Jumbled Numbers

<https://leetcode.com/problems/sort-the-jumbled-numbers>

此題目標是將 `nums` 透過 `mapping` 轉換後進行 non-decreasing order 的排序  
`mapping` 是一個長度為 `10` 的 int 陣列，`mapping[i] = j`  
`nums` 中為 `i` 的 digit 都要轉換成 `j`

若多個 `nums` 中的數字經轉換後的值相等，則排序的順序需要保持和原先 `nums` 中的順序相同

## Takeaway

- Comparison Sort
