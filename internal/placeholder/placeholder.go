package placeholder

import (
	"bufio"
	"fmt"
	"io"
	"path/filepath"
	"strings"

	"github.com/spf13/afero"
	"stash.tools.deloitteinnovation.us/projects/CRE/repos/concourse-ssh-resource/internal/models"
)

// ReplacePlaceholders replaces placeholders in Script with Placeholders
func ReplacePlaceholders(stderr io.Writer, fs afero.Fs, baseDir string, params *models.Params) (string, error) {
	var err error

	script := params.Script
	placeholders := params.Placeholders

	// replacing all placeholders, either given as static value using .Value
	// or as dynamic using .File
	for _, placeholder := range placeholders {
		var value string

		// File should always be used if conflicting
		if placeholder.File != "" {
			// Load content from File
			value, err = readLineFromFile(fs, filepath.Join(baseDir, placeholder.File))
			if err != nil {
				return "", err
			}
			if value == "" {
				return "", fmt.Errorf("File for placeholder '%s' seems to be empty", placeholder.Name)
			}
		} else if placeholder.Value != "" {
			// static Value
			value = placeholder.Value
		} else {
			fmt.Fprintf(stderr, "WARNING: Neither File nor Value are set for placeholder '%s'", placeholder.Name)
		}

		if strings.Contains(script, placeholder.Name) {
			script = strings.Replace(script, placeholder.Name, value, -1)
		} else {
			fmt.Fprintf(stderr, "WARNINIG: Placeholder '%s' is not found in script, maybe a typo?", placeholder.Name)
		}
	}

	return script, nil
}

func readLineFromFile(fs afero.Fs, filepath string) (string, error) {
	if filepath == "" {
		return "", nil
	}

	file, err := fs.Open(filepath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var contents []string
	// Read all lines of file
	for scanner.Scan() {
		contents = append(contents, scanner.Text())
	}
	return strings.Join(contents, "\n"), nil
}
