package state

/*
func Names(){
        var KY string = "Kylling"
        var RE string = "Rev"
        var KO string = "Korn"
        var HS string = "HomoSapiens"
}
*/

func ViewState() string {
    return "[V| kylling rev korn HS\\________________/ __________________ |Ø]"           // Sjekke data som er lagret ("kylling til venstre", "rev til venstre")
}

func PutInBoat() string {
    return "[V| _________________ \\kylling rev korn HS/ ________________ |Ø]"           //Flytter kylling rev korn i båten ->
}

func ChangeStateEast() string {
    return "[V| _________________ \\________________/ kylling rev korn HS |Ø]"           //Flytter kylling rev korn i båten ->
}