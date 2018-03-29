package interact

import (
	"fmt"
	"regexp"
	"strings"
)

// MatchingStringInput waits for user to input a string that matches expr.
// If lowercase is true then, before matching, string is transformed to lowercase.
func MatchingStringInput(expr string, lowercase bool, question string) (string, error) {
	var input string
	r, err := regexp.Compile(expr)
	if err != nil {
		return "", err
	}

	for {
		fmt.Print(question)
		_, err := fmt.Scanf("%s", &input)
		if err != nil {
			return "", err
		}

		if lowercase {
			input = strings.ToLower(input)
		}

		if r.MatchString(input) {
			return input, nil
		}

		fmt.Println("Invalid input format")
	}
}
