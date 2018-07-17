package ConvertToBin

import (
	"strconv"
	"os"
	"bufio"
	"fmt"
)

func ConvertToBin(n int) string{
	if n==0 {
		return strconv.Itoa(0)
	}
	res := ""
	for ; n>0 ; n/=2 {
		lsb := n % 2
		res = strconv.Itoa(lsb) + res
	}
	return res;
}





