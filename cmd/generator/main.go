package main

import (
	"codex-documents/tools/generator"
	"log"
)

func main() {
	whitelist := []string{
		"Patient", "DocumentReference", "Binary", "Observation",
		"Bundle", "OperationOutcome", "Encounter", "Practitioner",
	}

	gen := generator.NewGenerator("fhir/spec", "fhir/models", whitelist)

	log.Println("Loading official FHIR StructureDefinitions...")

	if err := gen.Load("profiles-types.json"); err != nil {
		log.Fatal("Failed to load types:", err)
	}

	if err := gen.Load("profiles-resources.json"); err != nil {
		log.Fatal("Failed to load resources:", err)
	}

	log.Println("Generating clean Go models...")
	if err := gen.Generate(); err != nil {
		log.Fatal("Generation failed:", err)
	}

	log.Println("Done! Check 'fhir' directory.")
}
