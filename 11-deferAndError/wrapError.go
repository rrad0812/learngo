/*
Omotavanje grešaka
==================

Omotavanje grešaka je proces kapsuliranja jedne greške u drugu. Recimo da imamo
veb server koji pristupa bazi podataka i pokušava da preuzme zapis iz baze
podataka. Ako poziv baze podataka vrati grešku, možemo da odlučimo da li da
omotamo ovu grešku ili da pošaljemo sopstvenu prilagođenu grešku iz veb servisa.

Hajde da napišemo mali program da bismo ovo razumeli.
*/
package de

import (
	"errors"
	"fmt"
)

var errNoRows = errors.New("'no rows found'")

func getRecord() error {
	return errNoRows
}

func webService() error {
	if err := getRecord(); err != nil {
		return fmt.Errorf("Error %s when calling DB", err)
	}
	return nil
}

func errWrapp() {

	fmt.Println("\n --- errWrapping ---")

	if err := webService(); err != nil {
		fmt.Printf("Error: %s when calling webservice\n", err)
		return
	}
	fmt.Println("webservice call successful")

}

/*
U gornjem programu šaljemo string opisa greške koja se javlja prilikom poziva
"getRecord" funkcije.

Iako ovo može izgledati kao omotavanje grešaka, nije :).

Hajde da razumemo kako se greške prelamaju u sledećem odeljku.

Omotavanje greške i funkcija Is
-------------------------------
Funkcija "Is" u paketu "errors" izveštava da li se neka od grešaka u lancu
podudara sa ciljem. U našem slučaju, vraćamo "errNoRows" grešku iz "getRecord"
funkcije. Format stringa ove greške vraća se iz "webService" funkcije. Hajde da
modifikujemo glavnu funkciju unkciju ovog programa i koristimo "Is" funkciju da
proverimo da li se neka greška u lancu podudara sa "errNoRows" greškom.

func errWrappIs() {
	if err := webService(); err != nil {
		if errors.Is(err, errNoRows) {
			fmt.Printf("The searched record cannot be found. Error returned from DB is %s", err)
			return
		}
		fmt.Println("unknown error when searching record")
		return
	}
	fmt.Println("webservice call successful")
}

U gornjoj funkciji "Is" funkcija će proveriti da li bilo koja greška u lancu
grešaka "err" sadrži "errNoRows" grešku. Kod u svom trenutnom stanju neće
raditi jer if uslov gornje glavne funkcije će biti neuspešan. Da bi ovo
funkcionisalo, potrebno je da omotamo errNoRows grešku kada je vrati webService
funkcija. Jedan od načina da se to uradi je da se koristi %w specifikator
formata prilikom vraćanja greške umesto %s. Dakle, ako izmenimo red koji vraća
grešku na:

	return fmt.Errorf("Error %w when calling DB", err)

To znači da je novovraćena greška omotava preklapa originalnu grešku errNoRows
i da if uslov gornje glavne funkcije će biti uspešno ispunjen.
Kompletan program sa preklapanjem greške je dat u nastavku:
*/

var errNoRows2 = errors.New("no rows found")

func getRecord2() error {
	return errNoRows2
}

func webService2() error {
	if err := getRecord2(); err != nil {
		return fmt.Errorf("Error %w when calling DB", err)
	}
	return nil
}

func errWrappIs() {

	fmt.Println("\n --- errWrappingIS ---")

	if err := webService2(); err != nil {
		if errors.Is(err, errNoRows2) {
			fmt.Printf("The searched record cannot be found. Error returned from DB is %s\n", err)
			return
		}
		fmt.Println("unknown error when searching record")
		return
	}
	fmt.Println("webservice call successful")
}

/*
Kada se ovaj program pokrene, on će štampati.

	>> The searched record cannot be found. Error returned from DB is Error
	   no rows found when calling DB

As funkcija
-----------
Funkcja "As" u paketu "errors" će pokušati da konvertuje grešku koja je
prosleđena kao ulaz u ciljni tip greške. Uspeće ako se bilo koja od grešaka u
lancu grešaka podudara sa ciljem. Ako je uspešan, vratiće vrednost true i
postaviće cilj na prvu grešku u lancu grešaka koja se podudara.
Program će olakšati razumevanje stvari :)
*/

type DBError3 struct {
	desc string
}

func (dbError DBError3) Error() string {
	return dbError.desc
}

func getRecord3() error {
	return DBError3{
		desc: "no rows found",
	}
}

func webService3() error {
	if err := getRecord3(); err != nil {
		return fmt.Errorf("Error %w when calling DB", err)
	}
	return nil
}

func errWrappAs() {

	fmt.Println("\n --- errWrappingAs ---")

	if err := webService3(); err != nil {
		var dbError DBError3
		if errors.As(err, &dbError) {
			fmt.Printf("The searched record cannot be found. Error returned from DB is %s", dbError)
			return
		}
		fmt.Println("unknown error when searching record")
		return
	}
	fmt.Println("webservice call successful")
}

/*
U gornjem programu, izmenili smo getRecord funkciju da vrati prilagođenu grešku
tipa DBError.

U glavnoj funkciji, pokušavamo da konvertujemo grešku vraćenu pozivom
webService() funkcije u tip DBError. if izjava će biti uspešna jer smo obmotali
grešku DBError kada je funkcija vratila grešku webService().
Pokretanje ovog programa će ispisati:

	>> The searched record cannot be found.
	   Error returned from DB is no rows found.

Da li treba da omotavamo greške?

Odgovor na ovo pitanje je da zavisi. Ako grešku "omotamo", izlažemo je
pozivaocima naše biblioteke/funkcije. Generalno ne želimo da "omotamo" greške
koje sadrže interne detalje implementacije funkcije. Još jedna važna stvar koju
treba zapamtiti je da, ako vratimo "omotanu" grešku, a kasnije odlučimo da
uklonimo "omotanu" grešku, kod koji koristi našu biblioteku će početi da
otkazuje. Dakle, "omotane" greške treba smatrati delom API-ja i treba izvršiti
odgovarajuće izmene verzije ako odlučimo da izmenimo grešku koju vraćamo.
*/

func WrappError() {
	fmt.Println("\n --- Wrapping Error ---")

	errWrapp()
	errWrappIs()
	errWrappAs()
}
