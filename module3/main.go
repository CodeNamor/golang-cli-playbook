package module3

import (
	"fmt"
	"github.com/nathan-osman/go-sunrise"
	"time"
)

func main() {
	fmt.Println("hello world")
	rise, set := sunrise.SunriseSunset(
		43.65, -79.38, // Toronto, CA
		2000, time.January, 1, // 2000-01-01
	)
	fmt.Printf("sunrise at: %v", rise)
	fmt.Printf("sunset at: %v\n", set)
}
