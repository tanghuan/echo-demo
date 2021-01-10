package jobs

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
)

func init() {

	scheduler := gocron.NewScheduler(time.UTC)

	scheduler.Every(30).Minutes().Do(sync)

	scheduler.StartAsync()
}

func sync() {
	fmt.Printf("%s sync ...... \n", time.Now().Format("2006-01-02 15:04:05"))
}
