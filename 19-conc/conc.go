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
*/

func ConcFunc() {
	fmt.Println("\n --- Uvod u konkurentnost  ---")

	goroutineFunc()
	goroutineWithTimeOutFunc()
	goroutineMultiFunc()
	channelFunc()
}

/*
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

func goroutineFunc() {
	fmt.Println("\n --- Goroutine ---")
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

func goroutineWithTimeOutFunc() {

	fmt.Println("\n --- Goroutine With TimeOut Func---")

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

func goroutineMultiFunc() {

	fmt.Println("\n Multi goroutine ---")

	go numbers()
	go alphabets()
	time.Sleep(3000 * time.Millisecond)
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
|    |    |    |    |    |										<< numbers
     1    2    3    4    5
0       400    800    1200    1600     2000
|       |       |       |       |       |						<< alphabets
        a       b       c       d       e
0                                                        3000
|                                                          |    << main

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

func channelFunc() {

	fmt.Println("\n ---Funkcija kanala")

	var a chan int
	if a == nil {
		fmt.Println("channel a is nil, going to define it")
		a = make(chan int)
		fmt.Printf("Type of a is %T", a)
	}
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

1data := <- a // read from channel a
2a <- data // write to channel a

Smer strelice u odnosu na kanal određuje da li se podaci šalju ili primaju.

U prvom redu, strelica pokazuje ka spolja ai stoga čitamo iz kanala ai čuvamo vrednost u promenljivoj data.

U drugom redu, strelica pokazuje ka ai stoga pišemo na kanal a.
Slanje i primanje se podrazumevano blokiraju

Slanje i prijem na kanal se podrazumevano blokira. Šta to znači? Kada se podaci šalju na kanal, kontrola je blokirana u naredbi za slanje dok neka druga Gorutina ne pročita podatke sa tog kanala. Slično, kada se podaci čitaju sa kanala, čitanje je blokirano dok neka Gorutina ne upiše podatke u taj kanal.

Ovo svojstvo kanala je ono što pomaže Gorutinama da efikasno komuniciraju bez upotrebe eksplicitnih zaključavanja ili uslovnih promenljivih koje su prilično uobičajene u drugim programskim jezicima.

U redu je ako ovo sada nema smisla. U narednim odeljcima će biti jasnije kako se kanali podrazumevano blokiraju.
Primer programa kanala

Dosta teorije :). Hajde da napišemo program da bismo razumeli kako Gorutine komuniciraju koristeći kanale.

Zapravo ćemo prepisati program koji smo napisali kada smo učili o Gorutinama koristeći kanale ovde.

Dozvolite mi da ovde citiram program iz poslednjeg tutorijala.

 1package main
 2
 3import (
 4    "fmt"
 5    "time"
 6)
 7
 8func hello() {
 9    fmt.Println("Hello world goroutine")
10}
11func main() {
12    go hello()
13    time.Sleep(1 * time.Second)
14    fmt.Println("main function")
15}

Izvedite program na igralištu

Ovo je bio program iz prethodnog tutorijala. Ovde koristimo sleep da bismo naterali glavnu Gorutinu da sačeka da se završi hello Gorutina. Ako vam ovo nema smisla, preporučujem vam da pročitate tutorijal o Gorutinama.

Prepisaćemo gornji program koristeći kanale.

 1package main
 2
 3import (
 4	"fmt"
 5)
 6
 7func hello(done chan bool) {
 8	fmt.Println("Hello world goroutine")
 9	done <- true
10}
11func main() {
12	done := make(chan bool)
13	go hello(done)
14	<-done
15	fmt.Println("main function")
16}

Izvedite program na igralištu

U gornjem programu, kreiramo donebool kanal u liniji br. 12 i prosleđujemo ga kao parametar Gorutini hello. U liniji br. 14 primamo podatke iz donekanala. Ova linija koda je blokirajuća, što znači da dok neka Gorutina ne upiše podatke u donekanal, kontrola neće preći na sledeću liniju koda. Stoga se eliminiše potreba za ` time.Sleepwhich` koji je bio prisutan u originalnom programu da bi se sprečio izlazak glavne Gorutine.

Linija koda <-doneprima podatke iz kanala „done“, ali ih ne koristi niti čuva ni u jednoj promenljivoj. Ovo je sasvim legalno.

Sada imamo našu mainGorutinu blokiranu i čeka podatke na završenom kanalu. helloGorutina prima ovaj kanal kao parametar, štampa, Hello world goroutinea zatim upisuje u donekanal. Kada se ovo pisanje završi, glavna Gorutina prima podatke sa završenog kanala, oni se deblokiraju, a zatim se ispisuje tekst glavne funkcije .

Ovaj program izlazi

Hello world goroutine
main function

Hajde da modifikujemo ovaj program uvođenjem sleep-a u helloGoroutine kako bismo bolje razumeli ovaj koncept blokiranja.

 1package main
 2
 3import (
 4	"fmt"
 5	"time"
 6)
 7
 8func hello(done chan bool) {
 9	fmt.Println("hello go routine is going to sleep")
10	time.Sleep(4 * time.Second)
11	fmt.Println("hello go routine awake and going to write to done")
12	done <- true
13}
14func main() {
15	done := make(chan bool)
16	fmt.Println("Main going to call hello go goroutine")
17	go hello(done)
18	<-done
19	fmt.Println("Main received data")
20}

Trčanje na igralištu

U gornjem programu, uveli smo spavanje od 4 sekunde u hellofunkciju u liniji br. 10.

Ovaj program će prvo ispisati Main going to call hello go goroutine. Zatim će se pokrenuti hello Goroutine i ispisaće se hello go routine is going to sleep. Nakon što se ovo ispiše, helloGoroutine će biti u stanju mirovanja 4 sekunde i tokom tog vremena mainGoroutine će biti blokirana jer čeka podatke iz kanala „done“ u liniji br. 18. <-doneNakon 4 sekunde hello go routine awake and going to write to donebiće ispisano , a zatim Main received data.
Još jedan primer za kanale

Hajde da napišemo još jedan program da bismo bolje razumeli kanale. Ovaj program će ispisati zbir kvadrata i kubova pojedinačnih cifara broja.

Na primer, ako je ulaz 123, onda će ovaj program izračunati izlaz kao

kvadrati = (1 * 1) + (2 * 2) + (3 * 3) kocke = (1 * 1 * 1) + (2 * 2 * 2) + (3 * 3 * 3) izlaz = kvadrati + kocke = 50

Strukturiraćemo program tako da se kvadrati izračunavaju u posebnoj Gorutini, kubovi u drugoj Gorutini, a konačno sabiranje se dešava u glavnoj Gorutini.

 1package main
 2
 3import (
 4    "fmt"
 5)
 6
 7func calcSquares(number int, squareop chan int) {
 8    sum := 0
 9    for number != 0 {
10        digit := number % 10
11        sum += digit * digit
12        number /= 10
13    }
14    squareop <- sum
15}
16
17func calcCubes(number int, cubeop chan int) {
18    sum := 0
19    for number != 0 {
20        digit := number % 10
21        sum += digit * digit * digit
22        number /= 10
23    }
24    cubeop <- sum
25}
26
27func main() {
28    number := 589
29    sqrch := make(chan int)
30    cubech := make(chan int)
31    go calcSquares(number, sqrch)
32    go calcCubes(number, cubech)
33    squares, cubes := <-sqrch, <-cubech
34    fmt.Println("Final output", squares + cubes)
35}

Izvedite program na igralištu

Funkcija calcSquaresu liniji br. 7 izračunava zbir kvadrata pojedinačnih cifara broja i šalje ga kanalu squareop. Slično tome, calcCubesfunkcija u liniji br. 17 izračunava zbir kubova pojedinačnih cifara broja i šalje ga kanalu cubeop.

Ove dve funkcije se izvršavaju kao odvojene Gorutine u liniji br. 31 i 32 i svakoj se prosleđuje kanal za pisanje kao parametar. Glavna Gorutina čeka podatke sa oba ova kanala u liniji br. 33. Kada se podaci prime sa oba kanala, oni se čuvaju u promenljivim squaresi cubes, a konačni izlaz se izračunava i štampa. Ovaj program će štampati

Final output 1536

Zastoj

Jedan važan faktor koji treba uzeti u obzir pri korišćenju kanala je zastoj. Ako Gorutina šalje podatke na kanalu, onda se očekuje da neka druga Gorutina treba da prima podatke. Ako se to ne desi, program će paničiti tokom izvršavanja sa Deadlock.

Slično tome, ako Gorutina čeka da primi podatke sa kanala, onda se očekuje da neka druga Gorutina zapiše podatke na tom kanalu, u suprotnom će program paničiti.

1package main
2
3
4func main() {
5	ch := make(chan int)
6	ch <- 5
7}

Izvedite program na igralištu

U gornjem programu, kreiran je kanal chi šaljemo podatke 5na kanal u liniji br. 6. ch <- 5U ovom programu nijedna druga Gorutina ne prima podatke sa kanala ch. Stoga će ovaj program prijaviti sledeću grešku tokom izvršavanja.

fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.main()
	/tmp/sandbox046150166/prog.go:6 +0x50

Jednosmerni kanali

Svi kanali o kojima smo do sada govorili su dvosmerni kanali, odnosno podaci se preko njih mogu i slati i primati. Takođe je moguće kreirati jednosmerne kanale, odnosno kanale koji samo šalju ili primaju podatke.

 1package main
 2
 3import "fmt"
 4
 5func sendData(sendch chan<- int) {
 6	sendch <- 10
 7}
 8
 9func main() {
10	sendch := make(chan<- int)
11	go sendData(sendch)
12	fmt.Println(<-sendch)
13}

Izvedite program na igralištu

U gornjem programu, kreiramo kanal samo za slanje sendchu liniji br. 10. chan<- intoznačava kanal samo za slanje jer strelica pokazuje na chan. Pokušavamo da primimo podatke sa kanala samo za slanje u liniji br. 12. Ovo nije dozvoljeno i kada se program pokrene, kompajler će se žaliti rekavši,

./prog.go:12:14: nevažeća operacija: <-sendch (primanje od tipa samo za slanje chan<- int)

Sve je u redu, ali koja je svrha pisanja na kanal samo za slanje ako se sa njega ne može čitati!

Ovde dolazi do izražaja konverzija kanala. Moguće je konvertovati dvosmerni kanal u kanal samo za slanje ili samo za prijem, ali ne i obrnuto.

 1package main
 2
 3import "fmt"
 4
 5func sendData(sendch chan<- int) {
 6	sendch <- 10
 7}
 8
 9func main() {
10	chnl := make(chan int)
11	go sendData(chnl)
12	fmt.Println(<-chnl)
13}

Izvedite program na igralištu

U liniji br. 10 gornjeg programa, chnlkreiran je dvosmerni kanal. On se prosleđuje kao parametar Gorutini sendDatau liniji br. 11. sendDataFunkcija konvertuje ovaj kanal u kanal samo za slanje u liniji br. 5 u parametru sendch chan<- int. Dakle, sada se kanal samo šalje unutar sendDataGorutine, ali je dvosmeran u glavnoj Gorutini. Ovaj program će ispisati 10kao izlaz.
Zatvaranje kanala i petlje za opseg na kanalima

Pošiljaoci imaju mogućnost da zatvore kanal kako bi obavestili primaoce da se više podaci neće slati preko tog kanala.

Prijemnici mogu koristiti dodatnu promenljivu dok primaju podatke sa kanala da bi proverili da li je kanal zatvoren.

v, ok := <- ch

U gornjoj izjavi okje tačno ako je vrednost primljena uspešnom operacijom slanja na kanal. Ako okje netačno, to znači da čitamo iz zatvorenog kanala. Vrednost pročitana iz zatvorenog kanala biće nulta vrednost tipa kanala. Na primer, ako je kanal int, onda će vrednost primljena iz zatvorenog kanala biti 0.

 1package main
 2
 3import (
 4	"fmt"
 5)
 6
 7func producer(chnl chan int) {
 8	for i := 0; i < 10; i++ {
 9		chnl <- i
10	}
11	close(chnl)
12}
13func main() {
14	ch := make(chan int)
15	go producer(ch)
16	for {
17		v, ok := <-ch
18		if ok == false {
19			break
20		}
21		fmt.Println("Received ", v, ok)
22	}
23}

Izvedite program na igralištu

U gornjem programu, producerGorutina upisuje vrednosti od 0 do 9 u chnlkanal, a zatim ga zatvara. Glavna funkcija ima beskonačnu forpetlju u liniji br. 16 koja proverava da li je kanal zatvoren koristeći promenljivu oku liniji br. 18. Ako okje false, to znači da je kanal zatvoren i stoga je petlja prekinuta. U suprotnom, primljena vrednost i vrednost okse ispisuju. Ovaj program ispisuje,

Received  0 true
Received  1 true
Received  2 true
Received  3 true
Received  4 true
Received  5 true
Received  6 true
Received  7 true
Received  8 true
Received  9 true

Oblik for range u for petlji može se koristiti za primanje vrednosti iz kanala dok se on ne zatvori.

Hajde da prepišemo gornji program koristeći petlju for range.

 1package main
 2
 3import (
 4	"fmt"
 5)
 6
 7func producer(chnl chan int) {
 8	for i := 0; i < 10; i++ {
 9		chnl <- i
10	}
11	close(chnl)
12}
13func main() {
14	ch := make(chan int)
15	go producer(ch)
16	for v := range ch {
17		fmt.Println("Received ",v)
18	}
19}

Izvedite program na igralištu

Petlja for rangeu liniji br. 16 prima podatke iz chkanala dok se ne zatvori. Kada chse zatvori, petlja se automatski izlazi. Ovaj program izbacuje,

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

Program iz odeljka Još jedan primer za kanale može se prepisati sa većom mogućnošću ponovne upotrebe koda korišćenjem petlje za opseg.

Ako pažljivije pogledate program, primetićete da se kod koji pronalazi pojedinačne cifre broja ponavlja i u calcSquaresfunkciji i calcCubesu funkciji. Premestićemo taj kod u njegovu zasebnu funkciju i pozivati je istovremeno.

 1package main
 2
 3import (
 4	"fmt"
 5)
 6
 7func digits(number int, dchnl chan int) {
 8	for number != 0 {
 9		digit := number % 10
10		dchnl <- digit
11		number /= 10
12	}
13	close(dchnl)
14}
15func calcSquares(number int, squareop chan int) {
16	sum := 0
17	dch := make(chan int)
18	go digits(number, dch)
19	for digit := range dch {
20		sum += digit * digit
21	}
22	squareop <- sum
23}
24
25func calcCubes(number int, cubeop chan int) {
26	sum := 0
27	dch := make(chan int)
28	go digits(number, dch)
29	for digit := range dch {
30		sum += digit * digit * digit
31	}
32	cubeop <- sum
33}
34
35func main() {
36	number := 589
37	sqrch := make(chan int)
38	cubech := make(chan int)
39	go calcSquares(number, sqrch)
40	go calcCubes(number, cubech)
41	squares, cubes := <-sqrch, <-cubech
42	fmt.Println("Final output", squares+cubes)
43}

Izvedite program na igralištu

Funkcija digitsu gornjem programu sada sadrži logiku za dobijanje pojedinačnih cifara iz broja i pozivaju je obe funkcije calcSquares, i , calcCubesistovremeno. Kada u broju više nema cifara, kanal se zatvara u liniji br. 13. Gorutine calcSquaresi calcCubesslušaju na svojim kanalima koristeći for rangepetlju dok se ne zatvore. Ostatak programa je isti. Ovaj program će takođe ispisati

Final output 1536

*/
