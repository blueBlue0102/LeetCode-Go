# 121. Best Time to Buy and Sell Stock

<https://leetcode.com/problems/best-time-to-buy-and-sell-stock/>

直覺的想法是 $O(n^2)$ 解法，每個元素都逐一與從後面的元素找最大差價  
但其實可以 one-pass

一定可以 one-pass 的理由是，只要循過一次整個陣列，則必定可以知道何時是最低點，及該最低點後的最高點為何  
抱著這個思維去解題，就能想出 one-pass 的解法

## Takeaway

無
