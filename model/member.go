package model
import (
	//"github.com/jmoiron/sqlx"
	//"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"

)
type Member struct {
	Id int64
	Name string
	Code string
}



func (m *Member) Create() (id int64 , err error){
	dbconn1 := Connect()

	//&conn.Preparex("select id, from member")
	res,err := dbconn1.Exec(`insert into member(code,name) values(?,?);`,m.Code,m.Name)
	if err != nil{
		fmt.Println(err)
	}
	id ,_ = res.LastInsertId()

	m.Id = id
	return m.Id,nil
}

func (m *Member) Update()(row int64,err error){
	conn := Connect()
	res,err := conn.Exec(`update member set code = ?,name = ? where id = ?`,m.Code,m.Name,m.Id)
	if err != nil{
		fmt.Println(err)
	}
	row , _ = res.RowsAffected()

	return row,err
}


func (m *Member) Reset(){
	conn :=Connect()
	_,err := conn.Exec(`delete from member`)
	if err != nil{
		fmt.Println(err)
	}

	_,err = conn.Exec(`alter table member  auto_increment=1`)
	if err != nil{
		fmt.Println(err)
	}


}

