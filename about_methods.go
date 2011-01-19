package about_methods

import "./koans"

type Int int

func (t Int) Add(v Int) Int {
	return Int(int(t) + int(v))
}

func AGlobalFunction(a, b int) int {
	return a + b
}

func FunctionWithMultipleReturnTypes() (int, string) {
	return 3, "hello"
}

func TestFunctionWithMultipleReturnTypes(t *koans.T) {
	intValue, stringValue := FunctionWithMultipleReturnTypes()

	t.AssertEqualInt(koans.Int__, intValue)
	t.AssertTrue(koans.String__ == stringValue)
}

func TestCallingAGlobalFunction(t *koans.T) {
	result := AGlobalFunction(3, 4)
	t.AssertEqualInt(koans.Int__, int(result))
}

func TestEveryTypeCanHaveMethods(t *koans.T) {
	v := Int(4)
	result := v.Add(3)
	t.AssertEqualInt(koans.Int__, int(result))
}

