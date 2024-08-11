package shortner

import (
	"database/sql"
	"fmt"
	//"fmt"
	_ "log"
	"net/http"

	"github.com/Ashutosh1910/url-shortner/models"
	"github.com/Ashutosh1910/url-shortner/templates"
)
type Router struct{
	Service *ShortnerService
}

func  NewRouter() *Router{
	return &Router{

	}
 
}
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request){
	//first check for registered paths and then for redirect paths
	// ++visits

	if req.URL.Path=="/"{
		r.Service.HomeForm(w,req)
	}else if  req.URL.Path=="/shorten" ||req.URL.Path=="/shorten/"{
		r.Service.Shorten(w,req)
	}else{

		r.Service.FindHandler(w,req)
	}
	
	

	//http.NotFound(w,req)

	



}

type ShortnerService struct{
	Db *sql.DB
}

func (s *ShortnerService)Shorten(w http.ResponseWriter, req *http.Request){

	oldUrl:=req.FormValue("name")
	fmt.Println(oldUrl)
	urlpair:=&models.UrlPair{OgUrl: oldUrl}
	urlpair.GenerateNewUrl(req,s.Db)
	// if err!=nil{
	// 	log.Fatalln(err)
	// }
	newurlpair,_:=models.FindUrlPair(s.Db,oldUrl)
	component:=templates.ResultPage(newurlpair.NewUrl)
	component.Render(req.Context(),w)
 
}

func (s *ShortnerService)HomeForm(w http.ResponseWriter, req *http.Request){

	component:=templates.FormPage()
	component.Render(req.Context(),w)

}


func (s *ShortnerService)FindHandler(w http.ResponseWriter, req *http.Request){

	rows,err:=s.Db.Query("SELECT ogurl FROM urlpair WHERE newurl= ? LIMIT 1",req.URL.Path)
	defer rows.Close()
	
	fmt.Println(req.URL.Path)
	if err==nil &&rows.Next(){
		var url string
		fmt.Println(rows)
		rows.Scan(&url)
		fmt.Println(url)
		http.Redirect(w,req,url,http.StatusMovedPermanently)
    }else{
		fmt.Println(rows.Next())
		http.NotFound(w,req)
	}

}

