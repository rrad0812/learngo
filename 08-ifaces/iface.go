/*
Interfejsi
==========

Šta je interfejs?
-----------------
U jeziku Go, interfejs je skup potpisa metoda. Kada tip pruža definiciju za sve
metode u interfejsu, kaže se da implementira interfejs. Interfejs određuje koje
metode tip treba da ima, a tip odlučuje kako da implementira te metode.

Interfejsi su ključni deo Go jezika i omogućavaju apstrakciju i fleksibilnost u
dizajnu programa. Oni omogućavaju različitim tipovima da budu tretirani kao
isti tip ako implementiraju isti interfejs, što olakšava pisanje generičkog
koda. Interfejsi su slični apstraktnim klasama u drugim programskim jezicima,
ali su u Go-u mnogo lakši za korišćenje.

Interfejsi se često koriste za definisanje zajedničkog ponašanja koje različiti
tipovi mogu imati, bez obzira na njihovu konkretno implementiranu strukturu.

Na primer, interfejs može da definiše metode za obradu plaćanja, a različiti
tipovi plaćanja (kao što su kreditne kartice, PayPal ili direktna zaduženja)
mogu implementirati te metode na različite načine.

Na primer, PaymentProcessor može biti interfejs sa potpisima metoda
ProcessPayment() i GenerateReceipt(). Za svaki tip koji pruža definicije za
ProcessPayment() i GenerateReceipt() metode se kaže da implementira
PaymentProcessor interfejs.

Ovo može da uključuje strukture poput CreditCardProcessor, PayPalProcessor ili
DirectDebitProcessor od kojih svaka implementira metode na način specifičan za
njihov sistem plaćanja.

Deklarisanje i implementacija interfejsa
----------------------------------------
Napisaćemo jednostavan program koji izračunava ukupne troškove za kompaniju na
osnovu plata zaposlenih. Radi kratkoće, pretpostavili smo da su svi troškovi u
američkim dolarima.
*/
package ifaces

import (
	"fmt"
)

/*
Deklarisanje interfejsa
SalaryCalculator interfejs definiše metodu CalculateSalary() int. Ova metoda
treba da bude implementirana od strane tipova koji implementiraju ovaj
interfejs.

Interfejs ne sadrži nikakve podatke, samo potpis metode. Ovo omogućava
različitim tipovima da implementiraju ovaj interfejs na različite načine,
čime se postiže fleksibilnost i apstrakcija u kodu.
*/
type SalaryCalculator interface {
	CalculateSalary() int
}

/*
Definisanje struktura koje implementiraju SalaryCalculator interfejs

Permanent i Contract su tipovi koji implementiraju SalaryCalculator
interfejs.

Permanent tip predstavlja stalno zaposlenog, dok Contract tip predstavlja
zaposlenog na ugovor. Obe strukture imaju empId i basicpay polja, a Permanent
takođe ima pf polje koje predstavlja doprinos za penzijsko osiguranje
(Provident Fund).

Plata stalno zaposlenog je zbir osnovne plate i doprinosa za penzijsko
osiguranje, dok je plata zaposlenog na ugovor samo osnovna plata.

Ove strukture implementiraju SalaryCalculator interfejs tako što definišu
metodu CalculateSalary() int.
*/

type Permanent struct {
	empId    int
	basicpay int
	pf       int
}

type Contract struct {
	empId    int
	basicpay int
}

/*
Metode koje implementiraju SalaryCalculator interfejs

Plata stalnog zaposlenika se izračunava kao zbir osnovne plate i doprinosa
za penzijsko osiguranje (Provident Fund). Ova metoda je implementirana na
Permanent tipu. Ova metoda vraća ukupnu platu stalnog zaposlenika.
*/

func (p Permanent) CalculateSalary() int {
	return p.basicpay + p.pf
}

/*
Plata zaposlenog na ugovor je osnovna plata sama. Ova metoda je implementirana
na Contract tipu. Ova metoda vraća osnovnu platu zaposlenog na ugovor.
*/

func (c Contract) CalculateSalary() int {
	return c.basicpay
}

/*
totalExpense funkcija
Ova funkcija prima slice tipa SalaryCalculator i izračunava ukupne troškove
kompanije tako što poziva CalculateSalary() metodu za svakog zaposlenog u
slice-u. Funkcija sabira sve plate i ispisuje ukupne mesečne troškove.
*/
func totalExpense(s []SalaryCalculator) {
	expense := 0
	for _, v := range s {
		// Pozivanje CalculateSalary() metode za svakog zaposlenog
		// u slice-u i dodavanje rezultata na ukupne troškove
		expense = expense + v.CalculateSalary()
	}
	fmt.Printf("Total Expense Per Month $%d\n", expense)
}

/*
ifaceBegining funkcija
Ova funkcija kreira nekoliko zaposlenih tipa Permanent i Contract, dodaje ih
u slice tipa SalaryCalculator i poziva totalExpense funkciju da izračuna i
prikaže ukupne mesečne troškove.
*/
func ifaceBegining() {

	fmt.Println("\n ---Interfeacesc elementary ---")

	pemp1 := Permanent{
		empId:    1,
		basicpay: 5000,
		pf:       20,
	}
	pemp2 := Permanent{
		empId:    2,
		basicpay: 6000,
		pf:       30,
	}
	cemp1 := Contract{
		empId:    3,
		basicpay: 3000,
	}

	employees := []SalaryCalculator{pemp1, pemp2, cemp1} // 2. ključna tačka

	totalExpense(employees) // 3. ključna tačka
}

/*
Ovo se prilično razlikuje od drugih jezika poput Jave gde klasa mora eksplicitno
da navede da implementira interfejs koristeći implements ključnu reč.
Ovo nije potrebno u Go-u, a Go interfejsi su implicitno implementirani ako tip
sadrži sve metode deklarisane u interfejsu.

Izlaz programa:

	>> Total Expense Per Month $14050

Generalno:
Priča o interfejsima u Go-u se može svesti na tri ključne tačke:
	1. Implementacija metoda interfejsa za određene tipove.
	2. Prikupljanje instanci tipova koji implementiraju isti interfejs u slice.
	3. Obrada slice-a pozivanjem metoda interfejsa.

Radosav agregirao 04.0.2025

Najveća prednost ovoga načina je što se funkcija totalExpense može proširiti na
bilo koji novi tip zaposlenog bez ikakvih promena koda.

Recimo da kompanija doda novi tip zaposlenog Freelancer sa drugačijom strukturom
plate. Tip Freelancer se može jednostavno proslediti u argumentu slice
totalExpense funkciji bez ijedne promene koda u totalExpense funkciji. Ova
metoda će uraditi ono što bi trebalo da uradi, a tip Freelancer će
implementirati SalaryCalculator interfejs :).

Hajde da modifikujemo ovaj program i dodamo novog Freelancer zaposlenog. Plata
za frilensera je proizvod satnice i ukupnog broja odrađenih sati.
*/

type Freelancer struct {
	empId       int
	ratePerHour int
	totalHours  int
}

/*
CalculateSalary metoda za Freelancer tip
Ova metoda izračunava ukupnu platu frilensera kao cene po satu i ukupnog broja
odrađenih sati. Ova metoda implementira SalaryCalculator interfejs, što
omogućava da se Freelancer tip koristi u totalExpense funkciji.

Ovo je ključna tačka gde se novi tip integriše u postojeći sistem bez potrebe
za promenom totalExpense funkcije. Ovo omogućava lako proširivanje sistema
sa novim tipovima koji implementiraju SalaryCalculator interfejs, čime se
postiže fleksibilnost i modularnost koda.
*/
func (f Freelancer) CalculateSalary() int {
	return f.ratePerHour * f.totalHours
}

/*
ifaceExt funkcija
Ova funkcija kreira nekoliko instanci Permanent, Contract i Freelancer tipova,
dodaje ih u slice tipa SalaryCalculator i poziva totalExpense funkciju da
izračuna i prikaže ukupne mesečne troškove.
*/
func ifaceExt() {

	fmt.Println("\n ---Interfaces with Freelancer---")

	pemp1 := Permanent{
		empId:    1,
		basicpay: 5000,
		pf:       20,
	}
	pemp2 := Permanent{
		empId:    2,
		basicpay: 6000,
		pf:       30,
	}
	cemp1 := Contract{
		empId:    3,
		basicpay: 3000,
	}
	freelancer1 := Freelancer{
		empId:       4,
		ratePerHour: 70,
		totalHours:  120,
	}
	freelancer2 := Freelancer{
		empId:       5,
		ratePerHour: 100,
		totalHours:  100,
	}

	employees := []SalaryCalculator{pemp1, pemp2, cemp1, freelancer1, freelancer2}

	totalExpense(employees)
}

/*
Izlaz programa:

	>> Total Expense Per Month $32450

Interna reprezentacija interfejsa
---------------------------------
Interfejs se može smatrati interno predstavljen parom (type, value). "type" je
osnovni konkretni tip koji implementira interfejs i "value" sadrži vrednost
konkretnog tipa.

Hajde da napišemo program da bismo bolje razumeli:

U ovom programu, definišemo Worker interfejs koji ima jednu metodu Work().

Worker interfejs predstavlja radnika koji ima sposobnost da radi. Ovaj interfejs
može biti implementiran od strane različitih tipova koji imaju sposobnost da
rade, kao što su osobe, mašine ili bilo koji drugi entitet koji može da
izvršava posao.

Interfejs omogućava apstrakciju i fleksibilnost u radu sa različitim tipovima
radnika, omogućavajući da se različiti tipovi tretiraju na isti način kada se
radi o radu. Ovo je korisno u situacijama gde želimo da radimo sa različitim
tipovima radnika bez potrebe da znamo tačno koji su to tipovi, sve dok
implementiraju Work() metodu.
*/
type Worker interface {
	Work()
}

/*
Person tip
Person tip predstavlja osobu koja može da radi. Ovaj tip sadrži polje name koje
predstavlja ime osobe.
*/
type Person struct {
	name string
	age  int
}

/*
Person tip implementira Worker interfejs tako što definiše Work() metodu.
Ova metoda ispisuje ime osobe i poruku da je na poslu. Ovo omogućava da se
Person tip koristi gde god je potreban Worker interfejs. Ovo je primer kako se
konkretni tip može koristiti u interfejsu.
*/
func (p Person) Work() {
	fmt.Println(p.name, "is working")
}

/*
describe funkcija
Ova funkcija prima Worker interfejs kao argument i ispisuje tip i vrednost
interfejsa. Ovo je korisno za razumevanje kako se interfejsi ponašaju u Go-u.

Kada se prosledi Worker interfejs, funkcija će ispisati konkretni tip koji
implementira interfejs i njegovu vrednost. Ovo je korisno za debagovanje i
razumevanje kako se interfejsi koriste u programu. Ova funkcija takođe
pokazuje kako se interfejsi mogu koristiti za apstrakciju i fleksibilnost u
radu sa različitim tipovima koji implementiraju isti interfejs.
*/
func describe(w Worker) {
	fmt.Printf("Interface type is %T and value is %v\n", w, w)
}

func ifaceInternRepr() {

	fmt.Println("\n ---Interfaces with internal representation---")

	p := Person{
		name: "Radosav",
		age:  65,
	}

	var w Worker = p // Person je tipa Worker jer implementira Worker interfejs

	describe(w) // Ispisuje konkretni tip i vrednost interfejsa

	w.Work() // Poziva Work() metodu interfejsa Worker
}

/*
Worker interfejs ima jednu metodu Work(), a tip strukture Person implementira
taj interfejs. Dodeljujemo promenljivu p tipa Person promenljivoj w tipa Worker.
Ovo je moguće jer tip Person implementira Worker interfejs.

Sada interfejs Worker sadrži konkretni tip i sadrži polje name i age. Funkcija
describe ispisuje vrednost i konkretan tip interfejsa, a Work metod interfejsa
štampa ime osobe i poruku da je na poslu.

Ovaj program štampa:

	>> Interface type is main.Person and value is {Radosav 65}
	>> Radosav is working

Više o tome kako izvući osnovnu vrednost interfejsa ćemo razgovarati u narednim
odeljcima.

Prazan interfejs
----------------
Interfejs koji ima nula metoda naziva se prazan interfejs. Predstavljen je kao
interface{}. Pošto prazan interfejs ima nula metoda, svi tipovi implementiraju
prazan interfejs.
*/

func describe2(i interface{}) {
	fmt.Printf("Type = %T, value = %v\n", i, i)
}

func ifaceEmpty() {

	fmt.Println("\n --- Empty interface---")

	s := "Hello World"
	describe2(s)

	i := 55
	describe2(i)

	strt := struct {
		name string
	}{
		name: "Radosav R",
	}
	describe2(strt)
}

/*
U gornjem programu, u liniji br. 7, describe(i interface{}) funkcija uzima
prazan interfejs kao argument i stoga se može proslediti bilo koji tip.

Prosleđujemo string, int i struct funkciji describe, respektivno. Ovaj program
ispisuje,

	>> Type = string, value = Hello World
	>> Type = int, value = 55
	>> Type = structRadosav Rame string }, value = {Radosav R}

Tvrdnja tipa
------------
U Go-u, interfejsi se mogu koristiti za apstrakciju i fleksibilnost u radu sa
različitim tipovima. Međutim, ponekad je potrebno dobiti osnovnu vrednost
interfejsa i raditi sa njom kao sa konkretnim tipom. Ovo se može postići
korišćenjem tvrdnje tipa (type assertion).

i.(T) je sintaksa koja se koristi za dobijanje osnovne vrednosti interfejsa
čiji je konkretni tip T.

Program vredi hiljadu reči 😀. Hajde da napišemo jedan za tvrdnju tipa.
*/

func assert(i interface{}) {
	s := i.(int) //get the underlying int value from i
	fmt.Printf("Konkretan tip je %T, vrednost je %d\n", s, s)
}

func ifaceTypeAssertion() {

	fmt.Println("\n --- Type Assertion int---")

	var s interface{} = 56
	assert(s)
}

/*
U gornjem programu konkretan tip je int. Koristimo sintaksu i.(int) da bismo
preuzeli osnovnu int vrednost i. Ovaj program ispisuje:

>> 56.

Šta će se desiti ako konkretan tip u gornjem programu nije int? Pa, hajde da
saznamo.
*/

// func ifaceTypeAssertion2() {

// 	fmt.Println("\n --- Type Assertion 2 ---")

// 	var s interface{} = "Steven Paul"
// 	assert(s)
// }

/*
Ako u gornjem programu pokušamo da prenosemo konkretan tip string na assert
funkciju, a unutar nje se provera radi na int, ovaj program će izazvati
paniku sa porukom:

	>> panic: interface conversion: interface {} is string, not int

Da bismo rešili gore navedeni problem, možemo koristiti sintaksu:

v, ok := i.(T)

Ako je konkretni tip T tada:
	- ok je true, v ima osnovnu vrednost,
	- ok je false, v ima osnovnu vrednost jednaku nultoj vrednosti tipa T i
	  program neće paničiti.
*/

func assertOk(i interface{}) {
	v, ok := i.(int)
	fmt.Println(v, ok)
}

func ifaceAssertOk() {

	fmt.Println("\n --- Type Assertion with ok ---")

	var s interface{} = 56
	assertOk(s)

	var i interface{} = "Steven Paul"
	assertOk(i)
}

/*
Kada se funkciji assertOk prosledi "Steven Paul", ok će biti false jer
konkretni tip nije int i v ima vrednost 0 što je nulta vrednost tipa int.
Ovaj program će ispisati,

	>> 56 true
	>> 0 false

Prekidač tipa
-------------
Prekidač tipa se koristi za upoređivanje konkretnog tipa interfejsa sa više
tipova navedenih u različitim case naredbama. Slično je switch case naredbi.
Jedina razlika je što slučajevi tipovi, a ne vrednosti kao kod  normalnog
switch-a.

Sintaksa za prekidač tipa je slična tvrdnji tipa. U sintaksi i.(T) za tvrdnju
tipa, tip  treba zameniti T ključnom reči type za prekidač tipa. Pogledajmo
kako ovo funkcioniše u programu ispod.
*/

func findType(i interface{}) {
	switch i.(type) {
	case string: // Ovaj kod je model i mora ostati iako nije po volji kompajleru
		fmt.Printf("I am a string and my value is %s\n", i.(string))
	case int:
		fmt.Printf("I am an int and my value is %d\n", i.(int))
	default:
		fmt.Printf("Unknown type\n")
	}
}

func ifaceTypeSwitch() {

	fmt.Println("\n --- Type Switch ---")

	findType("Naveen")
	findType(77)
	findType(89.98)
}

/*
U gornjem programu, findType funkcija koristi switch type da proveri tip
interfejsa i.(type). Ovo je slično tvrdnji tipa, ali se koristi u switch
naredbi.
3333
Svaka od case naredbi upoređuje konkretni tip interfejsa i sa određenim tipom.
Ako se bilo koja case poklapa, ispisuje se  odgovarajuća naredba.

Ovaj program štampa,

	>> I am a string and my value is Naveen
	>> I am an int and my value is 77
	>> Unknown type

Takođe je moguće uporediti tip sa interfejsom. Ako imamo tip i ako taj tip
implementira interfejs, moguće je uporediti ovaj tip sa interfejsom koji on
implementira.

Hajde da napišemo program radi veće jasnoće.
*/

type Describer interface {
	Describe()
}

type Person2 struct {
	name string
	age  int
}

func (p Person2) Describe() { // Implementacija Describe metode za Person2 tip
	fmt.Printf("%s is %d years old\n", p.name, p.age)
}

func findType2(i interface{}) {
	switch v := i.(type) {
	case Describer:
		v.Describe()
	default:
		fmt.Printf("unknown type\n")
	}
}

func ifaceTypeSwitch2() {

	fmt.Println("\n --- Type Switch 2 ---")
	p := Person2{
		name: "Radosav R",
		age:  65,
	}
	findType2(p)
	findType2("Radosav")
}

/*
U gornjem programu, Person2 struktura implementira Describer interfejs. U case
izjavi tip v se upoređuje sa Describer tipom interfejsa i pošto struktura Person2
implementira Describer, ovaj slučaj zadovoljen i Describe() metoda se poziva.

Ovaj program štampa

	>> Radosav R is 65 years old
	>> unknown type

Implementacija interfejsa pomoću pointer prijemnika vs vrednosnih prijemnika
---------------------------------------------------------------------------
Svi interfejsi o kojima smo do sada govorili implementirani su korišćenjem
vrednosnih prijemnika. Takođe je moguće implementirati interfejse korišćenjem
pointer prijemnika.

Postoji jedna suptilnost koju treba napomenuti pri implementaciji interfejsa
korišćenjem pointer prijemnika. Hajde da to razumemo koristeći sledeći program:
*/

type Describer2 interface {
	Describe()
}

type Person3 struct {
	name string
	age  int
}

func (p Person3) Describe() { //implemented using value receiver
	fmt.Printf("%s is %d years old\n", p.name, p.age)
}

type Address struct {
	state   string
	country string
}

func (a *Address) Describe() { //implemented using pointer receiver
	fmt.Printf("%s is in %s", a.state, a.country)
}

// ifacePointerReceiver funkcija
// Ova funkcija pokazuje kako interfejsi mogu implementirati metode koristeći
// i vrednosne i pointer prijemnike. Ovo je korisno za razumevanje kako se
// interfejsi ponašaju u Go-u kada se koriste različiti tipovi prijemnika.
func ifacePointerReceiver() {

	fmt.Println("\n --- Interfaces with pointer receiver ---")

	var i1 Describer2
	p1 := Person3{"Sam", 25}
	i1 = p1
	i1.Describe()

	p2 := Person3{"James", 32}
	i1 = &p2
	i1.Describe()

	var i2 Describer
	addr := Address{"Washington", "USA"}

	// compilation error if the following line is uncommented
	// ./prog.go:45:7: cannot use a (variable of type Address) as Describer value
	// in assignment: Address does not implement Describer
	// (method Describe has pointer receiver)
	//i2 = addr

	// Ovo radi jer Describer interfejs kao pointer prijemnik adrese
	i2 = &addr
	i2.Describe()
	fmt.Println()
}

/*
Kao što smo već naučili tokom naše diskusije o metodama, metode sa vrednosnim
prijemnicima prihvataju i pointere i vrednosne prijemnike. Legalno je pozvati
metodu sa vrednosnim prijemnikom na bilo čemu što je vrednost ili čija se
vrednost može dereferencirati.

U našem slučaju, vrednost p1 tipa Person3 dodeljena je interfejsu i1. Person3
implementira Describer2 interfejs i stoga je sve Ok. Slično tome, &p2 je adresa
vrednosti tipa Person3 i dodeljena je interfejsu i1, po gornjem stavu kompajler
će je dereferencirati i ovde je sve Ok.

Struktura Address implementira Describer2 interfejs koristeći pointer prijemnik.
Ako linija sa i2 = addr gornjeg programa nije komentarisana, dobićemo grešku
pri kompajliranju:

>> ./prog.go:45:7: cannot use a (variable of type Address) as Describer2 value
>>	in assignment: Address does not implement Describer2 (method Describe has
>> pointer receiver).

To je zato što tip Address implementira Describer2 interfejs korišćenjem pointer
prijemnika a mi pokušavamo da dodelimo addr koji je vrednosni tip, a nije
implementiran interfejs Describer2 sa vrednosnim prijemnikom.

Ovo će vas definitivno iznenaditi, jer smo ranije saznali da metode sa
pointer prijemnicima pointera prihvataju i pointer i vrednosne prijemnike.

Zašto onda kod u i2 = addr ne uspeva?

Razlog je taj što je legalno pozivati pointer ili vrednosni metod koji se
odnosi na bilo šta što je već pointer ili čija se adresa može uzeti.

Međutim, kada se radi o interfejsima, to ne važi. Interfejsi su samo skup
potpisa metoda i ne mogu se direktno koristiti sa vrednostima koje nisu
implementirale interfejs.

To znači da se konkretna vrednost sačuvana u interfejsu ne može direktno
adresirati.

Dakle, iako je addr tipa Address, koji implementira Describer2 interfejs,
konkretna vrednost sačuvana u interfejsu i2 nije adresabilna. To znači da
kompajler ne može uzeti adresu iz linije sa kodom i2 = addr, jer i2 očekuje
pointer prijemnik, a addr je vrednost tipa Address.

Međutim, ako uradimo i2 =&addr, program se uspešno kompajlira. Ostatak programa
je sam po sebi razumljiv. Ovaj program će ispisati,

	>> Sam is 25 years old
	>> James is 32 years old
	>> Washington is in USA

Implementacija više interfejsa
------------------------------
Tip može implementirati više od jednog interfejsa. Pogledajmo kako se to radi
u sledećem programu.
*/

type SalaryCalc interface {
	DisplaySalary()
}

type LeaveCalc interface {
	CalculateLeavesLeft() int
}

type Employee struct {
	firstName   string
	lastName    string
	basicPay    int
	pf          int
	totalLeaves int
	leavesTaken int
}

func (e Employee) DisplaySalary() {
	fmt.Printf("%s %s has salary $%d",
		e.firstName, e.lastName, (e.basicPay + e.pf))
}

func (e Employee) CalculateLeavesLeft() int {
	return e.totalLeaves - e.leavesTaken
}

func ifaceImplementMoreInterfaces() {

	fmt.Println("\n --- Types with more than one interface ---")

	e := Employee{
		firstName:   "Naveen",
		lastName:    "Ramanathan",
		basicPay:    10000,
		pf:          200,
		totalLeaves: 30,
		leavesTaken: 5,
	}

	var s SalaryCalc = e
	s.DisplaySalary()

	var l LeaveCalc = e
	fmt.Println("\nLeaves left =", l.CalculateLeavesLeft())
}

/*
U gorenjem programu, tip Employee implementira dva interfejsa SalaryCalc i
LeaveCalc.

Nadalje dodeljujemo promenljivu e tipa Employee promenljivoj s tipa SalaryCalc
interfejsa, a potom dodeljujemo istu promenljivu e promenljivoj l tipa LeaveCalc.
Ovo je moguće jer e tipa Employee koji implementira i SalaryCalc i LeaveCalc
interfejse.

Ovaj program štampa,

	>> Naveen Ramanathan has salary $10200
	>> Leaves left = 25

Ugrađivanje interfejsa
----------------------
Iako Go ne podržava nasleđivanje, ugrađivanje interfejsa je moguće. Moguće je
kreirati novi interfejs ugrađivanjem drugih interfejsa.

Da vidimo kako se ovo radi.
*/

type SalaryCalcu2 interface {
	DisplaySalary()
}

type LeaveCalc2 interface {
	CalculateLeavesLeft() int
}

type EmployeeOperations interface {
	SalaryCalcu2
	LeaveCalc2
}

type Employee2 struct {
	firstName   string
	lastName    string
	basicPay    int
	pf          int
	totalLeaves int
	leavesTaken int
}

func (e Employee2) DisplaySalary() {
	fmt.Printf("%s %s has salary $%d", e.firstName, e.lastName, (e.basicPay + e.pf))
}

func (e Employee2) CalculateLeavesLeft() int {
	return e.totalLeaves - e.leavesTaken
}

func ifaceEmbedid() {

	fmt.Println("\n --- Embedded interfaces ---")

	e := Employee2{
		firstName:   "Radosav",
		lastName:    "Radovanović",
		basicPay:    10000,
		pf:          200,
		totalLeaves: 30,
		leavesTaken: 5,
	}

	var empOp EmployeeOperations = e
	empOp.DisplaySalary()
	fmt.Println("\nLeaves left =", empOp.CalculateLeavesLeft())

	if v, ok := empOp.(Employee2); ok {
		// Tvrdnja tipa da biste dobili osnovne vrednosti Employee2.Ovo je
		// korisno ako želite da pristupite specifičnim poljima Employee2 ili
		// ako želite da pozovete metode koje nisu deo interfejsa
		// EmployeeOperations.

		// Napomena: Ovo nije neophodno za metode definisane u interfejsu
		// EmployeeOperations za tip Employee2.

		fmt.Printf("Employee2 struct:\n firstName: %s type is: %T\n",
			v.firstName, v.firstName)
	}
	fmt.Println()
}

/*
EmployeeOperations interfejs je kreiran ugrađivanjem interfejsa SalaryCalcu2 i
LeaveCalc2. Za bilo koji tip se kaže da implementira EmployeeOperations
interfejs ako pruža definicije metoda za metode prisutne u interfejsima
SalaryCalcu2 i LeaveCalc2.

Struktura Employee2 implementira EmployeeOperations interfejs jer pruža
definiciju za obe metode DisplaySalary i CalculateLeavesLeft.

Vrednost tipa Employee2 je dodeljena empOp vrednosti tipa interfejsa
EmployeeOperations. U naredna dva reda, metode DisplaySalary() i
CalculateLeavesLeft() se pozivaju na empOp.

!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
!!! Obratite pažnju, funkcije iz grupe Print..., za varijadičke parametre
!!! prihvaju konkretne vrednosti osnovnih tipova, ali ne i interfejs, jer
!!! interfejs nema vrednost kojoj se može direktno pristupiti.
!!!
!!! Ako želite da koristite interfejs sa funkcijama koje očekuju konkretne
!!! vrednosti, potrebno je koristiti tvrdnju tipa ili konverziju tipa da
!!! biste dobili osnovnu vrednost interfejsa.
!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

Ovaj program će štampati:

	>> Naveen Ramanathan has salary $5200
	>> Leaves left = 25

Nulta vrednost interfejsa
-------------------------
Nulta vrednost interfejsa je nil. nil interfejs ima i svoju osnovnu vrednost i
konkretan tip iako je nil.
*/

type Describer3 interface {
	Describe()
}

func ifaceNilInterface() {

	fmt.Println("\n --- Nil interface ---")

	var d1 Describer3

	fmt.Printf("d1 is nil and has type %T value %v\n", d1, d1)

}

/*
U gornjem programu >> d1 je nil i ovaj program će štampati:

	>> d1 is nil and has type <nil> value <nil>

Ako pokušamo da pozovemo metodu na nil interfejsu, program će paničiti jer nil
interfejs nema ni osnovnu vrednost niti konkretan tip.

// type Describer4 interface {
// 	Describe()
// }

// func main() {
// 	var d1 Describer4
// 	d1.Describe()
// }

Pošto je d1 u gornjem programu nil, ovaj program će paničiti sa greškom tokom
izvršavanja.

	>> panic: runtime error: invalid memory address or nil pointer dereference
	>> signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x4664b0]
	>> goroutine 1 [running]:
	>> 		main.main()
	>> 		/tmp/sandbox2797051632/prog.go:9 +0x10

To je zato što pokušavamo da pozovemo metodu Describe() na nil interfejsu

Da bismo izbegli ovu paniku, možemo proveriti da li je interfejs nil pre
pozivanja metode. Na primer:
*/

func ifaceNilInterfaceSafe() {

	fmt.Println("\n --- Nil interface safe ---")

	var d1 Describer3

	if d1 != nil {
		d1.Describe()
	} else {
		fmt.Println("d1 is nil, cannot call Describe()")
	}
}

func InterfaceFuncs() {

	fmt.Println("\n --- Interfaces ---")

	ifaceBegining()
	ifaceExt()
	ifaceInternRepr()
	ifaceEmpty()
	ifaceTypeAssertion()
	// ifaceTypeAssertion2()
	ifaceAssertOk()
	ifaceTypeSwitch()
	ifaceTypeSwitch2()
	ifacePointerReceiver()
	ifaceImplementMoreInterfaces()
	ifaceEmbedid()
	ifaceNilInterface()
	ifaceNilInterfaceSafe()
}
