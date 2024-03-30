# 744. Find Smallest Letter Greater Than Target

<https://leetcode.com/problems/find-smallest-letter-greater-than-target/>

最直覺的暴力解是，輪巡 letters 陣列，然後和 target 比較大小即可  
但這樣的做法，時間複雜度是 $O(n)$

時間複雜度 $O(log(n))$ 的寫法是透過 Binary Search 的做法  
由於給定的 array 已經經過排序  
所以可以使用 Binary Search 來解題

## Takeaway

- Binary Search 的練習
