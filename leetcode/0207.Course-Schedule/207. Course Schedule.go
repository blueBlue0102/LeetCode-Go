package leetcode

func CourseSchedule(numCourses int, prerequisites [][]int) bool {
	courseMap := map[int][]int{}
	visitMap := map[int]bool{}
	okCourseMap := map[int]bool{}
	for _, prerequisite := range prerequisites {
		if _, existed := courseMap[prerequisite[0]]; !existed {
			courseMap[prerequisite[0]] = make([]int, 0)
		}
		courseMap[prerequisite[0]] = append(courseMap[prerequisite[0]], prerequisite[1])
	}

	var dfs func(course int, visitMap map[int]bool) bool
	dfs = func(course int, visitMap map[int]bool) bool {
		if visitMap[course] {
			return false
		}
		if okCourseMap[course] {
			return true
		}

		visitMap[course] = true
		for _, prerequisite := range courseMap[course] {
			if !dfs(prerequisite, visitMap) {
				return false
			}
		}
		visitMap[course] = false

		okCourseMap[course] = true
		return true
	}

	for course := range courseMap {
		if !dfs(course, visitMap) {
			return false
		}
	}
	return true
}
