package kata

func latest(n int) int {
	nCopy := n
	sum := 0
	for i := 0; i < 4; i++ {
		sum += n % 10
		n /= 10
	}
	newNumber := nCopy - sum
	return newNumber
}

func SubtractSum(n int) string {
	n = latest(n)
	if n > 100 {
		newNumber := latest(n)
		return SubtractSum(newNumber)
	} else {
		fruitArray := []string{
			"kiwi",
			"pear",
			"kiwi",
			"banana",
			"melon",
			"banana",
			"melon",
			"pineapple",
			"apple",
			"pineapple",
			"cucumber",
			"pineapple",
			"cucumber",
			"orange",
			"grape",
			"orange",
			"grape",
			"apple",
			"grape",
			"cherry",
			"pear",
			"cherry",
			"pear",
			"kiwi",
			"banana",
			"kiwi",
			"apple",
			"melon",
			"banana",
			"melon",
			"pineapple",
			"melon",
			"pineapple",
			"cucumber",
			"orange",
			"apple",
			"orange",
			"grape",
			"orange",
			"grape",
			"cherry",
			"pear",
			"cherry",
			"pear",
			"apple",
			"pear",
			"kiwi",
			"banana",
			"kiwi",
			"banana",
			"melon",
			"pineapple",
			"melon",
			"apple",
			"cucumber",
			"pineapple",
			"cucumber",
			"orange",
			"cucumber",
			"orange",
			"grape",
			"cherry",
			"apple",
			"cherry",
			"pear",
			"cherry",
			"pear",
			"kiwi",
			"pear",
			"kiwi",
			"banana",
			"apple",
			"banana",
			"melon",
			"pineapple",
			"melon",
			"pineapple",
			"cucumber",
			"pineapple",
			"cucumber",
			"apple",
			"grape",
			"orange",
			"grape",
			"cherry",
			"grape",
			"cherry",
			"pear",
			"cherry",
			"apple",
			"kiwi",
			"banana",
			"kiwi",
			"banana",
			"melon",
			"banana",
			"melon",
			"pineapple",
			"apple",
			"pineapple"}
		return fruitArray[n-1]
	}
}
