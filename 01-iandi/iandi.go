/*
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

package iandi

import "fmt"

func IAndI() {

	fmt.Println("\n --- Verifying Go installation ---")
	fmt.Println("$go version")
	fmt.Println("$go version go1.19.2 linux/amd64")
}

/*
1.19.2 je bila najnovija verzija Goa kada pišem ovaj tutorijal.Ovo potvrđuje
da je Go pokrenut uspešno.U sledećem tutorialu ćemo napisati naš prvi
"Hello World" program u Go :)
*/
