package main
import (
	"fmt"
	"net/http"
	"html/template"
)

func indexHandler(w http.ResponseWriter,r *http.Request) {
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "index", nil)
}
func main() {
	fmt.Println("Listen port 3000")
	http.HandleFunc("/",indexHandler)
	http.ListenAndServe(":3000",nil)
}


