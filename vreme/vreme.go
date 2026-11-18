package vreme

import (
	"fmt"
	"sort"
)

type Postaja struct {
	ime     string
	kraj    string
	meritve []float64
}

func DodajPostajo(postaje map[string]Postaja, oznaka, ime, kraj string, meritve ...float64) {
	kopijaMeritev := append([]float64(nil), meritve...)
	postaje[oznaka] = Postaja{
		ime:     ime,
		kraj:    kraj,
		meritve: kopijaMeritev,
	}
}

func DodajMeritev(postaje map[string]Postaja, oznaka string, temperatura, minTemp, maxTemp float64) {
	if temperatura < minTemp || temperatura > maxTemp {
		fmt.Println("Izven dovoljenega območja")
		return
	}

	postaja, obstaja := postaje[oznaka]
	if !obstaja {
		fmt.Println("Postaja ne obstaja")
		return
	}

	postaja.meritve = append(postaja.meritve, temperatura)
	postaje[oznaka] = postaja
}

func statistika(postaje map[string]Postaja, oznaka string, stMeritev int) (najnizja, najvisja, povprecna float64, ok bool) {
	postaja, obstaja := postaje[oznaka]
	if !obstaja || len(postaja.meritve) < stMeritev || len(postaja.meritve) == 0 {
		return 0, 0, 0, false
	}

	najnizja = postaja.meritve[0]
	najvisja = postaja.meritve[0]
	vsota := 0.0

	for _, m := range postaja.meritve {
		if m < najnizja {
			najnizja = m
		}
		if m > najvisja {
			najvisja = m
		}
		vsota += m
	}

	povprecna = vsota / float64(len(postaja.meritve))
	return najnizja, najvisja, povprecna, true
}

func IzpisMeritev(postaje map[string]Postaja) {
	fmt.Println("MERITVE:")
	oznake := uredi(postaje)

	for _, oznaka := range oznake {
		p := postaje[oznaka]
		fmt.Printf("%s - %s (%s): %v\n", oznaka, p.ime, p.kraj, p.meritve)
	}
}

func IzpisPorocila(postaje map[string]Postaja, stMeritev int) {
	oznake := uredi(postaje)

	for _, oznaka := range oznake {
		p := postaje[oznaka]
		min, max, avr, ok := statistika(postaje, oznaka, stMeritev)

		if !ok {
			fmt.Printf("%s: premalo meritev\n", p.ime)
			continue
		}

		var opis string
		switch {
		case avr >= 25:
			opis = "Vroče"
		case avr >= 15:
			opis = "Toplo"
		case avr >= 5:
			opis = "Hladno"
		default:
			opis = "Mrzlo"
		}

		fmt.Printf("%s: min %.1f °C, max %.1f °C, povprečje %.1f °C -> %s\n", p.ime, min, max, avr, opis)
	}
}

func uredi(postaje map[string]Postaja) []string {
	oznake := make([]string, 0, len(postaje))
	for k := range postaje {
		oznake = append(oznake, k)
	}
	sort.Strings(oznake)
	return oznake
}