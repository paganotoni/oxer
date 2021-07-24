package docker

import (
	"io/ioutil"
	"os"
	"oxer/internal/meta"
	"path/filepath"
	"strings"
	"testing"
)

func TestReplaceDockerfile(t *testing.T) {
	d := t.TempDir()
	os.Chdir(d)

	// Create a Dockerfile
	err := ioutil.WriteFile(filepath.Join(d, "Dockerfile"), []byte("FROM scratch"), 0777)
	if err != nil {
		t.Fatal(err)
	}

	err = ReplaceDockerfile(meta.Application{ModuleShort: "hnib", Module: "github.com/eventpipe/hnib"})
	if err != nil {
		t.Fatalf("error processing:%v", err)
	}

	// Check that the Dockerfile has been replaced
	b, err := ioutil.ReadFile("Dockerfile")
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(string(b), "FROM golang:1.16-alpine ") {
		t.Fatalf("Dockerfile not replaced:%s", string(b))
	}

	if !strings.Contains(string(b), "RUN ox build --static -o bin/hnib") {
		t.Fatalf("Dockerfile not replaced:%s", string(b))
	}

	if !strings.Contains(string(b), "WORKDIR /hnib") {
		t.Fatalf("Dockerfile not replaced:%s", string(b))
	}

	if !strings.Contains(string(b), "COPY --from=builder /hnib/bin/* /bin/") {
		t.Fatalf("Dockerfile not replaced:%s", string(b))
	}

}
