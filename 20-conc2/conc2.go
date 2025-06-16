/*
Baferovani kanali i worker pool-ovi
==================================

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

Hajde da napišemo malo koda i kreiramo baferovani kanal.
*/

package conc2

import (
	"fmt"
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
U gornjem programu, u liniji br. 9 kreiramo baferovani kanal kapaciteta 2. Pošto kanal ima kapacitet 2, moguće je upisati 2 stringa u kanal bez blokiranja. Upisujemo 2 stringa u kanal u linijama br. 10 i 11 i kanal se ne blokira. Čitamo 2 stringa zapisana u linijama br. 12 i 13 respektivno. Ovaj program ispisuje,

naveen
paul

Još jedan primer

Pogledajmo još jedan primer baferovanog kanala u kome se vrednosti ka kanalu zapisuju u istovremenu Gorutinu i čitaju iz glavne Gorutine. Ovaj primer će nam pomoći da bolje razumemo kada se piše u baferovani blok kanala.

 1package main
 2
 3import (
 4    "fmt"
 5    "time"
 6)
 7
 8func write(ch chan int) {
 9    for i := 0; i < 5; i++ {
10        ch <- i
11        fmt.Println("successfully wrote", i, "to ch")
12    }
13    close(ch)
14}
15func main() {
16    ch := make(chan int, 2)
17    go write(ch)
18    time.Sleep(2 * time.Second)
19    for v := range ch {
20        fmt.Println("read value", v,"from ch")
21        time.Sleep(2 * time.Second)
22
23    }
24}

Izvedite program na igralištu

U gornjem programu, baferovani kanal chkapaciteta 2je kreiran u liniji br. 16 Gorutine maini prosleđen writeGorutini u liniji br. 17. Zatim glavna Gorutina prelazi u režim spavanja 2 sekunde. Tokom ovog vremena, writeGorutina se izvršava istovremeno. writeGorutina ima forpetlju koja upisuje brojeve od 0 do 4 u chkanal. Kapacitet ovog baferovanog kanala je 2i stoga će pisanje Goroutinemoći da upisuje vrednosti 0i 1u chkanal odmah, a zatim se blokira dok se barem jedna vrednost ne pročita iz chkanala. Dakle, ovaj program će odmah ispisati sledeća 2 reda.

successfully wrote 0 to ch
successfully wrote 1 to ch

Nakon štampanja gornja dva reda, pisanje u chkanal u writeGorutini je blokirano dok neko ne pročita sa chkanala. Pošto glavna Gorutina miruje 2 sekunde pre nego što počne da čita sa kanala, program neće ništa ispisati naredne 2 sekunde. mainGorutina se budi nakon 2 sekunde i počinje da čita sa chkanala koristeći for rangepetlju u redu br. 19, štampa pročitanu vrednost, a zatim ponovo miruje 2 sekunde i ovaj ciklus se nastavlja dok chse ne zatvori. Dakle, program će ispisati sledeće redove nakon 2 sekunde,

read value 0 from ch
successfully wrote 2 to ch

Ovo će se nastaviti dok se sve vrednosti ne upišu u kanal i on se ne zatvori u writeGorutini. Konačni izlaz bi bio:

successfully wrote 0 to ch
successfully wrote 1 to ch
read value 0 from ch
successfully wrote 2 to ch
read value 1 from ch
successfully wrote 3 to ch
read value 2 from ch
successfully wrote 4 to ch
read value 3 from ch
read value 4 from ch

Zastoj

 1package main
 2
 3import (
 4	"fmt"
 5)
 6
 7func main() {
 8	ch := make(chan string, 2)
 9	ch <- "naveen"
10	ch <- "paul"
11	ch <- "steve"
12	fmt.Println(<-ch)
13	fmt.Println(<-ch)
14}

Izvedite program na igralištu

U gornjem programu, upisujemo 3 niza u baferovani kanal kapaciteta 2. Kada kontrola dođe do trećeg pisanja u liniji br. 11, pisanje je blokirano jer je kanal premašio svoj kapacitet. Sada neka Gorutina mora da čita iz kanala da bi pisanje moglo da se nastavi, ali u ovom slučaju nema istovremenog čitanja rutine iz ovog kanala. Stoga će doći do zastoja i program će paničiti tokom izvršavanja sa sledećom porukom,

fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.main()
	/tmp/sandbox091448810/prog.go:11 +0x8d

Zatvaranje baferovanih kanala

Već smo razgovarali o zatvaranju kanala u prethodnom tutorijalu . Pored onoga što smo naučili u prethodnom tutorijalu, postoji još jedna suptilnost koju treba uzeti u obzir prilikom zatvaranja baferovanih kanala.

Moguće je čitati podatke iz već zatvorenog baferovanog kanala. Kanal će vratiti podatke koji su već upisani u kanal i kada se svi podaci pročitaju, vratiće nultu vrednost kanala.

Hajde da napišemo program da bismo ovo razumeli.

 1package main
 2
 3import (
 4	"fmt"
 5)
 6
 7func main() {
 8	ch := make(chan int, 5)
 9	ch <- 5
10	ch <- 6
11	close(ch)
12	n, open := <-ch
13	fmt.Printf("Received: %d, open: %t\n", n, open)
14	n, open = <-ch
15	fmt.Printf("Received: %d, open: %t\n", n, open)
16	n, open = <-ch
17	fmt.Printf("Received: %d, open: %t\n", n, open)
18}

Izvedite program na igralištu

U gornjem programu, kreirali smo baferovani kanal kapaciteta 5u liniji br. 8. Zatim upisujemo 5i 6u kanal. Kanal se nakon toga zatvara u liniji br. 11. Čak i ako je kanal zatvoren, možemo čitati vrednosti koje su već zapisane u kanal. To se radi u linijama br. 12 i 14. Vrednost nwill be 5i open biće trueu liniji br. 12. Vrednost nwill be 6i open će trueponovo biti u liniji br. 14. Sada smo završili sa čitanjem 5i 6iz kanala i nema više podataka za čitanje. Sada kada se kanal ponovo čita u liniji br. 16, vrednost nwill be 0, što je nulta vrednost inti , openbiće , što falseukazuje da je kanal zatvoren.

Ovaj program će štampati

Received: 5, open: true
Received: 6, open: true
Received: 0, open: false

Isti program se može napisati i korišćenjem petlje for range.

 1package main
 2
 3import (
 4	"fmt"
 5)
 6
 7func main() {
 8	ch := make(chan int, 5)
 9	ch <- 5
10	ch <- 6
11	close(ch)
12	for n := range ch {
13		fmt.Println("Received:", n)
14	}
15}

Izvedite program na igralištu

Petlja for rangeu liniji br. 12 gornjeg programa će pročitati sve vrednosti zapisane u kanal i završiće se kada više nema vrednosti za čitanje, jer je kanal već zatvoren.

Ovaj program će štampati,

Received: 5
Received: 6

Dužina naspram kapaciteta

Kapacitet baferovanog kanala je broj vrednosti koje kanal može da sadrži. To je vrednost koju navodimo prilikom kreiranja baferovanog kanala pomoću makefunkcije .

Dužina baferovanog kanala je broj elemenata koji se trenutno nalaze u njemu u redu čekanja.

Program će razjasniti stvari 😀

 1package main
 2
 3import (
 4	"fmt"
 5)
 6
 7func main() {
 8	ch := make(chan string, 3)
 9	ch <- "naveen"
10	ch <- "paul"
11	fmt.Println("capacity is", cap(ch))
12	fmt.Println("length is", len(ch))
13	fmt.Println("read value", <-ch)
14	fmt.Println("new length is", len(ch))
15}

Izvedite program na igralištu

U gornjem programu, kanal je kreiran sa kapacitetom od 3, odnosno može da sadrži 3 stringa. Zatim upisujemo 2 stringa u kanal u redovima br. 9 i 10, respektivno. Sada kanal ima 2 stringa u redu čekanja i stoga je njegova dužina 2. U redu br. 13, čitamo string iz kanala. Sada kanal ima samo jedan string u redu čekanja i stoga njegova dužina postaje 1. Ovaj program će ispisati,

capacity is 3
length is 2
read value naveen
new length is 1

Grupa čekanja

Sledeći odeljak u ovom tutorijalu je o radničkim bazenima . Da bismo razumeli radničke bazene, prvo moramo znati WaitGroupkako će se koristiti u implementaciji radničke baze.

Grupa čekanja (WaitGroup) se koristi za čekanje da se završi izvršavanje kolekcije gorutina (Goroutine). Kontrola je blokirana dok se sve gorutine ne završe sa izvršavanjem. Recimo da imamo 3 gorutine koje se istovremeno izvršavaju, a koje su nastale iz maingorutine. mainGorutine moraju da sačekaju da se završe ostale 3 gorutine pre nego što se završe. To se može postići korišćenjem grupe čekanja (WaitGroup).

Hajde da prestanemo sa teorijom i odmah napišemo neki kod 😀

 1package main
 2
 3import (
 4	"fmt"
 5	"sync"
 6	"time"
 7)
 8
 9func process(i int, wg *sync.WaitGroup) {
10	fmt.Println("started Goroutine ", i)
11	time.Sleep(2 * time.Second)
12	fmt.Printf("Goroutine %d ended\n", i)
13	wg.Done()
14}
15
16func main() {
17	no := 3
18	var wg sync.WaitGroup
19	for i := 0; i < no; i++ {
20		wg.Add(1)
21		go process(i, &wg)
22	}
23	wg.Wait()
24	fmt.Println("All go routines finished executing")
25}

Trčanje na igralištu

WaitGroup je tipa struktura i kreiramo promenljivu nulte vrednosti tog tipa WaitGroupu liniji br. 18. Način WaitGrouprada je korišćenjem brojača. Kada pozovemo Addi WaitGroupprosledimo mu int, WaitGroupbrojač se povećava za vrednost prosleđenu na Add. Način za smanjenje brojača je pozivanjem Done()metode na WaitGroup. Wait()Metoda blokira Goroutineu kojoj se poziva dok brojač ne postane nula.

U gornjem programu, pozivamo wg.Add(1)u liniji br. 20 unutar forpetlje koja se ponavlja 3 puta. Tako brojač sada postaje 3. forPetlja takođe stvara 3 processGorutine, a zatim wg.Wait()poziv u liniji br. 23 tera mainGorutinu da čeka dok brojač ne postane nula. Brojač se smanjuje pozivom wg.Doneu processGorutini u liniji br. 13. Kada sve 3 generisane Gorutine završe svoje izvršavanje, odnosno kada wg.Done()budu pozvane tri puta, brojač će postati nula, a glavna Gorutina će biti deblokirana.

Važno je proslediti pokazivač wgu liniji br. 21. Ako pokazivač nije prosleden, onda će svaka Gorutina imati svoju kopiju WaitGroupi mainneće biti obaveštena kada završi sa izvršavanjem.

Ovaj program izvodi.

started Goroutine  2
started Goroutine  0
started Goroutine  1
Goroutine 0 ended
Goroutine 2 ended
Goroutine 1 ended
All go routines finished executing

Vaš izlaz može biti drugačiji od mog jer redosled izvršavanja Gorutina može da varira :).
Implementacija baze radnika

Jedna od važnih upotreba baferovanog kanala je implementacija baze radnika .

Generalno, radnički pul je skup niti koje čekaju da im se dodele zadaci. Kada završe dodeljeni zadatak, ponovo se stavljaju na raspolaganje za sledeći zadatak.

Implementiraćemo pul radnika koristeći baferovane kanale. Naš pul radnika će izvršiti zadatak pronalaženja zbira cifara ulaznog broja. Na primer, ako se prosledi 234, izlaz bi bio 9 (2 + 3 + 4). Ulaz u pul radnika biće lista pseudo-slučajnih celih brojeva.

Sledeće su osnovne funkcionalnosti našeg fonda radnika

    Kreiranje baze Gorutina koje slušaju na ulaznom baferovanom kanalu čekajući da se zadaci dodele
    Dodavanje poslova u ulazni baferovani kanal
    Upisivanje rezultata u izlazni baferovani kanal nakon završetka posla
    Čitanje i štampanje rezultata iz izlaznog baferovanog kanala

Ovaj program ćemo pisati korak po korak kako bi bio lakši za razumevanje.

Prvi korak će biti kreiranje struktura koje predstavljaju posao i rezultat.

1type Job struct {
2	id       int
3	randomno int
4}
5type Result struct {
6	job         Job
7	sumofdigits int
8}

Svaka Jobstruktura ima a idi a randomnoza koje se mora izračunati zbir pojedinačnih cifara.

Struktura Resultima jobpolje koje predstavlja zadatak za koji se u polju čuva rezultat (zbir pojedinačnih cifara) sumofdigits.

Sledeći korak je kreiranje baferovanih kanala za prijem poslova i pisanje izlaza.

1var jobs = make(chan Job, 10)
2var results = make(chan Result, 10)

Radničke gorutine osluškuju nove zadatke na jobsbaferovanom kanalu. Kada se zadatak završi, rezultat se upisuje u resultsbaferovani kanal.

Funkcija digitsispod obavlja stvarni posao pronalaženja zbira pojedinačnih cifara celog broja i vraćanja tog rezultata. Dodaćemo period spavanja od 2 sekunde ovoj funkciji samo da bismo simulirali činjenicu da je potrebno neko vreme da ova funkcija izračuna rezultat.

 1func digits(number int) int {
 2	sum := 0
 3	no := number
 4	for no != 0 {
 5		digit := no % 10
 6		sum += digit
 7		no /= 10
 8	}
 9	time.Sleep(2 * time.Second)
10	return sum
11}

Zatim, napisaćemo funkciju koja kreira radničku Gorutinu.

1func worker(wg *sync.WaitGroup) {
2	for job := range jobs {
3		output := Result{job, digits(job.randomno)}
4		results <- output
5	}
6	wg.Done()
7}

Gore navedena funkcija kreira radnik koji čita iz jobskanala, kreira Resultstrukturu koristeći trenutnu jobi povratnu vrednost funkcije digits, a zatim upisuje rezultat u resultsbaferovani kanal. Ova funkcija uzima WaitGroup wgkao parametar na kojem će pozvati Done()metodu kada se sve jobszavrši.

Funkcija createWorkerPoolće kreirati skup radnih Gorutina.

1func createWorkerPool(noOfWorkers int) {
2	var wg sync.WaitGroup
3	for i := 0; i < noOfWorkers; i++ {
4		wg.Add(1)
5		go worker(&wg)
6	}
7	wg.Wait()
8	close(results)
9}

Gore navedena funkcija uzima broj radnika koji treba da se kreiraju kao parametar. wg.Add(1)Pre kreiranja Gorutine poziva funkciju da bi se povećao brojač WaitGroup. Zatim kreira radne Gorutine prosleđivanjem pokazivača WaitGroup wgfunkciji worker. Nakon kreiranja potrebnih radnih Gorutina, čeka da sve Gorutine završe svoje izvršavanje pozivanjem funkcije wg.Wait(). Nakon što sve Gorutine završe sa izvršavanjem, zatvara resultskanal jer su sve Gorutine završile svoje izvršavanje i niko drugi više neće pisati u resultskanal.

Sada kada imamo spreman skup radnika, hajde da napišemo funkciju koja će dodeliti poslove radnicima.

1func allocate(noOfJobs int) {
2	for i := 0; i < noOfJobs; i++ {
3		randomno := rand.Intn(999)
4		job := Job{i, randomno}
5		jobs <- job
6	}
7	close(jobs)
8}

Gore navedena funkcija allocateuzima broj poslova koji treba da se kreiraju kao ulazni parametar, generiše pseudo slučajne brojeve sa maksimalnom vrednošću 998, kreira Jobstrukturu koristeći slučajni broj i brojač iz petlje for ikao identifikator, a zatim ih upisuje u jobskanal. Zatvara jobskanal nakon što upiše sve poslove.

Sledeći korak bi bio kreiranje funkcije koja čita resultskanal i štampa izlaz.

1func result(done chan bool) {
2	for result := range results {
3		fmt.Printf("Job id %d, input random no %d , sum of digits %d\n", result.job.id, result.job.randomno, result.sumofdigits)
4	}
5	done <- true
6}

Funkcija resultčita resultskanal i ispisuje ID posla, uneti slučajni broj i zbir cifara slučajnog broja. Funkcija rezultata takođe uzima donekanal kao parametar u koji upisuje nakon što ispiše sve rezultate.

Sada smo sve podesili. Hajde da završimo poslednji korak pozivanja svih ovih funkcija iz main()funkcije.

 1func main() {
 2	startTime := time.Now()
 3	noOfJobs := 100
 4	go allocate(noOfJobs)
 5	done := make(chan bool)
 6	go result(done)
 7	noOfWorkers := 10
 8	createWorkerPool(noOfWorkers)
 9	<-done
10	endTime := time.Now()
11	diff := endTime.Sub(startTime)
12	fmt.Println("total time taken ", diff.Seconds(), "seconds")
13}

Prvo čuvamo vreme početka izvršavanja programa u redu br. 2 glavne funkcije, a u poslednjem redu (red br. 12) izračunavamo vremensku razliku između vremena kraja (endTime) i vremena početka (startTime) i prikazujemo ukupno vreme potrebno za izvršavanje programa. Ovo je potrebno jer ćemo vršiti neka testiranja promenom broja Gorutina (Goutines).

je noOfJobspodešeno na 100, a zatim allocatese poziva da bi se dodali poslovi u jobskanal.

Zatim donese kreira kanal i prosleđuje Gorutini resulttako da može da počne sa štampanjem izlaza i obavesti kada je sve odštampano.

10Konačno , poziv funkcije kreira skup radnih gorutina createWorkerPool, a zatim funkcija main čeka na donekanalu da se svi rezultati odštampaju.

Evo kompletnog programa za vašu referencu. Uvezao sam i potrebne pakete.

 1package main
 2
 3import (
 4	"fmt"
 5	"math/rand"
 6	"sync"
 7	"time"
 8)
 9
10type Job struct {
11	id       int
12	randomno int
13}
14type Result struct {
15	job         Job
16	sumofdigits int
17}
18
19var jobs = make(chan Job, 10)
20var results = make(chan Result, 10)
21
22func digits(number int) int {
23	sum := 0
24	no := number
25	for no != 0 {
26		digit := no % 10
27		sum += digit
28		no /= 10
29	}
30	time.Sleep(2 * time.Second)
31	return sum
32}
33func worker(wg *sync.WaitGroup) {
34	for job := range jobs {
35		output := Result{job, digits(job.randomno)}
36		results <- output
37	}
38	wg.Done()
39}
40func createWorkerPool(noOfWorkers int) {
41	var wg sync.WaitGroup
42	for i := 0; i < noOfWorkers; i++ {
43		wg.Add(1)
44		go worker(&wg)
45	}
46	wg.Wait()
47	close(results)
48}
49func allocate(noOfJobs int) {
50	for i := 0; i < noOfJobs; i++ {
51		randomno := rand.Intn(999)
52		job := Job{i, randomno}
53		jobs <- job
54	}
55	close(jobs)
56}
57func result(done chan bool) {
58	for result := range results {
59		fmt.Printf("Job id %d, input random no %d , sum of digits %d\n", result.job.id, result.job.randomno, result.sumofdigits)
60	}
61	done <- true
62}
63func main() {
64	startTime := time.Now()
65	noOfJobs := 100
66	go allocate(noOfJobs)
67	done := make(chan bool)
68	go result(done)
69	noOfWorkers := 10
70	createWorkerPool(noOfWorkers)
71	<-done
72	endTime := time.Now()
73	diff := endTime.Sub(startTime)
74	fmt.Println("total time taken ", diff.Seconds(), "seconds")
75}

Trčanje na igralištu

Molimo vas da pokrenete ovaj program na vašem lokalnom računaru radi veće tačnosti u izračunavanju ukupnog potrebnog vremena.

Ovaj program će štampati,

Job id 1, input random no 636, sum of digits 15
Job id 0, input random no 878, sum of digits 23
Job id 9, input random no 150, sum of digits 6
...
total time taken  20.01081009 seconds

Ukupno će biti ispisano 100 redova koji odgovaraju 100 zadataka, a zatim će na kraju u poslednjem redu biti ispisano ukupno vreme potrebno za izvršavanje programa. Vaš izlaz će se razlikovati od mog jer se Gorutine mogu izvršavati bilo kojim redosledom, a ukupno vreme će takođe varirati u zavisnosti od hardvera. U mom slučaju, potrebno je približno 20 sekundi da se program završi.

Sada povećajmo noOfWorkersu mainfunkciji na 20. Udvostručili smo broj radnika. Pošto su se radnički Gorutine povećali (udvostručili, da budem precizan), ukupno vreme potrebno za završetak programa trebalo bi da se smanji (za polovinu, da budem precizan). U mom slučaju, to je postalo 10,004364685 sekundi i program je ispisao,

...
total time taken  10.004364685 seconds

*/

func Conc2Func() {
	fmt.Println("\n --- Conc2 Func ---")

	conc2BuffChannels()
}
