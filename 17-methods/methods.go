/*
Methods
=======

Uvod
----
Metoda je samo funkcija sa tipom prijemnika između func ključne reči i imena
metode. Prijemnik može biti strukturnog ili nestrukturnog tipa.

Sintaksa deklaracije metode
---------------------------

func (t Type) methodName(parameter list) {
}

Gornji kod kreira metodu methodName sa tipom prijemnika Type. t se naziva
prijemnikom (receiverom) može mu se pristupiti unutar metode.

Primer
------
Hajde da napišemo jednostavan program koji kreira metodu na tipu strukture i
poziva je.
*/

package methods

import (
	"fmt"
	"math"
)

type Employee struct {
	name     string
	salary   int
	currency string
}

// displaySalary() method has Employee as the receiver type
func (e Employee) displaySalary() {
	fmt.Printf("Salary of %s is %s%d\n", e.name, e.currency, e.salary)
}

func methodCreate() {

	fmt.Println("\n --- Creating a method on a struct type ---")

	emp1 := Employee{
		name:     "Sam Adolf",
		salary:   5000,
		currency: "$",
	}
	emp1.displaySalary() //Calling displaySalary() method of Employee type
}

/*
U gornjem programu, kreirali smo metodu displaySalary sa prijemnikom tipa
Employee struct. displaySalary() metoda ima pristup objektu e unutar sebe,
korišćenjem prijemnika pristupamo poljima strukture. U ovom slučaju,
ovde koristimo prijemnik e i ispisujemo ime, valutu i ​​platu zaposlenog.

Na kraju smo pozvali metodu koristeći sintaksu emp1.displaySalary().

Ovaj program štampa

	>> Salary of Sam Adolf is $5000.

Metode na spram funkcija
------------------------
Gore navedeni program se može prepisati koristeći samo funkcije i bez metoda.
*/

// displaySalary() method converted to function with Employee as parameter
func displaySalary(e Employee) {
	fmt.Printf("Salary of %s is %s%d\n", e.name, e.currency, e.salary)
}

func methodConvMethodToFunc() {

	fmt.Println("\n --- Converting method to function ---")

	emp1 := Employee{
		name:     "Sam Adolf",
		salary:   5000,
		currency: "$",
	}
	displaySalary(emp1)
}

/*
U gornjem programu, displaySalary metoda se konvertuje u funkciju i Employee
struktura joj se prosleđuje kao parametar. Ovaj program takođe proizvodi
potpuno isti izlaz

	>> Salary of Sam Adolf is $5000.

Pa zašto onda imamo metode kada možemo da napišemo isti program koristeći
funkcije. Postoji nekoliko razloga za to. Hajde da ih pogledamo jedan po jedan.

Go nije čisto objektno orijentisan programski jezik i ne podržava klase. Stoga
su "metode na tipovima" način da se postigne ponašanje slično klasama. Metode
omogućavaju logičko grupisanje ponašanja povezanih sa tipom sličnim klasama. U
gornjem primeru programa, sva ponašanja povezana sa Employee tipom mogu se
grupisati kreiranjem metoda koristeći Employee tip prijemnika. Na primer, možemo
dodati metode kao što su calculatePension, calculateLeaves i tako dalje.

Metode sa istim imenom mogu biti definisane na različitim tipovima, dok funkcije
sa istim imenima nisu dozvoljene. Pretpostavimo da imamo strukturu Square i
Circle. Moguće je definisati metodu sa imenom Area i na Square i na Circle tipu.
To je urađeno u programu ispod.
*/

type Rectangle struct {
	length int
	width  int
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() int {
	return r.length * r.width
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func methodName() {

	fmt.Println("\n --- Method with same name on different types ---")

	r := Rectangle{
		length: 10,
		width:  5,
	}
	fmt.Printf("Area of rectangle %d\n", r.Area())

	c := Circle{
		radius: 12,
	}

	fmt.Printf("Area of circle %f\n", c.Area())
	fmt.Println()
}

/*
Ovaj program štampa:

	>> Area of rectangle 50
	>> Area of circle 452.389342

Gore navedeno svojstvo metoda se koristi za implementaciju interfejsa. O tome
ćemo detaljno razgovarati u sledećem tutorijalu kada se budemo bavili
interfejsima.

Pointer prijemnici u odnosu na vrednosne prijemnike
----------------------------------------------------
Do sada smo videli metode samo sa vrednosnim prijemnicima. Moguće je kreirati
metode sa pointer prijemnicima. Razlika između vrednosnih prijemnika i pointer
prijemnika je u tome što su promene napravljene unutar metode sa pointer
prijemnikom vidljive pozivaocu, dok to nije slučaj kod vrednosnih prijemnika.
Hajde da ovo razumemo uz pomoć programa.
*/

type Employee2 struct {
	name string
	age  int
}

// Method with value receiver
func (e Employee2) changeName2(newName string) {
	e.name = newName
}

// Method with pointer receiver
func (e *Employee2) changeAge2(newAge int) {
	e.age = newAge
}

func methodReceivers() {
	fmt.Println("\n --- Method with value receiver vs pointer receiver ---")
	e := Employee2{
		name: "Mark Andrew",
		age:  50,
	}

	fmt.Println("\nValue receiver")
	fmt.Printf("Employee name before change: %s", e.name)
	e.changeName2("Michael Andrew")
	fmt.Printf("\nEmployee name after change: %s", e.name)

	fmt.Println("\n\nPointer receiver")
	fmt.Printf("Employee age before change: %d", e.age)
	(&e).changeAge2(51)
	fmt.Printf("\nEmployee age after change: %d", e.age)
	fmt.Println()
}

/*
U gornjem programu, changeName metoda ima vrednosni prijemnik (e Employee),
dok changeAge metoda ima pointer prijemnik (e *Employee2). Promene napravljene
u polju "name" Employee2 strukture unutar metode sa vrednosnim prijemnikom neće
biti vidljive pozivaocu i stoga program ispisuje isto ime i pre i posle poziva
metode changeName. Metoda changeAge ima pointer prijemnik, promene napravljene
u polju "age" nstrukture Employee2 nakon  poziva metode biće vidljive pozivaocu.

Ovaj program ispisuje:

	>> Employee name before change: Mark Andrew
	>> Employee name after change: Mark Andrew

	>> Employee age before change: 50
	>> Employee age after change: 51

U gornjem programu, koristimo (&e).changeAge(51) za poziv changeAge metode.
Pošto changeAge ima pointer prijemnik, koristili smo (&e) za poziv metode.
Ovo nije potrebno i kompajler nam daje mogućnost da jednostavno koristimo
e.changeAge(51). Kompajler će interpretirati e.changeAge(51) kao
(&e).changeAge(51)

Sledeći program je prepisan da se koristi e.changeAge(51) umesto
(&e).changeAge(51) i štampa isti izlaz.
*/

func methodReceiverAltSyntax() {

	fmt.Println("\n --- Method with value receiver vs pointer receiver (alternate syntax) ---")
	e := Employee2{
		name: "Mark Andrew",
		age:  50,
	}

	fmt.Println("\nValue receiver")
	fmt.Printf("Employee name before change: %s", e.name)
	e.changeName2("Michael Andrew")
	fmt.Printf("\nEmployee name after change: %s", e.name)

	fmt.Println("\n\nPointer receiver")
	fmt.Printf("Employee age before change: %d", e.age)
	e.changeAge2(51) // Using alternate syntax for pointer receiver
	fmt.Printf("\nEmployee age after change: %d\n", e.age)
}

/*
Kada koristiti pointer prijemnik, a kada vrednosni prijemnik
------------------------------------------------------------
Generalno, pointer prijemnici se mogu koristiti kada promene napravljene na
prijemniku unutar metode treba da budu vidljive pozivaocu.

Pointer prijemnici mogu se koristiti i na mestima gde je skupo kopirati
strukturu podataka. Razmotrite strukturu koja ima mnogo polja. Korišćenje ove
strukture kao prijemnika vrednosti u metodi zahtevaće kopiranje cele strukture,
što će biti skupo. U ovom slučaju, ako se koristi pointer prijemnik, struktura
se neće kopirati i samo će se pointer na nju koristiti u metodi.

U svim ostalim situacijama, mogu se koristiti prijemnici vrednosti.

Metode anonimnih strukturnih polja
----------------------------------
Metode koje pripadaju anonimnim poljima strukture mogu se pozivati kao da
pripadaju strukturi u kojoj je anonimno polje definisano.
*/

type address struct {
	city  string
	state string
}

func (a address) fullAddress() {
	fmt.Printf("Full address is: %s, %s\n", a.city, a.state)
}

type person struct {
	firstName string
	lastName  string
	address   // Anonymous field of type address
}

func methodAnonymousFieldMethod() {

	fmt.Println("\n --- Method on anonymous field in struct ---")

	p := person{
		firstName: "Elon",
		lastName:  "Musk",
		address: address{
			city:  "Los Angeles", // Promoting city field to person struct
			state: "California",  // Promoting state field to person struct
		},
	}

	p.fullAddress() //accessing fullAddress method of address struct
}

/*
U gornjem programu, pozivamo fullAddress() metod strukture tipa address
koristeći p.fullAddress(). Eksplicitni poziv p.address.fullAddress() nije
potreban. Ovaj program ispisuje:

	>> Full address: Los Angeles, California

Vrednosni prijemnik u metodama u odnosu na vrednosni argument u funkcijama
--------------------------------------------------------------------------
Ova tema je zanimljiva većini početnika. Pokušaću da bude što jasnija 😀.

Kada funkcija ima vrednosni argument, prihvatiće samo taj tip argumenta -
vrednosni argument.

Kada metoda ima vrednosni prijemnik, ona će prihvatiti i pointer prijemnik i
vrednosni prijemnik.

Hajde da ovo razumemo pomoću jednog primera.
*/

type rectangle2 struct {
	length int
	width  int
}

// Function with value argument
func area2(r rectangle2) {
	fmt.Printf("Area Function result: %d\n", (r.length * r.width))
}

// Method with value receiver
func (r rectangle2) area2() {
	fmt.Printf("Area Method result: %d\n", (r.length * r.width))
}

func methodValueReceiverVsValueArgument() {

	fmt.Println("\n --- Value receiver in methods vs value argument in functions ---")

	r := rectangle2{
		length: 10,
		width:  5,
	}

	fmt.Println("Calling function with value argument")
	area2(r) //calling function with value argument

	fmt.Println("Calling method with value receiver")
	r.area2() //calling method with value receiver

	p := &r

	fmt.Println("Calling function with pointer argument")
	//area2(p) // Uncommenting this line will cause a compilation error
	fmt.Println("Cannot use pointer function argument instead of value argument in function")

	fmt.Println("Calling method with pointer receiver")
	p.area2()
}

/*
Funkcija func area2(r rectangle2)u prihvata vrednosni argument, a metoda
func (r rectangle) area2() prihvata vrednosni prijemnik.

U gornjem programu pozivamo funkciju area2 sa vrednosnim argumentom area2(r) i
ona će raditi. Slično, pozivamo metodu area2 r.area2() koristeći vrednosni
prijemnik i ovo će takođe raditi.

Potom kreiramo pointer na strukturu r. Ako pokušamo da prosledimo ovaj pointer
funkciji area2 koja prihvata samo vrednost, kompajler će se žaliti.

Komentarisao sam tu liniju dq bi ostatak programa radio. Ako dekomentišete ovu
liniju, kompajler će izbaciti grešku

	>> * compilation error, cannot use p (type rectangle) as type rectangle in
	>> argument to area2.

Ovo funkcioniše tačno kako se očekuje.

Sada dolazi najteži deo koda, p.area2() poziva metodu area2 koja prihvata samo
vrednosni prijemnik koristeći pointer p. Ovo je sasvim ispravno. Razlog je taj
što će p.area2() radi praktičnosti, biti interpretirana od strane Go-a
kompajlera, (*p).area2(), pošto area2 ima vrednosni prijemnik.

Ovaj program će štampati,

	>> Calling function with value argument
	>> Area Function result: 50
	>> Calling method with value receiver
	>> Area Method result: 50
	>> Calling function with pointer argument
	>> Cannot use pointer function argument instead of value argument in function
	>> Calling method with pointer receiver
	>> Area Method result: 50

Pointer prijemnici u metodama u odnosu na pointer argumente u funkcijama
--------------------------------------------------------------------------
Slično vrednosnim argumentima, funkcije sa pointer argumentima prihvataće samo
pointere, dok će metode sa pointer prijemnicima prihvatati i pointer i
vrednosne prijemnike.
*/

type rectangle3 struct {
	length int
	width  int
}

func perimeter3(r *rectangle3) {
	fmt.Println("perimeter function output:", 2*(r.length+r.width))
}

func (r *rectangle3) perimeter3() {
	fmt.Println("perimeter method output:", 2*(r.length+r.width))
}

func methodPointerReceiverVsPointerArgument() {

	fmt.Println("\n --- Pointer receiver in methods vs pointer argument in functions ---")

	r := rectangle3{
		length: 10,
		width:  5,
	}

	p := &r

	perimeter3(p)  // calling function with pointer argument
	p.perimeter3() // calling method with pointer receiver

	fmt.Println("cannot use r (type rectangle3) as type *rectangle3 in argument to perimeter function")
	//perimeter3(r) // calling function with value argument istead of pointer argument
	r.perimeter3() // calling method with pointer receiver using value receiver
}

/*
U gornjem programu definišemo funkciju perimeter3 koja prihvata pointerski
argument, i definišemo metodu perimeter3 koja ima pointerski prijemnik.

Pozivamo funkciju perimeter3 sa pointer argumentom, a potom pozivamo metod
perimeter3 sa pointer prijemnikom. Sve je u redu.

U komentarisanoj liniji pokušavamo da pozovemo funkciju perimeter3 sa vrednosnim
argumentom r. Ovo nije dozvoljeno jer funkcija sa pointer argumentom pointera
neće prihvatiti vrednosni argument. Ako se ova linija dekomentiše i program se
pokrene, kompilacija će biti neuspešno sa greškom

	>> * main.go:33: cannot use r (variable of struct type rectangle3) as
	>> *rectangle3 value in argument to perimeter3

Na kraju pozivamo pozivamo metodu sa pointer prijemnikom perimeter3 sa
vrednosnim prijemnikom. Ovo je dozvoljeno. Ovaj program će ispisati:

	>> perimeter function output: 30
	>> perimeter method output: 30
	>> perimeter method output: 30

Metode sa prijemnicima koji nisu strukturni
-------------------------------------------
Do sada smo definisali metode samo na strukturnim tipovima. Takođe je moguće
definisati metode na tipovima koji nisu strukture, ali postoji jedna začkoljica.
Da bi se definisala metoda na tipu, definicija tipa prijemnika i definicija
metode "treba da budu prisutne u istom paketu". Do sada su sve strukture i
metode na strukturama koje smo definisali bile smeštene u istom paketu i stoga
su uvek radile.

package main

func (a int) add(b int) {
}

func main() {

}

U gornjem programu, pokušavamo da dodamo metodu po imenu add ugrađenom tipu int.
Ovo nije dozvoljeno ako definicija metode add i definicija tipa int nisu u
istom paketu. Ovaj program će izbaciti grešku pri kompajlaciji:

	>>„ Ne mogu da definišem nove metode na nelokalnom tipu int“.

Način da se ovo pokrene jeste da se kreira "alias" tipa za ugrađeni tip int, a
zatim da se kreira metod sa ovim aliasom tipa kao prijemnikom.
*/

type myInt int

func (a myInt) add(b myInt) myInt {
	return a + b
}

func methodUnstructuredType() {

	fmt.Println("\n --- Method on unstructured type ---")

	num1 := myInt(4)
	num2 := myInt(11)

	sum := num1.add(num2)
	fmt.Println("Sum is", sum)
}

/*
U gornjem programu, kreirali smo alias tipa myInt za tip int. Definisali smo metodu
add sa myInt kao prijemnikom.

Ovaj program će štampati

	>> Sum is 15.
*/

func MethodFuncs() {

	fmt.Println("\n --- Methods ---")

	methodCreate()
	methodConvMethodToFunc()
	methodName()
	methodReceivers()
	methodReceiverAltSyntax()
	methodAnonymousFieldMethod()
	methodValueReceiverVsValueArgument()
	methodPointerReceiverVsPointerArgument()
	methodUnstructuredType()
}
