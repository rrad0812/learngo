/*
Funkcije
=========

Šta je funkcija?
-------------------
Funkcija je blok koda koji obavlja određeni zadatak.Funkcija uzima ulaz,
obavlja neke operacije na ulazu i stvara rezultate. Na primer, funkcija
može da preuzme radijus kao ulaz i izračuna područje i obim kao rezultat.

Deklaracija funkcije
--------------------
Sledeće je sintaksa za deklaraciju funkcije u Go-u

func functionname(parametername datatype) returntype {

 //function body

}

Deklaracija funkcije počinje sa ključnom rečju func-a praćena imenom funkcije.
Parametri su navedeni između ( i ) nakon čega sledi povratni tip funkcije.
Sintaksa za određivanje parametra je: ime parametra praćeno tipom. Bilo koji
broj parametara može se odrediti. Zatim postoji blok koda između { i } koji
je telo funkcije.

Parametri i tip povratka su opcionalni u funkciji. Otuda je sledeće takođe
valjana deklaracija funkcije.

func functionname() {

}

Primer funkcije
---------------
Napišimo funkciju koja uzima cenu jednog proizvoda i količinu kao ulazne
parametre i vraća ukupnu cenu množenjem ove dve vrednosti.

func calculateBill(price int, quantity int) int {

	var totalPrice = price * quantity

	return totalPrice

}

Gornja funkcija ima dva ulazne parametre "price" i "quantity", tipa int, i vraća
"totalPrice" koja je proizvod "price" i "quantity". Povratna vrednost je takođe
tipa int.

Ako su uzastopni parametri istog tipa, možemo da izbegnemo da pišemo tip svaki
put i dovoljno je da se jednom napišete na kraju.

Dakle,  price int, quantity int se može napisati kao price, quantity int.Gornja
funkcija može biti prepisana kao:

func calculateBill(price, quantity int) int {

	var totalPrice = price * quantity

	return totalPrice

}

Sad kada imamo spremnu funkciju, pozovima je negde u kodu.Sintaksa za pozivanje
funkcije je name_of_func(parameters).Gore navedena funkcija može se nazvati
pomoću sledećeg koda.

calculateBill(10, 5)

Evo kompletnog programa koji koristi gornju funkciju i štampa ukupnu cenu.
*/

package funcs

import (
	"fmt"
)

func calculateBill(price, quantity int) int {

	var totalPrice = price * quantity
	return totalPrice
}

func calculateAndPrintBill() {

	fmt.Println("\n --- calculateAndPrintBill ---")
	price, quantity := 90, 6

	totalPrice := calculateBill(price, quantity)
	fmt.Println("Total price is", totalPrice)
}

/*
Gornji program će štampati

Total price is 540

Višestruke povratne vrednosti
-----------------------------
Moguće je vratiti više vrednosti sa funkcije. Napišimo funkcijsko "rectProps"
koja iz dužine i širine strana pravougaonika računa i vraće njegovu površinu
i obim.
*/

func rectProps(length, width float64) (float64, float64) {

	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}

func rectPropsPrint() {

	fmt.Println("\n --- rectPropsPrint ---")
	area, perimeter := rectProps(10.8, 5.6)
	fmt.Printf("Area %f Perimeter %f \n", area, perimeter)
}

/*

Ako funkcija treba da vrati višestruke povratne vrednosti, tada one moraju
biti određene između ( i ).

Funkcija "rectProps(length, width float64) (float64, float64)" ima dva
parametra tipa float64 "length" i "width" i takođe vraća dve float64
vrednosti. Gornji program štampa:

	>> Area 60.480000 Perimeter 32.800000

Imenovane povratne vrednosti
----------------------------
Moguće je vratiti imenove vrednosti iz funkcije. Ako je imenovana povratna
vrednost, ona se može smatrati promenljivom kao promenljiva u prvom redu
tela funkcije.

Gornja rectProps se može prepisati koristeći imenovane povratne vrednosti:

func rectProps(length, width float64)(area, perimeter float64) {

    area = length * width

    perimeter = (length + width) * 2

    return // no explicit return value

}

"area" i "perimetar" su imenovane povratne vrednosti u gornjoj funkciji.
Imajte na umu da izjava o povratku u funkciji ne vraća izričito vrednost.
Pošto su "area" i "perimetar" navedeni u deklaraciji funkcije kao povratne
vrednosti, automatski se vraćaju iz funkcije kada se naiđe na "return".

Prazni identifikator
--------------------
_ poznat je kao prazan identifikator u Gou. Može se koristiti umesto bilo
koje vrednosti bilo kog tipa. Da vidimo koja je upotreba ovog praznog
identifikatora.

Funkcija rectProps vraća "area"  i "perimetar" pravougaonika.Šta ako nam je
potrebno samo "area" i želimo da odbacimo "perimetar".Ovde _ dolazi na svoje.

Program ispod koristi samo "area" vraćeno iz funkcije rectProps.
*/

func rectPropsPrintOnlyArea() {

	fmt.Println("\n --- rectPropsOnlyArea ---")
	area, _ := rectProps(10.8, 5.6) // perimeter is discarded
	fmt.Printf("Area %f \n", area)
}

/*
Ovde koristimo samo "area", identifikator "_" koristi se za odbacivanje
"perimetra".
*/

func Funcs() {
	calculateAndPrintBill()
	rectPropsPrint()
	rectPropsPrintOnlyArea()
}
