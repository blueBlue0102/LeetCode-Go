# 207. Course Schedule

<https://leetcode.com/problems/course-schedule/>

`numCourses`: 總共有幾種課要修  
`prerequisites`: 課程之間的依賴關係

判斷課程是否有辦法完成

例如 `numCourses = 2, prerequisites = [[1,0]]`  
總共有兩種課程  
若要修`課程_1`需要先完成`課程_0`  
這是可能達成的，所以答案是 `true`

例如 `numCourses = 2, prerequisites = [[1,0], [0,1]]`  
總共有兩種課程  
若要修`課程_1`需要先完成`課程_0`；若要修`課程_0`需要先完成`課程_1`  
這是不可能達成的，`課程_0`和`課程_1`是一個閉環，所以答案是 `false`

以有向圖的概念思考  
若視 `numCourses` 為 Vertex 數量，`prerequisites` 則為 Edge  
然後要確認這個有向圖當中沒有閉環

進行圖的 Traversal，記錄遇過的節點，若重複遇到節點則表示有閉環  
已確認沒有閉環的節點需要記錄下來，避免重複搜索

## Takeaway

- Graph Traversal
