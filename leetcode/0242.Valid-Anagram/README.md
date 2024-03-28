# 242. Valid Anagram

<https://leetcode.com/problems/valid-anagram/>

檢查字串 `s` 和 `t` 的所有字元和個數是否相等，無視順序

有兩種做法：

1. Sort

   將字串做 sort 後進行比較  
   由於 sort 的時間複雜度是 $O(n*log(n))$，所以時間複雜度是 $O(n*log(n))$

   空間複雜度是 $O(1)$

2. Hash Table

   直覺的做法可能是字串 `s` 和 `t` 各自建立 table 然後進行比較  
   但其實只要建立一個 table 就可以了，譬如在字串 `s` 時，每遇到一個字就將數字加一，字串 `t` 時則減一  
   最終結果若 table 中每個 key 值皆為 `0` 則表示符合

   時間複雜度是 $O(n)$  
   空間複雜度是 $O(1)$，因為 table 的 size 並不受字串長短而有改變

在 Go 語言，string 是由一連串的 byte 組成  
此題已經限定字元只會是英文的 26 個小寫字母，所以都剛好是 1 byte  
但如果沒有這個限制的話，字串中可能包含任意字元，就很可能會出現單個字超過 1 byte 的情形  
因此，遞迴字串時需要透過 `range` 的方式，來確保遞迴中的每個 value 都是完整的單個字  
<https://go.dev/blog/strings#:~:text=strings%20are%20normalized.-,Range%20loops,-Besides%20the%20axiomatic>

## Takeaway

無
