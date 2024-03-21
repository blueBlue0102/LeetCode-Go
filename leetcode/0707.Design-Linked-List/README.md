# 707. Design Linked List

<https://leetcode.com/problems/design-linked-list/>

可以自行選擇要用 singly linked list 或是 doubly linked list 來實作  
singly linked list 的設計較簡單，`addAtHead` 是 constant time，`addAtTail` 則是 linear time  
doubly linked list 則是在 `addAtHead` 和 `addAtTail` 都是 constant time

我實作的是 singly linked list  
實作完後看解題討論，才知道如果有 sentinel node 的話，程式邏輯會更簡單些

## Takeaway

- 解 linked list 相關題目，搭配 sentinel node 常常能使情境更簡單
