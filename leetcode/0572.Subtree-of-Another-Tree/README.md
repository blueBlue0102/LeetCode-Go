# 572. Subtree of Another Tree

<https://leetcode.com/problems/subtree-of-another-tree/>

給兩棵樹 root 和 subRoot  
驗證 subRoot 是否為 root 中的 sub tree

寫一個 `IsSameTree` function，若 root != subRoot 時  
則接著判定 root.left 或 root.right 是否等於 subRoot

## Takeaway

- Tree Traversal
