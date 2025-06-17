/*
Mutex
=====

Kritična sekcija
----------------
Pre nego što pređemo na mutex, važno je razumeti koncept kritične sekcije u
konkurentnom programiranju. Kada se program izvršava konkurentno, delovi koda
koji menjaju deljene resurse ne bi trebalo da budu pristupačni višestrukim
goroutinama konkurentno. Ovaj deo koda koji menja deljene resurse naziva se
"kritična sekcija". Na primer, pretpostavimo da imamo neki deo koda koji
povećava promenljivu x za 1.

	>> x = x + 1

Sve dok se gore navedenom delu koda pristupa iz jedne gorutine, ne bi trebalo
da bude problema.

Da vidimo zašto će ovaj kod otkazati kada se više gorutina izvršava konkurentno.
Radi jednostavnosti, pretpostavimo da imamo 2 gorutine koje konkurentno
izvršavaju gornju liniju koda.

Interno, gornja linija koda će biti izvršena od strane sistema u sledećim
koracima (postoje dodatni tehnički detalji koji se odnose na registre, kako
funkcioniše sabiranje i tako dalje, ali za potrebe ovog tutorijala
pretpostavimo da su ovo ta tri koraka),

- dobiti trenutnu vrednost x
- izračunati x + 1
- dodeliti izračunatu vrednost u koraku 2 u x

Kada ova tri koraka sprovede samo jedna gorutina, sve je u redu.

Hajde da razgovaramo o tome šta se dešava kada 2 gorutine istovremeno pokreću
ovaj kod. Slika ispod prikazuje jedan scenario šta bi se moglo dogoditi kada
dve gorutine konkurentno pristupe liniji koda x = x + 1.

	Pretpostavimo da je početna vrednost x 0. Gorutina 1 dobija početnu
	vrednost x, izračunava x + 1 i pre nego što je mogla da dodeli izračunatu
	vrednost x, sistemski kontekst se prebacuje na gorutinu 2. Sada gorutine 2
	dobija početnu vrednost x, koja je i dalje 0, izračunava x + 1.

	Nakon ovoga, sistemski kontekst se ponovo prebacuje na gorutinu 1. Sada
	gorutina 1 dodeljuje svoju izračunatu vrednost 1 u x i stoga x postaje 1.
	Zatim gorutina 2 ponovo počinje izvršavanje, a zatim dodeljuje svoju
	izračunatu vrednost, koja je ponovo 1 u x stoga x je 1 nakon što se obe
	gorutine izvrše.

Sada da vidimo drugačiji scenario šta bi se moglo dogoditi.

	Goroutina 1 počinje izvršavanje i završava sva svoja tri koraka i stoga
	vrednost x postaje 1. Zatim Goroutine 2 počinje izvršavanje. Sada je
	vrednost x jednaka 1, a kada se gorutina 2 završi, vrednost x je 2.

Dakle, iz ova dva slučaja možete videti da je konačna vrednost x jednaka 1 ili
2 u zavisnosti od toga kako se dešava promena konteksta. Ova vrsta nepoželjne
situacije gde izlaz programa zavisi od redosleda izvršavanja gorutina naziva se
"uslov trke".

U gore navedenom scenariju, uslov trke je mogao biti izbegnut ako bi samo
jednoj gorutini bilo dozvoljeno da pristupi kritičnom delu koda u bilo kom
trenutku. Ovo je omogućeno korišćenjem mutex-a.

Mutex
-----
Mutex se koristi za obezbeđivanje mehanizma zaključavanja kako bi se osiguralo
da samo jedna gorutina pokreće kritični deo koda u bilo kom trenutku, kako bi
se sprečilo nastajanje uslova trke.

Mutex je dostupan u paketu sync. Na mutex-u su definisane dve metode, naime
Lock i Unlock. Bilo koji kod koji se nalazi između poziva na Lock i Unlock biće
izvršen samo od strane jedne gorutine, čime se izbegava uslov trke.

	>> mutex.Lock()
	>> x = x + 1
	>> mutex.Unlock()

U gornjem kodu, x = x + 1 biće izvršen od strane samo jedne Gorutine u bilo kom
trenutku, čime se sprečava uslov trke.

Ako jedna gorutina već drži zaključavanje i ako nova gorutina pokušava da ga
dobije, nova gorutina će biti blokirana dok se muteks ne otključa.

Program sa uslovom trke
-----------------------
U ovom odeljku ćemo napisati program koji ima uslov trke, a u narednim
odeljcima ćemo ispraviti uslov trke.
*/

package mut

import (
	"fmt"
	"sync"
	"time"
)

var x = 0

func increment(wg *sync.WaitGroup) {
	x = x + 1
	wg.Done()
}

func mutRaceCond() {

	fmt.Println("\n --- mutRaceCond ---")

	var w sync.WaitGroup
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w)
	}
	w.Wait()
	fmt.Println("final value of x", x)
}

/*
U gornjem programu, increment funkcija povećava vrednost promenljive x za 1 a
zatim poziva Done() WaitGroup da obavesti o njenom završetku.

Generišemo 1000 increment gorutina. Svaka od ovih gorutina se izvršava
istovremeno, a uslov trke se javlja pri pokušaju povećanja x jer više gorutina
pokušava da pristupi vrednosti x konkurentno.

Molimo vas da pokrenete ovaj program na vašoj lokalnoj mašini jer je igralište
determinističko i uslov trke se neće pojaviti na igralištu. Pokrenite ovaj
program više puta na vašoj lokalnoj mašini i videćete da će izlaz biti
drugačiji svaki put zbog uslova trke. Neki od izlaza na koje sam naišao su
final value of x 941, final value of x 928, final value of x 922 i tako dalje.

Rešavanje uslova trke pomoću mutex-a
------------------------------------
U gornjem programu, generišemo 1000 Gorutina. Ako svaka poveća vrednost x za 1,
konačna željena vrednost x treba da bude 1000. U ovom odeljku, popravićemo
uslov trke u gornjem programu koristeći mutex.
*/

var x1 = 0

func increment1(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	x1 = x1 + 1
	m.Unlock()
	wg.Done()
}

func mutRaceCondWithMutex() {

	fmt.Println("\n --- mutRaceCondWithMutex ---")

	timeStart := time.Now()

	var w sync.WaitGroup
	var m sync.Mutex
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment1(&w, &m)
	}
	w.Wait()

	timeEnd := time.Now()

	diff := timeEnd.Sub(timeStart)
	fmt.Println("final value of x", x1, "seconds:", diff.Seconds())
}

/*
Mutex je tipa struct i kreiramo nil promenljivu m tipa Mutex. U gornjem programu
smo promenili increment funkciju tako da kod koji povećava x, x = x + 1 bude
između m.Lock() i m.Unlock(). Sada je ovaj kod bez ikakvih uslova trke jer je
samo jednoj gorutini dozvoljeno da izvrši ovaj deo koda u bilo kom trenutku.

Sada, ako se ovaj program pokrene, on će ispisati

	>> final value of x 1000

Važno je proslediti adresu muteksa. Ako se muteks prosledi po vrednosti umesto
po adresi, svaka gorutina će imati svoju kopiju muteksa i uslov trke će se i
dalje javljati.

Rešavanje uslova trke korišćenjem kanala
----------------------------------------
Možemo rešiti uslov trke korišćenjem kanala. Da vidimo kako se to radi.
*/

var x2 = 0

func increment2(wg *sync.WaitGroup, ch chan bool) {
	ch <- true
	x2 = x2 + 1
	<-ch
	wg.Done()
}

func mutRaceCondWithChannel() {

	fmt.Println("\n --- mutRaceCondWithChannel ---")

	timeStart := time.Now()

	var w sync.WaitGroup
	ch := make(chan bool, 1)
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment2(&w, ch)
	}
	w.Wait()

	timeEnd := time.Now()

	diff := timeEnd.Sub(timeStart)
	fmt.Println("final value of x", x2, "seconds:", diff.Seconds())
}

/*
U gornjem programu, kreirali smo baferovani kanal kapaciteta 1 i on se
prosleđuje gorutini increment2. Ovaj baferovani kanal se koristi da bi se
osiguralo da samo jedna gorutina pristupa kritičnoj sekciji koda koja povećava
x. To se postiže prosleđenjem true na baferovani kanal neposredno pre nego što
x se poveća. Pošto baferovani kanal ima kapacitet od 1, sve ostale gorutine koje
pokušavaju da pišu u ovaj kanal su blokirane dok se vrednost ne pročita iz ovog
kanala nakon povećanja x. Ovo efikasno dozvoljava samo jednoj gorutini da
pristupi kritičnoj sekciji.

Ovaj program takođe štampa

	>> final value of x 1000

Mutex vs Channels
-----------------
Rešili smo problem uslova trke koristeći i mutekse i kanale. Pa kako da
odlučimo šta kada da koristimo? Odgovor leži u problemu koji pokušavate da
rešite.

Ako je problem koji pokušavate da rešite bolje prilagođen muteksima, onda samo
napred i koristite muteks. Ne oklevajte da koristite muteks ako je
potrebno. Ako se čini da problem bolje odgovara kanalima, onda ga koristite :).

Većina početnika u Gou pokušava da reši svaki problem konkurentnosti koristeći
kanal, jer je to sjajna karakteristika jezika. To je pogrešno. Jezik nam daje
mogućnost da koristimo Mutex ili Channel i nema greške u izboru bilo kog od
njih.

Generalno, koristite kanale kada gorutine treba da komuniciraju jedna sa drugom,
a mutekse kada samo jedna gorutina treba da pristupi kritičnom delu koda.

U slučaju problema koji smo rešili gore, više bih voleo da koristim mutex jer
ovaj problem ne zahteva nikakvu komunikaciju između gorutina. Stoga bi mutex
bio prirodno rešenje.

Moj savet bi bio da izaberete alat za problem i da ne pokušavate da prilagodite
problem alatu :)
*/

func MutFunc() {

	fmt.Println("\n --- MutFunc ---")

	mutRaceCond()
	mutRaceCondWithMutex()
	mutRaceCondWithChannel()
}
