package main

import (
		"./networkmodule"
        "./elevator"
        "fmt"
        "time"
        
)

func main(){
   
   	fmt.Println("main")
   	
	// Create all neccesary channels:
	ButtonEventChan 	:= make(chan elevator.Button)
	FloorEventChan 		:= make(chan int)
	InitFloorChan       := make(chan int)
	
	OrderTakenChan      := make(chan elevator.OrderSetLight)

	SetLightChan 		:= make(chan elevator.OrderSetLight)
	BtnPanelToOrderChan := make(chan elevator.Button)
	

	//BtnFromNetworkChan    := make(chan elevator.Button)
	
	//OrderChan	 		:= make(chan [4][3]int)
	OrderToFSMChan		:= make(chan elevator.OrderToFSM)
	OrderToNetChan      := make(chan elevator.Button)

	BtnFromNetChan		:= make(chan elevator.Button)
	BtnNetToOrderChan	:= make(chan elevator.Button)
	
	//LocalClientChan		:= make(chan elevator.LocalClient)


	elevator.Init(ButtonEventChan, FloorEventChan, InitFloorChan)

	go elevator.PanelHandler(ButtonEventChan, SetLightChan, BtnPanelToOrderChan)
	
	go elevator.OrderHandler(BtnPanelToOrderChan, SetLightChan, OrderToFSMChan, OrderTakenChan, OrderToNetChan, BtnNetToOrderChan) //, LocalClientChan)
	
	go networkmodule.NetworkHandler(OrderToNetChan, BtnFromNetChan, BtnNetToOrderChan)

    go elevator.UpdateState(FloorEventChan, OrderToFSMChan, OrderTakenChan) //, LocalClientChan)
    
   

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




















