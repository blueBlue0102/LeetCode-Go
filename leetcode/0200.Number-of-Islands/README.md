# 200. Number of Islands

<https://leetcode.com/problems/number-of-islands/>

`1` 表示島嶼；`0` 表示海洋  
島嶼的定義是垂直和水平的四個點都是水，斜角的連接並不算連接

理想上，只需要將整個 `grid` 掃過一遍即可算出島嶼的數量  
假設 `grid` 的大小是 `m * n` 的話，則時間複雜度就會是 $O(m \cdot n)$

這題有三個解決方向，分別是 DFS, BFS 以及 Disjoint Set  
這邊使用 Disjoint Set 的方式來解題

Disjoint Set 的概念可以參考：  
<https://leetcode.com/explore/featured/card/graph/618/disjoint-set/3881/>

## Takeaway

- Disjoint Set
