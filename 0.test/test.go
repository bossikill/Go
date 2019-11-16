package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
)

func init() {

}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Println(path.Ext(input.Text()))
	}
}
