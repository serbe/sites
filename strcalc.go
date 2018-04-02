package sites

import (
	"strconv"
	"strings"
)

func ps2s(ps []string, src string) string {
	src = strings.Replace(src, "ps[", "", -1)
	src = strings.Replace(src, "]", "", -1)
	i, _ := strconv.Atoi(src)
	return ps[i]
}

func mul(src string) string {
	values := strings.Split(src, "*")
	v1, _ := strconv.Atoi(values[0])
	v2, _ := strconv.Atoi(values[1])
	return strconv.Itoa(v1 * v2)
}

func rem(src string) string {
	values := strings.Split(src, "%")
	v1, _ := strconv.Atoi(values[0])
	v2, _ := strconv.Atoi(values[1])
	return strconv.Itoa(v1 % v2)
}

func add(src string) string {
	values := strings.Split(src, "+")
	v1, _ := strconv.Atoi(values[0])
	v2, _ := strconv.Atoi(values[1])
	return strconv.Itoa(v1 + v2)
}

func sub(src string) string {
	values := strings.Split(src, "-")
	v1, _ := strconv.Atoi(values[0])
	v2, _ := strconv.Atoi(values[1])
	return strconv.Itoa(v1 - v2)
}

func brackets(src string) string {
	length := len(src)
	if length < 3 {
		return src
	}
	start := -1
	end := length + 1
	noclose := false
	for i := 0; i < length; i++ {
		if src[i] == '(' {
			start = i
			noclose = true
		}
		if noclose && src[i] == ')' {
			end = i
			noclose = false
		}
	}
	if start > -1 && !noclose {
		return src[start+1 : end]
	}
	return src
}

func findMul(src string) (string, int) {
	return find(src, '*')
}

func findRem(src string) (string, int) {
	return find(src, '%')
}

func findAdd(src string) (string, int) {
	return find(src, '+')
}

func findSub(src string) (string, int) {
	return find(src, '-')
}

func find(src string, b byte) (string, int) {
	length := len(src)
	if length < 3 {
		return "", -1
	}
	pos := 0
	for i := 1; i < length-1; i++ {
		if src[i] == b {
			pos = i
			break
		}
	}
	if pos > 0 {
		return getNearNumbers(src, pos)
	}
	return "", -1
}

func getNearNumbers(src string, pos int) (string, int) {
	leftValue, rightValue, leftPos := getTwoNumber(src, pos)
	return leftValue + string(src[pos]) + rightValue, leftPos
}

func getTwoNumber(src string, pos int) (string, string, int) {
	findLeft := false
	findRight := false
	leftValue := ""
	rightValue := ""
	leftPos := -1
	length := len(src)
	for i := 1; i < length; i++ {
		if pos-i < 0 {
			findLeft = true
		}
		if pos+i > length-1 {
			findRight = true
		}
		if !findLeft {
			if isNumeric(src[pos-i]) {
				leftValue = string(src[pos-i]) + leftValue
				leftPos = pos - i
			} else {
				findLeft = true
			}
		}
		if !findRight {
			if isNumeric(src[pos+i]) {
				rightValue = rightValue + string(src[pos+i])
			} else {
				findRight = true
			}
		}
		if findLeft && findRight {
			break
		}
	}
	return leftValue, rightValue, leftPos
}

func isNumeric(b byte) bool {
	return !(b == '+' || b == '-' || b == '*' || b == '%' || b == '(' || b == ')')
}

func calc(src string) string {
	f := brackets(src)
	if f != src {
		src = strings.Replace(src, "("+f+")", calc(f), 1)
		src = calc(src)
	}

	fm, pm := findMul(src)
	if pm != -1 {
		src = strings.Replace(src, fm, mul(fm), 1)
		src = calc(src)
	}
	fa, pa := findAdd(src)
	fs, ps := findSub(src)
	if pa != -1 && ps != -1 {
		if pa < ps {
			src = strings.Replace(src, fa, add(fa), 1)
			src = calc(src)
		} else {
			src = strings.Replace(src, fs, sub(fs), 1)
			src = calc(src)
		}
	}
	if pa != -1 {
		src = strings.Replace(src, fa, add(fa), 1)
		src = calc(src)
	}
	if ps != -1 {
		src = strings.Replace(src, fs, sub(fs), 1)
		src = calc(src)
	}
	fr, pr := findRem(src)
	if pr != -1 {
		src = strings.Replace(src, fr, rem(fr), 1)
		src = calc(src)
	}
	return src
}

func rotate(src []string, r int) []string {
	r = r % len(src)
	for i := 0; i < r; i++ {
		src = append(src[1:], src[0])
	}
	return src
}
