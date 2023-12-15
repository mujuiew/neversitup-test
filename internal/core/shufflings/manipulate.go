package shufflings

// BodyData : body for shufflings
type BodyData struct {
	Data string
}

// Process :  generates all unique permutations of the given string.
func Process(data BodyData) []string {

	result := make([]string, 0)
	visited := make(map[string]bool)
	generatePermutations([]rune(data.Data), 0, &result, visited)

	return result
}

func generatePermutations(chars []rune, index int, result *[]string, visited map[string]bool) {
	if index == len(chars)-1 {
		currentPermutation := string(chars)
		if !visited[currentPermutation] {
			*result = append(*result, currentPermutation)
			visited[currentPermutation] = true
		}
		return
	}

	for i := index; i < len(chars); i++ {
		chars[index], chars[i] = chars[i], chars[index]
		generatePermutations(chars, index+1, result, visited)
		chars[index], chars[i] = chars[i], chars[index]
	}
}
