/*
Prilagođene greške
==================

Kreiranje prilagođenih grešaka pomoću funkcije New
--------------------------------------------------
Najjednostavniji način za kreiranje prilagođene greške je korišćenje funkcije
"New" iz paketa "errors".

Pre nego što upotrebimo funkciju "New" za kreiranje prilagođene greške, hajde da
razumemo kako je ona implementirana. Implementacija funkcije "New" u paketu
"errors" je data u nastavku:

	package errors

	// New returns an error that formats as the given text.
	// Each call to New returns a distinct error value even if the text is
	// identical.
	func New(text string) error {
	        return &errorString{text}
	}

	// errorString is a trivial implementation of error.
	type errorString struct {
	        s string
	}

	func (e *errorString) Error() string {
	        return e.s
	}

Implementacija je prilično jednostavna. "errorString" je tipa strukture sa
jednim string poljem "s". Metoda "Error() string" interfejsa "error" je
implementirana korišćenjem "errorString" pokazivača prijemnika.

Funkcija New uzima string parametar, kreira vrednost tipa errorStringkoristeći
taj parametar i vraća njegovu adresu. Tako se kreira i vraća nova greška.

Sada kada znamo kako Newfunkcija radi, hajde da je koristimo u našem programu
da bismo kreirali prilagođenu grešku.

Napravićemo jednostavan program koji izračunava površinu kruga i vraća grešku
ako je poluprečnik negativan.
*/

package de

import (
	"errors"
	"fmt"
	"math"
)

func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		return 0, errors.New("area calculation failed, radius is less than zero")
	}
	return math.Pi * radius * radius, nil
}

func errCustom() {

	fmt.Println("\n --- errCustom ---")

	radius := -20.0
	area, err := circleArea(radius)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Area of circle %0.2f", area)
}

/*
U gornjem programu proveravamo da li je radijus manji od nule. Ako jeste,
vraćamo nulu za površinu zajedno sa odgovarajućom porukom o grešci. Ako je
radijus veći od 0, onda se površina izračunava i vraća nil kao greška.

U glavnoj funkciji proveravamo da li greška nije nil. Ako nije nil, ispisujemo
grešku i vraćamo se, u suprotnom se ispisuje površina kruga.

U ovom programu radijus je manji od nule i stoga će se ispisati,

	>> Area calculation failed, radius is less than zero

Dodavanje dodatnih informacija grešci pomoću funkcije Errorf
------------------------------------------------------------
Gore navedeni program dobro funkcioniše, ali zar ne bi bilo lepo kada bismo
ispisali stvarni radijus koji je izazvao grešku. Tu nam dobro dođe funkcija
"Errorf" iz "fmt" paketa. Ova funkcija formatira grešku prema specifikatoru
formata i vraća string kao vrednost koja zadovoljava "error" interfejs.

Hajde da iskoristimo Errorffunkciju i poboljšamo program.
*/

func circleArea2(radius float64) (float64, error) {

	if radius < 0 {
		return 0, fmt.Errorf("area calculation failed, radius %0.2f is less than zero", radius)
	}
	return math.Pi * radius * radius, nil
}

func errCustomErrorf() {

	fmt.Println("\n --- errCustomErrorf ---")
	radius := -20.0

	area, err := circleArea2(radius)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Area of circle %0.2f", area)
}

/*
U gornjem programu, Errorf se koristi za ispis stvarnog radijusa koji je
izazvao grešku. Pokretanje ovog programa će ispisati,

	>> Area calculation failed, radius -20.00 is less than zero

Pružanje više informacija o grešci korišćenjem tipa strukture i polja
---------------------------------------------------------------------
Takođe je moguće koristiti strukturne tipove koji implementiraju interfejs error
kao greške. Ovo nam daje veću fleksibilnost u rukovanju greškama. U našem
prethodnom primeru, ako želimo da pristupimo radijusu koji je izazvao grešku,
jedini način sada je da analiziramo opis greške

	>> Area calculation failed, radius -20.00 is less than zero.

Ovo nije pravilan način da se to uradi jer ako se opis promeni, naš kod će se
pokvariti.

Koristićemo strategiju koju prati standardna biblioteka objašnjena u prethodnom
tutorijalu u odeljku "Konvertovanje greške u osnovni tip i preuzimanje dodatnih
informacija iz strukturnih polja" i koristiti strukturna polja da bismo
omogućili pristup radijusu koji je izazvao grešku. Napravićemo strukturni tip
koji implementira interfejs greške i koristiti njegova polja da bismo pružili
više informacija o grešci.

Prvi korak bi bio kreiranje strukturnog tipa koji predstavlja grešku.
Konvencija imenovanja tipova grešaka je da se ime završava sa tekstom Error.
Dakle, hajde da nazovemo naš strukturni tip kao "areaError":

type areaError struct {
	err    string
	radius float64
}

Gore navedeni tip strukture ima polje "radius" koje čuva vrednost radijusa
odgovornog za grešku i "err" polje koje čuva stvarnu poruku o grešci.

Sledeći korak je implementacija interfejsa za greške.

func (e *areaError) Error() string {
    return fmt.Sprintf("radius %0.2f: %s", e.radius, e.err)
}

U gornjem isečku koda, implementiramo "Error() string" metodu interfejsa greške
koristeći pokazivački prijemnik *areaError. Ova metoda ispisuje radijus i opis
greške.

Hajde da završimo program pisanjem mainfunkcije i circleAreafunkcije.
*/

type areaError struct {
	err    string
	radius float64
}

func (e *areaError) Error() string {
	return fmt.Sprintf("radius %0.2f: %s", e.radius, e.err)
}

func circleArea3(radius float64) (float64, error) {
	if radius < 0 {
		return 0, &areaError{
			err:    "radius is negative",
			radius: radius,
		}
	}
	return math.Pi * radius * radius, nil
}

func errCustomErrorStruct() {

	fmt.Println("\n --- errCustomErrorStruct ---")
	radius := -20.0
	area, err := circleArea3(radius)
	if err != nil {
		var areaError *areaError
		if errors.As(err, &areaError) {
			fmt.Printf("Area calculation failed, radius %0.2f is less than zero\n", areaError.radius)
			return
		}
		fmt.Println(err)
		return
	}
	fmt.Printf("Area of rectangle %0.2f\n", area)
}

/*
U gornjem programu, circleArea3 izračunava površinu kruga. Ova funkcija prvo
proverava da li je prečnik manji od nule, ako jeste, kreira vrednost tipa
areaError koristeći poluprečnik odgovoran za grešku i odgovarajuću poruku o
grešci, a zatim vraća adresu istog sa 0 površinom. Tako smo pružili više
informacija o grešci, u ovom slučaju o poluprečniku koji je izazvao grešku,
koristeći polja prilagođene strukture greške.

Ako radijus nije negativan, ova funkcija izračunava i vraća površinu zajedno sa
nil greškom.

U glavnoj funkciji, pokušavamo da pronađemo površinu kruga sa poluprečnikom -20.
Pošto je poluprečnik manji od nule, biće vraćena greška.

Proveravamo da li je greška nil i ako nije pokušavamo da je konvertujemo u tip
*areaError. Ako je greška tipa *areaError, dobijamo poluprečnik koji je izazvao
grešku koristeći "areaError.radius", ispisujemo prilagođenu poruku o grešci i
vraćamo se iz programa.

Ako greška nije tipa "*areaError", jednostavno ispisujemo grešku i vraćamo. Ako
nema greške, površina će biti ispisana.

Program će štampati,

	>> Area calculation failed, radius -20.00 is less than zero

Sada hajde da koristimo drugu strategiju opisanu u prethodnom tutorijalu i
koristimo metode za prilagođene tipove grešaka kako bismo pružili više
informacija o grešci.

Pružanje više informacija o grešci korišćenjem metoda na tipovima struktura
-------------------------------------------------------------------------------
U ovom odeljku ćemo napisati program koji izračunava površinu pravougaonika.
Ovaj program će ispisati grešku ako je dužina ili širina manja od nule.

Prvi korak bi bio kreiranje strukture koja predstavlja grešku.

type areaError2 struct {
	err    string //error description
	length float64 //length which caused the error
	width  float64 //width which caused the error
}

Gore navedeni tip strukture greške sadrži polje za opis greške zajedno sa
dužinom i širinom koje su izazvale grešku.

Sada kada imamo tip greške, hajde da implementiramo interfejs za greške i
dodamo nekoliko metoda na tip greške kako bismo pružili više informacija o
grešci.

func (e *areaError) Error() string {
	return e.err
}

func (e *areaError) lengthNegative() bool {
	return e.length < 0
}

func (e *areaError) widthNegative() bool {
	return e.width < 0
}

U gornjem isečku koda, vraćamo opis greške iz "Error() string" metode.
"lengthNegative() bool" metod vraća vrednost "true" kada je dužina manja od
nule, a "widthNegative() bool" metod vraća vrednost "true" kada je širina manja
od nule. Ove dve metode pružaju više informacija o grešci, u ovom slučaju govore
da li je izračunavanje površine neuspešno zbog negativne dužine ili negativne
širine. Stoga smo koristili metode za greške tipa strukture da bismo pružili
više informacija o grešci.

Sledeći korak je pisanje funkcije za izračunavanje površine.

func rectArea(length, width float64) (float64, error) {
	err := ""
	if length < 0 {
		err += "length is less than zero"
	}
	if width < 0 {
		if err == "" {
			err = "width is less than zero"
		} else {
			err += ", width is less than zero"
		}
	}
	if err != "" {
		return 0, &areaError2{
			err:    err,
			length: length,
			width:  width,
		}
	}
	return length * width, nil
}

Gore navedena funkcija "rectArea proverava" da li je dužina ili širina manja od
nule, ako jeste, vraća grešku tipa "*areaError", u suprotnom vraća površinu
pravougaonika sa nil greškom

Završićemo ovaj program kreiranjem glavne funkcije.

func errCustomErrorStructMethod() {
	length, width := -5.0, -9.0
	area, err := rectArea(length, width)
	if err != nil {
		var areaError *areaError2
		if errors.As(err, &areaError) {
			if areaError.lengthNegative() {
				fmt.Printf("error: length %0.2f is less than zero\n", areaError.length)

			}
			if areaError.widthNegative() {
				fmt.Printf("error: width %0.2f is less than zero\n", areaError.width)

			}
			return
		}
		fmt.Println(err)
		return
	}
	fmt.Println("area of rect", area)
}

U glavnoj funkciji proveravamo da li greška nije nil. Ako nije nil, pokušavamo
da je konvertujemo u tip "*areaError". Zatim, koristeći metode
"lengthNegative()" i "widthNegative()"", proveravamo da li je greška nastala
zbog činjenice da je dužina negativna ili širina negativna. Ispisujemo
odgovarajuću poruku o grešci i vraćamo se iz programa. Stoga smo koristili
metode na strukturi tipa greške da bismo pružili više informacija o grešci.

Ako nema greške, biće ispisana površina pravougaonika.

Evo kompletnog programa za vašu referencu.
*/

type areaError2 struct {
	err    string  //error description
	length float64 //length which caused the error
	width  float64 //width which caused the error
}

func (e *areaError2) Error() string {
	return e.err
}

func (e *areaError2) lengthNegative() bool {
	return e.length < 0
}

func (e *areaError2) widthNegative() bool {
	return e.width < 0
}

func rectArea(length, width float64) (float64, error) {
	err := ""
	if length < 0 {
		err += "length is less than zero"
	}
	if width < 0 {
		if err == "" {
			err = "width is less than zero"
		} else {
			err += ", width is less than zero"
		}
	}
	if err != "" {
		return 0, &areaError2{
			err:    err,
			length: length,
			width:  width,
		}
	}
	return length * width, nil
}

func errCustomErrorStructMethod() {

	fmt.Println("\n --- errCustomErrorStructMethod ---")
	length, width := -5.0, -9.0
	area, err := rectArea(length, width)
	if err != nil {
		var areaError *areaError2
		if errors.As(err, &areaError) {
			if areaError.lengthNegative() {
				fmt.Printf("error: length %0.2f is less than zero\n", areaError.length)

			}
			if areaError.widthNegative() {
				fmt.Printf("error: width %0.2f is less than zero\n", areaError.width)
			}
			return
		}
		fmt.Println(err)
		return
	}
	fmt.Println("area of rect", area)
}

/*
Ovaj program će ispisati izlaz:

	>> error: length -5.00 is less than zero
	>> error: width -9.00 is less than zero

Videli smo primere za dva od tri načina opisana u tutorijalu za rukovanje
greškama kako bismo pružili više informacija o greškama.

Treći način korišćenja direktnog poređenja je prilično jednostavan. Ostavio bih
vam to kao vežbu da shvatite kako da koristite ovu strategiju da biste pružili
više informacija o našim prilagođenim greškama.
*/

func CustomError() {
	fmt.Println("\n --- Custom Error ---")

	errCustom()
	errCustomErrorf()
	errCustomErrorStruct()
	errCustomErrorStructMethod()
}
