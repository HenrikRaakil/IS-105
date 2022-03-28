package Listener

import (
	"crypto/rand"
	"io"
	"net"
	"testing"
)

func TestReadIntoBuffer(t *testing.T) {
	// Generer mock data
	payload := make([]byte, 1<<24) // "make" alokerer plass i minne. << skift operatør: 0001 1<<1 0010 Forskyver en desimal
	// I dette tilfelle alokkerer 2^24 = 16777216 bytes
	_, err := rand.Read(payload) // leser inn tilfeldige byte i payload
	if err != nil {
		t.Fatal(err)
	}

	listener, err := net.Listen("tcp","127.0.0.1:")
	if err != nil {
		t.Fatal(err)
	}

	go func() {
		conn, err := listener.Accept()
		if err != nil {
			t.Log(err)
			return
		}
		defer conn.Close()

		_, err = conn.Write(payload) // den første delen brukes til å sende x antall bytes. _ betyr at man ignorerer den
		if err != nil {
			t.Fatal(err)
		}
	}()

	conn, err := net.Dial("tcp", listener.Addr().String())

	buf := make([]byte, 1<<19) // 512 KiB
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err != io.EOF {
				t.Error(err)
			}
			break

		}
		t.Logf("read %d bytes", n)
	}
	conn.Close()
}










