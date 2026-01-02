FHIR_DIR = fhir
SCHEMA = $(FHIR_DIR)/fhir.schema.json
OUTPUT = $(FHIR_DIR)/models.go
CORE_DIR = core/domain/fhir
MAPPERS_DIR = adapters/fhir

.PHONY: generate generate-fhir generate-core generate-mappers clean-fhir clean-core clear-mappers clean

generate: clean generate-fhir generate-core generate-mappers

generate-fhir:
	go-jsonschema -p fhir -o $(OUTPUT) --tags json --capitalization ID $(SCHEMA)

generate-core:
	go run fhir/generator.go

generate-mappers:
	goverter gen ./adapters/fhir

clean-fhir:
	rm -f $(OUTPUT)

clean-core:
	rm -rf $(CORE_DIR)/*.go

clean-mappers:
	rm -rf $(MAPPERS_DIR)/*.go

clean: clean-fhir clean-core clean-mappers