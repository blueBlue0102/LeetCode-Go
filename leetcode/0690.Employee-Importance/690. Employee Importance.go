package leetcode

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

func EmployeeImportance(employees []*Employee, id int) int {
	// return BFS(employees, id)
	return DFS(employees, id)
}

func BFS(employees []*Employee, id int) int {
	sum := 0
	deque := []int{id}

	// hash map
	employeeMap := map[int](*Employee){}
	for _, employee := range employees {
		employeeMap[employee.Id] = employee
	}

	for len(deque) > 0 {
		// pop front
		id := deque[0]
		deque = deque[1:]

		// push back
		deque = append(deque, employeeMap[id].Subordinates...)

		sum += employeeMap[id].Importance
	}

	return sum
}

func DFS(employees []*Employee, id int) int {
	// hash map
	employeeMap := map[int]*Employee{}
	for _, employee := range employees {
		employeeMap[employee.Id] = employee
	}

	return dfsRecur(employeeMap, id)
}

func dfsRecur(employeeMap map[int]*Employee, id int) int {
	sum := employeeMap[id].Importance
	for _, subEmployeeId := range employeeMap[id].Subordinates {
		sum += dfsRecur(employeeMap, subEmployeeId)
	}
	return sum
}
