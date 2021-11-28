package main

import "fmt"

func intToRoman(num int) string {
	var valueSymbols = []struct {
		value  int
		symbol string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}
	roman := ""
	for _, vs := range valueSymbols {
		for num >= vs.value {
			num -= vs.value
			roman += vs.symbol
		}
		if num == 0 {
			break
		}
	}
	return roman
}
// 最快
func intToRoman2(num int) string {
	m := map[int]string{
		3000: "MMM",
		2000: "MM",
		1000: "M",
		900:  "CM",
		800:  "DCCC",
		700:  "DCC",
		600:  "DC",
		500:  "D",
		400:  "CD",
		300:  "CCC",
		200:  "CC",
		100:  "C",
		90:   "XC",
		80:   "LXXX",
		70:   "LXX",
		60:   "LX",
		50:   "L",
		40:   "XL",
		30:   "XXX",
		20:   "XX",
		10:   "X",
		9:    "IX",
		8:    "VIII",
		7:    "VII",
		6:    "VI",
		5:    "V",
		4:    "IV",
		3:    "III",
		2:    "II",
		1:    "I",
		0:    "",
	}
	return m[num/1000*1000] + m[num%1000/100*100] + m[num%100/10*10] + m[num%10]
}


var (
	thousands = []string{"", "M", "MM", "MMM"}
	hundreds  = []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	tens      = []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	ones      = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
)

// 最优
func intToRoman3(num int) string {
	return thousands[num/1000] + hundreds[num%1000/100] + tens[num%100/10] + ones[num%10]
}

func main() {

	fmt.Println(intToRoman2(2000))
}
