package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Slurp a whole file into memory
	dat, err := os.ReadFile("./0_hello.go")
	check(err)
	fmt.Println(string(dat))

	f, err := os.Open("./0_hello.go")
	check(err)
	defer f.Close()

	// Read (up to) a precise number of bytes from the file
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	// Seek and read from a particular location in a file
	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	// io package includes more robust input/output functionality
	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3[:n3]))

	// No built in rewind
	_, err = f.Seek(0, 0)
	check(err)

	// bufio implements more ergonomic buffered reads / writes
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	// Scan a file line by line
	_, err = f.Seek(0, 0)
	check(err)

	fileScanner5 := bufio.NewScanner(f)
	fileScanner5.Split(bufio.ScanLines)
	fmt.Println("File contents:")
	for fileScanner5.Scan() {
		fmt.Println(fileScanner5.Text())
	}

	// Scan line by line with line numbers
	_, err = f.Seek(0, 0)
	check(err)

	fileScanner6 := bufio.NewScanner(f)
	fileScanner6.Split(bufio.ScanLines)
	fmt.Println("File contents with line numbers:")
	for i := 0; fileScanner6.Scan(); i++ {
		fmt.Printf("[%d] %s\n", i, fileScanner6.Text())
	}
}
