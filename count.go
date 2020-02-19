package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		kata :=strings.Split(scanner.Text()," ")
		hitung :=make(map[string]int)
		for i := range kata {
			hitung[kata[i]]++	
	}
	fmt.Println(hitung);
	}
}