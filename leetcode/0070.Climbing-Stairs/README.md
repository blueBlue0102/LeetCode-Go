# 70. Climbing Stairs

<https://leetcode.com/problems/climbing-stairs/>

由於正在練習遞迴，因此使用遞迴來解題  
[可以參考這個影片學習遞迴的解題思路](https://youtu.be/ngCos392W4w?si=fYEm6G-f5LGP5D8y)  
不過如果用純遞迴寫法的話，在 Leetcode 上會 TLE，所以需要搭配 hash table 紀錄重覆的計算結果

爬第 n 層的方法等於【n-1 層的爬法後再走 "1"】+【n-2 層的爬法後再走 "2"】

所以：

$$
f(n) = \begin{cases}
f(n-1) + f(n-2) & \text{若} \ n > 2, \ n \in \mathbb{Z}^+ \\
2 & \text{若} \ n = 2 \\
1 & \text{若} \ n = 1
\end{cases}
$$

這也是所謂的[費波那契數](https://zh.wikipedia.org/wiki/%E6%96%90%E6%B3%A2%E9%82%A3%E5%A5%91%E6%95%B0)

## Takeaway

無
