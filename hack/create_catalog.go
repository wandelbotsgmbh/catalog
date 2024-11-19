package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

const (
	catalogDir     = "catalog"
	catalogOutFile = "catalog.yaml"
	catalogHeader  = "# This file is autogenerated by 'create_catalog.go'. Do not edit manually.\n"
)

type CatalogEntry struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
}

type Catalog struct {
	SchemaVersion string         `yaml:"schemaVersion"`
	Version       string         `yaml:"version"`
	Entries       []CatalogEntry `yaml:"entries"`
}

func main() {
	entries, err := collectEntries()
	if err != nil {
		log.Fatal(err)
	}

	catalog := Catalog{
		SchemaVersion: "0.0.1",
		Version:       "0.0.1",
		Entries:       entries,
	}

	out, err := yaml.Marshal(catalog)
	if err != nil {
		log.Fatal(err)
	}
	out = []byte(catalogHeader + string(out))
	os.WriteFile(catalogOutFile, out, 0644)
}

func collectEntries() ([]CatalogEntry, error) {
	files, err := os.ReadDir(catalogDir)
	if err != nil {
		return nil, err
	}

	entries := []CatalogEntry{}
	for _, file := range files {
		if !file.IsDir() {
			continue
		}
		e, err := createEntry(file.Name())
		if err != nil {
			return nil, err
		}
		entries = append(entries, e)
	}
	return entries, nil
}

func createEntry(entryDir string) (CatalogEntry, error) {
	manifestPath := catalogDir + "/" + entryDir + "/manifest.yaml"
	m, err := os.ReadFile(manifestPath)
	if err != nil {
		return CatalogEntry{}, err
	}

	parsed := make(map[string]interface{})
	err = yaml.Unmarshal(m, &parsed)
	if err != nil {
		return CatalogEntry{}, err
	}

	version, ok := parsed["version"].(string)
	if !ok {
		return CatalogEntry{}, fmt.Errorf("version not found in manifest")
	}

	if version == "" {
		return CatalogEntry{}, fmt.Errorf("version is empty")
	}

	return CatalogEntry{
		Name:    entryDir,
		Version: version,
	}, nil
}
