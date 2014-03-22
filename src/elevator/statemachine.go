// State Machine 

package elevator

import (
		"time"
		"fmt"
		"net"		
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
    CurrentDir Direction
	Floor int
	IpAddr net.IP

}

// Private variables:

var timeStart time.Time

var event Event

var localClient LocalClient //{STANDSTILL,NONE,-1}//,{0, NONE} ," "}

var OrderMatrix = InitOrderMatrix()
var doorClose chan bool


// Statemachine:
func UpdateState(floorEventChan <-chan int, OrderToFSMChan <-chan Button, OrderTakenChan chan<- OrderSetLight, LocalClientChan chan<- LocalClient) {  
	var event Event
	//stateHasChanged := make(chan bool, 10)
	doorClose = make(chan bool, 2)
    var doorIsOpen bool
	var prevDir Direction

	// initialize localClient
	localClient.currentState = STANDSTILL
	localClient.CurrentDir = NONE
	localClient.Floor = -1
	localClient.IpAddr,_ = LocalIP()
	
	for{
		time.Sleep(25*time.Millisecond)
		
		
		//Read order(s) from ordersystem:
		select {
			case LocalClientChan <- localClient:
				break
		    case readOrder := <- OrderToFSMChan:
        		//Save order in ordermatrix
		        OrderMatrix = SaveOrder(readOrder, OrderMatrix)
	            fmt.Println("newOrder:", OrderMatrix)
		        //Check for other orders
                if OrderAbove(localClient.Floor, OrderMatrix) || OrderBelow(localClient.Floor, OrderMatrix){
                    if !doorIsOpen { 
                        event = MOVE
                    }
                }
				/*
				if readOrder.Floor == localClient.Floor {
	                SetDoorOpenLight(ON)
					doorIsOpen = true
					go timeAfter(doorClose, 3*time.Second)
		            OrderMatrix = DeleteOrder(localClient.Floor, localClient.CurrentDir, OrderMatrix, OrderTakenChan)
				}
				*/
		        /*if StopAtFloor(localClient.CurrentDir, localClient.Floor, OrderMatrix) {
		            event = HALT		            
		            // Delete order from ordermatrix and tell panel to turn off lights
		            OrderMatrix = DeleteOrder(localClient.Floor, localClient.CurrentDir, OrderMatrix, OrderTakenChan)
		            fmt.Println("Order was at this floor:", OrderMatrix)
		        }*/
		        
		    case newFloor := <- floorEventChan:
				// If floor is updated, check if stop is needed
		        localClient.Floor = newFloor
				SetFloorLight(newFloor)
		        if StopAtFloor(localClient.CurrentDir, localClient.Floor, OrderMatrix) {
		            
		            event = HALT
		            
		            // Delete order from ordermatrix and tell panel to turn off lights
		            OrderMatrix = DeleteOrder(localClient.Floor, localClient.CurrentDir, OrderMatrix, OrderTakenChan)
		            fmt.Println(OrderMatrix)
		        }
		    //case <- stateHasChanged:
		        //fmt.Println("Event: State has changed")
	        case <- doorClose:
				//If timer is out -> close door and get next direction
	            SetDoorOpenLight(OFF)
	            doorIsOpen = false
	            newDir := GetNextDirection(localClient.CurrentDir, prevDir, localClient.Floor, OrderMatrix)
	            if newDir != NONE {
	                event = MOVE
	            }
		    }

		switch localClient.currentState {
		
		    case MOVING:
		        switch event {
		            case MOVE:
		                break 
		            case HALT:
		                //Stop elevator and open door for 3 sec
		                ElevatorStop(localClient.CurrentDir)
		                SetDoorOpenLight(ON)
		                doorIsOpen = true
                		go timeAfter(doorClose, 3*time.Second)
		          		//Update client
		                localClient.currentState = STANDSTILL
		                //stateHasChanged <- true
						prevDir = localClient.CurrentDir
		                localClient.CurrentDir = NONE
                        break
                }
                break      
                       
		    case STANDSTILL:
		        switch event {		            

		            case MOVE:
		                //Figure out new direction and move
		                newDir := GetNextDirection(localClient.CurrentDir, prevDir, localClient.Floor, OrderMatrix)
		                SetMotorDir(newDir)
						//Update client
		                localClient.currentState = MOVING
		                //stateHasChanged <- true
		                localClient.CurrentDir = newDir

		                break
		            
		            case HALT:
						/*if (OrderMatrix[localClient.Floor][UP] != 0 || OrderMatrix[localClient.Floor][DOWN] != 0) {
							SetDoorOpenLight(ON)
		                	doorIsOpen = true
                			go timeAfter(doorClose, 3*time.Second)
							//OrderMatrix = DeleteOrder(localClient.Floor, UP, OrderMatrix, OrderTakenChan)
							//OrderMatrix = DeleteOrder(localClient.Floor, DOWN, OrderMatrix, OrderTakenChan)

						}
						
						if StopAtFloor(localClient.CurrentDir, localClient.Floor, OrderMatrix) {
		            
		            		event = HALT
		      				OrderMatrix = DeleteOrder(localClient.Floor, localClient.CurrentDir, OrderMatrix, OrderTakenChan)    
		        		}*/
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



