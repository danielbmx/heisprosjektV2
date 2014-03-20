// Ordersystem
package elevator

import (
	//"fmt"
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


/*
//Needs to be changed!
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

*/

func OrderHandler(BtnPanelToOrderChan <-chan Button, SetLightChan chan<- SetLightFromOrder, OrderToFSMChan chan<- OrderToFSM){ //,buttonEventChan chan Button, orderchan chan [4][3]int, )
   	/*
   	Provide neccesary order-handling based on information from driver via channels.
   	communication? 
	
	Send buttons to be turned on/off to "panel" via setLightChan
   	*/
	
	var btnFromPanel Button
	var setBtnToPanel SetLightFromOrder
	//var btnFromNetwork Button 
	var orderToFSM OrderToFSM

	for{
		time.Sleep(25*time.Millisecond)

		select{
			// Case 1: Button pressed on own panel
			case btnFromPanel = <- BtnPanelToOrderChan:
				//fmt.Println("btnFromPanel Button", btnFromPanel)
				// Calculate cost related to taking order:

				// Compare cost with other elevators:
				
				// If should take order, tell statemachine and other participants
				// Else, someone else takes order.
				
				if takeOrder == true {
					// store order in ordermatrix
					//fmt.Println("HandleOrder: true")
					
					// Tell panel to set light on
					setBtnToPanel.Floor = btnFromPanel.Floor
					setBtnToPanel.Dir = btnFromPanel.Dir
					setBtnToPanel.Light = ON
					SetLightChan <- setBtnToPanel
					
					// tell runElevator to go to that floor.
					
					orderToFSM.Floor = btnFromPanel.Floor
					orderToFSM.Dir = btnFromPanel.Dir
					OrderToFSMChan <- orderToFSM

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
	
		}
	}
}












