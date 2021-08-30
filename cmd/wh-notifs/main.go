package main

import (
	"fmt"
	"log"
	"time"

	"github.com/LompeBoer/wh-notifs/internal/database"
	"github.com/LompeBoer/wh-notifs/internal/database/whdbv1"
	"github.com/LompeBoer/wh-notifs/internal/discord"
)

const (
	DiscordWebHookURL = ""
	VersionNumber     = "0.2-alpha"
	StorageFile       = "storage.db"
	SleepTime         = 10 * time.Second
)

func main() {
	log.Printf("Starting wh-notifs v%s\n", VersionNumber)

	db := whdbv1.New(StorageFile)
	defer db.Close()

	discordWebHook := discord.DiscordWebHook{
		Enabled: true,
		URL:     DiscordWebHookURL,
	}
	checkTrades(db, discordWebHook)
}

func checkTrades(db database.DatabaseService, discordWebHook discord.DiscordWebHook) {
	check := time.Now()
	for {
		log.Printf("check trades since %s\n", check.Format("2006-01-02T15:04:05.999999999Z07:00"))
		trades, err := db.SelectLatestTrades(check)
		if err != nil {
			log.Printf("ERROR: %s\n", err.Error())
		}
		check = time.Now()

		log.Printf("trades: %d\n", len(trades))

		for _, t := range trades {
			log.Printf("New trade row: %+v\n", t)
			datetime, err := time.Parse("2006-01-02T15:04:05.999999999Z", t.DateTime)
			if err != nil {
				datetime = time.Unix(0, 0)
			}
			msg := fmt.Sprintf(
				"**Status** %s **Symbol:** %s **Side:** %s **BuyCount:** %d **AvgPrice**: %f **TPLimitPrice:** %s **Reason:** %s **Time:** %s",
				t.Status,
				t.Symbol,
				t.Side,
				t.BuyCount,
				t.AveragePrice,
				t.TakeProfitLimitPrice,
				t.Reason,
				datetime.Format("2006-01-02 15:04:05.9"),
			)
			discordWebHook.SendMessage(msg)
		}

		time.Sleep(SleepTime)
	}
}
