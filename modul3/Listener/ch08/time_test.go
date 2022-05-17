package ch08

import (
	"net/http"
	"testing"
	"time"
)

func TestHeadTime(t *testing.T) {
	resp, err := http.Head("https://www.time.gov")
	if err != nil {
		t.Fatal(err)
	}
	_ = resp.Body.Close()		// Avslutter transaksjonen / forbindelsen

	now := time.Now().Round(time.Second)	// bruker Time pakken
	date := resp.Header.Get("Date")
	if date == "" {
		t.Fatal("Ingen dato header mottatt fra time.gov")
	}

	dt, err := time.Parse(time.RFC1123, date)		// RFC: Request For Comment, forskjellige dataformater som definerer tid
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("time.gov: %s (avvik %s)", dt, now.Sub(dt))	// Subtract. tiden vi har - en annen tid
}
