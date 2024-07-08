package main

import (
    "os"
    "fmt"
    "time"
    "bufio"
    "strings"
    "github.com/fatih/color"
)


func main() {
    // Check arguments
    if len(os.Args) < 2 {
        fmt.Println("Usage: log <level> <message>")
        os.Exit(1)
    }

    // Read loglevel as first positional argument of the cli
    var logLevel string = strings.ToUpper(os.Args[1])

    // Format a timestamp
    var timestamp string = time.Now().Format("2006-01-02 15:04:05")

    // Set the generic bold white color as filler
    var filler func(a ...interface{}) string = color.New(color.FgWhite, color.Bold).SprintFunc()

    // If there is anything on stdin, read that, otherwise use args
    stat, _ := os.Stdin.Stat()
    var message string

    if (stat.Mode() & os.ModeCharDevice) == 0 {
        reader := bufio.NewReader(os.Stdin)
        message, _ = reader.ReadString('\n')
    } else {
        message = strings.Join(os.Args[2:], " ")
    }

    var colfunc func(a ...interface{}) string

    switch logLevel {
        case "DEBUG":
            colfunc = color.New(color.FgMagenta, color.Bold).SprintFunc()
        case "INFO" , "OK":
            colfunc = color.New(color.FgGreen, color.Bold).SprintFunc()
        case "WARNING", "WARN", "DANGER":
            colfunc = color.New(color.FgYellow, color.Bold).SprintFunc()
        case "ERROR", "CRIT", "CRITICAL", "FATAL":
            colfunc = color.New(color.FgRed, color.Bold).SprintFunc()
        default:
            colfunc = color.New(color.FgCyan, color.Bold).SprintFunc()
    }

    fmt.Printf(
        "%s %s - %s - %s\n",
        colfunc("[+]"),
        filler(timestamp),
        colfunc(logLevel),
        filler(message),
    )
}
