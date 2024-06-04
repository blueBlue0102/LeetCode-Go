# 105. Construct Binary Tree from Preorder and Inorder Traversal

<https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/>

給定 `preorder` 和 `inorder` 的數字陣列，轉換成樹回傳  
已知這兩個陣列都由獨一無二的數字組成  
且 `preorder` 的陣列可能由 `[root, lefts..., rights...]` 所組成  
`inorder` 的陣列可能由 `[lefts..., root, rights...]` 所組成  

1. `preorder` 的第一個 element 即為 root
1. 在 `inorder` 中找尋此 element。inorder 陣列中，該 element 的左邊即是左子樹的內容；反之右子樹
1. 以 `preorder[left:right]` 和 `inorder[:root]` 建立左子樹；以 `preorder`[right:]` 和 `inorder[root+1:]` 建立右子樹

由於需要不斷查詢特定元素於 array 的位置，所以先建立一個 `value:index` 的 hash 對應表來避免 $O(n^2)$ 的時間複雜度

## Takeaway

- Tree Traversal
