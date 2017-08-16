package application

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

const emulatorHostEnvVar = "DATASTORE_EMULATOR_HOST"

// Config is configurations
type Config struct {
	ProjectID string
	Public    string
}

// App is application globals
type App struct {
	Config *Config
}

// NewApp returns new App instance
func NewApp(cnf *Config) (*App, error) {
	flag.Parse()

	// Check DATASTORE_EMULATOR_HOST env var is set
	emuHost := os.Getenv(emulatorHostEnvVar)
	if strings.TrimSpace(emuHost) == "" {
		return nil, fmt.Errorf("%s is not set in env", emulatorHostEnvVar)
	}
	log.Println("DATASTORE_EMULATOR_HOST", emuHost)

	return &App{
		Config: cnf,
	}, nil
}
