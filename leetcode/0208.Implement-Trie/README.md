# 208. Implement Trie

<https://leetcode.com/problems/implement-trie-prefix-tree/>

Trie 是一種資料結構，常見的應用場景在於 autocomplete

可先參考 Neetcode 的介紹  
<https://youtu.be/oobqoCJlHA0?si=ei5hvEJX29wZnMqk>

和 Hash Map 的一大差別在於 `startsWith` function  
例如字典內含有 `"apple"` 這個單詞時，則預期輸入 `startsWith("app")` 要能得到 `true`  
若使用 Hash Map 需要以 $O(n)$ 的時間複雜度才能查詢  
但若使用 Trie 的話則預期不需要 $O(n)$ 時間複雜度

## Takeaway

- Trie
