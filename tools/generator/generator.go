package generator

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Generator struct {
	SpecPath    string
	OutputPath  string
	Whitelist   map[string]struct{}
	Definitions map[string]StructureDefinition
}

func NewGenerator(specPath, outputPath string, whitelist []string) *Generator {
	wl := make(map[string]struct{})
	for _, v := range whitelist {
		wl[v] = struct{}{}
	}
	return &Generator{
		SpecPath:    specPath,
		OutputPath:  outputPath,
		Whitelist:   wl,
		Definitions: make(map[string]StructureDefinition),
	}
}

func (g *Generator) Load(filename string) error {
	path := filepath.Join(g.SpecPath, filename)
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	var bundle StructureDefinitionBundle
	if err := json.Unmarshal(data, &bundle); err != nil {
		return fmt.Errorf("unmarshal %s: %w", filename, err)
	}

	for _, entry := range bundle.Entry {
		if entry.Resource.Name != "" {
			g.Definitions[entry.Resource.Name] = entry.Resource
		}
	}
	return nil
}

func (g *Generator) Generate() error {
	if err := os.MkdirAll(g.OutputPath, 0755); err != nil {
		return err
	}

	entries, err := os.ReadDir(g.OutputPath)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("read output directory: %w", err)
	}
	for _, entry := range entries {
		if !entry.IsDir() && filepath.Ext(entry.Name()) == ".go" {
			if err := os.Remove(filepath.Join(g.OutputPath, entry.Name())); err != nil {
				return fmt.Errorf("remove old file %s: %w", entry.Name(), err)
			}
		}
	}

	for name := range g.Whitelist {
		if def, ok := g.Definitions[name]; ok {
			if err := g.WriteResource(def); err != nil {
				return err
			}
		}
	}

	for tName := range usedTypes {
		if _, isWhite := g.Whitelist[tName]; isWhite {
			continue
		}
		if def, ok := g.Definitions[tName]; ok {
			if err := g.WriteResource(def); err != nil {
				return err
			}
		}
	}
	return nil
}
