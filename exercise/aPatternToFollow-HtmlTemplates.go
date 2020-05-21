package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

// a package that will load the HTML in from the file and insert signatures into it for us: the html/template package//

type GuestBook struct {
	SignatureCount int
	Signatures     []string
}

func main() {
	fmt.Println("a pattern to follow: HTML Templates")

	http.HandleFunc("/guestbook", viewHandler)
	http.HandleFunc("/guestbook/new", newHandler)
	http.HandleFunc("/guestbook/create", createHandler)
	err := http.ListenAndServe("localhost:8090", nil)
	log.Fatal(err)

	// executeTemplate()

	// text := "Here is your template!!! and Action: {{.}} \n"
	// executeTemplateWithAction(text, "ABC")

}
func createHandler(writer http.ResponseWriter, req *http.Request) {
	signature := req.FormValue("signature")
	// _, err := writer.Write([]byte(signature))
	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	file, err := os.OpenFile("exercise/data.txt", options, os.FileMode(0600))
	checkError(err)
	_, err = fmt.Fprintln(file, signature)
	checkError(err)
	err = file.Close()
	checkError(err)

	http.Redirect(writer, req, "/guestbook", http.StatusFound)
}

func newHandler(writer http.ResponseWriter, req *http.Request) {
	html, err := template.ParseFiles("./exercise/aPattern-new.html")
	checkError(err)
	err = html.Execute(writer, nil)
}

func viewHandler(writer http.ResponseWriter, req *http.Request) {
	signatures := getStrings("exercise/data.txt")
	fmt.Printf("%#v \n", signatures)
	// resp := []byte("Welcome to guestbook")
	html, err := template.ParseFiles("./exercise/aPattern-view.html")
	checkError(err)

	guestBook := GuestBook{
		SignatureCount: len(signatures),
		Signatures:     signatures,
	}
	err = html.Execute(writer, guestBook)
	checkError(err)

}

// func executeTemplate() {
// 	text := "Here is your template!!!"
// 	tmpl, err := template.New("test").Parse(text)
// 	checkError(err)
// 	err = tmpl.Execute(os.Stdout, nil)
// 	checkError(err)
// }

// func executeTemplateWithAction(text string, data interface{}) {
// 	// text := "Here is your template!!!"
// 	tmpl, err := template.New("test").Parse(text)
// 	checkError(err)
// 	err = tmpl.Execute(os.Stdout, data)
// 	checkError(err)
// }

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getStrings(fileName string) []string {
	var lines []string
	file, err := os.Open(fileName)
	if os.IsNotExist(err) {
		return nil
	}
	checkError(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	checkError(scanner.Err())
	return lines
}
