/*
Petlje
======

Petlja se koristi za izvršavanje bloka koda ponovljeno sve dok se ne ispuni
određeni uslov.

for je jedina petlja dostupna u Gou. Go nema while ili do while petlje koje su
prisutne u drugim jezicima poput C-a.

Sintaksa petlje for
-------------------

for initialisation; condition; post {

}

Workflow je sledeći:

	Inicijalizacije će se izvršiti samo jednom.
	Nakon što je petlja inicijalizovana, uslov se proverava.
	Ako uslov evaluira kao true, telo petlje unutar bloka { } će se izvršiti.
	Naredba post će se izvršiti nakon svake uspešne iterace petlje.
	Nakon što se naredba post izvrši, uslov će se ponovo proveriti.
	Ako je true, petlja će nastaviti s izvršavanjem.
	U suprotnom petlja završava.

Sve tri komponente, inicijalizacija, uslov i post, su opcionalne u Go-u.
Pogledajmo primer kako bismo bolje razumeli for petlju.

Napišimo program koji koristi for petlju za ispis brojeva od 1 do 10.
*/

package loops

import (
	"fmt"
)

func loopsElem() {

	fmt.Println("\n --- For loop ---")

	for i := 1; i <= 10; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}

/*
U gornjem programu, i je inicijalizovano na 1. Uslovna naredba će proveriti
da li je i <= 10. Ako je uslov true, vrednost "i" je ispisana, u suprotnom se
petlja završava. Naredba post povećava "i" za 1 na kraju svake iteracije. Kada
"i" bude veće od 10, petlja se završava.

Gornji program će ispisati:

	>> 1 2 3 4 5 6 7 8 9 10

Varijable deklarisane u for petlji dostupne su samo unutar opsega petlje. Stoga
im se ne može pristupiti izvan tela petlje.

Break naredba
-------------
Break naredba se koristi za nagli prekid for petlje i premeštanje kontrole u
liniju koda odmah nakon for petlje.

Hajde da modifikujemo gornji program tako da se for petlja prekida nakon ispisa
broja 5.
*/

func loopsBreak() {

	fmt.Println("\n --- loopsBreak ---")

	for i := 1; i <= 10; i++ {
		if i > 5 {
			break //loop is terminated if i > 5
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println("\nloop ended")
}

/*
U gornjem programu, vrednost "i" se proverava tokom svake iteracije. Ako je veće
od 5 tada se break izvršava i petlja se završava. Zatim se izvršava naredba
printf odmah nakon petlje for. Gornji program će ispisati,

	>> 1 2 3 4 5
	>> loop ended

Contiunue naredba
-----------------
Naredba continue se koristi za preskakanje trenutne iteracije for petlje. Sav
kod prisutan u for petlji nakon naredbe continue neće biti izvršen za trenutnu
iteraciju. Petlja će preći na sledeću iteraciju.

Napišimo program za ispis svih neparnih brojeva od 1 do 10 koristeći naredbu
continue.
*/

func loopsContinue() {

	fmt.Println("\n --- loopsContinue ---")

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}

/*
U gornjem programu, kod proverava se da li je ostatak deljenja "i" sa 2 jednak 0.
Ako je nula, tada je broj paran i naredba continue izvršava, a kontrola prelazi
na sledeću iteraciju petlje. Stoga se naredba printf nakon naredbe continue neće
pozvati i petlja nastavlja na sledeću iteraciju. Izlaz gornjeg programa je

	>> 1 3 5 7 9

Ugnežđene for petlje
--------------------
Petlja for koja ima drugu for petlju unutar sebe naziva se ugnežđenom for
petljom.

Hajde da shvatimo ugnežđene for petlje tako što ćemo napisati program koji
ispisuje sekvencu ispod.

*
**
***
****
*****

Program u nastavku koristi ugnežđene for petlje za ispis sekvence. Varijabla n
pohranjuje broj linija u sekvenci. U našem slučaju to je 5. Spoljnja for petlja
iterira "i" od 0 do 4, a unutrašnja for petlja iterira jod 0 do trenutne
vrednosti i. Unutrašnja petlja ispisuje * za svaku iteraciju, a spoljnja petlja
ispisuje novi red na kraju svake iteracije. Pokrenite ovaj program i videćete
sekvencu ispisanu na izlazu.
*/

func loopsNested() {

	fmt.Println("\n --- loopsNested ---")

	n := 5
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

/*
Labele
------
Labele se mogu koristiti za prekid spoljnje for petlje iz unutarašnje for petlje.
Hajde da shvatimo šta mislim koristeći jednostavan primer.

	func main() {
		for i := 0; i < 3; i++ {
			for j := 1; j < 4; j++ {
				fmt.Printf("i = %d , j = %d\n", i, j)
			}
		}
	}

Gornji program je samorazumljiv i ispisat će:

	>> i = 0 , j = 1
	>> i = 0 , j = 2
	>> i = 0 , j = 3
	>> i = 1 , j = 1
	>> i = 1 , j = 2
	>> i = 1 , j = 3
	>> i = 2 , j = 1
	>> i = 2 , j = 2
	>> i = 2 , j = 3

Ništa posebno u ovome :)

Šta ako želimo prestati sa ispisivanjem kada su i i j jednaki? Da bismo to
uradili, moramo učiniti break iz vanjske for petlje. Dodavanje break u
unutrašnju for petlju kada su i i j jednaki će prekinuti ispis samo iz
unutrašnje for petlje.

	func main() {
		for i := 0; i < 3; i++ {
			for j := 1; j < 4; j++ {
				fmt.Printf("i = %d , j = %d\n", i, j)
				if i == j {
					break
				}
			}
		}
	}

U gornjem programu, dodao sam break unutar unutrašnje for petlje kada su i i j
jednaki. Ovo će se izvršavati break samo iz unutrašnje for petlje, a vanjska
petlja će se nastaviti. Ovaj program će ispisati.

	>> i = 0 , j = 1
	>> i = 0 , j = 2
	>> i = 0 , j = 3
	>> i = 1 , j = 1
	>> i = 2 , j = 1
	>> i = 2 , j = 2

Ovo nije nameravani izlaz. Moramo prestati sa ispisom kada su oba i i j jednaka,
tj. kada su jednaka 1.

Tu nam u pomoć priskaču labele. Labela se može koristiti za prekid vanjske
petlje. Prepišimo gornji program koristeći labele:
*/

func loopsLabel() {

	fmt.Println("\n --- loopsLabel ---")

outer:
	for i := 0; i < 3; i++ {
		for j := 1; j < 4; j++ {
			fmt.Printf("i = %d , j = %d\n", i, j)
			if i == j {
				break outer
			}
		}
	}
}

/*
U gornjem programu, dodali smo oznaku "outer" na spoljnjoj for petlji i
prekidamo spoljnju for petlju navođenjem ove oznake pored break iz unutrašnje
petlje. Ovaj program će prestati sa ispisom kada su oba i i j jednaka. Ovaj
program će ispisati:

	>> i = 0 , j = 1
	>> i = 0 , j = 2
	>> i = 0 , j = 3
	>> i = 1 , j = 1

Petlja While koristeći petlju for
---------------------------------
Ranije smo raspravljali o tome da je for petlja jedina naredba petlje dostupna
u Gou. Moguće je koristiti varijaciju for petlje kako bi se postigla
funkcionalnost petlje while.

Hajde da razgovaramo o tome kako se ovo može uraditi. Program ispod ispisuje
sve parne brojeve od 0 do 10.
*/

func loopsForWhile() {

	fmt.Println("\n --- loopsForWhile ---")

	i := 0
	for i <= 10 { // initialisation and post are omitted
		fmt.Printf("%d ", i)
		i += 2
	}
	fmt.Println()
}

/*
Kao što već znamo, sve tri komponente for petlje, inicijalizacija, uslov i post,
su opcionalne. U gornjem programu, inicijalizacija i post su izostavljeni. i je
inicijalizovan na 0 izvan for petlje. Petlja će se izvršavati sve dok i <= 10 ne
bude true. i se povećava za 2 unutar for petlje. Gornji program daje izlaz:

	>> 0 2 4 6 8 10

Tačka-zarez u for petlji gornjeg programa se takođe može izostaviti. Ovaj format
se može smatrati alternativom for - while petljom. Gornji program se može
prepisati kao:
*/

func loopsWhile() {

	fmt.Println("\n --- loopsWhile ---")

	i := 0
	for i <= 10 { //semicolons are ommitted and only condition is present.
		fmt.Printf("%d ", i)
		i += 2
	}
	fmt.Println()
}

/*
Višestruke deklaracije varijabli
--------------------------------
Moguće je deklarisati i operisati nad više varijabli u for petlji. Napišimo
program koji ispisuje donji niz koristeći više deklaracija varijabli:

10 * 1 = 10
11 * 2 = 22
12 * 3 = 36
13 * 4 = 52
14 * 5 = 70
15 * 6 = 90
16 * 7 = 112
17 * 8 = 136
18 * 9 = 162
19 * 10 = 190
*/

func loopsMultiVars() {

	fmt.Println("\n --- loopsMultiVars ---")

	//multiple initialisation and increment
	for i, j := 1, 11; i <= 10 && j <= 20; i, j = i+1, j+1 {
		fmt.Printf("%d * %d = %d\n", i, j, i*j)
	}
}

/*
U gornjem programu i i j su deklarisani i inicijalizirani na 1 i 11
respektivno. Na kraju svake iteracije se povećavaju za 1. Boolean operator &&
se koristi u uslovu kako bi se osiguralo da je i manje ili jednako 10, a takođe
i j je manje ili jednako 20.

Beskonačna petlja
-----------------
Sintaksa za kreiranje beskonačne petlje je,

for {
}

Sledeći program će nastaviti sa štampanjem Hello World kontinuirano bez
prekidanja.
*/

// func loopsForever() {

// 	fmt.Println("\n --- loopsForever ---")

// 	for {
// 		fmt.Println("Hello World")
// 	}
// }

/*
Ako pokušate pokrenuti gornji program na Go Playgroundu, dobićete grešku
"timeout running program". Molimo pokušajte ga pokrenuti na vašem lokalnom
sistemu da biste beskonačno ispisivali "Hello World".

Postoji još jedan konstruktivni raspon - range koji se može koristiti u for
petljama za manipulaciju nizovima. O tome ćemo govoriti kada budemo učili o
nizovima.
*/

func Loops() {

	fmt.Println("\n --- Loops ---")

	loopsElem()
	loopsBreak()
	loopsContinue()
	loopsNested()
	loopsLabel()
	loopsForWhile()
	loopsWhile()
	loopsMultiVars()
	// loopsForever()
}
