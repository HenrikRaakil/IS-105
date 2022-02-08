package state

import "testing"
import "fmt"

// Test funksjon
func TestViewState(t *testing.T) {
    wanted := "[V| kylling rev korn HS\\________________/ __________________ |Ø]"
	got := ViewState()
    if got == wanted {
         fmt.Printf("Feil %s, ønsket %s.", got, wanted)
    } else {
        t.Errorf("Fikk %s, ønsket %s.", got, wanted)
    }
}