/*
Defer
=====

Naredba "defer" se koristi za izvršavanje poziva funkcije neposredno pre nego
što se vrati okolna funkcija u kojoj se nalazi naredba "defer". Definicija
može delovati složeno, ali je prilično jednostavna za razumevanje pomoću
primera.

package main

import (
	"fmt"
	"time"
)

func totalTime(start time.Time) {
	fmt.Printf("Total time taken %f seconds", time.Since(start).Seconds())
}

func test() {
	start := time.Now()
	defer totalTime(start)
	time.Sleep(2 * time.Second)
	fmt.Println("Sleep complete")
}

func main() {
	test()
}

Gore navedeno je jednostavan program koji ilustruje upotrebu defer. U gornjem
programu, deferse koristi za pronalaženje ukupnog vremena potrebnog za
izvršavanje funkcije test(). Vreme početka izvršavanja test() funkcije se
prosleđuje kao argument u defer totalTime(start) redu br. 14. Ovaj poziv
odlaganja se izvršava neposredno pre nego što test() vrati. totalTime ispisuje
razliku između start i trenutnog vremena koristeći time.Since. Da bi se
simuliralo neko izračunavanje koje se dešava u , dodate su test()2 sekunde.sleep

Pokretanje ovog programa će ispisati:

	>> Sleep complete
	>> Total time taken 2.000000 seconds

Izlaz se odnosi na dodato vreme spavanja od 2 sekunde. Pre nego št est() funkcija
vrati, totalTime poziva se i ispisuje ukupno vreme potrebno za test() izvršavanje.

Procena argumenata
------------------
Argumenti odložene funkcije se izračunavaju kada deferse izraz izvrši, a ne kada
se izvrši stvarni poziv funkcije.

Hajde da ovo razumemo pomoću jednog primera.

package main

import (
	"fmt"
)

func displayValue(a int) {
	fmt.Println("value of a in deferred function", a)
}
func main() {
	a := 5
	defer displayValue(a)
	a = 10
	fmt.Println("value of a before deferred function call", a)
}


U gornjem programu a početno ima vrednost 5. Kada defer se izraz izvrši, vrednost
a je 5 i stoga će ovo biti argument funkcije displayValuekoja je odložena. Menjamo
vrednost a na 10 u redu br. 13. Sledeći red ispisuje vrednost a. Ovaj program
izbacuje,

	>> value of a before deferred function call 10
	>> value of a in deferred function 5

Iz gornjeg izlaza može se razumeti da iako se vrednost amenja na 10nakon
izvršavanja naredbe defer, stvarni poziv odložene funkcije displayValue(a)i dalje
ispisuje 5.

Odložene metode

Odlaganje nije ograničeno samo na funkcije . Sasvim je legalno odložiti i poziv
metode . Hajde da napišemo mali program da bismo ovo testirali.

package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName string
}

func (p person) fullName() {
	fmt.Printf("%s %s",p.firstName,p.lastName)
}

func main() {
	p := person {
		firstName: "John",
		lastName: "Smith",
	}
	defer p.fullName()
	fmt.Printf("Welcome ")
}

U gornjem programu smo odložili poziv metode. Ostatak programa je sam po sebi
razumljiv. Ovaj program izbacuje,

	>> Welcome John Smith

Više odloženih poziva se stavlja u stek.

Kada funkcija ima više odloženih poziva, oni se stavljaju na stek i izvršavaju
po redosledu „Poslednji ušao, prvi izašao“ (LIFO).

Napisaćemo mali program koji ispisuje string unazad koristeći stek odlaganja.

package main

import (
	"fmt"
)

func main() {
	str := "Gopher"
	fmt.Printf("Original String: %s\n", string(str))
	fmt.Printf("Reversed String: ")
	for _, v := range str {
		defer fmt.Printf("%c", v)
	}
}

U gornjem programu, for range petlja iterira string i poziva

defer fmt.Printf("%c", v).

Ovi odloženi pozivi će biti dodati na stek.

Gornja slika predstavlja sadržaj steka nakon što su dodani odloženi pozivi. Stek je struktura podataka tipa „poslednji ulazi, prvi izlazi“. Odloženi poziv koji je poslednji stavljen na stek biće izbačen i prvi izvršen. U ovom slučaju, defer fmt.Printf("%c", 'n')biće izvršen prvi i stoga će string biti ispisan obrnutim redosledom.

Ovaj program će štampati

Original String: Gopher
Reversed String: rehpoG

Praktična upotreba odlaganja

U ovom odeljku ćemo razmotriti neke praktičnije primene odlaganja.

Funkcija „defer“ se koristi na mestima gde poziv funkcije treba da se izvrši bez obzira na tok koda. Hajde da ovo razumemo na primeru programa koji koristi funkciju „waitGroup“ . Prvo ćemo napisati program bez korišćenja funkcije „defer“ defer, a zatim ćemo ga modifikovati da bi se koristila deferi razumeti koliko je funkcija „defer“ korisna.

package main

import (
	"fmt"
	"sync"
)

type rect struct {
	length int
	width  int
}

func (r rect) area(wg *sync.WaitGroup) {
	if r.length < 0 {
		fmt.Printf("rect %v's length should be greater than zero\n", r)
		wg.Done()
		return
	}
	if r.width < 0 {
		fmt.Printf("rect %v's width should be greater than zero\n", r)
		wg.Done()
		return
	}
	area := r.length * r.width
	fmt.Printf("rect %v's area %d\n", r, area)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	r1 := rect{-67, 89}
	r2 := rect{5, -67}
	r3 := rect{8, 9}
	rects := []rect{r1, r2, r3}
	for _, v := range rects {
		wg.Add(1)
		go v.area(&wg)
	}
	wg.Wait()
	fmt.Println("All go routines finished executing")
}

U gornjem programu, kreirali smo rectstrukturu u liniji br. 8 i metodu area u
rect liniji br. 13 koja izračunava površinu pravougaonika. Ova metoda proverava
da li su dužina i širina pravougaonika manje od nule. Ako jesu, ispisuje
odgovarajuću poruku o grešci, u suprotnom ispisuje površinu pravougaonika.

Funkcija mainkreira 3 promenljive tipa . One se zatim dodaju segmentu. Ovaj
segment se zatim iterira pomoću petlje i metoda se poziva kao konkurentna
Gorutina. WaitGroup se koristi da bi se osiguralo da glavna funkcija čeka dok
se r1sve Gorutine ne završe. Ova WaitGroup se prosleđuje metodi area kao
argument, a metod area poziva da bi obavestila glavnu funkciju da je
Gorutina završila svoj posao. Ako pažljivo pogledate, možete videti da se
ovi pozivi dešavaju neposredno pre nego što se metoda area vrati. wg.Done()
treba pozvati pre nego što se metoda vrati bez obzira na putanju kojom tok
koda ide i stoga se ovi pozivi mogu efikasno zameniti jednim pozivom.

r2r3rectrectsfor rangearea wgwg.Done()defer

Hajde da prepišemo gornji program koristeći komandu defer.

U programu ispod, uklonili smo 3 wg.Done() poziva iz gornjeg programa i zamenili
ga jednim defer wg.Done()pozivom u liniji br. 14. Ovo čini kod jednostavnijim i
čitljivijim.

package main

import (
	"fmt"
	"sync"
)

type rect struct {
	length int
	width  int
}

func (r rect) area(wg *sync.WaitGroup) {
	defer wg.Done()
	if r.length < 0 {
		fmt.Printf("rect %v's length should be greater than zero\n", r)
		return
	}
	if r.width < 0 {
		fmt.Printf("rect %v's width should be greater than zero\n", r)
		return
	}
	area := r.length * r.width
	fmt.Printf("rect %v's area %d\n", r, area)
}

func main() {
	var wg sync.WaitGroup
	r1 := rect{-67, 89}
	r2 := rect{5, -67}
	r3 := rect{8, 9}
	rects := []rect{r1, r2, r3}
	for _, v := range rects {
		wg.Add(1)
		go v.area(&wg)
	}
	wg.Wait()
	fmt.Println("All go routines finished executing")
}

Ovaj program izlazi,

	>> rect {8 9}'s area 72
	>> rect {-67 89}'s length should be greater than zero
	>> rect {5 -67}'s width should be greater than zero
	>> All go routines finished executing

Postoji još jedna prednost korišćenja metode defer u gornjem programu. Recimo
da dodamo još jednu putanju povratka metodi areakoristeći novi if uslov. Ako
poziv metode wg.Done()nije bio odložen, moramo biti oprezni i osigurati da
pozovemo wg.Done()ovu novu putanju povratka. Ali pošto je poziv metode wg.Done()
odložen, ne moramo da brinemo o dodavanju novih putanja povratka ovoj metodi.
*/
