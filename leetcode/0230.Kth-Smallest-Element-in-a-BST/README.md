# 230. Kth Smallest Element in a BST

<https://leetcode.com/problems/kth-smallest-element-in-a-bst/>

給定一個 BST(Binary Search Tree)，尋找第 $K_{th}$ 小的數字  
BST 是一個已經排序好的樹，其定義"通常"是【左子樹 < root < 右子樹】  
題目裡並沒有說明清楚是否有重複值的可能性  
也就是【左子樹 <= root < 右子樹】或【左子樹 < root <= 右子樹】的可能性

但這並不影響這題的做法

直覺的解法是進行 Inorder Traversal  
第一個 visit 的 node 即為 $K_{1}$，第二個即為 $K_{2}$，以此類推

## Takeaway

- Binary Search Tree
- DFS - Inorder Traversal
