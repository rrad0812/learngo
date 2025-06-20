package conc

import (
	"fmt"
	"time"
)

/*
Uvod u konkurentnost
====================

Go je konkurentni, a ne paralelni jezik. Pre nego što počnemo da razgovaramo o
tome kako se u Go-u vodi računa o konkurentnosti, moramo prvo razumeti šta je
konkurentnost i kako se razlikuje od paralelizma.

Šta je konkurentnost?
---------------------
Konkurentnost je sposobnost bavljenja mnogim stvarima konkurentno. Najbolje je
objasniti primerom:

Zamislimo osobu koja trči. Tokom jutarnjeg trčanja, recimo da joj se odvežu
pertle. Sada osoba prestaje da trči, veže pertle, a zatim ponovo počinje da
trči. Ovo je klasičan primer konkurentnog trčanja. Osoba je sposobna da se nosi
i sa trčanjem i sa vezivanjem pertli, odnosno da je u stanju da se nosi sa
mnogo stvari istovremeno :)

Šta je paralelizam i kako se razlikuje od konkurentnosti?
---------------------------------------------------------
Paralelizam je obavljanje mnogo stvari istovremeno. Možda zvuči slično
konkurentnosti, ali je zapravo drugačije.

Hajde da to bolje razumemo na istom primeru džogiranja. U ovom slučaju,
pretpostavimo da osoba džogira i istovremeno sluša muziku na svom iPod-u.
U ovom slučaju, osoba džogira i sluša muziku u isto vreme, odnosno radi mnogo
stvari istovremeno. Ovo se naziva paralelizam.

Konkurencija i paralelizam - tehnička tačka gledišta
----------------------------------------------------
Razumeli smo šta je konkurentnost i kako se razlikuje od paralelizma koristeći
primere iz stvarnog sveta. Sada hajde da ih pogledamo sa tehničke tačke
gledišta, pošto smo gikovi :).

Recimo da programiramo veb pregledač. Veb pregledač ima različite komponente.
Dve od njih su oblast za prikazivanje veb stranica i program za preuzimanje
datoteka sa interneta. Pretpostavimo da smo strukturirali kod našeg pregledača
na takav način da se svaka od ovih komponenti može izvršavati nezavisno (Ovo se
radi korišćenjem niti u jezicima kao što je Java, a u Go-u to možemo postići
korišćenjem gorutina, više o tome kasnije). Kada se ovaj pregledač pokreće na
jednojezgranom procesoru, procesor će prebacivati kontekst između dve komponente
pregledača. Možda će neko vreme preuzimati datoteku, a zatim će preći na
prikazivanje html koda veb stranice koju je korisnik zahtevao. Ovo je poznato
kao konkurentnost. Konkurentni procesi počinju u različitim vremenskim tačkama
i njihovi ciklusi izvršavanja se preklapaju. U ovom slučaju, preuzimanje i
prikazivanje počinju u različitim vremenskim tačkama i njihova izvršavanja se
preklapaju.

Recimo da isti pregledač radi na višejezgarnom procesoru. U ovom slučaju,
komponenta za preuzimanje datoteka i komponenta za renderovanje HTML-a mogu da
rade istovremeno na različitim jezgrima. Ovo je poznato kao paralelizam.

Go - konkurentnost-paralelizam
------------------------------
Paralelizam neće uvek rezultirati bržim vremenom izvršavanja. To je zato što
komponente koje rade paralelno moraju da komuniciraju jedna sa drugom.
Na primer, u slučaju našeg pregledača, kada je preuzimanje datoteke završeno,
to bi trebalo da bude saopšteno korisniku, recimo pomoću iskačućeg prozora.
Ova komunikacija se odvija između komponente odgovorne za preuzimanje i
komponente odgovorne za prikazivanje korisničkog interfejsa. Ovo komunikaciono
opterećenje je malo u konkurentnim sistemima. U slučaju kada komponente rade
paralelno u više jezgara, ovo komunikaciono opterećenje je veliko. Stoga,
paralelni programi ne rezultiraju uvek bržim vremenom izvršavanja!

Podrška za konkurentnost u Gou
-------------------------------
Konkurentnost je sastavni deo programskog jezika Go. Konkurentnost se u Go-u
obrađuje pomoću goroutina i kanala. O njima ćemo detaljno razgovarati u
narednim tutorijalima.

Gorutine
========

Gorutine su funkcije ili metode koje se izvršavaju konkurentno sa drugim
funkcijama ili metodama. Gorutine se mogu smatrati laganim nitima. Troškovi
kreiranja gorutine su mali u poređenju sa jednom niti. Stoga je uobičajeno da
Go aplikacije imaju hiljade gorutina koje se izvršavaju konkurentno.

Prednosti Gorutina u odnosu na niti
-----------------------------------

- Gorutine su izuzetno jeftine u poređenju sa nitima. Veličina steka je samo
  nekoliko kb i stek može da raste i smanjuje se u skladu sa potrebama
  aplikacije, dok u slučaju niti veličina steka mora biti navedena i fiksna je.
- Gorutine su multipleksirane na manji broj OS niti. Može postojati samo jedna
  nit u programu sa hiljadama gorutina. Ako bilo koja gorutina u toj niti
  blokira čekanjem korisničkog unosa, onda se kreira druga OS nit, a preostale
  gorutine se premeštaju u novu OS nit. O svemu tome brine okruženje za
  izvršavanje, a mi kao programeri smo apstrahovani od ovih složenih detalja i
  dobijamo čist API za rad sa konkurentnošću.
- Gorutine komuniciraju pomoću kanala. Kanali su po svojoj prirodi sprečavaju
  pojavu uslova trke pri pristupu deljenoj memoriji pomoću gorutina. Kanali se
  mogu smatrati cevovodom pomoću kog Gorutine komuniciraju. Kanale ćemo
  detaljno razmotriti u sledećem tutorijalu.

Kako pokrenuti Gorutinu?
------------------------
Dodajte ključnu reč go ispred poziva funkcije ili metode i imaćete novu
gorutinu koja se pokreće konkurentno.

Hajde da napravimo gorutinu :)
*/

func hello() {
	fmt.Println("Hello world goroutine")
}

func concGoFunc() {

	fmt.Println("\n --- concGoFunc ---")

	go hello()
	fmt.Println("main function")
}

/*
go hello() pokreće novu gorutinu. Sada će se hello() funkcija izvršavati
konkurentno sa main()gorutinom - funkcijom. Main funkcija se izvršava u
sopstvenoj gorutini i naziva se main gorutina.

Pokrenite ovaj program i imaćete iznenađenje!

Ovaj program prikazuje samo tekst "main function". Šta se desilo sa gorutinom
koju smo pokrenuli? Moramo da razumemo dva glavna svojstva gorutina da bismo
razumeli zašto se to dešava.

- Kada se pokrene nova gorutina, poziv gorutine odmah vraća. Za razliku od
  funkcija, kontrola ne čeka da gorutina završi izvršavanje. Kontrola se odmah
  vraća na sledeći red koda nakon poziva nove gorutine i sve povratne vrednosti
  iz Gorutine se ignorišu.
- Main gorutina treba da bude pokrenuta da bi se pokrenule sve ostale gorutine.
  Ako se main gorutina prekine, program će biti prekinut i nijedna druga
  gorutina se neće pokrenuti.

Pretpostavljam da sada možete razumeti zašto se naša gorutina nije pokrenula.
Nakon poziva go hello(), kontrola se odmah vratila na sledeću liniju koda bez
čekanja da se završi hello gorutina i ispisala "main function". Zatim je glavna
gorutina prekinuta jer nije bilo drugog koda za izvršavanje i stoga hello
gorutina nije dobila priliku da se pokrene.

Hajde da ovo sada popravimo.
*/

func concGoWithTimeOutFunc() {

	fmt.Println("\n --- concGoWithTimeOutFunc---")

	go hello()
	time.Sleep(1 * time.Second) // Make pause 1 sec.
	fmt.Println("main function")
}

/*
U gornjem programu pozvali smo metod "Sleep" paketa "time" koji uspava main go
rutinu u kojoj se izvršava. U ovom slučaju, main gorutina se stavlja u stanje
mirovanja na 1 sekundu. Sada poziv metode go hello() ima dovoljno vremena za
izvršenje pre nego što se glavna gorutina završi. Ovaj program prvo ispisuje
"Hello world goroutine", čeka 1 sekundu, a zatim ispisuje "main function".

Ovaj način korišćenja režima spavanja (sleep) u glavnoj gorutini da bi se
sačekalo da druge gorutine završe svoje izvršavanje je trik koji koristimo da
bismo razumeli kako gorutine funkcionišu. Kanali se mogu koristiti za
blokiranje main gorutine dok sve ostale gorutine ne završe svoje izvršavanje.
O kanalima ćemo razgovarati u sledećem tutorijalu.

Pokretanje više gorutina
------------------------
Hajde da napišemo još jedan program koji pokreće više gorutina kako bismo ih
bolje razumeli.
*/

func numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}

func concGoMultiFunc() {

	fmt.Println("\n concGoMultiFunc ---")

	go numbers()
	go alphabets()
	time.Sleep(3000 * time.Millisecond)
	fmt.Println()
	fmt.Println("main terminated")
}

/*
Gore navedeni program pokreće dve gorutine. Ove dve gorutine rade konkurentno.
numbers gorutina je u početku u stanju mirovanja 250 milisekundi, a zatim
ispisuje 1, zatim ponovo prelazi u stanje mirovanja i ispisuje 2, i isti ciklus
se ponavlja dok se ne ispiše 5. Slično tome, alphabets gorutina ispisuje abecde
od a do e i ima 400 milisekundi vremena mirovanja. Main gorutina pokreće
numbers i alphabets gorutine, miruje 3000 milisekundi, a zatim se završava.

Ovaj program ispisuje

1 a 2 3 b 4 c 5 d e main terminated

Sledeća slika prikazuje kako ovaj program radi. Molimo vas da otvorite sliku u
novoj kartici radi bolje vidljivosti :)

Prvi deo slike u plavoj boji predstavlja numbers gorutine, drugi deo u bordo
boji predstavlja alfabets gorutine, treći deo u zelenoj boji predstavlja main
gorutinu, a poslednji deo u crnoj boji spaja sva gore navedena tri i pokazuje
nam kako program radi. Nizovi poput 0 ms, 250 ms na vrhu svakog polja
predstavljaju vreme u milisekundama, a izlaz je predstavljen na dnu svakog
polja kao 1, 2, 3 i tako dalje. Plavo polje nam govori da je ispisano 1 posle
250 ms, 2 je ispisano posle 500 ms i tako dalje. Dno poslednjeg crnog polja
sadrži vrednosti 1 a 2 3 b 4 c 5 d e main terminated koje su takođe izlaz
programa. Slika je sama po sebi jasna i moći ćete da razumete kako program radi.

0   250  500  750 1000 1250
|    |    |    |    |    |	<< numbers
     1    2    3    4    5
0       400    800    1200    1600     2000
|       |       |       |       |       |	<< alphabets
        a       b       c       d       e
0                                                         3000
|                                                           |    << main

Kanali
======

Kanali se mogu smatrati cevima pomoću kojih gorutine komuniciraju. Slično kao
što voda teče od jednog do drugog kraja u cevi, podaci se mogu slati sa jednog
kraja i primati sa drugog kraja kanala.

Deklarisanje kanala
-------------------
Svaki kanal ima tip koji je povezan sa njim. Ovaj tip je tip podataka koje
kanal može da prenosi. Nije dozvoljeno da se prenosi bilo koji drugi tip pomoću
kanala.

chan T je kanal tipa T

Nulta vrednost kanala je nil. nil kanali nisu od koristi i stoga kanal mora
biti definisan koristeći make, slično mapama i isečcima.

Hajde da napišemo kod koji deklariše kanal.
*/

func concChannelFunc() {

	fmt.Println("\n --- concChannelFunc ---")

	var a chan int
	fmt.Println("channel a is nil, going to define it")
	a = make(chan int)
	fmt.Printf("Type of a is %T\n", a)

}

/*
Kanal adeklarisan u gornjem programu je nil. Nulta vrednost kanala nil. Stoga
se izvršavaju naredbe unutar if uslova i kanal se definiše pomoću make. U
gornjem programu je kanal tipa int. Ovaj program će ispisati,

	>> channel a is nil, going to define it
	>> Type of a is chan int

Kao i obično, skraćena deklaracija je takođe validan i koncizan način za
definisanje kanala.

	>> a := make(chan int)

Gornja linija koda takođe definiše int kanal a.

Slanje i primanje sa kanala
---------------------------
Sintaksa za slanje i primanje podataka sa kanala je data u nastavku,

	>> data := <- a // read from channel a
	>> a <- data // write to channel a

Smer strelice u odnosu na kanal određuje da li se podaci šalju ili primaju.
U prvom redu, strelica pokazuje ka spolja, od a i stoga čitamo iz kanala a
i čuvamo vrednost u promenljivoj data.

U drugom redu, strelica pokazuje ka a i stoga pišemo na kanal a. Slanje i
primanje se podrazumevano blokiraju.

Slanje i prijem na kanal se podrazumevano blokira. Šta to znači? Kada se podaci
šalju na kanal, kontrola je blokirana u naredbi za slanje dok neka druga
gorutina ne pročita podatke sa tog kanala. Slično, kada se podaci čitaju sa
kanala, čitanje je blokirano dok neka gorutina ne upiše podatke u taj kanal.

Ovo svojstvo kanala je ono što pomaže gorutinama da efikasno komuniciraju bez
upotrebe eksplicitnih zaključavanja ili uslovnih promenljivih koje su prilično
uobičajene u drugim programskim jezicima.

U redu je ako ovo sada nema smisla. U narednim odeljcima će biti jasnije kako se
kanali podrazumevano blokiraju.

Primer programa sa kanalima
---------------------------
Hajde da napišemo program da bismo razumeli kako gorutine komuniciraju koristeći
kanale. Zapravo ćemo prepisati program koji smo napisali kada smo učili o
gorutinama koristeći kanale.

Dozvolite mi da ovde citiram program iz poslednjeg tutorijala.

	func hello() {
	    fmt.Println("Hello world goroutine")
	}

	func main() {
	    go hello()
	    time.Sleep(1 * time.Second)
	    fmt.Println("main function")
	}

Ovo je program iz prethodnog tutorijala. Ovde koristimo sleep da bismo naterali
main gorutinu da sačeka da se završi hello gorutina. Ako vam ovo nema smisla,
preporučujem vam da pročitate tutorijal o gorutinama.

Prepisaćemo gornji program koristeći kanale.
*/

func helloChannel(done chan bool) {
	fmt.Println("Hello world goroutine")
	done <- true // Write to channel done
}

func concGoChannelFunc() {

	fmt.Println("\n --- concGoChannelFunc ---")

	done := make(chan bool)
	go helloChannel(done)
	<-done // This is blocking line, waits for read data from data channel
	fmt.Println("main function")
}

/*
U gornjem programu, kreiramo done kanal tipa bool i prosleđujemo ga kao
parametar gorutini hello. Potom primamo podatke iz done kanala. Ova linija
koda je blokirajuća, što znači da dok neka gorutina hello ne upiše podatke u
done kanal, kontrola neće preći na sledeću liniju koda. Stoga se eliminiše
potreba za "time.Sleep" pozivom koji je bio prisutan u originalnom programu
da bi se sprečio izlazak iz glavne gorutine.

Linija koda <-done prima podatke iz kanala "done", ali ih ne koristi niti čuva
ni u jednoj promenljivoj. Ovo je sasvim legalno.

Sada imamo našu glavnu gorutinu blokiranu koja čeka podatke na kanalu. hello
gorutina prima ovaj kanal kao parametar, štampa, Hello world goroutine a zatim
upisuje u done kanal. Kada se ovo pisanje završi, glavna gorutina prima podatke
sa završenog kanala, ona se deblokira, a zatim se ispisuje tekst -
"main function".

Ovaj program ispisuje

	>> Hello world goroutine
	>> main function

Hajde da modifikujemo ovaj program uvođenjem "sleep"-a u hello goroutinu kako
bismo bolje razumeli ovaj koncept blokiranja.
*/
func helloChannelSleep(done chan bool) {
	fmt.Println("hello go routine is going to sleep 4 secs")
	time.Sleep(4 * time.Second)
	fmt.Println("hello go routine awake and going to write to done")
	done <- true
}

func concGoChannelSleepFunc() {

	fmt.Println("\n --- concGoChannelSleepFunc ---")

	done := make(chan bool)
	fmt.Println("Main going to call hello go goroutine")
	go helloChannelSleep(done)
	<-done
	fmt.Println("Main received data")
}

/*
U gornjem programu, uveli smo spavanje od 4 sekunde u helloChannelSleep funkciji.
Ovaj program će prvo ispisati:

	>> "Main going to call hello go goroutine".
	>> "hello go routine is going to sleep 4 secs".
	>> "hello go routine awake
	>> "Main received data".

Još jedan primer za kanale
--------------------------
Hajde da napišemo još jedan program da bismo bolje razumeli kanale. Ovaj program
će ispisati zbir kvadrata i kubova pojedinačnih cifara broja.

# Na primer, ako je ulaz 123, onda će ovaj program izračunati izlaz kao

square = (1 * 1) + (2 * 2) + (3 * 3)
cubes = (1 * 1 * 1) + (2 * 2 * 2) + (3 * 3 * 3)
results = square + cubes = 50

Strukturiraćemo program tako da se kvadrati izračunavaju u posebnoj gorutini,
kubovi u posebnoj Gorutini, a konačno sabiranje se dešava u glavnoj gorutini.
*/

func calcSquares(number int, squareop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	squareop <- sum
}

func calcCubes(number int, cubeop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	cubeop <- sum
}

func concGoCalcSquaresAndCubes() {

	fmt.Println("\n ---concGoCalcSquareAndCubes ---")

	number := 567
	sqrch := make(chan int)
	cubech := make(chan int)
	go calcSquares(number, sqrch)
	go calcCubes(number, cubech)
	squares, cubes := <-sqrch, <-cubech
	fmt.Println("Final output", squares+cubes)
}

/*
Funkcija calcSquares izračunava zbir kvadrata pojedinačnih cifara broja i šalje
ga kanalu squareop. Slično tome, calcCubes funkcija izračunava zbir kubova
pojedinačnih cifara broja i šalje ga kanalu cubeop.

Ove dve funkcije se izvršavaju kao odvojene gorutine i svakoj se prosleđuje kanal
za pisanje kao parametar. Glavna gorutina čeka podatke sa oba ova kanala. Kada
se podaci prime sa oba kanala, oni se čuvaju u promenljivim squares i cubes, a
konačni izlaz se izračunava i štampa. Ovaj program će štampati:

	>> Final output 1536

Zastoj
------
Jedan važan faktor koji treba uzeti u obzir pri korišćenju kanala je zastoj.
Ako gorutina šalje podatke na kanal, onda se očekuje da neka druga gorutina
treba da prima podatke. Ako se to ne desi, program će paničiti tokom
izvršavanja sa Deadlock-om.

Slično tome, ako gorutina čeka da primi podatke sa kanala, onda se očekuje da
neka druga gorutina piše podatke na tom kanalu, u suprotnom će program paničiti.
*/
/*
func channelPanic() {
	ch := make(chan int)
	ch <- 5
}
*/
/*
U gornjem programu, kreiran je kanal ch i šaljemo podatak 5 na kanal. U ovom
programu nijedna druga gorutina ne prima podatke sa kanala ch. Stoga će ovaj
program prijaviti sledeću grešku tokom izvršavanja.

	>> fatal error: all goroutines are asleep - deadlock!

		goroutine 1 [chan send]:
		main.main()
			/tmp/sandbox046150166/prog.go:6 +0x50

Jednosmerni kanali
------------------
Svi kanali o kojima smo do sada govorili su dvosmerni kanali, odnosno podaci se
preko njih mogu i slati i primati. Takođe je moguće kreirati jednosmerne kanale,
odnosno kanale koji samo šalju ili samo primaju podatke.
*/
/*
func sendData(sendch chan<- int) {
	sendch <- 10
}

func unidirectionalCannels() {
	sendch := make(chan<- int)
	go sendData(sendch)
	fmt.Println(<-sendch)
}
*/
/*
U gornjem programu, kreiramo kanal samo za slanje sendch. chan<- int označava
kanal samo za slanje jer strelica pokazuje na chan. Pokušavamo da primimo
podatke sa kanala samo za slanje. Ovo nije dozvoljeno i kada se program
pokrene, kompajler će se žaliti rekavši,

	>>  invalid operation: cannot receive from send-only channel sendch
		(variable of type chan<- int)

Sve je u redu, ali koja je svrha pisanja na kanal samo za slanje ako se sa njega
ne može čitati!

Ovde dolazi do izražaja konverzija kanala. Moguće je konvertovati dvosmerni kanal
 u kanal samo za slanje ili samo za prijem, ali ne i obrnuto.
*/

func sendData(sendch chan<- int) {
	sendch <- 10
}

func concConvBiToUniChannel() {

	fmt.Println("\n --- concConvBiToUniChannel --- ")

	chnl := make(chan int)
	go sendData(chnl)   // Conversion
	fmt.Println(<-chnl) // Conversion
}

/*
chnl je dvosmerni kanal. On se prosleđuje kao parametar gorutini sendData.
sendData funkcija konvertuje ovaj dvosmerni kanal u kanal samo za slanje
u parametru sendch chan<- int. Dakle, sada se kanal šalje kao unudirekcionalni
unutar sendData gorutine, ali je dvosmeran u glavnoj Gorutini. Ovaj program će
ispisati 10 kao izlaz.

Zatvaranje kanala i range petlje na kanalima
--------------------------------------------
Pošiljaoci imaju mogućnost da zatvore kanal kako bi obavestili primaoce da se
više podaci neće slati preko tog kanala.

Prijemnici mogu koristiti dodatnu promenljivu dok primaju podatke sa kanala da
bi proverili da li je kanal zatvoren.

v, ok := <- ch

U gornjoj izjavi ok je tačno ako je vrednost uspešno primljena. Ako je ok
netačno, to znači da čitamo iz zatvorenog kanala. Vrednost pročitana iz
zatvorenog kanala biće nulta vrednost tipa kanala. Na primer, ako je kanal
tipa int, onda će vrednost primljena iz zatvorenog kanala biti 0.
*/

func producer(chnl chan int) {
	for i := 0; i < 10; i++ {
		chnl <- i
	}
	close(chnl)
}

func concGoChannelClose() {

	fmt.Println("\n --- concGoChannelClose --- ")

	ch := make(chan int)
	go producer(ch)
	for {
		v, ok := <-ch
		if !ok {
			break
		}
		fmt.Println("Received ", v, ok)
	}
}

/*
U gornjem programu, producer gorutina upisuje vrednosti od 0 do 9 u chnl kanal,
a zatim ga zatvara. Glavna funkcija ima beskonačnu for petlju koja proverava da
li je kanal zatvoren koristeći promenljivu ok. Ako je ok false, to znači da je
kanal zatvoren i stoga je petlja prekinuta. U suprotnom, primljena vrednost i
vrednost ok se ispisuju. Ovaj program ispisuje,

	>> Received  0 true
	>> Received  1 true
	>> Received  2 true
	>> Received  3 true
	>> Received  4 true
	>> Received  5 true
	>> Received  6 true
	>> Received  7 true
	>> Received  8 true
	>> Received  9 true

Oblik for range petlje može se koristiti za primanje vrednosti iz kanala dok
se on ne zatvori.

Hajde da prepišemo gornji program koristeći petlju for range.
*/
func producerForRange(chnl chan int) {
	for i := 0; i < 10; i++ {
		chnl <- i
	}
	close(chnl)
}

func concGoChannelCloseForRange() {

	fmt.Println("\n --- concGoChannelCloseForRange ---")

	ch := make(chan int)
	go producerForRange(ch)

	for v := range ch {
		fmt.Println("Received ", v)
	}
}

/*
Petlja for range prima podatke iz ch kanala dok se ne zatvori. Kada
se ch zatvori, petlja automatski izlazi. Ovaj program ispisuje,

Received  0
Received  1
Received  2
Received  3
Received  4
Received  5
Received  6
Received  7
Received  8
Received  9

Još jedan primer za kanale može se prepisati sa većom mogućnošću ponovne
upotrebe koda korišćenjem petlje za opseg.

Ako pažljivije pogledate program, primetićete da se kod koji pronalazi
pojedinačne cifre broja ponavlja i u calcSquares i calcCubes funkciji.
Premestićemo taj kod u njegovu zasebnu funkciju i pozivati je konkurentno.
*/

func digits(number int, dchnl chan int) {
	for number != 0 {
		digit := number % 10
		dchnl <- digit
		number /= 10
	}
	close(dchnl)
}

func calcSquares2(number int, squareop chan int) {
	sum := 0
	dch := make(chan int)
	go digits(number, dch)
	for digit := range dch {
		sum += digit * digit
	}
	squareop <- sum
}

func calcCubes2(number int, cubeop chan int) {
	sum := 0
	dch := make(chan int)
	go digits(number, dch)
	for digit := range dch {
		sum += digit * digit * digit
	}
	cubeop <- sum
}

func concGoMultiFunc2() {

	fmt.Println("\n --- concGoMultiFunc2 ---")

	number := 333
	sqrch := make(chan int)
	cubech := make(chan int)
	go calcSquares2(number, sqrch)
	go calcCubes2(number, cubech)
	squares, cubes := <-sqrch, <-cubech
	fmt.Println("Final output", squares+cubes)
}

/*
Funkcija digit su gornjem programu sada sadrži logiku za dobijanje pojedinačnih
cifara iz broja i pozivaju je obe funkcije calcSquares i calcCubes konkurentno.
Kada u broju više nema cifara, kanal se zatvara. Gorutine calcSquares i calcCubes
slušaju na svojim kanalima koristeći for range petlju dok se ne zatvore. Ostatak
programa je isti.

Ovaj program će takođe ispisati:

	>> Final output 1536
*/
func ConcFunc() {
	fmt.Println("\n --- Intro to concurency  ---")

	concGoFunc()
	concGoWithTimeOutFunc()
	concGoMultiFunc()
	concChannelFunc()
	concGoChannelFunc()
	concGoChannelSleepFunc()
	concGoCalcSquaresAndCubes()
	// channelPanic()
	// unidirectionalCannels()
	concConvBiToUniChannel()
	concGoChannelClose()
	concGoChannelCloseForRange()
	concGoMultiFunc2()
}
