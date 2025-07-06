package utils

import (
	"math/rand"
	"strings"
)

func Coalesce[T comparable](vals ...T) T {
	var empty T
	for _, v := range vals {
		if v != empty {
			return v
		}
	}
	return empty
}

func RandChars(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var sb strings.Builder
	for range n {
		sb.WriteByte(letterBytes[rand.Intn(len(letterBytes))])
	}
	return sb.String()
}
