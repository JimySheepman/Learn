package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

const input = "It is not the critic who counts; not the man who points out how the strong man stumbles, or where the doer of deeds could have done them better. The credit belongs to the man who is actually in the arena, whose face is marred by dust and sweat and blood; who strives valiantly; who errs, who comes short again and again, because there is no effort without error and shortcoming; but who does actually strive to do the deeds; who knows great enthusiasms, the great devotions; who spends himself in a worthy cause; who at the best knows in the end the triumph of high achievement, and who at the worst, if he fails, at least fails while daring greatly, so that his place shall never be with those cold and timid souls who neither know victory nor defeat. "

func rune_are_numbers() {
	letter := 'A'
	fmt.Println(letter)
	fmt.Printf("%T \n", letter)
}

func strings_to_rune_conversion() {
	letter := rune("A"[0])
	fmt.Println(letter)
}

func string_index_access() {
	word := "hello"
	letter := rune(word[0])
	fmt.Println(letter)
}

func remainder_bucket_selection() {
	for i := 65; i < 122; i++ {
		fmt.Println(i, " - ", string(rune(i)), " - ", i%12)
	}
}

func hash_function() {
	n := hashBucket("Go", 12)
	fmt.Println(n)
}

func hashBucket(word string, buckets int) int {
	letter := int(word[0])
	bucket := letter % buckets
	return bucket
}

func get() {
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil {
		log.Fatal(err)
	}

	bs, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", bs)
}

func scanners() {
	scanner := bufio.NewScanner(strings.NewReader(input))

	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
}

func moby_dicks_words() {
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()

	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func int_slice_plus_plus() {
	buckets := make([]int, 1)
	fmt.Println(buckets[0])
	buckets[0] = 42
	fmt.Println(buckets[0])
	buckets[0]++
	fmt.Println(buckets[0])
}

func hash_letter_buckets() {
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()

	scanner.Split(bufio.ScanWords)

	buckets := make([]int, 200)

	for scanner.Scan() {
		n := hashBuckets(scanner.Text())
		buckets[n]++
	}
	fmt.Println(buckets[65:123])
}

func hashBuckets(word string) int {
	return int(word[0])
}

func hash_remainder_bucket() {
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()

	scanner.Split(bufio.ScanWords)

	buckets := make([]int, 12)

	for scanner.Scan() {
		n := hashBucket(scanner.Text(), 12)
		buckets[n]++
	}
	fmt.Println(buckets)
}

func main() {
	rune_are_numbers()
	strings_to_rune_conversion()
	string_index_access()
	remainder_bucket_selection()
	hash_function()
	get()
	scanners()
	moby_dicks_words()
	int_slice_plus_plus()
	hash_letter_buckets()
	hash_remainder_bucket()
}
