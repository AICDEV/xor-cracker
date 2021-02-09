package core

// ProcessInputMessage discover the input message put as an cmd arg
func ProcessInputMessage(m string, kl int) float32 {
	mf := []string{}

	splitMessageIntoParts(m, kl, &mf)
	return discoverMessageParts(mf, kl)
}

// GetMessageGroups split message into groups of same index
func GetMessageGroups(m string, gl int) []string {
	p := []string{}

	for i := 0; i < gl; i++ {
		g := ""
		c := i
		for c < len(m) {
			g += string(m[c])
			c += gl
		}

		p = append(p, g)
	}

	return p
}

func splitMessageIntoParts(m string, l int, fragments *[]string) {
	for i := 0; i < len(m); i += l {
		if (i + l) < len(m) {
			*fragments = append(*fragments, m[i:i+l])
		}
	}
}

func discoverMessageParts(p []string, kl int) float32 {
	s := 0
	for i := 0; i < 6; i++ {
		if (i + 1) < len(p) {
			s += HammingStringDistance(p[i], p[i+1])
		}
	}

	return (float32(s) / float32(kl))
}
