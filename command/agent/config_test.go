package agent

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestConfig_ParseConfigFile(t *testing.T) {
	// Fails if the file doesn't exist
	if _, err := ParseConfigFile("/wosniak/jobs"); err == nil {
		t.Fatalf("expected error, got nothing")
	}

	fh, err := ioutil.TempFile("", "replcaitor")
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	defer os.RemoveAll(fh.Name())

	// Invalid content returns error
	if _, err := fh.WriteString("throwingcoins"); err != nil {
		t.Fatalf("err: %s", err)
	}
	if _, err := ParseConfigFile(fh.Name()); err == nil {
		t.Fatalf("expected load error, got nothing")
	}

	// Valid content parses successfully
	if err := fh.Truncate(0); err != nil {
		t.Fatalf("err: %s", err)
	}
	if _, err := fh.Seek(0, 0); err != nil {
		t.Fatalf("err: %s", err)
	}
	if _, err := fh.WriteString(`{"aws_region":"us-east-1"}`); err != nil {
		t.Fatalf("err: %s", err)
	}

	config, err := ParseConfigFile(fh.Name())
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if config.Region != "us-east-1" {
		t.Fatalf("bad aws region: %q", config.Region)
	}
}

func TestConfig_LoadConfigDir(t *testing.T) {

	// Fails if the dir doesn't exist.
	if _, err := LoadConfigDir("/wosniak/jobs"); err == nil {
		t.Fatalf("expected error, got nothig")
	}

	dir, err := ioutil.TempDir("", "replicator")
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	defer os.RemoveAll(dir)

	// Returns empty config on empty dir
	config, err := LoadConfig(dir)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if config == nil {
		t.Fatalf("should not be nil")
	}

	file1 := filepath.Join(dir, "replicator.hcl")
	err = ioutil.WriteFile(file1, []byte(`{"aws_region":"us-east-1"}`), 0600)
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	file2 := filepath.Join(dir, "replicator_1.hcl")
	err = ioutil.WriteFile(file2, []byte(`{"scaling_interval":1}`), 0600)
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	// Works if configs are valid
	config, err = LoadConfigDir(dir)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if config.Region != "us-east-1" || config.ScalingInterval != 1 {
		t.Fatalf("bad: %#v", config)
	}
}