package strmanip

// bruteForce performs substring search using a brute force algorithm.
func bruteForce(text string, pattern string) int {
	// An empty pattern is always present, regardless of the text.
	if len(pattern) == 0 {
		return 0
	}

	for i := 0; i <= len(text)-len(pattern); i++ {
		j := 0
		for j < len(pattern) {
			if text[i+j] != pattern[j] { // no match, advance in the text
				break
			}
			j++
		}
		if j == len(pattern) { // match
			return i
		}
	}

	return -1
}

// prefixFunction implements the prefix function of the Knuth-Morris-Pratt algorithm.
func prefixFunction(pattern string) []int {
	prefix := make([]int, len(pattern))
	for i := 1; i < len(pattern); i++ {
		border := prefix[i-1]
		// Tricky case: The next character in the pattern differs
		// from the next character in the border. Thus, we have to
		// compare the next character of the pattern with the next
		// character of each previous border.
		for border > 0 && pattern[i] != pattern[border] {
			border = prefix[border-1]
		}
		// Trivial case: The next character in the pattern is
		// equal to the next character of the border, so we
		// just increase the border length.
		if pattern[i] == pattern[border] {
			border++
		}
		prefix[i] = border
	}
	return prefix
}

// KMPSearch performs string search using the Knuth-Morris-Pratt algorithm.
// It returns the index of the first occurrence of the pattern in the text,
// or -1 if the pattern is not present in the text.
func kmp(text string, pattern string) int {
	// An empty pattern is always present in the text.
	if len(pattern) == 0 {
		return 0
	}

	prefix := prefixFunction(pattern)
	i, j := 0, 0
	for i < len(text) {
		// Started a match: advance pointer on both text and pattern.
		if text[i] == pattern[j] {
			i++
			j++
		} else {
			if j != 0 {
				// Backtrack in the pattern according to the prefix function.
				j = prefix[j-1]
			} else {
				// Found a mismatch at the beginning of the pattern, so
				// advance only the text pointer.
				i++
			}
		}
		// Found an occurrence of the pattern in the text.
		if j == len(pattern) {
			return i - len(pattern)
		}
	}
	return -1
}

// primeRK is the prime base used in Rabin-Karp algorithm.
const primeRK = 16777619

// hashStr returns the hash and the appropriate multiplicative
// factor for use in Rabin-Karp algorithm.
func hashStr(sep string) (uint32, uint32) {
	hash := uint32(0)
	for i := 0; i < len(sep); i++ {
		hash = hash*primeRK + uint32(sep[i])
	}
	var pow, sq uint32 = 1, primeRK
	for i := len(sep); i > 0; i >>= 1 {
		if i&1 != 0 {
			pow *= sq
		}
		sq *= sq
	}
	return hash, pow
}

func rabinKarp(s, substr string) int {
	// Rabin-Karp search
	hashss, pow := hashStr(substr)
	n := len(substr)
	var h uint32
	for i := 0; i < n; i++ {
		h = h*primeRK + uint32(s[i])
	}
	if h == hashss && s[:n] == substr {
		return 0
	}
	for i := n; i < len(s); {
		h *= primeRK
		h += uint32(s[i])
		h -= pow * uint32(s[i-n])
		i++
		if h == hashss && s[i-n:i] == substr {
			return i - n
		}
	}
	return -1
}

func boyerMoore(text string, pattern string) int {
	skip := 0
	for i := 0; i < len(text)-len(pattern); i += skip {
		skip = 0
		for j := len(pattern) - 1; j >= 0; j-- {
			if pattern[j] != text[i+j] {
				//skip = int(math.Max(1, j-right[i+j]))
				break
			}
		}
		if skip == 0 {
			return i
		}
	}
	return -1
}

// SubStrSearch searches for pattern in text using the specified method.
func SubStrSearch(text string, pattern string, method string) int {
	switch n := len(pattern); {
	case n == 0:
		return 0
	case n == len(text):
		if pattern == text {
			return 0
		}
		return -1
	case n > len(text):
		return -1
	}

	switch method {
	case "brute-force":
		return bruteForce(text, pattern)
	case "kmp":
		return kmp(text, pattern)
	case "rabin-karp":
		return rabinKarp(text, pattern)
	case "boyer-moore":
		return boyerMoore(text, pattern)
	default:
		panic("available methods: brute-force, kmp, rabin-karp, boyer-moore.")
	}
}
