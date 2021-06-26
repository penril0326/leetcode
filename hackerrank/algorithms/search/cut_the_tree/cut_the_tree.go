package cut_the_tree

func cutTheTree(data []int32, edges [][]int32) int32 {
	var totalValue int32
	value := make([]int32, len(data)+1)      // 每個節點的值，+1是因為input頂點的編號從1開始
	nodeList := make([][]int32, len(data)+1) // 第一層是頂點，第二層的slice存跟這個頂點有連接的點
	visited := make([]bool, len(data)+1)     // 記錄是否走訪過
	sum := make([]int32, len(data)+1)        // index表示那個節點為根的時候的總和

	// 計算所有頂點的總和
	for idx, v := range data {
		totalValue += v
		value[idx+1] = v
	}

	// 為每個頂點加入邊，無向圖，所以兩個方向都要有
	for _, edge := range edges {
		node1, node2 := edge[0], edge[1]
		nodeList[node1] = append(nodeList[node1], node2)
		nodeList[node2] = append(nodeList[node2], node1)
	}

	// 把1當作root做深度搜尋，更新sum
	dfs(1, sum, visited, value, nodeList)

	result := totalValue
	for i := 1; i <= len(data); i++ {
		t1 := totalValue - sum[i]
		t2 := totalValue - t1

		result = min(result, abs(t1-t2))
	}

	return result
}

func dfs(nodeIdx int, sum []int32, visited []bool, value []int32, nodeList [][]int32) int32 {
	var result int32
	visited[nodeIdx] = true
	depth := len(nodeList[nodeIdx])
	for i := 0; i < depth; i++ {
		next := nodeList[nodeIdx][i]
		if visited[next] == false {
			result += dfs(int(next), sum, visited, value, nodeList)
		}
	}

	sum[nodeIdx] = value[nodeIdx] + result

	return sum[nodeIdx]
}

func min(a, b int32) int32 {
	if a < b {
		return a
	}

	return b
}

func abs(a int32) int32 {
	if a < 0 {
		return -a
	}

	return a
}
