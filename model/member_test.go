
package model_test

import (
	"testing"
	"github.com/member/model"
	"fmt"

	//"github.com/revel/modules/db/app"
)


func TestMember_Create(t *testing.T){
	//arrange
	m1 := &model.Member{Name: "satit", Code: "0816398388"}
	fmt.Println(m1)

	//act
	m1.Reset()

	id,err := m1.Create()

	fmt.Println("new id is : " , id)
	fmt.Println(m1)
	if err != nil {
		t.Error(err)
	}
}

func TestMember_Update(t *testing.T){
	m1 := &model.Member{Id:1,Name: "satitx", Code: "0816398388"}
	fmt.Println(`update to `,m1)


	//act
	//m1.Reset()

	row,err := m1.Update()

	fmt.Println("row effected is  : " , row)
	fmt.Println(m1)
	if err != nil {
		t.Error(err)
	}
}


func TestMember_DuplicateName(t *testing.T){
	m1 := &model.Member{Name: "satit", Code: "0816398388"}
	fmt.Println(m1)

	//act
	m1.Reset()


	//fist insert code 0816398388
	id,err := m1.Create()
	fmt.Println("new id is : " , id)
	fmt.Println(m1)
	if err != nil {
		t.Error(err)
	}

	//fist insert code 0816398388
	m1.Code = "999"
	id,err = m1.Create()
	fmt.Println("new id is : " , id)
	fmt.Println(m1)
	if err!=nil {
		t.Error(err)
	}


}


func TestMember_DuplicateCode(t *testing.T){
	m1 := &model.Member{Name: "satit", Code: "0816398388"}
	fmt.Println(m1)

	//act
	m1.Reset()


	//fist insert code 0816398388
	id,err := m1.Create()
	fmt.Println("new id is : " , id)
	fmt.Println(m1)
	if err != nil {
		t.Error(err)
	}

	//fist insert code 0816398388
	m1.Name = "999"
	id,err = m1.Create()
	fmt.Println("new id is : " , id)
	fmt.Println(m1)
	if err != nil {
		t.Error(err)
	}


}
