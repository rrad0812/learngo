/*
Baferovani kanali i worker pool-ovi
===================================

Svi kanali o kojima smo govorili u prethodnom tutorijalu su u osnovi bili
nebaferovani. Kao što smo detaljno objasnili u tutorijalu o kanalima, slanje i
prijem na nebaferovani kanal blokira.

Moguće je kreirati kanal sa baferom. Slanja na baferovani kanal se blokira samo
kada je bafer pun. Slično tome, prijem iz baferovanog kanala se blokira samo
kada je bafer prazan.

Baferovani kanali mogu se kreirati dodavanjem dodatnog parametra "capacity"
funkciji "make", koji određuje veličinu bafera.

    >> ch := make(chan type, capacity)

Kapacitet u gornjoj sintaksi treba da bude veći od 0 da bi kanal imao bafer.
Kapacitet za nebaferovani kanal je podrazumevano 0 i stoga smo izostaviljali
parametar kapaciteta prilikom kreiranja nebaferovanih kanala u prethodnom
tutorijalu.

Primer 1
--------
*/

package conc2

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func conc2BuffChannels() {

	fmt.Println("\n --- conc2BuffChannels ---")

	ch := make(chan string, 2)
	ch <- "naveen"
	ch <- "paul"
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

/*
U gornjem programu, kreiramo baferovani kanal tipa string kapaciteta 2. Moguće
je upisati 2 stringa u kanal bez blokiranja. Upisujemo 2 stringa u kanal, i
kanal se ne blokira. Čitamo 2 stringa iz kanala. Ovaj program ispisuje,

	>> naveen
	>> paul

Primer 2
--------
Pogledajmo još jedan primer baferovanog kanala u kome se vrednosti ka kanalu
zapisuju u konkurentnoj gorutini i čitaju iz glavne gorutine. Ovaj primer će
nam pomoći da bolje razumemo kada se piše u baferovani blok kanala.
*/

func write(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("successfully wrote", i, "to ch")
	}
	close(ch)
}

func conc2BuffChannels2() {

	fmt.Println("\n --- conc2BuffChannels2 ---")

	ch := make(chan int, 2)
	go write(ch)
	time.Sleep(2 * time.Second)
	for v := range ch {
		fmt.Println("read value", v, "from ch")
		time.Sleep(2 * time.Second)
	}
}

/*
U gornjem programu, baferovani kanal "ch" kapaciteta 2 je kreiran u glavnoj
gorutini i prosleđen je "write" gorutini. Zatim glavna gorutina prelazi u režim
spavanja 2 sekunde. Tokom ovog vremena, "write" gorutina se izvršava konkurentno.
"write" gorutina ima for petlju koja piše brojeve od 0 do 4 u "ch" kanal.
Kapacitet ovog baferovanog kanala je 2 i stoga će "write" goroutina moći da
upiše vrednosti 0 i 1 u "ch" kanal odmah, a zatim se blokira dok se barem jedna
vrednost ne pročita iz "ch" kanala. Dakle, ovaj program će odmah ispisati sledeća
2 reda:

	>> successfully wrote 0 to ch
	>> successfully wrote 1 to ch

Nakon štampanja gornja dva reda, pisanje u "ch" kanal u "write" gorutini je
blokirano dok neko ne pročita sa "ch" kanala. Pošto glavna gorutina miruje 2
sekunde pre nego što počne da čita sa kanala, program neće ništa ispisati
naredne 2 sekunde. Glavna gorutina se budi nakon 2 sekunde i počinje da čita sa
"ch" kanala koristeći "for range" petlju, štampa pročitanu vrednost, a zatim
ponovo miruje 2 sekunde i ovaj ciklus se nastavlja dok se "ch" se ne zatvori.
Dakle, program će ispisati sledeće redove nakon 2 sekunde,

	>> read value 0 from ch
	>> successfully wrote 2 to ch

Ovo će se nastaviti dok se sve vrednosti ne upišu u kanal i on se ne zatvori u
"write" gorutini. Konačni izlaz bi bio:

	>> successfully wrote 0 to ch
	>> successfully wrote 1 to ch
	>> read value 0 from ch
	>> successfully wrote 2 to ch
	>> read value 1 from ch
	>> successfully wrote 3 to ch
	>> read value 2 from ch
	>> successfully wrote 4 to ch
	>> read value 3 from ch
	>> read value 4 from ch

Zastoj
------
*/
/*
func conc2BuffChannelPanic() {

	fmt.Println("\n --- conc2BuffChannelPanic  ---")
	ch := make(chan string, 2)
	ch <- "naveen"
	ch <- "paul"
	ch <- "steve" // This is panic. All gorutines are asleep - deadlock!

	go fmt.Println(<-ch)
	go fmt.Println(<-ch)
}
*/
/*
U gornjem programu, upisujemo 3 stringa u baferovani kanal kapaciteta 2. Kada
kontrola dođe do trećeg pisanja u liniji br. 11, pisanje je blokirano jer je
kanal dostigao svoj svoj kapacitet. Sada neka gorutina mora da čita iz kanala
da bi pisanje moglo da se nastavi, ali u ovom slučaju nema konkurentnog čitanja
iz ovog kanala. Stoga će doći do zastoja i program će paničiti tokom izvršavanja
sa sledećom porukom:

	>> fatal error: all goroutines are asleep - deadlock!
	>>
	>> goroutine 1 [chan send]:
	>> main.main()
	>> 	/tmp/sandbox091448810/prog.go:11 +0x8d

Zatvaranje baferovanih kanala
-----------------------------

Već smo razgovarali o zatvaranju kanala u prethodnom tutorijalu. Pored onoga
što smo naučili u prethodnom tutorijalu, postoji još jedna suptilnost koju
treba uzeti u obzir prilikom zatvaranja baferovanih kanala.

Moguće je čitati podatke iz već zatvorenog baferovanog kanala. Kanal će vratiti
podatke koji su već upisani u kanal i kada se svi podaci pročitaju, vratiće
nultu vrednost kanala.
*/

func conc2BuffChannelClosed() {

	fmt.Println("\n --- conc2BuffChannelClosed ---")

	ch := make(chan int, 5)
	ch <- 5
	ch <- 6

	close(ch)

	n, open := <-ch
	fmt.Printf("Received: %d, open: %t\n", n, open)

	n, open = <-ch
	fmt.Printf("Received: %d, open: %t\n", n, open)

	n, open = <-ch
	fmt.Printf("Received: %d, open: %t\n", n, open)
}

/*
U gornjem programu, kreirali smo baferovani kanal kapaciteta 5. Zatim upisujemo
5 i 6 u kanal. Kanal se nakon toga zatvara. Čak i ako je kanal zatvoren, možemo
čitati vrednosti koje su već zapisane u kanal. Vrednost n će biti 5 i open biće
true posle prvog čitanja sa kanala. Vrednost n će biti 6 i open će true posle
drugog čitanja sa kanala. Posle trećeg čita nja n će biti 0, što je nulta
vrednost za tip kanala int i open će biti false što ukazuje da je kanal zatvoren.

Ovaj program će štampati:

	>> Received: 5, open: true
	>> Received: 6, open: true
	>> Received: 0, open: false

Isti program se može napisati i korišćenjem petlje for range petlje:
*/

func conc2BuffChannelClosedForRange() {

	fmt.Println("\n --- conc2BuffChannelClosedForrange ---")

	ch := make(chan int, 5)
	ch <- 5
	ch <- 6

	close(ch)

	for n := range ch {
		fmt.Println("Received:", n)
	}
}

/*
Petlja for range gornjeg programa će pročitati sve vrednosti zapisane u kanal i
završiće se kada više nema vrednosti za čitanje, jer je kanal već zatvoren.

Ovaj program će štampati,

Received: 5
Received: 6

Dužina naspram kapaciteta
-------------------------
Kapacitet baferovanog kanala je broj vrednosti koje kanal može da sadrži. To je
vrednost koju navodimo prilikom kreiranja baferovanog kanala pomoću make
funkcije.

Dužina baferovanog kanala je broj elemenata koji se trenutno nalaze u njemu u
redu čekanja.

Program će razjasniti stvari:
*/

func conc2BuffCapVsLen() {

	fmt.Println("\n --- conc2BuffCapVsLen ---")

	ch := make(chan string, 3)
	ch <- "naveen"
	ch <- "paul"

	fmt.Println("capacity is", cap(ch))
	fmt.Println("length is", len(ch))
	fmt.Println("read value", <-ch)
	fmt.Println("new length is", len(ch))
}

/*
U gornjem programu, kanal je kreiran sa kapacitetom od 3, odnosno može da
sadrži 3 stringa. Zatim upisujemo 2 stringa u kanal. Sada kanal ima 2 stringa
u redu čekanja i stoga je njegova dužina 2. Potom čitamo string iz kanala. Sada
kanal ima samo jedan string u redu čekanja i stoga njegova dužina postaje 1.
Ovaj program će ispisati:

	>> capacity is 3
	>> length is 2
	>> read value naveen
	>> new length is 1

WaitGroup
=========

Sledeći odeljak u ovom tutorijalu je o worker pool-ovima. Da bismo razumeli
worker pool-ove, prvo moramo znati WaitGroup i kako ih koristiti u implementaciji
worker poola.

Grupa čekanja (WaitGroup) se koristi za čekanje da se završi izvršavanje
kolekcije gorutina. Kontrola je blokirana dok sve gorutine ne završe sa
izvršavanjem. Recimo da imamo 3 gorutine koje se istovremeno izvršavaju, a koje
su nastale iz glavne gorutine. Glavna gorutina mora da sačeka da se završe
ostale 3 gorutine pre nego što se ona završi. To se može postići korišćenjem
grupe čekanja (WaitGroup).

Hajde da prestanemo sa teorijom i odmah napišemo neki kod:
*/

func process(i int, wg *sync.WaitGroup) {
	fmt.Println("started goroutine ", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("goroutine %d ended\n", i)
	wg.Done()
}

func conc2WaitGroup() {

	fmt.Println("\n --- conc2WaitGroup ---")
	no := 3
	var wg sync.WaitGroup
	for i := 0; i < no; i++ {
		wg.Add(1)
		go process(i, &wg)
	}
	wg.Wait()
	fmt.Println("All go routines finished executing")
}

/*
WaitGroup je tipa strukture. Kreiramo promenljivu nulte vrednosti tipa
WaitGroupe. Način WaitGroupra je korišćenjem brojača. Kada pozovemo Add i
WaitGroupI prosledimo int, WaitGroup brojač se povećava za vrednost
prosleđenu na Add. Način za smanjenje brojača je pozivanjem Done() metode na
WaitGroup. Wait() metoda blokira goroutinu U kojoj se poziva, dok brojač
ne postane nula.

U gornjem programu, pozivamo wg.Add(1) unutar for petlje koja se ponavlja 3
puta. Tako brojač sada postaje 3. for petlja takođe stvara 3 process gorutine,
a zatim wg.Wait() poziv tera glavnu gorutinu da čeka dok brojač ne postane nula.
Brojač se smanjuje pozivom wg.Done u process gorutini. Kada sve 3 generisane
gorutine završe svoje izvršavanje, odnosno kada wg.Done() bude pozvana tri puta,
brojač će postati nula, a glavna gorutina će biti deblokirana.

Važno je proslediti pokazivač wg u gorutinu koja se pkreće. Ako pokazivač nije
prosleden, onda će svaka gorutina imati svoju kopiju WaitGroup i glavna gorutina
neće biti obaveštena  kada završi sa izvršavanjem.

Ovaj program ispisuje:

	>> started goroutine  2
	>> started goroutine  0
	>> started goroutine  1
	>> goroutine 0 ended
	>> goroutine 2 ended
	>> goroutine 1 ended
	>> All go routines finished executing

Vaš izlaz može biti drugačiji od mog jer redosled izvršavanja Gorutina može da
varira :).

Implementacija worker poola
---------------------------
Jedna od važnih upotreba baferovanog kanala je implementacija worker poola.
Generalno, worker pool je skup niti koje čekaju da im se dodele zadaci. Kada
završe dodeljeni zadatak, ponovo se stavljaju na raspolaganje za sledeći
zadatak.

Implementiraćemo worker pool koristeći baferovane kanale. Naš worker pool će
izvršiti zadatak pronalaženja zbira cifara ulaznog broja. Na primer, ako se
prosledi 234, izlaz bi bio 9 (2 + 3 + 4). Ulaz u worker pool biće lista
pseudo-slučajnih celih brojeva.

Sledeće su osnovne funkcionalnosti našeg worker pool-a:

- Kreiranje poola gorutina koje slušaju na ulaznom baferovanom kanalu čekajući
  da se zadaci dodele
- Dodavanje poslova u ulazni baferovani kanal
- Upisivanje rezultata u izlazni baferovani kanal nakon završetka posla
- Čitanje i štampanje rezultata iz izlaznog baferovanog kanala

Ovaj program ćemo pisati korak po korak kako bi bio lakši za razumevanje.

Prvi korak će biti kreiranje struktura koje predstavljaju posao i rezultat.

type Job struct {
	id       int
	randomno int
}
type Result struct {
	job         Job
	sumofdigits int
}

Svaka "Job" struktura ima a "randomno" za koje se mora izračunati zbir
pojedinačnih cifara. Struktura "Result" ima "job" polje koje predstavlja zadatak
za koji se u polju "sumofdigits" čuva rezultat (zbir pojedinačnih cifara).

Sledeći korak je kreiranje baferovanih kanala za prijem poslova i pisanje izlaza.

var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

Worker gorutine osluškuju nove zadatke na jobs baferovanom kanalu. Kada se
zadatak završi, rezultat se upisuje u results baferovani kanal.

Funkcija "digits" ispod obavlja stvarni posao pronalaženja zbira pojedinačnih
cifara celog broja i vraćanja tog rezultata. Dodaćemo period spavanja od 2
sekunde ovoj funkciji samo da bismo simulirali činjenicu da je potrebno neko
vreme da ova funkcija izračuna rezultat.

func digits(number int) int {
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}
	time.Sleep(2 * time.Second)
	return sum
}

Zatim, napisaćemo funkciju koja kreira worker gorutinu.

func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		output := Result{job, digits(job.randomno)}
		results <- output
	}
	wg.Done()
}

Gore navedena funkcija kreira "worker" koji čita iz "jobs" kanala, kreira
"Result" strukturu koristeći trenutnu job i povratnu vrednost funkcije
"digits", a zatim upisuje rezultat u "results" baferovani kanal.

Ova funkcija uzima WaitGroup "wg" kao parametar na kojem će pozvati "Done()"
metodu kada sve "jobs" završi.

Funkcija "createWorkerPool" će kreirati skup radnih gorutina.

func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}

Gore navedena funkcija uzima broj workera koji treba da se kreiraju kao
parametar. Pre kreiranja gorutine poziva funkciju wg.Add(1) da bi se povećao
brojač WaitGroup. Zatim kreira worker gorutinu prosleđivanjem pokazivača
WaitGroup "wg" funkciji worker. Nakon kreiranja potrebnih worker gorutina,
čeka da sve gorutine završe svoje izvršavanje pozivanjem funkcije wg.Wait().
Nakon što sve gorutine završe sa izvršavanjem, zatvara results kanal jer su
sve gorutine završile svoje izvršavanje i niko drugi više neće pisati u results
kanal.

Sada kada imamo spreman skup workera, hajde da napišemo funkciju koja će
dodeliti poslove radnicima.

func allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		randomno := rand.Intn(999)
		job := Job{i, randomno}
		jobs <- job
	}
	close(jobs)
}

Gore navedena funkcija allocate uzima broj poslova koji treba da se kreiraju
kao ulazni parametar, generiše pseudo slučajne brojeve sa maksimalnom vrednošću
998, kreira Jobstrukturu koristeći slučajni broj i brojač iz petlje for i kao
identifikator, a zatim ih upisuje u jobskanal. Zatvara jobskanal nakon što upiše
sve poslove.

Sledeći korak bi bio kreiranje funkcije koja čita results kanal i štampa izlaz.

func result(done chan bool) {
	for result := range results {
		fmt.Printf("Job id %d, input random no %d , sum of digits %d\n", result.job.id, result.job.randomno, result.sumofdigits)
	}
	done <- true
}

Funkcija resultčita resultskanal i ispisuje ID posla, uneti slučajni broj i zbir cifara slučajnog broja. Funkcija rezultata takođe uzima donekanal kao parametar u koji upisuje nakon što ispiše sve rezultate.

Sada smo sve podesili. Hajde da završimo poslednji korak pozivanja svih ovih funkcija iz main()funkcije.

func main() {
	startTime := time.Now()
	noOfJobs := 100
	go allocate(noOfJobs)
	done := make(chan bool)
	go result(done)
	noOfWorkers := 10
	createWorkerPool(noOfWorkers)
	<-done
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}

Prvo sačuvamo vreme početka izvršavanja programa, a na kraju izračunavamo
vremensku razliku između vremena kraja (endTime) i vremena početka
(startTime) i prikazujemo ukupno vreme potrebno za izvršavanje programa.
Ovo je potrebno jer ćemo vršiti neka testiranja promenom broja gorutina.

noOfJobs je podešeno na 100, a zatim se poziva allocate da bi dodali poslove
u jobs kanal.

Zatim se kreira done kanal i prosleđuje result gorutini tako da može da počne
sa štampanjem izlaza i obavesti kada je sve odštampano.

Konačno postavi se noWorkers na 10 i pozivq funkcijq createWorkerPool da kreira
skup worker gorutina, a zatim glavna funkcija čeka na done kanalu da se svi
rezultati odštampaju.

Evo kompletnog programa za vašu referencu. Uvezao sam i potrebne pakete.
*/

type Job struct {
	id       int
	randomno int
}
type Result struct {
	job         Job
	sumofdigits int
}

var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

func digits(number int) int {
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}
	time.Sleep(2 * time.Second)
	return sum
}
func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		output := Result{job, digits(job.randomno)}
		results <- output
	}
	wg.Done()
}
func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}
func allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		randomno := rand.Intn(999)
		job := Job{i, randomno}
		jobs <- job
	}
	close(jobs)
}
func result(done chan bool) {
	for result := range results {
		fmt.Printf("Job id %d, input random no %d , sum of digits %d\n", result.job.id, result.job.randomno, result.sumofdigits)
	}
	done <- true
}
func conc2WorkerPool() {

	fmt.Println("\n --- conc2WorkerPool ---")

	startTime := time.Now()
	noOfJobs := 100
	go allocate(noOfJobs)
	done := make(chan bool)
	go result(done)
	noOfWorkers := 10
	createWorkerPool(noOfWorkers)
	<-done
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}

/*
Molimo vas da pokrenete ovaj program na vašem lokalnom računaru radi veće
tačnosti u izračunavanju ukupnog potrebnog vremena.

Ovaj program će štampati,

Job id 1, input random no 636, sum of digits 15
Job id 0, input random no 878, sum of digits 23
Job id 9, input random no 150, sum of digits 6
...
total time taken  20.01081009 seconds

Ukupno će biti ispisano 100 redova koji odgovaraju za 100 zadataka, a zatim će
na kraju u poslednjem redu biti ispisano ukupno vreme potrebno za izvršavanje
programa. Vaš izlaz će se razlikovati od mog jer se gorutine mogu izvršavati
bilo kojim redosledom, a ukupno vreme će takođe varirati u zavisnosti od
hardvera. U mom slučaju, potrebno je približno 20 sekundi da se program završi.

Sada povećajmo noOfWorkersu main funkciji na 20. Udvostručili smo broj radnika.
Pošto su se radnički Gorutine povećali (udvostručili, da budem precizan), ukupno
vreme potrebno za završetak programa trebalo bi da se smanji (za polovinu,
da budem precizan). U mom slučaju, to je postalo 10,004364685 sekundi i program
je ispisao:

...
total time taken  10.004364685 seconds
*/

func Conc2Func() {
	fmt.Println("\n --- Conc2 Func ---")

	conc2BuffChannels()
	conc2BuffChannels2()
	// conc2BuffChannelPanic()
	conc2BuffChannelClosed()
	conc2BuffChannelClosedForRange()
	conc2BuffCapVsLen()
	conc2WaitGroup()
	conc2WorkerPool()
}
