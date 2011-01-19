package about_control_statements

import "./koans"

func TestIfThenElseStatements(t *koans.T) {
	var result string
	if true {
		result = "true value"
	} else {
		result = "false value"
	}

	t.AssertTrue(koans.String__ == result)
}

func TestIfThenStatement(t *koans.T) {
	var result string = "default value"
	if true {
		result = "true value"
	}
	t.AssertTrue(koans.String__ == result)
}