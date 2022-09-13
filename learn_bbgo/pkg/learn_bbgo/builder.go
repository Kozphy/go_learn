package learn_bbgo

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// FIXME:
var wrapperTemplate = template.Must(template.New("main").Parse(`// Code generated by bbgo; DO NOT EDIT.
package main
import (
	"github.com/c9s/bbgo/pkg/bbgo"
	"github.com/c9s/bbgo/pkg/cmd"
{{- range .Imports }}
	_ "{{ . }}"
{{- end }}
)
func init() {
	bbgo.SetWrapperBinary()
}
func main() {
	cmd.Execute()
}
`))

func generateRunFile(filepath string, config *Config, imports []string) error {
	var buf = bytes.NewBuffer(nil)
	if err := wrapperTemplate.Execute(buf, struct {
		Config  *Config
		Imports []string
	}{
		Config:  config,
		Imports: imports,
	}); err != nil {
		return err
	}

	return os.WriteFile(filepath, buf.Bytes(), 0644)
}

func compilePackage(packageDir string, userConfig *Config, imports []string) error {
	if _, err := os.Stat(packageDir); os.IsNotExist(err) {
		if err := os.MkdirAll(packageDir, 0777); err != nil {
			return errors.Wrapf(err, "can not create wrapper package directory: %s", packageDir)
		}
	}
	mainFile := filepath.Join(packageDir, "main.go")
	if err := generateRunFile(mainFile, userConfig, imports); err != nil {
		return errors.Wrap(err, "compile error")
	}
	return nil
}

func Build(ctx context.Context, userConfig *Config, targetConfig BuildTargetConfig) (string, error) {
	imports := append(userConfig.Imports, targetConfig.Imports...)

	buildDir := userConfig.BuildDir
	if len(buildDir) == 0 {
		buildDir = "build"
	}

	packageDir, err := os.MkdirTemp(buildDir, "bbgow-")
	if err != nil {
		return "", err
	}

	if err := compilePackage(packageDir, userConfig, imports); err != nil {
		return "", err
	}

	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	var buildEnvs []string

	if targetConfig.OS != runtime.GOOS {
		buildEnvs = append(buildEnvs, "GOOS="+targetConfig.OS)
	}
	if targetConfig.Arch != runtime.GOARCH {
		buildEnvs = append(buildEnvs, "GOARCH="+targetConfig.Arch)
	}

	buildTarget := filepath.Join(cwd, packageDir)

	binary := targetConfig.Name
	if len(binary) == 0 {
		binary = fmt.Sprintf("bbgow-%s-%s", targetConfig.OS, targetConfig.Arch)
	}

	output := filepath.Join(buildDir, binary)
	args := []string{"build", "-tags", "wrapper", "-o", output, buildTarget}
	log.Debugf("building binary %s from %s: go %v", output, buildTarget, args)
	buildCmd := exec.CommandContext(ctx, "go", args...)
	buildCmd.Env = append(os.Environ(), buildEnvs...)
	buildCmd.Stdout = os.Stdout
	buildCmd.Stderr = os.Stderr

	if err := buildCmd.Run(); err != nil {
		return output, err
	}
	return output, os.RemoveAll(packageDir)

}