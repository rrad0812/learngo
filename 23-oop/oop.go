/*
OOP
===

Go nije objektno orijentisan programski jezik.

	Iako Go poseduju tipove i metode i omogućava objektno orijentisan stil
	programiranja, nema hijerarhije tipa. Koncept "interfejsa" u Gou pruža
	drugačiji pristup konceptu nasleđivanja. Verujte da je jednostavan za
	upotrebu i na neki način uopšteniji.

	Postoje i načini za ugrađivanje tipova u druge tipove koje pružaju nešto
	analogno - ali ne identično - nasleđivanju.

	Štaviše, metode u Gou su uopštenije nego u C ++ ili Java: mogu se definisati
	za bilo kakve podatke, 	čak i ugrađene vrste kao što su obični,	"bezbojni"
	celi brojevi. Nisu ograničeni samo na strukture.

	U narednim tutorijalima ćemo razgovarati o tome kako se koncepti objektno
	orijentisanog programiranja mogu implementirati pomoću Goa. Neki od njih se
	prilično razlikuju u implementaciji u poređenju sa drugim objektno
	orijentisanim jezicima kao što je Java.

Strukture umesto klasa
----------------------
Go ne pruža klase, ali pruža strukture. Metode se mogu dodavati strukturama.
Ovo pruža ponašanje objedinjavanja podataka i metoda koje operišu na podacima
zajedno, slično klasi.

Hajde odmah da počnemo sa primerom radi boljeg razumevanja.

U ovom primeru ćemo kreirati prilagođeni paket jer pomaže da se bolje razume
kako strukture mogu biti efikasna zamena za klase.

Napravite poddirektorijum unutra "~/go/src/learngo" i nazovite ga "23-oop".

Napravite poddirektorijum "employee" unutar "23-oop". Unutar "employee"
direktorijuma kreirajte datoteku pod nazivom "employee.go".

Struktura direktorijuma bi izgledala ovako:

	├── go/src/learngo
	    └──23-oop
		    ├── oop.go
			└── employee
			    └── employee.go

Zamenite sadržaj "employee.go" sa sledećim kodom:

	package employee

	import (
		"fmt"
	)

	type Employee struct {
	    FirstName   string
	    LastName    string
	    TotalLeaves int
	    LeavesTaken int
	}

	func (e Employee) LeavesRemaining() {
	    fmt.Printf("%s %s has %d leaves remaining\n",
			e.FirstName, e.LastName, (e.TotalLeaves - e.LeavesTaken))
	}

U gornjem programu, izjava "package employee" određuje da ova datoteka pripada
paketu employee. Deklarišemo strukturu Employee i metodu pod nazivom
"LeavesRemaining", koja je dodata strukturi Employee. Ona izračunava i prikazuje
broj preostalih odsustava zaposlenog. Sada imamo strukturu i metodu koja radi na
strukturi objedinjenoj zajedno, slično klasi.

Napravite datoteku sa imenom "oop.go" (ova datoteka) unutar "23-oop"
direktorijuma.

Sada bi struktura direktorijuma izgledala ovako,

	├── go/src/learngo
	│   └── 23-oop
	│       ├── employee
	│       │   └── employee.go
	│       └── oop.go

Sadržaj "oop.go" je dat u nastavku:

	package oop

	import (
		"fmt"
		"learngo/23-oop/employee"
	)

	func oopExample() {
		fmt.Println("\n --- oopExample ---")

		e := employee.Employee{
			FirstName:   "Sam",
			LastName:    "Adolf",
			TotalLeaves: 30,
			LeavesTaken: 20,
		}
		e.LeavesRemaining()
	}

Uvozimo paket employee. Metod LeavesRemaining() strukture Employee se poziva.

Ovaj program ne može da se pokrene na Playground-u jer ima prilagođeni paket.
Ako pokrenete ovaj program na vašem lokalnom računaru, program će ispisati:

	>> Sam Adolf has 10 leaves remaining.

Funkcija New() umesto konstruktora
----------------------------------
Program koji smo napisali gore izgleda dobro, ali postoji jedan suptilni
problem. Da vidimo šta se dešava kada definišemo strukturu zaposlenih sa nultim
vrednostima. Zamenimo sadržaj "oop.go" sledećim kodom,

	package oop

	import "golearn/23-oop/employee"

	func oopNilStruct() {
		var e employee.Employee
		e.LeavesRemaining()
	}

Jedina promena koju smo napravili je kreiranje strukture Employee nulte
vrednosti. Ovaj program će ispisati,

	>> has 0 leaves remaining

Kao što vidite, Employee promenljiva kreirana sa nultom vrednošću je
neupotrebljiva. Nema važeće ime, prezime, a takođe nema ni važeće podatke o
odsustvu.

U drugim OOP jezicima kao što je Java, ovaj problem se može rešiti korišćenjem
konstruktora. Validan objekat se može kreirati korišćenjem parametrizovanog
konstruktora.

Go ne podržava konstruktore. Ako nulta vrednost tipa nije upotrebljiva, zadatak
programera je da deeksportuje tip kako bi sprečio pristup iz drugih paketa, a
takođe i da obezbedi funkciju sa imenom NewT(parameters) koja inicijalizuje tip
T sa potrebnim vrednostima.

Konvencija u Gou je da se funkcija koja kreira vrednost tipa T imenuje sa
NewT(parameters). Ovo će delovati kao konstruktor.

Ako paket definiše samo jedan tip, onda je konvencija u Gou da se ova funkcija
imenuje samo New(parameters) umesto NewT(parameters).

Hajde da napravimo izmene u programu koji smo napisali tako da svaki put kada
se kreira Employee, on bude upotrebljiv.

Prvi korak je da se deeksportuje Employee struktura i kreira funkcija New()
koja će kreirati novi Employee. Zamenite kod employee.go sa sledećim,

	package employee

	import (
		"fmt"
	)

	type employee struct {
		firstName   string
		lastName    string
		totalLeaves int
		leavesTaken int
	}

	func New(firstName string, lastName string, totalLeave int, leavesTaken int) employee {
		e := employee {firstName, lastName, totalLeave, leavesTaken}
		return e
	}

	func (e employee) LeavesRemaining() {
		fmt.Printf("%s %s has %d leaves remaining\n",
			e.firstName, e.lastName, (e.totalLeaves - e.leavesTaken))
	}

Ovde smo napravili neke važne izmene. Početno slovo strukture Employee je
napisano malim slovom, odnosno promenili smo "type Employee struct" ga na "type
employee struct". Time smo uspešno opozvali izvoz employee strukture i sprečili
pristup iz drugih paketa.

Dobra je praksa da se sva polja neeksportovane strukture takođe opozovu na
izvoz, osim ako ne postoji posebna potreba za njihovim izvozom. Pošto nam nije
potrebno da pristupamo poljima strukture employee bilo gde van employee paketa,
opozvali smo izvoz svih polja.

Promenili smo imena polja u skladu sa tim u LeavesRemaining() metodi.

Pošto employee nije izvezen, nije moguće kreirati vrednosti tipa Employee iz
drugih paketa. Stoga, dajemo izvezenu New funkciju koja uzima potrebne parametre
kao ulaz i vraća novokreiranog zaposlenog.

Ovaj program još uvek treba da se promeni da bi radio, ali hajde da ga pokrenemo
da bismo razumeli efekat dosadašnjih promena. Ako se ovaj program pokrene, doći
će do greške pri kompajliranju,

	>> # oop
	>> ./main.go:6:8: undefined: employee.Employee

To je zato što imamo neeksportovanu funkciju employee u employee paketu i njoj
se ne može pristupiti iz oop paketa. Zbog toga kompajler izbacuje grešku da
ovaj tip nije definisan u main.go.

Savršeno. Baš ono što smo želeli. Sada nijedan drugi paket neće moći da kreira
nultu vrednost employee. Uspešno smo sprečili kreiranje neupotrebljive vrednosti
strukture employee. Jedini način da se sada kreira employee je da se koristi
New funkcija.

Zamenite sadržaj "oop.go" sa sledećim kodom,
*/
package oop

import (
	"fmt"
	"learngo/23-oop/employee"
)

func oopExampleNew() {
	fmt.Println("\n --- oopExampleNew ---")
	e := employee.New("Sam", "Adolf", 30, 20)
	e.LeavesRemaining()
}

/*
Jedina promena u ovoj datoteci je u kreiranju novog zaposlenog tako što smo
funkciji New prosledili potrebne parametre.

Pokretanje ovog programa će ispisati,

	>> Sam Adolf has 10 leaves remaining

Dakle, možete razumeti da iako Go ne podržava klase, strukture se mogu efikasno
koristiti umesto klasa, a metode New(parameters) se mogu koristiti umesto
konstruktora.

Kompozicija umesto nasleđivanja
===============================

Kompozicija ugrađivanjem struktura
----------------------------------
Kompozicija se u Go-u može postići ugrađivanjem jednog tipa strukture u drugi.
Blog post je savršen primer kompozicije. Svaki blog post ima naslov, sadržaj i
podatke o autoru. Ovo se može savršeno predstaviti pomoću kompozicije.

Prvo kreirajmo author strukturu.

	type author struct {
		firstName string
		lastName  string
		bio       string
	}

	func (a author) fullName() string {
		return fmt.Sprintf("%s %s", a.firstName, a.lastName)
	}

U gornjem isečku koda, kreirali smo author strukturu sa poljima firstName,
lastName i bio. Takođe smo dodali metodu fullName() sa author vrednosnim tipom
prijemnika koja vraća puno ime autora.

Sledeći korak bi bio kreiranje blogPost strukture.

	type blogPost struct {
		title     string
		content   string
		author
	}

	func (b blogPost) details() {
		fmt.Println("Title: ", b.title)
		fmt.Println("Content: ", b.content)
		fmt.Println("Author: ", b.author.fullName())
		fmt.Println("Bio: ", b.author.bio)
	}

Struktura blogPost ima polja title, content. Takođe ima ugrađeno anonimno polje
author. Ovo polje označava da je u blogPost strukturu ugrađena struktura author.
Sada blogPost struktura ima pristup svim poljima i metodama strukture author.
Takođe smo dodali details() metodu blogPost strukturi koja ispisuje naslov, sadržaj,
puno ime i biografiju autora.

Kad god je jedno polje strukture ugrađeno, Go nam daje mogućnost da pristupimo
ugrađenim poljima kao da su deo spoljašnje strukture. To znači da se
p.author.fullName() može zameniti sa p.fullName().

Stoga details() metod može prepisati kao što je prikazano ispod,

	func (p blogPost) details() {
		fmt.Println("Title: ", p.title)
		fmt.Println("Content: ", p.content)
		fmt.Println("Author: ", p.fullName())
		fmt.Println("Bio: ", p.bio)
	}

Sada kada imamo strukture author i blogPost spremne, hajde da završimo ovaj
program kreiranjem blog posta.
*/
type author struct {
	firstName string
	lastName  string
	bio       string
}

func (a author) fullName() string {
	return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

type blogPost struct {
	title   string
	content string
	author
}

func (b blogPost) details() {
	fmt.Println("Title: ", b.title)
	fmt.Println("Content: ", b.content)
	fmt.Println("Author: ", b.fullName())
	fmt.Println("Bio: ", b.bio)
}

func compExample() {
	fmt.Println("\n --- compExample ---")

	author1 := author{
		"Radosav",
		"Radovanović",
		"Golang Enthusiast",
	}

	blogPost1 := blogPost{
		"Inheritance in Go",
		"Go supports composition instead of inheritance",
		author1,
	}

	blogPost1.details()
}

/*
Glavna funkcija gore navedenog programa kreira novog autora. Nova objava se
kreira u sa ugrađivanjem author1. Ovaj program ispisuje,

	>> Title:  Inheritance in Go
	>> Content:  Go supports composition instead of inheritance
	>> Author:  Naveen Ramanathan
	>> Bio:  Golang Enthusiast

Ugrađivanje isečka
--------------------------------
Možemo ovaj primer odvesti korak dalje i napraviti veb stranicu koristeći
isečak blog postova :).

Hajde prvo da definišemo website strukturu. Dodajmo sledeći kod iznad glavne
funkcije postojećeg programa i pokrenete ga.

	type website struct {
	    []blogPost
	}

	func (w website) contents() {
	    fmt.Println("Contents of Website\n")
	    for _, v := range w.blogPosts {
	        v.details()
	        fmt.Println()
	    }
	}

Kada pokrenete gornji program nakon dodavanja gornjeg koda, kompajler će se
žaliti na sledeću grešku,

	>> main.go:31:9: syntax error: unexpected [, expecting field name or
	   embedded type

Ova greška ukazuje na ugrađeni isečak []blogPost. Razlog je taj što nije moguće
anonimno ugraditi isečak u strukturu. Potrebno je ime polja. Zato hajde da
ispravimo ovu grešku i usrećimo kompajler.

	type website struct {
		blogPosts []blogPost
	}

Dodao sam polje blogPosts koje je tipa []blogPosts. Sada hajde da modifikujemo
glavnu funkciju i kreiramo nekoliko objava za naš novi veb sajt.
*/
type website struct {
	blogPosts []blogPost
}

func (w website) contents() {
	fmt.Println("\n Contents of Website")
	for _, v := range w.blogPosts {
		v.details()
		fmt.Println()
	}
}

func compWithSlices() {

	fmt.Println("\n --- compWithSlices ---")
	author1 := author{
		"Naveen",
		"Ramanathan",
		"Golang Enthusiast",
	}
	blogPost1 := blogPost{
		"Inheritance in Go",
		"Go supports composition instead of inheritance",
		author1,
	}
	blogPost2 := blogPost{
		"Struct instead of Classes in Go",
		"Go does not support classes but methods can be added to structs",
		author1,
	}
	blogPost3 := blogPost{
		"Concurrency",
		"Go is a concurrent language and not a parallel one",
		author1,
	}
	w := website{
		blogPosts: []blogPost{blogPost1, blogPost2, blogPost3},
	}
	w.contents()
}

/*
U gornjoj funkciji, kreirali smo autora author1 i tri objave post1, post2 i
post3. Konačno, kreirali smo veb stranicu w ugrađivanjem ove 3 objave i
prikazivanjem sadržaja u sledećem redu.

Ovaj program će ispisati,

	>> # Contents of Website
	>>
	>> Title:  Inheritance in Go
	>> Content:  Go supports composition instead of inheritance
	>> Author:  Naveen Ramanathan
	>> Bio:  Golang Enthusiast
	>>
	>> Title:  Struct instead of Classes in Go
	>> Content:  Go does not support classes but methods can be added to structs
	>> Author:  Naveen Ramanathan
	>> Bio:  Golang Enthusiast
	>>
	>> Title:  Concurrency
	>> Content:  Go is a concurrent language and not a parallel one
	>> Author:  Naveen Ramanathan
	>> Bio:  Golang Enthusiast

Polimorfizam
==================

Polimorfizam korišćenjem interfejsa
-----------------------------------
Za svaki tip koji pruža definiciju za sve metode interfejsa kaže se da
implicitno implementira taj interfejs. Ovo će biti jasnije kada uskoro
budemo razmatrali primer polimorfizma.

Promenljiva tipa interfejs može da sadrži vrednost bilo kog tipa koji
implementira taj interfejs. Ovo svojstvo interfejsa se koristi za postizanje
polimorfizma u Go jeziku.

Hajde da razumemo polimorfizam u Go jeziku uz pomoć programa koji izračunava
neto prihod organizacije. Radi jednostavnosti, pretpostavimo da ova zamišljena
organizacija ima prihod od dve vrste projekata, fiksnog fakturisanja i vremena
i materijala . Neto prihod organizacije se izračunava kao zbir prihoda
od ovih projekata. Da bismo pojednostavili ovaj tutorijal, pretpostavićemo da
je valuta dolar i da se nećemo baviti centima. Biće predstavljeno pomoću int.

Hajde prvo da definišemo interfejs Income.

	type Income interface {
		calculate() int
		source() string
	}

Interfejs Income definisan gore sadrži dve metode "calculate()"" koja izračunava
i vraća prihod od izvora i "source()"" koje vraćaju naziv izvora.

Zatim, definišimo strukturu za "FixedBilling" tip projekta.

	type FixedBilling struct {
		projectName string
		biddedAmount int
	}

Projekat "FixedBilling" ima dva polja "projectName" koja predstavljaju naziv
projekta i "biddedAmount" iznos koji je organizacija ponudila za projekat.

Struktura "TimeAndMaterial" će predstavljati projekte tipa "TimeAndMaterial".

	type TimeAndMaterial struct {
		projectName string
		noOfHours  int
		hourlyRate int
	}

Struktura "TimeAndMaterial" ima tri polja, "projectName", "noOfHours" i
"hourlyRate".

Sledeći korak bi bio definisanje metoda na ovim tipovima struktura koje
izračunavaju i vraćaju stvarni prihod i izvor prihoda.

	func (fb FixedBilling) calculate() int {
		return fb.biddedAmount
	}

	func (fb FixedBilling) source() string {
		return fb.projectName
	}

	func (tm TimeAndMaterial) calculate() int {
		return tm.noOfHours * tm.hourlyRate
	}

	func (tm TimeAndMaterial) source() string {
		return tm.projectName
	}

U slučaju FixedBilling projekata, prihod je samo iznos ponuđen za projekat.
Stoga ga vraćamo iz calculate() metode FixedBilling tipa.

U slučaju TimeAndMaterial projekata, prihod je proizvod noOfHours i
hourlyRate. Ova vrednost se vraća iz calculate() metode sa tipom prijemnika
TimeAndMaterial.

Vraćamo naziv projekta kao izvor prihoda iz source() metode.

Pošto obe strukture FixedBilling i TimeAndMaterial pružaju definicije za
calculate() i source() metode interfejsa Income, obe strukture implementiraju
Income interfejs.

Hajde da deklarišemo "calculateNetIncome" funkciju koja će izračunati i ispisati
ukupan prihod.

	func calculateNetIncome(ic []Income) {
		var netincome int = 0
		for _, income := range ic {
			fmt.Printf("Income From %s = $%d\n", income.source(), income.calculate())
			netincome += income.calculate()
		}
		fmt.Printf("Net income of organization = $%d", netincome)
	}

Gore navedena funkcija prihvata isečak Income kao argument. Izračunava ukupan
prihod iteracijom kroz isečak Income i pozivanjem metode calculate() na svakoj
od njegovih stavki. Takođe prikazuje izvor prihoda pozivanjem metode source().
U zavisnosti od konkretnog tipa Income interfejsa, pozivaće se metode
calculate() i source() različitih konkretnih tipova. Time smo postigli
polimorfizam  u funkciji "calculateNetIncome".

U budućnosti, ako organizacija doda novu vrstu izvora prihoda, ova funkcija će
i dalje pravilno izračunati ukupan prihod bez ijedne izmene koda :).

Jedini deo koji je preostao u programu je glavna funkcija.

	func poliExample() {
		project1 := FixedBilling{projectName: "Project 1", biddedAmount: 5000}
		project2 := FixedBilling{projectName: "Project 2", biddedAmount: 10000}
		project3 := TimeAndMaterial{projectName: "Project 3", noOfHours: 160, hourlyRate: 25}
		incomeStreams := []Income{project1, project2, project3}
		calculateNetIncome(incomeStreams)
	}

U gornjoj funkciji smo kreirali tri projekta, dva tipa "FixedBilling" i jedan
tipa "TimeAndMaterial". Zatim, kreiramo isečak tipa "Income" sa ova 3 projekta.
Pošto svaki od ovih projekata ima implementiran "Income" interfejs, moguće je
dodati sva tri projekta u isečak tipa "Income". Konačno, pozivamo
"calculateNetIncome" funkciju i prosleđujemo ovaj isečak kao parametar.
Ona će prikazati različite izvore prihoda i prihod od njih.

Evo kompletnog programa za vašu referencu.
*/

type Income interface {
	calculate() int
	source() string
}

type FixedBilling struct {
	projectName  string
	biddedAmount int
}

type TimeAndMaterial struct {
	projectName string
	noOfHours   int
	hourlyRate  int
}

func (fb FixedBilling) calculate() int {
	return fb.biddedAmount
}

func (fb FixedBilling) source() string {
	return fb.projectName
}

func (tm TimeAndMaterial) calculate() int {
	return tm.noOfHours * tm.hourlyRate
}

func (tm TimeAndMaterial) source() string {
	return tm.projectName
}

func calculateNetIncome(ic []Income) {
	var netincome int = 0
	for _, income := range ic {
		fmt.Printf("Income From %s = $%d\n", income.source(), income.calculate())
		netincome += income.calculate()
	}
	fmt.Printf("Net income of organization = $%d\n", netincome)
}

func poliExample() {
	fmt.Println("\n --- poliExample ---")
	project1 := FixedBilling{projectName: "Project 1", biddedAmount: 5000}
	project2 := FixedBilling{projectName: "Project 2", biddedAmount: 10000}
	project3 := TimeAndMaterial{projectName: "Project 3", noOfHours: 160, hourlyRate: 25}
	incomeStreams := []Income{project1, project2, project3}
	calculateNetIncome(incomeStreams)
}

/*
Ovaj program će ispisati:

	>> Income From Project 1 = $5000
	>> Income From Project 2 = $10000
	>> Income From Project 3 = $4000
	>> Net income of organization = $19000

Dodavanje novog izvora prihoda gore navedenom programu
------------------------------------------------------
Recimo da je organizacija pronašla novi izvor prihoda putem reklama. Da vidimo
koliko je jednostavno dodati ovaj novi izvor prihoda i izračunati ukupan prihod
bez ikakvih promena u calculateNetIncome funkciji. Ovo postaje moguće
zahvaljujući polimorfizmu.

Prvo definišimo Advertisement tip i metode calculate() i source() nad tim
Advertisement tipom.

	type Advertisement struct {
		adName     string
		CPC        int
		noOfClicks int
	}

	func (a Advertisement) calculate() int {
		return a.CPC * a.noOfClicks
	}

	func (a Advertisement) source() string {
		return a.adName
	}

Tip Advertisement ima tri polja adName, CPC(cena po kliku) i noOfClicks
(broj klikova). Ukupan prihod od oglasa je proizvod CPC i noOfClicks.

Hajde da mainmalo modifikujemo funkciju kako bismo uključili ovaj novi tok prihoda.

	func poliExampleExt() {
		project1 := FixedBilling{projectName: "Project 1", biddedAmount: 5000}
		project2 := FixedBilling{projectName: "Project 2", biddedAmount: 10000}
		project3 := TimeAndMaterial{projectName: "Project 3", noOfHours: 160, hourlyRate: 25}
		bannerAd := Advertisement{adName: "Banner Ad", CPC: 2, noOfClicks: 500}
		popupAd := Advertisement{adName: "Popup Ad", CPC: 5, noOfClicks: 750}
		incomeStreams := []Income{project1, project2, project3, bannerAd, popupAd}
		calculateNetIncome(incomeStreams)
	}

Napravili smo dva oglasa, naime "bannerAd" i "popupAd". Isečak "incomeStreams"
sadrži dva oglasa koja smo upravo kreirali.

Evo dodatnog dela programa nakon dodavanja oglasa.
*/

type Advertisement struct {
	adName     string
	CPC        int
	noOfClicks int
}

func (a Advertisement) calculate() int {
	return a.CPC * a.noOfClicks
}

func (a Advertisement) source() string {
	return a.adName
}

func poliExampleExt() {
	fmt.Println("\n --- poliExampleExt ---")
	project1 := FixedBilling{projectName: "Project 1", biddedAmount: 5000}
	project2 := FixedBilling{projectName: "Project 2", biddedAmount: 10000}
	project3 := TimeAndMaterial{projectName: "Project 3", noOfHours: 160, hourlyRate: 25}
	bannerAd := Advertisement{adName: "Banner Ad", CPC: 2, noOfClicks: 500}
	popupAd := Advertisement{adName: "Popup Ad", CPC: 5, noOfClicks: 750}
	incomeStreams := []Income{project1, project2, project3, bannerAd, popupAd}
	calculateNetIncome(incomeStreams)
}

/*
Gore navedeni program će ispisati,

	>> Income From Project 1 = $5000
	>> Income From Project 2 = $10000
	>> Income From Project 3 = $4000
	>> Income From Banner Ad = $1000
	>> Income From Popup Ad = $3750
	>> Net income of organization = $23750

Primeti ste da nismo napravili nikakve izmene u "calculateNetIncome" funkciji
iako smo dodali novi tok prihoda. Radi jednostavno zbog polimorfizma. Pošto je
novi "Advertisement" tip takođe implementirao "Income" interfejs, mogli smo da
ga dodamo u "incomeStreams" isečak. "calculateNetIncome" funkcija takođe
radi bez ikakvih izmena jer može da poziva metode tipa "Advertisment",
calculate() i source().
*/

func OOPFunc() {

	fmt.Println("\n --- oopFunc ---")

	oopExampleNew()
	compExample()
	compWithSlices()
	poliExample()
	poliExampleExt()
}
