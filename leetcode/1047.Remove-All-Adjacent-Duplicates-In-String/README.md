# 1047. Remove All Adjacent Duplicates In String

<https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string/>

~~只要字串中一旦有重複的連續字元，就所有連在一起的字元都刪除  
且刪除後，剩餘的字串也不能連在一起，否則也都要被刪除  
直到不再有字元被刪除為止~~

> A duplicate removal consists of choosing "two" adjacent and equal letters and removing them

所以只刪除連接在一起的"二個"字元  
`"aaa"` => `"a"`  
`"aaaa"` => `""`

## Takeaway

- Stack
