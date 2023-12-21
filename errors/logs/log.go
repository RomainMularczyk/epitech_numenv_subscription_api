package logs

import (
  "os"
  "log"
  "fmt"
)

type LogLevel int

const (
  INFO LogLevel = iota + 1
  WARNING
  ERROR 
)

func (logLevel LogLevel) String() string {
  return [...]string{"INFO", "WARNING", "ERROR"}[logLevel - 1]
}

func Output(logLevel LogLevel, errMsg string) {
  l := log.New(
    os.Stderr,
    fmt.Sprintf("%s: ", logLevel.String()),
    1,
  )
  l.Print(errMsg)
}
