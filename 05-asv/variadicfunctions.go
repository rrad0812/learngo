/*
Varijadičke funkcije
====================

Šta je varijadička funkcija?
----------------------------
Funkcije generalno prihvataju fiksni broj argumenata. Varijadička funkcija je
funkcija koja prihvata promenljiv broj argumenata. Ako poslednji parametar
definicije funkcije ima prefiks ..., onda funkcija može da prihvati bilo koji
broj argumenata za taj parametar.

Samo poslednji parametar funkcije može biti promenljiv. Naučićemo zašto je to
slučaj u sledećem odeljku ovog tutorijala.

Sintaksa
--------

func hello(a int, b ...int) {

}

U gornjoj funkciji, parametar b je promenljiv jer mu tip ...int prethodi sa
elipsom i može da prihvati bilo koji broj argumenata. Ova funkcija se može
pozvati korišćenjem sintakse:

hello(1, 2) //passing one argument "2" to b
hello(5, 6, 7, 8, 9) //passing arguments "6, 7, 8 and 9" to b

U gornjem delu koda, u prvom redu pozivamo hello sa jednim argumentom 2 za
parametar b, a u sledećem redu prosleđujemo četiri argumenta 6, 7, 8, 9
parametru b.

Takođe je moguće proslediti nula argumenata varijadičkoj funkciji.

hello(1)

U gornjem kodu, pozivamo hello bez argumenata za b. Ovo je sasvim u redu.

Do sada pretpostavljam da ste razumeli zašto varijabilni parametar treba da
bude na kraju.

Hajde da pokušamo da prvi parametar funkcije hello učinimo varijabilnim.

Sintaksa će izgledati ovako:

func hello(b ...int, a int) {

}

U gornjoj funkciji nije moguće proslediti argumente parametru a jer će bilo
koji argument koji prosledimo biti dodeljen prvom parametru b pošto je on
varijabilan. Stoga varijabilni parametri mogu biti prisutni samo na poslednjem
mestu u definiciji funkcije. Gornja funkcija neće uspeti da se kompajlira:

	>> syntax error: cannot use ... with non-final parameter b

Primeri i razumevanje kako funkcionišu varijadičke funkcije
-----------------------------------------------------------

Hajde da kreiramo sopstvenu varijabilnu funkciju. Napisaćemo jednostavan
program koji će utvrditi da li ceo broj postoji u ulaznoj listi celih
brojeva.
*/
package asv

import (
	"fmt"
)

func find(num int, nums ...int) {
	fmt.Printf("type of nums is %T\n", nums)
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			found = true
		}
	}
	if !found {
		fmt.Println(num, "not found in ", nums)
	}
	fmt.Printf("\n")
}

func varfuncsFind() {

	fmt.Println("\n --- Find ---")

	find(89, 89, 90, 95)
	find(45, 56, 67, 45, 90, 109)
	find(78, 38, 56, 98)
	find(87)
}

/*
U gornjem programu,

func find(num int, nums ...int)

prihvata promenljiv broj argumenata za nums parametar. Unutar funkcije find,
tip nums je []int, tj. celobrojni isečak.

Varijadičke funkcije rade tako što konvertuju promenljivi broj argumenata u
tip isečka tipa varijadičkog parametra. Na primer, u gornjem programu
promenljivi broj argumenata funkcije find je 89, 90, 95. Funkcija find očekuje
varijadički int argument. Stoga će kompajler konvertovati ova tri argumenta u
isečak tipa int - []int{89, 90, 95}, a zatim će ih proslediti funkciji find.

Nadalje, for petlja se kreće preko nums isečka i ispisuje poziciju num ako je
prisutan u isečku. Ako nije, ispisuje da broj nije pronađen.

Gore navedeni programski štampa:

	>> type of nums is []int
	>> 89 found at index 0 in [89 90 95]

	>> type of nums is []int
	>> 45 found at index 2 in [56 67 45 90 109]

	>> type of nums is []int
	>> 78 not found in  [38 56 98]

	>> type of nums is []int
	>> 87 not found in  []

Poslednji poziv funkcije find ima samo jedan argument. Nismo prosledili nijedan
argument parametru variadic nums ...int. Kao što je ranije rečeno, ovo je sasvim
legalno i u ovom slučaju, nums biće nil isečak dužine i kapaciteta 0.

Argumenti sa isečcima u odnosu na varijabilne argumente
--------------------------------------------------------
Definitivno bi trebalo da vam se sada mota po glavi jedno pitanje. U prethodnom
odeljku smo saznali da su varijadični argumenti funkcije zapravo konvertovani u
isečke. Zašto nam onda uopšte trebaju varijadičke funkcije kada istu
funkcionalnost možemo postići koristeći isečke?

Prepisao sam gornji program koristeći isečke.
*/

func findSlice(num int, nums []int) {

	fmt.Printf("type of nums is %T\n", nums)
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			found = true
		}
	}
	if !found {
		fmt.Println(num, "not found in ", nums)
	}
	fmt.Printf("\n")
}

func varfuncsFindSlice() {

	fmt.Println("\n --- Find slice---")

	findSlice(89, []int{89, 90, 95})
	findSlice(45, []int{56, 67, 45, 90, 109})
	findSlice(78, []int{38, 56, 98})
	findSlice(87, []int{})
}

/*
Slede prednosti korišćenja varijadičkih argumenata umesto isečaka.

    Nema potrebe za kreiranjem isečka tokom svakog poziva funkcije. Ako
	pogledate program iznad, kreirali smo nove isečke tokom svakog poziva
	funkcije findSlice. Ovo dodatno kreiranje isečka  se može izbeći
	korišćenjem varijadičnih funkcija.

    Kod poslenjeg poziva funkcije finSlice u gornjrm programu, kreiramo prazan
	isečak samo da bismo zadovoljili potpis funkcije findSlice. Ovo uopšte nije
	potrebno u slučaju varijadičnih funkcija. Ova linija koda može biti samo
	findSlice(87) kada se koristi varijadička funkcija.

	Lično smatram da je program sa varijadnim funkcijama čitljiviji nego
	onaj sa kriškama :)

Append je varijadička funkcija
------------------------------
Da li ste se ikada zapitali kako funkcija append iz standardne biblioteke
koja se koristi za dodavanje vrednosti na isečak prihvata bilo koji broj
argumenata? To je zato što je to varijadička funkcija.

func append(slice []Type, elems ...Type) []Type

Gore navedeno je definicija append funkcije. U ovoj definiciji elems je
varijabilni parametar. Stoga, append može prihvatiti promenljiv broj
argumenata.

Prenošenje sloja varijabilnoj funkciji
--------------------------------------
Hajde da prosledimo isečak varijadičkoj funkciji i saznamo šta se dešava iz
donjeg primera.
*/

// func findSlice2(num int, nums ...int) {

// 	fmt.Printf("type of nums is %T\n", nums)
// 	found := false
// 	for i, v := range nums {
// 		if v == num {
// 			fmt.Println(num, "found at index", i, "in", nums)
// 			found = true
// 		}
// 	}
// 	if !found {
// 		fmt.Println(num, "not found in ", nums)
// 	}
// 	fmt.Printf("\n")
// }

// func varfuncsFindSlice2() {

// 	fmt.Println("\n --- Find slice 2---")

// 	nums := []int{89, 90, 95}
// 	findSlice2(89, nums)
// }

/*
U gornjem programu, prosleđujemo isečak funkciji koja očekuje promenljiv broj
argumenata.

Ovo neće raditi. Gore navedeni program će se otkazati zbog greške u kompilaciji

>> ../prog.go:23:10: cannot use nums (type []int) as type int in argument to find

Zašto ovo ne funkcioniše? Pa, prilično je jednostavno. Potpis funkcije find je dat
ispod,

func findSlice2(num int, nums ...int)

Prema definiciji varijadičke funkcije findSlice2 to znači da će ona prihvatiti
promenljiv broj argumenata tipa int.

U gornjem programu, slice nums koji je tipa []int se prosleđuje funkciji
findSlice2 koja očekuje varijadički int argument. Kao što smo već pomenuli, ovi
varijadički argumenti će biti konvertovani u tip int isečka, int jer findSlice2
očekuje varijadičke celobrojne argumente. U ovom slučaju, nums je već je []int
isečak i kompajler pokušava da kreira novi []int, tj. kompajler pokušava da uradi

findSlice2(89, []int{nums})

što će propasti jer nums je []int a ne int.

Dakle, postoji li način da se prosledi isečak varijabilnoj funkciji? Odgovor
je da.

Postoji sintaksički šećer koji se može koristiti za prosleđivanje isečka
varijadičkoj funkciji. Morate ispred isečka dodati elipsu ... .Ako se to uradi,
isečak se direktno prosleđuje funkciji bez kreiranja novog isečka.

U gornjem programu, ako zamenite findSlice2(89, nums) sa findSlice2(89, nums...),
program će se kompajlirati i ispisati sledeći izlaz:

	>> type of nums is []int
	>> 89 found at index 0 in [89 90 95]

Evo kompletnog programa za vašu referencu.
*/

func findSlice3(num int, nums ...int) {

	fmt.Printf("type of nums is %T\n", nums)
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			found = true
		}
	}
	if !found {
		fmt.Println(num, "not found in ", nums)
	}
	fmt.Printf("\n")
}

func varfuncsFindSlice3() {

	fmt.Println("\n --- Find slice 3---")

	nums := []int{89, 90, 95}
	findSlice3(89, nums...)
}

/*
Modifikovanje isečka unutar varijadičke funkcije
------------------------------------------------
Samo budite sigurni da znate šta radite kada modifikujete isečak unutar
varijadičke funkcije.

Pogledajmo jednostavan primer.
*/

func change(s ...string) {
	s[0] = "Go"
}

func varfuncsSliceElipsis() {

	fmt.Println("\n --- Slice elipsis ---")

	welcome := []string{"hello", "world"}
	change(welcome...)
	fmt.Println(welcome)
}

/*
Šta mislite da će biti izlaz gornjeg programa? Ako mislite da će biti...
[Go world] Čestitam! Razumeli ste varijadičke funkcije i isečke. Ako ste
pogrešili, nije strašno, dozvolite mi da objasnim kako dobijamo ovaj izlaz.

U gornjem programu, koristimo sintaksički šećer ... i prosleđujemo isečak
kao varijabilni argument funkciji change.

Kao što smo već razgovarali, ako se koristi  ..., welcome isečak će biti
prosleđen kao argument bez kreiranja novog isečka. Stoga welcome će biti
prosleđen funkciji change kao argument.

Unutar funkcije change, prvi element segmenta se menja u Go. Stoga, ovaj
program ispisuje:

[Go world]

Evo još jednog programa za razumevanje varijadičkih funkcija.
*/

func change2(s ...string) {
	s[0] = "Go"
	s = append(s, "playground")
	fmt.Println(s)
}

func varfuncsSliceElipsis2() {

	fmt.Println("\n --- Slice elipsis 2 ---")

	welcome := []string{"hello", "world"}
	change2(welcome...)
	fmt.Println(welcome)
}

func VarFuncs() {

	fmt.Println("\n --- Varriadic functions ---")

	varfuncsFind()
	varfuncsFindSlice()
	// varfuncsFindSlice2()
	varfuncsFindSlice3()
	varfuncsSliceElipsis()
	varfuncsSliceElipsis2()
}
