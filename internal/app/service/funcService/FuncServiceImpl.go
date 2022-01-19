package funcService

import (
	"math"
	"strconv"
	"strings"
)

type funcService struct {
}

func NewFuncService() IFuncService {
	return &funcService{}
}

func (f funcService) BinaryToDecimal(inp string) int {
	des := 0
	count := 0.0
	idx := 0
	num, _ := strconv.Atoi(inp)
	for num != 0 {
		idx = num % 10
		des += idx * int(math.Pow(2.0, count))
		num = num / 10
		count++
	}
	return des
}
func (f funcService) DecimalToBinary(inp string) int {
	bin := 0
	count := 1
	idx := 0
	num, _ := strconv.Atoi(inp)
	for num != 0 {
		idx = num % 2
		num = num / 2
		bin += idx * count
		count *= 10
	}
	return bin
}
func (f funcService) Polyndrome(inp string) string {
	result1 := ""
	result2 := ""

	words := strings.Fields(inp)
	for i := 0; i < len(words); i++ {
		chars := []rune(words[i])
		idx1 := 0
		if i != len(words)-1 {
			idx1 = i + 1
		} else {
			break
		}
		tamp1 := ""
		tamp2 := ""

		indexmundur := 1
		x := 0
		for {
			KataPembanding := []rune(words[idx1])
			if len(KataPembanding)-indexmundur < 0 || len(chars) == x {
				break
			}
			if chars[x] == KataPembanding[len(KataPembanding)-indexmundur] {
				tamp1 += string(chars[x])
				tamp2 = string(KataPembanding[len(KataPembanding)-indexmundur]) + tamp2
				x++
			}
			indexmundur++
		}
		if len(result1) < len(tamp1) {
			result1 = tamp1
			result2 = tamp2
		}
	}
	return result1 + " " + result2
}
