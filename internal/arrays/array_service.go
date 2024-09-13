package arrays

type ArrayService struct{}

type Step struct {
	Number      int  `json:"number"`
	IsDuplicate bool `json:"isDuplicate"`
}

type ContainsDuplicatesResult struct {
	ContainsDuplicates bool   `json:"containsDuplicates"`
	Steps              []Step `json:"steps"`
}

func (arrayService *ArrayService) ContainsDuplicates(numbers []int) ContainsDuplicatesResult {
	result := ContainsDuplicatesResult{
		ContainsDuplicates: false,
		Steps:              make([]Step, 0, len(numbers)),
	}

	if len(numbers) <= 1 {
		for _, number := range numbers {
			result.Steps = append(result.Steps, Step{Number: number, IsDuplicate: false})
		}
		return result
	}

	seen := make(map[int]struct{})

	for _, number := range numbers {
		if _, found := seen[number]; found {
			result.Steps = append(result.Steps, Step{Number: number, IsDuplicate: true})
			result.ContainsDuplicates = true
			return result
		}
		seen[number] = struct{}{}
		result.Steps = append(result.Steps, Step{Number: number, IsDuplicate: false})
	}

	return result
}

type SudokuFeedback struct {
	Digit  int
	Row    int
	Column int
	Valid  bool
	Step   string
}

func (arrayService *ArrayService) IsValidSudoku(board [][]byte) []SudokuFeedback {
	rows := [9][9]bool{}
	cols := [9][9]bool{}
	grid := [3][3][9]bool{}
	sudokuFeedback := []SudokuFeedback{}
	var reason string

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			cell := board[r][c]

			if cell == '.' {
				sudokuFeedback = append(sudokuFeedback, SudokuFeedback{
					Row:    r,
					Column: c,
					Digit:  -1,
					Valid:  true,
					Step:   "Empty cell",
				})
				continue
			}
			digit := int(cell) - 49 

			isValid := true

			if rows[r][digit] {
				isValid = false
				reason = "Row conflict"
			}

			if cols[c][digit] {
				isValid = false
				reason = "Column conflict"
			}

			if grid[r/3][c/3][digit] {
				isValid = false
				reason = "Grid conflict"
			}

			if isValid {
				rows[r][digit] = true
				cols[c][digit] = true
				grid[r/3][c/3][digit] = true
				reason = "Valid placement"
			}			

			sudokuFeedback = append(sudokuFeedback, SudokuFeedback{
				Row:    r,
				Column: c,
				Digit:  digit,
				Valid:  isValid,
				Step:   reason,
			})

			if !isValid {
				return sudokuFeedback
			}
		}
	}

	return sudokuFeedback
}
