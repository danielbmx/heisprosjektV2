package main

import (
		//"./networkmodule"
        "./elevator"
        "fmt"
        "time"
)

func main(){
   
   	fmt.Println("main")
   	
	// Create all neccesary channels:
	ButtonEventChan 	:= make(chan elevator.Button)
	FloorEventChan 		:= make(chan int)
	DirEventChan 		:= make(chan elevator.Direction)

	SetLightChan 		:= make(chan elevator.SetLightFromOrder)
	BtnPanelToOrderChan := make(chan elevator.Button)
	//BtnFromNetworkChan	:= make(chan elevator.Button)
	
	//OrderChan	 		:= make(chan [4][3]int)
	//OrderToFSMChan		:= make(chan elevator.OrderToFSM)



	//elevator.InitOrderMatrix(networkmodule.OrderChan)
   	
	//go elevator.UpdateState()
	
	//go ordersystem.OrderHandler(elevator.ButtonEventChan, networkmodule.OrderChan)

	elevator.Init(ButtonEventChan, FloorEventChan, DirEventChan)

	go elevator.PanelHandler(ButtonEventChan, SetLightChan, BtnPanelToOrderChan)
	
	go elevator.OrderHandler(BtnPanelToOrderChan)





	//fmt.Println(<-networkmodule.OrderChannel)
	
	
	//fmt.Println("Starting again")
	
   	//elevdriver.SetButtonLight(1, elevdriver.UP, elevdriver.ON)
   	/*

    go elevdriver.Poller(buttonEventChan, floorEventChan)
    
    
    go HandleOrder()
   
    */
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






/*
func foo(get, set chan []int){

	var arr []int

	for {
		select {
			case get <- arr:
			
			case arr := <- set:
		}
	}
}


setter := make(chan []int)
getter := make(chan []int)

go foo(getter, setter)


func bar(getter, setter chan []int){

	data := <- getter
	data := <- getter
	// data
	
	setter <- data
}


func bar(getter chan []int){

	data := <- getter
	data := <- getter
	// data
	
}



func getButtonEventChanCopy(ButtonEventChan chan elevdriver.Button){
	



}




*/



















