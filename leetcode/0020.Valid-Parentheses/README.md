# 20. Valid Parentheses

<https://leetcode.com/problems/valid-parentheses/>

關於合法的組合，就是當從左至右遍歷字串時，每當遇到任何一種右括號，之前最後出現過的左括號必須和右括號同種類  
題目已保證輸入字串只會由 `'('`, `')'`, `'{'`, `'}'`, `'['` and `']'` 組成  
所以只要遇到左括號時就 push 進 stack，遇到右括號時就 pop  
若 stack 為空無法 pop，或 pop 出的左括號和遇到的右括號不同種類時，即表示非法  
當所有字串遍歷完後 stack 必須為空，否則也是非法組合

## Takeaway

- stack
