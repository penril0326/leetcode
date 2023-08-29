package keysandrooms

// Time: O(N+E), E is total keys, N is total rooms
// Space: O(N)
func canVisitAllRooms(rooms [][]int) bool {
	visited := make([]bool, len(rooms))

	components := 0
	for i := 0; i < len(rooms); i++ {
		if !visited[i] {
			components++
			dfs(i, rooms, &visited)
		}
	}

	return (components == 1)
}

func dfs(room int, rooms [][]int, visited *[]bool) {
	(*visited)[room] = true

	for _, key := range rooms[room] {
		if !(*visited)[key] {
			dfs(key, rooms, visited)
		}
	}
}
