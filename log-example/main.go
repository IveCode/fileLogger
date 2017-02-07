package main

import (
        "time"
        "fmt"

        "github.com/aiwuTech/fileLogger"
)

var (
        log *fileLogger.FileLogger
)

func main() {
        log = fileLogger.NewSizeLogger("./", "test.log", "", 10, 1, fileLogger.MB, 60, fileLogger.DEFAULT_LOG_SEQ)
        log.SetLogLevel(0)

        for i := 0; i < 5; i++ {
                go func(i int) {
                        for {
                                log.I("test log file", i)
                                time.Sleep(time.Duration(time.Microsecond))
                                fmt.Println("test log file", i)
                        }
                }(i)
        }

        select {}
}
