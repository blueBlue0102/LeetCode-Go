# 104. Maximum Depth of Binary Tree

<https://leetcode.com/problems/maximum-depth-of-binary-tree/>

給定一棵 Binary Tree，求樹的最大高度  
（若只有 root 則高度為 1）

樹需要從 root 開始進行遍歷，但位於 root 時並無法得知右或左子樹能否達到最遠節點  
所以勢必需要遍歷整棵樹

要用 pre 或 in-order 都無所謂，因為勢必就是需要遍歷每個節點

## Takeaway

- Tree Traversal
