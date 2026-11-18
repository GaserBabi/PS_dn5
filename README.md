# vreme

Paket `vreme` ponuja podatkovno strukturo in funkcije za upravljanje vremenskih postaj ter obdelavo temperaturnih meritev.

Paket omogoča dodajanje vremenskih postaj v slovar, beleženje novih temperaturnih meritev znotraj dovoljenega območja ter izpis zbranih meritev in statističnega poročila (najnižja, najvišja in povprečna temperatura ter opis vremena).

Primer uporabe:

```go
postaje := make(map[string]vreme.Postaja)
vreme.DodajPostajo(postaje, "LJ01", "Ljubljana Bežigrad", "Ljubljana", 18.2, 21.5)
vreme.DodajMeritev(postaje, "LJ01", 19.8, -50, 50)

vreme.IzpisMeritev(postaje)
vreme.IzpisPorocila(postaje, 3)
```

Privzete nastavitve pri obdelavi meritev:
- Dovoljeno območje temperature: -50–50 °C
- Najmanjše število meritev za statistiko: 3