/*
Strukture
=========

Šta je struktura?
-----------------
Struktura je korisnički definisan tip koji predstavlja kolekciju polja. Može se
koristiti na mestima gde je smisleno grupisati podatke u jednu programsku
jedinicu, umesto da se svaki od njih prikazuje kao zasebna vrednost.

Na primer, zaposleni ima ime (firstName), prezime (lastName) i godine (age).
Ima smisla grupisati ova tri svojstva u jednu strukturu pod nazivom Employee.

Deklarisanje tipa strukture
---------------------------

type Employee struct {
	firstName string
	lastName  string
	age       int
}

Gornji isečak deklariše tip strukture Employee sa poljima firstName, lastName i
age. Gornji tip strukture Employee se naziva imenovana struktura jer kreira novi
tip podataka nazvan Employee pomoću kojeg se mogu kreirati instance ove
strukture.

Ova struktura se takođe može učiniti kompaktnijom deklarisanjem polja koja
pripadaju istom tipu u jednom redu, nakon čega sledi ime tipa. U gornjoj
strukturi firstName i lastName pripadaju istom tipu string i stoga se struktura
može prepisati kao

type Employee struct {
	firstName, lastName string
	age                 int
}

Iako gore navedena sintaksa štedi nekoliko redova koda, ona ne čini deklaracije
polja eksplicitnim. Molimo vas da se uzdržite od korišćenja gore navedene
sintakse.

Kreiranje imenovanih struktura
------------------------------
Hajde da deklarišemo imenovanu strukturu Employee i kreiramo njene instance
koristeći sledeći jednostavan program.
*/

package structs

import (
	"fmt"
	stexp "learngo/16-structs/structsexported"
)

type Employee struct {
	firstName string
	lastName  string
	age       int
	salary    int
}

func structDeclAndCreateNamedType() {

	fmt.Println("\n --- Struct declaration and creation with struct literal ---")

	// Creating instance of Employee struct "specifying field names" and its
	// values. Order of fields can be different from the order in which
	// they are declared in the struct type declaration.
	// This is called struct literal.
	emp1 := Employee{
		firstName: "Sam",
		age:       25,
		salary:    500,
		lastName:  "Anderson",
	}

	// Creating instance of struct without specifying field names
	emp2 := Employee{"Thomas", "Paul", 29, 800}

	fmt.Println("Employee 1", emp1)
	fmt.Println("Employee 2", emp2)
}

/*
U gornjem programu, kreiramo imenovani tip strukture Employee. emp1 instanca
strukture Employee tipa je definisana navođenjem vrednosti za svako ime polja.
Redosled polja ne mora nužno biti isti kao redosled imena polja prilikom
deklarisanja tipa strukture.

U ovom slučaju, promenili smo poziciju lastName i pomerili ga na kraj. Ovo će
raditi bez ikakvih problema.

U gornjem programu, instanca emp2 je definisano izostavljanjem imena polja. U
ovom slučaju, neophodno je održati redosled polja istim kao što je navedeno u
deklaraciji tipa strukture.

Molimo vas da se uzdržite od korišćenja ove sintakse jer otežava određivanje
koja vrednost je za koje polje. Naveli smo ovaj format ovde samo da bismo
razumeli da je i ovo validna sintaksa :)

Gore navedeni program štampa

	>> Employee 1 {Sam Anderson 25 500}
	>> Employee 2 {Thomas Paul 29 800}

Kreiranje anonimnih struktura
-----------------------------
Moguće je deklarisati strukture bez kreiranja novog tipa podataka. Ovi tipovi
struktura se nazivaju anonimne strukture i koriste se "on place", po potrebi.
*/
func structDeclAndCreateAnonymousType() {

	fmt.Println("\n --- Struct declaration and creation with anonymous struct literal ---")
	emp3 := struct {
		firstName string
		lastName  string
		age       int
		salary    int
	}{
		firstName: "Andreah",
		lastName:  "Nikola",
		age:       31,
		salary:    5000,
	}

	fmt.Println("Employee 3", emp3)
}

/*
U gornjem programu definisana je anonimna strukturna promenljiva emp3. Kao što
smo već pomenuli, ova struktura se naziva anonimna jer samo kreira novu
strukturnu promenljivu emp3 i ne definiše nijedan novi tip strukture poput
imenovanih struktura.

Ovaj program štampa

	>> Employee 3 {Andreah Nikola 31 5000}

Pristup pojedinačnim poljima strukture
--------------------------------------
Operator tačka . se koristi za pristup pojedinačnim poljima strukture.
*/

func structAccessFields() {
	fmt.Println("\n --- Accessing individual fields of struct ---")
	emp6 := Employee{
		firstName: "Sam",
		lastName:  "Anderson",
		age:       55,
		salary:    6000,
	}
	fmt.Println("First Name:", emp6.firstName)
	fmt.Println("Last Name:", emp6.lastName)
	fmt.Println("Age:", emp6.age)
	fmt.Printf("Salary: $%d\n", emp6.salary)

	fmt.Println("Reasign salary field of emp6")
	emp6.salary = 6500
	fmt.Printf("New Salary: $%d\n", emp6.salary)
}

/*
Izraz emp6.firstName u gornjem programu pristupa firstName polju strukture emp6.
U liniji sa emp6.salary = 6500 menjamo platu zaposlenog. Ovaj program ispisuje,

	>> First Name: Sam
	>> Last Name: Anderson
	>> Age: 55
	>> Salary: $6000
	>> New Salary: $6500

Nulta vrednost strukture
------------------------
Kada je struktura definisana i nije eksplicitno inicijalizovana ni sa jednom
vrednošću, poljima strukture se podrazumevano dodeljuju nulte vrednosti.
*/

func structZeroValued() {

	fmt.Println("\n --- Zero valued struct ---")

	var emp4 Employee //zero valued struct

	fmt.Println("First Name:", emp4.firstName)
	fmt.Println("Last Name:", emp4.lastName)
	fmt.Println("Age:", emp4.age)
	fmt.Println("Salary:", emp4.salary)
}

/*
Gore navedeni program deklariše emp4 ali nije inicijalizovan nijednom vrednošću.
Stoga su firstName i lastName dodeljene nulte vrednosti stringa, koji je prazan
string "", a age i salary su dodeljene nulte vrednosti celobrojnog tipa, koji je
0. Ovaj program ispisuje,

	>> First Name:
	>> Last Name:
	>> Age: 0
	>> Salary: 0

Takođe je moguće navesti vrednosti za neka polja, a ostala ignorisati. U svakom
slučaju, ignorisanim poljima se dodeljuju nulte vrednosti.
*/

func structPartInit() {

	fmt.Print("\n --- Partially initialized struct ---\n")

	emp5 := Employee{
		firstName: "John",
		lastName:  "Paul",
	}

	fmt.Println("First Name:", emp5.firstName)
	fmt.Println("Last Name:", emp5.lastName)
	fmt.Println("Age:", emp5.age)
	fmt.Println("Salary:", emp5.salary)
}

/*
U gornjem programu firstName i lastName su inicijalizovani, dok age i salary
nisu. Stoga su age i salary dodeljene njihove nulte vrednosti. Ovaj program daje
rezultat:

	>> First Name: John
	>> Last Name: Paul
	>> Age: 0
	>> Salary: 0

Pointer na strukturu
--------------------
Takođe je moguće kreirati pointer na strukturu.
*/
func structPointer() {

	fmt.Println("\n --- Pointer to struct ---")

	emp8 := &Employee{
		firstName: "Sam",
		lastName:  "Anderson",
		age:       55,
		salary:    6000,
	}

	fmt.Printf("Type of emp8 is %T, value of *emp8 is %v\n", emp8, *emp8)
	fmt.Println("First Name:", (*emp8).firstName)
	fmt.Println("Age:", (*emp8).age)
}

/*
emp8 u gornjem programu je pointer na Employee strukturu. (*emp8).firstName je
sintaksa za pristup "firstName" polju emp8 strukture. Ovaj program ispisuje:

	>> First Name: Sam
	>> Age: 55

Go nam daje mogućnost da koristimo emp8.firstName umesto eksplicitne dereference
(*emp8).firstName za pristup firstName polju.
*/

func structPointer2() {

	fmt.Println("\n --- Pointer to struct without dereferencing ---")

	emp8 := &Employee{
		firstName: "Sam",
		lastName:  "Anderson",
		age:       55,
		salary:    6000,
	}
	fmt.Printf("Type of emp8 is %T, value of *emp8 is %v\n", emp8, *emp8)
	fmt.Println("First Name:", emp8.firstName)
	fmt.Println("Age:", emp8.age)
}

/*
Koristili smo emp8.firstName za pristup firstName polju u gornjem programu i
ovaj program takođe daje izlaz

	>> First Name: Sam
	>> Age: 55

!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
!!! Sintaksa emp8.firstName je ekvivalentna (*emp8).firstName   			!!!
!!! Generalno, kao što je (*array)[i] je ekvivalentno array[i], 			!!!
!!!(*struct).fieldname je ekvivalentmo struct.fieldname 					!!!
!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

Anonimna polja
--------------
Moguće je kreirati strukture sa poljima koja sadrže samo tip bez imena polja.
Ovakve vrste polja se nazivaju anonimna polja.

Isečak koda ispod kreira strukturu Person koja ima dva anonimna polja string i
int:

	type Person struct {
		string
		int
	}

Iako anonimna polja nemaju eksplicitno ime, podrazumevano ime anonimnog polja
je ime njegovog tipa. Na primer, u slučaju strukture Person iznad, iako su polja
anonimna, podrazumevano ona uzimaju ime tipa polja.

Dakle, Person struktura ima 2 polja sa imenom string i int.
*/

type Person struct {
	string
	int
}

func structAnonymousFields() {

	fmt.Println("\n --- Anonymous fields in struct ---")

	p1 := Person{
		string: "naveen",
		int:    50,
	}
	fmt.Println(p1.string)
	fmt.Println(p1.int)
}

/*
U gornjem programu, pristupamo anonimnim poljima strukture Person koristeći
njihove tipove kao ime polja, što je string i int respektivno. Izlaz gornjem
programu je:

	>> naveen
	>> 50

Ugnežđene strukture
-------------------
Moguće je da struktura sadrži polje koje je struktura. Ovakve vrste struktura
se nazivaju ugnežđene strukture.
*/

type Address struct {
	city  string
	state string
}

type Person2 struct {
	name    string
	age     int
	address Address
}

func structNested() {

	fmt.Println("\n --- Nested struct ---")

	p := Person2{
		name: "Naveen",
		age:  50,
		address: Address{
			city:  "Chicago",
			state: "Illinois",
		},
	}

	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
	fmt.Println("City:", p.address.city)
	fmt.Println("State:", p.address.state)
}

/*
Struktura Person2 u gornjem programu ima polje address koje je struktura tipa
Address. Ovaj program ispisuje:

	>> Name: Naveen
	>> Age: 50
	>> City: Chicago
	>> State: Illinois

Promovisana polja
-----------------
Polja koja su anonimna strukturna polja u strukturi nazivaju se promovisana
polja jer im se može pristupiti kao da pripadaju strukturi koja sadrži anonimno
strukturno polje.

Razumem da je ova definicija prilično složena, pa hajde da odmah zaronimo u kod
da bismo ovo razumeli :).
*/

type Address1 struct {
	city  string
	state string
}

type Person3 struct {
	name string
	age  int
	Address1
}

/*
U gornjem isečku koda, Person3 struktura ima anonimno polje Address1 koje je
struktura. Sada se polja Address1 strukture, city i state nazivaju promovisanim
poljima jer im se može pristupiti kao da su direktno deklarisana u Person3
strukturi.
*/

func structPromotedFields() {

	fmt.Println("\n --- Promoted fields in struct ---")

	p := Person3{
		name: "Naveen",
		age:  50,
		Address1: Address1{
			city:  "Chicago",
			state: "Illinois",
		},
	}

	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
	fmt.Println("City:", p.city)   //city is promoted field
	fmt.Println("State:", p.state) //state is promoted field
}

/*
U gornjem programu, promovisanim poljima city i state se pristupa kao da su
deklarisana u samoj strukturi p koristeći sintaksu p.city i p.state. Ovaj
program ispisuje:

	>> Name: Naveen
	>> Age: 50
	>> City: Chicago
	>> State: Illinois

Izvezene strukture i polja
--------------------------

Ako tip strukture počinje velikim slovom, onda je to izvezeni tip i može mu se
pristupiti iz drugih paketa. Slično tome, ako polja strukture počinju velikim
slovima, može im se pristupiti iz drugih paketa.

Hajde da napišemo program koji ima prilagođene pakete da bismo ovo bolje
razumeli.

Napravite direktorijum sa imenom structs u vašem "$GOPATH/src/16-structs"
direktorijumu.

mkdir $GOPATH/16-structs/structs

Hajde da kreiramo go subpaket pod nazivom "structsexported".

cd $GOPATH/16-structs/structsexported

Napravite još jedan direktorijum pod nazivom "computer" unutar
"structsexported."

mkdir computer

Unutar "computer" direktorijuma, kreirajte datoteku spec.go sa sledećim
sadržajem:

package computer

type Spec struct { //exported struct

		Maker string //exported field
		Price int    //exported field
		model string //unexported field
	}

Gore navedeni koda kreira paket computer koji sadrži izvezenu strukturu "Spec"
sa dva izvezena polja "Maker" i "Price" i jednim neizvezenim poljem "model".

Hajde da uvezemo ovaj paket iz roditeljskog paketa "structsexported"  i
koristimo Spec strukturu.

Napravite datoteku sa imenom structsexported.go unutar "structsexported"
direktorijuma i napišite sledeći program u njoj:

package structsexported

import (

	"learngo/16-structs/structsexported/computer"
	"fmt"

)

func main() {
	spec := computer.Spec{
		Maker: "apple",
		Price: 50000,
	}

	fmt.Println("Maker:", spec.Maker)
	fmt.Println("Price:", spec.Price)
}

Direktorijum structsexported treba da ima sledeću strukturu:

├── structsexported
    ├── computer
    │   └── spec.go
    └── structsexported.go

U gornjem programu, uvozimo computer paket. Potom pristupamo dvama izvezenim
poljima Maker i Price strukture Spec.
*/

func structExported() {

	fmt.Println("\n --- Exported struct and fields ---")

	stexp.StructExported()
}

/*
Kada pokrenete ovaj program, dobićete sledeći izlaz:

	>> Maker: apple
	>> Price: 50000

Ako pokušamo da pristupimo neizvezenom polju model, kompajler će se žaliti.
Zamenite sadržaj structsexported.go sledećim kodom.

	func structsexported() {
		spec := computer.Spec{
			Maker: "apple",
			Price: 50000,
			model: "Mac Mini",	// trying to access unexported field. Error!!!
		}
		fmt.Println("Maker:", spec.Maker)
		fmt.Println("Price:", spec.Price)
	}

U gornjem programu, pokušavamo da pristupimo neizvezenom polju "model".
Pokretanje ovog programa će rezultirati greškom pri kompajlaciji.

./main.go:12:13: unknown field 'model' in struct literal of type computer.Spec

Pošto model polje nije izvezeno, ne može mu se pristupiti iz drugih paketa.

Jednakost struktura
-------------------
Strukture su "vrednosni tipovi" i uporedive su ako je svako od njihovih polja
uporedivo. Dve strukturne promenljive se smatraju jednakim ako su im
odgovarajuća polja jednaka.
*/

type name struct {
	firstName string
	lastName  string
}

func structEqu() {

	fmt.Println("\n --- Struct equality ---")

	name1 := name{
		firstName: "Steve",
		lastName:  "Jobs",
	}
	name2 := name{
		firstName: "Steve",
		lastName:  "Jobs",
	}

	if name1 == name2 {
		fmt.Println("name1 and name2 are equal")
	} else {
		fmt.Println("name1 and name2 are not equal")
	}

	name3 := name{
		firstName: "Steve",
		lastName:  "Jobs",
	}
	name4 := name{
		firstName: "Steve",
	}

	if name3 == name4 {
		fmt.Println("name3 and name4 are equal")
	} else {
		fmt.Println("name3 and name4 are not equal")
	}
}

/*
U gornjem programu, struktura tipa name sadrži dva string polja. Pošto su
stringovi uporedivi, moguće je uporediti dve strukturne promenljive tipa name.

U gornjem programu name1 i name2 su jednaki, dok name 3i name4 nisu. Ovaj
program će ispisati,

	>> name1 and name2 are equal
	>> name3 and name4 are not equal

Strukturne promenljive nisu uporedive ako sadrže polja koja nisu uporediva
(hvala alasiji sa reddit-a što je ukazao na ovo).
*/

// type image struct {
// 	data map[int]int
// }

// func main() {
// 	image1 := image{
// 		data: map[int]int{
// 			0: 155,
// 		}}
// 	image2 := image{
// 		data: map[int]int{
// 			0: 155,
// 		}}
// 	if image1 == image2 {
// 		fmt.Println("image1 and image2 are equal")
// 	}
// }

/*
U gornjem programu, struktura tipa image sadrži polje data tipa map. Mape nisu
uporedive, stoga image1 i image2 se ne mogu porediti. Ako pokrenete ovaj
program, kompilacija će propasti sa greškom.

	>> ./prog.go:20:12: invalid operation: image1 == image2 (struct containing
	>> map[int]int cannot be compared)
*/

func StructFuncs() {
	fmt.Println("\n ---Structs---")
	structDeclAndCreateNamedType()
	structDeclAndCreateAnonymousType()
	structAccessFields()
	structZeroValued()
	structPartInit()
	structPointer()
	structPointer2()
	structAnonymousFields()
	structNested()
	structPromotedFields()
	structExported()
	structEqu()
}
