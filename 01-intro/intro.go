/*
Intro
=====

Uvod i instalacija
==================

Ovo je prvi tutorial u našoj seriji golang tutorijala.Ovaj tutorial pruža
Uvod za Go i takođe raspravlja o prednostima izbora prelazeći druge programiramske
jezike.Takođe ćemo naučiti kako instalirati Go na Mac OS, Windows i Linux.

Uvod
====

Poznat i kao Golang je otvorenog izvora, kompajliran i jako statički tipiziran
programski jezik koji je razvio Google.Ključni ljudi koji stoje iza stvaranja
Goa su Rob Pike, Ken Thompson i Robert Griesemer. Go je postao open source u
novembru 2009. godine.

Go je program za programiranje opšte namene sa jednostavnom sintaksom i podržan
robusnom standardnom bibliotekom.Jedna od ključnih područja u kojoj Go sjaji je
kreairanje visoko dostupnih i skalabilnih web aplikacija.Go se takođe može
koristiti za kreiranje aplikacija komandne linije, desktop aplikacija i čak i
mobilnih aplikacija.

Prednosti Goa
-------------
Zašto biste izabrali Go kao jezik na strani servera kada postoje tone
drugih jezika kao što su Python, Ruby, Nodejs ... koji rade isti posao.

Evo nekih od prednosti:

	Jednostavna sintaksa
	Sintaksa je jednostavna i sažeta i jezik se ne omamljuje sa nepotrebnim
	karakteristikama. To olakšava pisanje koda koji je čitljiv i održiv.

	Lako pisanje konkurentnih programa
	Konkurencija je svojstveni deo jezika.Kao rezultat, pisanje multithreaded
	programa je komad komad torte. To se postiže gorutinama i kanalima o kojima
	ćemo razgovarati u predstojećim tutorijalima.

	Kompajlirani jezik
	Go je kompajlirani jezik.Izvorni kod je kompajliran na izvršni binarni.
	Ovo je karika koja nedostaje u interpretiranim jezicima kao što je
	JavaScript.

	Brza kompilacija
	Kompajler Goa je dizajniran je da bude brz od samog početaka.

	Statičko linkovanje
	Go kompajler podržava statičko povezivanje.Ceo Go projekat može biti
	statički povezan u jednu veliku binarnu datoteku i lako se može rasporediti
	na serverima oblaka bez brige o zavisnosti.

	Go alati
	Alati zaslužuje poseban spominje u Go. Go golazi u paketu sa moćnim alatima
	koji pomažu programerima da napišu bolji kod. Evo najčešće korišćenih alata

    	gofmt - gofmt se koristi za automatsko formatiranje izvornog koda.
		Koristi	tabove za uvlačenje i praznine za poravnanje.
		vet - vet analizira kod i izveštava o mogućem sumnjivomm kodu.
		Sve što vet prijavi možda nije istinski problem, ali ima mogućnost
		da uhvati greške koje kompajler uvek ne prijavljuje.
		staticcheck - staticcheck se koristi za sprovođenje stalnih provera
		u kodu.

	Prikupljanje smeća
	Go koristi skeniranje radi prikupljanja smeća i otuda menadžment memorije
	je prilično automatski i programer ne treba da brine o upravljanju memorijom.
	Ovo takođe pomaže da se lako napiše konkurentne programe.

	Jednostavna specifikacija jezika
	Jezičke specifikacije su prilično jednostavne.Celokupni Go je dobro
	dokumentovan i možete ga koristiti da napišete sopstveni kompajler :)

	OpenSource
	Poslednje, ali ne najmanje bitno, Go je projekat otvorenog koda. Možete
	učestvovati i doprineti projektu.

Popularni proizvodi sagrađeni koristeći Go
------------------------------------------
	Google je razvio "Kubernetes" koristeći Go.

	"Docker", svetski poznata platforma za kontejnerizaciju razvijena na Gou.

	"Dropbox" je sve prebacio svoje po performansama kritične komponente iz
	Python-a u Go.

	Infoblox-ovi "Next Generation Networking Products" se razvijaju koristeći Go.

Instalacija
===========

Go se može instalirati na sve tri platforme Mac, Windows i Linux.
Može preuzeti binarne instalacije za odgovarajuću platformu sa
https://go.dev/dl/.

Mac OS
------
Preuzmite Mac OS Installer sa https://go.dev/dl/. Dva puta dodirnite da biste
započeli instalaciju.Sledite uputstva i to će za rezultat imati instaliran Go
u /us /local/go i dodaće direktorijum /usr/local/go/bin u PATH varijablu
okruženja. Na vama ostaje da dodate GOROOT=/usr/local/go varijablu okruženja.

Windows
-------
Preuzmite MSI Installer sa https://go.dev/dl/. Dva puta dodirnite da biste
započeli instalaciju i sledite uputstva.Ovo će instalirati Go na lokaciju
C:\go i dodati direktorijum C:\go\bin u PATH varijablu okruženja. Na vama je
da dodate GOROOT=c:\go varijablu okruženja.

Linux
-----
Preuzmite tar datoteku sa https://go.dev/dl/ i raspakujte je na /usr/local/go.
Dodajte /usr/local/go/bin u vašu PATH promenljivu okruženja i dodajte
GOROOT=/usr/local/go varijablu okruženja.

Provera instalacije
------------------------------
Da biste proverili da li je uspešno instaliran, unesite komandu "go version" u
Terminal i to će izlaziti instaliranu idu verziju.Evo izlaza
U mom terminalu.
*/

package intro

import "fmt"

func IAndI() {
	fmt.Println("\n --- Intro ---")

	fmt.Println("\n --- Verifying Go installation ---")
	fmt.Println("$go version")
	fmt.Println("$go version go1.19.2 linux/amd64")
}

/*
1.19.2 je bila najnovija verzija Goa kada pišem ovaj tutorijal.Ovo potvrđuje
da je Go pokrenut uspešno.U sledećem tutorialu ćemo napisati naš prvi
"Hello World" program u Go :)
*/

/*
Hello World
===========

Nema boljeg načina da se naučite programski jezik nego da uprljate ruke pišući
kod.Hajde da pišemo naš prvi GO program.

Postavljanje razvojnog okruženja
--------------------------------------

Prvo kreiramo direktorijume:

	$ mkdir ~/go
	$ mkdir ~/go/src
	$ mkdir ~/go/pkg
	$ mkdir ~/go/bin

a zatim promenljivu okruženja GOPATH=~/go. Ova promenljiva definiše naše
razvojno okruženje.

Kreirajmo direktorijum u kome želimo da napišemo naš "Hello World" program.
Otvorite terminal i pokrenite sledeću naredbu:

	$ mkdir ~/go/src/learngo/

Gornja naredba kreira direktorijum sa imenom "learngo" unutar /go/src
direktorijuma korisnika.

Stvaranje go modula
-------------------

Sjedeći korak je kreiranje modula, nazvaćemo ga isto learngo u ~/go/lerango
direktorijumu. Go moduli se koriste za praćenje zavisnosti naše instalacije
i njihove verzije. Detaljnije ćemo raspravljati o Go modulima kada naučimo
o paketima.

Pokrenite

	cd ~/go/src/learngo
	go mod init learngo

Ovo će stvoriti datoteku pod nazivom go.mod. Sledeće će biti odštampano nakon
pokretanja gornje komande:

Go: Create new go.mod: Modul learngo

Sadržaj datoteke go.mod je dat u nastavku:

"module learngo"

"go 1.21.0"

Naredba module u prvom redu određuje ime modula - learngo.
Sledeća linija go 1.21.0 označava da su datoteke u ovom modulu kompajlirane
koristeći go verziju 1.21.0

Helloworld
==========

Kreirajte datoteku pod nazivom main.go u learngo direktorijumu pomoću svog
omiljenog uređivača teksta sa sledećim sadržajem.*/

// package main

// import "fmt"

// func main() {

// 	fmt.Println("Hello World")
// }

/*
Konvencija je da datoteka koja sadrži main funkciju se zove main.go, ali i
druga imena rade.

Kompajliranje pokretanje Go programa
------------------------------------
Postoji nekoliko različitih načina za kompajliranje (izgradnju) i pokretanje
Go programa. Pogledajmo ih jedan po jedan.

1.	go install

	Prva metoda za pokretanje GO programa je korišćenje komande "go install".
	Hajde da odemo sa "cd" odemo u "learngo" direktorijum, koji smo malo pre
	kreirali.

	$ cd ~/go/src/learngo

	Sada pokrenite sledeću naredbu:

	$ go install

	Gornja naredba će kompajlirati program i pokušati da instalira (kopira)
	binarni fajl na lokaciju ~/go/bin.Ime binarne datoteke će biti ime Go
	modula.U našem slučaju to će biti ime learngo.

	Možda ćete naići na sledeću grešku kada ovo pokušate:

	Go install: Nema instalacione lokacije za $GOPATH/bin
	Za više detalja pogledajte: '$ go help GOPATH'

	Ono što je gornje greške zapravo znači, "go install" ne možete da pronađe
	lokaciju za instaliranje kompajliranog programa. Pa dajmo tu lokaciju.

	**************************************************************

	Kao rekapitualcija potrebne su nam sledeće promenljive okoline

		GOROOT=/usr/go
		GOPATH=~/go

	i sledeći dodaci u PATH promenljivu okruženja.

		PATH=$GOROOT/bin:$GOPATH/bin:$PATH

	***************************************************************

	Gore navedene promenljive okoline preciziraju lokaciju gde se binarni fajlovi
	kopiraju posle kompajliranja i to je put ~/go/bin/. Ovo je konvencionalna
	lokacija za smeštanje kompajliranih datoteka, ali slobodno je promenite ako
	želite drugu lokaciju.

	Sada pokušajte ponovo "go install" i program treba da se kompajlira bez ikakvih
	problema. Možete da unesete kmandu "ls -al ~/go/bin/"" u terminal i možete da
	vidite  da je tu naš prevedeni program "learngo".

	Sada ćemo pokrenuti kompajlirani fajl.

	$ ./learngo // jer je u ~go/bin/ u PATH-u.

	Gornja naredba će pokrenuti binarni fajl "learngo" i odštampaju sledeći izlaz.

		>> Hello World

	Uspešno ste kompajlirali i pokrenuli svoj prvi go program.

	Možda se pitate šta će se dogoditi kada sadrže learngo direktorijum imenik
	sadrži više go datoteka umesto samo main.go. Kako će se u ovom slučaju
	instalirati kompajlirani fajl. Molim vas, sačekajte, razgovaraćemo o tome
	kada naučimo o paketima i Go modulima.

2. go build

	Druga opcija za kompajliranje i pokretanje programa koristi komandu
	"go build". "go build" je sličan "go install", osim što se ne instalira
	(kopira) kompajlirani fajl $GOPATH/bin, radije to smešta binarni na
	lokaciji an kojoj je urađena komanda "go build".

	U terminalu otkucajte sledeću naredbu

	$ cd ~/go/src/learngo/

	da biste promenili trenutni direktorij u learngo direktorijum.

	Nakon toga, unesite sledeću naredbu:

	$ go build

	Gornja naredba stvoriće binarni learngo fajl trenutnom
	direktorijumu. Pokrenite "ls -al" i videćete da je stvorena datoteka
	learngo tu.

	Unesite ./learngo za pokretanje programa. Ovo će takođe odštampati:

	>> Hello World

	Uspešno smo izgradili i pokrenuli svoj prvi Go program koristeći
	"go build" naredbu. :)

3. go run

	Treći način pokretanja programa koristi komandu "go run" za
	kompajliranje i pokretanje.

	Unesite komandnu:

	$ cd ~/go/src/learngo/

	u terminalu da biste promenili trenutni direktorij na "learngo".

	Nakon toga, unesite sledeću naredbu.

	$ go run main.go

	Nakon unosa gore navedene naredbe, možemo videti izlaz

	>> Hello World

	Jedna razlika između komandi "go run" i "go build" je potreban
	naziv datoteke .go kao argument.

	Ispod haube, idite na vožnju mnogo sličnoj kao da se gradi.Umesto
	kompajlirate i instalirate program u trenutni direktorijum,
	"go run" kompajlira datoteku na privremenoj trenutnoj lokaciji i
	pokreće je sa te lokacije.

	Ako ste zainteresovani da znate lokaciju na kojoj je kompajlirana
	datoteka, pokrenite sledeću naredbu:

	$ go run --work main.go

	Pokretanjem  gore navedene naredbe dobićete nešto slično ovime:

	>> work=/tmp/go-build199689936
	>> Hello World

	Vrednost work ključa određuje privremenu lokaciju kojoj će program biti
	kompajliran. Ovo bi moglo varirati u vašem slučaju :)

4. Go Playground

Konačni način pokretanja programa koristi Go plaground.Iako to ima ograničenja,
ova metoda je prikladno kada želimo pokrenuti jednostavne programe, jer koristi
pregledač i ne treba ga instalirati na vaš lokalniračunar :).

Napravio sam igralište za "Hello World" program. Kliknite ovde da biste
pokrenuli program na mreži.

Takođe možete da koristite Go playground da biste podelili svoj izvorni kod sa
drugima.

Sada kada znamo 4 različita načina da pokrenemo Go program, možda ćete biti u
zbrci da biste odlučili koji način za upotrebu.Odgovor je, to zavisi. Obično
koristim igralište kada želim da napravim brzu proveru logike ili saznam kako
funkcioniše standardna funkcija biblioteke.U većini ostalih slučajeva, više
volim da instaliram jer mi daje mogućnost da pokrenem program iz bilo kojeg
direktorija u terminalu, jer kompajlira sve programe na standardnu ~/go/bin/
PATH.

Kratko objašnjenje Hello World Programa
---------------------------------------
Evo Hello World programa koji smo upravo napisali
*/

func HelloWorld() {
	fmt.Println("\n ---Hello World ---")
	fmt.Println("Hello World")
}

/*
Ukratko ćemo raspravljati o tome šta svaki red programa radi. U svakom od sledećih
tutorijala ići ćemo duboko u u svaki deo programa.

1. package main
Svaka go datoteka mora početi sa izjavom imena paketa kome pripada.Paketi se
koriste za čuvanje i ponovnu upotrebu koda. Ovde se koristi ime paketa helloworld.
Hello world je jedina funkcija ovog paketa, dok main funkcija treba da bude u
main fajlu main paketa.

2. import “fmt”
Izjava o uvozu koristi se za uvoz ostalih potrebnih paketa. U našem slučaju se
uvozi "fmt" paket i koristiće se unutar glavne funkcije paketa za štampanje
teksta na standardni izlaz.

3. func HelloWorld()
Ključna reč func označava početak funkcije. main funkcija je posebna funkcija.
Izvršenje programa započinje od main funkcije. { i } zagrade označavaju početak
i kraj tela main funkcije.

3. fmt.Println(“Hello World”) - Funkcija Println-a FMT paketa koristi se za pisanje
teksta na standardni izlaz.

Package.Function ()

je sintaksa da pozove funkcija iz paketa.
*/
