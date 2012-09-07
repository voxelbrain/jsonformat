package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"text/template"
)

var (
	fatalFlag = flag.Bool("fatal", false, "Do not continue on error")
	helpFlag  = flag.Bool("help", false, "Show this help")
)

func main() {
	flag.Parse()

	if *helpFlag || flag.NArg() != 1 {
		fmt.Println("Usage: jsonformat [options] <template>")
		flag.PrintDefaults()
		return
	}

	t, err := template.New("jsonformat").Parse(flag.Arg(0))
	if err != nil {
		log.Fatalf("Template invalid: %s", err)
	}

	dec := json.NewDecoder(os.Stdin)
	logFn := NewLogFn(*fatalFlag)
	for {
		var input interface{}
		err := dec.Decode(&input)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Could not decode input: %s", err)
			continue
		}
		err = t.Execute(os.Stdout, input)
		if err != nil {
			logFn("Could not apply input to template: %s", err)
			continue
		}
		io.WriteString(os.Stdout, "\n")
	}
}

type LogFn func(format string, v ...interface{})

func NewLogFn(fatal bool) LogFn {
	if fatal {
		return func(format string, v ...interface{}) {
			log.Fatalf(format, v...)
		}
	}
	return func(format string, v ...interface{}) {
		log.Printf(format, v...)
	}
}
