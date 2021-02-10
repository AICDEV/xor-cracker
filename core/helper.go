package core

import (
	"log"
	"regexp"
	"strings"
)

// KeyScore represents the score of a key
type KeyScore struct {
	KeyLength int
	Score     float32
}

// Check for error
func Check(err error) {
	if err != nil {
		log.Fatalf("error occurred: %v", err)
	}
}

// CleanTextLine Cleans a line of text from all chars that are not [a-zA-z]
func CleanTextLine(l string) string {
	reg, err := regexp.Compile("[^a-zA-z] ")
	Check(err)

	l = strings.ToLower(l)
	l = reg.ReplaceAllString(l, "")
	return l
}

// XOR to bytes
func XOR(a byte, b byte) byte {
	return a ^ b
}

// HammingByteDistance Compute Hamming Distance from two bytes
func HammingByteDistance(a byte, b byte) int {
	c := 0
	r := a ^ b
	for i := 0; i < 8; i++ {
		if (r & 0x01) == 1 {
			c++
		}
		r = r >> 1
	}

	return c
}

// HammingStringDistance Compute Hamming Distance of two strings
func HammingStringDistance(a string, b string) int {
	if len(a) != len(b) {
		log.Fatal("error in calc hamming distance from string. length of strings are not equal")
	}

	d := 0
	for i := 0; i < len(a); i++ {
		d += HammingByteDistance(a[i], b[i])
	}

	return d
}

// GetTopValuesFromKeyMap Get most frequent elements from map by limit
func GetTopValuesFromKeyMap(target map[int]float32, l int) []KeyScore {
	tv := make([]KeyScore, 0, len(target))

	for k, v := range target {
		tv = append(tv, KeyScore{k, v})
	}

	for i, v := range tv {
		for a, x := range tv {
			if v.Score < x.Score {
				c := tv[i]
				tv[i] = x
				tv[a] = c
			}
		}
	}

	return tv[:l]
}

// GetStringKeys Print GroupAnalyse ascii to string in terminal
func GetStringKeys(k [][]GroupAnalyse) (string, string) {
	ko := ""
	kt := ""

	for _, v := range k {
		ko += string(v[0].ascii)
		kt += string(v[1].ascii)
	}

	return ko, kt
}
