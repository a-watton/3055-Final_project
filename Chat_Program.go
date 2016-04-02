package main

import (
  "fmt"
  "github.com/gorilla/mux"

  "net/http"
)

var router = mux.NewRouter()



func helloPage(response http.ResponseWriter, request *http.Request){
  fmt.Fprintf(response, "Hello!")
}


const indexPage = `
<h1>Login</h1>
<form method="post" action="/login">
    <label for="name">User name</label>
    <input type="text" id="name" name="name">

    <button type="submit">Login</button>
</form>
`

func loginPage(response http.ResponseWriter, request *http.Request){
  fmt.Fprintf(response, indexPage)
  name := request.FormValue("Name: ")
  redirectTarget := "/"
  if name != "" {
    //setSession(name, response)
    redirectTarget = "/inside"
  }
}


func insidePage(response http.ResponseWriter, request *http.Request){
  fmt.Fprintf(response, "Inside")
}


func main(){
  router.HandleFunc("/", helloPage)
  router.HandleFunc("/login", loginPage)
    router.HandleFunc("/inside", insidePage)
  http.Handle("/", router)
  http.ListenAndServe(":8000", nil)
}
