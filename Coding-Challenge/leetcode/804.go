func uniqueMorseRepresentations(words []string) int {
	morseCodeMap := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	wordMorseCodeMap := map[string]bool{}
	for _, word := range words {
		wordMorseCode := morseCodeMap[word[0]-'a']
		for i := 1; i < len(word); i++ {
			wordMorseCode += morseCodeMap[word[i]-'a']
		}
		wordMorseCodeMap[wordMorseCode] = true
	}
	return len(wordMorseCodeMap)
}