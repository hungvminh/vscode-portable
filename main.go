//go:generate go install -v github.com/josephspurrier/goversioninfo/cmd/goversioninfo
//go:generate goversioninfo -icon=res/papp.ico -manifest=res/papp.manifest
package main

import (
	"os"
	"path"

	"github.com/portapps/portapps/v3"
	"github.com/portapps/portapps/v3/pkg/log"
	"github.com/portapps/portapps/v3/pkg/utl"
)

type config struct {
	Cleanup bool `yaml:"cleanup" mapstructure:"cleanup"`
}

var (
	app *portapps.App
	cfg *config
)

func init() {
	var err error

	// Default config
	cfg = &config{
		Cleanup: false,
	}

	// Init app with custom branding
	if app, err = portapps.NewWithCfg("vscode-portable", "Visual Studio Code Portable (Custom Build)", cfg); err != nil {
		log.Fatal().Err(err).Msg("Cannot initialize application. See log file for more info.")
	}
}

func main() {
	// Create data directory
	utl.CreateFolder(app.DataPath)
	
	// Set VSCode executable path
	app.Process = utl.PathJoin(app.AppPath, "Code.exe")
	
	// VSCode arguments for better debugging
	app.Args = []string{
		"--log", "info",
	}

	// Cleanup on exit if enabled
	if cfg.Cleanup {
		defer func() {
			log.Info().Msg("Cleaning up temporary files...")
			utl.Cleanup([]string{
				path.Join(os.Getenv("APPDATA"), "Code"),
				path.Join(os.Getenv("LOCALAPPDATA"), "Programs", "Microsoft VS Code"),
			})
		}()
	}

	// Set VSCode environment variables for portable mode
	os.Setenv("VSCODE_APPDATA", utl.PathJoin(app.DataPath, "appdata"))
	os.Setenv("VSCODE_EXTENSIONS", utl.PathJoin(app.DataPath, "extensions"))
	
	// Only set logs if logging is enabled
	if !app.Config().Common.DisableLog {
		os.Setenv("VSCODE_LOGS", utl.PathJoin(app.DataPath, "logs"))
	}
	
	// Set portable flag
	os.Setenv("VSCODE_PORTABLE", "1")
	
	// Custom data directory
	os.Setenv("VSCODE_CWD", app.RootPath)

	// Log startup info
	log.Info().
		Str("version", "1.101-50").
		Str("build_by", "hungvminh").
		Str("data_path", app.DataPath).
		Msg("Starting VSCode Portable")

	// Cleanup and launch
	defer app.Close()
	app.Launch(os.Args[1:])
}