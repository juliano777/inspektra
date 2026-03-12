package main

import (
        "fmt"
        "inspektra/sys"
)


func main(){
    params := sys.LoadParams()
    
    fmt.Println("Long live Rock in Roll!")
   
    fmt.Println("Host: ", params.Host)
    fmt.Println("Port: ", params.Port)
    fmt.Println("Database: ", params.Database)
    fmt.Println("Username: ", params.User)
    fmt.Println("")
}

