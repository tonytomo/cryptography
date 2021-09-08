package main

import (
	// "strconv"
	"fmt"
	"strings"
	// "github.com/lxn/walk"
	// . "github.com/lxn/walk/declarative"
)

// func main() {
// 	var plainText, cipherText, key *walk.TextEdit

// 	MainWindow{
// 		Title:   "Cipher",
// 		MinSize: Size{240, 160},
// 		Layout:  VBox{},
// 		Children: []Widget{
// 			Label{
// 				Text: "Plain text",
// 			},
// 			TextEdit{
// 				AssignTo: &plainText,
// 			},

// 			Label{
// 				Text: "Cipher text",
// 			},
// 			TextEdit{
// 				AssignTo: &cipherText,
// 			},

// 			Label{
// 				Text: "Key",
// 			},
// 			TextEdit{
// 				AssignTo: &key,
// 			},

// 			Composite{
// 				Layout: Grid{Columns: 2},
// 				Children: []Widget{
// 					PushButton{
// 						Text: "Enkripsi",
// 						OnClicked: func() {
// 							var integer int
// 							if i, err := strconv.Atoi(key.Text()); err == nil {
// 								integer = i
// 							}
// 							cipherText.SetText(shiftCipherE(strings.ToUpper(plainText.Text()), integer))
// 						},
// 					},
// 					PushButton{
// 						Text: "Dekripsi",
// 						OnClicked: func() {
// 							var integer int
// 							if i, err := strconv.Atoi(key.Text()); err == nil {
// 								integer = i
// 							}
// 							plainText.SetText(shiftCipherD(strings.ToUpper(cipherText.Text()), integer))
// 						},
// 					},
// 				},
// 			},
// 		},
// 	}.Run()
// }

func main() {
	plainText := "fathoni satrio utomo"
	shiftKey := 2
	keyText := "ZEBRASCDFGHIJKLMNOPQTUVWXY"
	vKey := "lmao"
	a := 5
	b := 8

	fmt.Println("Plain text = ", strings.ReplaceAll(strings.ToUpper(plainText), " ", ""))
	fmt.Println("Shift E = ", shiftCipherE(plainText, shiftKey))
	fmt.Println("Shift D = ", shiftCipherD(shiftCipherE(plainText, shiftKey), shiftKey))
	fmt.Println("Subs E = ", subsCipherE(plainText, keyText))
	fmt.Println("Subs D = ", subsCipherD(subsCipherE(plainText, keyText), keyText))
	fmt.Println("Affine E = ", affCipherE(plainText, a, b))
	fmt.Println("Affine D = ", affCipherD(affCipherE(plainText, a, b), a, b))
	fmt.Println("Vigenere E =", vigCipherE(plainText, vKey))
	fmt.Println("Vigenere D =", vigCipherD(vigCipherE(plainText, vKey), vKey))
}

var alfabet string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func replaceAtIndex(str string, replacement rune, index int) string {
	return str[:index] + string(replacement) + str[index+1:]
}

func shiftCipherE(teks string, shift int) string {
	result := strings.ReplaceAll(strings.ToUpper(teks), " ", "")
	length := len(result)
	var index int
	flag := false
	for i := 0; i < length; i++ {
		for j := 0; j < 26; j++ {
			if rune(result[i]) == rune(alfabet[j]) {
				index = j + shift
				if index > 25 {
					index -= 26
				}
				j = 26
				flag = true
			}
		}
		if flag {
			result = replaceAtIndex(result, rune(alfabet[index]), i)
			flag = false
		}
	}
	return result
}

func shiftCipherD(teks string, shift int) string {
	result := strings.ReplaceAll(strings.ToUpper(teks), " ", "")
	length := len(teks)
	var index int
	flag := false
	for i := 0; i < length; i++ {
		for j := 0; j < 26; j++ {
			if rune(result[i]) == rune(alfabet[j]) {
				index = j - shift
				if index < 0 {
					index = 26 + index
				}
				j = 26
				flag = true
			}
		}
		if flag {
			result = replaceAtIndex(result, rune(alfabet[index]), i)
			flag = false
		}
	}
	return result
}

func subsCipherE(teks string, key string) string {
	result := strings.ReplaceAll(strings.ToUpper(teks), " ", "")
	length := len(result)
	key = strings.ToUpper(key)
	var index int
	flag := false
	for i := 0; i < length; i++ {
		for j := 0; j < 26; j++ {
			if rune(result[i]) == rune(alfabet[j]) {
				index = j
				j = 26
				flag = true
			}
		}
		if flag {
			result = replaceAtIndex(result, rune(key[index]), i)
			flag = false
		}
	}
	return result
}

func subsCipherD(teks string, key string) string {
	result := strings.ReplaceAll(strings.ToUpper(teks), " ", "")
	length := len(result)
	key = strings.ToUpper(key)
	var index int
	flag := false
	for i := 0; i < length; i++ {
		for j := 0; j < 26; j++ {
			if rune(result[i]) == rune(key[j]) {
				index = j
				j = 26
				flag = true
			}
		}
		if flag {
			result = replaceAtIndex(result, rune(alfabet[index]), i)
			flag = false
		}
	}
	return result
}

func affCipherE(teks string, a int, b int) string {
	result := strings.ToUpper(strings.ReplaceAll(teks, " ", ""))
	length := len(result)
	var index int
	flag := false
	for i := 0; i < length; i++ {
		for j := 0; j < 26; j++ {
			if rune(result[i]) == rune(alfabet[j]) {
				index = (a*j + b) % 26
				j = 26
				flag = true
			}
		}
		if flag {
			result = replaceAtIndex(result, rune(alfabet[index]), i)
			flag = false
		}
	}
	return result
}

func affCipherD(teks string, a int, b int) string {
	result := strings.ToUpper(strings.ReplaceAll(teks, " ", ""))
	length := len(result)
	var index int
	flag := false
	for i := 0; i < length; i++ {
		for j := 0; j < 26; j++ {
			if rune(result[i]) == rune(alfabet[j]) {
				dec := (26 - a) * (j - b)
				index = (dec%26 + 26) % 26
				if index < 0 {
					index *= -1
				}
				j = 26
				flag = true
			}
		}
		if flag {
			result = replaceAtIndex(result, rune(alfabet[index]), i)
			flag = false
		}
	}
	return result
}

func vigCipherE(teks string, key string) string {
	result := strings.ToUpper(strings.ReplaceAll(teks, " ", ""))
	length := len(result)

	keylen := len(key)
	var newkey string
	for i := 0; i < ((length-(length%keylen))/keylen)+1; i++ {
		newkey = newkey + key
	}
	newkey = strings.ToUpper(newkey)
	var shift []int
	for i := 0; i < length; i++ {
		for j := 0; j < 26; j++ {
			if rune(newkey[i]) == rune(alfabet[j]) {
				shift = append(shift, j)
				j = 26
			}
		}
	}
	var index int
	flag := false
	for i := 0; i < length; i++ {
		for j := 0; j < 26; j++ {
			if rune(result[i]) == rune(alfabet[j]) {
				index = j + shift[i]
				if index > 25 {
					index -= 26
				}
				j = 26
				flag = true
			}
		}
		if flag {
			result = replaceAtIndex(result, rune(alfabet[index]), i)
			flag = false
		}
	}
	return result
}

func vigCipherD(teks string, key string) string {
	result := strings.ToUpper(strings.ReplaceAll(teks, " ", ""))
	length := len(result)

	keylen := len(key)
	var newkey string
	for i := 0; i < ((length-(length%keylen))/keylen)+1; i++ {
		newkey = newkey + key
	}
	newkey = strings.ToUpper(newkey)
	var shift []int
	for i := 0; i < length; i++ {
		for j := 0; j < 26; j++ {
			if rune(newkey[i]) == rune(alfabet[j]) {
				shift = append(shift, j)
				j = 26
			}
		}
	}
	var index int
	flag := false
	for i := 0; i < length; i++ {
		for j := 0; j < 26; j++ {
			if rune(result[i]) == rune(alfabet[j]) {
				index = j - shift[i]
				if index < 0 {
					index = 26 + index
				}
				j = 26
				flag = true
			}
		}
		if flag {
			result = replaceAtIndex(result, rune(alfabet[index]), i)
			flag = false
		}
	}
	return result
}
