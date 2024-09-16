//I dont think my last server work and I have no idea what this server gone do
//now its a host server that talks to other servers/*Need to make this a client and host little self note*/

package main// I forgot go

import("fmt"
        "os"
        "net"
        "net/http"
       "bufio"
       "log"
       "context"
       "time"
      )

var WorldClock = time.ParseDuration("2m")//Work on time clock its for time out... reminder dude its food time

func main(){

    online()
    
}


func online(){// if its a func is it online(tinder(???))

    tinder()
    
}

const Tinder(){ // trying to relearn go
//This should make it a host and client server
    YelloPages := net.Dialer
    CWTO, cancel := context.WithoutTimeout(context.Background(), time.Minute)
    defer cancel()
    
    conn, err := YelloPages.DialContext(CWTO, "tcp", "operatingsystem.gov")
    if err != nil{// I tink this just the standard error code

        log.Fatalf("Connection Error Code: 404: could not dial /%v\n blessed day",err)

    }

    defer conn.Close()

    ear, err := net.Listener("tcp":"53")
    if err != nil{

        log.Fatal(err)
        
    }
    defer ear.Close()    
    for{

        conn,err := ear.Accept()
        if err != nil{

            log.Fatal(err)
            
        }
        
    }
    
    go func(plug net.conn){

        io.copy(plug,plug)
        c.Close()
        
    }(conn)
    
    var JoinHostPort("operatingsystem.gov","53")//still working on opening ports up
    
}

const popup(){//handle errors in here

    type Error interface{

        error
        Timeout() bool
        
    }
    
}

const locate(){//this gone be my addr func everything about an address I need to pass through YelloPages

    type Addr interface{

        network "tcp"//I think this is how I do it
        operatingsystem.gov string
        
        
    }

    type AddrError struct{

        Err string
        Addr string
        timeout() bool// Still working on the time out func
        
    }

    type conn interface( //Hopefully make connection actuall a thing on my localserver(LocalAddr)

        Read(a []byte) (n int, err error)
        write(a []byte) (n int, err error)

        Close() error
        LocalAddr() Addr
        remoteAddr() Addr

    )
    
    func (e *AddrError) Timeout() bool
    
}
