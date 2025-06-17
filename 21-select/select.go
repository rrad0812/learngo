/*
Select
======

Naredba select se koristi za izbor između više operacija slanja/primanja kanala.
Naredba select blokira dok jedna od operacija slanja/primanja ne bude spremna.
Ako je više operacija spremno, jedna od njih se bira nasumično. Sintaksa je
slična sa switch osim što će svaka od naredbi case biti operacija kanala.

Primer
------
*/
package sel

import (
	"fmt"
	"time"
)

func server1(ch chan string) {
	time.Sleep(6 * time.Second)
	ch <- "from server1"
}
func server2(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "from server2"

}
func selExample() {

	fmt.Println("\n --- selExample ---")

	output1 := make(chan string)
	output2 := make(chan string)
	go server1(output1)
	go server2(output2)
	select {
	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-output2:
		fmt.Println(s2)
	}
}

/*
U gornjem programu, server1 funkcija je u stanju mirovanja 6 sekundi, a zatim
upisuje tekst "from server1" na kanal ch. server2 funkcija je u stanju
mirovanja 3 sekunde, a zatim upisuje sa servera2 na kanal ch.

Glavna funkcija poziva goroutines server1 i server2.

Kontrola stiže do select naredbe. select naredba blokira dok jedan od njenih
slučajeva ne bude spreman. U našem programu iznad, server1 upisuje u output1
kanal nakon 6 sekundi, dok server2 upisuje u output2 kanal nakon 3 sekunde.
Dakle, naredba select će se blokirati 3 sekunde i čekaće da server2 gorutina
upiše u output2 kanal. Nakon 3 sekunde, program ispisuje,

	>> from server2

a zatim će prekinuti.

Praktična upotreba select
-------------------------
Razlog za imenovanje funkcija u gornjem programu kao server1i server2 jeste
ilustracija praktične upotrebe funkcije select.

Pretpostavimo da imamo aplikaciju kritične važnosti i da nam je potrebno da što
pre vratimo rezultat korisniku. Baza podataka za ovu aplikaciju je replicirana
i sačuvana na različitim serverima širom sveta. Pretpostavimo da funkcije
server1 i server2 zapravo komuniciraju sa 2 takva servera. Vreme odziva svakog
servera zavisi od opterećenja na svakom od njih i kašnjenja mreže. Šaljemo
zahtev na oba servera, a zatim čekamo odgovor na odgovarajućim kanalima
koristeći naredbu select. Server koji prvi odgovori bira se selectom, a drugi
odgovor se ignoriše. Na ovaj način možemo poslati isti zahtev na više servera i
vratiti najbrži odgovor korisniku :).

Podrazumevani slučaj
--------------------
Podrazumevani slučaj u select naredbi se izvršava kada nijedan drugi slučaj
nije spreman. Ovo se generalno koristi da bi se sprečilo blokiranje naredbe
SELECT.
*/

func process(ch chan string) {
	time.Sleep(10500 * time.Millisecond)

	ch <- "process successful"
}

func selDefault() {

	fmt.Println("\n --- selDefault ---")

	ch := make(chan string)
	go process(ch)
	for {
		time.Sleep(1000 * time.Millisecond)
		select {
		case v := <-ch:
			fmt.Println("received value: ", v)
			return
		default:
			fmt.Println("no value received")
		}
	}
}

/*
U gornjem programu, funkcija process je u stanju mirovanja 10500 milisekundi
(10,5 sekundi), a zatim upisuje process successful na ch kanal. Ova funkcija
se poziva konkurentnoj u glavnoj gorutini.

Nakon konkurentnog pozivanja process gorutine, u glavnoj gorutini se pokreće
beskonačna for petlja. Beskonačna petlja miruje 1000 milisekundi (1 sekundu)
tokom početka svake iteracije, a zatim izvršava operaciju selekcije. Tokom
prvih 10500 milisekundi, prvi slučaj selekcije naredbe,

	>> case v := <-ch:

neće biti spreman jer će process gorutina pisati u ch kanal tek nakon 10500
milisekundi. Stoga će se default slučaj izvršiti tokom ovog vremena i program
će ispisati no value received 10 puta.

Nakon 10,5 sekundi, process gorutina upisuje process successful u ch. Sada će
se izvršiti prvi slučaj naredbe SELECT, program će ispisati, received value:
process successful a zatim će se završiti. Ovaj program će ispisati:

	>> no value received
	>> no value received
	>> no value received
	>> no value received
	>> no value received
	>> no value received
	>> no value received
	>> no value received
	>> no value received
	>> no value received
	>> received value:  process successful

# Zastoj i slučaj neizvršenja
-----------------------------
*/
/*
func selDeadlock() {
	ch := make(chan string)
	select {
	case <-ch:
	}
}
*/
/*
U gornjem programu, kreirali smo kanal ch. Pokušavamo da čitamo iz ovog kanala
unutar komande select. Naredba select će se blokirati zauvek jer nijedna druga
gorutina ne piše u ovaj kanal i stoga će dovesti do zastoja. Ovaj program će
paničiti tokom izvršavanja sa sledećom porukom,

	>> fatal error: all goroutines are asleep - deadlock!
	>>
	>> goroutine 1 [chan receive]:
	>> main.main()
	>>
	>> 	/tmp/sandbox627739431/prog.go:6 +0x4d

Ako postoji podrazumevani slučaj, do ove blokade neće doći jer će se
podrazumevani slučaj izvršiti kada nijedan drugi slučaj nije spreman. Gornji
program je prepisan sa podrazumevanim slučajem ispod.
*/

func selDeadlockWithDefault() {

	fmt.Println("\n --- selDeadlockWithDefault ---")

	ch := make(chan string)
	select {
	case <-ch:
	default:
		fmt.Println("default case executed")
	}
}

/*
Gore navedeni program će ispisati,

	>> default case executed

Slično tome, podrazumevani slučaj će se izvršiti čak i ako select ima samo
nil kanale.
*/

func selDeadlockWithDefaultAndNil() {

	fmt.Println("\n --- selDeadlockWithDefaultAndNil ---")

	var ch chan string
	select {
	case v := <-ch:
		fmt.Println("received value", v)
	default:
		fmt.Println("default case executed")
	}
}

/*
U gornjem programu ch je nil i pokušavamo da čitamo iz ch u selektu. Da default
slučaj nije bio prisutan, select bi se blokirao zauvek i izazvao bi zastoj.
Pošto imamo podrazumevani slučaj unutar selekta, on će biti izvršen i program
će ispisati,

	>> default case executed

Slučajni izbor
--------------
Kada je više slučajeva u select izjavi spremno, jedan od njih će biti izvršen
nasumično.
*/

func server11(ch chan string) {
	ch <- "from server11"

}
func server12(ch chan string) {
	ch <- "from server12"

}
func selChoose() {

	fmt.Println("\n --- selChoose ---")

	output1 := make(chan string)
	output2 := make(chan string)
	go server11(output1)
	go server12(output2)
	time.Sleep(1 * time.Second)
	select {
	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-output2:
		fmt.Println(s2)
	}
}

/*
U gornjem programu, pozivaju se rutine "go server11" i "go server12". Zatim
glavni program prelazi u stanje mirovanja na 1 sekundu. Kada kontrola dođe
do select naredbe oba case su spremna za izvršenje. Ako pokrenete ovaj program
više puta, izlaz će varirati između "from server11" ili "from server12 u
zavisnosti od toga koji je slučaj nasumično izabran.

Molimo vas da pokrenete ovaj program na vašem lokalnom sistemu da biste dobili
ovu slučajnost. Ako se ovaj program pokrene na igralištu, ispisaće isti izlaz
jer je igralište determinističko.

Uhvaćen - Prazan izbor
----------------------
*/
/*
func selWithoutCase() {
	select {}
}
*/
/*
Šta mislite da će biti rezultat rada gore navedenog programa?
Znamo da će se naredba select blokirati dok se ne izvrši jedan od njenih slučajeva.
Ovde, naredba select nema nijedan slučaj i stoga će se blokirati zauvek, što će
rezultirati zastojem. Ovaj program će paničiti sa sledećim izlazom,

	>> fatal error: all goroutines are asleep - deadlock!
	>>
	>> goroutine 1 [select (no cases)]:
	>> main.main()
	>>
	>> 	/tmp/sandbox246983342/prog.go:4 +0x25
*/
func SelectFunc() {

	fmt.Println("\n --- selectFunc ---")

	selExample()
	selDefault()
	selDeadlockWithDefault()
	selDeadlockWithDefaultAndNil()
	selChoose()
	//selWithoutCase()
}
