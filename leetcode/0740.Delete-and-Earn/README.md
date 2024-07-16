# 740. Delete and Earn

<https://leetcode.com/problems/delete-and-earn/>

欲在有因果關聯的選擇中找出 maximize，所以使用 DP 的概念解題

由於 `nums` 中的數字是無序的，且可能有數字會重複出現  
所以建立一個 `points` hash map，使 `points[n]` 即為 `n * 出現次數` 的值

影響決定是否選擇數字 `n` 的是 `n-1` 和 `n+1` 的分數  
若設計 `dp(n)` 為「只考慮 `nums` 中所有小於等於 `n` 的值的最大總和」  
且若從 `nums` 中的最大值開始進行 `dp()` 的話  
`dp()` 內的邏輯就變成只需要考慮 `n`, `n-1` 和 `n-2` 之間的比較  
若選了 `n` 就勢必無法選擇 `n-1`，但可以選擇 `n-2`  
所以變成 `points[n]+dp(n-2)` 和 `dp(n-1)` 之間大小的比較  
（若 `n` 不存在於 `nums` 之中的話，`points[n]` 便是 `0`）

```
// Base Cases
dp(0) = 0
dp(1) = points[1]
// General Cases
dp(2) = max(points[2]+dp(0), dp(1))
dp(3) = max(points[3]+dp(1), dp(2))
dp(4) = max(points[4]+dp(2), dp(3))
dp(n) = max(points[n]+dp(n-2), dp(n-1))
```

### 若是類似 `[1, 2, 3, 10000]` 這樣的 `nums`？

上面的解法的缺點  
當 `nums = [1, 2, 3, 10000]`，就會將 `dp(10000) dp(9999) dp(9998) ... dp(0)` 逐一都計算一遍  
所以一個做法是可以先將 `nums` 進行 sort，計算 dp 時就可以按照數字順序越過不存在的數字進行計算  
只是這樣的缺點是 sort 會額外增加 $O(logn)$ 的時間複雜度  
在數字很密集的 case 下會比原來的方法來的慢，得不到 sort 來越過計算不存在數字的好處

## Takeaway

- Dynamic Programming
