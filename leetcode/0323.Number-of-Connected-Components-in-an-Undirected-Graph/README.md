# 323. Number of Connected Components in an Undirected Graph

<https://leetcode.com/problems/number-of-connected-components-in-an-undirected-graph/>

題目是 Premium 的，可以參考 <https://github.com/TheExplainthis/LeetCodeJourney/tree/main/solutions/323.%20Number%20of%20Connected%20Components%20in%20an%20Undirected%20Graph>

無向圖  
給定 node 數量 `n` 和 edge `edges`  
找出有幾個 group

此題是一個 Union Find 的情境  
也可以使用 DFS 或 BFS 解題

這邊使用 DFS 解題

## Takeaway

- Graph Traversal
