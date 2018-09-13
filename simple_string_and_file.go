package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"path"
	"strconv"
	"time"
)

func gettmpfile() string {
	// Without seeding you'll get the same "random number" every time
	rand.Seed(time.Now().UnixNano())
	result := path.Join(os.TempDir(), strconv.Itoa(rand.Intn(100))+"Test")

	return result
}

func main() {
	n := 2
	fmt.Printf("print number: %d\n", n)
	fmt.Print(fmt.Sprintf("sprint number: %d\n", n))

	fmt.Fprintf(os.Stderr, "stderr number: %d\n", n)

	fpath := gettmpfile()
	fmt.Printf("Attempting to write to %s \n", fpath)
	file, err := os.OpenFile(fpath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(file, "Hey look i'm in file %s \n", fpath)
	if err := file.Close(); err != nil {
		log.Fatal(err)
	}
}
