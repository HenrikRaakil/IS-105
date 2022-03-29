package Listener

import (
	"io"
	"net"
	"testing"
)

func TestListener(t *testing.T){
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		t.Fatal(err)
	}

	defer listener.Close()						// https://go.dev/tour/flowcontrol/12

	for {
		conn, err := listener.Accept()
		if err != nil{
			t.Fatal(err)						// Returnerer en error og avslutter programmet. Den hopper ut av løkken.
		}
		// Anonym funksjon, kan brukes som paramter/variabel
		go func() {								// ordet go blir lagret som en miniprosess inne i gomiljø.
			io.Copy(conn, conn)					// Kopier en strømm som kommer inn, og kopier en strømm som kommer ut (echo)
			conn.Close()						// Avslutter
		}()										// parametere til go func
	}

	t.Logf("Bundet til %q", listener.Addr())	// listener.Addr() Finner adressen og porten man bruker
}

/* Hører til package ch03

importerer Golang-pakkene "net" og "testing". Disse inneholder funksjoner som blir brukt i koden
parameterene (t *Testing.T) i funksjonen TestListener er satt opp etter standardaen til Testing pakken
til Golang som ble importert.

https://pkg.go.dev/net
https://pkg.go.dev/testing
https://pkg.go.dev/io
https://pkg.go.dev/bufio

Setter opp to variabeler som heter "listener" og "err". Setter verdiene til det samme som net.Listen.
net.Listen er en funksjon som hører til net-pakken som ble importert. paramteret til net.Listen spesifiserer hva slags type
adresse man hører etter, i dette tilfelle en TCP pakke, og hvilke IP adresse.

if statement. Hvis *ikke* err er det samme som nil (nil betyr ingen verdi) så kjør t.Fatal(err)
Dette er en funksjon som hører til testing pakken.

-

for-løkke uten noen parametere. Denne løkken vil gå i det uendelige helt til man dreper prossesen
*/