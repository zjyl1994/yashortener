package utils

import (
	"encoding/json"
	"math/rand"
	"strconv"
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

type PageMeta struct {
	Prev string
	Next string
	Size string
	List []string
}

func CalcPage(page, size int, total int64) PageMeta {
	var result PageMeta
	totalPage := total / int64(size)
	if total%int64(size) > 0 {
		totalPage++
	}
	result.Size = strconv.Itoa(size)
	if page > 1 {
		result.Prev = strconv.Itoa(page - 1)
	}
	if page < int(totalPage) {
		result.Next = strconv.Itoa(page + 1)
	}
	for i := 1; i <= int(totalPage); i++ {
		start := page - 2
		end := page + 2
		if start < 1 {
			start = 1
			end = 5
		}
		if end > int(totalPage) {
			end = int(totalPage)
			start = end - 4
			if start < 1 {
				start = 1
			}
		}
		for j := start; j <= end; j++ {
			result.List = append(result.List, strconv.Itoa(j))
		}
	}
	return result
}

func ToJson(v any) string {
	b, err := json.Marshal(v)
	if err != nil {
		return err.Error()
	} else {
		return string(b)
	}
}
