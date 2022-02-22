package square

import (
	"fmt"
	"math"
)

type intCustomType int //объявили кастомный тип типа int

const SidesTriangle intCustomType = 3 //объявили const кастомного типа intCustomType
const SidesSquare intCustomType = 4
const SidesCircle intCustomType = 0
const pi float64 = math.Pi

func CalcSquare(sideLen float64, sidesNum intCustomType) float64 {
	if sidesNum == 0 {
		return math.Pi * sideLen * sideLen
		
	} else if sidesNum == 3 {
		return sideLen * sideLen * math.Sqrt(3) / 4
		
	} else if sidesNum == 4 {
		return sideLen * sideLen
		
	} else {
		return 0
	}

}

func main() {
	
	fmt.Println(CalcSquare(10.0, SidesTriangle))
	fmt.Println(CalcSquare(10.0, SidesSquare))
	fmt.Println(CalcSquare(10.0, SidesCircle))
	fmt.Println(CalcSquare(10.0, 7))
}
