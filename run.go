// 2>/dev/null; go build -o "${0%.go}.bin" "$0" && exec "${0%.go}.bin"
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello, world")
	os.Exit(1)
}
