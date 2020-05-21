package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

// a package that will load the HTML in from the file and insert signatures into it for us: the html/template package//

func main() {
	fmt.Println("a pattern to follow: HTML Templates")

	// http.HandleFunc("/guestbook", viewHandler)
	// err := http.ListenAndServe("localhost:8090", nil)
	// log.Fatal(err)

	// executeTemplate()

	text := "Here is your template!!! and Action: {{.}} \n"
	executeTemplateWithAction(text, "ABC")

}

// func viewHandler(writer http.ResponseWriter, req *http.Request) {

// 	// resp := []byte("Welcome to guestbook")
// 	html, err := template.ParseFiles("./exercise/aPattern-view.html")
// 	checkError(err)

// 	err = html.Execute(writer, nil)
// 	checkError(err)

// }

func executeTemplate() {
	text := "Here is your template!!!"
	tmpl, err := template.New("test").Parse(text)
	checkError(err)
	err = tmpl.Execute(os.Stdout, nil)
	checkError(err)
}

func executeTemplateWithAction(text string, data interface{}) {
	// text := "Here is your template!!!"
	tmpl, err := template.New("test").Parse(text)
	checkError(err)
	err = tmpl.Execute(os.Stdout, data)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
