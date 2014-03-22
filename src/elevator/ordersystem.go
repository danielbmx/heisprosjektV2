// Ordersystem
package elevator

import (
	//"fmt"
	"time"

) 

// Dummy variable for running only one elevator:
var takeOrder bool = true

type OrderSetLight struct {
    Floor int
    Dir Direction
	Light LightVal
}


/*
func GetCost(client ClientStatus, external ClientExternalOrder, command map[string]ClientCommandOrders) int {
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

*/



func OrderHandler(SetLightChan chan<- OrderSetLight, BtnPanelToOrderChan <-chan Button, OrderTakenChan <-chan OrderSetLight, OrderToFSMChan chan<- Button, LocalClientFSMToOrderChan <-chan LocalClient, ClientOrderToNetChan chan<- LocalClient, ClientNetToOrderChan <-chan LocalClient, BtnOrderToNetChan chan<- Button, BtnNetToOrderChan <-chan Button) {
   	/*
   	Provide neccesary order-handling based on information from driver via channels.
   	communication? 
	
	Send buttons to be turned on/off to "panel" via setLightChan
   	*/
	
	var btnFromPanel Button
	var setBtnToPanel OrderSetLight
	var btnFromNetwork Button 
	var orderToFSM Button
	var localClient LocalClient
	
	for{
		time.Sleep(25*time.Millisecond)
		//fmt.Println("OrderHandler: her")
		select{
			// Case 1: Button pressed on own panel
			case btnFromPanel = <- BtnPanelToOrderChan:
				//fmt.Println("btnFromPanel Button", btnFromPanel)
				
				// Send order to network module
				if btnFromPanel.Dir != NONE {
					BtnOrderToNetChan <- btnFromPanel
				}
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
				    

			
			// Case 2: Order received via network
			case btnFromNetwork = <- BtnNetToOrderChan:
				// Calculate cost related to taking order:
				//Cost1 := GetCost(ClientFromNet, btnFromNetwork)
				//Cost2 := GetCost(localClient, btnFromNetwork)
				//Cost3 
				// Compare cost with other elevators:
				
			
				// Tell panel to set light on
				setBtnToPanel.Floor = btnFromNetwork.Floor
				setBtnToPanel.Dir = btnFromNetwork.Dir
				setBtnToPanel.Light = ON
				SetLightChan <- setBtnToPanel
				
				// tell runElevator to go to that floor.
				
				orderToFSM.Floor = btnFromNetwork.Floor
				orderToFSM.Dir = btnFromNetwork.Dir
				OrderToFSMChan <- orderToFSM
			//default:
				//continue
		//

		//select 
			case localClient = <- LocalClientFSMToOrderChan: // From FSM
				//fmt.Println("OrderHandler: localClient: ", localClient)
				select{				
					case ClientOrderToNetChan <- localClient:
						//fmt.Println("OrderHandler: her")
					default:
						//fmt.Println("OrderHandler: to default")
						continue
				}
			//case ClientOrderToNetChan <- localClient:
				//fmt.Println("Sending localclient to net")

			case clientFromNet := <- ClientNetToOrderChan:
				_ = clientFromNet
				//fmt.Println("This is received from net:  ", clientFromNet)

			default:
				//fmt.Println("OrderHandler: default")
				continue
		
	
		}
	}

}




































