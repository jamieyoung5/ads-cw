package dlx

import "testing"

func TestSearch(t *testing.T) {
	// Setup a simple exact cover problem that is known to have a solution
	matrix := NewMatrix([]string{"1", "2", "3", "4", "5", "6", "7"})
	matrix.AppendRow([]string{"1", "4", "7"})
	matrix.AppendRow([]string{"1", "4"})
	matrix.AppendRow([]string{"4", "5", "7"})
	matrix.AppendRow([]string{"3", "5", "6"})
	matrix.AppendRow([]string{"2", "3", "6", "7"})
	matrix.AppendRow([]string{"2", "7"})

	var solutions [][]*Node
	matrix.Search([]*Node{}, &solutions)

	if len(solutions) == 0 {
		t.Error("Search failed to find a solution for a solvable problem")
	}
}
