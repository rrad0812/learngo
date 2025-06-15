/*
Pointeri
========

Šta je pointer?
---------------
Pointer je promenljiva koja čuva memorijsku adresu druge promenljive.

Pointeri u programu Go
----------------------

	           a                  b
		+---------------+      +------+
		|   0x1040a124  | ---> | 156  |
	    +---------------+      +------+

Na gornjoj ilustraciji, promenljiva b ima vrednost 156 i čuva se na memorijskoj
adresi 0x1040a124. Promenljiva a sadrži adresu b. Kaže se da je a pointer koji
pokazuje na promenljivu b. U ranijim danima, ovo se nazivalo i indirektnim
adresiranjem vrednosti.

Deklarisanje pointera
---------------------
*T je tip pointerske promenljive koja pokazuje na vrednost tipa T.

Hajde da napišemo program koji deklariše pointer.
*/
package pointers

import (
	"fmt"
)

func pointerDecl() {

	fmt.Println("\n ---Pointer declaration ---")

	b := 255
	a := &b

	fmt.Printf("Type of a is %T\n", a)
	fmt.Println("Address of b is", a)
}

/*
Operator & se koristi za dobijanje adrese promenljive. U gornjem programu
dodeljujemo adresu promenljive b čiji je tip int, pointerskoj promenljivoj a
koja je tipa *int. Sada se kaže da a pokazuje na b. Kada ispišemo vrednost,
biće ispisana adresa promenljive b. Ovaj program štampa:

	>> Type of a is *int
	>> Address of b is 0x1040a124

Možda ćete dobiti drugačiju adresu za b jer lokacija b može biti bilo gde u
memoriji.

Nulta vrednost pointera
-----------------------
Nulta vrednost pointera je nil.
*/

func pointerNil() {

	fmt.Println("\n ---Pointer nil ---")

	a := 25

	var b *int // b is not initialized, so it is nil
	fmt.Println("b is", b)

	b = &a // b is now pointing to a
	fmt.Println("b after initialization is", b)
}

/*
b je inicijalno ništa u gornjem programu, a kasnije se dodeljuje adresi a. Ovaj
program daje rezultat:

	>> b is <nil>
	>> b after initialisation is 0x1040a124

Kreiranje pointera pomoću funkcije new
--------------------------------------
Go takođe pruža praktičnu funkciju new za kreiranje pointera nekog tipa. new
funkcija uzima tip kao argument i vraća pointer na novo alociranu nultu vrednost
tipa koji je prosleđen kao argument.

Sledeći primer će stvari učiniti jasnijim.
*/

func pointerNew() {

	fmt.Println("\n ---Pointer new ---")

	i := new(int) // i is a pointer to int
	fmt.Printf("Value of i is %d, type is %T, address is %v\n", *i, i, i)

	*i = 85 // Dereferencing the pointer i to assign a value
	fmt.Printf("Value of i is %d, type is %T, address is %v\n", *i, i, i)
}

/*
U gornjem programu, koristimo funkciju new za kreiranje pointera tipa int. Ova
funkcija će vratiti pointer na novo alociranu nultu vrednost tipa int. Nulta
vrednost tipa int je 0. Dakle, i će biti tipa *int i pokazivaće na 0, tj. *i će
biti 0.

Gore navedeni program će ispisati:
	>> Value of i is 0, type is *int, address is 0x414020
	>> Value of i is 85, type is *int, address is 0x414020

Dereferenciranje pointera
-------------------------
Dereferenciranje pointera znači pristupanje vrednosti promenljive na koju
pointer pokazuje. *a je sintaksa za deferenciranje pointerske promenljive a.

Da vidimo kako ovo funkcioniše u programu.
*/

func pointerDeref() {

	fmt.Println("\n ---Pointer dereferencing ---")

	// Declare a pointer to an int
	var a *int
	fmt.Println("Pointer a is", a)

	// Assign a value to the pointer
	b := 255
	a = &b
	fmt.Println("Address of b is", a)
	fmt.Println("Value of b is", *a)
}

/*
U gornjem programu, sa pointerom a referenciramo na b i ispisujemo vrednost.
Kao što se i očekivalo, ispisuje se vrednost b. Izlaz programa je:

	>> Address of b is 0x1040a124
	>> Value of b is 255

Hajde da napišemo još jedan program gde menjamo vrednost promenljivoj b
koristeći pointer.
*/

func pointerDeref2() {

	fmt.Println("\n ---Pointer dereferencing 2 ---")

	b := 255
	fmt.Printf("Value of b is %v, type is %T, address of b is %v\n", b, b, &b)

	a := &b
	*a++ // Incrementing the value pointed by a. This is statement not expression
	fmt.Printf("Value of b is %v, type is %T, address of b is %v\n", *a, *a, a)
}

/*
U gornjem programu, povećavamo vrednost na koju pokazuje a za 1, što menja
vrednost b pošto `a` pokazuje na `b`. Stoga vrednost `b` postaje 256.
Izlaz programa je:

	>> Value of b is 255, type is int, address of b is 0xc000098060
	>> Value of b is 256, type is int, address of b is 0xc000098060

Prenošenje pointera funkciji
----------------------------
*/

func change(val *int) {
	*val = 55
}

func pointerPassFunc() {

	fmt.Println("\n ---Pointer passing function ---")

	a := 58
	fmt.Println("Value of a before function call is", a)

	b := &a
	change(b) // Passing pointer b to change function
	fmt.Println("Value of a after function call is", a)
}

/*
U gornjem programu prosleđujemo pointersku promenljivu b koja sadrži adresu
promenljive a funkciji change. Unutar change funkcije, vrednost promenljive a
se menja korišćenjem dereferenciranja prosleđenog pointera. Ovaj program štampa,

	>> value of a before function call is 58
	>> value of a after function call is 55

Vraćanje pointera iz funkcije
-----------------------------
Sasvim je legalno da funkcija vrati pointer lokalne promenljive iz funkcije.
Go kompajler je dovoljno inteligentan i alociraće ovu promenljivu na hipu.
*/

func hello() *int {
	i := 5
	return &i
}

func pointerReturnFunc() {

	fmt.Println("\n ---Pointer returning function ---")

	d := hello()
	fmt.Println("Value of d is", *d)
}

/*
U gornjem programu, vraćamo adresu lokalne promenljive "i" iz funkcije hello.
Ponašanje ovog koda je nedefinisano u programskim jezicima kao što su C i C++,
jer promenljiva izlazi van opsega važenja kada se funkcija hello vrati. Ali u
slučaju Go-a, kompajler vrši analizu izlaska i alocira "i" na hipu kada se
adresa vrati iz lokalnog opsega važenja. Stoga će ovaj program raditi i
ispisaće,

	>> Value of d is 5

!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
!!! Ne prosleđujte pointer na niz kao argument funkcije.                    !!!
!!! Umesto toga koristite slice. 						                    !!!
!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

Pretpostavimo da želimo da napravimo neke izmene na nizu unutar funkcije i da
promene napravljene na tom nizu unutar funkcije treba da budu vidljive
pozivaocu. Jedan od načina da se to uradi jeste da se funkciji prosledi
pointer na niz kao argument.
*/

func modifyArray(arr *[3]int) {
	(*arr)[0] = 90
}

func pointerArray() {

	fmt.Println("\n ---Pointer array ---")

	a := [3]int{89, 90, 91}
	fmt.Println(a)

	modifyArray(&a)
	fmt.Println(a)
}

/*
U gornjem programu, prosleđujemo adresu niza "a" funkciji modifyArray. U
modifyArray funkciji dereferenciramo arr i dodeljujemo vrednost 90 prvom
elementu niza. Ovaj program ispisuje

	>> [90 90 91]

a[x] je skraćenica za (*a)[x]. Dakle, (*arr)[0] u gornjem programu može se
zameniti sa arr[0]. Hajde da prepišemo gornji program koristeći ovu skraćenu
sintaksu.
*/

func modifyArray2(arr *[3]int) {
	arr[0] = 90 // shortened syntax of (*arr)[0]
}

func pointerArray2() {

	fmt.Println("\n ---Pointer array 2---")

	a := [3]int{89, 90, 91}
	fmt.Println(a)

	modifyArray2(&a)
	fmt.Println(a)
}

/*
Ovaj program takođe ispisuje

	>> [90 90 91]

Iako ovaj način prenošenja pointera na niz kao argument funkcije i pravljenja
njegovih izmena unutar funkcije funkcioniše, to nije idiomatski način da se ovo
postigne u Go-u. Za ovo imamo isečke (slices).

Hajde da prepišemo isti program koristeći isečke.
*/

func modifyArray3(sls []int) {
	sls[0] = 90
}

func pointerArray3() {

	fmt.Println("\n ---Pointer array 3---")

	a := [3]int{89, 90, 91}
	fmt.Println(a)

	modifyArray3(a[:])
	fmt.Println(a)
}

/*
U gornjem programu, prosleđujemo isečak funkciji modifyArray3. Prvi element
isečka se menja u 90 unutar modifyArray3 funkcije. Ovaj program takođe ispisuje

>> [90 90 91].

Zato zaboravite na prosleđivanje pointera na nizove i umesto toga koristite
isečke :) . Ovaj kod je mnogo čistiji i idiomatski je u Go :).

Go ne podržava pointerku aritmetiku
-----------------------------------
Go ne podržava pointersku aritmetiku koja je prisutna u drugim jezicima poput
C i C++.
*/

// func main() {
// 	   b := [...]int{109, 110, 111}
// 	   p := &b
// 	   p++
// }

/*
Gore navedeni program će izbaciti grešku kompilacije

	>> * main.go:6: nevažeća operacija: p++ (nenumerički tip [3]int)
*/

func PointerFuncs() {

	fmt.Println("\n ---Pointers in Go ---")

	pointerDecl()
	pointerNil()
	pointerNew()
	pointerDeref()
	pointerDeref2()
	pointerPassFunc()
	pointerReturnFunc()
	pointerArray()
	pointerArray2()
	pointerArray3()
}
