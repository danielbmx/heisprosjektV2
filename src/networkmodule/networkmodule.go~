// Network module
package main
import (
	"fmt"
	"net"
	"time"
	"encoding/json"
	"./elevdriver"
) 
/*
// Confirm elevator order taken to other elevators in the network
func UdpConfirmOrder() {
	UdpSender()
}

// Broadcast order recieved to all elevators in the network
func PassOrder() {

}
*/

// main function for testing JSON package sending:
func main() {

    ch := make(chan int)

    elevdriver.Init(nil, ch, nil, nil)    
	
	fmt.Println("main")
	testbutton := elevdriver.Button{
		Floor : 2,
		Dir : elevdriver.DOWN,
	}
	
	// message_chan := make(chan elevdriver.Button) 
	
	connection := UdpConnect()
	fmt.Println("connected")
	
	UdpSender(testbutton, connection)
	fmt.Println("sent")
	

	/*knapp := elevdriver.Button{}
	go UdpReciver(message_chan)
	fmt.Println("før knapp")
	knapp = <- message_chan
	fmt.Println("recieved")
	fmt.Println(knapp)
	
	elevdriver.SetButtonLight(knapp.Floor, knapp.Dir, elevdriver.OFF)
	fmt.Println("ferdi")*/
}

// Recieve message via UDP
func UdpReciver(message_channel chan elevdriver.Button) {
    
    serverAddr_udp, err := net.ResolveUDPAddr("udp", ":20005")
	PrintError(err)

    con_udp, err := net.ListenUDP("udp", serverAddr_udp)
    PrintError(err)
    save := elevdriver.Button{} 
    buffer := make([]byte,1024)
	//connection, err := net.ListenUDP("udp", UDP_addr)
	//PrintError(err)
	
	for {
	    fmt.Println("hit2") 
        n, _,_ := con_udp.ReadFromUDP(buffer)
        fmt.Println("hit") 
        //PrintError(err)
        
        err1 := json.Unmarshal(buffer[0:n],&save)
        PrintError(err1)
        message_channel <- save
    }
    
}

// Create UDP connection
func UdpConnect() *net.UDPConn{
	serverAddr_udp, err := net.ResolveUDPAddr("udp", "129.241.187.255:20005")
	PrintError(err)

    con_udp, err := net.DialUDP("udp", nil, serverAddr_udp)
    PrintError(err)
    
    return con_udp
}




// Broadcast message via UDP using Json
func UdpSender(parameter elevdriver.Button, con_udp *net.UDPConn) {
    fmt.Println("in udpSender")
    message, err := json.Marshal(parameter) 
    PrintError(err)
	
	for {
	fmt.Println("for in udpSender")
		time.Sleep(1000 * time.Millisecond)
		_, err2 := con_udp.Write(message)
		PrintError(err2)
	}
}


func PrintError(err error) {
	if err != nil{
        fmt.Println(err)
	}
}



























