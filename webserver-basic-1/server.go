package main 

import (
	"fmt"
	"io/ioutil"
)

type Page struct { 
	Title string
	Body []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) *Page{
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return &Page{Title: title, Body: body}
}

func main(){
	p1 := &Page{Title: "simple web server", Body: []byte("This is a simple test webserver")}
	p1.save()
	p2 := loadPage("simple web server")
	fmt.Println(string(p2.Body))
}