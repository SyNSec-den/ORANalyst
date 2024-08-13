package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func networkOperation() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer conn.Close()

	// Sending some dummy data
	fmt.Fprintf(conn, "Hello, Server!\n")
}

func fileOperations(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(fmt.Sprintf("Writing line %d\n", i))
	}
	writer.Flush()

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}

func memoryOperations() {
	data := make(map[string][]int)
	for i := 0; i < 5; i++ {
		key := fmt.Sprintf("key_%d", i)
		data[key] = append(data[key], i, i+1, i+2)
	}

	if v, exists := data["key_2"]; exists {
		fmt.Println("Found:", v)
	}
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "net" {
		networkOperation()
	} else if len(os.Args) > 1 && os.Args[1] == "file" {
		fileOperations("test.txt")
	} else {
		memoryOperations()
	}
}
