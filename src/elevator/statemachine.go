// State Machine 

package elevator

import (
		"time"
		//"fmt"		
)


// States
type State int
const (
    INVALID State = iota
    MOVING
    STANDSTILL
    //EMG_STOPPED
    N_STATES
)

// Events
type Event int
const(
	INITIALIZE = iota
	NOEVENT
	MOVE
	HALT
	//EMGSTOP
	N_EVENTS
)


// Private variables:

var timeStart time.Time
var currentState State = INVALID
var direction Direction
var event Event
//var last_floor = make(chan int, 1)

/*

// Statemachine:
func UpdateState() { 
	for {

		time.Sleep(10*time.Millisecond)

		event := GetNextEvent(currentState, DirEventChan, FloorEventChan, ordersystem.OrderChannel)
		//fmt.Println("State:", currentState, "  Event:", event)
		switch currentState {
		    
		    case INVALID:
		         switch event {
		            case INITIALIZE:
		                Init(ButtonEventChan, FloorEventChan, DirEventChan)
		                currentState = STANDSTILL
		                fmt.Println("UpdateState: INVALID-INITIALIZE")
		            case NOEVENT:
		                break 
		            case MOVE:
		                break 
		            case HALT:
		                break 
		                }
		    
		    case MOVING:
		        switch event {
		            case INITIALIZE:
		                break
		            case NOEVENT:
		                break 
		            case MOVE:
		                break 
		            case HALT:
		                
		                //Stop car
		                //Setmot
		                //Open door for 3 sec
		                //Delete this order from queue HAPPENS IN ORDERSYSTEM!
		                //current state = STANDSTILL
		                fmt.Println("UpdateState: In HALT")
		                ElevatorStop(UP)
		                SetDoorOpenLight(ON)
		                timeStart = time.Now()
		                currentState = STANDSTILL
		                }
		                
		    case STANDSTILL:
		        switch event {
		            
		            case INITIALIZE:
		                break
		            
		            case NOEVENT:
		                //Close door if door has been open for more than 3 sec
		                closeTime := time.Now()
		                if closeTime.Sub(timeStart) > 3*time.Second {
		                    SetDoorOpenLight(OFF)
		                    }
		                    
		            case MOVE:
		                //move car in right direction
		                ordersystem.GetNextDirection(DirEventChan, FloorEventChan, ordersystem.OrderChannel)
		                nextDir := <- DirEventChan
		                DirEventChan <- nextDir
		                SetMotorDir(nextDir)
		                currentState = MOVING
		            
		            case HALT:
		                //Open door
		                timeStart = time.Now()
		                SetDoorOpenLight(ON)
		       }		// Delete order
		}
		<-FloorEventChan
	}
}


// Get next event for statemachine
func GetNextEvent(currentState State, dirEventChan chan Direction, floorEventChan chan int, orderChan chan[4][3] int) Event { 
    //fmt.Println("Inside GetNextEvent")
    switch currentState {
    
        case INVALID:
        	return INITIALIZE
        
        case MOVING:
        	//fmt.Println("case MOVING")
        	// if should stop
        	//return HALT

        	// "Compare"
        	if ordersystem.StopAtFloor(dirEventChan, floorEventChan, orderChan){
        		return HALT
        	}

        	// if shoud continue moving
        	// return NOEVENT 
        	return NOEVENT   
        
        case STANDSTILL:

        	fmt.Println("GNE: Inside Standstill")
        	//if order in queue
        	nextDir := <- DirEventChan
		    DirEventChan <- nextDir
		    fmt.Println("GNE: nextDir= ", nextDir)
		    
        	if nextDir != NONE {//&& time.Now().Sub(timeStart) > 3 
				fmt.Println("GNE: getting dir")
        		return MOVE
        	
        	}
        	
        	// if no order in queue
        	// return NOEVENT


	}
	fmt.Println("GNE: return NOEVENT")
    return NOEVENT

}


*/




















