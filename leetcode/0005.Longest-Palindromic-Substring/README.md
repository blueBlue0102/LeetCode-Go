# 5. Longest Palindromic Substring

<https://leetcode.com/problems/longest-palindromic-substring/>

要找出字串 `s` 中的最長回文  
回文的定義是從頭和從尾讀的順序都一樣的字串  
（單一個字元也算是回文）

`s` 中可能有不只一組的最長字串，只需要找到其中一組即可

回文根據有或沒有中心點而有兩種形式  
例如：`abc|cba`, `bbbcc|ccbbb` 屬於左右對稱型  
例如：`abc|e|cba`, `bbbcc|d|ccbbb` 屬於有中心點的類型  

要找特定字元的最長迴文，時間複雜度是 $O(n)$  
找完所有字元的迴文，時間複雜度是 $O(n^2)$

## Takeaway


