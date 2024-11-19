# 49. Group Anagrams

<https://leetcode.com/problems/group-anagrams/>

題目給定一個 string array  
找出有哪些 string 是由相同字母組成

作法一：Sort + Hash Table  
建立一個 Hash Table  
key 是 sort 後的 string；value 是 sort 前的 string  
掃完整個 array 後建立完 table，即可得到每一組相同字母的 string array

每次的 sort 時間複雜度是 $O(k \cdot log(k))$，k 為最長 string 的長度  
假設有 n 個 string，所以時間複雜度是 $O(n \cdot k \cdot log(k))$  
由於 Hash Table 需要存下所有 string，所以空間複雜度是 $O(n \cdot k)$

作法二：Hash Table  
由於已知給定的 string 只會由小寫字母組成，總共 26 個  
和做法一一樣建立一個 Hash Table  
然後 key 可以不必 sort，而是用一個長度 26 的 int array 紀錄各字母的出現次數  
接著將這個 array 轉換成 string 後做為 Hash Table 的 key  

這樣子的做法  
時間複雜度是 $O(n \cdot k)$，n 是 string 個數，k 是最長 string 的長度  
空間複雜度是 $O(n \cdot k)$

## Takeaway

- Hash Table
