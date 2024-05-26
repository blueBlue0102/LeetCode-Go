# 150. Evaluate Reverse Polish Notation

<https://leetcode.com/problems/evaluate-reverse-polish-notation/>

輪巡 string，使用 stack 將所遇到的數字儲存  
一旦遇到加減乘除符號就將 stack 的前兩個數字取出做運算，再將運算結果放回 stack

## Takeaway

- Stack
