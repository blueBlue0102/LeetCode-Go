# 22. Generate Parentheses

<https://leetcode.com/problems/generate-parentheses/>

看到列舉題就可以想到 Backtracking

思考合法的括號特性，必定是左括號出現後才能出現右括號  
也就是左括號的數量大於右括號的數量時，才可以放置右括號

所以進行遞迴時，就檢查當下 string 的左右括號數量  
依據這個條件進行遞迴即可

## Takeaway

- Backtracking
