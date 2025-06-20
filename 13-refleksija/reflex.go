package ref

import (
	"fmt"
	"reflect"
)

/*
Refleksija
==========

Refleksija je sposobnost programa da pregleda svoje promenljive i vrednosti
tokom izvršavanja i pronađe njihov tip. Možda ne razumete šta ovo znači, ali to
je u redu.

Zašto je potrebno ispitati promenljivu i pronaći njen tip?
----------------------------------------------------------
Prvo pitanje koje svako dobije kada uči o refleksiji je zašto uopšte moramo da
ispitujemo promenljivu i pronađemo njen tip tokom izvršavanja programa kada je
svaka promenljiva u našem programu definisana od strane nas i znamo njen tip
tokom same kompajlacije. Pa, ovo je tačno većinu vremena, ali ne uvek.

Hajde da napišemo jednostavan program.
*/

func refExample() {
	fmt.Println("\n --- refExample ---")
	i := 10
	fmt.Printf("%d %T\n", i, i)
}

/*
U gornjem programu, tip promenljive i je poznat u vreme kompajliranja i mi ga
ispisujemo. Ovde nema ništa magično.

Sada da razumemo potrebu za poznavanjem tipa promenljive tokom izvršavanja.

Recimo da želimo da napišemo jednostavnu funkciju koja će prihvatiti strukturu
kao argument i kreirati SQL upit za umetanje koristeći je.

Razmotrite sledeći program,

type order struct {
	ordId      int
	customerId int
}

func ref() {
	fmt.Println("\n --- refExample ---")
	o := order{
		ordId:      1234,
		customerId: 567,
	}
	fmt.Println(o)
}

Potrebno je da napišemo funkciju koja će prihvatiti strukturu o iz gornjeg
programa kao argument i vratiti SQL upit za umetanje,

	>> insert into order values(1234, 567)

Ova funkcija je jednostavna za pisanje. Hajde da to sada uradimo.
*/

type order struct {
	ordId      int
	customerId int
}

func createQuery(o order) string {
	i := fmt.Sprintf("insert into order values(%d, %d)", o.ordId, o.customerId)
	return i
}

func refExample2() {
	fmt.Println("\n --- refExample2 ---")
	o := order{
		ordId:      1234,
		customerId: 567,
	}
	fmt.Println(createQuery(o))
}

/*
Funkcija "createQuery" kreira upit za umetanje koristeći polja "ordId" i
"customerId" strukture "o" tipa "order". Ovaj program će ispisati,

	>> insert into order values(1234, 567)

Sada hajde da podignemo naš kreator upita na viši nivo. Šta ako želimo da
generalizujemo naš kreator upita i da ga nateramo da radi na bilo kojoj
strukturi?

Primer:

type order struct {
	ordId      int
	customerId int
}

type employee struct {
	name string
	id int
	address string
	salary int
	country string
}

func createQuery(q interface{}) string {
}

func main() {

}

Naš cilj je da završimo createQuery funkciju gornjeg programa tako da prihvati
any struct kao argument i kreira upit za umetanje na osnovu polja strukture.

Na primer, ako prosledimo strukturu ispod,

	o := order {
		ordId: 1234,
		customerId: 567
	}

"createQuery" funkcija bi trebalo da vrati,

	insert into order values (1234, 567)

Slično tome, ako prođemo sa

	e := employee {
	       name: "Naveen",
	       id: 565,
	       address: "Science Park Road, Singapore",
	       salary: 90000,
	       country: "Singapore",
	   }

"createQuery" funkcija trebalo bi da se vrati,

	insert into employee values("Naveen", 565, "Science Park Road, Singapore",
		90000, "Singapore")

Pošto "createQuery" funkcija treba da radi sa bilo kojom strukturom, ona uzima
interface{} kao argument. Radi jednostavnosti, bavićemo se samo strukturama
koje sadrže polja tipa string i int ali ovo se može proširiti za bilo koji tip.

Funkcija "createQuery" bi trebalo da radi na bilo kojoj strukturi. Jedini način
da se napiše ova funkcija jeste da se ispita tip argumenta strukture koji joj
se prosleđuje tokom izvršavanja, pronađe njegova polja, a zatim kreira upit.
Tu je refleksija korisna.

U sledećim koracima tutorijala, naučićemo kako to možemo postići koristeći
"reflect" paket.

Reflect paket
-------------
Paket "reflect" implementira refleksiju tokom izvršavanja u Go-u. Paket
"reflect" pomaže u identifikaciji konkretnog tipa i vrednosti promenljive
tipa interface{}. To je upravo ono što nam je potrebno. "createQuery" funkcija
prihvata interface{} argument, a upit treba kreirati na osnovu konkretnog tipa
i vrednosti argumenta interface{}. Upravo u tome nam paket "reflect" pomaže.

Postoji nekoliko tipova i metoda u paketu reflect koje treba prvo da znamo pre
nego što napišemo naš generički program za generisanje upita. Hajde da ih
pogledamo jednu po jednu.

reflect.Type i reflect.Value
----------------------------
Konkretan tip interface{} je predstavljen sa "reflect.Type", a osnovna vrednost
je predstavljena sa "reflect.Value". Postoje dve funkcije "reflect.TypeOf()" i
"reflect.ValueOf()"" koje vraćaju "reflect.Type" i "reflect.Value" respektivno.
Ova dva tipa su osnova za kreiranje našeg generatora upita.

Napišimo jednostavan primer da bismo razumeli ova dva tipa:
*/

type order2 struct {
	ordId      int
	customerId int
}

func createQuery2(q interface{}) {
	t := reflect.TypeOf(q)
	v := reflect.ValueOf(q)

	fmt.Println("Type (reflect.TypeOf)", t)
	fmt.Println("Value (reflect.ValueOf)", v)
}

func refBasicFunc() {
	fmt.Println("\n --- Basic methods: TypeOf and ValueOf ---")

	o := order2{
		ordId:      456,
		customerId: 56,
	}
	createQuery2(o)
}

/*
U gornjem programu, funkcija "createQuery" uzima interface{} kao argument.

Funkcija "reflect.TypeOf" uzima "interface{}"" kao argument i vraća
"reflect.Type" koji sadrži konkretan tip prosleđenog argumenta "interface{}".

Slično tome, funkcija "reflect.ValueOf" uzima "interface{}"" kao argument i
vraća "reflect.Value" koji sadrži osnovnu vrednost prosleđenog argumenta
"interface{}".

Gore navedeni program štampa,

	>> Typere  ref.order2
	>> Value  {456 56}

Iz izlaza možemo videti da program ispisuje konkretan tip i vrednost interfejsa.

reflect.Kind
------------
Postoji još jedan važan tip u "reflex" paketu koji se zove "Kind". Tipovi "Kind"
i "Type" u "reflex" paketu mogu izgledati slično, ali postoji razlika koja će
biti jasna iz programa u nastavku.
*/

type order3 struct {
	ordId      int
	customerId int
}

func createQuery3(q interface{}) {
	t := reflect.TypeOf(q)
	k := t.Kind()
	fmt.Println("Type ", t)
	fmt.Println("Kind ", k)
}

func refKindType() {

	fmt.Println("\n --- refKindType methods---")

	o := order3{
		ordId:      456,
		customerId: 56,
	}
	createQuery3(o)
}

/*
Gore navedeni program prikazuje,

	>> Type  ref.order3
	>> Kind  struct

Mislim da vam je sada jasno koja je razlika između njih dvoje. Type predstavlja
stvarni tip "interface{}"", u ovom slučaju ref.order3 a Kind predstavlja
specifičnu vrstu tipa. U ovom slučaju, to je struct.

Metode NumField() i Field()
----------------------------
Metoda NumField() vraća broj polja u strukturi, a metoda Field(i int) vraća
broj reflect.Value tog i polja.
*/

type order4 struct {
	ordId      int
	customerId int
}

func createQuery4(q interface{}) {
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		v := reflect.ValueOf(q)
		fmt.Println("Number of fields", v.NumField())
		for i := 0; i < v.NumField(); i++ {
			fmt.Printf("Field:%d type:%T value:%v\n", i, v.Field(i), v.Field(i))
		}
	}
}

func refKindType2() {

	fmt.Println("\n --- refKindType2 methods ---")

	o := order4{
		ordId:      456,
		customerId: 56,
	}
	createQuery4(o)
}

/*
U gornjem programu prvo proveravamo da li je Kind refelct.Struct jer metoda
"NumField()" radi samo na strukturi. Ostatak programa je sam po sebi razumljiv.
Ovaj program ispisuje,

	>> Number of fields 2
	>> Field:0 type:reflect.Value value:456
	>> Field:1 type:reflect.Value value:56

Metode Int() i String()
-----------------------
Metode "Int()"" i "String()"" pomažu u izdvajanju reflect.Value kao "int64" i
"string" respektivno.
*/

func refIntString() {

	fmt.Println("\n --- refIntString methods ---")

	a := 56
	x := reflect.ValueOf(a).Int()
	fmt.Printf("type:%T value:%v\n", x, x)

	b := "Naveen"
	y := reflect.ValueOf(b).String()
	fmt.Printf("type:%T value:%v\n", y, y)
}

/*
U gornjem programu izdvajamo reflect.Value kao int64, a zatim izdvajamo ga kao
string. Ovaj program ispisuje,

	>> type:int64 value:56
	>> type:string value:Naveen

Kompletan program
-----------------

Sada kada imamo dovoljno znanja da završimo naš generator upita, hajde da to
uradimo.
*/

type order5 struct {
	ordId      int
	customerId int
}

type employee struct {
	name    string
	id      int
	address string
	salary  int
	country string
}

func createQuery5(q interface{}) {
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		t := reflect.TypeOf(q).Name()
		query := fmt.Sprintf("insert into %s values(", t)
		v := reflect.ValueOf(q)
		for i := 0; i < v.NumField(); i++ {
			switch v.Field(i).Kind() {
			case reflect.Int:
				if i == 0 {
					query = fmt.Sprintf("%s%d", query, v.Field(i).Int())
				} else {
					query = fmt.Sprintf("%s, %d", query, v.Field(i).Int())
				}
			case reflect.String:
				if i == 0 {
					query = fmt.Sprintf("%s\"%s\"", query, v.Field(i).String())
				} else {
					query = fmt.Sprintf("%s, \"%s\"", query, v.Field(i).String())
				}
			default:
				fmt.Println("Unsupported type")
				return
			}
		}
		query = fmt.Sprintf("%s)", query)
		fmt.Println(query)
		return

	}
	fmt.Println("unsupported type")
}

func refComplete() {

	fmt.Println("\n --- Complete propram about reflexion ---")

	o := order5{
		ordId:      456,
		customerId: 56,
	}
	createQuery5(o)

	e := employee{
		name:    "Naveen",
		id:      565,
		address: "Coimbatore",
		salary:  90000,
		country: "India",
	}
	createQuery5(e)

	i := 90
	createQuery5(i)

}

/*
Prvo proveravamo da li je prosleđeni argument tip struct. Potom dobijamo ime
strukture iz njene "reflect.Type" koristeći "Name()"" metodu. U sledećem redu
koristimo ime i počinjemo sa kreiranjem upita.

Naredba case proverava da li je trenutno polje "reflect.Int", ako je to
slučaj, izdvajamo vrednost tog polja tipa int64 koristeći Int()metodu. Naredba
if else se koristi za obradu graničnih slučajeva. Dodajte logičke podatke da
biste razumeli zašto je potrebna. Slična logika se koristi za izdvajanje
vrednosti stringa.

Takođe smo dodali provere kako bismo sprečili pad programa kada se funkciji
"createQuery" proslede nepodržani tipovi. Ostatak programa je sam po sebi
razumljiv. Preporučujem dodavanje logova na odgovarajuća mesta i proveru
njihovog izlaza kako biste bolje razumeli ovaj program.

Ovaj program štampa,

	>> insert into order values(456, 56)
	>> insert into employee values("Naveen", 565, "Coimbatore", 90000, "India")
	>> unsupported type

Ostavio bih čitaocu kao vežbu da doda imena polja u izlazni upit. Molimo vas
pokušajte da promenite program da ispisuje upit u formatu,

	>> insert into order(ordId, customerId) values(456, 56)

Da li treba koristiti refleksiju?
---------------------------------
Nakon što smo pokazali praktičnu upotrebu refleksije, sada dolazi pravo pitanje.
Da li treba koristiti refleksiju? Želeo bih da citiram poslovicu Roba Pajka o
upotrebi refleksije koja odgovara na ovo pitanje.

	>> Jasno je bolje nego pametno. Refleksija nikada nije jasna.

Refleksija je veoma moćan i napredan koncept u Gou i treba je koristiti pažljivo.
Veoma je teško napisati jasan i održiv kod koristeći refleksiju. Treba je
izbegavati gde god je to moguće i treba je koristiti samo kada je to apsolutno
neophodno.
*/

func RefFunc() {
	refExample()
	refExample2()
	refBasicFunc()
	refKindType()
	refKindType2()
	refIntString()
	refComplete()
}
