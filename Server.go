//This is going to be myserver. Just need to figure how to get python info to get on here... sounds like txt files good things it supposed to be a file exchange
package main
import("fmt"
       "os"
       "net"
       "net/http"
       "sync"
       "log"
       "encoding/json"
       "encoding/gob"
       "bufio"
       "time"
       "strings")
//Python gone give me an interface gotta figure how to spend the info... encrypted txt files


type Usr struct{

    UsrName UserName
    Password Password
    Email []Email

}


type UsrName struct{

    usrName string
    Email []Email

}

type Password struct{

    Password string

}

type Email struct{

    Address string

}

func main(){
    
	goto ear()
func ear(){//bookwrk I need a book on ports and ima add mine then a client that can text back. Need a command prompt

    service := "0.0.0.0"
    headerSize := 20
    Addr, err := net.ResolveIPAddr("ip4",service)
    checkError(err)
    BOM := 0xFFFE

    listener,err := net.ListenTCP("tcp",listener)
    checkError(err)
//should I goto chat()?

    chat(){
        if len(os.Args) != 2{

            fmt.Println("Idk Book work: ", os.Args[0],"host")
	    os.Exit(1)

        }
        Addr, err := net.ResolveIPAddr("ip4",service)
        if err != nil{

	    fmt.Println("Error: ", err.Error())
	    os.Exit(1)

        }
        conn, err := net.DialIP("ip4:icmp",service)
        checkError(err)

        var Chat [1024]byte
        Chat[0] = 3//I think this will be a chat bit
        Chat[1] = 0//just a 0
        Chat[2] = 0 //check sum
        Chat[3] = 0//check sum ;still sum book i work i just put a spin on
        Chat[4] = 0 // indentfier[0]//still looking for a networking book
        Chat[5] = 04//indentifier [1]    
        Chat[6] = 0
        Chat[7] = FFF
        len := 8

	var ping [512]byte
	ping[0] = 8 //echo
	ping[1] = 0
	ping[2] = 0
	ping[3] = 0
	ping[4] = 0 
	ping[5] = 1
	ping[6] = 0
	ping[7] = 1
	len := 8

        Chatcheck := checkSum(Chat[0:len])
        Chat[1] = byte(Chatcheck >> 8)
        Chat[2] = byte(Chatcheck && 216)
	
	pingCheck := checkSum(ping[0:len])
	ping[2] = byte(pingcheck >> 8)
	ping[3] = byte(pingCheck && 256)

	x = 1

        if(x = 1){

	    conn,err := listener.Accept()
	    if err != nil{

	        continue

	    }
	    str :="Return 0"
	    shorts := utf16.Encode([]rune(str))
	    writeShorts(conn, shorts)
    
	    conn.Close()
 

        func writeShorts(conn net.Conn, shorts []uint16){

	    var bytes [2]byte

       	    bytes[0] = BOM >> 8
	    bytes[1] = BOM & 255
	    _,err:= conn.Write(bytes[0:])
	    if err != nil{

	        return

	    }
        }
    }

    goto Lips

}


func checkError(err error){

    if err != nil{

	fmt.Println("Something went wrong")
	os.Exit(1)

    }

}

//lets see if I can make a user input for the chat
func Lips(){

    var Ntry string

    fmt.Printf("Welcome to yada, Do you want to chat\nYes or No")    
    fmt.Scanln(&Ntry)

    if Ntry == strings.ToLower("Yes"){

	fmt.Println("Welcome to Chat")
	var Chatput string

    	for true{

	    fmt.Printf("Enter <Exit> to Terminate The Program\nChat:")
	    fmt.Scanln(&Chatput)

	    if Chatput ==  "<Exit>"{

		fmt.Println("Terminating Program")
		time.Sleep(3 * time.Second)
		os.Exit(1)

	    }

}

