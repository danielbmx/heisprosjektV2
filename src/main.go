package main

import (
		"./networkmodule"
        "./elevator"
        "fmt"
        "time"
        
)

func main(){
   
   	fmt.Println("main")


	// Driver channels:
	ButtonEventChan 	:= make(chan elevator.Button, 1)
	FloorEventChan 		:= make(chan int, 1)
	InitFloorChan       := make(chan int, 1)
   	
	// Channels between panel and ordersystem:
	SetLightChan 		:= make(chan elevator.OrderSetLight, 1)
	BtnPanelToOrderChan := make(chan elevator.Button, 1)

	// Channels between statemachine and ordersystem:
	OrderTakenChan      := make(chan elevator.OrderSetLight, 1)
	OrderToFSMChan		:= make(chan elevator.Button, 1)
	LocalClientFSMToOrderChan	:= make(chan elevator.LocalClient, 1)


	// Channels between net and ordersystem:
	ClientOrderToNetChan := make(chan elevator.LocalClient, 1)
	ClientNetToOrderChan 	:= make(chan elevator.LocalClient, 1)
	BtnOrderToNetChan      := make(chan elevator.Button, 1)
	BtnNetToOrderChan	:= make(chan elevator.Button, 1)


	// Channels between Net and UDP: 
	BtnFromUDPChan		:= make(chan elevator.Button, 1)
	ClientFromUDPChan	:= make(chan elevator.LocalClient, 1)

	elevator.Init(ButtonEventChan, FloorEventChan, InitFloorChan)

	go elevator.PanelHandler(ButtonEventChan, SetLightChan, BtnPanelToOrderChan)
	
	go elevator.OrderHandler(SetLightChan, BtnPanelToOrderChan, OrderTakenChan, OrderToFSMChan, LocalClientFSMToOrderChan, ClientOrderToNetChan , ClientNetToOrderChan, BtnOrderToNetChan, BtnNetToOrderChan)
	
	go networkmodule.NetworkHandler(BtnOrderToNetChan, BtnFromUDPChan, BtnNetToOrderChan, ClientOrderToNetChan, ClientFromUDPChan, ClientNetToOrderChan)
	

    go elevator.UpdateState(FloorEventChan, OrderToFSMChan, OrderTakenChan, LocalClientFSMToOrderChan)
    
	/*go func(){
		time.Sleep(10*time.Second)
		panic("\n")
	}()*/

	//fmt.Println(<-networkmodule.OrderChannel)
	
	
	//fmt.Println("Starting again")
	
   	//elevdriver.SetButtonLight(1, elevdriver.UP, elevdriver.ON)
   

    /*
    // Sender:
    
    testbutton := elevdriver.Button{1,1} 
    testalive := 10
    
    connection1 := networkmodule.UdpConnect("129.241.187.255:20005") 
    connection2 := networkmodule.UdpConnect("129.241.187.255:20006") 
    
    go networkmodule.UdpButtonSender(testbutton, connection1) 
    
    go networkmodule.UdpAliveSender(testalive, connection2)
    
    
    
    // Reciver:
    
    button_chan := make(chan elevdriver.Button) 
    alive_chan := make(chan int)

    knapp := elevdriver.Button{}
    alive := 0
    
    
    go networkmodule.UdpButtonReciver(button_chan)
    go networkmodule.UdpAliveReciver(alive_chan)
    
    for{
         select{
            case knapp = <- button_chan:
            	fmt.Println(knapp)
            	elevdriver.SetButtonLight(knapp.Floor, knapp.Dir, elevdriver.ON)
            	
            case alive = <- alive_chan:
            	fmt.Println(alive)
 
        }
    }
      
	
	fmt.Println(knapp.Floor)
	fmt.Println("før knapp")	
	fmt.Println("recieved")
		
	fmt.Println("ferdi")
	*/	
	for {
		time.Sleep(10000*time.Hour)
	}
}




















