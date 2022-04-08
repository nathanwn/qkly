package workspace

import (
	"testing"

	"github.com/spf13/afero"
)

func TestParseConfig(t *testing.T) {
	fs := afero.NewMemMapFs()
	f, err := fs.Create("./qkly.yaml")
	if err != nil {
		t.Fatalf("Error while creating config file: %s\n", err)
	}

	_, err = f.WriteString("port: 10042")
	if err != nil {
		t.Fatalf("Error while writing config file: %s\n", err)
	}
	f.Close()

	conf, err := NewConfig(fs)

	if err != nil {
		t.Fatalf("Error: Could not parse config file due to error %v\n", err)
	}

	if conf.Port() != "10042" {
		t.Fatalf(
			"Wrong port number: expected \"10042\", found %s\n",
			conf.Port(),
		)
	}
}
