package core

type GroupAnalyse struct {
	ascii int
	score int
}

func CrackKey(mg []string, rb []AnalyseStringResult, rl []AnalyseRuneResult) [][]GroupAnalyse {

	pkg := [][]GroupAnalyse{}

	for _, v := range mg {
		gpr := processGroup(v, rb, rl)
		pkg = append(pkg, gpr)
	}

	return pkg
}

func processGroup(g string, rb []AnalyseStringResult, rl []AnalyseRuneResult) []GroupAnalyse {
	gr := []GroupAnalyse{}
	for i := 0x00; i < 0xFF; i++ {
		xg := xorGroup(g, i)
		gr = append(gr, GroupAnalyse{i, calcMatchScore(rb, rl, xg)})
	}

	sortGroupAnalyse(&gr)
	return gr[:2]
}

func sortGroupAnalyse(t *[]GroupAnalyse) {
	for i, v := range *t {
		for a, x := range *t {
			if v.score > x.score {
				c := (*t)[i]
				(*t)[i] = x
				(*t)[a] = c
			}
		}
	}
}

func xorGroup(g string, k int) string {
	xm := []byte{}
	for _, c := range g {
		xm = append(xm, byte(c)^byte(k))
	}

	return string(xm)
}

func calcMatchScore(rb []AnalyseStringResult, rl []AnalyseRuneResult, xg string) int {
	s := 0
	for i := 0; i < len(xg); i++ {
		if (i + 1) < len(xg) {
			for _, b := range rb {
				if string(xg[i:(i+2)]) == b.value {
					s++
				}
			}
		}

		for _, l := range rl {
			if rune(xg[i]) == l.value {
				s++
			}
		}
	}

	return s
}
