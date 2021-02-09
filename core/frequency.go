package core

import (
	"bufio"
	"os"
	"sort"

	"strings"
)

type AnalyseStringResult struct {
	value string
	count int
}

type AnalyseRuneResult struct {
	value rune
	count int
}

// Process a list of txt files and return word analyse result, bigram anaylse result and letter anaylse result
func GetReferenceTextMetrics(fn []string) ([]AnalyseStringResult, []AnalyseRuneResult) {

	// bigram map
	bm := make(map[string]int)
	// letter map
	lm := make(map[rune]int)
	// overall lines from reference files
	lines := []string{}

	for _, f := range fn {
		processInputFile(f, &lines)
	}

	processContent(&lines, bm, lm)

	bma := toStringScoreArray(bm, 10)
	lma := toRuneScoreArray(lm, 5)
	return bma, lma
}

func toRuneScoreArray(target map[rune]int, limit int) []AnalyseRuneResult {
	arr := make([]AnalyseRuneResult, 0, len(target))
	for k, v := range target {
		arr = append(arr, AnalyseRuneResult{k, v})
	}

	sort.SliceStable(arr, func(i int, j int) bool {
		return arr[i].count > arr[j].count
	})

	return arr[:limit]
}

func toStringScoreArray(target map[string]int, limit int) []AnalyseStringResult {
	arr := make([]AnalyseStringResult, 0, len(target))
	for k, v := range target {
		arr = append(arr, AnalyseStringResult{k, v})
	}

	sort.SliceStable(arr, func(i int, j int) bool {
		return arr[i].count > arr[j].count
	})

	return arr[:limit]
}

func processInputFile(f string, l *[]string) {
	file, err := os.Open(f)
	Check(err)

	defer file.Close()

	sc := bufio.NewScanner(file)

	for sc.Scan() {
		*l = append(*l, CleanTextLine(sc.Text()))
	}
}

func processContent(lines *[]string, bm map[string]int, lm map[rune]int) {
	for _, line := range *lines {
		words := strings.Split(line, " ")

		for _, word := range words {
			for i := 0; i < len(word); {
				if (i + 1) < len(word) {
					bg := word[i:(i + 2)]
					bm[bg] = bm[bg] + 1
				}

				r := rune(word[i])
				lm[r] = lm[r] + 1
				i++
			}
		}
	}

}
