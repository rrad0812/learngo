/*
Promenljive
===========

Ovo je treći vodič našoj Tutorial Golang seriji i bavi se promenljivima
u Golangu.

Šta je promenljiva?
-------------------

Promenljivo je ime dato memorijskoj lokaciji za čuvanje vrednosti specifičnog
tipa.Postoje razne sintakse za deklaracija promenljivih u Gou.Pogledajmo ih
jednu po jednu.

Deklaracija jedne promenljive
-----------------------------

	var name type
*/
package variables

import (
	"fmt"
	"math"
)

func singleVarDecl() {

	fmt.Println("\n --- Single var declaration ---")
	var age int // variable declaration
	fmt.Println("My initial age is", age)
}

/*
Izjava var age int deklariše promenljivu age tipa int.Promenljivoj nije
dodeljena vrednost.Ako promenljivoj nije dodeljena vrednost, Go automatski
inicijalizira sa nultom vrednosti promenljive za dati tip.U ovom slučaju,
promenljivoj age je dodeljena vrednost 0, koja je nulta vrednost tipa int.
Ako pokrenete ovaj program, možete videti sledeći izlaz.

>> My initial age is 0

Deklaracija i doela vrednosti
-----------------------------

Varijabli se može dodeliti bilo kojoj vrednosti njenog tipa.U gornjem programu,
promeljivoj age se može dodeliti bilo koju celobrojna vrednost.
*/
func singleVarDeclAndAssigment() {

	fmt.Println("\n --- Single var declaration and assigment---")
	var age int // variable declaration
	fmt.Println("My initial age is", age)

	age = 29 //assignment
	fmt.Println("My age after first assignment is", age)

	age = 54 //assignment
	fmt.Println("My age after second assignment is", age)
}

/*
Gornji program će odštampati sledeći izlaz:

>> My initial age is 0
>> My age after first assignment is 29
>> My age after second assignment is 54

Deklaracija promenljive sa početnom vrednosti
-----------------------------------------------

Promenljiva se takođe može inicijatizovati sa vrednosti kada je proglašena.
Sledi sintaksa koja proglašava promenljivu sa početnom vrednošću.

	var name type = initialvalue
*/
func singleVarDeclAndInit() {

	fmt.Println("\n --- Single var declaration and init ---")
	var age int = 29 // variable declaration with initial value
	fmt.Println("My initial age is", age)
}

/*
U gornjem programu, promenljiva age je tipa int i ima inicijalnu
vrednost 29. Gornji program će odštampati sledeći izlaz.

>> My initial age is 29

Potvrđuje da je starost inicijalizirana sa vrednosti 29.

Deklaracija sa zaključivanjem tipa
------------------------------------

Ako promenljiva ima početnu vrednost, Go će automatski da se zaključi
tip te promenljive koristeći njenu početnu vrednost.Otuda ako promenljiva
ima početna vrednost, deklaracija tipa promenljive može se izostaviti.

Ako se promenljiva proglašava pomoću sledeće sintakse

	var name = initialvalue

Go će automatski zaključiti tip te promenljive na osnovu početne vrednosti.

U sledećem primeru možemo videti da je tip int promenljive age uklonjen u redu
br.3 Pošto promenljiva age ima početnu vrednost 29, Go će zaključiti da je tip
promenljive int.
*/
func singleVarDelcAndInfered() {

	fmt.Println("\n --- Single var declaration, init and infered ---")
	var age = 29 // type will be inferred - int
	fmt.Println("My initial age is", age)
}

/*
Višestruka deklaracija promenljivih
-----------------------------------

Više promenljivih se može deklarisati pomoću jedinstvene izjave.

	var name1, name2 type = initialvalue1, initialvalue2

je sintaksa za više deklaracija promenljivih u jednom izjavi.
*/
func multipleVarDeclAndInit() {

	fmt.Println("\n --- Multiple var declaration and init ---")
	var price, quantity int = 5000, 100 //declaring multiple variables
	fmt.Println("price is", price, "quantity is", quantity)
}

/*
Deklaracija više promenljivh zaklučivanjem tipa
-----------------------------------------------

Tip se može ukloniti ako promenljive imaju početnu vrednost.
U programu iznad promenljive imaju početne vrednosti, tip int se može ukloniti.
*/
func multipleVarDeclAndInfered() {

	fmt.Println("\n --- Multiple var declaration and infered ---")
	var price, quantity = 5000, 100 //declaring multiple variables with type inference
	fmt.Println("price is", price, "quantity is", quantity)
}

/*
Gornji program će štampati

>> price is 5000 quantity is 100

Podrazumevana - nulta vrednost za više promenljivih
---------------------------------------------------

Kao što ste verovatno pretpostavili, ako početna vrednost promenljive nije
navedena imaće dodeljenu vrednost  0, koja je nulta vrednost za tip int u Gou.
*/
func multipleVarDecl() {

	fmt.Println("\n --- Multiple var declaration and assigment ---")
	var price, quantity int
	fmt.Println("price is", price, "quantity is", quantity)

	price = 3000
	quantity = 500
	fmt.Println("new price is", price, "new quantity is", quantity)
}

/*
Gornji program će štampati

>> price is 0 quantity is 0
>> new price is 3000 new quantity is 500

Grupna deklaracija više promenljivih različitog tipa
----------------------------------------------------

Možda će biti slučajeva u kojima bismo želeli da deklarišemo i inicijalizujemo
promenljive koje pripadaju različitim tipovima u jednoj izjavi. Sintaksa za to je

	var (
	      name1 = initialvalue1
	      name2 = initialvalue2
	)

Sledeći program koristi gornju sintaksu za deklaraciju promenljivih različitih tipova.
*/
func multipleVarGroupDecl() {

	fmt.Println("\n --- Multiple var group declaration and infered ---")
	var (
		name   = "Naveen"
		age    = 38
		height int
	)

	fmt.Println("my name is", name)
	fmt.Println("my age is", age)
	fmt.Println("my height is", height)
}

/*
Ovde proglašavamo promenljivu name tipa string, age i height tipa int.
(Razgovaraćemo o različitim tipovima dostupnim u Golang-u u sledećem tutorialu).

# pokretanje gornjeg programa će štampati

>> my name is Naveen
>> my age is 38
>> my height is 0

Kratka deklaracija
------------------

Go pruža još jedan sažet način da deklariše promenljive.Ovo je poznato kao kratka
deklaracija i koristi walrus := operater.

	name := initialvalue

je sintaksa kratkih ruku da proglasi promenljivu.

Sledeći program koristi sintaksu kratke ruke za deklaraciju i inicijalizaciju
promenljive count sa celobrojnom vrednosti 10. Go će automatski zaključiti da
je count tipa int.
*/
func shorthandVarDeclInferAndInit() {

	fmt.Println("\n --- Shorthand var declaration ---")
	count := 10
	fmt.Println("Count =", count)
}

/*
Gornji program će štampati,

>> Count = 10

Takođe je moguće deklarisati više promenljivih u jednoj liniji koristeći sintaksu
kratake deklaracije.
*/
func shorthandMultipleVarInOneLine() {

	fmt.Println("\n --- Shorthand multiple var declaration ---")
	name, age := "Naveen", 29 //short hand declaration
	fmt.Println("my name is", name)
	fmt.Println("my age is", age)
}

/*
Gornji program deklariše dva varijable name i age tipa string i int redom.
# Ako pokrenete gornji program, možete videti

>> my name is Naveen
>> my age is 29

odštampano.

Kratka deklaracija zahteva početne vrednosti za sve promenljive na levoj strani
strana znaka walrus operatora.

Sve varijable deklarisane sa walrus operatorom moraju imati početne vrednosti
-----------------------------------------------------------------------------

Sledeći program će ispisati grešaku je jednoj od dve promenljive nije dodeljena

	vrednost.
*/
func shorthandMultipleVar() {

	fmt.Println("\n --- Shorthand multiple var declaration and infered ---")
	name, age := "Naveen", 29 // if age is not init thi is error
	fmt.Println("my name is", name, "age is", age)
}

/*
Bar jedna promenljiva mora biti novo deklarisana
------------------------------------------------

Sintaksa kratke ruke može se koristiti samo kada je bar jedna od promenljivih
na levoj strani warus operatora (:=) novodeklarisana.
Razmotrimo sledeći program:
*/
func shorthandOnlyOneVarDec() {

	fmt.Println("\n --- Mininimum One Var Must be New Decl in shorthand decl ---")
	a, b := 20, 30 // declare variables a and b
	fmt.Println("a is", a, "b is", b)

	b, c := 40, 50 // b is already declared but c is new
	fmt.Println("b is", b, "c is", c)

	b, c = 80, 90 // assign new values to already declared variables b and c
	fmt.Println("changed b is", b, "c is", c)
}

/*
U gornjem programu, na liniji br.7, promenljiva b je već dekarisana, ali c
promenljiva je je novodeklarisana i otuda je sve OK i štampa

>> a is 20 b is 30
>> b is 40 c is 50
>> changed b is 80 c is 90

Kratka deklaracija ne može se izvršiti ponovljena
-------------------------------------------------

Dok ako pokrenemo program ispod
*/
func shorthandVarDuplError() {

	fmt.Println("\n --- Shorthand decl cant be duplicate ---")
	a, b := 20, 30 //a and b declared
	fmt.Println("a is", a, "b is", b)

	// a, b := 40, 50 //error, no new variables
}

/*
Greška će biti ispisana ./prog.go: 8:7: Nema novih promenljivih na levoj
strani: =. Ovo je jer su i promenljive a i b već deklarisane i nema nove
promenljive na levoj strani: = u redu br.6.

Runtime dodela vrednosti promenljivoj
-------------------------------------

Promenljivima se mogu dodeliti vrednosti koje se računaju tokom runtime-a.
Razmotrimo sledeći program,
*/
func shorthandRuntimeVarEval() {

	fmt.Println("\n --- Shorthand decl runtime var eval ---")
	a, b := 145.8, 543.8
	c := math.Min(a, b)
	fmt.Println("Minimum value is", c)
}

/*
Za gotnji program potreban je math paket i min funkcija u tom paketu.
Ne brinite o tome, razgovaraćemo detaljno i o paketima i o funkcijama u predstojećim
tutorijalima. Sve što trebamo znati je da se vrednost C izračunava u runtime-u
i to je minimum vrednosti a i b.Program iznad će odštampati

>> Minimum value is  145.8

Ne možete da promenite tip deklarisane promenljive
--------------------------------------------------

Pošto je Go jako tipizirani jezik, vrednost promenljive deklarisane kao
pripadnice jednog tipa ne može biti dodeljen kao vrednost promenljivoj
deklarisanoj kao pripadnici drugog tipa.

Sledeći program će ispisati grešku Ne možete koristite "Naveen" (string konstanta)
kao int vrednost u zadatku, jer je age proglašeno kao tip Int a mi pokušavamo da
joj dodelimo vrednost tipa string niza.
*/
func dontChangeTypeOfVar() {

	fmt.Println("\n --- cant change type of var declaration ---")
	age := 29 // age is int
	//age = "Naveen"    // error since we are trying to assign a string to
	// a variable of type int
	fmt.Println(age)
}

func Variables() {
	singleVarDecl()
	singleVarDeclAndAssigment()
	singleVarDeclAndInit()
	singleVarDelcAndInfered()
	multipleVarDeclAndInit()
	multipleVarDeclAndInfered()
	multipleVarDecl()
	multipleVarGroupDecl()
	shorthandVarDeclInferAndInit()
	shorthandMultipleVarInOneLine()
	shorthandMultipleVar()
	shorthandOnlyOneVarDec()
	shorthandVarDuplError()
	shorthandRuntimeVarEval()
	dontChangeTypeOfVar()
}
