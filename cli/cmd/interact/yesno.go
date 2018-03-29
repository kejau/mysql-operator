package interact

import "github.com/grtl/mysql-operator/cli/cmd/util"

// YesNoInput waits for user to input yes or no.
func YesNoInput(options *util.Options) (bool, error) {
	if options.AssumeYes {
		return true, nil
	}

	input, err := MatchingStringInput("^(y(es)?|no?)?$", true, "Is that correct [y/N]: ")
	if err != nil {
		return false, err
	}

	return isYes(input), err
}

func isYes(input string) bool {
	return input == "y" || input == "yes"
}
