package state

import "github.com/HenrikRaakil/is105-del3test/state"
import "testing"
import "fmt"

// Test funksjon
func TestViewState(t *testing.T) {
    wanted := "[V| kylling rev korn HS\\________________/ __________________ |Ø]"
    if ViewState() != wanted {
         fmt.Println("Funksjon ViewState er en suksess! De er identiske!")
    }
    else {
        t.Error("Feil, fikk %q, ønsket %q.", ViewState, wanted)
    }
}

func TestPutInBoat(t *testing.T) {
    wanted := "[V| kylling rev korn HS\\________________/ __________________ |Ø]"
    if PutInBoat() == wanted {
         fmt.Println("Funksjon ViewState er en suksess! De er identiske!")
    }
    else {
        t.Error("Feil, fikk %q, ønsket %q.", ViewState, wanted)
    }
}