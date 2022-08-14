func toArray(set map[string]struct{}) []string {
	var vs []string
	for v := range set {
		vs = append(vs, v)
	}

	return vs
}

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	found := false
	for _, w := range wordList {
		if w == endWord {
			found = true
			break
		}
	}
	if !found {
		return nil
	}

	// Construct buckets of words that only differ by one letter
	fullList := append([]string{}, wordList...)
	fullList = append(fullList, beginWord)
	buckets := map[string]map[string]struct{}{}
	for i := 0; i < len(beginWord); i++ {
		for _, w := range fullList {
			underscore := fmt.Sprintf("%s_%s", w[0:i], w[i+1:len(w)])
			if buckets[underscore] == nil {
				buckets[underscore] = map[string]struct{}{}
			}
			buckets[underscore][w] = struct{}{}
		}
	}

	// Map word masks (one letter obscured by an underscore) to corresponding words
	graph := map[string][]string{}
	for _, b := range buckets {
		words := toArray(b)
		for i, w := range words {
			related := append([]string{}, words[0:i]...)
			related = append(related, words[i+1:len(words)]...)
			graph[w] = append(graph[w], related...)
		}
	}

	dict := map[string]struct{}{}
	for _, w := range wordList {
		dict[w] = struct{}{}
	}

	// The layer maps all encountered words to sequences leading up to them
	layer := map[string][][]string{
		beginWord: {
			{beginWord},
		},
	}

	// Layer by layer, derive word sequences until we arrive at the end word
	for len(layer) > 0 {
		newLayer := map[string][][]string{}
		for w := range layer {
			if w == endWord {
				// Return all possible sequences
				return layer[w]
			}

			for _, nw := range graph[w] {
				if _, ok := dict[nw]; !ok {
					continue
				}

				// Add all sequences leading to current word + new word
				var toAppend [][]string
				// Append new word to all sequences corresponding to word in this layer
				for _, j := range layer[w] {
					seq := append([]string{}, j...)
					toAppend = append(toAppend, append(seq, nw))
				}
				// Add new word to all sequences and form new layer element
				newLayer[nw] = append(newLayer[nw], toAppend...)
			}
		}

		// Prune dictonary to prevent loops
		for k := range newLayer {
			delete(dict, k)
		}
		// Move down to new layer
		layer = newLayer
	}

	return nil
}