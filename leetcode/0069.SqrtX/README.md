# 69. SqrtX

<https://leetcode.com/problems/sqrtx/>

由於只需要取 rounded down to the nearest integer  
所以不需要真的進行 sqrt  
而是考慮如何快速找到某個數字其平方最接近並且小於等於 `x`

假設 $sqrt(x)=floor(k)$，則我們的目標就是要找到 $k^2 <= x$ 並且 $(k+1)^2 > x$  
然後就可以在 `0~x` 之間進行 binary search 來找出這個數字

## Takeaway

- Binary Search
