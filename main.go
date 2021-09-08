package main

import (

	// "fmt"

	"log"
	"strconv"
	"strings"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func main() {
	walk.AppendToWalkInit(func() {
		walk.FocusEffect, _ = walk.NewBorderGlowEffect(walk.RGB(0, 63, 255))
		walk.InteractionEffect, _ = walk.NewDropShadowEffect(walk.RGB(63, 63, 63))
		walk.ValidationErrorEffect, _ = walk.NewBorderGlowEffect(walk.RGB(255, 0, 0))
	})

	var mw *walk.MainWindow

	if _, err := (MainWindow{
		AssignTo: &mw,
		Title:    "My Cipher",
		MinSize:  Size{300, 200},
		Layout:   VBox{},
		Children: []Widget{
			PushButton{
				Text: "Shift Cipher",
				OnClicked: func() {
					if _, err := RunShiftDialog(mw); err != nil {
						log.Print(err)
					}
				},
			},
			PushButton{
				Text: "Substitution Cipher",
				OnClicked: func() {
					if _, err := RunSubsDialog(mw); err != nil {
						log.Print(err)
					}
				},
			},
			PushButton{
				Text: "Affine Cipher",
				OnClicked: func() {
					if _, err := RunAffDialog(mw); err != nil {
						log.Print(err)
					}
				},
			},
			PushButton{
				Text: "Vigenere Cipher",
				OnClicked: func() {
					if _, err := RunVigDialog(mw); err != nil {
						log.Print(err)
					}
				},
			},
		},
	}.Run()); err != nil {
		log.Fatal(err)
	}
}

// func main() {
// 	plainText := "fathoni satrio utomo"
// 	shiftKey := 2
// 	keyText := "ZEBRASCDFGHIJKLMNOPQTUVWXY"
// 	vKey := "lmao"
// 	a := 5
// 	b := 8

// 	fmt.Println("Plain text = ", strings.ReplaceAll(strings.ToUpper(plainText), " ", ""))
// 	fmt.Println("Shift E = ", shiftCipherE(plainText, shiftKey))
// 	fmt.Println("Shift D = ", shiftCipherD(shiftCipherE(plainText, shiftKey), shiftKey))
// 	fmt.Println("Subs E = ", subsCipherE(plainText, keyText))
// 	fmt.Println("Subs D = ", subsCipherD(subsCipherE(plainText, keyText), keyText))
// 	fmt.Println("Affine E = ", affCipherE(plainText, a, b))
// 	fmt.Println("Affine D = ", affCipherD(affCipherE(plainText, a, b), a, b))
// 	fmt.Println("Vigenere E =", vigCipherE(plainText, vKey))
// 	fmt.Println("Vigenere D =", vigCipherD(vigCipherE(plainText, vKey), vKey))
// }

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

func RunShiftDialog(owner walk.Form) (int, error) {
	var plainText, cipherText, key *walk.TextEdit

	return Dialog{
		Title:   "Shift Cipher",
		MinSize: Size{300, 300},
		Layout:  VBox{},
		Children: []Widget{
			Label{
				Text: "Plain text",
			},
			TextEdit{
				AssignTo: &plainText,
			},

			Label{
				Text: "Cipher text",
			},
			TextEdit{
				AssignTo: &cipherText,
			},

			Label{
				Text: "Key",
			},
			TextEdit{
				AssignTo: &key,
			},

			Composite{
				Layout: Grid{Columns: 2},
				Children: []Widget{
					PushButton{
						Text: "Enkripsi",
						OnClicked: func() {
							var integer int
							if i, err := strconv.Atoi(key.Text()); err == nil {
								integer = i
							}
							cipherText.SetText(shiftCipherE(strings.ToUpper(plainText.Text()), integer))
						},
					},
					PushButton{
						Text: "Dekripsi",
						OnClicked: func() {
							var integer int
							if i, err := strconv.Atoi(key.Text()); err == nil {
								integer = i
							}
							plainText.SetText(shiftCipherD(strings.ToUpper(cipherText.Text()), integer))
						},
					},
				},
			},
		},
	}.Run(owner)
}

func RunSubsDialog(owner walk.Form) (int, error) {
	var plainText, cipherText, key *walk.TextEdit

	return Dialog{
		Title:   "Substitution Cipher",
		MinSize: Size{300, 300},
		Layout:  VBox{},
		Children: []Widget{
			Label{
				Text: "Plain text",
			},
			TextEdit{
				AssignTo: &plainText,
			},

			Label{
				Text: "Cipher text",
			},
			TextEdit{
				AssignTo: &cipherText,
			},

			Label{
				Text: "Key",
			},
			TextEdit{
				AssignTo: &key,
			},

			Composite{
				Layout: Grid{Columns: 2},
				Children: []Widget{
					PushButton{
						Text: "Enkripsi",
						OnClicked: func() {
							cipherText.SetText(subsCipherE(strings.ToUpper(plainText.Text()), key.Text()))
						},
					},
					PushButton{
						Text: "Dekripsi",
						OnClicked: func() {
							plainText.SetText(subsCipherD(strings.ToUpper(cipherText.Text()), key.Text()))
						},
					},
				},
			},
		},
	}.Run(owner)
}

func RunAffDialog(owner walk.Form) (int, error) {
	var plainText, cipherText, a, b *walk.TextEdit

	return Dialog{
		Title:   "Affine Cipher",
		MinSize: Size{300, 300},
		Layout:  VBox{},
		Children: []Widget{
			Label{
				Text: "Plain text",
			},
			TextEdit{
				AssignTo: &plainText,
			},

			Label{
				Text: "Cipher text",
			},
			TextEdit{
				AssignTo: &cipherText,
			},

			Label{
				Text: "A",
			},
			TextEdit{
				AssignTo: &a,
			},

			Label{
				Text: "B",
			},
			TextEdit{
				AssignTo: &b,
			},

			Composite{
				Layout: Grid{Columns: 2},
				Children: []Widget{
					PushButton{
						Text: "Enkripsi",
						OnClicked: func() {
							var aValue int
							var bValue int
							if i, err := strconv.Atoi(a.Text()); err == nil {
								aValue = i
							}
							if i, err := strconv.Atoi(b.Text()); err == nil {
								bValue = i
							}
							cipherText.SetText(affCipherE(strings.ToUpper(plainText.Text()), aValue, bValue))
						},
					},
					PushButton{
						Text: "Dekripsi",
						OnClicked: func() {
							var aValue int
							var bValue int
							if i, err := strconv.Atoi(a.Text()); err == nil {
								aValue = i
							}
							if i, err := strconv.Atoi(b.Text()); err == nil {
								bValue = i
							}
							plainText.SetText(affCipherD(strings.ToUpper(cipherText.Text()), aValue, bValue))
						},
					},
				},
			},
		},
	}.Run(owner)
}

func RunVigDialog(owner walk.Form) (int, error) {
	var plainText, cipherText, key *walk.TextEdit

	return Dialog{
		Title:   "Vigenere Cipher",
		MinSize: Size{300, 300},
		Layout:  VBox{},
		Children: []Widget{
			Label{
				Text: "Plain text",
			},
			TextEdit{
				AssignTo: &plainText,
			},

			Label{
				Text: "Cipher text",
			},
			TextEdit{
				AssignTo: &cipherText,
			},

			Label{
				Text: "Key",
			},
			TextEdit{
				AssignTo: &key,
			},

			Composite{
				Layout: Grid{Columns: 2},
				Children: []Widget{
					PushButton{
						Text: "Enkripsi",
						OnClicked: func() {
							cipherText.SetText(vigCipherE(strings.ToUpper(plainText.Text()), key.Text()))
						},
					},
					PushButton{
						Text: "Dekripsi",
						OnClicked: func() {
							plainText.SetText(vigCipherD(strings.ToUpper(cipherText.Text()), key.Text()))
						},
					},
				},
			},
		},
	}.Run(owner)
}
