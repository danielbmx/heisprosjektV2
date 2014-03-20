// Execute order functions for FSM

package elevator

import(
    "fmt"
)



func InitOrderMatrix() [4][3]int {
	
	var ordermatrix[4][3]int
	
	for i := 0; i<4; i++ {
		for j := 0; j<3; j++{
			ordermatrix[i][j] = 0;
		}
	}
	return ordermatrix
}



func SaveOrder(readorder OrderToFSM, ordermatrix [4][3]int) [4][3]int{
	
	//time.Sleep(50*time.Millisecond)
	ordermatrix[readorder.Floor/* - 1*/][readorder.Dir] = 1
    return ordermatrix
}

// Her har Kristine begynt å rette opp ifra! Se over det!

func DeleteOrder(readorder OrderToFSM, ordermatrix [4][3]int) [4][3]int{
   ordermatrix[readorder.Floor/* - 1*/][readorder.Dir] = 0
   return ordermatrix
}



func StopAtFloor(dir Direction, floor int, ordermatrix [4][3]int) bool {
	
	if ordermatrix[floor/* - 1*/][dir] != 0 || ordermatrix[floor/* - 1*/][NONE] != 0 {
		fmt.Println("true returned")
		return true
	}else{
		return false
	}
}

// Returns true if there exists an order above
func OrderAbove(floor int, ordermatrix [4][3]int) bool {
	for floor+=1; floor < 4; floor++{
		if ordermatrix[floor][NONE] != 0 || ordermatrix[floor][UP] != 0 || ordermatrix[floor][DOWN] != 0{
			return true
			}
	}
	return false

}

// Returns true if there exists an order above
func OrderBelow(floor int, ordermatrix [4][3]int) bool {
	for floor-=1; floor >= 0; floor--{
		if ordermatrix[floor][NONE]!=0 || ordermatrix[floor][UP]!=0 || ordermatrix[floor][DOWN]!=0{
			return true
			}
	}
	return false

}

// This now returns a direction
func GetNextDirection(dir Direction, floor int, ordermatrix [4][3]int) Direction{
    fmt.Println("Getting dir, floor:", floor)
	switch dir {
		case NONE:
			if OrderBelow(floor, ordermatrix) {
			    fmt.Println("NONE -> DOWN")
				return DOWN
			}
			if OrderAbove(floor, ordermatrix) {
			    fmt.Println("NONE -> UP")
				return UP
			}
		case UP:
			if OrderAbove(floor, ordermatrix) {
			    fmt.Println("UP -> UP")
				return UP
			}
			if OrderBelow(floor, ordermatrix) {
			    fmt.Println("UP -> DOWN")
				return DOWN
			}
		case DOWN:
			if OrderBelow(floor, ordermatrix) {
			    fmt.Println("DOWN -> DOWN")
				return DOWN
			}
			if OrderAbove(floor, ordermatrix) {
			    fmt.Println("DOWN -> UP")
				return UP
			}
		}
    return NONE	
}

// Resets orders if one elevator is lost
func ResetOrder(elevator int, ordermatrix [4][3]int) [4][3]int {
	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			if ordermatrix[i][j] == elevator{
				ordermatrix[i][j] = 1
			}
		}
	}
	return ordermatrix
}




