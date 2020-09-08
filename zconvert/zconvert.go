package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "PAYPALISHIRING"
	numRows := 3
	jz := numRows - 1
	curIndex := 0
	var flag bool
	var sub string
	var tt [][]byte

	a := 0
	for {
		lf := (jz - (a % jz)) % jz
		rt := a % jz
		if lf == 0 {
			tp := numRows
			sub, flag = getElement(curIndex, s, tp)
			curIndex = curIndex + numRows
		} else {
			tp := 1
			sub, flag = getElement(curIndex, s, tp)
			curIndex = curIndex + 1
			sub = strings.Repeat(" ", lf) + sub + strings.Repeat(" ", rt)
		}
		//	fmt.Println(lf, rt)
		fmt.Println(sub)
		tt = append(tt, []byte(sub))
		a = a + 1
		if flag {
			break
		}
	}
	fmt.Println(tt)

	//fmt.Println(upserver(tt))
}

func getElement(curIndex int, s string, tp int) (string, bool) {
	if curIndex+tp >= len(s) {
		return s[curIndex:] + strings.Repeat(" ", tp-len(s[curIndex:])), true
	}
	return s[curIndex : curIndex+tp], false

}

func upserver(Aaaay3 [][]byte) string {
	var rArray string
	var srAarry []byte

	for i := 0; i < len(Aaaay3[0]); i++ {
		for j := 0; j < len(Aaaay3); j++ {
			//	fmt.Println(Aaaay3[j][i])
			srAarry = append(srAarry, Aaaay3[j][i])
		}

		rArray = rArray + string(srAarry)
		srAarry = []byte{}
	}
	rArray = strings.Replace(rArray, " ", "", -1)
	return rArray
}
