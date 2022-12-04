package job

import "github.com/robfig/cron/v3"

func Init() {
    c := cron.New(cron.WithSeconds())
    c.AddJob("00 00 00 * * ?", &test{})

    c.Start()
    select {}
}
