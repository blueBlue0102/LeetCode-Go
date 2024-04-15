# 394. Decode String

<https://leetcode.com/problems/decode-string/>

可以用 stack 或是遞迴的方法解題  
目前使用的是遞迴

使用遞迴的解題思考方式，重點在於想清楚每層遞迴的處理範疇

e.g.: 

- `f("a") = "a"`
- `f("2[a]") = f("aa") = "aa"`
- `f("2[a3[b]]") = f("2[abbb]") = f("abbbabbb") = "abbbabbb"`
- `f("a2[b]3[c]") = f("abb3[c]") = f("abbccc") = "abbccc"`

## Takeaway

無
