package PrintFile

import (
	"os"
	"bufio"
	"fmt"
)

func PrintFile(filename string){
	file ,err := os.Open(filename)
	if err != nil{
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}