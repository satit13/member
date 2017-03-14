package model_test

import (
	"testing"
	"fmt"
	"github.com/member/model"

	//"strconv"
)

func TestConnectMsSQL(t *testing.T) {
	_, err := model.ConnectMssql()
	if err != nil {
		fmt.Println("MsSql Connect Error : ", err)
		return
	}
	fmt.Println("connected")

}
func TestSale_GetSalebyMember(t *testing.T) {
	s := model.Sale{Arcode:"990002543"}
	res := s.GetbyMember()
	var sum int
	sum = 0
	for _,k := range res {
		fmt.Println("docno:", k.Docno , "totalamount : ",k.Totalamount )
		sum = sum+ int(k.Totalamount)
	}
	fmt.Println("sum totalamount is : ",sum)
}