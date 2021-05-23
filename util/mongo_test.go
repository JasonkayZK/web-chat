package util

import (
	"fmt"
	"strconv"
	"testing"
)

func Test_InitMongo(t *testing.T) {
	InitMongo()
}

func TestTemp(t *testing.T) {
	fmt.Println(strconv.ParseInt("9223372036854775807", 10, 64))

}