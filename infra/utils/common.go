package utils

import (
	"encoding/json"
	"fmt"
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

type PageMeta struct {
	Prev string
	Next string
	Size string
	Page string
	List []string
}

func CalcPage(page, size int, total int64) PageMeta {
	var meta PageMeta
	meta.Size = fmt.Sprintf("%d", size)
	meta.Page = fmt.Sprintf("%d", page)

	totalPages := int(total+int64(size)-1) / size // Calculate total pages
	if totalPages == 0 {
		totalPages = 1
	}

	if page > 1 {
		meta.Prev = fmt.Sprintf("%d", page-1)
	}
	if page < totalPages {
		meta.Next = fmt.Sprintf("%d", page+1)
	}

	startPage := max(1, min(page-2, totalPages-4))
	endPage := min(startPage+4, totalPages)

	for i := startPage; i <= endPage; i++ {
		meta.List = append(meta.List, fmt.Sprintf("%d", i))
	}

	return meta
}

func ToJson(v any) string {
	b, err := json.Marshal(v)
	if err != nil {
		return err.Error()
	} else {
		return string(b)
	}
}
