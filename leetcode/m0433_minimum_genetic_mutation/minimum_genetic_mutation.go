package minimumgeneticmutation

import "strings"

// Time: ?
// Space: ?
func minMutation(startGene string, endGene string, bank []string) int {
	bankMap := map[string]bool{}
	for _, validGene := range bank {
		bankMap[validGene] = true
	}

	if !bankMap[endGene] {
		return -1
	}

	type mutation struct {
		gene string
		mute int
	}

	queue := make([]mutation, 0)
	queue = append(queue, mutation{startGene, 0})
	visited := map[string]bool{}
	visited[startGene] = true

	for len(queue) > 0 {
		gene, mute := queue[0].gene, queue[0].mute
		queue = queue[1:]
		if gene == endGene {
			return mute
		}

		for _, neighbor := range neighbors(gene) {
			if !visited[neighbor] && bankMap[neighbor] {
				visited[neighbor] = true
				queue = append(queue, mutation{neighbor, mute + 1})
			}
		}
	}

	return -1
}

func neighbors(gene string) []string {
	change := "ACGT"
	neighborList := []string{}
	for i, r1 := range gene {
		for _, r2 := range change {
			if r1 == r2 {
				continue
			}

			var sb strings.Builder
			sb.WriteString(gene[:i])
			sb.WriteRune(r2)
			sb.WriteString(gene[i+1:])

			neighborList = append(neighborList, sb.String())
		}
	}

	return neighborList
}
