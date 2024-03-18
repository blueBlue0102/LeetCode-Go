# 209. Minimum Size Subarray Sum

<https://leetcode.com/problems/minimum-size-subarray-sum/>

在知道要使用 sliding window 之後，問題就會在於迴圈要如何設計  
自己想不通時，迴圈的條件判斷可能就會非常複雜  
但是一旦知道最佳做法時，就會變得很簡單

程式碼的寫法是雙層迴圈  
乍看之下是 `O(n^2)`，但其實是 `O(n)`  
因為陣列中的每一個點實際上都只會被巡兩次  
所以 `O(2n) = O(n)`

## Takeaway

- 雙指針，sliding windows
