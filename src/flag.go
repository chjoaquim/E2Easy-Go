package e2easy

import "flag"

func ConfigureFlags() {
	flag.String("file", "example.yaml", "Data file full path")
	flag.Parse()
}
