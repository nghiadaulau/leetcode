# LeetCode Solutions

This repository contains my solutions to various LeetCode problems, organized by difficulty level. Each solution is implemented in Go and includes detailed explanations where applicable.

## Repository Structure
```
leetcode/
├── easy/
│ ├── two-sum/
│ ├── palindrome-number/
│ └── ...
├── medium/
│ ├── longest-palindromic-substring/
│ ├── longest-substring-without-repeating/
│ └── ...
└── hard/
└── ...
```

## Current Solutions

### Easy
- [Two Sum](easy/two-sum)
- [Palindrome Number](easy/palindrome-number)
- [Roman to Integer](easy/roman-to-integer)
- [Longest Common Prefix](easy/longest-common-prefix)
- [Valid Parentheses](easy/valid-parentheses)
- [Remove Duplicates from Sorted Array](easy/remove-duplicates-from-sorted-array)

### Medium
- [Longest Palindromic Substring](medium/longest-palindromic-substring)
- [Longest Substring Without Repeating Characters](medium/longest-substring-without-repeating)

### Hard
*(Coming soon)*

## How to Contribute

1. Fork the repository
2. Create a new branch for your solution: `git checkout -b problem-name`
3. Add your solution following this structure:
   ```
   difficulty-level/
   └── problem-name/
       ├── main.go       # Solution implementation
       └── README.md     # Problem description and explanation (optional)
   ```
4. Ensure your code:
   - Passes all LeetCode test cases
   - Includes comments explaining the approach
   - Follows Go best practices and conventions
5. Create a Pull Request with:
   - Clear description of the problem and solution
   - LeetCode problem link
   - Time and space complexity analysis

## Updating the README

When adding a new solution, please update the README.md file following these steps:

1. Add your solution to the appropriate section under "Current Solutions"
   ```markdown
   ### [Difficulty Level]
   - [Problem Name](difficulty-level/problem-name)
   ```

2. Ensure the link follows the correct folder structure
   - Example: `[Two Sum](easy/two-sum)`
   
3. Keep solutions sorted:
   - Within each difficulty section, maintain alphabetical order
   - If adding a new difficulty section, follow the order: Easy → Medium → Hard

4. If adding the first solution to a difficulty level:
   - Replace "*(Coming soon)*" with your solution
   - Ensure proper heading structure (###)

Example addition:
```markdown
### Easy
- [Existing Solution](easy/existing-solution)
- [New Solution](easy/new-solution)  <!-- Add your solution here -->
```

## Testing

Each solution can be tested directly on LeetCode's platform. Local tests can be added in the future.


## Contact

Feel free to reach out if you have any questions or suggestions!

---
⭐ Star this repository if you find it helpful!
