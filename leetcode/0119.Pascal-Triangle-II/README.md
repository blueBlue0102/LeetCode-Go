# 119. Pascal Triangle II

<https://leetcode.com/problems/pascals-triangle-ii/>

關於 Pascal Triangle  
首先要知道 $T(row, col) = T(row - 1, col - 1) + T(row - 1, col)$  
且 $T(row, col) = 1$ *where* $row = col$ *or* $col = 0$

而且若要得出第 n 列的所有值時，只需要 n-1 列的值  
因此可以透過遞迴的方式來取得答案

目前程式碼的時間複雜度是 $O(n^2)$  
不同層遞迴的執行次數不一樣，$1+2+...+n = \frac{n(n+1)}{2} \approx n^2$  
空間複雜度則因為每層遞迴都是兩個 array 的 size，因此是 $O(2n) = O(n)$

Leetcode 解答中裡有時間和空間複雜度都是 $O(n)$ 的解法，且不需要遞迴  
但需要了解 Pascal's rule，待日後有機會再來了解

## Takeaway

- Pascal Triangle 特性
