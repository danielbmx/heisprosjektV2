// Panel

package elevator

import(
		
		"fmt"
		)

func PanelHandler(buttonEventChan chan Button, setLightChan chan SetLightFromOrder, btnPanelToOrder chan Button) {
	fmt.Println("PanelHandler running")
	var passOn Button
	var setLight SetLightFromOrder
	
	for{
		select{
			case passOn = <- buttonEventChan: 
				fmt.Println("Panel: button pressed", passOn)
				// Pass on buttons to ordersystem on the btnPanelToOrder channel
				btnPanelToOrder <- passOn
				//fmt.Println("Panel: button passed on")

			case setLight = <- setLightChan:
				// Set lights on panel
				SetButtonLight(setLight.Floor, setLight.Dir, setLight.Light)
		}

	}

}
