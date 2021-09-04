package main

import (
	"strconv"
	"strings"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func main() {
	var plainText, cipherText, key *walk.TextEdit

	MainWindow{
		Title:   "Shift Cipher",
		MinSize: Size{240, 160},
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
	}.Run()
}

var alfabet string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func shiftCipherE(teks string, shift int) string {
	teks = strings.ReplaceAll(teks, " ", "")
	length := len(teks)
	var result string = teks
	var index int
	for i := 0; i < length; i++ {
		for j := 0; j < 26; j++ {
			if rune(result[i]) == rune(alfabet[j]) {
				index = j + shift
				if index > 25 {
					index -= 26
				}
				j = 26
			}
		}
		result = replaceAtIndex(result, rune(alfabet[index]), i)
	}
	return result
}

func shiftCipherD(teks string, shift int) string {
	length := len(teks)
	var result string = teks
	var index int
	for i := 0; i < length; i++ {
		for j := 0; j < 26; j++ {
			if rune(teks[i]) == rune(alfabet[j]) {
				index = j - shift
				if index < 0 {
					index = 26 + index
				}
				j = 26
			}
		}
		result = replaceAtIndex(result, rune(alfabet[index]), i)
	}
	return result
}

func replaceAtIndex(str string, replacement rune, index int) string {
	return str[:index] + string(replacement) + str[index+1:]
}
