### Cara Pakai :

Instalasi `go get -u github.com/zahrulsch/golog`

```go
package main

import (
    stdLog "log"

    "github.com/zahrulsch/golog"
)

func main() {
    logFilePath := "app.log"
    appName := "My App"

    log, err := golog.NewLogger(logFilePath, appName)
    if err != nil {
        stdLog.Fatal(err)
    }
    defer log.Sync()
    // contoh penggunaan
    log.Info("Ini info")
}

```

| Parameter       | Type   | IsRequired |
| --------------- | ------ | ---------- |
| **logFilePath** | string | true       |
| **appName**     | string | true       |
