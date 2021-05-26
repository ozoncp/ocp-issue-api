package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func readFile(filename string, times uint64) {
	for i := uint64(0); i < times; i++ {
		func() {
			file, err := os.Open(filename)

			if err != nil {
				fmt.Printf("error: %s", err)
			}

			_, err = io.Copy(os.Stdout, bufio.NewReader(file))

			if err != nil {
				fmt.Printf("error: %s", err)
			}

			defer func() {
				err = file.Close()

				if err != nil {
					fmt.Printf("error: %s", err)
				}
			}()
		}()
	}
}

func main() {
	fmt.Println("ocp-issue-api")

	if len(os.Args) > 2 {
		filename := os.Args[1]

		times, err := strconv.ParseUint(os.Args[2], 10, 64)

		if err != nil {
			fmt.Printf("Error: %s. Usage: %s <filename: string> <times_to_repeat: int>", err, os.Args[0])
		} else {
			readFile(filename, times)
		}
	} else {
		fmt.Printf("Usage: %s <filename: string> <times_to_repeat: int>", os.Args[0])
	}
}
