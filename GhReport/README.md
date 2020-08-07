GhReport jalutab läbi kausta (ja kõik selle alamkaustad), parsib .feature failid
ja koostab nimekirja Gherkinis (kergkeeles) spetsifitseeritud omadustest (feature)
ja stsenaariumitest (scenario). Iga stsenaariumi kohta väljastab sammude arvu.

Näiteks:

fail `Hilinemine.feature`:

````gherkin
# language: et
@Kevade
Omadus: Kooli jõudmine

Stsenaarium: Hilinemine
    Eeldades Päev on koolipäev
    Ja Tunnid algavad kell 8
    Ja Arno läheb kooli
    Ja Isa läheb temaga kaasa
    Ja Neil on pikk tee minna
    Kui Arno isaga jõuavad koolimajja
    Siis On tunnid juba alanud
````

fail `Vahetund.feature`:

````gherkin
# language: et
@Kevade
Omadus: Vahetund

Stsenaarium: Vahetund
    Eeldades Päev on koolipäev
    Ja Tunnid algavad kell 8
    Ja Arno läks kooli
    Kui kell saab 8.45
    Siis Algab Vahetund
````

Aruanne:

````
$ go run .
ARUANNE
.
Omadus: Kooli jõudmine
    Testistsenaarium: Hilinemine (7 sammu)
Omadus: Vahetund
    Testistsenaarium: Vahetund (5 sammu)

Faile: 2
````

Kasutamine:

`go run . [-dir <.feature-failide juurkaust>]`

Gherkin-keele eesti dialekt: https://cucumber.io/docs/gherkin/reference/#spoken-languages.
