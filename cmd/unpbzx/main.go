package main

import (
    "context"
    "log"
    "os"
    "runtime"

    "github.com/palantir/stacktrace"
    "github.com/zhangyoufu/unpbzx"
)

func main() {
    var src *os.File
    switch len(os.Args) {
    case 1:
        src = os.Stdin
    case 2:
        var err error
        src, err = os.Open(os.Args[1])
        if err != nil {
            log.Fatal(stacktrace.Propagate(err, "cannot open input file"))
        }
    default:
        log.Fatal("invalid usage")
    }

    ctx := context.Background()
    dst := os.Stdout
    numWorker := runtime.NumCPU()

    if err := unpbzx.Extract(ctx, src, dst, numWorker); err != nil {
        log.Fatal(stacktrace.Propagate(err, "error while extracting"))
    }
}
