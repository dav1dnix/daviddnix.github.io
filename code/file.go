package main

import (
	"io/ioutil"
	"flag"
	"fmt"
	"time"
	"github.com/pkg/errors"
	"math/rand"
)

var (
	// Declare flag file (string) stored in the pointer f with type *string
	file = flag.String("f", "", "File name")

	// What is a rune?: https://stackoverflow.com/a/18311218
	letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

// Random sequence of letters
func randomSequence(n int) string {
	// Create slice using make, a dynamically sized array of "[]rune"s.
	// What make does: https://tour.golang.org/moretypes/13
	b := make([]rune, n)
	// This will iterate over the index in the slice "b".
	for i := range b {
		// https://golang.org/pkg/math/rand/#Intn
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// Read file named by filename and return contents.
func read(filename string) ([]byte, error) {
	rand.Seed(time.Now().UnixNano())
	readfile, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, errors.Wrap(err, "Couldn't read file")
	}
	// Now I need to make the content of the file go into a .html with a random name.
	// This can be done with ioutil.TempFile function https://golang.org/pkg/io/ioutil/#TempFile
	file, err := ioutil.TempFile("./", randomSequence(4))
	if err != nil {
		return nil, errors.Wrap(err, "Could not make temporary file")
	}

	// This will print where the file is stored 
	fmt.Printf("Content: %s\n", file.Name())

	if _, err := file.Write([]byte(readfile)); err != nil {
		return nil, errors.Wrap(err, "Couldn't write to file")
	}

	if err := file.Close(); err != nil {
		return nil, errors.Wrap(err, "Couldn't close file")
	}

	return []byte(readfile), nil
}



func main() {
	// Call Parse() to parse cmd line into defined flags
	flag.Parse()
	read(*file)

}
