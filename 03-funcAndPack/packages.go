/*
Go paketi
=========

Šta su paketi i zašto se koriste?
---------------------------------
Do sada smo videli Go programe koji imaju samo jednu datoteku sa glavnom
funkcijom i nekoliko drugih funkcija.U scenarijima u stvarnom svetu, ovaj
pristup pisanja izvornog koda u jednoj datoteci nije skalabilan. To postaje
nemoguće ponovno upotrebiti i održavati kod na ovaj način.Ovde su  paketi
korisni.

Paketi se koriste za organizovanje izvornog koda za bolju ponovnu upotrebu
i čitljivost. Paketi su kolekcija Go datoteka izvornog koda koji borave u istom
direktorijumu. Paketi pružaju decentralizaciju i stoga postaje lako održavanje
projekata.

Na primer, recimo da pišemo Fintech aplikaciju u Gou, a neke funkcionalnosti su
jednostavno izračunavanje kamate, složene kamate i i obračun otplate kredita.
Jedan jednostavan način organizovanja ove aplikacije je funkcionalni.
Mi možemo da kreiramo pakete "simpleinterest", "compoundinterest" i
"loanrepayment".Ako paket za kredit mora da izračuna jednostavni interes, to
jednostavno možemo učiniti uvozom "simpleinterest" paketa. Na ovaj način se kod
koristi ponovo.

Naučićemo pakete stvaranjem jednostavne aplikacije za utvrđivanje jednostavnog
interesa od glavnice, kamatnu stopu i vreme trajanja kredita u godinama.

Go paketi
----------
Strukturiraćemo kod na takav način da su sve funkcionalnosti u vezi sa prostim
interesom u "simpleinterest" paketu. Da biste to učinili moramo da kreiramo
prilagođeni paket koji će sadržavati funkciju za izračunavanje jednostavnog
interesa.Pre kreiranja prilagođenih paketa,  prvo moramo da razumemo Go module,
jer su potrebni moduli za kreiranje prilagođenih paketa.

Go modul nije ništa drugo nego kolekcija Go paketa.Ovo pitanje bi moglo doći
do vašeg uma: "Zašto nam trebaju moduli za kreiranje prilagođenog paketa?""
Odgovor je import put za prilagođeni paket koji stvaramo se izvodi iz imena Go
modula. Pored ovoga, sve ostale pakete trećih strana (kao što su izvorni kod
Github-a), zajedno sa njihovim verzijama. Našom upotrebom će upravljati go.mod
datoteka. Ova datoteka go.mod kreira se kada kreiramo novi modul. To ćete bolje
razumeti u sledećem odeljku.

Kreiranje go paketa
-------------------
Pokrenite naredbu u nastavku da biste kreirali direktorij nazvan learnpackage u
direktorijumu $GOPATH/src trenutnog korisnika.

	$ mkdir $GOPATH/src/learngo/learnpackage

Proverite da li ste unutar direktorijuma learnpackage, sa

	$ cd /go/src/learngo/learnpackage/.

Kreirajte "simpleinterset" prilagođeni paket
--------------------------------------------
Sve datoteke koje pripadaju paketu treba da budu postavljeni u svoje posebne
direktorijume. To je konvencija u Gou da se imenuje ovaj direktorijum sa istim
imenom kao i paket.

Kreirajmo direktorijum "simpleinterest" u osnovnom direktorijumu paketa
learnpackage. Komanda

	$ mkdir simpleinterest

će kreirati ovaj direktorijum simpleinterest za nas. Naredba

package packagename

specificira da određena .go datoteka pripada paketu "packagename". Ovo bi trebalo
da bude prva linija svake Go izvorne datoteke. Otuda sve datoteke unutar
direktorijuma "simpleinterest" trebalo bi da počnu sa linijom

package simpleinterest

jer sve pripadaju simpleinterest paketu.

Kreirajte datoteku simpleinterest.go unutar direktorijuma simpleinterest.

Sledeće će biti struktura direktorijuma naše aplikacije.

learngo
├── go.mod
└── learnpackage/
	├── learnpackages.go
	└── simpleinterest
	    └── simpleinterest.go

Dodajte sledeći kod u simpleinterest.go datoteku.

package simpleinterest

// calulate and return simple interest for principal p ($),
// interest rate r (%), and time of credit t(years)
func Calculate(p float64, r float64, t float64) float64 {

	interest := p * (r / 100) * t
	return interest
}

U gornjem kodu smo kreirali funkciju Calculate koja izračunava i vraća prosti
interes.Ova funkcija je samorazumljiva.

Imajte na umu da naziv funkcije Calculate započinje sa velikim slovom.Ovo je
neophodno i ubrzo ćemo objasniti zašto je to potrebno.

learnpackage paket and LearnPackage funkcija
--------------------------------------------
Na kraju learnpackage.go fajla postavićemo LearnPackage eksportovanu funkciju,
koju ćemo uvrstiti u main funkciju main paketa.

Sledeći korak je import paketa koga smo upravo stvorili i koristimo ga. Uvozimo
simpleinterest paket u learnpackages paket.

func learnPackagesPrint() {

	fmt.Println("Simple interest calculation")
}

Linija koda "package" navodi da ova datoteka pripada "learnpackages" paketu.
Izjava uvoza koristi se za uvoz postojećeg paketa simpleinterest.

Packagename.FunctionName() je sintaksa za pozivanje funkcije iz paketa.

U redu br.4, uvozimo fmt paket za upotrebu funkcije Println. fmt je standardni
paket i dostupan je kao deo Go standardne biblioteke. Tada funkcija koja štampa
jednostavan naslov se kompajlira iz direktorijuma learngo koristeći

	$ cd $GOPATH/src/learngo/

i unesite sledeće komande u terminal

	$ go build

Ako je sve dobro prošlo, naša binarna datoteka će biti tu i biće spremna za
izvršenje. Unesite naredbu:

	$ ./learngo

u terminal i videćete na kraju sledeći izlaz:

	>> Simple interest calculation

Ako ne razumete kako "go build" radi, posetite "Hello, world" tutorijal.

Uvoz prilagođenog "simpleinterest" paketa
-----------------------------------------
Da biste koristili prilagođeni paket, prvo ga moramo uvesti. Uvozni put je naziv
Go modula u kombinaciji direktorijuma u kome se paket boravi i naziva paketa.

U našem slučaju ime Go modula je "learngo", i paket je "learnpackages", a
podpaket je  "simpleinterest" udirektorijumu "simpleinterest".

├── learnpackage
│   └── simpleinterest

Dakle, linija uvoza je

import "learngo/learnpackage/simpleinterest".

U slučaju da imamo ovakvu strukturu direktorija

learnpackage
│   └── finance
│       └── simpleinterest

Tada bi bila izjava o uvozu bila:

import "learngo/learnpackage/finance/simpleinterest"

Sledeća funkcija je malo promenjena i računa i štampa interes.

func learnPackagesSimpleInterest() {

	p := 5000.0
	r := 10.0
	t := 1.0
	fmt.Println("For principal =", p, "rate =", r, "time of credit =", t)

	si := simpleinterest.Calculate(p, r, t)
	fmt.Println("Simple interest is", si)
}

Gornji kod uvozi "simpleinterest" paket i koristi funkciju "Calculate" da bi se
dobio prosti interes.

Paketi iz standardne Go biblioteke ne trebaju prefiks imenu modula  i otuda fmt
radi bez prefiksa modula. Kada se aplikacija pokrene, izlaz će biti

	>> For principal = 5000 rate = 10 time of credit = 1.5
	>> Simple interest calculation
	>> Simple interest is 500

Malo više o go build
--------------------

Sad kad razumemo kako paketi rade, vreme je da malo više razgovaramo o go build.
Go alati poput go build rade u kontekstu trenutnog direktorijuma.Shvatimo šta
to znači. Do sada smo pokrenuli go build iz direktorijuma $GOAPTH/src/learngo.
Ako pokušamo da ga izgradimo i pokrenemo iz bilo kog drugog direktorija, neće
uspeti.

Pokušajte sa cd $GOPATH/src i zatim pokrenite go build. Neće uspeti sa sledećom
greškom: package learngo is not in std (/usr/local/go/src/learngo)

Shvatimo razlog ove greške.Go build uzima opcionalno ime modula kao parametar
(u našem slučaju learngo) i pokušava da kompajlira main funkciju ako paket
main postoji u trenutnom direktorijumu iz kojeg se pokreće ili je u matičnom
direktorijumu i tako dalje.

Mi smo u direktorijumu $GOPATH/src i tu nije datoteka go.mod, a samim tim ne
može da se gradi, kompajler se žali da ne može da nađe modul learngo.

Kada pređemo na sa cd $GOPATH/src/learngo, go.mod postoji sa imenom modula.
Zato, Go će izgraditi learngo iz $GOPATH/src/learngo direktorijuma.

Ali do sada smo samo koristili go build i nismo odredili ime paketa. Ako nije
navedeno ni jedno ime paketa, go build će podrazumevati ime modula u trenutnom
radnom direktorijumu.Zato će go build bez imena paketa uraditi izgradnu. Dakle,
sledeće 3 komande su ekvivalentne kada se pokreću u $GOPATH/src/learngo:

	$ go build
	$ go build .
	$ go build learngo

Takođe sam napomenuo da go build ima mogućnost rekurzivnog pretraživanja
matičnog direktorija za go.mod datotekom. Proverimo da li to funkcioniše:

	$ cd $GOPATH/src/learnpackage/simpleinterest/

Gornja naredba će nas odvesti u direktorijum simpleinterest.Iz tog direktorija

	$ go build learnpackage

go build će uspešno pronaći datoteku go.mod u matičnom direktorijumu learngo
koji je definisao modul sa go.mod i stoga radi :).

Takođe je moguće promeniti ime izlazne binarne datoteke pomoću kada radimo sa
go buid. Pomerite se na $GOPATH/src/learngo i unesite:

	$ go build -o fintechapp

-o argument se koristi za određivanje imena izlazne binarne datoteke.U ovom
slučaju će se izgraditi binarna datoteka sa imenom fintechapp.

Pokrenite sa ./fintechapp i binarni fajl će se uspešno pokrenuti.

Izvezena imena
--------------
Kapitalizirali smo funkciju Calculate u simpleinterest paketu. Ovo ima posebno
značenje u Go. Bilo koja promenljiva ili funkcija koja počinje velikim slovom
se izvozi. Samo izvezenim funkcijama i promenljivim može se pristupiti iz drugih
paketa. U našem slučaju želimo da pristupimo Calculate funkciji iz main paketa.
Otuda je njeno ime kapitalizovano.

Ako se naziv funkcije promeni iz Calculate u calculate i pokušate da pozovete
simpleinterest.calculate u main.go, compajler će prijaviti grešku:

	$ # learnpackage
	$ ./main.go:13:8: undefined: simpleinterest.calculate

Dakle, ako želite da pristupite funkciji izvan paketa, njeno ime mora da bude
kapitalizovano.

init funkcija
-------------
Svaki paket u Go-u može da sadrži init funkciju. Funkcija "init" ne sme da ima
nikakve tipove povratne vrednosti i ne sme da ima parametre. Funkcija init ne
može se nazvati eksplicitno u našem izvornom kodu. Poziviće se automatski kada
se paket inicijalizira. Funkcija "init" ima sledeću sintaksu:

func init() {

}

Funkcija init može se koristiti za obavljanje zadataka inicijalizacije i takođe
se može koristiti za verifikaciju ispravnosti programa pre nego što se izvršenje
pokrene.

Redosled inicijalizacije paketa je sledeći:

	* Promenljive nivoa paketa prvo se inicijalizuju
	* Init funkcija se poziva sledeća. Paket može imati više init funkcija
	  (bilo u jednoj datoteci ili distribuirano u više datoteka) i one se pozivaju
	  po redosledu u kome su predstavljeni kompajleru.
	* Ako paket uvozi druge pakete, uvezeni paketi su prvo inicijalizovani.
	* Paket će se inicijalizovati samo jednom čak i ako se uvozi iz više paketa.

Napravimo neke modifikacije naše aplikacije da bismo razumeli init funkcije.

Da bismo započeli, dodajmo init funkciju u simpleinterest.go datoteku.

// init function added
func init() {

	fmt.Println("Simple interest package initialized")
}

// Calculate calculates and returns the simple interest for
// principal p ($), rate of interest r (%) and for time duration t (years)
func Calculate(p float64, r float64, t float64) float64 {

	interest := p * (r / 100) * t
	return interest
}

Dodali smo jednostavnu init funkciju koja samo štampa da je paket simpleinterest
inicijalizovan.

Sada izmenimo learnpackages paket.Znamo da principal, rate interest i vreme
trajanja kredita treba da budu veći od nule prilikom izračunavanja jednostavnog
interesa.Ovu proveru definišemo pomoću init funkcija i varijabli nivoa paketa u
learnpackages.go datoteci.

Modifikujte learnapckages.go na sledeći način,
*/

package funcs

import (
	"fmt"
	"learngo/03-funcAndPack/simpleinterest" //importing custom package
	"log"
)

var p, r, t = 5000.0, 10.0, 1.0

// init function to check if p, r and t are greater than zero
func init() {

	fmt.Println("Packages package initialized")

	if p < 0 {
		log.Fatal("Principal is less than zero")
	}

	if r < 0 {
		log.Fatal("Rate of interest is less than zero")
	}

	if t < 0 {
		log.Fatal("Duration is less than zero")
	}
}

func LearnPackages() {

	fmt.Println("\n --- Packages ---")
	fmt.Println("Simple interest calculation")

	si := simpleinterest.Calculate(p, r, t)
	fmt.Println("Simple interest is", si)
}

/*
Sledeće su promene u learnpackages paketu:

p, r i t promenljive se premeštaju na nivo paketa sa nivoa funkcije. Dodata je
init funkcija. Init funkcija štampa log i prekida izvršenje programa ako je p
ili r ili t manje od nula koristeći funkciju log.fatal. Obratite pažnju da smo
uvezli log paket.

Redosled inicijalizacije je sledeći:

	1.	Uvezeni paketi su prvo inicijalizirani. Otuda je prvo simpleinterest
		paket inicijaliziran i pokrenuta je njegova init metoda.
	2. 	Varijable nivoa learnpackages paketa p, r i t inicijalizirani su sledeće.
	3. 	Funkcija init learnpackages paketa se poziva.
	4. 	Poziva se LearnPackages funkcija.

Ako pokrenete program, dobićete sledeći izlaz:

	>> Simple interest package initialized
	>> Main package initialized
	>> Simple interest calculation
	>> Simple interest is 500

Kao što se očekivalo, init funkcija paketa simpleinterest se poziva prva praćena
inicijalizacijom varijabli main paketa r, p i r. Sledeće se poziva init funkcija
main paketa. Proverava da li su p, r i t manji od nule i napušta program ako je
stanje true. Mi ćemo detaljno saznati o if izjavi u odvojenom tutorialu.Za sada
možete pretpostaviti da sa if p < 0 proveravamo da li je p manji od 0 i ako
jeste, program će se prekinuti. Napisali smo slično za r i t. U ovom slučaju,
svi ovi uslovi su false i izvršenje programa se nastavlja. Konačno se poziva
main funkcija.

Hajde da izmenimo ovaj program malo da naučimo upotrebu init funkcije.

Promenite liniju

var p, r, t = 5000.0, 10.0, 1.0

u main.go u

var p, r, t = -5000.0, 10.0, 1.0

Inicijalizirali smo p na negativnu vrednost.

Ako pokrenete aplikaciju, videćete:

Simple interest package initialized

	>> Simpleinterest package initialized
	>> Learnpackages package initialized
	>> 2025/05/31 01:24:08 Principal is less than zero

p je negativano. Otuda kada se init funkcija pokreće, program prestaje nakon
što je odštampano da je principal manje od nule.

Upotreba praznog identifikatora u import sekciji
------------------------------------------------
Ilegalno je uvesti paket a da ga ne koristite u kodu.Kompajler će se žaliti ako
to učinite. Razlog za to je izbegavanje uvoza neiskorištenih paketa koji će
značajno povećati vreme kompilacije. Zamenite kod u main.go sa sledećim:

package main

import (
    "learnpackage/simpleinterest"
)

func main() {

}

Gornji program će biti greška

# learnpackage
./main.go:4:2: imported and not used: "learnpackage/simpleinterest"

Ali prilično je uobičajeno uvoziti pakete kada je aplikacija pod aktivnim
razvojem i kasnije ih koristiti negde u kodu ako ne odmah.Prazan identifikator
nas štedi u tim situacijama.

Greška u gornjem programu može se ugušiti sledećim kodom,

package main

import (
    "learnpackage/simpleinterest"
)

var _ = simpleinterest.Calculate

func main() {

}

Linija

var _ = simpleinterest.Calculate

prigušuje grešku. Treba da pratimo ove prigušivača grešaka i uklonimo ih,
uključujući uvezeni paket na kraju razvoja aplikacije ako se paket ne koristi.
Otuda se preporučuje pisanje prigušenja grešaka na nivou paketa neposredno
nakon import izjave.

Ponekad moramo da uvezemo paket samo da bismo bili sigurni da se inicijalizacija
odvija iako ne moramo da koristimo nikakvu funkciju ili promenljivu paketa.
Na primer, možda ćemo morati da osiguramo da se init funkcija paketa
simpleinterest poziva iako ne planiramo da koristimo taj paket u našem kodu.
U ovom slučaju se može koristiti i prazan identifikator _  kao što je prikazano
u nastavku.

package main

import (
	_ "learnpackage/simpleinterest"
)

func main() {

}

Pokretanje gornjeg programa će se inicijalizirati simpleinterest. Uspešno smo
inicijalizovali simpleinterest paket iako se ne koristi nigde u kodu.
*/
