/*
Greške
======

Greške ukazuju na bilo koje abnormalno stanje koje se javlja u programu. Recimo
da pokušavamo da otvorimo datoteku, a datoteka ne postoji u sistemu datoteka.
Ovo je abnormalno stanje i predstavlja se kao greška.

Greške u Gou su obične vrednosti. Baš kao i svaki drugi ugrađeni tip kao što su
int, float64, ... vrednosti grešaka mogu se čuvati u promenljivim, prosleđivati
kao parametri funkcijama, vraćati iz funkcija i tako dalje.

Greške su predstavljene korišćenjem ugrađenog "error" tipa. Više o ovom "error"
tipu ćemo saznati kasnije u ovom tutorijalu.

Počnimo odmah sa primerom programa koji pokušava da otvori datoteku koja ne
postoji.
*/

package de

import (
	"errors"
	"fmt"
	"os"
)

func errFileNotFound() {

	fmt.Println("\n --- errFileNotFound ---")

	f, err := os.Open("/test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(f.Name(), "opened successfully")
}

/*
Pokušavamo da otvorimo datoteku na putanji /test.txt (koja očigledno neće
postojati na igralištu). Funkcija "os.Open" ima sledeći potpis,

	>> func Open(name string) (*File, error)

Ako je datoteka uspešno otvorena, funkcija Open će vratiti handler datoteke i
error će biti nil. Ako dođe do greške prilikom otvaranja datoteke, vratiće se
error koja nije nil.

Ako funkcija ili metod vrati grešku, onda po konvenciji to mora biti poslednja
vrednost koju je funkcija vratila. Stoga "Open" funkcija vraća "error" kao
poslednja vrednost.

Idiomatski način obrade grešaka u programskom jeziku Go je upoređivanje vraćene
greške sa nil. Vrednost nil ukazuje na to da nije došlo do greške, a vrednost
koja nije nil ukazuje na prisustvo greške. U našem slučaju, proveravamo da li
se u error nalazi nil. Ako nije nil, jednostavno ispisujemo error i vraćamo
se iz glavne funkcije.

Pokretanje ovog programa će ispisati

	>> open /test.txt: No such file or directory

Dobijamo grešku koja navodi da datoteka ne postoji.

Reprezentacija tipa greške
--------------------------
Hajde da malo dublje istražimo i vidimo kako je definisan ugrađeni "error" tip.
"error" je tip interfejsa sa sledećom definicijom,

type error interface {
    Error() string
}

Sadrži jednu metodu sa potpisom "Error() string". Bilo koji tip koji
implementira ovaj interfejs može se koristiti kao error. Ova metoda pruža opis
greške.

Prilikom ispisivanja greške, fmt.Println funkcija interno poziva Error() string
metodu da bi dobila opis greške. Ovako je opis greške ispisan gornjeg primera
programa.

Različiti načini za izvlačenje više informacija o grešci
--------------------------------------------------------

Sada kada znamo da je "error" tip interfejsa, hajde da vidimo kako možemo izvući
više informacija o grešci.

U primeru koji smo videli iznad, upravo smo ispisali opis greške. Šta ako želimo
stvarnu putanju do datoteke koja je izazvala grešku? Jedan od mogućih načina da
se ovo dobije jeste da se analizira string greške. Ovo je bio izlaz našeg
programa,

open /test.txt: No such file or directory

Možemo analizirati ovu poruku o grešci i dobiti putanju do datoteke "/test.txt"
koja je izazvala grešku, ali ovo je prljav način. Opis greške se može promeniti
u bilo kom trenutku u novijim verzijama Go-a i naš kod će se pokvariti.

Postoji i bolji način da se više informacija o grešci:

1. Konvertovanje greške u osnovni tip i preuzimanje dodatnih informacija iz
   strukturnih polja

   Ako pažljivo pročitate dokumentaciju funkcije Open, možete videti da ona
   vraća grešku tipa *PathError. PathError, što je struktura, a njena
   implementacija u standardnoj biblioteci je sledeća:

   	type PathError struct {
   		Op   string
    	Path string
    	Err  error
	}

	func (e *PathError) Error() string {
		return e.Op + " " + e.Path + ": " + e.Err.Error()
	}

Ukoliko ste zainteresovani da znate gde se nalazi gore navedeni izvorni kod,
možete ga pronaći ovde
https://cs.opensource.google/go/go/+/refs/tags/go1.19:src/io/fs/fs.go;l=250.

Iz gornjeg koda možete razumeti da *PathError implementira error interface
deklarisanjem Error() string metode. Ova metoda spaja operaciju, putanju i
stvarnu grešku i vraća je. Tako smo dobili poruku o grešci,

	>> open /test.txt: No such file or directory

Polje Paths trukture PathError sadrži putanju datoteke koja je izazvala grešku.
Možemo koristiti funkciju As iz paketa errors da konvertujemo grešku u njen
osnovni tip.

Opis funkcije As govori o lancu grešaka. Molimo vas da ga za sada ignorišete.
Razumećemo kako lanac grešaka i prelamanje funkcionišu u posebnom tutorijalu.
Jednostavan opis As je da pokušava da konvertuje grešku u tip greške i vraća
vrednost "true" ili  "false", što ukazuje da li je konverzija uspešna ili ne.

Program će stvari razjasniti. Hajde da izmenimo program koji smo gore napisali
i ispišemo putanju koristeći "As" funkciju.
*/

func errAs() {
	fmt.Println("\n --- Error As ---")

	f, err := os.Open("test.txt")
	if err != nil {
		var pErr *os.PathError
		if errors.As(err, &pErr) {
			fmt.Println("Failed to open file at path", pErr.Path)
			return
		}
		fmt.Println("Generic error", err)
		return
	}
	fmt.Println(f.Name(), "opened successfully")
}

/*
Trči na igralištu

U gornjem programu, prvo proveravamo da li je greška nilu liniji br. 11, a zatim koristimo Asfunkciju u liniji br. 13 da je konvertujemo erru *os.PathError. Ako je konverzija uspešna, Asvratiće se true. Zatim ispisujemo putanju koristeći pErr.Pathu liniji br. 14.

Ako se pitate zašto pErrje pokazivač, razlog je taj što je interfejs za greške implementiran pokazivačem PathErrori stoga pErrje pokazivač. Donji kod pokazuje da *PathErrorimplementira interfejs za greške.

func (e *PathError) Error() string { return e.Op + " " + e.Path + ": " + e.Err.Error() }

Funkcija Aszahteva da drugi argument bude pokazivač na tip koji implementira grešku. Stoga prosleđujemo &perr.

Ovaj program izlazi,

Failed to open file at path test.txt

U slučaju da osnovna greška nije tipa *os.PathError, kontrola će doći do linije br. 17 i biće ispisana generička poruka o grešci.

Odlično 😃. Uspešno smo koristili Asfunkciju da dobijemo putanju do datoteke iz greške.
2. Prikupljanje više informacija korišćenjem metoda

Drugi način da se dobije više informacija o grešci jeste da se otkrije osnovni tip i dobije više informacija pozivanjem metoda na tipu strukture .

Hajde da ovo bolje razumemo pomoću jednog primera.

Tip strukture DNSError u standardnoj biblioteci je definisan na sledeći način,

type DNSError struct {
    ...
}

func (e *DNSError) Error() string {
    ...
}
func (e *DNSError) Timeout() bool {
    ...
}
func (e *DNSError) Temporary() bool {
    ...
}

Struktura DNSErrorima dve metode Timeout() booli Temporary() boolkoje vraćaju bulovsku vrednost koja pokazuje da li je greška nastala zbog isteka vremena ili je privremena.

Hajde da napišemo program koji konvertuje grešku u *DNSErrortip i poziva gore pomenute metode da bismo utvrdili da li je greška privremena ili je nastala zbog isteka vremena.

package main

import (
	"errors"
	"fmt"
	"net"
)

func main() {
	addr, err := net.LookupHost("golangbot123.com")
	if err != nil {
		var dnsErr *net.DNSError
		if errors.As(err, &dnsErr) {
			if dnsErr.Timeout() {
				fmt.Println("operation timed out")
				return
			}
			if dnsErr.Temporary() {
				fmt.Println("temporary error")
				return
			}
			fmt.Println("Generic DNS error", err)
			return
		}
		fmt.Println("Generic error", err)
		return
	}
	fmt.Println(addr)
}

Napomena: DNS pretrage ne rade u Playground-u. Molimo vas da pokrenete ovaj program na vašem lokalnom računaru.

U gornjem programu, u redu br. 9, pokušavamo da dobijemo IP adresu nevažećeg imena domena golangbot123.com. U redu br. 13 dobijamo osnovnu vrednost greške korišćenjem Asfunkcije i konvertovanjem u *net.DNSError. Zatim proveravamo da li je greška nastala zbog isteka vremena ili je privremena u redovima br. 14 i 18, respektivno.

U našem slučaju, greška nije ni privremena niti je nastala zbog isteka vremena i stoga će program ispisati,

Generic DNS error lookup golangbot123.com: no such host

Ako je greška bila privremena ili je nastala zbog isteka vremena, onda bi se odgovarajuća if naredba izvršila i možemo je obraditi na odgovarajući način.
3. Direktno poređenje

Treći način da se dobije više detalja o grešci je direktno poređenje sa promenljivom tipa error. Hajde da ovo razumemo pomoću primera.

Funkcija Glob paketa filepathse koristi za vraćanje imena svih datoteka koje odgovaraju šablonu. Ova funkcija vraća grešku ErrBadPatternkada je šablon pogrešno oblikovan.

ErrBadPattern je definisan u filepathpaketu kao globalna promenljiva.

1var ErrBadPattern = errors.New("syntax error in pattern")

errors.New() se koristi za kreiranje nove greške. O tome ćemo detaljno razgovarati u sledećem tutorijalu .

Funkcija Glob vraća grešku ErrBadPattern kada je šablon neispravan.

Hajde da napišemo mali program za proveru ove greške.

package main

import (
	"errors"
	"fmt"
	"path/filepath"
)

func main() {
	files, err := filepath.Glob("[")
	if err != nil {
		if errors.Is(err, filepath.ErrBadPattern) {
			fmt.Println("Bad pattern error:", err)
			return
		}
		fmt.Println("Generic error:", err)
		return
	}
	fmt.Println("matched files", files)
}

U gornjem programu tražimo datoteke šablona [koji je neispravan šablon. Proveravamo da li greška nije jednaka nil. Da bismo dobili više informacija o grešci, direktno je upoređujemo sa greškom filepath.ErrBadPatternu liniji br. 11 koristeći funkciju Is . Slično kao As, Isfunkcija radi na lancu grešaka. Više o ovome ćemo saznati u našem sledećem tutorijalu . Za potrebe ovog tutorijala, Isfunkcija se može smatrati vraćajućom trueako su obe greške koje su joj prosleđene iste.

Vraća vrednost Is„true“ u redu br. 12 jer je greška nastala zbog pogrešno oblikovanog obrasca. Ovaj program će ispisati,

Bad pattern error: syntax error in pattern

Standardna biblioteka koristi bilo koji od gore navedenih načina da pruži više informacija o grešci. Koristićemo ove načine u sledećem tutorijalu da kreiramo sopstvene prilagođene greške .

Ne ignorišite greške

Nikada ne ignorišite grešku. Ignorisanje grešaka je poziv na probleme. Dozvolite mi da prepišem primer koji navodi imena svih datoteka koje odgovaraju obrascu ignorišući greške.

package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	files, _ := filepath.Glob("[")
	fmt.Println("matched files", files)
}

Trčanje na igralištu

Već znamo iz prethodnog primera da je šablon nevažeći. Ignorisao sam grešku koju je Globfunkcija vratila koristeći _prazan identifikator u redu br. 9. Jednostavno sam ispisao podudarne datoteke u redu br. 10. Ovaj program će ispisati,

matched files []

Pošto smo ignorisali grešku, izlaz izgleda kao da nijedna datoteka ne odgovara šablonu, ali je sam šablon zapravo pogrešno oblikovan. Zato nikada ne ignorišite greške.
*/

func ErrorFunc() {
	fmt.Println("\n --- Errors ---")

	errFileNotFound()
	errAs()

}
