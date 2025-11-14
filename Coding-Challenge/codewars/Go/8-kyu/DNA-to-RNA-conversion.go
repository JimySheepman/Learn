package kata

import "strings"

func DNAtoRNA(dna string) string {
	for i := 0; i < len(dna); i++ {
		indexFind := strings.Index(dna, "T")
		if indexFind == -1 {
			return dna
		} else {
			dna = dna[:indexFind] + "U" + dna[indexFind+1:]
		}
	}
	return dna
}

// return strings.Replace(dna, "T", "U", -1) best practice
