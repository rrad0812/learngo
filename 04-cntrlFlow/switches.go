/*
Switch naredba
==============

Šta je switch naredba?
----------------------
Switch je uslovna naredba koja evaluira izraz i poredi ga sa listom mogućih
podudaranja, te izvršava odgovarajući blok koda. Može se smatrati idiomatskim
načinom zamene složenih if else klauzula.

Primer programa govori stotinu reči. Počnimo s jednostavnim primerom koji će
kao ulaz uzeti broj prsta, a na izlazu ispisati naziv tog prsta :) .
Na primer, 1 je palac, 2 je kažiprst i tako dalje.
*/

package cntrl

import (
	"fmt"
	"math/rand"
)

func switchElem() {

	fmt.Println("\n --- SwitchElem ---")

	finger := 4
	fmt.Printf("Finger %d is ", finger)
	switch finger {
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3:
		fmt.Println("Middle")
	case 4:
		fmt.Println("Ring")
	case 5:
		fmt.Println("Pinky")
	}
}

/*
U gornjem programu, switch finger upoređuje vrednost finger sa svakom od case
naredbi. Slučajevi se evaluiraju od vrha do dna i izvršava se prvi slučaj koji
odgovara vrednosti switch izraza. U ovom slučaju, finger ima vrednost 4 i stoga:

	>> Finger 4 is Ring

je odštampan.

Duplikati slučajeva sa istom konstantnom vrednošću nisu dozvoljeni.
Ovde je primer:

func main() {

	finger := 4
 	fmt.Printf("Finger %d is ", finger)
	switch finger {
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3:
		fmt.Println("Middle")
	case 4:
		fmt.Println("Ring")
	case 4: //duplicate case
		fmt.Println("Another Ring")
	case 5:
		fmt.Println("Pinky")
	}
}

Pokretanje gornjeg programa će rezultirati sledećom greškom pri kompilaciji

	>> ./prog.go:19:7: duplicate case 4 (constant of type int) in expression
	>> switch ./prog.go:17:7: previous case

Podrazumevani slučaj
--------------------
Imamo samo 5 prstiju na rukama. Šta će se desiti ako unesemo pogrešan broj
prsta? Tu na scenu stupa podrazumevani slučaj. Podrazumevani slučaj će se
izvršiti kada se nijedan drugi slučaj ne podudara.
*/

func switchDefault() {

	fmt.Println("\n --- SwitchDefault ---")

	switch finger := 8; finger { // init finger and expression
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3:
		fmt.Println("Middle")
	case 4:
		fmt.Println("Ring")
	case 5:
		fmt.Println("Pinky")
	default: //default case
		fmt.Println("incorrect finger number")
	}
}

/*
U gornjem programu finger je 8 i ne odgovara nijednom od slučajeva i stoga

	>> incorrect finger numberse

iz podrazumevanog slučaja se ispisuje. Nije neophodno da default bude posljednji
slučaj u switch naredbi. Može biti prisutan bilo gdje u switch naredbi.

Možda ste primetili i malu promenu u deklaraciji finger. Deklarisan je u samom
bloku switch. Switch može uključivati ​​opcionalnu naredbu koja se izvršava pre
nego što se izraz evaluira. Ovde, finger se prvo deklariše, a zatim se koristi
u izrazu. Opseg fingeru ovom slučaju je ograničen na blok switch.
Višestruki izrazi u slučaju

Više slučaja u jednom case izrazu
---------------------------------
Moguće je uključiti više izraza u jedan slučaj odvajajući ih zarezom.
*/

func switchMultiExpr() {

	fmt.Println("\n --- SwitchMultiExpr ---")

	letter := "i"
	switch letter {
	case "a", "e", "i", "o", "u": //multiple expressions in case
		fmt.Printf("%s is a vowel", letter)
	default:
		fmt.Printf("%s is not a vowel", letter)
	}
	fmt.Println()
}

/*
Gornji program pronalazi da li letter je samoglasnik ili ne. Kod

case "a", "e", "i", "o", "u":

odgovara bilo kome od samoglasnika. Pošto je i samoglasnik, ovaj program ispisuje

	>> i is a vowel

Prekidač bez izraza
-------------------
Izraz u switchu je opcionalan i može se izostaviti. Ako se izraz izostavi,
switch se smatra ispravnim (kao sa true izrazom) i svaki izraz case se
procjenjuje na istinitost, a zatim se izvršava odgovarajući blok koda.
*/

func switchWithoutExpr() {

	fmt.Println("\n --- SwitchWithOutWxpr ---")

	hour := 15 // hour in 24 hour format

	switch { // Using switch to determine the work shift
	case hour >= 6 && hour < 12:
		fmt.Println("It's the morning shift.")
	case hour >= 12 && hour < 17:
		fmt.Println("It's the afternoon shift.")
	case hour >= 17 && hour < 21:
		fmt.Println("It's the evening shift.")
	case (hour >= 21 && hour <= 24) || (hour >= 0 && hour < 6):
		fmt.Println("It's the night shift.")
	default:
		fmt.Println("Invalid hour.")
	}
}

/*
U gornjem programu, izraz nedostaje u switch naredbi i stoga se smatra true i
svaki od slučajeva se evaluira. U ovom slučaju "hour >= 12 && hour < 17" je
true i program ispisuje

	>> It's the afternoon shift.

Ova vrsta prekidača može se smatrati alternativom za višestruke if elseif
klauzule.

Fallthrough
-----------
U Go-u, kontrola se prenosi iz switch naredbe odmah nakon izvršenja slučaja.
fallthrough naredba se koristi za prenos kontrole na prvu naredbu slučaja
koja se nalazi odmah nakon izvršenog slučaja.

Napišimo program za razumevanje funkcije "fallthrough". Naš program će
proveriti da li je uneseni broj manji od 50, 100 ili 200. Na primer, ako
unesemo 75, program će ispisati da je 75 manje i od 100 i od 200.
To ćemo postići korištenjem fallthrough.
*/

func number() int {
	num := 15 * 5
	return num
}

func switchFallthrough() {

	fmt.Println("\n --- SwitchFallthrough ---")

	switch num := number(); { //num is not a constant, expr is missed
	case num < 50:
		fmt.Printf("%d is lesser than 50\n", num)
		fallthrough
	case num < 100:
		fmt.Printf("%d is lesser than 100\n", num)
		fallthrough
	case num < 200:
		fmt.Printf("%d is lesser than 200", num)
	}
	fmt.Println()
}

/*
Izrazi u switch i case ne moraju biti samo konstante. Mogu se evaluirati za
vreme izvršavanja. U gornjem programu num je inicijalizirano na povratnu
vrednost funkce number(). Kontrola dolazi unutar switch-a i slučajevi se
evaluiraju. "case num < 100:" je true i program ispisuje

	>> 75 is lesser than 100.

Sledeća naredba je fallthrough. Sada kontrola prelazi na prvu naredbu sledećeg
slučaja i takođe ispisuje

	>> 75 is lesser than 200.

Izlaz programa je

	>> 75 is lesser than 100
	> 75 is lesser than 200

fallthrough treba da bude poslednja naredba u case. Ako se nalazi negde u
sredini, kompajler će se žaliti da je "fallthrough statement out of place".

Postoji suptilnost koju treba uzeti u obzir pri korištenju fallthrough.
Propadanje će se dogoditi čak i kada se slučaj proceni kao netačan.
*/

func switchCaseFalseFallthrough() {

	fmt.Println("\n --- switchCaseFalseFallthrough ---")

	switch num := 25; { // Expe is missed
	case num < 50:
		fmt.Printf("%d is lesser than 50\n", num)
		fallthrough
	case num > 100:
		fmt.Printf("%d is greater than 100\n", num)
	}
	fmt.Println()
}

/*
U gornjem programu, num je 25 što je manje od 50 i stoga se ovaj slučaj
evaluira kao true. Zbog prisutnosti fallthrough iako je sledeći slučaj
case num > 100:u false jer je num < 100. Ali fallthrough ovo ne uzima u obzir.
Fallthrough će se dogoditi čak i ako se slučaj evaluira kao false.

Gornji program će ispisati

	>> 25 is lesser than 50
	>> 25 is greater than 100

Zato budite sigurni da razumete šta radite kada koristite fallthrough.

Još jedna stvar, fallthrough se ne može koristiti u poslednjem slučaju switch-a
jer nema više slučajeva za propadanje. Ako je fallthrough je prisutan u
posljednjem slučaju, to će rezultirati sledećom greškom pri kompajliranju:
"cannot fallthrough final case in switch"

Break
-----
Naredba break se može koristiti za rano prekidanje switch-a pre nego što se
normalno završi. Hajde da samo modifikujemo gornji primer kako bismo razumeli
kako break funkcioniše.

Dodajmo uslov da ako je num manje od 0, tada switch završava.
*/

func switchBreak() {

	fmt.Println("\n --- switchBreak ---")

	switch num := -5; {
	case num < 50:
		if num < 0 {
			fmt.Println("num is less than 0")
			break
		}
		fmt.Printf("%d is lesser than 50\n", num)
		fallthrough
	case num < 100:
		fmt.Printf("%d is lesser than 100\n", num)
		fallthrough
	case num < 200:
		fmt.Printf("%d is lesser than 200", num)
	}
}

/*
U gornjem programu num je -5. Kada kontrola dođe do if naredbe uslov je ispunjen
jer je num < 0. Naredba break prekida izmenu switcha pre nego što se završi i
program ispisuje:

>> num is less tnan 0

i završava.

Prekid spoljnje for petlje
--------------------------
Kada se slučaj switcha nalazi unutar for petlje, možda će biti potrebno rano
prekinuti for petlju. To se može učiniti labeliranjem for petlje i prekidanjem
for petlje korištenjem te labele unutar switch naredbe.

Napišimo program za generisanje slučajnog parnog broja.
Kreiraćemo beskonačnu for petlju i koristiti switch case da utvrdimo da li je
generisani slučajni broj paran. Ako je paran, generirani broj se ispisuje i
for petlja se završava korištenjem labele. Funkcija Intn paketa rand se koristi
za generisanje nenegativnih pseudoslučajnih brojeva.
*/

func switchRand() {

	fmt.Println("\n --- switchRand ---")

randloop:
	for {
		switch i := rand.Intn(100); {
		case i%2 != 0:
			fmt.Printf("Generated odd number is %d\n", i)
			continue
		case i%2 == 0:
			fmt.Printf("Generated even number is %d", i)
			break randloop
		}
	}
}

/*
U gornjem programu, for petlja je označena randloopu labelom. Slučajni broj se
generira između 0 i 99 (100 nije uključeno) korištenjem Intn funkcije. Ako je
generisani broj neparan petlja ide u novu iteraciju, ako je petlja generisani
broj paran petlja se se prekida break randloop labele.

Imajte na umu da ako se naredba break koristi bez labele, naredba switch će se
prekinuti i petlja će nastaviti s izvršavanjem. Dakle, označavanje petlje i
njeno korištenje u naredbi break unutar naredbe switch je neophodno za prekid
spoljnje for petlje.
*/

func SwitchFunc() {

	fmt.Println("\n --- Switch ---")

	switchElem()
	switchDefault()
	switchMultiExpr()
	switchWithoutExpr()
	switchFallthrough()
	switchCaseFalseFallthrough()
	switchBreak()
	switchRand()
}
