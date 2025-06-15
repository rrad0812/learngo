/*
Mape
====

Šta je mapa?
------------
Mapa je ugrađeni tip podataka u jeziku Go koji se koristi za čuvanje parova
ključ-vrednost. Praktična upotreba mape je čuvanje kodova valuta i
odgovarajućih naziva valuta.

	USD - United States Dollar
	EUR - Euro
	INR - India Rupee

Mapa će biti savršeno rešenje za gore navedeni slučaj upotrebe. Šifra valute
može biti ključ, a naziv valute može biti vrednost. Mape su slične rečnicima
u drugim jezicima kao što je Python.

Kako kreirati mapu?
-------------------
Mapa se može kreirati prosleđivanjem tipa podataka ključ i vrednost funkciji
make. Sledi sintaksa za kreiranje nove mape.

make(map[type of key]type of value)

Na primer:

currencyCode := make(map[string]string)

Gornja linija koda kreira mapu pod nazivom "currencyCode" koja sadrži string
ključeve i string vrednosti.
*/
package maps

import (
	"fmt"
)

func mapFuncMake() {

	fmt.Println("\n --- Map func make ---")

	currencyCode := make(map[string]string) // init, empty map
	fmt.Println("currency code =", currencyCode)
}

/*
Gornji program kreira mapu currencyCode sa string ključem i string vrednost.
Gornji program će ispisati:

	>> map[]

Pošto nismo dodali nikakve elemente na mapu, ona je prazna.

Dodavanje stavki na mapu
------------------------
Sintaksa za dodavanje novih stavki na mapu je ista kao i za nizove. Program
ispod dodaje neke kodove valuta i nazive valuta na currencyCodemapu.
*/

func mapFuncMakeInitAppend() {

	fmt.Println("\n --- Map func make init append ---")

	currencyCode := make(map[string]string)
	currencyCode["USD"] = "US Dollar"
	currencyCode["GBP"] = "Pound Sterling"
	currencyCode["EUR"] = "Euro"
	currencyCode["INR"] = "Indian Rupee"

	fmt.Println("currencyCode map is:\n", currencyCode)
}

/*
Dodali smo 4 ključa,naime USD, GBP, EUR, INR i vrednosti - njihova odgovarajuća
imena.

Gore navedeni program štampa:

	>> currencyCode map is:
	>> map[EUR:Euro GBP:Pound Sterling INR:Indian Rupee USD:US Dollar]

Kao što ste možda prepoznali iz gornjeg rezultata, redosled preuzimanja
vrednosti sa mape nije garantovano isti kao redosled kojim su elementi dodati
na mapu.

Takođe je moguće inicijalizovati mapu tokom same deklaracije.
*/

func mapFuncInit() {

	fmt.Println("\n --- Map func init ---")

	currencyCode := map[string]string{
		"USD": "US Dollar",
		"GBP": "Pound Sterling",
		"EUR": "Euro",
	}

	currencyCode["INR"] = "Indian Rupee"

	fmt.Println("currencyCode map is:\n", currencyCode)
}

/*
Gore navedeni program deklariše currencyCodemap i dodaje joj 3 elementa tokom
same deklaracije. Kasnije dodaje se još jedan element sa ključem INR. Program
ispisuje

	>> currencyCode map contents:
	>> map[EUR:Euro GBP:Pound Sterling INR:Indian Rupee USD:US Dollar]

Nije neophodno da tipovi ključeva budu samo stringova. Svi uporedivi tipovi kao
što su bulov, ceo broj, broj sa pokretnim brojem, kompleksni, stringovi itd.
takođe mogu biti ključevi. Čak i korisnički definisani tipovi kao što su
strukture mogu biti ključevi. Ako želite da saznate više o uporedivim tipovima,
 posetite https://go.dev/ref/spec#Comparison_operators.

Panika na nil mapi
------------------
Nulta vrednost mape je nil. Ako pokušate da dodate elemente na nil mapu, doći
će do panike tokom izvršavanja. Stoga, mapa mora biti inicijalizovana pre
dodavanja elemenata.

func main() {
	var currencyCode map[string]string
	currencyCode["USD"] = "US Dollar"
}

U gornjem programu, currencyCodeje je nil i mi pokušavamo da dodamo novi ključ
na nil mapu. Program će prikazati grešku.

	>> panic: assignment to entry in nil map

Preuzimanje vrednosti sa mape
-----------------------------
Sada kada smo dodali neke elemente na mapu, hajde da naučimo kako da ih
preuzmemo.

map[key] je sintaksa za preuzimanje elemenata mape.
*/

func mapFuncAccess() {

	fmt.Println("\n --- Map func access ---")

	currencyCode := map[string]string{
		"USD": "US Dollar",
		"GBP": "Pound Sterling",
		"EUR": "Euro",
	}

	currency := "USD"
	currencyName := currencyCode[currency]

	fmt.Println("Currency name for currency code", currency, "is", currencyName)
}

/*
Gore navedeni program je prilično jednostavan. Naziv valute za kod valute USD
se preuzima i ispisuje. Program ispisuje:

	>> Currency name for currency code USD is US Dollar

Šta će se desiti ako element nije prisutan? Funkcija map će vratiti nultu
vrednost tipa tog elementa. U slučaju funkcije currencyCodemap, ako pokušamo da
pristupimo stavci koja nije prisutna, vraća se nulta vrednost stringa, ""
(prazan string).
*/

func mapFuncAccesNotPresent() {

	fmt.Println("\n --- Map func access not present ---")

	currencyCode := map[string]string{
		"USD": "US Dollar",
		"GBP": "Pound Sterling",
		"EUR": "Euro",
	}

	fmt.Println("Currency name for currency code INR is", currencyCode["INR"])
}

/*
Izlaz gornjeg programa je

>> Currency name for currency code INR is

Gore navedeni program vraća prazan string kao naziv valute za INR. Neće biti
greške tokom izvršavanja kada pokušamo da preuzmemo vrednost za ključ koji nije
prisutan u mapi.

Provera da li ključ postoji
---------------------------
U prethodnom odeljku smo saznali da kada ključ nije prisutan, biće vraćena
nulta vrednost tipa. Ovo ne pomaže kada želimo da saznamo da li ključ zaista
postoji u mapi.

Na primer, želimo da znamo da li je ključ valutnog koda prisutan na currencyCode
mapi. Sledeća sintaksa se koristi da bi se utvrdilo da li je određeni ključ
prisutan na mapi.

value, ok := map[key]

ok u gornjoj liniji koda će biti true kada je ključ prisutan i vrednost za ključ
je prisutna u promenljivoj value. Ako ključ nije prisutan, ok biće false i vraća
se nulta vrednost za value.
*/

func mapFuncAccesOk() {

	fmt.Println("\n --- Map func access ok ---")

	currencyCode := map[string]string{
		"USD": "US Dollar",
		"GBP": "Pound Sterling",
		"EUR": "Euro",
	}

	cyCode := "INR"

	if currencyName, ok := currencyCode[cyCode]; ok {
		fmt.Println("Currency name for currency code", cyCode, " is", currencyName)
		return
	}
	fmt.Println("Currency name for currency code", cyCode, "not found")
}

/*
U gornjem programu, u liniji br. 14, okbiće falsepošto INRključ nije prisutan.
Stoga će program ispisati,

	>> Currency name for currency code INR not found

Iteriranje elemenata mape
-------------------------
Oblik range petlje for se koristi za iteraciju kroz sve elemente mape.
*/

func mapFuncForRange() {

	fmt.Println("\n --- Map func for range ---")

	currencyCode := map[string]string{
		"USD": "US Dollar",
		"GBP": "Pound Sterling",
		"EUR": "Euro",
	}

	for code, name := range currencyCode {
		fmt.Printf("Currency name for currency code %s is %s\n", code, name)
	}
}

/*
Gore navedeni programski ispiši:

	>> Currency Name for currency code GBP is Pound Sterling
	>> Currency Name for currency code EUR is Euro
	>> Currency Name for currency code USD is US Dollar

Jedna važna činjenica koju treba napomenuti jeste da redosled preuzimanja
vrednosti iz mape pri korišćenju for range nije garantovano isti za svako
izvršavanje programa. Takođe nije isti kao redosled kojim su elementi dodani na
mapu.

Brisanje stavki sa mape
-----------------------
Funkcija delete(map, key) je sintaksa za brisanje key iz map brisanja ne vraća
nikakvu vrednost.
*/

func mapFuncDelete() {

	fmt.Println("\n --- Map func delete ---")

	currencyCode := map[string]string{
		"USD": "US Dollar",
		"GBP": "Pound Sterling",
		"EUR": "Euro",
	}

	fmt.Println("map before deletion is", currencyCode)

	delete(currencyCode, "EUR")

	fmt.Println("map after deletion is", currencyCode)
}

/*
Gore navedeni program briše ključ EUR i štampa:

	>> map before deletion is map[EUR:Euro GBP:Pound Sterling USD:US Dollar]
	>> map after deletion is map[GBP:Pound Sterling USD:US Dollar]

Čak i ako pokušamo da obrišemo ključ koji nije prisutan na mapi, neće biti
greške tokom izvršavanja.

Mapa struktura
--------------
Do sada smo čuvali samo naziv valute u mapi. Zar ne bi bilo lepo kada bismo
mogli da sačuvamo i simbol valute? To se može postići korišćenjem mape strukture.
Valuta se može predstaviti kao struktura koja sadrži polja naziv valute i simbol
valute. Vrednost ove strukture može se sačuvati u mapi sa kodom valute kao
ključem. Hajde da napišemo program da bismo razumeli kako se to može uraditi.
*/

type currency struct {
	name   string
	symbol string
}

func mapFuncStructs() {

	fmt.Println("\n --- Map structs ---")
	curUSD := currency{
		name:   "US Dollar",
		symbol: "$",
	}
	curGBP := currency{
		name:   "Pound Sterling",
		symbol: "£",
	}
	curEUR := currency{
		name:   "Euro",
		symbol: "€",
	}

	currencyCode := map[string]currency{
		"USD": curUSD,
		"GBP": curGBP,
		"EUR": curEUR,
	}

	for cyCode, cyInfo := range currencyCode {
		fmt.Printf("Currency Code: %s, Name: %s, Symbol: %s\n",
			cyCode, cyInfo.name, cyInfo.symbol)
	}
}

/*
U gornjem programu, currency struktura sadrži polja name i symbol. Kreiramo tri
valute curUSD, curGBPi curEUR. Zatim, inicijalizujemo mapu sa string ključem i
vrednošću tipa currency sa tri valute koje smo kreirali.

Mapa se iterira, a detalji o valuti se štampaju u sledećem redu. Ovaj program
će štampati:

	>> Currency Code: USD, Name: US Dollar, Symbol: $
	>> Currency Code: GBP, Name: Pound Sterling, Symbol: £
	>> Currency Code: EUR, Name: Euro, Symbol: €

Dužina mape
-----------
Dužina mape se može odrediti pomoću funkcije len .
*/

func mapFuncLen() {

	fmt.Println("\n --- Map func len ---")

	currencyCode := map[string]string{
		"USD": "US Dollar",
		"GBP": "Pound Sterling",
		"EUR": "Euro",
	}
	fmt.Println("length is", len(currencyCode))
}

/*
len(currencyCode) u gornjem programu vraća dužinu mape. Gornji program ispisuje,

	>> length is 3

Mape su referentni tipovi
-------------------------
Slično kao i isečci, mape su referentni tipovi. Kada se map dodeli novoj
promenljivoj, obe ukazuju na istu osnovnu strukturu podataka. Stoga će se
promene napravljene u jednoj odraziti na drugu.
*/

func funcMapRef() {

	fmt.Println("\n --- Map func ref ---")

	employeeSalary := map[string]int{
		"steve": 12000,
		"jamie": 15000,
		"mike":  9000,
	}
	fmt.Println("Original employee salary", employeeSalary)
	modified := employeeSalary
	modified["mike"] = 18000
	fmt.Println("Employee salary changed", employeeSalary)

}

/*
U redu br. 14 gornjeg programa, employeeSalaryje dodeljeno modified. U sledećem
redu, plata mikese menja u 18000na modifiedmapi. Majkova plata će sada takođe
biti 18000prisutna employeeSalary. Program izbacuje,

	>> Original employee salary map[jamie:15000 mike:9000 steve:12000]
	>> Employee salary changed map[jamie:15000 mike:18000 steve:12000]

Slično je i slučaj kada se mape prosleđuju kao parametri funkcijama. Kada se
napravi bilo kakva promena na mapi unutar funkcije, ona će biti vidljiva i
pozivaocu.

Poređenjee mape
---------------
Mape se ne mogu upoređivati pomoću == operatora . == se može koristiti samo za
proveru da li je mapa nil.
*/

func mapFuncEqu() {

	fmt.Println("\n --- Map func equ ---")

	map1 := map[string]int{
		"one": 1,
		"two": 2,
	}

	map2 := map[string]int{
		"two": 2,
		"one": 1,
	}

	fmt.Println("map1 is", map1)
	fmt.Println("map2 is", map2)

	// if map1 == map2 { // This is syntax error
	// }

	for i1 := range map1 {
		if map1[i1] != map2[i1] {
			fmt.Println("Maps map1 and map2 is not equ")
			return
		}
	}
	fmt.Println("Maps map1 and map2 is equ")
}

/*
Gore navedeni program se neće kompajlirati uz grešku:

	>> invalid operation: map1 == map2 (map can only be compared to nil)

Jedan od načina da se proveri da li su dve mape jednake jeste da se uporede
pojedinačni elementi svake od njih jedan po jedan. Drugi način je korišćenje
refleksije. Preporučio bih vam da napišete program za ovo i da ga naterate da
radi :).
*/

func MapFuncs() {

	fmt.Println("\n --- Map funcs ---")

	mapFuncMake()
	mapFuncMakeInitAppend()
	mapFuncInit()
	mapFuncAccess()
	mapFuncAccesNotPresent()
	mapFuncAccesOk()
	mapFuncForRange()
	mapFuncDelete()
	mapFuncStructs()
	mapFuncLen()
	funcMapRef()
	mapFuncEqu()
}
