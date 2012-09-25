package main

import (
	"encoding/json"
	"github.com/voxelbrain/goptions"
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
	options = struct {
		Fatal bool	 `goptions:"-f, --fatal, description='Do not continue on error'"`
		Input string `goptions:"-i, --input, description='Read input from file instead of stdin'"`
		Format string `goptions:"-f, --format, obligatory, description='Name for the formatter'"`
		goptions.Help
	}{
		Format: "csv",
		Fatal: false,
		Input: "",
	}
)

func main() {
	err := goptions.Parse(&options)
	if err != nil {
		goptions.PrintHelp()
		fmt.Println("Formatters:")
		for name, format := range Formats {
			fmt.Printf("\t\"%s\": %s\n", name, format.Description)
		}
		fmt.Printf("Version %s\n", VERSION)
		return
	}

	format, ok := Formats[options.Format]
	if !ok {
		log.Fatalf("Unknown format %s", options.Format)
	}

	f, err := format.Compiler(formatString())
	if err != nil {
		log.Fatalf("Template invalid: %s", err)
	}

	dec := json.NewDecoder(os.Stdin)
	logFn := NewLogFn(options.Fatal)
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
	if len(options.Input) > 0 {
		d, e := ioutil.ReadFile(options.Input)
		if e != nil {
			log.Fatalf("Could not read file \"%s\": %s", options.Input, e)
		}
		return string(d)
	}
	// return flag.Arg(0)
	return ""
}
