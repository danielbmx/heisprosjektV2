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

var OrderMatrix = InitOrderMatrix()
var doorClose chan bool


// Statemachine:
func UpdateState(floorEventChan <-chan int, OrderToFSMChan <-chan OrderToFSM, OrderTakenChan chan<- OrderSetLight/*, LocalClientChan chan<- LocalClient*/) {  
	var event Event
	stateHasChanged := make(chan bool, 10)
	doorClose = make(chan bool, 2)
    var doorIsOpen bool

	
	for{
		time.Sleep(25*time.Millisecond)
		
		//fmt.Println("this is where it is hanging 2")
		//LocalClientChan <- localClient
		//fmt.Println("this is where it is hanging")
		
		// Read order(s) from ordersystem:
		fmt.Println("Selecting new event...")
		select {
		    case readOrder := <- OrderToFSMChan:
        		//fmt.Println("Event: FSM: order read: ", readOrder)
        		// Order saved in ordermatrix
		        OrderMatrix = SaveOrder(readOrder, OrderMatrix)
		        //fmt.Println("Event: FSM: ordermatrix: ", OrderMatrix)
                if OrderAbove(localClient.Floor, OrderMatrix) || OrderBelow(localClient.Floor, OrderMatrix){
                    if !doorIsOpen { 
                        event = MOVE
                    }
                }
		        
		    case newFloor := <- floorEventChan:
		        //fmt.Println("Event: Arrived at new floor: ", newFloor)
		        localClient.Floor = newFloor
		        if StopAtFloor(localClient.currentDir, localClient.Floor, OrderMatrix) {
		            
		            event = HALT
		            
		            // Delete order from ordermatrix and tell panel to turn off lights
		            
		            OrderMatrix = DeleteOrder(localClient.Floor, localClient.currentDir, OrderMatrix, OrderTakenChan)
		            fmt.Println(OrderMatrix)
		        }
		    case <- stateHasChanged:
		        //fmt.Println("Event: State has changed")
	        case <- doorClose:
	            //fmt.Println("Event: Door close")
	            SetDoorOpenLight(OFF)
	            doorIsOpen = false
	            newDir := GetNextDirection(localClient.currentDir, localClient.Floor, OrderMatrix)
	            if newDir != NONE {
	                event = MOVE
	            }
		    }

		//fmt.Println("State:", localClient.currentState, "  Event:", event)
		
		//fmt.Println("FSM: Floor: ", localClient.Floor, "  Dir: ", localClient.currentDir)
		
		switch localClient.currentState {
		
		    case MOVING:
		        switch event {
		            case MOVE:
		                break 
		            case HALT:
		                //Stop car
		                //Open door for 3 sec
		                //Delete this order from queue HAPPENS IN ORDERSYSTEM!

		                //fmt.Println("UpdateState: In HALT")
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
		                //fmt.Println("Update state: in state MOVE") 
		                //move car in right direction
		                newDir := GetNextDirection(localClient.currentDir, localClient.Floor, OrderMatrix)
		                
		                SetMotorDir(newDir)
		                localClient.currentState = MOVING
		                stateHasChanged <- true
		                //fmt.Println("current state Moving: ", localClient.currentState)
		                localClient.currentDir = newDir

		                break
		            
		            case HALT:
                		break
		       }
    		   break
		   }	    
	}
	
	

}

func timeAfter(ch chan bool, t time.Duration){
    time.Sleep(t)
    ch <- true
    fmt.Println("timer event has passed")
}



