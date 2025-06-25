/*
Upisivanje u datoteku
=====================

Upisivanje stringa u datoteku
-----------------------------

Jedna od najčešćih operacija pisanja datoteka je pisanje stringa u datoteku.
Ovo je prilično jednostavno za uraditi. Sastoji se od sledećih koraka:

	1. Napravite datoteku
	2. Zapišite string u datoteku

Hajde odmah da pređemo na kod.
*/

package files

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
)

func writeString() {

	fmt.Println("\n --- Write string to file ---")
	f, err := os.Create("wtest.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	l, err := f.WriteString("Hello World")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}

	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

/*
Funkcija "Create" gornjeg programa kreira datoteku pod nazivom "wtest.txt". Ako
datoteka sa tim imenom već postoji, onda funkcija create skraćuje datoteku na
nultu dužinu. Ova funkcija vraća deskriptor datoteke f.

Zatim upisujemo string "Hello World" u datoteku koristeći "WriteString" gmetodu.
Ova metoda vraća broj zapisanih bajtova i grešku ako postoji.

Konačno, zatvaramo datoteku u redu br. 21.

Gore navedeni program će ispisati

	>> 11 bytes written successfully

Ako otvorite datoteku pomoću bilo kog uređivača teksta, videćete da sadrži tekst
"Hello world.

Upisivanje bajtova u datoteku
-----------------------------
Upisivanje bajtova u datoteku je prilično slično pisanju stringa u datoteku.
Koristićemo metod "os.Write" za pisanje bajtova u datoteku. Sledeći program upisuje
isečak bajtova u datoteku.
*/

func writeBytes() {

	fmt.Println("\n --- Write bytes to file ---")

	f, err := os.Create("wbytes")
	if err != nil {
		fmt.Println(err)
		return
	}

	d2 := []byte{104, 101, 108, 108, 111, 32, 98, 121, 116, 101, 115}
	n2, err := f.Write(d2)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}

	fmt.Println(n2, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

/*
U gornjem programu, koristimo metod "Write" da zapišemo isečak bajtova u
datoteku pod nazivom wbytes u direktorijum ""./14-files". Možete promeniti ovaj
direktorijum u neki. Preostali deo programa je sam po sebi razumljiv.

Ovaj program će ispisati

>> 11 bytes written successfully

i kreirati datoteku pod nazivom bytes. Otvorite datoteku i videćete da ona
sadrži tekst "hello bytes".

Upisivanje stringova linija po linija u datoteku
------------------------------------------------
Još jedna uobičajena operacija sa datotekama je potreba za upisivanjem stringova
u datoteku red po red. U ovom odeljku ćemo napisati program za kreiranje
datoteke sa sledećim sadržajem:

	Welcome to the world of Go.
	Go is a compiled language.
	It is easy to learn Go.

Hajde odmah da pređemo na kod.
*/

func writeSliceOfStrings() {

	fmt.Println("\n --- Write slice strings to file ---")

	f, err := os.Create("wlines")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}

	d := []string{
		"Welcome to the world of Go1.",
		"Go is a compiled language.",
		"It is easy to learn Go.",
	}
	for _, v := range d {
		if _, err := fmt.Fprintln(f, v); err != nil {
			fmt.Println(err)
			return
		}
	}

	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file written successfully")
}

/*
U gornjem programu, kreiramo novu datoteku pod nazivom wlines. Iteriramo kroz
isečak stringova koristeći petlju for range i koristimo funkciju "Fprintln" da
upišemo linije u datoteku. Funkcija "Fprintln" uzima io.writer parametar a kao
i dodaje novu liniju, upravo ono što smo želeli.

Pokretanje ovog programa će ispisati rezultat

	>> file written successfully

i datoteka wlines će biti kreirana u trenutnom direktorijumu. Sadržaj datoteke
lines je dat u nastavku:

	>> Welcome to the world of Go1.
	>> Go is a compiled language.
	>> It is easy to learn Go.

Dodavanje u datoteku
--------------------
U ovom odeljku, dodaćemo još jedan red datoteci "wlines" koju smo kreirali u
prethodnom odeljku. Dodaćemo red "File handling is easy" u wlines.
datoteku.

Datoteka mora biti otvorena u režimu dodavanja i pisanja. Ove zastavice se
prosleđuju kao parametri funkciji Open. Nakon što se datoteka otvori u režimu
dodavanja, dodajemo novi red u datoteku.
*/

func writeAppend() {

	fmt.Println("\n --- Write append to file ---")

	f, err := os.OpenFile("wlines", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	newLine := "File handling is easy."
	_, err = fmt.Fprintln(f, newLine)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}

	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("file appended successfully")
}

/*
U gornjem programu, otvaramo datoteku u režimu dodavanja i pisanja. Nakon što
je datoteka uspešno otvorena, dodajemo novi red u datoteku. Ovaj program će
ispisati

	>> file appended successfully.

Nakon pokretanja ovog programa, sadržaj datoteke wlines će biti,

	>> Welcome to the world of Go1.
	>> Go is a compiled language.
	>> It is easy to learn Go.
	>> File handling is easy.

Istovremeno pisanje u datoteku
------------------------------
Kada više gorutina istovremeno piše u datoteku, doći će do uslova trke. Stoga,
istovremena pisanja u datoteku moraju biti koordinisana korišćenjem kanala.

Napisaćemo program koji kreira 100 gorutina. Svaka od ovih gorutina će
generisati slučajni broj konkurentno, čime će se generisati ukupno stotinu
slučajnih brojeva. Ovi slučajni brojevi će biti zapisani u datoteku.

Problem uslova trke rešićemo koristeći sledeći pristup:

	1. Napravimo kanal koji će se koristiti za čitanje i pisanje generisanih
	   slučajnih brojeva.
	2. Napravite 100 gorutina - "produce"-ra. Svaka gorutina će generisati
	   slučajni broj i takođe će upisati slučajni broj u kanal.
	3. Napravimo gorutinu consumer koja će čitati iz kanala i upisivati
	   generisani slučajni broj u datoteku. Na taj način imamo samo jednu
	   gorutinu koja istovremeno piše u datoteku, čime se izbegava uslov trke :)
	4. Zatvorimo datoteku kada završimo.

Prvo ćemo napisati "produce" funkciju koja generiše slučajne brojeve.
*/

func produce(data chan int, wg *sync.WaitGroup) {
	n := rand.Intn(999)
	data <- n
	wg.Done()
}

/*
Gore navedena "produce" funkcija generiše slučajni broj i upisuje ga u kanal
data a zatim poziva Done grupe čekanja wg da je obavesti da je završila svoj
zadatak.

Hajde sada da pređemo na funkciju koja piše u datoteku.
*/

func consume(data chan int, done chan bool) {

	f, err := os.Create("wconcurrent")
	if err != nil {
		fmt.Println(err)
		return
	}

	for d := range data {
		_, err = fmt.Fprintln(f, d)
		if err != nil {
			fmt.Println(err)
			f.Close()
			done <- false
			return
		}
	}

	err = f.Close()
	if err != nil {
		fmt.Println(err)
		done <- false
		return
	}
	done <- true
}

/*
Funkcija "consume" kreira datoteku pod nazivom "concurrent". Zatim čita slučajne
brojeve iz data kanala i upisuje ih u datoteku. Kada pročita i upiše sve
slučajne brojeve, upisuje true u done kanal kako bi obavestila da je završila
svoj zadatak.

Hajde da napišemo glavnu funkciju.
*/
func writeConcurently() {

	fmt.Println("\n --- Write concurently to file ---")

	data := make(chan int)
	done := make(chan bool)

	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go produce(data, &wg)
	}

	go consume(data, done)

	go func() {
		wg.Wait()
		close(data)
	}()

	d := <-done

	if d {
		fmt.Println("File concurently written successfully")
	} else {
		fmt.Println("File concurently writing failed")
	}
}

/*
Glavna funkcilja kreira data kanal iz kojeg se čitaju i upisuju slučajni brojevi.
Kanal done koristi consume gorutina da obavesti glavnu funkciju da je završila
svoj zadatak. Grupa čekanja wg se koristi da sačeka da svih 100 gorutina završi
generisanje slučajnih brojeva.

Petja for kreira 100 gorutina. Poziv wait() grupe čekanja kaže da sačeka da svih
100 gorutina završi kreiranje slučajnih brojeva. Nakon toga, zatvara kanal.
Kada se kanal zatvori i consume gorutina završi sa upisivanjem svih generisanih
slučajnih brojeva u datoteku, ona upisuje true u done kanal, a glavna gorutina
se deblokira i ispisuje

	>> File written successfully.

Sada možete konkuretntno otvoriti datoteku u bilo kom uređivaču teksta i videti
100 generisanih slučajnih brojeva :)
*/

func WriteFiles() {

	fmt.Println("\n --- Write files ---")

	writeString()
	writeBytes()
	writeSliceOfStrings()
	writeAppend()
	writeConcurently()
}
