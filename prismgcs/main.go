package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/apache/beam/sdks/v2/go/pkg/beam"
	_ "github.com/apache/beam/sdks/v2/go/pkg/beam/io/filesystem/gcs"
	_ "github.com/apache/beam/sdks/v2/go/pkg/beam/io/filesystem/local"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/io/textio"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/runners/prism"
	"os"
)

var (
	// Change this to a GCS location to reproduce the bugs
	inputs = flag.String("inputs", "hackedlogs.json", "Input file")
)

func run() error {
	beam.Init()
	p := beam.NewPipeline()
	s := p.Root()

	all := textio.Read(s, *inputs)

	spans := beam.ParDo(s, func(line string, emit func(string)) {
		output := fmt.Sprintf("line had length %d", len(line))
		fmt.Print(output + "\n")
		emit(output)
	}, all)

	textio.Write(s, "/tmp/output-test-pipeline", spans)
	if _, err := prism.Execute(context.Background(), p); err != nil {
		return err
	}
	return nil
}

func main() {
	flag.Parse()

	if err := run(); err != nil {
		fmt.Printf("Error running the runner;\n %+v\n", err)
		os.Exit(1)
	}
}
