package strmanip

import (
	"math/rand"
	"reflect"
	"strings"
	"testing"
	"time"
)

var prefixTests = []struct {
	pattern  string // input
	expected []int  // expected result
}{
	{"aabaaab", []int{0, 1, 0, 1, 2, 2, 3}},
	{"abcabcd", []int{0, 0, 0, 1, 2, 3, 0}},
	{"abababcaab", []int{0, 0, 1, 2, 3, 4, 0, 1, 1, 2}},
}

func TestPrefixFunction(t *testing.T) {
	for _, tt := range prefixTests {
		got := prefixFunction(tt.pattern)
		if !reflect.DeepEqual(got, tt.expected) {
			t.Errorf("prefixFunction(%s): got %v, want %v", tt.pattern, got, tt.expected)
		}
	}
}

const letters = "abcdefghijklmnopqrstuvwxyz"

func generateRandomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func TestSearch(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	methods := []string{"brute-force", "kmp", "rabin-karp"}
	for _, method := range methods {
		for i := 0; i < 100; i++ {
			pattern := generateRandomString(i)
			for j := 0; j < 1000; j++ {
				text := generateRandomString(j)
				want := strings.Index(text, pattern)
				if got := SubStrSearch(text, pattern, method); got != want {
					t.Errorf("SubStrSearch(%s, %s, %s): got %d, want %d", text, pattern, method, got, want)
				}
			}
		}
	}
}
