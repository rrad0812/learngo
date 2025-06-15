/*
Stingovi
========

Šta je string?
--------------
U programskom jeziku Go string je isečak bajtova. Stringovi se mogu kreirati
stavljanjem skupa znakova unutar dvostrukih navodnika "".

Pogledajmo jednostavan primer koji kreira stringi štampa ga.
*/

package strings

import (
	"fmt"
	"unicode/utf8"
)

func stringElem() {

	fmt.Println("\n --- stringElem ---")

	name := "Hello World"
	name_bytes := []byte(name) // String conversion to byte slice

	fmt.Println(name)
	fmt.Println(name_bytes)
}

/*
Gore navedeni program će ispisati Hello World.

Stringovi u Go programu su kompatibilni sa Unicode standardom i kodirani su u
UTF-8 standardu.

Pristup pojedinačnim bajtovima stringa
--------------------------------------
Pošto je string isečak bajtova, moguće je pristupiti svakom bajtu stringa.
*/

func printBytes(s string) {
	fmt.Printf("Bytes: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()
}

func stringAccessBytes() {

	fmt.Println("\n --- stringAccessBytes ---")

	name := "Hello World"
	fmt.Printf("String: %s\n", name)
	printBytes(name)
}

/*
%s je specifikator formata za štampanje stringa. U gornjem programu, štampa se
ulazni string. Dalje, len(s) vraća broj bajtova u stringu i koristimo for petlju
da ispišemo te bajtove u heksadecimalnom zapisu. %x je specifikator formata za
heksadecimalni zapis. Gornji program štampa:

	>> String: Hello World
	>> Bytes: 48 65 6c 6c 6f 20 57 6f 72 6c 64

Ovo su vrednosti Hello World kodirane po Unicode UTF-8 standardu. Potrebno je
osnovno razumevanje Unicode-a i UTF-8 standarda da biste bolje razumeli
stringove. Preporučujem čitanje:
https://naveenr.net/unicode-character-set-and-utf-8-utf-16-utf-32-encoding/
da biste saznali više o Unicode-u i UTF-8 standardu.

Pristup pojedinačnim znakovima stringa
--------------------------------------
Hajde da malo izmenimo gornji program da bi ispisao karaktere stringa.
*/

func printChars(s string) {
	fmt.Printf("Characters: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i])
	}
	fmt.Println()
}

func stringAccessChars() {

	fmt.Println("\n --- stringAccessChars ---")

	name := "Hello World"
	fmt.Printf("String: %s\n", name)

	printChars(name)
	printBytes(name)
}

/*
U gornjem programu, specifikator formata %c se koristi za ispisivanje karaktera
stringa u printChars metodi. Program ispisuje:

	>> String: Hello World
	>> Characters: H e l l o   W o r l d
	>> Bytes: 48 65 6c 6c 6f 20 57 6f 72 6c 64

Iako gore navedeni program izgleda kao legitiman način za pristup pojedinačnim
karakterima stringa, on ima ozbiljnu grešku. Hajde da saznamo šta je ta greška.
*/

func stringError() {

	fmt.Println("\n --- stringError ---")

	name := "Hello World"
	fmt.Printf("String: %s\n", name)
	printChars(name)
	printBytes(name)

	fmt.Println()

	name = "Señor"
	fmt.Printf("String: %s\n", name)
	printChars(name)
	printBytes(name)
}

/*
Izlaz gornjem programu je

	>> String: Hello World
	>> Characters: H e l l o   W o r l d
	>> Bytes: 48 65 6c 6c 6f 20 57 6f 72 6c 64

	>> String: Señor
	>> Characters: S e Ã ± o r
	>> Bytes: 53 65 c3 b1 6f 72

U gornjem programu, pokušavamo da odštampamo karaktere `Señor` ali program
štampa `S e Ã ± o r`, što je pogrešno. Zašto ovaj program ne radi za "Señor",
kada radi savršeno dobro za "Hello World". Razlog je taj što je `Unicode`
kodna tačka "ñi" U+00F1, a njegovo UTF-8 kodiranje zauzima 2 bajta "c3i b1".
Pokušavamo da odštampamo karaktere pod pretpostavkom da će svaka kodna tačka
biti dugačka jedan bajt, što je pogrešno. U UTF-8 kodiranju, kodna tačka može
zauzimati više od 1 bajta. Pa kako da rešimo ovo? Tu nas spasava novi tip
podataka "Rune".

Rune

Runa je ugrađeni tip u programskom jeziku Go i to je alias tipa int32. Runa
predstavlja Unikod kodnu tačku u programskom jeziku Go. Nije bitno koliko
bajtova kodna tačka zauzima, može biti predstavljena runom. Hajde da
modifikujemo gornji program da ispisuje znakove koristeći runu.
*/

func printChars2(s string) {
	fmt.Printf("Characters: ")
	runes := []rune(s) // String conversion to rune slice
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c ", runes[i])
	}
	fmt.Println()
}

func stringRune() {

	fmt.Println("\n --- stringRune ---")

	name := "Hello World"
	fmt.Printf("String: %s\n", name)
	printChars2(name)
	printBytes(name)

	fmt.Println()

	name = "Señor"
	fmt.Printf("String: %s\n", name)
	printChars2(name)
	printBytes(name)
}

/*
U gornjem programu, string se konvertuje u isečak runa. Zatim ga iteriramo for
petljom i prikazujemo znakove. Ovaj program ispisuje

	>> String: Hello World
	>> Characters: H e l l o   W o r l d
	>> Bytes: 48 65 6c 6c 6f 20 57 6f 72 6c 64

	>> String: Señor
	>> Characters: S e ñ o r
	>> Bytes: 53 65 c3 b1 6f 72

Gore navedeni rezultat je savršen. Baš ono što smo želeli 😀.

Pristup pojedinačnim runama korišćenjem for range petlje
--------------------------------------------------------
Gore navedeni program je savršen način za iteraciju kroz pojedinačne rune u
stringu. Ali Go nam nudi mnogo lakši način da to uradimo koristeći for range
petlju.
*/

func charsAndBytePosition(s string) {
	for index, rune := range s {
		fmt.Printf("%c starts at byte %d\n", rune, index)
	}
}

func stringForRangeRune() {

	fmt.Println("\n --- stringForRangeRune ---")

	name := "Señor"
	charsAndBytePosition(name)
}

/*
U gornjem programu, string se iterira pomoću for range petlje. Petlja vraća
poziciju bajta gde runa počinje, zajedno sa runom. Ovaj program štampa:

	>> S starts at byte 0
	>> e starts at byte 1
	>> ñ starts at byte 2
	>> o starts at byte 4
	>> r starts at byte 5

Iz gornjeg izlaza je jasno da ñ zauzima 2 bajta jer sledeći znak o počinje od
bajta 4 umesto od bajta 3 😀.

Kreiranje stringa od isečka bajtova
-----------------------------------
*/

func stringFromSliceBytes() {

	fmt.Println("\n --- stringFromSliceHexaBytes ---")

	byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
	str := string(byteSlice)
	fmt.Println(str)
}

/*
byteSlice sadrži UTF-8 kodirane heksadecimalne bajtove stringa Café.
Program ispisuje:

	>> Café

Šta ako imamo decimalni ekvivalent heksadecimalnih vrednosti. Da li će gornji
program raditi? Hajde da proverimo.
*/

func stringFromSliceDecimalBytes() {

	fmt.Println("\n --- stringFromSliceDecimalBytes ---")
	//decimal equivalent of {'\x43', '\x61', '\x66', '\xC3', '\xA9'}
	byteSlice := []byte{67, 97, 102, 195, 169}
	str := string(byteSlice)
	fmt.Println(str)
}

/*
Decimalne vrednosti takođe rade i gornji program će takođe ispisati

	>> Café.

Kreiranje stringa od isečka runa
--------------------------------
*/

func stringFromSliceRune() {
	fmt.Println("\n --- stringFromSliceRune ---")

	runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
	str := string(runeSlice)
	fmt.Println(str)
}

/*
Gore navedeni program sadrži Unicode kodne tačke stringa Señor u heksadecimalnom
obliku. Program štampa:

	>> Señor

Dužina stringa
--------------
Funkcija RuneCountInString(s string) (n int) iz paketa utf8 može se koristiti
za pronalaženje dužine stringa. Ova metoda uzima string kao argument i vraća
broj runa u njemu.

Kao što smo ranije pomenuli, len(s) se koristi za pronalaženje broja bajtova u
stringu i ne vraća dužinu stringa. Neki Unicode znakovi imaju kodne tačke koje
zauzimaju više od 1 bajta. Korišćenje len za pronalaženje dužine tih stringova
vratiće pogrešnu dužinu stringa.
*/

func stringLen() {

	fmt.Println("\n --- stringLen ---")

	word1 := "Señor"
	fmt.Printf("String: %s\n", word1)
	fmt.Printf("Length: %d runes\n", utf8.RuneCountInString(word1))
	fmt.Printf("Number of bytes: %d\n", len(word1))

	fmt.Printf("\n")

	word2 := "Pets"
	fmt.Printf("String: %s\n", word2)
	fmt.Printf("Length: %d runes\n", utf8.RuneCountInString(word2))
	fmt.Printf("Number of bytes: %d", len(word2))

	fmt.Printf("\n")
}

/*
Izlaz gornjeg programa je

	>> String: Señor
	>> Length: 5 runes
	>> Number of bytes: 6
	>>
	>> String: Pets
	>> Length: 4 runes
	>> Number of bytes: 4

Gore navedeni izlaz to potvrđuje. len(s)i RuneCountInString(s) vraćaju
različite vrednosti 😀.

Poređenje stringova
-------------------
Operator == se koristi za upoređivanje strinova radi poređenja jednakosti.
Ako su stringovi jednaki, onda je rezultat true, inače je false.
*/

func compareStrings(str1 string, str2 string) {
	if str1 == str2 {
		fmt.Printf("%s and %s are equal\n", str1, str2)
		return
	}
	fmt.Printf("%s and %s are not equal\n", str1, str2)
}

func stringCompare() {

	fmt.Println("\n --- stringCompare ---")

	string1 := "Señor"
	string2 := string([]rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}) //"Señor"
	compareStrings(string1, string2)

	string3 := "hello"
	string4 := "world"
	compareStrings(string3, string4)
}

/*
U compareStrings funkciji upoređujemo da li su dva stringa str1 i str2 jednaka
koristeći == operator. Ako su jednaki, ispisuje odgovarajuću poruku i funkcija
se vraća.

Gore navedeni program štampa,

	>> Go and Go are equal
	>> hello and world are not equal

Spajanje stringova
------------------
Postoji više načina za spajanje stringova u Gou. Hajde da pogledamo nekoliko
njih.

Najjednostavniji način za spajanje stringova je korišćenje + operatora.
*/

func stringConcat() {

	fmt.Println("\n --- stringConcat ---")

	string1 := "Go"
	string2 := "is awesome"

	result := string1 + " " + string2
	fmt.Println(result)
}

/*
U gornjem programu, string1 je spojen sa string2 a razmakom u sredini. Ovaj
program ispisuje:

	>> Go is awesome

Drugi način spajanja stringova je korišćenje funkcije Sprintf iz fmt paketa.

Funkcija Sprintf formira string prema specifikatoru ulaznog formata i vraća
rezultujući string. Hajde da prepišemo gornji program koristeći Sprintf
funkciju.
*/
func stringSprintf() {

	fmt.Println("\n --- StringSprintf ---")

	string1 := "Go"
	string2 := "is awesome"
	result := fmt.Sprintf("%s %s", string1, string2)
	fmt.Println(result)
}

/*
Specifikatori %s %s su specifikatori formata za Sprintf. Ovaj specifikator
formata uzima dva stringa kao ulaz i ima razmak između njih. Ovo će spojiti dva
stringa sa razmakom u sredini. Rezultujući string se čuva u result. Ovaj program
takođe ispisuje:

	>> Go is awesome

Stringovi su nepromenjivi
-------------------------
Stringovi su nepromenjivi u Gou. Kada se string kreira, nije ga moguće promeniti.
*/

// func mutateString(s string) string {
// 	s[0] = 'a' //any valid unicode character within single quote is a rune
// 	return s
// }

// func stringMutate() {

// 	fmt.Println("\n --- stringMutate ---")

// 	h := "hello"
// 	fmt.Println(mutateString(h))
// }

/*
U gornjem programu, pokušavamo da promenimo prvi karakter stringa na 'a'. Bilo
koji važeći Unicode karakter unutar jednog navodnika je runa. Pokušavamo da
dodelimo runu a nultoj poziciji isečka. Ovo nije dozvoljeno jer je string
nepromenjiv i stoga program ne uspeva da se kompajlira. Greška je:

	>> ./prog.go:8:7: ne može se dodeliti s[0]

Da bi se zaobišla ova nepromenjivost nizova, nizovi se konvertuju u isečak runa.
Zatim se taj deo mutira sa svim potrebnim promenama i konvertuje nazad u novi
string.
*/

func mutateString(s []rune) string {
	s[0] = 'a'
	return string(s)
}

func stringMutate() {

	fmt.Println("\n --- stringMutate ---")

	h := "čćžšđ"
	fmt.Println(mutateString([]rune(h)))
}

/*
U gornjem programu, mutateString funkcija prihvata isečak rune kao argument.
Zatim menja prvi element isečka u 'a', konvertuje runu nazad u string i vraća ga.

Ova metoda se poziva, h se konvertuje u isečak runa i prosleđuje mutateStringu
Ovaj program štampa:

aćžšđ
*/

func StringFuncs() {
	stringElem()
	stringAccessBytes()
	stringAccessChars()
	stringError()
	stringRune()
	stringForRangeRune()
	stringFromSliceBytes()
	stringFromSliceDecimalBytes()
	stringFromSliceRune()
	stringLen()
	stringCompare()
	stringConcat()
	stringSprintf()
	stringMutate()
}
