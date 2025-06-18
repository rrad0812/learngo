/*
Konstante

Šta je konstanta?
-----------------
Konstante u Gou se koriste za označavanje fiksnih statičkih vrednosti kao što su:

	95
	"I love Go"
	67.89

i tako dalje. Konstante se uglavnom koriste za predstavljanje vrednosti koje se
ne menjaju tokom životnog veka aplikacije.

Deklarisanje konstante
----------------------
Ključna reč "const" se koristi za deklarisanje konstante u Go-u. Pogledajmo
primer:
*/

package vars

import (
	"fmt"
	"math"
)

func ctoDecl() {

	fmt.Println("\n --- ctoDecl ---")

	const a = 50
	fmt.Println(a)
}

/*
U gornjem kodu u "a" se nalazi konstanta vrednost 50.

Deklarisanje grupe konstanti
----------------------------
Postoji i druga sintaksa za definisanje grupe konstanti korištenjem jedne
naredbe. Primer definisanja grupe konstanti:
*/

func ctoMultiDecl() {

	fmt.Println("\n --- ctoMultiDecl ---")

	const (
		retryLimit = 4
		httpMethod = "GET"
	)

	fmt.Println(retryLimit)
	fmt.Println(httpMethod)
}

/*
U gornjem programu deklarisali smo 2 konstante "retryLimit" i "httpMethod".
Gornji program ispisuje

	>> 4
	>> GET

Konstantama, kao što im samo ime govori, ne mogu se ponovo dodeliti druge
vrednosti. U programu ispod, pokušavamo dodeliti drugu vrednost 89 na "a". To
nije dozvoljeno jer "a" je konstanta. Ovaj program neće uspeti pri
kompajliranju i dobićemo grešku:

	>> cannot assign to a (neither addressable nor a map index expression).
*/

func ctoReAssignError() {

	fmt.Println("\n --- ctoReAssignError ---")

	const a = 55 //allowed
	fmt.Println("a is", a)
	// a = 89 //reassignment not allowed
	fmt.Println("a is", a)
	fmt.Println("In Go, reassigment of constant is forbiden.")
}

/*
Vrednost konstante treba biti poznata u vreme kompajliranja. Stoga se ne može
dodeliti vrednost koju vraća poziv funkce jer se poziv funkce odvija u vreme
izvršavanja.
*/

func ctoCompileTime() {

	fmt.Println("\n --- ctoCompileTime ---")

	var a = math.Sqrt(4) //allowed
	fmt.Println("var a is", a)

	// const b = math.Sqrt(4) //not allowed
	const b = 2
	fmt.Println("const b is", b)
}

/*
U gornjem programu, "a" je varijabla i stoga se može dodeliti rezultat funkcije
math.Sqrt(4).

b je konstanta i vrednost b mora biti poznata u vreme kompajliranja. Funkcija
math.Sqrt(4)će biti evaluirana samo tokom izvršavanja i stoga c
onst b = math.Sqrt(4) se kompajlira sa greškom:

	>> ./prog.go:11:12: math.Sqrt(4) (value of type float64) is not constant

String konstante, tipizirane i netipizirane konstante
-------------------------------------------------------
Bilo koja vrednost zatvorena između dvostrukih navodnika je string konstanta u
Gou. Na primer, stringovi poput "Hello World", "Sam" su svi konstante u Gou.

Kojem tipu pripada string konstanta? Odgovor je da su to netipizirane konstante.

String konstanta poput „Zdravo svete“ nema nijedan tip.

const hello = "Hello World"

U gornjoj liniji koda, konstanta hello nema tip.

Go je jezik sa strogo definisanim tipom. Sve varijable zahtevaju eksplicitni
tip. Kako funkcioniše sledeći program koji dodjeljuje varijabli name
netipiziranu konstantu n?
*/

func ctoUntyped() {

	fmt.Println("\n --- ctoUntiped ---")

	const n = "Sam"
	var name = n

	fmt.Printf("Type of name is %T, and value is %v\n", name, name)
}

/*
Odgovor je da netipizirane konstante imaju zadani tip koji je s njima povezan
i one ga dodjeljuju ako i samo ako to zahteva kod. U naredbi

var name = n

promenljiva name zahteva tip i dobija ga iz zadanog tipa konstante n, koji je
string.

Postoji li način za kreiranje tipizirane konstante?
Odgovor je da. Sledeći kod kreira tipiziranu konstantu.

const name string = "Hello World"

name u gornjem kodu je konstanta tipa string.

Go je jezik sa strogo definisanim tipovima. Mešanje tipova tokom izvršenja
nije dozvoljeno. Pogledajmo šta to znači uz pomoć jednog programa:
*/

func ctoMixedType() {

	fmt.Println("\n --- ctoMixedType ---")
	var defaultName = "Sam" //allowed

	// Define a new type myString
	// if we write type myString = string it is alias for string type
	type myString string
	var customName myString = "Sam" //allowed

	// customName = defaultName //not allowed
	fmt.Printf("Type of defaultName is %T\n", defaultName)
	fmt.Printf("Type of customName is %T\n", customName)
}

/*
U gornjem kodu, prvo kreiramo varijablu defaultName i dodjeljujemo je konstantu
vrednost Sam. Podrazumevani tip konstante Sam je string, tako da je nakon dodele
defaultName tipa string.

U sledećem redu kreiramo novi tip myString.
PS. Ako želimo da to bude alias onda:

type myString = string

Zatim kreiramo varijablu customName tipa myString i dodelimo joj konstantu Sam.
Budući da konstanta Sam string tipa, može se dodeliti bilo kojoj string
varijabli. Stoga je ovo dodeljivanje dozvoljeno i customName dobija tip myString.

Sada imamo varijablu defaultName tipa string i drugu varijablu customName tipa
myString. Iako znamo da je myString string, Go-ova politika strogog tipiziranja
ne dozvoljava da se varijable jednog tipa dodele drugom. Stoga dodeljivanje

customName = defaultNamene

nije dozvoljeno i kompajler izbacuje grešku:

	>> ../prog.go:9:15: cannot use defaultName (variable of type string)
	>> as myString value in assignment

Da bi gornji program radio, defaultName se mora pretvoriti u tip myString.
To je urađeno u sledećem programu:
*/

func ctoConv() {

	fmt.Println("\n --- ctoConv ---")

	var defaultName = "Sam" //allowed
	type myString string
	var customName myString = "Sam" //allowed

	customName = myString(defaultName) //allowed

	fmt.Println(customName)
}

/*
Gornji program će ispisati:

	>> Sam

Booleove konstante
------------------
Booleove konstante se ne razlikuju od string konstanti. To su dve netipizirane
konstante true i false. Ista pravila za string konstante primjenjuju se na
booleove, tako da ih ovde nećemo ponavljati. Sledeći je jednostavan program za
objašnjenje booleovih konstanti:
*/

func ctoBool() {

	fmt.Println("\n --- ctoBool ---")

	const trueConst = true
	type myBool bool

	var defaultBool = trueConst       //allowed
	var customBool myBool = trueConst //allowed

	// defaultBool = customBool       // not allowed
	defaultBool = bool(customBool) // allowed

	fmt.Printf("Type of defaultBool is %T\n", defaultBool)
	fmt.Printf("Type of customBool is %T\n", customBool)
}

/*
Gore navedeni program je samorazumljiv.

Numeričke konstante
-------------------
Numeričke konstante uključuju cele, s pomičnim zarezom i kompleksne konstante.
Postoje neke suptilnosti kod numeričkih konstanti.

Pogledajmo nekoliko primera kako bismo stvari razjasnili:
*/

func ctoNumeric() {

	fmt.Println("\n --- ctoNumeric ---")

	const c = 5
	var intVar int = c

	var int32Var int32 = c
	var float64Var float64 = c
	var complex64Var complex64 = c

	fmt.Println("intVar", intVar, "\nint32Var", int32Var, "\nfloat64Var",
		float64Var, "\ncomplex64Var", complex64Var)
}

/*
U gornjem programu, konstanta c je untyped i ima vrednost 5. Možda se pitate
koji je podrazumevani tip promenljive c i ako ga ima, kako ga onda dodelimo
varijablama različitih tipova. Odgovor leži u sintaksi promenljive c. Sledeći
program će stvari pojasniti.
*/

func ctoNumeric2() {

	fmt.Println("\n --- ctoNumeric2 ---")

	var i = 5
	var f = 5.6

	var c = 5 + 6i
	fmt.Printf("i's type is %T, f's type is %T, c's type is %T\n", i, f, c)
}

/*
U gornjem programu, tip svake promenljive određen je sintaksom numeričke
konstante. 5 je celi broj, 5.6 je broj s pomičnim zarezom, a 5 + 6i je
kompleksan broj. Kada se gornji program pokrene, ispisuje se

	>> i's type is int, f's type is float64, c's type is complex128

S ovim znanjem, pokušajmo razumeti kako sledeći program funkcioniše:
*/

func ctoNumeric3() {

	fmt.Println("\n --- ctoNumeric3 ---")

	const c = 5

	var intVar int = c
	var int32Var int32 = c
	var float64Var float64 = c
	var complex64Var complex64 = c

	fmt.Println("intVar", intVar, "\nint32Var", int32Var, "\nfloat64Var",
		float64Var, "\ncomplex64Var", complex64Var)
}

/*
U gornjem programu, vrednost c je 5, a sintaksa c je generička. Može
predstavljati broj s pomičnim zarezom, celi broj ili čak kompleksan broj bez
imaginarnog dela. Stoga je moguće da se dodeli bilo kojem kompatibilnom tipu.

Zadani tip ovih vrsta konstanti može se smatrati generičkim na osnovu konteksta
u kojem se koriste.

U sledećoj dodeli:

var intVar int = c

zahteva se da "c" bude int, pa postaje int konstanta.

Dok u sledećoj dodeli:

var complex64Var complex64 = c

zahteva se da "c" bude kompleksan broj i stoga postaje kompleksna konstanta.
Prilično zgodno :). Naravno ovo je moguće samo kod komapatabilnih tipova.

Sledeća dodela neće uspeti:

const d = 5.11
var intVar int = d

sa greškom:

	>> 05-constants/constants.go:308:19: cannot use d (untyped float constant
	>> 5.11) as int value in variable declaration (exit status 1)

Numerički izrazi
------------------
Numeričke konstante se mogu slobodno mešati i uparivati ​​u izrazima, a tip je
potreban samo kada su dodeljene promenljivim ili korištene na bilo kojem mestu
u kodu koje zahteva tip.
*/

func ctoExpr() {

	fmt.Println("\n --- ctoExpr ---")
	var a = 5.9 / 8
	fmt.Printf("a = 5.9/8 = %v, and type of a is %T\n", a, a)
}

/*
U gornjem programu, 5.9 je float po sintaksi, a 8 je celi broj po sintaksi.
Ipak, 5.9/8 je dozvoljeno jer su oba numeričke konstante. Rezultat deljenja je
0.7375 što je je float  stoga je promenljiva "a" tipa float. Izlaz programa je:

	>> a = 5.9/8 = 0.7375, and type of a is float64
*/
func CtoFunc() {

	fmt.Println("\n --- Constants ---")

	ctoDecl()
	ctoMultiDecl()
	ctoReAssignError()
	ctoCompileTime()
	ctoUntyped()
	ctoMixedType()
	ctoConv()
	ctoBool()
	ctoNumeric()
	ctoNumeric2()
	ctoNumeric3()
	ctoExpr()
}
