package kinaco

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

type kinaco struct {
	delimiter string
	num       int
}

func Generate() string {
	kinaco := kinaco{
		delimiter: "-",
		num:       3,
	}
	return kinaco.generate()
}

func Delimiter(delimiter string) *kinaco {
	kinaco := kinaco{
		delimiter: delimiter,
	}
	return &kinaco
}

func Num(num int) *kinaco {
	kinaco := kinaco{
		num: num,
	}
	return &kinaco
}

func (k *kinaco) Delimiter(delimiter string) *kinaco {
	k.delimiter = delimiter
	return k
}

func (k *kinaco) Num(num int) *kinaco {
	k.num = num
	return k
}

func (k *kinaco) Generate() string {
	if k.delimiter == "" {
		k.delimiter = "-"
	}

	if k.num == 0 {
		k.num = 3
	}

	return k.generate()
}

func (k *kinaco) generate() string {
	words := readLine("words.txt")
	wordNum := len(words)
	rand.Seed(time.Now().UnixNano())

	var res []string
	for range make([]int, k.num) {
		res = append(res, words[rand.Intn(wordNum)])
	}
	return strings.Join(res, k.delimiter)
}

func readLine(filename string) []string {
	_, b, _, _ := runtime.Caller(1)
	basePath := filepath.Dir(b)
	file, err := os.Open(filepath.Join(basePath, filename))
	if err != nil {
		log.Fatalf("file can not open: %v", err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal("file can not close")
		}
	}(file)

	var res []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}

	return res
}
