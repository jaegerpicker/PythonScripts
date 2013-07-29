package main

import (
	"os"
	"bufio"
	"bytes"
	"io"
	"fmt"
	"math"
	"strconv"
)

func readLines(path string) (lines []string, err error) {
	var (
		file *os.File
		part []byte
		prefix bool
		)
	if file, err = os.Open(path); err != nil {
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	buffer := bytes.NewBuffer(make([]byte, 0))
	for {
		if part, prefix, err = reader.ReadLine(); err != nil {
			break
		}
		buffer.Write(part)
		if !prefix {
			lines = append(lines, buffer.String())
			buffer.Reset()
		}
	}
	if err == io.EOF {
		err = nil
	}
	return
}

func doubSq(i float64) {
	result := 0
    for x := 0; x < int(math.Sqrt(i))+1; x++ {
        y := math.Sqrt(float64(int(i) - x*x))
        if y == float64(int(y)) {
            if x*x == int(y) {
                result += 2
            } else {
                result +=1
            }
        }
    }
    fmt.Println(result/2)
}


func main() {
	lines, err := readLines(os.Args[1])
	if err != nil {
		fmt.Println("Error: %s\n", err)
		return
	}
	j := 0
	for _, line := range lines {
		//fmt.Println(line)
		if j > 0 {
			val, err := strconv.ParseFloat(line, 64)
			if err == nil {
				doubSq(val)
			}
		} else {
			j++
		}
		//fmt.Println()
	}

}
