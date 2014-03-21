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

type OrderSetLight struct {
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

func GetCost(client ClientStatus,external ClientExternalOrder,command map[string]ClientCommandOrders)int{
	currentPos:=client.Floor
	OrderedPos:=external.Floor
	penalty := abs(client.Floor-external.Floor)
	if(client.Direction == "STOP"){
		return penalty
	}
	for k,v:=range command{
		if(k==client.IP){
			if(v.HasCommandOrder && v.Floor==currentPos){
				penalty+=1
			}
		
		}
		currentPos+=1
		if(currentPos==OrderedPos){
			break
		}
	}
	return penalty
}


func OrderHandler(BtnPanelToOrderChan <-chan Button, SetLightChan chan<- OrderSetLight, OrderToFSMChan chan<- OrderToFSM, OrderTakenChan <-chan OrderSetLight, OrderToNetChan chan<- Button, BtnNetToOrderChan <-chan elevator.Button){ 

   	/*
   	Provide neccesary order-handling based on information from driver via channels.
   	communication? 
	
	Send buttons to be turned on/off to "panel" via setLightChan
   	*/
	
	var btnFromPanel Button
	var setBtnToPanel OrderSetLight
	//var btnFromNetwork Button 
	var orderToFSM OrderToFSM

	for{
		time.Sleep(25*time.Millisecond)

		select{
			// Case 1: Button pressed on own panel
			case btnFromPanel = <- BtnPanelToOrderChan:
				//fmt.Println("btnFromPanel Button", btnFromPanel)
				
				// Send order to network module
				OrderToNetChan <- btnFromPanel
				
				
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
				
			case DeleteOrder := <- OrderTakenChan:
			    setBtnToPanel.Floor = DeleteOrder.Floor
				setBtnToPanel.Dir = DeleteOrder.Dir
				setBtnToPanel.Light = DeleteOrder.Light
				SetLightChan <- setBtnToPanel
				    

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












