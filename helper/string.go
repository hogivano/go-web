package helper

import "math/rand"

func RandomString(n int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

	s := make([]rune, n)
	for i := range s {
		s[i] = letter[rand.Intn(len(letter))]
	}

	return string(s)
}
