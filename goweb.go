package main
import (
	"fmt"
	// "log"
	// "html/template"
	// "github.com/gin-gonic/gin"
	// "net/http"
	
)
type IndexData struct {
	Title   string
	Content string
}
var UserData map[string]string

func init() {
	UserData = map[string]string{
		// "test": "test",
		"Daniel": "honte511",
	}
}

// func index(c *gin.Context) {
// 	data := new(IndexData)
// 	data.Title = "首頁"
// 	data.Content = "我的第一支 gin 專案"
// 	c.HTML(http.StatusOK, "index.html", data)
// }
func main() {
	// http.HandleFunc("/", index)
	// http.HandleFunc("/index", index)
	// log.Fatal(http.ListenAndServe(":8080", nil))
	tt, isExist := UserData["daniel"]
	fmt.Println(tt, isExist)
	// server := gin.Default()
	// server.LoadHTMLGlob("template/*")
	// server.GET("/", index)
	// server.Run(":8888")
}
