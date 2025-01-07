package roman_to_integer

func romanToInt(s string) int {
	symbols := make(map[string]int)
	symbols["I"] = 1
	symbols["V"] = 5
	symbols["X"] = 10
	symbols["L"] = 50
	symbols["C"] = 100
	symbols["D"] = 500
	symbols["M"] = 1000
	symbols["IV"] = 4
	symbols["IX"] = 9
	symbols["XL"] = 40
	symbols["XC"] = 90
	symbols["CD"] = 400
	symbols["CM"] = 900

	var total int
	i := 0
	for i < len(s) {
		if i < len(s)-1 {
			pair := string(s[i]) + string(s[i+1])
			if value, exists := symbols[pair]; exists {
				total += value
				i += 2
				continue
			}
		}
		total += symbols[string(s[i])]
		i++
	}
	return total
}
