package helpers

import (
	"fmt"
	"time"
)

type PrintDividerOptions struct {
	Step               int
	Label              string
	ShowTime           bool
	DrawEmptyLines     bool
	DrawEmptyLineAbove bool
	DrawEmptyLineBelow bool
}

func NewPrintDividerOptionsWithLine(label string) *PrintDividerOptions {
	return &PrintDividerOptions{
		DrawEmptyLineAbove: true,
		DrawEmptyLineBelow: true,
		Label:              label,
	}
}

func NewPrintDividerOptions(label string, step int) *PrintDividerOptions {
	return &PrintDividerOptions{
		Step:  step,
		Label: label,
	}
}

func NewPrintDividerOptionsDefault() *PrintDividerOptions {
	return &PrintDividerOptions{}
}

func PrintDivider(options *PrintDividerOptions) {
	dividerValue := fmt.Sprintf(
		"%[1]v----------------%[2]v %[3]v",
		ternaryOperator(options.Step != 0, fmt.Sprintf(" %v ", options.Step), ""),
		ternaryOperator(options.Label != "", fmt.Sprintf(" %v ", options.Label), ""),
		ternaryOperator(options.ShowTime, time.Now().Format("2006-01-02 15:04:05"), ""),
	)

	if options.DrawEmptyLines || options.DrawEmptyLineAbove {
		fmt.Println()
	}

	fmt.Println(dividerValue)

	if options.DrawEmptyLines || options.DrawEmptyLineBelow {
		fmt.Println()
	}
}

// func SetDefaultValue[T any](rawValue *T, defaultValue T) T {
// 	value := reflect.ValueOf(*rawValue)

// 	if value.IsZero() {
// 		return defaultValue
// 	}
// 	return *rawValue
// }

// func IsZeroValue[T any](rawValue *T) bool {
// 	value := reflect.ValueOf(*rawValue)

// 	return value.IsZero()
// }

func ternaryOperator[T any](cond bool, vtrue, vfalse T) T {
	if cond {
		return vtrue
	}
	return vfalse
}
