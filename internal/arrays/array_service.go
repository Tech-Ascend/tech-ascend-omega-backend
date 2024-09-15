package arrays

import "fmt"

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
type ValidAnagramResults struct {
	ValidAnagram   bool
	FrequencyLogs  []string
	Report         string
}

func (arrayService *ArrayService) ValidAnagram(s string, t string) ValidAnagramResults {
	feedback := ValidAnagramResults{
		FrequencyLogs: []string{},
	}
fmt.Println(s, t)
	if len(s) != len(t) {
		feedback.ValidAnagram = false
		feedback.Report = "Words do not match in length, not a valid anagram"
		return feedback
	}

	var freq [26]int

	for idx := 0; idx < len(s); idx++ {
		freq[s[idx]-'a']++
		freq[t[idx]-'a']--
		feedback.FrequencyLogs = append(feedback.FrequencyLogs, "added one to "+string(s[idx]))
		feedback.FrequencyLogs = append(feedback.FrequencyLogs, "subtracted one from "+string(t[idx]))
	}

	for idx := 0; idx < len(freq); idx++ {
		if freq[idx] != 0 {
			feedback.ValidAnagram = false
			feedback.Report = "Words are not valid anagrams"
			return feedback
		}
	}

	feedback.ValidAnagram = true
	feedback.Report = "Words are valid anagrams"
	return feedback
}
