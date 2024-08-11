package models

import (
	//"context"
	//	"crypto/aes"
	"database/sql"
	"errors"
	"math/rand"

	//"fmt"
	//"io"
	"log"
	//"math"
	//"math/big"
	"net/http"
	"strings"
)

type Url struct{
	Domain string
	Path string
	Link string
	Protocol string

}
type UrlPair struct{
	ID int
	//UserId int
	OgUrl string
	NewUrl string  //unique
	Visits int
}

func MigrateUrlPair(db *sql.DB){
	
	_,err:=db.Exec(`CREATE TABLE IF NOT EXISTS urlpair(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		ogurl TEXT,
		newurl TEXT)`)
	if err!=nil{
		log.Fatalln(err)
		
	}


}

func FindUrlPair(db *sql.DB,url string) (*UrlPair,error){
	row:=db.QueryRow("SELECT newurl FROM urlpair WHERE ogurl=? LIMIT 1",url)
	var NewUrl string
//	defer row.Err().Error()
	row.Scan(&NewUrl)

	return &UrlPair{NewUrl: NewUrl},nil

}

// func (u *UrlPair) Save(){

// }
func (u *UrlPair) IsValidUrl() bool{

	url:=u.OgUrl

	if url[:8]=="https://" && len(strings.Split(url, "."))>=2{
		return true
	}
	return false
}
func generate_code(length int) string{ 
	letters:="a b c d e f g h i j k l m n o p q r s t u v w x y z "
	letters_slice:=strings.Split(letters," ")
	path:=""
	for i:=1;i<=length;i++{
		index:=rand.Intn(25)
		path+=letters_slice[index]

	}
	return "/"+path+"/"
}

func (u *UrlPair) GenerateNewUrl(req *http.Request,db *sql.DB) error {
	if !u.IsValidUrl(){
		// var w http.ResponseWriter
		// http.Error(w,"Invalid url",400)
		return errors.New("Invalid Url")

	}
	url:=u.OgUrl


	newurl:=generate_code(4)
	u.NewUrl=newurl

	_,err:=db.Exec("INSERT INTO urlpair (ogurl, newurl) VALUES (?, ?)",url,newurl)
	
	//affected,_:=q.RowsAffected()
	//fmt.Println(affected)

	if err!=nil{
		log.Fatal(err)
	}
	return nil

}





