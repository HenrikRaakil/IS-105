# Del 1
Installasjon av verktøy og testing av funksjonaliteten (se ressursen V&GB: Kommandolinje ). Alle punktene skal demonstreres med skjermopptak inkludert konsist forklaring om det som demonstreres. 

De som bruker Microsoft Windows 10 operativsystemet installerer WSL/Ubuntu ( enkelte kan oppleve at de må aktivere virtualisering i BIOS; for mer informasjon se (craigloewen-msft, 2021a og 2021b) ). For å demonstrere bruken av kommandoer fra kommandolinje må MS Win 10 brukere installere Ubuntu app, starte den og installere Go i det virtuelle miljøet og vise at den funksjonerer med kommandoen $ go version 

1) Golang og Java 

a) bevis hvilke versjoner som er installert med kommandoer i CLI-applikasjonen ( for Golang: $ go version og Java: $ java --version )

b) bevis at man kan kompilere et Golang-program (lagret i en fil;bruk kode fra slide 6 (IS-105, 2022a)) og utføre den med kommandoer i CLI-applikasjonen ( $ go build <din_fil>.go og for å utføre $ ./<din_fil> )

c) bevis at man kan kompilere et Java-program til "bytecode" og fortolke den med kommanoder i CLI-applikasjonen ( kompilering $ javac <din_fil>.java og fortolking $ java <din_fil> )

d) drøft forskjellene mellom prosessen fra å skrive kildekode til å kunne utføre program på en datamaskin for Golang og Java programmeringspråk

2) Installere Wireshark (starte Wireshark og fange pakker på en lokal grensesnitt mellom en golang klient og server ( eksempelkode gitt i https://github.com/digitnow/demos-wireshark (Lenker til en ekstern side.) )

Eksempelkoden kan kompileres og utføres på den samme måten, som illustrert i punkt 1)-b) av Del 1. Man kan også kompilere og utføre et Golang-program med et kommando $ go run <din_fil>.go 

Det kan være fordel å bruke $ go run ... i steden for $ go build ... når man endrer kildekoden ofte for å teste resultatet som endringen gir.

a) endre responsen, som server-prosessen sender til klient-prosessen, slik at server-prosessen sender ditt navn (må endre kildekoden),

b) finn pakken i det grafiske grensesnittet til Wireshark, som inneholder data som server-prosessen sender til klient-prosessen

c) forklare hva kodene i TCP pakkene betyr (TCP push ACK - Develop Paper, 2021) og hvilke koder som brukes for den pakken, som inneholder ditt navn

d) finn ut de binære kodene for hver symbol i ditt navn (de kan du se i Wireshark)

# Del 2
Go modeller og Github repositorier (som skal være "public"). Se ressursen V&GB: Go moduler . Lage en skjermopptak av prosessen, inkludert en konsist forklaring av hva du viser. Levere en lenke i en skriftlig innleveringen for denne oppgaven (kan både skrive inn tekst direkte i Canvas, eller laste opp en fil).

1) Demonstrere inkludering av den eksterne modulen "rsc.io/quote" i en ny modul på en lokal node (din datamaskin). Demonstrere hvordan funksjonene Glass(), Go(), Hello() og Opt() i modulen "rsc.io/quote" funksjonerer (quote package - rsc.io/quote - pkg.go.dev, 2018), ved å påkalle disse fra en pakke i din modul (kan kalle denne pakken "myquote"). 

2) Lage en ny tom repository i Github og bruke den som et "hus" for den, foreløpig, lokale Go modulen. 

3) Markere distribusjon av module med en "tag" i Github repository. 

4) Demonstrere inkludering av den nye moduler (som ble laget i 1)-3) ) i et lokalt Go-program (man må da inkludere navn og versjon til din modul fra 1)-3) og påkalle metodene fra din pakke "myquote", som skriver ut resultater fra de fire nevnte funksjonene).   

5) Kan du foreslå bruk av Go "testing" modulen for denne oppgaven (du trenger ikke å skrive kode, du kan drøfte problemstillinger/utfordringer med testing i dette tilfelle)?

# Del 3
Vi betrakter en abstrakt verden for "Elvekryssing". Se beskrivelse av RC1 (River Crossing 1) problemet her V&GB: Enhetstesting i Golang

Skrive koden i ditt repository for test-drevet utvikling av RC1-verden. Du kan ta utgangspunkt i koden vi har gjennomgått på samlinger https://github.com/digitnow/rivercrossing (Lenker til en ekstern side.) men det er nødvendig med en del endringer, derfor anbefales det ikke å kopiere hele repository, men begynne å lage Go module og Github repository fra bunnen av slik som lært i Del 2.

For å kunne bruke testing ved hjelp av $ go test verktøyet, må man skrive tester i tråd med Go "testing" modul. 

Utforske fmt pakken kan være nyttig for å generer den grafiske fremstillingene av tilstandene i RC1-verden (fmt package - fmt - pkg.go.dev, 2021). 

I tillegg til ViewState-funksjonen, implementer også en PutInBoat-funksjon (kan sette HS, Rev, Kylling eller Korn i Båt) og CrossRiver-funksjon (flytter båt til andre siden av Elven), og som returnerer en ny tilstand. Forutsetning er å starte med test funksjon for put(), TestPut(t *testing T) (bruk Go sin testpakke).
