// Ordersystem
package elevator

import (
	"fmt"
	//"net"
	"time"
	//"../networkmodule"
	
) 

// Dummy variable for running only one elevator:
var takeOrder bool = true

type SetLightFromOrder struct {
    Floor int
    Dir Direction
	Light LightVal
}

type OrderToFSM struct {
    Floor int
    Dir Direction
}

type Client struct {
	Floor int
	Dir Direction
	IpAddr string

}

var OrderChannel = make(chan [4][3]int, 10)

func InitOrderMatrix(orderchan chan [4][3]int) {
	var ordermatrix[4][3]int
	
	for i := 0; i<4; i++ {
		for j := 0; j<3; j++{
			ordermatrix[i][j] = 0;
		}
	}
	orderchan <- ordermatrix
}
/*
func SaveOrder(buttonEventChan chan Button, orderchan chan [4][3]int){
	for{
		button := <-buttonEventChan
		buttonEventChan <- button
		ordermatrix := <- orderchan
		ordermatrix[button.Floor - 1][button.Dir] = 1
		orderchan <- ordermatrix
		time.Sleep(50*time.Millisecond)
	}
}


Needs to be changed
func CalculateCost(buttonEventChan chan Button, floorEventChan chan int, floorDirectionChan chan Direction) int{
	dir := <- floorDirectionChan
	floorDirectionChan <- dir
	floor := <- floorEventChan
	floorEventChan <- floor
	button := <- buttonEventChan
	buttonEventChan <- button
	score := 0
	if dir != button.Dir {
		score += 1
	}
	if floor != button.Floor {
		score += 1
	}
	if button.Dir == DOWN {
		if floor >= button.Floor{
			score += (floor - button.Floor)
		}else{
			score += 4
		}
	}
	if button.Dir == UP {
		if floor <= button.Floor{
			score += (button.Floor - floor)
		}else{
			score += 4
		}
	}
	return score
}


func ResetOrder(elevator int, orderchan chan[4][3]int) {
	ordermatrix := <- orderchan
	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			if ordermatrix[i][j] == elevator{
				ordermatrix[i][j] = 1
			}
		}
	}
}

func DeleteOrder(button Button, orderchan chan[4][3]int){
   ordermatrix := <- orderchan
   ordermatrix[button.Floor][button.Dir] = 0
   orderchan <- ordermatrix
}



func StopAtFloor(dirEventChan chan Direction, floorEventChan chan int, orderChan chan [4][3]int) bool {
	//fmt.Println("inside StopAtFloor")
	dir := <- dirEventChan
	dirEventChan <- dir
	
	floor := <- floorEventChan
	floorEventChan <- floor
	//fmt.Println(floor)
	matrix := <- orderChan
	orderChan <- matrix
	if floor == -1{
		return false
	}
	if matrix[floor][dir] != 0 || matrix[floor][NONE] != 0 {
		fmt.Println("true returned")
		<-dirEventChan
		dirEventChan <- NONE
		return true
	}else{
		return false
	}
}


func OrderAbove(floor int, ordermatrix [4][3]int) bool {

	for floor+=1; floor < 4; floor++{
		if ordermatrix[floor][NONE] != 0 || ordermatrix[floor][UP] != 0 || ordermatrix[floor][DOWN] != 0{
			return true
			}
	}
	return false

}

func OrderBelow(floor int, ordermatrix [4][3]int) bool {
	
	for floor-=1; floor >= 0; floor--{
		if ordermatrix[floor][NONE]!=0 || ordermatrix[floor][UP]!=0 || ordermatrix[floor][DOWN]!=0{
			return true
			}
	}
	return false

}

func GetNextDirection(dirEventChan chan Direction, floorEventChan chan int, orderChan chan[4][3]int) {
	
	dir := <- dirEventChan
	dirEventChan <- dir
	
	matrix := <- orderChan
	orderChan <- matrix
	
	floor := <- floorEventChan
	floorEventChan <- floor
	
	switch dir {
		case NONE:
			if OrderBelow(floor, matrix) {
				<-dirEventChan
				dirEventChan <- DOWN
				
			}
			if OrderAbove(floor, matrix) {
				<-dirEventChan
				dirEventChan <- UP
				
			}
		case UP:
			if OrderAbove(floor, matrix) {
				<-dirEventChan
				dirEventChan <- UP
				
			}
			if OrderBelow(floor, matrix) {
				<-dirEventChan
				dirEventChan <- DOWN
				
			}
		case DOWN:
			if OrderBelow(floor, matrix) {
				<-dirEventChan
				dirEventChan <- DOWN
				
			}
			if OrderAbove(floor, matrix) {
				<-dirEventChan
				dirEventChan <- UP
				
			}

	}
	
}
*/
func OrderHandler(BtnPanelToOrderChan chan Button){ //,buttonEventChan chan Button, orderchan chan [4][3]int, )
   	/*
   	Provide neccesary order-handling based on information from driver via channels.
   	communication? 
	
	Send buttons to be turned on/off to "panel" via setLightChan
   	*/
	
	var btnFromPanel Button
	//var btnFromNetwork Button 

	for{
		time.Sleep(25*time.Millisecond)

		select{
			// Case 1: Button pressed on own panel
			case btnFromPanel = <- BtnPanelToOrderChan:
				fmt.Println("btnFromPanel Button", btnFromPanel)
				// Calculate cost related to taking order:

				// Compare cost with other elevators:
				
				// If should take order, tell statemachine and other participants
				// Else, someone else takes order.
				if takeOrder == true {
					// store order in ordermatrix
					fmt.Println("HandleOrder: true")
					// tell runElevator to go to that floor.

				}

			/*
			// Case 2: Order received via network
			case btnFromNetwork = <- BtnFromNetworkChan:
				// Calculate cost related to taking order:

				// Compare cost with other elevators:
			*/
			default:
				//fmt.Println("HandleOrder: default")
				continue
	/*
		for{
			fmt.Println("In handleOrder")
			time.Sleep(25*time.Millisecond)
			
			button := <- buttonEventChan
			buttonEventChan <- button
			
			// Set lights
			SetButtonLight(button.Floor, button.Dir, ON)
			
			// Push ordermatrix back to orderchan
			ordermatrix := <- orderchan
			orderchan <- ordermatrix
			ordermatrix[button.Floor - 1][button.Dir] = 1
			fmt.Println(ordermatrix)
			orderchan <- ordermatrix
			toemptybuttonChan := <- buttonEventChan
			
			fmt.Println("trykket:  ", toemptybuttonChan)
		}
			//UdpButtonSender(button, con_udp)
			//UdpButtonReciver(buttonEventChan)
	*/
		}
	}
}












