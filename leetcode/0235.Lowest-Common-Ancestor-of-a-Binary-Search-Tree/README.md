# 235. Lowest Common Ancestor of a Binary Search Tree

<https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/>

給定一個 BST，要找到其中二個節點 `p` 和 `q` 的最近共同祖先  
既然是共同祖先，則從根節點前往共同祖先的路上，節點都必須要相同

利用 BST 的特性，從 `root` 開始尋找 `p` 和 `q`  
若 cursor 的 val 比 `p` and `q` 還大，就往左邊找；反之往右邊找  
當出現 cursor 的值介於 `p` 和 `q` 之間時(包含值相同)，即表示該節點為最近的共同祖先

## Takeaway

- Binary Search Tree
