package goutil

// Encrypt is an unsafe option for encrypting Strings.
func Encrypt(key []rune, text string) string {
	currKeyByte := 0
	textRuneList := []rune(text)
	for i, textByte := range []rune(text) {
		textRuneList[i] = textByte + key[currKeyByte]
		currKeyByte++
		if currKeyByte >= len(key) {
			currKeyByte = 0
		}
	}
	return string(textRuneList)
}

// Decrypt is the function to decrypt data encrypted by Encrypt
func Decrypt(key []rune, text string) string {
	currKeyByte := 0
	textRuneList := []rune(text)
	for i, textByte := range []rune(text) {
		textRuneList[i] = textByte - key[currKeyByte]
		currKeyByte++
		if currKeyByte >= len(key) {
			currKeyByte = 0
		}
	}
	return string(textRuneList)
}
