/*
Panic i recover
===============

Šta je panic?
-------------
Idiomatski način rešavanja abnormalnih stanja u Go programu je korišćenje
grešaka. Greške su dovoljne za većinu abnormalnih stanja koja se javljaju u
programu.

Ali postoje neke situacije u kojima program ne može da nastavi izvršavanje nakon
abnormalnog stanja. U ovom slučaju, koristimo "panic" za prevremeno prekidanje
programa. Kada funkcija naiđe na "panic", njeno izvršavanje se zaustavlja,
sve odložene funkcije se izvršavaju, a zatim se kontrola vraća pozivaocu.
Ovaj proces se nastavlja sve dok se sve funkcije trenutne gorutine ne vrate, u
kom trenutku program ispisuje poruku panic, nakon čega sledi trag steka, a
zatim se prekida. Ovaj koncept će biti jasniji kada napišemo primer programa.

Moguće je povratiti kontrolu nad programom koji izaziva "panic", o "recover"
čemu ćemo kasnije govoriti u ovom tutorijalu.

„panic“ i „recover“ se mogu smatrati sličnim idiomu „try-catch-finally“ u
drugim jezicima kao što je Java, osim što se retko koriste u Gou.

Kada treba koristiti paniku?
----------------------------
Jedan važan faktor je da treba izbegavati paniku i oporavljati se od grešaka
kad god je to moguće. Samo u slučajevima kada program jednostavno ne može da
nastavi sa izvršavanjem, trebalo bi koristiti mehanizam za paniku i oporavak.

Postoje dva validna slučaja upotrebe panike.

  - Nepopravljiva greška gde program jednostavno ne može da nastavi svoje
    izvršavanje. Jedan primer je veb server koji ne uspeva da se poveže sa
    potrebnim portom. U ovom slučaju, razumno je paničiti jer nema šta drugo da
    se uradi ako samo povezivanje porta ne uspe.

  - Greška programera. Recimo da imamo metodu koja prihvata pokazivač kao
    parametar i neko pozove ovu metodu koristeći nil argument. U ovom slučaju,
    možemo paničiti jer je to greška programera pozvati metodu sa ni largumentom
    koji je očekivao validan pokazivač.

Primer panike
-------------
Potpis ugrađene panic funkcije je dat u nastavku,

	func panic(interface{})

Argument prosleđen funkciji panic biće ispisan kada se program završi. Upotreba
ovoga će biti jasna kada budemo pisali primer programa. Zato hajde da to uradimo
odmah.

Počećemo sa izmišljenim primerom koji pokazuje kako panika funkcioniše.
*/

package de

import (
	"fmt"
	"runtime/debug"
)

// func fullName(firstName *string, lastName *string) {
// 	if firstName == nil {
// 		panic("runtime error: first name cannot be nil")
// 	}
// 	if lastName == nil {
// 		panic("runtime error: last name cannot be nil")
// 	}
// 	fmt.Printf("%s %s\n", *firstName, *lastName)
// 	fmt.Println("returned normally from fullName")
// }

// func panicExample() {

// 	fmt.Println("\n --- Panic example ---")

// 	firstName := "Elon"
// 	fullName(&firstName, nil)
// 	fmt.Println("returned normally from main")
// }

/*
Gore navedeno je jednostavan program za ispisivanje punog imena osobe. Funkcija
fullName ispisuje puno ime osobe. Ova funkcija proverava da li su pokazivači
firstName i lastName nil. Ako jesu, nil funkcija se poziva panic sa
odgovarajućom porukom. Ova poruka će biti ispisana kada se program završi.

Pokretanje ovog programa će ispisati sledeći izlaz,

	>> panic: runtime error: last name cannot be nil
	>>
	>> goroutine 1 [running]:
	>> main.fullName(0xc00006af58, 0x0)
	>> 	/tmp/sandbox210590465/prog.go:12 +0x193
	>> main.main()
	>> 	/tmp/sandbox210590465/prog.go:20 +0x4d

Hajde da analiziramo ovaj izlaz da bismo razumeli kako panika funkcioniše i kako
se ispisuje trag steka kada program paniči.

Prvo dodeljujemo Elon u "firstName". Pozivamo funkciju "fullName" sa "nil" u
"lastName". Stoga će uslov "lastName == nil" biti zadovoljen i program će
paničiti. Kada se pojavi panika, izvršavanje programa se prekida, ispisuje se
argument prosleđen funkciji "panic", a zatim i trag steka. Pošto se program
prekida nakon poziva funkcije "panic", daljnji kod u glavnoj funkciji neće biti
izvršen.

Ovaj program prvo ispisuje poruku koja je prosleđena funkciji panic,

	>> panic: runtime error: last name cannot be nil

a zatim ispisuje trag steka.

Još jedan primer
----------------
Panike takođe mogu biti uzrokovane greškama koje se dešavaju tokom izvršavanja,
kao što je pokušaj pristupa indeksu koji nije prisutan u isečku.

Hajde da napišemo izmišljen primer koji stvara paniku zbog pristupa isečku van
granica.
*/

// func slicePanic() {
// 	n := []int{5, 7, 4}
// 	fmt.Println(n[4])
// 	fmt.Println("normally returned from a")
// }

// func panicSlice() {

// 	fmt.Println("\n --- Panic slice ---")

// 	slicePanic()
// 	fmt.Println("normally returned from main")
// }

/*
U gornjem programu, da pristupimo n[4] što je nevažeći indeks u isečku. Ovaj
program će izbaciti sledeći izlaz,

	>> panic: runtime error: index out of range [4] with length 3
	>>
	>> goroutine 1 [running]:
	>> main.slicePanic()
	>> 	/tmp/sandbox942516049/prog.go:9 +0x1d
	>> main.main()
	>> 	/tmp/sandbox942516049/prog.go:13 +0x22

Odloženi pozivi tokom panike
----------------------------
Podsetimo se šta panika radi. Kada funkcija naiđe na paniku, njeno izvršavanje
se zaustavlja, sve odložene funkcije se izvršavaju, a zatim se kontrola vraća
pozivaocu. Ovaj proces se nastavlja sve dok se sve funkcije trenutne gorutine
ne vrate, u kom trenutku program ispisuje poruku panike, nakon čega sledi trag
steka, a zatim se završava.

U gornjem primeru, nismo odložili nijedan poziv funkcije. Ako postoji odloženi
poziv funkcije, on se izvršava, a zatim se kontrola vraća pozivaocu.

Hajde da malo izmenimo gornji primer i koristimo naredbu "defer".
*/

// func fullName(firstName *string, lastName *string) {

// 	defer fmt.Println("deferred call in fullName")

// 	if firstName == nil {
// 		panic("runtime error: first name cannot be nil")
// 	}
// 	if lastName == nil {
// 		panic("runtime error: last name cannot be nil")
// 	}
// 	fmt.Printf("%s %s\n", *firstName, *lastName)
// 	fmt.Println("returned normally from fullName")
// }

// func panicWithDefer() {

// 	fmt.Println("\n --- Panic with defer ---")
// 	defer fmt.Println("deferred call in main")

// 	firstName := "Elon"
// 	fullName(&firstName, nil)
// 	fmt.Println("returned normally from main")
// }

/*
Jedine izmene koje su napravljene su dodavanje odloženih poziva funkcija.

Ovaj program štampa,

	>> deferred call in fullName
	>> deferred call in main
	>>
	>> panic: runtime error: last name cannot be nil
	>>
	>> goroutine 1 [running]:
	>> main.fullName(0xc00006af28, 0x0)
	>> 	/tmp/sandbox451943841/prog.go:13 +0x23f
	>> main.main()
	>> 	/tmp/sandbox451943841/prog.go:22 +0xc6

Kada program doživi paniku u liniji svi odloženi pozivi funkcija se prvo
izvršavaju, a zatim se kontrola vraća pozivaocu čiji se odloženi pozivi
izvršavaju i tako dalje dok se ne dođe do pozivaoca najvišeg nivoa.

U našem slučaju, prvo se izvršava defer naredba u fullName funkcije. Ovo
ispisuje sledeću poruku:

	>> deferred call in fullName

A zatim se kontrola vraća glavnoj funkciji čiji se odloženi pozivi izvršavaju i
stoga se ovo ispisuje,

	>> deferred call in main

Sada je kontrola dostigla funkciju najvišeg nivoa i stoga program ispisuje
poruku panike praćenu tragom steka, a zatim se završava.

Oporavak od panike
------------------
"recover" je ugrađena funkcija koja se koristi za ponovno preuzimanje kontrole
nad programom koji izaziva paniku.

Potpis funkcije oporavka je dat ispod,

	func recover() interface{}

Funkcija "recover" je korisna samo kada se poziva unutar odloženih funkcija.
Izvršavanje poziva "recover" unutar odložene funkcije zaustavlja sekvencu koja
izaziva "panic" obnavljanjem normalnog izvršavanja i preuzima poruku o grešci
koja je prosleđena pozivu funkcije panic.

Ako se funkcija "recover" pozove izvan odložene funkcije, neće zaustaviti
sekvencu koju izaziva "panic".

Hajde da modifikujemo naš program i koristimo "recover" da bismo vratili
normalno izvršavanje nakon "panic".
*/

func recoverFullName() {
	if r := recover(); r != nil {
		fmt.Println("recovered from ", r)
	}
}

func fullName(firstName *string, lastName *string) {

	defer recoverFullName()

	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}

func recoverExample() {

	fmt.Println("\n --- Recover example ---")

	defer fmt.Println("deferred call in recoverExample")

	firstName := "Elon"
	fullName(&firstName, nil)
	fmt.Println("returned normally from recoverExampe()")
}

/*
Funkcija "recoverFullName()"" poziva funkciju "recover()"" koja vraća vrednost
prosleđenu pozivu "panic" funkcije. Ovde samo ispisujemo vrednost koju vraća
funkcija recover, "recoverFullName()"" se odlaže u funkcije "fullName".

Ovaj program će štampati:

	>> recovered from  runtime error: last name cannot be nil
	>> returned normally from main
	>> deferred call in main

Nakon izvršavanja recover(), panika prestaje i kontrola se vraća pozivaocu,
u ovom slučaju,funkciji "recoverExample()". Program nastavlja da se normalno
izvršava u funkciji "recoverExample()" pošto je panika prevaziđena.

Pogledajmo još jedan primer gde se oporavljamo od panike izazvane pristupom
nevažećem indeksu isečka.
*/

func recoverInvalidAccess() {
	if r := recover(); r != nil {
		fmt.Println("Recovered", r)
	}
}

func invalidSliceAccess() {
	defer recoverInvalidAccess()

	n := []int{5, 7, 4}
	fmt.Println(n[4])
	fmt.Println("normally returned from a")
}

func recoverInvalidSliceAccess() {

	fmt.Println("\n --- Recover invalid slice access ---")

	invalidSliceAccess()
	fmt.Println("normally returned from main")
}

/*
Pokretanje gore navedenog programa će ispisati,

	>> Recovered runtime error: index out of range [4] with length 3
	>> normally returned from main

Iz rezultata možete shvatiti da smo se oporavili od panike.

Dobijanje traga steka nakon oporavka
------------------------------------
Ako se oporavimo od panike, gubimo trag steka o panici. Čak i u programu iznad,
nakon oporavka, izgubili smo trag steka.

Postoji način da se odštampa trag steka koristeći funkciju PrintStack iz Debug
paketa:
*/

func recoverFullName2() {
	if r := recover(); r != nil {
		fmt.Println("recovered from ", r)
		debug.PrintStack()
	}
}

func fullName2(firstName *string, lastName *string) {

	defer recoverFullName2()

	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName2")
}

func recoverExample2() {

	fmt.Println("\n --- Recover example2---")

	defer fmt.Println("deferred call in recoverExample2")

	firstName := "Elon"
	fullName2(&firstName, nil)
	fmt.Println("returned normally from recoverExample2")
}

/*
U gornjem programu, koristimo "debug.PrintStack()" za ispis traga steka.

Ovaj program će štampati:

	>> recovered from  runtime error: last name cannot be nil
	>> goroutine 1 [running]:
	>> runtime/debug.Stack(0x37, 0x0, 0x0)
	>> 	/usr/local/go-faketime/src/runtime/debug/stack.go:24 +0x9d
	>> runtime/debug.PrintStack()
	>> 	/usr/local/go-faketime/src/runtime/debug/stack.go:16 +0x22
	>> main.recoverFullName()
	>> 	/tmp/sandbox771195810/prog.go:11 +0xb4
	>> panic(0x4a1b60, 0x4dc300)
	>> 	/usr/local/go-faketime/src/runtime/panic.go:969 +0x166
	>> main.fullName(0xc0000a2f28, 0x0)
	>> 	/tmp/sandbox771195810/prog.go:21 +0x1cb
	>> main.main()
	>> 	/tmp/sandbox771195810/prog.go:30 +0xc6
	>> returned normally from main
	>> deferred call in main

Iz izlaza možete videti da je panika oporavljena i

	>> recovered from runtime error: last name cannot be nil

da je ispisana. Nakon toga, ispisuje se trag steka. Zatim

	>> returned normally from main
	>> deferred call in main

štampa se nakon što se panika oporavi.

Panika, oporavak i gorutine
---------------------------
Funkcija „recover“ radi samo kada se poziva iz iste gorutine koja izaziva
paniku. Nije moguće oporaviti se od panike koja se dogodila u drugoj gorutini.
Hajde da ovo razumemo na primeru:
*/

func recovery() {
	if r := recover(); r != nil {
		fmt.Println("recovered:", r)
	}
}

func sum(a int, b int) {

	defer recovery()

	fmt.Printf("%d + %d = %d\n", a, b, a+b)
	done := make(chan bool)
	go divide(a, b, done)
	<-done
}

func divide(a int, b int, done chan bool) {
	fmt.Printf("%d / %d = %d", a, b, a/b)
	done <- true
}

func recoverGoroutine() {
	fmt.Println("\n --- Recover goroutine ---")
	sum(5, 0)
	fmt.Println("normally returned from main")
}

/*
U gornjem programu, funkcija divide()će izazvati paniku jer je b nula a nije
moguće podeliti broj sa nulom. Funkcija sum() poziva odloženu funkciju recovery()
koja se koristi za oporavak od panike. Funkcija divide() se poziva kao posebna
gorutina. Čekamo na done kanalu da bismo osigurali da se divide() izvršavanje
završi.

Šta mislite da će biti izlaz programa. Da li će se panika oporaviti? Odgovor je
ne. Panika se neće oporaviti. To je zato što recovery je funkcija prisutna u
drugoj gorutini i panika se dešava u divide() funkciji u drugoj gorutini.
Stoga oporavak nije moguć.

Pokretanje ovog programa će ispisati,

	>> 5 + 0 = 5
	>> panic: runtime error: integer divide by zero
	>>
	>> goroutine 18 [running]:
	>> main.divide(0x5, 0x0, 0xc0000a2000)
	>> 	/tmp/sandbox877118715/prog.go:22 +0x167
	>> created by main.sum
	>> 	/tmp/sandbox877118715/prog.go:17 +0x1a9

Iz rezultata možete videti da se oporavak nije dogodio.

Ako divide() se funkcija pozove u istoj gorutini, oporavili bismo se od panike.

Ako se red sa  go divide(a, b, done) na divide(a, b, done) oporavak će se
desiti jer se panika dešava u istoj gorutini. Ako se program pokrene sa gore
navedenom promenom, ispisaće

	>> 5 + 0 = 5
	>> recovered: runtime error: integer divide by zero
	>> normally returned from main
*/

func PanicRecoverFunc() {
	fmt.Println("\n --- Panic Recover ---")

	// panicExample()
	// panicSlice()
	// panicWithDefer()
	recoverExample()
	recoverInvalidSliceAccess()
	recoverExample2()
	recoverGoroutine()

}
