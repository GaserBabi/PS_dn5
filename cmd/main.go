package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/GaserBabi/PS_dn5/vreme"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "postaje",
		Usage: "vodenje meritev in izpis statistike",
		Flags: []cli.Flag{
			&cli.FloatFlag{
				Name:  "minTemp",
				Value: -50.0,
				Usage: "najnižja temperatura",
			},
			&cli.FloatFlag{
				Name:  "maxTemp",
				Value: 50.0,
				Usage: "najvišja temperatura",
			},
			&cli.IntFlag{
				Name:  "stMeritev",
				Value: 3,
				Usage: "število meritev",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			minTemp := cmd.Float("minTemp")
			maxTemp := cmd.Float("maxTemp")
			stMeritev := cmd.Int("stMeritev")

			postaje := make(map[string]vreme.Postaja)

			vreme.DodajPostajo(postaje, "LJ01", "Ljubljana Bežigrad", "Ljubljana", 18.2, 21.5, 19.8, 16.5)
			vreme.DodajPostajo(postaje, "MB01", "Maribor Tabor", "Maribor", 12.4)
			vreme.DodajPostajo(postaje, "KR01", "Kredarica", "Triglav", 2, -1.5, 3.5)
			vreme.DodajPostajo(postaje, "KP01", "Koper Luka", "Koper", 24.5, 27, 26.1, 28.4)
			vreme.DodajPostajo(postaje, "CE01", "Celje jug", "Celje", 24.5, 17, 20.1, 19.4)
			vreme.DodajPostajo(postaje, "NM01", "Novo mesto sever", "Novo Mesto", 5, -1)
			vreme.DodajPostajo(postaje, "SG01", "Slovenj Gradec zahod", "Slovenj Gradec", 18.2, 19.5, 20.8, -16.5)
			vreme.DodajPostajo(postaje, "KK01", "Krško vzhod", "Krško", 15.4, 11.3, 12.4)

			vreme.IzpisMeritev(postaje)
			fmt.Println()

			vreme.IzpisPorocila(postaje, stMeritev)
			fmt.Println()

			vreme.DodajMeritev(postaje, "LJ01", 60.5, minTemp, maxTemp)
			vreme.DodajMeritev(postaje, "NEOBSTOJECA", 20.0, minTemp, maxTemp)
			vreme.DodajMeritev(postaje, "CE01", 104.2, minTemp, maxTemp)
			vreme.DodajMeritev(postaje, "NM01", 15.1, minTemp, maxTemp)
			fmt.Println()

			vreme.IzpisMeritev(postaje)
			fmt.Println()

			vreme.IzpisPorocila(postaje, stMeritev)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
