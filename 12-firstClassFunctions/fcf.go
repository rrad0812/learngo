/*
Funkcije prve klase
====================

Jezik koji podržava funkcije prve klase omogućava da se funkcije dodeljuju
promenljivim, prosleđuju kao argumenti drugim funkcijama i vraćaju iz drugih
funkcija. Go ima podršku za funkcije prve klase.

Anonimne funkcije
-----------------
Počnimo sa jednostavnim primerom koji dodeljuje funkciju promenljivoj.
*/
package fcf

import (
	"fmt"
)

func fcfAnonimousF() {
	fmt.Println("\n --- Anonimous functions ---")

	a := func() { // Decl var a and assign func literal
		fmt.Println("hello world first class function")
	}

	a()                                    // Call a() func
	fmt.Printf("Type of var a is %T\n", a) // Type of var a
}

/*
U gornjem programu, dodelili smo funkciju promenljivoj. Ovo je sintaksa za
dodeljivanje funkcije promenljivoj. Ako pažljivo primetite, funkcija koja je
dodeljena nema ime. Ovakve funkcije se nazivaju anonimne funkcije.

Jedini način da se pozove ova funkcija je korišćenjem promenljive a. To smo
uradili u sledećem redu. a() poziva funkciju, potom ispisujemo tip a promenljive.
Ovo će ispisati:

	>> hello world first class function
	>> Type of var a is func()

Takođe je moguće pozvati anonimnu funkciju bez njenog dodeljivanja promenljivoj.
Pogledajmo kako se to radi u sledećem primeru.
*/
func fcfAnonimousF2() {
	fmt.Println("\n --- Anonimous function without var ---")
	func() {
		fmt.Println("hello world first class function")
	}()
}

/*
U gornjem programu, anonimna funkcija je definisana i odmah nakon definicije
funkcije pozivana, (). Ovaj program će ispisati:

	>> hello world first class function

Takođe je moguće proslediti argumente anonimnim funkcijama baš kao i bilo kojoj
drugoj funkciji.
*/

func fcfAnonimousF3() {

	fmt.Println("\n --- Anonimous function with parameters ---")

	func(n string) {
		fmt.Println("Welcome", n)
	}("Gophers")
}

/*
U gornjem programu, string argument se prosleđuje anonimnoj funkciji. Pokretanje
ovog programa će ispisati:

	>> Welcome Gophers

Korisnički definisani tipovi funkcija
-------------------------------------
Baš kao što definišemo sopstvene tipove struktura, moguće je definisati
sopstvene tipove funkcija.

	type add func(a int, b int) int

Isečak koda iznad kreira novi tip funkcije add koji prihvata dva cela argumenta
i vraća ceo broj. Sada možemo definisati promenljive tipa add.

Hajde da napišemo program koji definiše promenljivu tipa add.
*/

type add func(a int, b int) int

func fcfUserTypesOfF() {
	fmt.Println("\n --- User def types of functions ---")

	var a add = func(a int, b int) int {
		return a + b
	}

	s := a(5, 6)
	fmt.Println("Sum", s)
}

/*
U gornjem programu definišemo promenljivu a tipa add i dodeljujemo joj funkciju
čiji potpis odgovara tipu add. Pozivamo funkciju u liniji br. 13 i dodeljujemo
rezultat promenljivoj s. Ovaj program će ispisati:

	>> Sum 11

Funkcije višeg reda
-------------------

Definicija funkcije višeg reda sa wiki-ja je funkcija koja obavlja barem jednu
od sledećih stvari:

- prihvata jednu ili više funkcija kao argumente
- vraća funkciju kao rezultat

Pogledajmo nekoliko jednostavnih primera za gornja dva scenarija.

Prenošenje funkcija kao argumenata drugim funkcijama
----------------------------------------------------
*/

func simple(a, b int, f func(a, b int) int) {
	fmt.Println(f(a, b))
}

func fcfPassFuncArg() {

	fmt.Println("\n --- Pass func arguments to function ---")

	a := 60
	b := 7

	sum := func(a, b int) int {
		return a + b
	}

	diff := func(a, b int) int {
		return a - b
	}

	mul := func(a, b int) int {
		return a * b
	}

	div := func(a, b int) int {
		return a / b
	}

	mod := func(a, b int) int {
		return a % b
	}

	simple(a, b, sum)
	simple(a, b, diff)
	simple(a, b, mul)
	simple(a, b, div)
	simple(a, b, mod)
}

/*
U gornjem primeru definišemo funkciju "simple" koja prihvata dva celobrojna
argumenta i vraća ceo broj. Unutar funkcije glavne funkcije kreiramo anonimnu
funkciju f1 čiji se potpis podudara sa func. parametrom funkcije simple.
Pozivamo simple i prosleđujemo sum kao argument u sledećem redu. Isto to radimo
sa funkcijama diff, mul, div i mod.

Ovaj program štampa:

	>> 67
	>> 53
	>> 420
	>> 8
	>> 4

kao izlaz.

Vraćanje funkcija iz funkcija
------------------------------------
Sada hajde da prepišemo gornji program i vratimo funkciju iz simple funkcije.
*/

func simpleRetF() func(a, b int) int {
	f := func(a, b int) int {
		return a + b
	}
	return f
}

func fcfFuncRetF() {
	fmt.Println("\n --- Funnction return function ---")
	s := simpleRetF()
	fmt.Println(s(60, 7))
}

/*
U gornjem programu, "simpleRetF" funkcija vraća funkciju koja prihvata dva int
argumenta i vraća jedan int argument.

Ova funkcija se poziva u glavnoj funkciji. Povratna vrednost iz te funkcije se
dodeljuje promenljivoj s. Sada s sadrži funkciju koju vraća simpleRetF funkcija.
Pozivamo s i prosleđujemo joj dva celobrojna argumenta. Ovaj program štampa

	>> 67

Zatvaranja
----------
Zatvaranja su poseban slučaj anonimnih funkcija. Zatvaranja su anonimne funkcije
koje pristupaju promenljivim definisanim izvan tela funkcije. Primer će stvari
učiniti jasnijim.
*/

func fcfClosure() {

	fmt.Println("\n --- Closures ---")

	a := 5
	func() {
		fmt.Println("a =", a)
	}()
}

/*
U gornjem programu, anonimna funkcija pristupa promenljivoj a koja se nalazi
izvan njenog tela. Stoga je ova anonimna funkcija zatvaranje.

Svako zatvaranje je vezano za sopstvenu okolnu promenljivu. Hajde da razumemo
šta to znači koristeći jednostavan primer.
*/

func appendStr() func(string) string {
	t := "Hello"
	c := func(b string) string {
		t = t + " " + b
		return t
	}
	return c
}

func fcfClosures2() {

	fmt.Println("\n --- Closures full example  ---")

	a := appendStr()
	b := appendStr()

	fmt.Println(a("World"))
	fmt.Println(b("Everyone"))

	fmt.Println(a("Gopher"))
	fmt.Println(b("!"))
}

/*
U gornjem programu, funkcija appendStr vraća zatvaranje. Ovo zatvaranje je
anonimna funkcija plus njen enviroment od značaja za nju. U našem slučaju to je
funkcija koja uzima string kao parametar i vraća string. Pored toga njoj je
na raspolaganju promenljiva t.

Promenljive a i b deklarisane kao zatvarači i vezane su za vrednost promenljive
t.

Ovaj program će štampati,

	>> Hello World
	>> Hello Everyone
	>> Hello World Gopher
	>> Hello Everyone !

Praktična upotreba funkcija prve klase
--------------------------------------
Do sada smo definisali šta su funkcije prve klase i videli smo nekoliko
osmišljenih primera da bismo naučili kako funkcionišu. Sada hajde da napišemo
konkretan program koji prikazuje praktičnu upotrebu funkcija prve klase.

Napravićemo program koji filtrira isečak učenika na osnovu nekih kriterijuma.
Hajde da pristupimo ovome korak po korak.

Prvo, hajde da definišemo tip studenta.

	type student struct {
		firstName string
		lastName string
		grade string
		country string
	}

Sledeći korak je pisanje filter funkcije. Ova funkcija uzima kao parametre
isečak učenika i funkciju koja određuje da li učenik ispunjava kriterijume
filtracije. Bolje ćemo razumeti kada napišemo ovu funkciju. Hajde da to uradimo.

	func filter(s []student, f func(student) bool) []student {
		var r []student
		for _, v := range s {
			if f(v) == true {
				r = append(r, v)
			}
		}
		return r
	}

U gornjoj filter funkciji, drugi parametar f je funkcija koja uzima studenta
kao parametar i vraća bool. Ova funkcija određuje da li određeni student
ispunjava kriterijum ili ne.

Iteriramo kroz isečak studenata i prosleđujemo svakog studenta kao parametar
funkciji f. Ako funkcija vrati true, to znači da je student ispunio kriterijume
filtera i da je dodat u isečak r.

Možda ste malo zbunjeni oko stvarne upotrebe ove funkcije, ali biće jasno kada
završimo program. Dodao sam glavnu funkciju i dao sam kompletan program ispod.
*/

type student struct {
	firstName string
	lastName  string
	grade     string
	country   string
}

func filter(s []student, f func(student) bool) []student {
	var r []student
	for _, v := range s {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}

func fcfFilterFunc() {
	fmt.Println("\n --- Filter function ---")

	s1 := student{
		firstName: "Naveen",
		lastName:  "Ramanathan",
		grade:     "A",
		country:   "India",
	}

	s2 := student{
		firstName: "Samuel",
		lastName:  "Johnson",
		grade:     "B",
		country:   "USA",
	}

	s := []student{s1, s2}

	f := filter(s, func(s student) bool {
		return s.grade == "B"
	})

	c := filter(s, func(s student) bool {
		return s.country == "India"
	})

	fmt.Println(f)
	fmt.Println(c)
}

/*
U glavnoj funkciji, prvo kreiramo dva studenta s1 i s2 i dodajemo ih u slice s.
Sada, recimo da želimo da saznamo koji su svi studenti koji imaju ocenu B.

Ovo smo utvrdili u gornjem programu tako što smo prosledili funkciju koja
proverava da li student ima ocenu B i ako ima, vraća vrednost true, kao
funkciji filter. Gornji program će ispisati,

	>> [{Samuel Johnson B USA}]

Recimo da želimo da pronađemo sve studente iz Indije. To se može lako uraditi
promenom parametra funkcije na funkciju filtera. U nastavku sam naveo kod koji
to radi,

	c := filter(s, func(s student) bool {
	    return s.country == "India"
	})
	fmt.Println(c)

Molim vas, dodajte ovo glavnoj funkciji i proverite izlaz.

Zaključimo ovaj odeljak pisanjem još jednog programa. Ovaj program će izvršiti
iste operacije na svakom elementu isečka i vratiti rezultat. Na primer, ako
želimo da pomnožimo sve cele brojeve u isečku sa 5 i vratimo izlaz, to se lako
može uraditi korišćenjem funkcija prve klase. Ovakve funkcije koje operišu na
svakom elementu kolekcije nazivaju se map funkcije. Program sam naveo u nastavku.
Kod je sam po sebi jasan.
*/

func iMap(s []int, f func(int) int) []int {
	var r []int
	for _, v := range s {
		r = append(r, f(v))
	}
	return r
}

func fcfMapFunc() {

	fmt.Println("\n --- Map function---")

	i := []int{5, 6, 7, 8, 9}

	r := iMap(i, func(n int) int {
		return n * 5
	})

	fmt.Println(r)
}

/*
Gore navedeni program će ispisati:

	>> [25 30 35 40 45]
*/
func FcfFunc() {
	fmt.Println("\n --- First class functions---")

	fcfAnonimousF()
	fcfAnonimousF2()
	fcfAnonimousF3()
	fcfUserTypesOfF()
	fcfPassFuncArg()
	fcfFuncRetF()
	fcfClosure()
	fcfClosures2()
	fcfFilterFunc()
	fcfMapFunc()
}
