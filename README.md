# Strmanip: a string manipulation library in Golang

This package contains several algorithms that operate on strings:


  - ~Brute force string search (DONE)~
  - ~KMP string search (DONE)~
  - ~Rabin-Karp string search (DONE)~
  - Boyer-Moore string search (TODO)
  - Suffix Array (Kasai) (TODO)
  - LCP (TODO)
  - Manacher's Algorithm (TODO)
  - Lyndon factorization (TODO)
  - Trie (TODO)

https://www.geeksforgeeks.org/%C2%AD%C2%ADkasais-algorithm-for-construction-of-lcp-array-from-suffix-array/

go test -test.v

https://golang.org/src/strings/search.go

## Usage

You can import the package in the usual way:

```golang
import "github.com/dougct/strmanip"
```

Here's an example of how to perform substring search using the Knuth-Morris-Pratt algorithm:

```golang
index = SubStrSearch("abcxabcdabcdabcy", "abcdabcy", /* method = */ "kmp")
if index != -1 { // we found the pattern in the text
    // do something
}
```

