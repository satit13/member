package model_test

import (
	"testing"
	"fmt"
	"github.com/member/model"

	//"strconv"
)

func TestSale_GetbyMember(t *testing.T) {
	s := model.Sale{Arcode:"990002543"}
	res := s.GetbyMember()
	var sum int
	sum = 0
	for k := range res {
		fmt.Println("docno:", res[k].Docno , "totalamount : ",res[k].Totalamount )
		sum = sum+ int(res[k].Totalamount)
	}

	fmt.Println("sum totalamount is : ",sum)


}