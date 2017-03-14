package model

import (
	_ "database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"flag"
	"log"
	"database/sql"
	"github.com/jmoiron/sqlx"
)


var debug = flag.Bool("debug", false, "enable debugging")
var password = flag.String("password", "[ibdkifu", "the database password")
var port *int = flag.Int("port", 1433, "the database port")
var server = flag.String("server", "nebula2", "the database server")
var user = flag.String("user", "sa", "the database user")

func NewDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mssql", dsn)
	if err != nil {

		log.Panic("sql.Open() Error >> ", err)
		return nil , err

	}
	if err = db.Ping(); err != nil {
		log.Panic("db.Ping() Error >> ", err)
		return nil, err
	}
	log.Println("db=", db)
	return db,nil

}
func ConnectMssql()(db *sql.DB,err error){
	flag.Parse() // parse the command line args

	if *debug {
		fmt.Printf(" password:%s\n", *password)
		fmt.Printf(" port:%d\n", *port)
		fmt.Printf(" server:%s\n", *server)
		fmt.Printf(" user:%s\n", *user)
	}
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d", *server, *user, *password, *port)
	fmt.Println(connString)
	if *debug {
		fmt.Printf(" connString:%s\n", connString)
	}

	db, err = NewDB(connString)


	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}


	//defer db.Close()



	fmt.Println("connected from init.go")
	return db , nil
}
func Connect()(conn *sqlx.DB){
	dsn := "satit:[ibdkifu@tcp(192.168.0.89:3306)/satit"
	conn = sqlx.MustConnect("mysql",dsn)
	return conn
}
