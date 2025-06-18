/*
Defer
=====

Naredba "defer" se koristi za izvršavanje poziva funkcije neposredno pre nego
što se vrati okolna funkcija u kojoj se nalazi naredba "defer". Definicija
može delovati složeno, ali je prilično jednostavna za razumevanje pomoću
primera.
*/

package de

import (
	"fmt"
	"sync"
	"time"
)

func totalTime(start time.Time) {
	fmt.Printf("Total time taken %f seconds\n", time.Since(start).Seconds())
}

func test() {
	start := time.Now()
	defer totalTime(start)
	time.Sleep(2 * time.Second)
	fmt.Println("Sleep complete")
}

func deferExample() {
	fmt.Println("\n --- Defer example ---")
	test()
}

/*
Gore navedeno je jednostavan program koji ilustruje upotrebu defer. U gornjem
programu, defer se koristi za pronalaženje ukupnog vremena potrebnog za
izvršavanje funkcije test(). Vreme početka izvršavanja test() funkcije se
prosleđuje kao argument u defer totalTime(start). Ovaj poziv defer se izvršava
neposredno pre nego što test() vrati. totalTime ispisuje razliku između start i
trenutnog vremena koristeći time.Since. Da bi se simuliralo neko izračunavanje
koje se dešava u test(), dodate su sleep 2 sekunde.

Pokretanje ovog programa će ispisati:

	>> Sleep complete
	>> Total time taken 2.000000 seconds

Izlaz se odnosi na dodato vreme spavanja od 2 sekunde. Pre nego što test()
funkcija vrati, totalTime se poziva i ispisuje ukupno vreme potrebno za test()
izvršavanje.

Procena argumenata
------------------
Argumenti defer funkcije se izračunavaju kada se defer izraz izvrši, a ne kada
se izvrši stvarni poziv funkcije.

Hajde da ovo razumemo pomoću jednog primera.
*/

func displayValue(a int) {
	fmt.Println("value of a in deferred function", a)
}

func deferEvaluteArgs() {
	fmt.Println("\n --- deferEvaluteArgs ---")
	a := 5
	defer displayValue(a)
	a = 10
	fmt.Println("value of a before deferred function call", a)
}

/*
U gornjem programu a početno ima vrednost 5. Kada se defer izraz izvrši, vrednost
a je 5 i stoga će ovo biti argument funkcije displayValue koja je odložena.
Menjamo vrednost a na 10 u redu br. 13. Sledeći red ispisuje vrednost a. O
vaj program ispisuje,

	>> value of a before deferred function call 10
	>> value of a in deferred function 5

Iz gornjeg izlaza može se razumeti da iako se vrednost a menja na 10 nakon
izvršavanja naredbe defer, stvarni poziv odložene funkcije displayValue(a) i
dalje ispisuje 5.

Odložene metode
---------------
Odlaganje nije ograničeno samo na funkcije . Sasvim je legalno odložiti i poziv
metode . Hajde da napišemo mali program da bismo ovo testirali.
*/

type person struct {
	firstName string
	lastName  string
}

func (p person) fullName() {
	fmt.Printf("%s %s\n", p.firstName, p.lastName)
}

func deferMethod() {

	fmt.Println("\n --- Defer method ---")
	p := person{
		firstName: "John",
		lastName:  "Smith",
	}
	defer p.fullName()
	fmt.Printf("Welcome \n")
}

/*
U gornjem programu smo odložili poziv metode fullName(). Ostatak programa je
sam po sebi razumljiv. Ovaj program ispisuje,

	>> Welcome John Smith

Više odloženih poziva se stavlja na stek
----------------------------------------
Kada funkcija ima više odloženih poziva, oni se stavljaju na stek i izvršavaju
po redosledu "Poslednji ušao, prvi izašao" (LIFO).

Napisaćemo mali program koji ispisuje string unazad koristeći stek defer.
*/

func deferStack() {
	fmt.Println("\n --- Defer stack ---")
	str := "Gopher"
	fmt.Printf("Original String: %s\n", string(str))
	fmt.Printf("Reversed String: ")
	for _, v := range str {
		defer fmt.Printf("%c", v)
	}
}

/*
U gornjem programu, for range petlja iterira string i poziva

	>> defer fmt.Printf("%c", v).

Ovi odloženi pozivi će biti dodati na stek.

	>> defer fmt.Printf("%c", 'G')
	>> defer fmt.Printf("%c", 'o')
	>> defer fmt.Printf("%c", 'p')
	>> defer fmt.Printf("%c", 'h')
	>> defer fmt.Printf("%c", 'e')
	>> defer fmt.Printf("%c", 'r')

Gornja slika odlaganja predstavlja sadržaj steka nakon što su dodani odloženi
pozivi.

Stek je struktura podataka tipa "Last input, first output". Odloženi poziv
koji je poslednji stavljen na stek biće prvi izbačen i prvi izvršen. U ovom
slučaju, "defer fmt.Printf("%c", 'r')" biće izvršen prvi i stoga će string biti
ispisan obrnutim redosledom.

Ovaj program će štampati:

	>> Original String: Gopher
	>> Reversed String: rehpoG

Praktična upotreba Defer
------------------------
U ovom odeljku ćemo razmotriti neke praktičnije primene odlaganja.

Funkcija "defer" se koristi na mestima gde poziv funkcije treba da se izvrši
bez obzira na tok koda. Hajde da ovo razumemo na primeru programa koji koristi
funkciju "waitGroup".

Prvo ćemo napisati program bez korišćenja funkcije "defer", a zatim ćemo ga
modifikovati da bi se koristila deferi razumeti koliko je funkcija "defer"
korisna.
*/

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

func deferWithout() {
	fmt.Println("\n\n --- deferWithoutDefer ---")

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

/*
U gornjem programu, kreirali smo rect strukturu i metodu area koja izračunava
površinu pravougaonika. Ova metoda proverava da li su dužina i širina
pravougaonika manje od nule. Ako jesu, ispisuje odgovarajuću poruku o grešci,
u suprotnom ispisuje površinu pravougaonika.

- Glavna funkcija kreira 3 promenljive tipa react.
- One se zatim dodaju isečku reacts.
- Ovaj isečak se zatim iterira pomoću for range petlje i metoda se poziva kao
  konkurentna gorutina.
- WaitGroup se koristi da bi se osiguralo da glavna funkcija čeka dok se sve
  gorutine ne završe. Ova WaitGroup se prosleđuje metodi area kao argument, a
  metod area poziva wg.Done() da bi obavestila glavnu funkciju da je gorutina
  završila svoj posao.

Ako pažljivo pogledate, možete videti da se ovi pozivi dešavaju neposredno pre
nego što se metoda area vrati. wg.Done() treba pozvati pre nego što se metoda
vrati bez obzira na putanju kojom tok koda ide i stoga se ovi pozivi mogu
efikasno zameniti jednim pozivom.

Hajde da prepišemo gornji program koristeći komandu defer.

U programu ispod, uklonili smo 3 x wg.Done() poziva iz gornjeg programa i
zamenili ga jednim defer wg.Done() pozivom. Ovo čini kod jednostavnijim i
čitljivijim.
*/

type rect1 struct {
	length int
	width  int
}

func (r rect1) area1(wg *sync.WaitGroup) {

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

func deferWith() {
	fmt.Println("\n\n --- deferWithDefer ---")

	var wg sync.WaitGroup
	r1 := rect1{-67, 89}
	r2 := rect1{5, -67}
	r3 := rect1{8, 9}
	rects := []rect1{r1, r2, r3}

	for _, v := range rects {
		wg.Add(1)
		go v.area1(&wg)
	}

	wg.Wait()
	fmt.Println("All go routines finished executing")
}

/*
Ovaj program ispisuje,

	>> rect {8 9}'s area 72
	>> rect {-67 89}'s length should be greater than zero
	>> rect {5 -67}'s width should be greater than zero
	>> All go routines finished executing

Postoji još jedna prednost korišćenja metode defer u gornjem programu. Recimo
da dodamo još jednu putanju povratka metodi area koristeći novi if uslov. Ako
poziv metode wg.Done() nije bio odložen, moramo biti oprezni i osigurati da
pozovemo wg.Done() ovu novu putanju povratka. Ali pošto je poziv metode wg.Done()
odložen, ne moramo da brinemo o dodavanju novih putanja povratka ovoj metodi.
*/

func DeferFunc() {
	fmt.Println("\n --- deferErrorFunc ---")

	deferExample()
	deferEvaluteArgs()
	deferMethod()
	deferStack()
	deferWithout()
	deferWith()
}
