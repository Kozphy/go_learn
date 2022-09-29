package fixedpoint

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const DefaultPow = 1e8

type Value int64

const PosInf = Value(math.MaxInt64)
const NegInf = Value(math.MinInt64)

func NewFromFloat(val float64) Value {
	if math.IsInf(val, 1) {
		fmt.Println(1)
		return PosInf
	} else if math.IsInf(val, -1) {
		fmt.Println(2)
		return NegInf
	}
	fmt.Println(3)
	return Value(int64(math.Trunc(val * DefaultPow)))
}

func Execute_DefaultFeeRate() {
	// a := NewFromFloat(0.01 * 0.075)

	// a := NewFromFloat(float64(math.Inf(1)))
	// a := NewFromFloat(float64(math.Inf(-1)))
	a := NewFromFloat(float64(math.MaxFloat64 + 1))
	fmt.Println(a)
}

func NewFromString(input string) (Value, error) {
	length := len(input)
	if length == 0 {
		return 0, nil
	}
	isPercentage := input[length-1] == '%'

	if isPercentage {
		input = input[0 : length-1]
	}

	dotIndex := -1
	hasDecimal := false
	decimalCount := 0
	// if is decimal, we don't need this
	hasScientificNotion := false
	hasIChar := false
	// default scientific notation position is -1, equaling nil
	scIndex := -1
	/*
		A rune is used to represent a Unicode character, whereas only ASCII characters
		can be represented solely by an int32 data type.
	*/
	for i, c := range input {
		if hasDecimal {
			if c <= '9' && c >= '0' {
				// U+0030 ~ U+0039
				decimalCount++
			} else {
				break
			}

		} else if c == '.' {
			dotIndex = i
			hasDecimal = true
		}
		if c == 'e' || c == 'E' {
			hasScientificNotion = true
			scIndex = i
			break
		}
		if c == 'i' || c == 'I' {
			hasIChar = true
			break
		}
	}
	if hasDecimal {
		fmt.Println("hasDecimal")
		fmt.Printf("dotIndex: %v\n", dotIndex)
		fmt.Printf("DecimalCount: %v\n", decimalCount)
		// get following number of dot
		after := input[dotIndex+1:]
		if decimalCount >= 8 {
			/*
				ex: 10.00000000
				after * 1e8 -> 1000000000
			*/
			after = after[0:8] + "." + after[8:]
		} else {
			/*
				ex: 1.001e, decimalCount = 3
				001 + 1e(8-decimalCount) + addNotDecimal = 00100000e
			*/
			fmt.Printf("after[0:decimalCount] = %v\n", after[0:decimalCount])
			fmt.Printf("after[decimalCount:] = %v\n", after[decimalCount:])
			after = after[0:decimalCount] + strings.Repeat("0", 8-decimalCount) + after[decimalCount:]
		}
		// get previous number of dot and plus modified after
		input = input[0:dotIndex] + after
		fmt.Printf("input: %v\n", input)
		v, err := strconv.ParseFloat(input, 64)
		fmt.Printf("v = %v\n", v)
		if err != nil {
			return 0, err
		}

		if isPercentage {
			v = v * 0.01
		}

		return Value(int64(math.Trunc(v))), nil

	} else if hasScientificNotion {
		/*
			ex: 1001e+1, scindex = 4
		*/
		fmt.Println("hasScientificNotion")
		fmt.Printf("scIndex: %v\n", scIndex)
		// "+1"
		exp, err := strconv.ParseInt(input[scIndex+1:], 10, 32)
		fmt.Printf("exp: %v\n", exp)
		if err != nil {
			return 0, err
		}
		fmt.Printf("input[0: scIndex+1]: %v\n", input[0:scIndex+1])
		fmt.Printf("strconv.FormatInt(exp+8, 10): %v\n", strconv.FormatInt(exp+8, 10))
		// 1001e9
		v, err := strconv.ParseFloat(input[0:scIndex+1]+strconv.FormatInt(exp+8, 10), 64)
		if err != nil {
			return 0, err
		}
		return Value(int64(math.Trunc(v))), nil
	} else if hasIChar {
		if floatV, err := strconv.ParseFloat(input, 64); nil != err {
			return 0, err
		} else if math.IsInf(floatV, 1) {
			return PosInf, nil
		} else if math.IsInf(floatV, -1) {
			return NegInf, nil
		} else {
			return 0, fmt.Errorf("fixedpoint.Value parse error, invalid input string %s", input)
		}
	} else {
		v, err := strconv.ParseInt(input, 10, 64)
		if err != nil {
			return 0, err
		}
		if isPercentage {
			v = v * DefaultPow / 100
		} else {
			v = v * DefaultPow
		}
		return Value(v), nil
	}

}

func Execute_NewFromString() {
	a, err := NewFromString("1001e+1")
	fmt.Printf("res: %v\n", a)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}
