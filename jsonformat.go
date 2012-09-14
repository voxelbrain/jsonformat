package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

const (
	VERSION = "0.3.1"
)

var (
	formatFlag = flag.String("format", "csv", "Name of formatter")
	fatalFlag  = flag.Bool("fatal", false, "Do not continue on error")
	inputFlag  = flag.String("input", "", "Read input from file instead of stdin")
	helpFlag   = flag.Bool("help", false, "Show verbose help")
)

func main() {
	flag.Parse()

	if *helpFlag || (flag.NArg() != 1 && len(*inputFlag) <= 0) {
		fmt.Println("Usage: jsonformat [options] <format string>")
		flag.PrintDefaults()
		fmt.Println("Formatters:")
		for name, format := range Formats {
			fmt.Printf("\t\"%s\": %s\n", name, format.Description)
		}
		fmt.Printf("Version %s\n", VERSION)
		return
	}

	format, ok := Formats[*formatFlag]
	if !ok {
		log.Fatalf("Unknown format %s", *formatFlag)
	}

	f, err := format.Compiler(formatString())
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
		err = f.Execute(os.Stdout, input)
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

func formatString() string {
	if len(*inputFlag) > 0 {
		d, e := ioutil.ReadFile(*inputFlag)
		if e != nil {
			log.Fatalf("Could not read file \"%s\": %s", *inputFlag, e)
		}
		return string(d)
	}
	return flag.Arg(0)
}
