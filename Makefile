FHIR_DIR = fhir
SCHEMA = $(FHIR_DIR)/fhir.schema.json
OUTPUT = $(FHIR_DIR)/models.go
CORE_DIR = core/domain/fhir

.PHONY: generate-fhir generate-core generate clean-fhir clean-core clean

generate: clean generate-fhir generate-core

generate-fhir:
	go-jsonschema -p fhir -o $(OUTPUT) --tags json --capitalization ID $(SCHEMA)

generate-core:
	go run fhir/gen_domain.go

clean-fhir:
	rm -f $(OUTPUT)

clean-core:
	rm -rf $(CORE_DIR)/*.go

clean: clean-fhir clean-core