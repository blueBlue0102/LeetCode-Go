# 690. Employee Importance

<https://leetcode.com/problems/employee-importance/>

為了可以快速查詢特定 id 所對應的 importance，所以可以建立 hash map  
之後就是如何進行 importance 加總的問題

由於題目已限定一個員工只會有一個或沒有主管，不會有附屬於兩個以上主管的情形  
因此是一個樹的結構  
所以 traversal 的方式常見的就是 BFS 或 DFS

## Takeaway

- tree traversal
- recursion
