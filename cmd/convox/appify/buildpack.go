package appify

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strings"

	yaml "gopkg.in/yaml.v2"

	"github.com/convox/rack/cmd/convox/helpers"
	"github.com/convox/rack/manifest"
)

// Buildpack type representing a Heroku Buildpack
type Buildpack struct {
	manifest  manifest.Manifest
	setup     bool
	tmplInput map[string]interface{}
}

// EnvEntry is an environment entry from an app.json file
type EnvEntry struct {
	Value string
}

// Appfile represent specific fields of an app.json file
type Appfile struct {
	Addons []string
	Env    map[string]EnvEntry
}

// ProcfileEntry is an entry in a Procfile
type ProcfileEntry struct {
	Name    string
	Command string
}

// Procfile represents a Procfile used in Heroku-based apps
type Procfile []ProcfileEntry

var reProcfile = regexp.MustCompile("^([A-Za-z0-9_]+):\\s*(.+)$")

// Appify generates the files needed for an app
// Must be called after Setup()
func (bp *Buildpack) Appify() error {
	if !bp.setup {
		return fmt.Errorf("must call Setup() first")
	}

	if err := writeAsset("entrypoint.sh", "appify/templates/buildpack/entrypoint.sh", nil); err != nil {
		return err
	}

	if err := writeAsset("Dockerfile", "appify/templates/buildpack/Dockerfile", bp.tmplInput); err != nil {
		return err
	}

	if err := writeAsset(".dockerignore", "appify/templates/buildpack/.dockerignore", nil); err != nil {
		return err
	}

	data, err := yaml.Marshal(bp.manifest)
	if err != nil {
		return err
	}

	return writeFile("docker-compose.yml", data, 0644)
}

// Setup reads the location for Buildpack specifc files and settings
func (bp *Buildpack) Setup(location string) error {
	kind := detectBuildpack(location)
	if kind == "unknown" {
		// TODO: track this event
		return fmt.Errorf("unknown Buildpack type")
	}

	bp.tmplInput = map[string]interface{}{
		"framework":   kind,
		"environment": buildpackEnvironment(kind),
	}

	pf, err := readProcfile("Procfile")
	if err != nil {
		return fmt.Errorf("reading Procfile : %s", err)
	}

	af, err := readAppfile("app.json")
	if err != nil {
		return fmt.Errorf("reading app.json : %s", err)
	}

	bp.manifest = generateManifest(pf, af)

	if len(af.Addons) > 0 {
		bp.manifest = parseAddons(af, bp.manifest)
	}

	bp.setup = true

	return nil
}

func detectBuildpack(location string) string {
	switch {
	case helpers.Exists(filepath.Join(location, "requirements.txt")) || helpers.Exists(filepath.Join(location, "setup.py")):
		return "python"
	case helpers.Exists(filepath.Join(location, "package.json")):
		return "nodejs"
	case helpers.Exists(filepath.Join(location, "Gemfile")):
		return "ruby"
	}

	return "unknown"
}

// buildpackEnvironment creates environment variables that are buildpack specific
func buildpackEnvironment(kind string) map[string]string {
	switch kind {
	case "ruby":
		return map[string]string{
			"CURL_CONNECT_TIMEOUT": "0", // default timeout is too aggressive causing failure
			"STACK":                "cedar-14",
		}
	default:
		return map[string]string{}
	}
}

func readProcfile(path string) (Procfile, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	pf := Procfile{}

	scanner := bufio.NewScanner(bytes.NewReader(data))

	for scanner.Scan() {
		parts := reProcfile.FindStringSubmatch(scanner.Text())

		if len(parts) == 3 {
			pf = append(pf, ProcfileEntry{
				Name:    parts[1],
				Command: parts[2],
			})
		}
	}

	return pf, nil
}

func readAppfile(path string) (Appfile, error) {
	af := Appfile{}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return af, err
	}

	if err := json.Unmarshal(data, &af); err != nil {
		return af, err
	}

	return af, nil
}

// generateManifest generates a Manifest from a Procfile
func generateManifest(pf Procfile, af Appfile) manifest.Manifest {

	m := manifest.Manifest{
		Services: make(map[string]manifest.Service),
	}

	for _, e := range pf {
		me := manifest.Service{
			Build: &manifest.Build{
				Context: ".",
			},
			Command: &manifest.Command{
				String: e.Command,
			},
			Environment: make(manifest.Environment),
			Labels:      make(manifest.Labels),
			Ports:       make(manifest.Ports, 0),
		}

		switch e.Name {
		case "web":
			me.Labels["convox.port.443.protocol"] = "tls"

			me.Ports = append(me.Ports, manifest.Port{
				Balancer:  80,
				Container: 4001,
				Public:    true,
			})
			me.Ports = append(me.Ports, manifest.Port{
				Balancer:  443,
				Container: 4001,
				Public:    true,
			})

			me.Environment["PORT"] = "4001"

			for k, v := range af.Env {
				me.Environment[k] = v.Value
			}
		}

		m.Services[e.Name] = me
	}

	return m
}

func parseAddons(af Appfile, m manifest.Manifest) manifest.Manifest {

	for _, addon := range af.Addons {
		switch {
		case strings.Contains(addon, "postgres"):
			m = postgresAddon(af, m)
		}
	}

	return m
}

func postgresAddon(af Appfile, m manifest.Manifest) manifest.Manifest {
	s := manifest.Service{
		Image: "convox/postgres",
		Name:  "database",
		Ports: manifest.Ports{
			{Container: 5432, Public: false},
		},
		Volumes: []string{
			"/var/lib/postgresql/data",
		},
	}

	web := m.Services["web"]
	web.Links = append(web.Links, "database")
	m.Services["web"] = web

	m.Services["database"] = s

	return m
}
