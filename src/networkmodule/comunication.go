// Network module
package networkmodule

import (
	"fmt"
	"net"
	"time"
	"encoding/json"
	elevator "../elevator" 
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



func NetworkHandler(BtnOrderToNetChan <-chan elevator.Button, BtnFromUDPChan chan elevator.Button, BtnNetToOrderChan chan<- elevator.Button, ClientOrderToNetChan <-chan elevator.LocalClient, ClientFromUDPChan chan elevator.LocalClient, ClientNetToOrderChan chan<- elevator.LocalClient  ) { // 6

	// Storing Local IP	
	var client elevator.LocalClient
	client.IpAddr, _ = elevator.LocalIP()
	

	// map with all peers joined:
	AllPeers := make(map[string]elevator.LocalClient)

    // Make connections
    BtnPort := "129.241.187.255:20005"
    //AlivePort := "129.241.187.255:20006"
	ClientPort := "129.241.187.255:20007"
    BtnConnection := UdpConnect(BtnPort)
    //AliveConnection := UdpConnect(AlivePort)
	ClientConnection := UdpConnect(ClientPort)
    
    go UdpButtonReciver(BtnFromUDPChan)
    go UdpClientReciver(ClientFromUDPChan)
    
    for{
        select{
            case btnFromOrder := <- BtnOrderToNetChan:
                //fmt.Println("btnFromOrder: ", btnFromOrder)
				//fmt.Println("NWH: case 1 runs")
                // send via UDP to other elevators using the same port
                UdpButtonSender(btnFromOrder, BtnConnection)

				// Calculate cost related to taking the order:
				                



            case btnFromNet := <- BtnFromUDPChan:
				//fmt.Println("NWH: case 2 runs")

				// Calculate cost related to taking the order:
				best_peer := OrderDistribute(AllPeers, btnFromNet)				
							
				if best_peer.IpAddr.String() == client.IpAddr.String(){
					// Send to Ordersystem
					BtnNetToOrderChan <- btnFromNet
				}
				
				


			case clientFromOrder := <- ClientOrderToNetChan:
//				fmt.Println("NWH: case 3 runs")
//				fmt.Println("!!!! Sending Client via UDP!!!!!")
				UdpClientSender(clientFromOrder, ClientConnection)
			
			case clientFromNet := <- ClientFromUDPChan:
//				fmt.Println("NWH: case 4 runs")
//				fmt.Println("!!!! Receiving Client via UDP!!!!!")
				AllPeers[clientFromNet.IpAddr.String()] = clientFromNet
				fmt.Println(AllPeers)
				ClientNetToOrderChan <- clientFromNet
        }
    }
}


// Create UDP connection
func UdpConnect(address string) *net.UDPConn{
	serverAddr_udp, err := net.ResolveUDPAddr("udp", address)
	PrintError(err)

    con_udp, err := net.DialUDP("udp", nil, serverAddr_udp)
    PrintError(err)
    
    return con_udp
}

// Broadcast message via UDP using Json
func UdpButtonSender(parameter elevator.Button, con_udp *net.UDPConn) {

    message, err := json.Marshal(parameter) 
    PrintError(err)
	
//	for {
	
//		time.Sleep(1000 * time.Millisecond)
		_, err2 := con_udp.Write(message)
		PrintError(err2)
//	}
}

// Recieve message via UDP
func UdpButtonReciver(message_channel chan elevator.Button) {
    
    serverAddr_udp, err := net.ResolveUDPAddr("udp", ":20005")
	PrintError(err)

    con_udp, err := net.ListenUDP("udp", serverAddr_udp)
    PrintError(err)
    save := elevator.Button{} 
    buffer := make([]byte,1024)
	//connection, err := net.ListenUDP("udp", UDP_addr)
	//PrintError(err)
	
	for {
        n, _,err := con_udp.ReadFromUDP(buffer)
        PrintError(err)
        
        err1 := json.Unmarshal(buffer[0:n],&save)
        PrintError(err1)
        message_channel <- save
    }
    
}


func UdpClientSender(parameter elevator.LocalClient, con_udp *net.UDPConn) {

    message, err := json.Marshal(parameter) 
    PrintError(err)
	
	for i := 0; i<2; i++ {
		time.Sleep(10 * time.Millisecond)
		_, err2 := con_udp.Write(message)
		PrintError(err2)
	}
}

func UdpClientReciver(message_channel chan elevator.LocalClient) {
    
    serverAddr_udp, err := net.ResolveUDPAddr("udp", ":20007")
	PrintError(err)

    con_udp, err := net.ListenUDP("udp", serverAddr_udp)
    PrintError(err)
    save := elevator.LocalClient{} 
    buffer := make([]byte,1024)
	//connection, err := net.ListenUDP("udp", UDP_addr)
	//PrintError(err)
	
	for {
        n, _,err := con_udp.ReadFromUDP(buffer)
        PrintError(err)
        
        err1 := json.Unmarshal(buffer[0:n],&save)
        PrintError(err1)
        message_channel <- save
    }
    
}


func UdpAliveSender(parameter int, con_udp *net.UDPConn) {
    message, err := json.Marshal(parameter) 
    PrintError(err)
	
	for {
	fmt.Println("for in udpSender")
		time.Sleep(1000 * time.Millisecond)
		_, err2 := con_udp.Write(message)
		PrintError(err2)
	}
}

func UdpAliveReciver(message_alive chan int) {
    
    serverAddr_udp, err := net.ResolveUDPAddr("udp", ":20006")
	PrintError(err)

    con_udp, err := net.ListenUDP("udp", serverAddr_udp)
    PrintError(err)
    save := 0
    buffer := make([]byte,1024)
	//connection, err := net.ListenUDP("udp", UDP_addr)
	//PrintError(err)
	
	for {
        n, _,_ := con_udp.ReadFromUDP(buffer)
        //PrintError(err)
        
        err1 := json.Unmarshal(buffer[0:n],&save)
        PrintError(err1)
        message_alive <- save
    }
}



func PrintError(err error) {
	if err != nil{
        fmt.Println(err)
	}
}



func GetCost(client elevator.LocalClient, btn elevator.Button) int {
	var temp int
	if client.CurrentDir == elevator.NONE{
		temp = client.Floor-btn.Floor
		if temp < 0 {
			temp = temp * -1
		}	
		temp += 1
	}
	
	if client.CurrentDir == elevator.UP && btn.Dir == elevator.UP && btn.Floor >= client.Floor {
		temp = btn.Floor - client.Floor
	}
	if client.CurrentDir == elevator.DOWN && btn.Dir == elevator.DOWN && btn.Floor <= client.Floor {
		temp = client.Floor - btn.Floor
	}
	if client.CurrentDir != btn.Dir {
		temp := client.Floor-btn.Floor
		if temp < 0 {
			temp = temp * -1
		}	
		temp += 3
	}else{
		temp = 6
	}
	
	penalty := temp
	
	return penalty
	
}


func GetCost2(client elevator.LocalClient, btn elevator.Button) int {

	temp := client.Floor-btn.Floor
	if temp < 0 {
		temp = temp * -1
	}	

	penalty := temp

	return penalty

}


func OrderDistribute(peers map[string]elevator.LocalClient, btn elevator.Button) elevator.LocalClient {
	least_cost := 10
	var best_peer elevator.LocalClient
	
	for k,_ := range peers {
		best_peer = peers[k]
		break

	}
	
	for key,_ := range peers {
		temp_cost := GetCost(peers[key],btn)
		if temp_cost < least_cost {
			least_cost = temp_cost
			best_peer = peers[key]		

		}
	}
	return best_peer
}


















































