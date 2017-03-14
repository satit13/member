package model

import (
	"time"
//	"fmt"
)

type  Sale struct{

	Docno string
	Docdate time.Time
	Totalamount float64
	Point int
	Arcode string
}
type Sales []*Sale

func (s *Sale)GetbyMember()(sl Sales){

	db,err := ConnectMssql()
	rows , err := db.Query("select docno,docdate,totalamount from bcnp.dbo.bcarinvoice where arcode = ?",s.Arcode)

	if err != nil {
		return nil
	}
	defer rows.Close()
	var cts Sales

	for rows.Next() {
		var s = new(Sale)
		rows.Scan(
			&s.Docno,
			&s.Docdate,
			&s.Totalamount,
		)

		cts =  append(cts, s)


	}
	return cts




}



