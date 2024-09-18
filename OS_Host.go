// I need to make a server server... guess data base? idk
/*
I
Got
Hungry
*///Plus I gotta read on this port connection is stupid
package main

import("fmt"
      "net"
      )

func main(){

    ear, err := net.Listen("tcp":"9001")
    if err != nil{

        log.Fatal(err)
    
  }
    
    for{

            conn, err := ln.Accept()
            if err != nil{

                  Timeout()
                  
            }
            go handleConnection(conn)    
    }
    
}

func Timeout(){

    var stuckhand = time.Duration("2m")
    var stopwatch = time.After(2 * time.Minute)

    if stopwatch != stuckhand{
        fmt.Println("...")//if this works I'll take it out but I have no compilation time so let a note 
    }else stopwatch == stuckhand{
        fmt.Println("408")
    }
    
}
