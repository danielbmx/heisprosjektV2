// State Machine 

package elevator

import (
		"time"
		"fmt"		
)


// States
type State int
const (
    MOVING State = iota
    STANDSTILL
    //EMG_STOPPED
    N_STATES
)

// Events
type Event int
const(
	MOVE = iota
	HALT
	//EMGSTOP
	N_EVENTS
)

type LocalClient struct {
    currentState State
    currentDir Direction
	Floor int
	//orderedBtn Button
	//IpAddr string

}

// Private variables:

var timeStart time.Time

var event Event

var localClient = LocalClient{STANDSTILL,NONE,-1}//,{0, NONE} ," "}
//localClient.currentState = STANDSTILL
//localClient.currentDir   = NONE

//var readOrder OrderToFSM
//var last_floor = make(chan int, 1)

var OrderMatrix = InitOrderMatrix()
var doorClose chan bool

func timeAfter(ch chan bool, t time.Duration){
    time.Sleep(t)
    ch <- true
    fmt.Println("timer event has passed")
}

// Statemachine:
func UpdateState(floorEventChan <-chan int, OrderToFSMChan <-chan OrderToFSM, OrderTakenChan chan<- OrderSetLight) {  // , LocalClientChan chan<- LocalClient
		var event Event
		stateHasChanged := make(chan bool, 10)
		doorClose = make(chan bool, 2)
	    var doorIsOpen bool
		//var DeleteLight OrderSetLight   
		

    /*go func(){
        for {
            time.Sleep(10*time.Millisecond)
            localClient.Floor = <- floorEventChan 
            fmt.Println("thread, Floor: ",localClient.Floor)
        }
    }()*/

	for {


		
		// Read order(s) from ordersystem:
		fmt.Println("Selecting new event...")
		select {
		    case readOrder := <- OrderToFSMChan:
        		fmt.Println("Event: FSM: order read: ", readOrder)
        		// Order saved in ordermatrix
		        OrderMatrix = SaveOrder(readOrder, OrderMatrix)
		        fmt.Println("Event: FSM: ordermatrix: ", OrderMatrix)
                if OrderAbove(localClient.Floor, OrderMatrix) || OrderBelow(localClient.Floor, OrderMatrix){
                    if !doorIsOpen { 
                        event = MOVE
                    }
                }
		        
		    case newFloor := <- floorEventChan:
		        fmt.Println("Event: Arrived at new floor: ", newFloor)
		        localClient.Floor = newFloor
		        if StopAtFloor(localClient.currentDir, localClient.Floor, OrderMatrix) {
		            
		            event = HALT
		            

		            /*
		            // Turn off light(s)
		            DeleteLight.Floor = localClient.Floor
		            DeleteLight.Dir = localClient.currentDir
		            DeleteLight.Light = OFF
		            OrderTakenChan <- DeleteLight
		            */
		            // Delete order from ordermatrix
		            fmt.Println("Sletter ordre")
		            OrderMatrix = DeleteOrder(localClient.Floor, localClient.currentDir, OrderMatrix, OrderTakenChan)
		            fmt.Println(OrderMatrix)
		        }
    //		case event = <- GetNextEvent(localClient.currentState, localClient.currentDir, localClient.Floor, OrderMatrix):
		    case <- stateHasChanged:
		        fmt.Println("Event: State has changed")
	        case <- doorClose:
	            fmt.Println("Event: Door close")
	            SetDoorOpenLight(OFF)
	            doorIsOpen = false
	            newDir := GetNextDirection(localClient.currentDir, localClient.Floor, OrderMatrix)
	            if newDir != NONE {
	                event = MOVE
	            }
//		case <- time.After(10*time.Millisecond):
		    }
//		fmt.Println("FSM: order read: ", readOrder)
		

//		OrderMatrix = SaveOrder(readOrder, OrderMatrix)
//		fmt.Println("FSM: ordermatrix: ", OrderMatrix)

 
//		event := GetNextEvent(localClient.currentState, localClient.currentDir, localClient.Floor, OrderMatrix)
		fmt.Println("State:", localClient.currentState, "  Event:", event)
		
		fmt.Println("FSM: Floor: ", localClient.Floor, "  Dir: ", localClient.currentDir)
		switch localClient.currentState {
		
		    case MOVING:
		        switch event {
		            case MOVE:
		                break 
		            case HALT:
		                //Stop car
		                //Open door for 3 sec
		                //Delete this order from queue HAPPENS IN ORDERSYSTEM!

		                fmt.Println("UpdateState: In HALT")
		                ElevatorStop(localClient.currentDir)
		                SetDoorOpenLight(ON)
		                doorIsOpen = true
                		go timeAfter(doorClose, 3*time.Second)
		                //timeStart = time.Now()

		                localClient.currentState = STANDSTILL
		                stateHasChanged <- true
		                localClient.currentDir = NONE
                        break
                }
                break      
                       
		    case STANDSTILL:
		        switch event {		            

		            case MOVE:
		                fmt.Println("Update state: in state MOVE") 
		                //move car in right direction
		                newDir := GetNextDirection(localClient.currentDir, localClient.Floor, OrderMatrix)
		                fmt.Println(newDir)
		                SetMotorDir(newDir)
		                localClient.currentState = MOVING
		                stateHasChanged <- true
		                fmt.Println("current state Moving: ", localClient.currentState)
		                localClient.currentDir = newDir
		                break
		            
		            case HALT:
		                /*
		                fmt.Println("in halt")
		                if OrderMatrix[localClient.Floor][UP] != 0 || OrderMatrix[localClient.Floor][DOWN] != 0{
		                    SetDoorOpenLight(ON)
		                    go timeAfter(doorClose, 3*time.Second)
		                }*/
		                //Open door
		                //timeStart = time.Now()
		                //SetDoorOpenLight(ON)
		                //fmt.Println("Starter doortimer :-)")
                		//go timeAfter(doorClose, 3*time.Second)
                		break
		       }		// Delete order
    		   break
		   }	    
	}
}

/*
// Get next event for statemachine
func GetNextEvent(currentState State, dir Direction, floor int, ordermatrix [4][3] int)  Event { 
    //fmt.Println("Inside GetNextEvent")
    
    switch currentState {
    
        case MOVING:
            fmt.Println("GNE: in event moving")
        	if StopAtFloor(dir, floor, ordermatrix){
        	    fmt.Println("GNE: getting stop at floor")
                return HALT
        	}else{
        	    return NOEVENT
        	}  
        
        case STANDSTILL:
            if OrderAbove(floor, ordermatrix) || OrderBelow(floor, ordermatrix){
                return MOVE
            }else{
                return NOEVENT
            }        	
        	
	}
	fmt.Println("GNE: default. return NOEVENT")
    return NOEVENT


}



*/


















