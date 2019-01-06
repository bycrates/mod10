package mod10

import (
	"errors"
	"fmt"
	"strconv"
)

const (
	// StringNotDigits is the error returned if
	// the specified string to String functions
	// cannot be converted into an integer
	StringNotDigits = "specified_string_is_not_int"
)

// reverse a string
func reverse(s string) string {
	rs := []rune(s)
	for i, j := 0, len(rs)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	return string(rs)
}

// doubleCondRet does the computation required for each digit
// We're cycling between doubling and not doubling each time
// That is handled in the main function
func doubleCondRet(num int, cycle bool) int {
	if cycle {
		// Double the input number
		double := num * 2

		// The subtraction should be 0 unless
		subNum := 0

		// The double number is more than 9
		// you usually have to sum each separate digit
		// but subtracting 9 from the number
		// accomplishes the same task
		if double > 9 {
			subNum = 9
		}

		return double - subNum
	}

	return num
}

// totalSum calls doubleCondRet for each digit
// and does cycling
func totalSum(digits string, cycle bool) int {
	endSum := 0
	currentCycle := cycle

	for _, digit := range digits {
		digitInt, _ := strconv.Atoi(string(digit))
		endSum += doubleCondRet(digitInt, currentCycle)

		// Flip the cycle
		currentCycle = !currentCycle
	}

	return endSum
}

func common(digits int, cycle bool) int {
	return totalSum(reverse(strconv.Itoa(digits)), cycle)
}

// AddControlBit computes mod10 and appends controlbit
// to end of input
func AddControlBit(digits int) string {
	// The control bit is 10 subtracted with the modulo 10
	// of the Mod10 algorithmic operation
	controlBit := 10 - (common(digits, true) % 10)

	// If the control bit is 10 then set it to 0
	if controlBit == 10 {
		controlBit = 0
	}

	return fmt.Sprintf("%d%d", digits, controlBit)
}

// AddControlBitString does the same action as AddControlBit
// but accepts a string for digits
func AddControlBitString(digits string) (string, error) {
	intDigits, err := strconv.Atoi(digits)
	if err != nil {
		return "", errors.New(stringNotDigits)
	}

	return AddControlBit(intDigits), nil
}

// Check checks a mod10 string with a control bit if it's
// control bit is valid
func Check(digits int) bool {
	return common(digits, false)%10 == 0
}

// CheckString does the same action as Check
// but accepts a string for digits
func CheckString(digits string) (bool, error) {
	intDigits, err := strconv.Atoi(digits)
	if err != nil {
		return false, errors.New(stringNotDigits)
	}

	return Check(intDigits), nil
}
