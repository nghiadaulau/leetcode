# isValid Function

## Author: **ai**

## Overview
This repository contains an implementation of the `isValid` function in Go, which checks if a given string containing brackets is valid. A string is considered valid if:
1. Open brackets are closed by the same type of brackets.
2. Open brackets are closed in the correct order.

The brackets considered are:
- `()`
- `{}`
- `[]`

## Function Signature
```go
func isValid(s string) bool
```

### Parameters
- `s` (string): The input string containing brackets to be validated.

### Returns
- `bool`: Returns `true` if the string is valid, otherwise `false`.

## How It Works
1. The function uses a stack (`fruits`) to keep track of open brackets.
2. It iterates through each character of the input string:
    - If an open bracket (`(`, `{`, `[`) is encountered, it is pushed onto the stack.
    - If a closing bracket (`)`, `}`, `]`) is encountered, it checks if the top of the stack has the corresponding opening bracket. If not, the string is invalid, and the function returns `false`.
    - If the match is correct, the opening bracket is popped from the stack.
3. After processing all characters, the stack should be empty for the string to be valid. If not, the function returns `false`.

## Example Usage
```go
package main

import (
	"fmt"
)

func main() {
	examples := []string{
		"()",
		"()[]{}",
		"(]",
		"([)]",
		"{[]}",
	}

	for _, example := range examples {
		fmt.Printf("%s: %v\n", example, isValid(example))
	}
}
```

### Output
```
(): true
()[]{}: true
(]: false
([)]: false
{[]}: true
```

## Notes
- The function runs in O(n) time complexity, where n is the length of the input string, as each character is processed once.
- Space complexity is O(n) in the worst case due to the stack usage.
