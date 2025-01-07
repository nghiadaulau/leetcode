package valid_parentheses

func isValid(s string) bool {
	runes := []rune(s)
	var fruits []string
	for i := 0; i < len(runes); i++ {
		if string(runes[i]) == "(" {
			fruits = append(fruits, "(")
		} else if string(runes[i]) == "{" {
			fruits = append(fruits, "{")
		} else if string(runes[i]) == "[" {
			fruits = append(fruits, "[")
		}

		if string(runes[i]) == ")" {
			if fruits == nil || len(fruits) < 1 || string(fruits[len(fruits)-1]) != "(" {
				return false
			}
			fruits = fruits[:len(fruits)-1]
		} else if string(runes[i]) == "}" {
			if fruits == nil || len(fruits) < 1 || string(fruits[len(fruits)-1]) != "{" {
				return false
			}
			fruits = fruits[:len(fruits)-1]

		} else if string(runes[i]) == "]" {
			if fruits == nil || len(fruits) < 1 || string(fruits[len(fruits)-1]) != "[" {
				return false
			}
			fruits = fruits[:len(fruits)-1]
		}
	}

	if len(fruits) >= 1 {
		return false
	}
	return true
}
