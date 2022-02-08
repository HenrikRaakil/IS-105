package state

import "testing"
import "fmt"

// Test funksjon
func TestViewState(t *testing.T) {
    wanted := "[V| kylling rev korn HS\\________________/ __________________ |Ø]"
	got := ViewState()
    if got == wanted {
         fmt.Printf("Suksess. Fikk %s, ønsket %s.", got, wanted)
    } else {
        t.Errorf("Feil. Fikk %s, ønsket %s.", got, wanted)
    }
}