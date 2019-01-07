package main
 
import (
   "fmt"
    "log"
    "net/http"
)
 
func sayHello(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Hello World ggggg!")
 
}
func main() {
    http.HandleFunc("/hello", sayHello)
    er := http.ListenAndServe(":9090", nil) 
    if er != nil {
        log.Fatal("ListenAndServe: ", er)
    }
}