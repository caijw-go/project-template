package main

import (
    "flag"
    "fmt"
    "github.com/caijw-go/library/base"
    "github.com/caijw-go/library/log"
    _log "log"
    "service/applications/admin"
    "service/applications/job"
)

func init() {
    base.Init(base.Template{
        AppYamlPath: []string{"application"},
    })
    if err := log.Init(base.Config().GetString("log.path")); err != nil {
        _log.Fatalln("log init error", err)
    }
}

func main() {
    var app string
    flag.StringVar(&app, "app", "", "program start type web|job")
    flag.Parse()

    switch app {
    case "job":
        job.Init()
    case "admin":
        admin.Init()
    default:
        fmt.Printf("app:【%s】不被支持", app)
    }
}
