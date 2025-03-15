package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}
func compare(a, b string) int {
	if a == b {
	} else if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}
func ToUpper(s string) string {
	phrase := []rune(s)
	for i := 0; i < len(phrase); i++ {
		currentLetter := phrase[i]
		if currentLetter >= 'a' && currentLetter <= 'z' {
			phrase[i] = phrase[i] - 32
		}
	}
	return string(phrase)
}
func ToLower(s string) string {
	phrase := []rune(s)
	for i := 0; i < len(phrase); i++ {
		currentLetter := phrase[i]
		if currentLetter >= 'A' && currentLetter <= 'Z' {
			phrase[i] = phrase[i] + 32
		}
	}
	return string(phrase)
}
func FirstRune(s string) string {
	t := []rune(s)
	return string(t[0])
}
func Capitalise(s string) string {
	runes := []rune(s)
	strlen := 0
	for i := range runes {
		strlen = i + 1
	}
	for i := 0; i < strlen; i++ {
		if i != 0 && (!((runes[i-1] >= 'a' && runes[i-1] <= 'z') || (runes[i-1] >= 'A' && runes[i-1] <= 'Z'))) {
			if runes[i] >= 'a' && runes[i] <= 'z' {
				runes[i] = rune(runes[i] - 32)
			}
		} else if i == 0 {
			if runes[i] >= 'a' && runes[i] <= 'z' {
				runes[i] = rune(runes[i] - 32)
			}
		} else {
			if runes[i] >= 'A' && runes[i] <= 'Z' {
				runes[i] = rune(runes[i] + 32)
			}
		}
	}
	return string(runes)
}
func split_white_spaces(s string) []string {
	var str []string
	var word string
	l := len(s) - 1
	for i, v := range s {
		if i == l {
			word = word + string(v)
			str = append(str, word)
		} else if v == 32 || v == 15 || v == 10 {
			if s[i+1] == ' ' || s[i+1] == '	' || s[i+1] == 10 {
			} else {
				str = append(str, word)
				word = ""
			}
		} else {
			word = word + string(v)
		}
	}
	return str
}
func quotes(s string) string {
	str := ""
	var removeSpace bool
	for i, char := range s {
		if char == 39 && s[i-1] == ' ' {
			if removeSpace {
				str = str[:len(str)-1]
				str = str + string(char)
				removeSpace = false
			} else {
				str = str + string(char)
				removeSpace = true
			}
		} else if i > 1 && s[i-2] == 39 && s[i-1] == ' ' {
			if removeSpace {
				str = str[:len(str)-1]
				str = str + string(char)
			} else {
				str = str + string(char)
			}
		} else {
			str = str + string(char)
		}
	}
	return str
}
func remove_tags(s []string) string {
	str := ""
	for i, tag := range s {
		if tag == "(cap," || tag == "(low," || tag == "(up," {
			s[i] = ""
			s[i+1] = ""
		} else if tag != "(up)" && tag != "(hex)" && tag != "(bin)" && tag != "(cap)" && tag != "(low)" && tag != "" {
			if i == 0 {
				str = str + tag
			} else {
				str = str + " " + tag
			}
		}
	}
	return str
}
func remove_spaces(s string) string {
	len := len(s) - 1
	if s[len-1] == ' ' {
		return remove_spaces(s[:len])
	}
	return s[:len]
}
func main() {

	filename := os.Args[1]
	data, err := ioutil.ReadFile(filename)
	check(err)
	input := string(data)
	result := split_white_spaces(input)
	for i, v := range result {
		if compare(v, "(hex)") == 0 {
			j, _ := strconv.ParseInt(result[i-1], 16, 64)
			result[i-1] = fmt.Sprint(j)
		}
		if compare(v, "(bin)") == 0 {
			j, _ := strconv.ParseInt(result[i-1], 2, 64)
			result[i-1] = fmt.Sprint(j)
		}
		if compare(v, "(low)") == 0 {
			result[i-1] = ToLower(result[i-1])
		}
		if compare(v, "(low,") == 0 {
			result[i-1] = ToLower(result[i-1])

			le := len(result[i+1])
			numb := result[i+1][:le-1]
			nu, err := strconv.Atoi(numb)
			check(err)
			for j := 1; j <= nu; j++ {
				result[i-j] = ToLower(result[i-j])
			}
		}
		if compare(v, "(up)") == 0 {
			result[i-1] = ToUpper(result[i-1])
		}
		if compare(v, "(up,") == 0 {
			result[i-1] = ToUpper(result[i-1])
			le := len(result[i+1])
			numb := result[i+1][:le-1]
			nu, err := strconv.Atoi(numb)
			check(err)
			for j := 1; j <= nu; j++ {
				result[i-j] = ToUpper(result[i-j])
			}
		}
		if compare(v, "(cap)") == 0 {
			result[i-1] = Capitalise(result[i-1])
		}
		if compare(v, "(cap,") == 0 {
			result[i-1] = Capitalise(result[i-1])
			le := len(result[i+1])
			numb := result[i+1][:le-1]
			nu, err := strconv.Atoi(numb)
			check(err)
			for j := 1; j <= nu; j++ {
				result[i-j] = Capitalise(result[i-j])
			}
		}
		if compare(v, "a") == 0 && FirstRune(result[i+1]) == "a" {
			result[i] = "an"
		}
		if compare(v, "a") == 0 && FirstRune(result[i+1]) == "e" {
			result[i] = "an"
		}
		if compare(v, "a") == 0 && FirstRune(result[i+1]) == "i" {
			result[i] = "an"
		}
		if compare(v, "a") == 0 && FirstRune(result[i+1]) == "o" {
			result[i] = "an"
		}
		if compare(v, "a") == 0 && FirstRune(result[i+1]) == "u" {
			result[i] = "an"
		}
		if compare(v, "a") == 0 && FirstRune(result[i+1]) == "h" {
			result[i] = "an"
		}
	}
	notagResult := remove_tags(result)
	result2 := split_white_spaces(notagResult)
	str := ""
	for _, word := range result2 {
		str = str + word + " "
	}
	str = remove_spaces(str)
	word := ""
	for i, char := range str {
		if i == len(str)-1 {
			if char == '.' || char == ',' || char == '!' || char == '?' || char == ';' || char == ':' {
				if str[i-1] == ' ' {
					word = word[:len(word)-1]
					word = word + string(char)
				} else {
					word = word + string(char)
				}
			} else {
				word = word + string(char)
			}
		} else if char == '.' || char == ',' || char == '!' || char == '?' || char == ';' || char == ':' {
			if str[i-1] == ' ' {
				word = word[:len(word)-1]
				word = word + string(char)
			} else {
				word = word + string(char)
			}
			if str[i+1] != ' ' && str[i+1] != '.' && str[i+1] != ',' && str[i+1] != '!' && str[i+1] != '?' && str[i+1] != ';' && str[i+1] != ':' {
				word = word + " "
			}
		} else {
			word = word + string(char)
		}
	}
	word = quotes(word)
	output := []byte(word)
	words := ioutil.WriteFile(os.Args[2], output, 0644)
	check(words)
}
