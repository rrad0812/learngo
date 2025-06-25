package files

import (
	"bufio"
	"fmt"
	"io"
	fh "learngo/14-files/fh"
	fh1 "learngo/14-files/fh1"
	fh2 "learngo/14-files/fh2"
	fh3 "learngo/14-files/fh3"
	"log"
	"os"
)

/*
Čitanje datoteka
================

Čitanje celokupne datoteke u memoriju
-------------------------------------
Jedna od najosnovnijih operacija datoteka čita celu datoteku u memoriju.
To se radi uz pomoć funkcije "ReadFile" iz paketa "os".

Pročitajmo datoteku i štampjamo njen sadržaj
............................................
Napravio sam direktorijum filehandling u 14-files

	mkdir /home/radosav/go/src/leargo/14-files/filehandling

Imam text datoteku test.txt koji ćemo pročitati iz našeg programa. test.txt
sadrži sledeći string

	>> Hello World. Welcome to file handling in Go.

Evo moje strukture direktorijuma.

├── $GOPATH (/home/radosav/go)
	└── src
		└── learngo
			└── 14-files
				├── goreadFiles.Go
				└── filehandling
					└── filehandling.go

Idemo odmah do koda. Kreirajmo datoteku filehandling.go sa sledećim sadržajem:

package filehandling

import (
	"fmt"
	"os"
)

func ReadAll() {
	contents, err := os.ReadFile("test.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Contents of file:", string(contents))
}

Nakon toga pokrenimo compile/run ciklus sa funkcijom:
*/

func ReadAllOfLocalFile() {

	fmt.Println("\n --- ReadAllOfLocalFile ---")

	fh.ReadAll()

}

/*
Molimo pokrenite ovaj program iz svog lokalnog okruženja jer nije moguće raditi
sa datotekama na igralištu.

Molimo pokrenite ovaj program sa lokacije na kojoj je prisutan test.txt.

Kompiliranje i pokretanje:
	$ go build .
	$ cd /home/radosav/go/src/learngo/14-files/filehandling
	$ ../../learngo

Ovaj program će ispisati:

	>> Contents of file: Hello World. Welcome to file handling in Go.

Ako se ovaj program pokrene sa bilo koje druge lokacije, na primer:

cd ~/Documents/
filehandling

Odštampaće sledeću grešku:

	>> File reading error open test.txt: no such file or directory

Razlog je što je Go kompajlirajući jezik. Instalacija stvara binarni iz
izvornog koda.Binarni je nezavisan od izvornog koda i može se pokrenuti sa bilo
koje lokacije. Poštao test.txt nije pronađen na lokaciji iz koga se pokreće
binarni kod, program se žali da ne može da nađe navedenu test.txt datoteku.

Postoje tri načina za rešavanje ovog problema,

1. Korišćenje absolutnog puta datoteke
2. Prosleđenje lokacije datoteke kroz komandnu liniju
3. Pakovanju tekstualne datoteke zajedno sa binarnim kodom!!!

Hajde da razgovaramo o njima jedan po jedan.

1. Korišćenje absolutnog puta datoteke
Najjednostavniji način da rešite ovaj problem je da prođete sa absolutnim pute
datoteka u funkciju "Open". Modifikovao sam gornji program i promenio put u
absoluti. Promenite vaš put do lokacije testa.txt.

package filehandling1

import (
	"fmt"
	"os"
)

func ReadAll() {
	contents, err := os.ReadFile("/home/radosav/go/src/learngo/14-files/
		filehandling/test.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Contents of file:", string(contents))
}

Sada se program može pokrenuti sa bilo koje lokacije i štampaće sadržaj
test.txt.
*/

func ReadAllOfNonLocalFileWithAbsPath() {

	fmt.Println("\n --- ReadAllOfNonLocalFileWithAbsPath ---")
	fh1.ReadAll()
}

/*
Program će odštampati sadržaj test.txt

	>> Contents of file: Hello World. Welcome to file handling in Go.

Čini se da je to jednostavan način, ali dolazi sa zamkom da datoteka treba da
bude smeštena na putu navedenoj u programu, ako je na drugom mestu, ova metoda
neće uspeti.

2. Prosleđenje puta datoteke kroz komandnu liniju
Drugi način da reši ovaj problem je da se put datoteka prenese kao argument
komandne linije. Pomoću paketa "flag" možemo dobiti put datoteke kao ulazni
argument Iz komandne linije i zatim pročitati njegov sadržaj.

Prvo da razumemo kako funkcioniše paket "flag". Paket "flag" ima String funkciju.
Ova funkcija prihvata 3 argumenta.Prvi je ime zastavice, drugi je podrazumevana
vrednost i treći je kratak opis zastavice.

Napišimo mali program da bismo pročitali ime datoteke iz komandne linije.
Zameniti sadržaj datoteke datoteke .go sa sledećim,

package filehandling2

import (
	"flag"
	"fmt"
)

func ReadAll() {
	fptr := flag.String("fpath", "test.txt", "file path to read from")
	flag.Parse()
	fmt.Println("value of fpath is", *fptr)
}

Gornji program, stvara string zastavu pod nazivom "fpath" sa zadanim test.txt
vrednosti, i opis "file path to read from". Ova funkcija vraća adresu
promenljive stringa koja čuva vrednost zastave.

"flag.Parse()" bi trebalo pozvati pre pristupa bilo kojoj zastavi.

Štampamo vrednost zastave.

Kada se ovaj program pokrene:

	learngo -fpath=/path-of-file/test.txt

mi prosleđujemo /path-of-file/test.txt kao vrednost zastave "fpath".

Ovaj program štampa

	>> value of fpath is /path-of-file/test.txt

Ako je program pokrenut samo sa leargo bez prosledjenje vrednosti "fpath", on
će štampatii:

	>> value of fpath is test.txt

pošto je to podrazumevana vrednost zastva "fpath".

Zastava takođe pruža lepo formatiranu proizvodnju različitih argumenata koji
su dostupni.Ovo se može prikazati pokretanjem

	filehandling --help

Ova naredba će odštampati sledeći izlaz:

	Usage of learngo:
	  -fpath string
	    	file path to read from (default "test.txt")

Lepo zar ne?

Sada kada znamo kako pročitati put datoteka iz komandne linije, idemo
napred i završite naš program za čitanje datoteka.

package filehandling2

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fptr := flag.String("fpath", "test.txt", "file path to read from")
	flag.Parse()
	contents, err := os.ReadFile(*fptr)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Contents of file:", string(contents))
}
*/

func ReadAllOfNonLocalFileWithCmdLine() {

	fmt.Println("\n --- ReadAllOfNonLocalFileWithCmdLine ---")

	fh2.ReadAllFlag()
}

/*
Gornji program čita sadržaj Put datoteke preneseno iz komande
Linija.Pokrenite ovaj program pomoću komande

	filehandling -fpath=/path-of-file/test.txt

Molimo zamenite / Put-datoteku / sa apsolutnim putem test.txt. Na primer,
u mom slučaju, potrčao sam komandu

	filehandling --fpath=/Users/naveen/Documents/filehandling/test.txt

I program će ispisati:

	>> Contents of file: Hello World. Welcome to file handling in Go.

3. Pakovanje tekstualne datoteke zajedno sa binarnim kodom
Gornja opcija dobijanja puta datoteke iz komandne linije je dobra, ali postoji
još bolji način da se reši ovaj problem. Zar ne bi bilo fenomenalno ako možemo
da nosimo tekstualnu datoteku zajedno sa našim binarnim kodom? To ćemo dalje
raditi.

Paket "embed" iz standardne biblioteke pomoći će nam da to postignemo. Nakon
uvoza "embed" paketa,

	//go: embed test.txt

direktiva se može koristiti za čitanje sadržaja datoteke. Program će učiniti da
bolje razumemo stvari.

U datoteku filehandling3/filehandling3 smestimo sledeći sadržaj:

package filehandling3

import (
	_ "embed"
	"fmt"
)

//go:embed test.txt
var contents []byte

func ReadAllEmbed() {
	fmt.Println("Contents of file:", string(contents))
}

Uvozimo "embed" paket sa prefiksom _.Razlog je zato što se "embed" ne koristi
izričito u kodu, već //go: embed komentar. Budući da moramo da uvezemo paket
bez ikakvih eksplicitnih upotreba, prefiksiramo ga sa podvlakom da se kompajler
učini srećnim. Ako ne, kompajler će se žaliti navodeći da se paket nigde ne
koristi.

// go: embed test.txt govori kompajleru da pročita sadržaj test.txt i dodeli ga
promenljivoj nakon tog komentara. U našem slučaju "contents" promenljiva će
zadržati sadržaj datoteke.
*/

func ReadAllOfFileEmbed() {

	fmt.Println("\n --- ReadAllOfFileEmbed ---")
	fh3.ReadAllEmbed()
}

/*
Pokrenite program i program će štampati

	>> Contents of file: Hello World. Welcome to file handling in Go.

Sada je datoteka ubačena zajedno sa binarnom kodom  i dostupna je Gou bez obzira
na to gde se izvršava. Na primer, pokušajte da pokrenete program iz direktorija
u kojem test.txt ne boravi.

Kada se pokrene, ispisaće sadržaj datoteke.

Imajte na umu da dodela promenljive na koju treba dodeliti sadržaj datoteke mora
biti na nivou paketa. Lokalne promenljive neće raditi.

Program će propasti sa sledećom greškom.

	>> ./filehandling.go:9:4: go:embed cannot apply to var inside func

Čitanje datoteke u malim delovima
---------------------------------
U poslednjem odeljku naučili smo kako da učitamo celu datoteku u memoriju. Kada
je veličina datoteke izuzetno velika, nema smisla pročitati celu datoteku u
memoriju, posebno ako radite sa malo RAM-a.Optimalniji način je pročitati
datoteku u malim delovima. To se može učiniti uz pomoć "Bufio" paketa.

Hajde da napišemo program koji čita našu test.txt datoteku u komadima od 3 bajta.
*/

func ReadChunkByChunk() {

	fmt.Println("\n --- ReadChunkByChunk ---")

	// fptr := flag.String("fpath", "test.txt", "file path to read from")
	// flag.Parse()
	// f, err := os.Open(*fptr) // Open file

	filepath := "./14-files/test.txt" // root /home/radosav/go/src/learngo
	f, err := os.Open(filepath)       // Open file
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = f.Close(); err != nil { // Defer Close file
			log.Fatal(err)
		}
	}()

	r := bufio.NewReader(f) // Get newreader for file
	b := make([]byte, 3)    // Buffer 3 - bytes

	for {
		n, err := r.Read(b) // Read from file to buff, n is number reading chars
		if err == io.EOF {  // End of file
			fmt.Println("finished reading file")
			break
		}

		if err != nil {
			fmt.Printf("Error %s reading file", err)
			break
		}
		fmt.Println(string(b[0:n])) // Print all from buff b
	}
}

/*
// U gornjem programu, otvaramo datoteku pomoću puta iz zastave komandne linije.
U gornjem programu, otvaramo datoteku pomoću relativne putanje u odnosu na root
direktorijum modula learngo. Zatim odlažemo zatvaranje  datoteke.

U gornjem programu kreiramo novi buferisani čitač. U sledećoj liniji kreiramo
isečak bajtova dužine i kapaciteta 3 u koji će se čitati bajtovi datoteke.

Metoda Read čita do 3 bajta i vraća broj pročitanih bajtova.Čuvamo bajtove
vraćene u bufferu b. Isečak se čita 0 do N-1 indeksa, tj. do broja bajtova koji
su vraćeni metodom za čitanje i tada ih štampano.

Jednom kada se postigne kraj datoteke, Read će vratiti EOF grešku.Proveravamo
ovu grešku i prekidamo.

Ako pokrećemo program iznad, sledeći će biti izlaz

	>> Hel
	>> lo
	>> Wor
	>> ld.
	>>  We
	>> lco
	>> me
	>> to
	>> fil
	>> e h
	>> and
	>> lin
	>> g i
	>> n G
	>> o.
	>> finished reading file

Čitanje datoteke datoteke po liniji
-----------------------------------
U odeljku ćemo razgovarati o tome kako da pročitam datotečnu liniju linijom koristeći Go.To može učiniti pomoću Bufio paketa.
Molimo zamenite sadržaj u test.tkt sa sledećim

Pozdrav svijet.Dobrodošli u rukovanje datotekom u Idi.
Ovo je druga linija datoteke.
We have reached the end of the file.

Sledeći su koraci uključeni u čitanje datoteke datoteke po liniji.
    1. Open the file
    2. Create a new scanner from the file
    3. Scan the file and read it line by line.

Replace the contents of filehandling.go with the following
*/

func ReadLineByLine() {

	fmt.Println("\n --- ReadLineByLine ---")

	// fptr := flag.String("fpath", "test.txt", "file path to read from")
	// flag.Parse()
	// f, err := os.Open(*fptr)

	filepath := "./14-files/test.txt" // root /home/radosav/go/src/learngo
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	s := bufio.NewScanner(f)
	for s.Scan() {
		fmt.Println(s.Text())
	}

	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}
}

/*
U gornjem programu otvorimo datoteku pomoću staze koja je prođena iz zastave
komandne linije.Zatim kreiramo novi skener pomoću datoteke. Metoda "Scan()" čita
sledeći red datoteke i string koji je pročitala biće dostupno kroz metodu
"Text()".

Kada Scan() vrati false, metoda Err() će vratiti bilo koju grešku koja se
dogodila tokom skeniranja.Ako je greška kraj datoteke, greška će se vratiti nil.
Ako pokrenemo program sadržaj datoteke će biti štampana linija po liniji kao
što je prikazano u nastavku.

	>> Hello World. Welcome to file handling in Go.
	>> This is the second line of the file.
	>> We have reached the end of the file.
*/

func ReadFiles() {

	fmt.Println("\n --- Read files ---")

	ReadAllOfLocalFile()
	ReadAllOfNonLocalFileWithAbsPath()
	ReadAllOfNonLocalFileWithCmdLine()
	ReadAllOfFileEmbed()
	ReadChunkByChunk()
	ReadLineByLine()
}
