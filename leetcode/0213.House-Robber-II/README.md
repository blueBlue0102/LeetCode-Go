# 213. House Robber II

<https://leetcode.com/problems/house-robber-ii/>

給定的 `nums` 是一個環，`nums[0]` 和 `nums[n-1], n = len(nums)` 是相連的  
要找出能搶的最大金額，且當下的選擇會影響未來的選擇，所以是一個 DP problem

解題思維和 [198. House Robber](../0198.House-Robber/README.md) 一樣  
只是差別在於要解決 base case 如何設計的問題  
當選擇了 `nums[0]` 就意味不能選擇 `nums[n-1]`，除了頭尾的部分需特別考量，其餘部分都是一樣的

所以分成兩種 case，選擇或不選擇 `nums[0]`  
選擇 `nums[0]` 即表示 `nums[1]` 和 `nums[n-1]` 都必定不會被選擇  
所以可以當作在計算 `num[2:n-1]`  
不選擇 `nums[0]` 即表示計算 `nums[1:]`  
最後比較 `nums[0]+dp(nums[2:n-1])` 和 `dp(nums[1:])` 即可

## Takeaway

- Dynamic Programming
