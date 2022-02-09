package main

import (
    "fmt"
    "github.com/HenrikRaakil/is105-del3test/state"
)

func main(){
    fmt.Println("Dyrene ble venner og reven spiste ikke høna underturen")
    fmt.Println(state.ViewState())
    fmt.Println(state.PutInBoat())
    fmt.Println("På veien så delte de kornet")
    fmt.Println(state.ChangeStateEast())
}