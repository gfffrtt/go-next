package html

import (
	"crypto/rand"
	"encoding/base64"
	"maps"
	"strings"
)

func Map[T any, R any](s []T, fn func(T) R) []R {
	out := make([]R, len(s))
	for i, v := range s {
		out[i] = fn(v)
	}
	return out
}

type Entry[T any] struct {
	Key   string
	Value T
}

func Entries[T any](m map[string]T) []Entry[T] {
	out := make([]Entry[T], 0, len(m))
	for k, v := range m {
		out = append(out, Entry[T]{Key: k, Value: v})
	}
	return out
}

func Merge[T any](m1, m2 map[string]T) map[string]T {
	out := make(map[string]T)
	maps.Copy(out, m1)
	maps.Copy(out, m2)
	return out
}

func Id() (string, error) {
	b := make([]byte, 12)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	s := base64.RawURLEncoding.EncodeToString(b)
	return strings.TrimRight(s, "="), nil
}
