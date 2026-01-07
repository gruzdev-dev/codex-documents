package models

import (
	"encoding/json"
	"fmt"
)

// The detailed description of a substance, typically at a level beyond what is used for prescribing.
type SubstanceDefinition struct {
	ResourceType     string                                `json:"resourceType" bson:"resource_type"`                            // Type of resource
	Id               *string                               `json:"id,omitempty" bson:"id,omitempty"`                             // Logical id of this artifact
	Meta             *Meta                                 `json:"meta,omitempty" bson:"meta,omitempty"`                         // Metadata about the resource
	ImplicitRules    *string                               `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`      // A set of rules under which this content was created
	Language         *string                               `json:"language,omitempty" bson:"language,omitempty"`                 // Language of the resource content
	Text             *Narrative                            `json:"text,omitempty" bson:"text,omitempty"`                         // Text summary of the resource, for human interpretation
	Contained        []json.RawMessage                     `json:"contained,omitempty" bson:"contained,omitempty"`               // Contained, inline Resources
	Identifier       []Identifier                          `json:"identifier,omitempty" bson:"identifier,omitempty"`             // Identifier by which this substance is known
	Version          *string                               `json:"version,omitempty" bson:"version,omitempty"`                   // A business level edition or revision identifier
	Status           *CodeableConcept                      `json:"status,omitempty" bson:"status,omitempty"`                     // Status of substance within the catalogue e.g. active, retired
	Classification   []CodeableConcept                     `json:"classification,omitempty" bson:"classification,omitempty"`     // A categorization, high level e.g. polymer or nucleic acid, or food, chemical, biological, or lower e.g. polymer linear or branch chain, or type of impurity
	Domain           *CodeableConcept                      `json:"domain,omitempty" bson:"domain,omitempty"`                     // The applicable usage of the substance, as an example human or veterinary
	Grade            []CodeableConcept                     `json:"grade,omitempty" bson:"grade,omitempty"`                       // The quality standard, established benchmark, to which substance complies (e.g. USP/NF, BP)
	Description      *string                               `json:"description,omitempty" bson:"description,omitempty"`           // Textual description of the substance
	Note             []Annotation                          `json:"note,omitempty" bson:"note,omitempty"`                         // Textual comment about the substance's catalogue or registry record
	Manufacturer     []Reference                           `json:"manufacturer,omitempty" bson:"manufacturer,omitempty"`         // The entity that creates, makes, produces or fabricates the substance
	Supplier         []Reference                           `json:"supplier,omitempty" bson:"supplier,omitempty"`                 // An entity that is the source for the substance. It may be different from the manufacturer
	Moiety           []SubstanceDefinitionMoiety           `json:"moiety,omitempty" bson:"moiety,omitempty"`                     // Moiety, for structural modifications
	Characterization []SubstanceDefinitionCharacterization `json:"characterization,omitempty" bson:"characterization,omitempty"` // General specifications for this substance
	Property         []SubstanceDefinitionProperty         `json:"property,omitempty" bson:"property,omitempty"`                 // General specifications for this substance
	MolecularWeight  []SubstanceDefinitionMolecularWeight  `json:"molecularWeight,omitempty" bson:"molecular_weight,omitempty"`  // The average mass of a molecule of a compound
	Structure        *SubstanceDefinitionStructure         `json:"structure,omitempty" bson:"structure,omitempty"`               // Structural information
	Code             []SubstanceDefinitionCode             `json:"code,omitempty" bson:"code,omitempty"`                         // Codes associated with the substance
	Name             []SubstanceDefinitionName             `json:"name,omitempty" bson:"name,omitempty"`                         // Names applicable to this substance
	Relationship     []SubstanceDefinitionRelationship     `json:"relationship,omitempty" bson:"relationship,omitempty"`         // A link between this substance and another
	SourceMaterial   *SubstanceDefinitionSourceMaterial    `json:"sourceMaterial,omitempty" bson:"source_material,omitempty"`    // Material or taxonomic/anatomical source
}

func (r *SubstanceDefinition) Validate() error {
	if r.ResourceType != "SubstanceDefinition" {
		return fmt.Errorf("invalid resourceType: expected 'SubstanceDefinition', got '%s'", r.ResourceType)
	}
	if r.Meta != nil {
		if err := r.Meta.Validate(); err != nil {
			return fmt.Errorf("Meta: %w", err)
		}
	}
	if r.Text != nil {
		if err := r.Text.Validate(); err != nil {
			return fmt.Errorf("Text: %w", err)
		}
	}
	for i, item := range r.Identifier {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Identifier[%d]: %w", i, err)
		}
	}
	if r.Status != nil {
		if err := r.Status.Validate(); err != nil {
			return fmt.Errorf("Status: %w", err)
		}
	}
	for i, item := range r.Classification {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Classification[%d]: %w", i, err)
		}
	}
	if r.Domain != nil {
		if err := r.Domain.Validate(); err != nil {
			return fmt.Errorf("Domain: %w", err)
		}
	}
	for i, item := range r.Grade {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Grade[%d]: %w", i, err)
		}
	}
	for i, item := range r.Note {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Note[%d]: %w", i, err)
		}
	}
	for i, item := range r.Manufacturer {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Manufacturer[%d]: %w", i, err)
		}
	}
	for i, item := range r.Supplier {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Supplier[%d]: %w", i, err)
		}
	}
	for i, item := range r.Moiety {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Moiety[%d]: %w", i, err)
		}
	}
	for i, item := range r.Characterization {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Characterization[%d]: %w", i, err)
		}
	}
	for i, item := range r.Property {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Property[%d]: %w", i, err)
		}
	}
	for i, item := range r.MolecularWeight {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("MolecularWeight[%d]: %w", i, err)
		}
	}
	if r.Structure != nil {
		if err := r.Structure.Validate(); err != nil {
			return fmt.Errorf("Structure: %w", err)
		}
	}
	for i, item := range r.Code {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Code[%d]: %w", i, err)
		}
	}
	for i, item := range r.Name {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Name[%d]: %w", i, err)
		}
	}
	for i, item := range r.Relationship {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Relationship[%d]: %w", i, err)
		}
	}
	if r.SourceMaterial != nil {
		if err := r.SourceMaterial.Validate(); err != nil {
			return fmt.Errorf("SourceMaterial: %w", err)
		}
	}
	return nil
}

type SubstanceDefinitionRelationship struct {
	Id                                 *string          `json:"id,omitempty" bson:"id,omitempty"`                                                                    // Unique id for inter-element referencing
	SubstanceDefinitionReference       *Reference       `json:"substanceDefinitionReference,omitempty" bson:"substance_definition_reference,omitempty"`              // A pointer to another substance, as a resource or a representational code
	SubstanceDefinitionCodeableConcept *CodeableConcept `json:"substanceDefinitionCodeableConcept,omitempty" bson:"substance_definition_codeable_concept,omitempty"` // A pointer to another substance, as a resource or a representational code
	Type                               *CodeableConcept `json:"type" bson:"type"`                                                                                    // For example "salt to parent", "active moiety"
	IsDefining                         bool             `json:"isDefining,omitempty" bson:"is_defining,omitempty"`                                                   // For example where an enzyme strongly bonds with a particular substance, this is a defining relationship for that enzyme, out of several possible relationships
	AmountQuantity                     *Quantity        `json:"amountQuantity,omitempty" bson:"amount_quantity,omitempty"`                                           // A numeric factor for the relationship, e.g. that a substance salt has some percentage of active substance in relation to some other
	AmountRatio                        *Ratio           `json:"amountRatio,omitempty" bson:"amount_ratio,omitempty"`                                                 // A numeric factor for the relationship, e.g. that a substance salt has some percentage of active substance in relation to some other
	AmountString                       *string          `json:"amountString,omitempty" bson:"amount_string,omitempty"`                                               // A numeric factor for the relationship, e.g. that a substance salt has some percentage of active substance in relation to some other
	RatioHighLimitAmount               *Ratio           `json:"ratioHighLimitAmount,omitempty" bson:"ratio_high_limit_amount,omitempty"`                             // For use when the numeric has an uncertain range
	Comparator                         *CodeableConcept `json:"comparator,omitempty" bson:"comparator,omitempty"`                                                    // An operator for the amount, for example "average", "approximately", "less than"
	Source                             []Reference      `json:"source,omitempty" bson:"source,omitempty"`                                                            // Supporting literature
}

func (r *SubstanceDefinitionRelationship) Validate() error {
	if r.SubstanceDefinitionReference != nil {
		if err := r.SubstanceDefinitionReference.Validate(); err != nil {
			return fmt.Errorf("SubstanceDefinitionReference: %w", err)
		}
	}
	if r.SubstanceDefinitionCodeableConcept != nil {
		if err := r.SubstanceDefinitionCodeableConcept.Validate(); err != nil {
			return fmt.Errorf("SubstanceDefinitionCodeableConcept: %w", err)
		}
	}
	if r.Type == nil {
		return fmt.Errorf("field 'Type' is required")
	}
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	if r.AmountQuantity != nil {
		if err := r.AmountQuantity.Validate(); err != nil {
			return fmt.Errorf("AmountQuantity: %w", err)
		}
	}
	if r.AmountRatio != nil {
		if err := r.AmountRatio.Validate(); err != nil {
			return fmt.Errorf("AmountRatio: %w", err)
		}
	}
	if r.RatioHighLimitAmount != nil {
		if err := r.RatioHighLimitAmount.Validate(); err != nil {
			return fmt.Errorf("RatioHighLimitAmount: %w", err)
		}
	}
	if r.Comparator != nil {
		if err := r.Comparator.Validate(); err != nil {
			return fmt.Errorf("Comparator: %w", err)
		}
	}
	for i, item := range r.Source {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Source[%d]: %w", i, err)
		}
	}
	return nil
}

type SubstanceDefinitionProperty struct {
	Id                   *string          `json:"id,omitempty" bson:"id,omitempty"`                                       // Unique id for inter-element referencing
	Type                 *CodeableConcept `json:"type" bson:"type"`                                                       // A code expressing the type of property
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept,omitempty" bson:"value_codeable_concept,omitempty"` // A value for the property
	ValueQuantity        *Quantity        `json:"valueQuantity,omitempty" bson:"value_quantity,omitempty"`                // A value for the property
	ValueRange           *Range           `json:"valueRange,omitempty" bson:"value_range,omitempty"`                      // A value for the property
	ValueDate            *string          `json:"valueDate,omitempty" bson:"value_date,omitempty"`                        // A value for the property
	ValueBoolean         *bool            `json:"valueBoolean,omitempty" bson:"value_boolean,omitempty"`                  // A value for the property
	ValueAttachment      *Attachment      `json:"valueAttachment,omitempty" bson:"value_attachment,omitempty"`            // A value for the property
}

func (r *SubstanceDefinitionProperty) Validate() error {
	if r.Type == nil {
		return fmt.Errorf("field 'Type' is required")
	}
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	if r.ValueCodeableConcept != nil {
		if err := r.ValueCodeableConcept.Validate(); err != nil {
			return fmt.Errorf("ValueCodeableConcept: %w", err)
		}
	}
	if r.ValueQuantity != nil {
		if err := r.ValueQuantity.Validate(); err != nil {
			return fmt.Errorf("ValueQuantity: %w", err)
		}
	}
	if r.ValueRange != nil {
		if err := r.ValueRange.Validate(); err != nil {
			return fmt.Errorf("ValueRange: %w", err)
		}
	}
	if r.ValueAttachment != nil {
		if err := r.ValueAttachment.Validate(); err != nil {
			return fmt.Errorf("ValueAttachment: %w", err)
		}
	}
	return nil
}

type SubstanceDefinitionStructureRepresentation struct {
	Id             *string          `json:"id,omitempty" bson:"id,omitempty"`                         // Unique id for inter-element referencing
	Type           *CodeableConcept `json:"type,omitempty" bson:"type,omitempty"`                     // The kind of structural representation (e.g. full, partial)
	Representation *string          `json:"representation,omitempty" bson:"representation,omitempty"` // The structural representation as a text string in a standard format
	Format         *CodeableConcept `json:"format,omitempty" bson:"format,omitempty"`                 // The format of the representation e.g. InChI, SMILES, MOLFILE (note: not the physical file format)
	Document       *Reference       `json:"document,omitempty" bson:"document,omitempty"`             // An attachment with the structural representation e.g. a structure graphic or AnIML file
}

func (r *SubstanceDefinitionStructureRepresentation) Validate() error {
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	if r.Format != nil {
		if err := r.Format.Validate(); err != nil {
			return fmt.Errorf("Format: %w", err)
		}
	}
	if r.Document != nil {
		if err := r.Document.Validate(); err != nil {
			return fmt.Errorf("Document: %w", err)
		}
	}
	return nil
}

type SubstanceDefinitionName struct {
	Id           *string                           `json:"id,omitempty" bson:"id,omitempty"`                     // Unique id for inter-element referencing
	Name         string                            `json:"name" bson:"name"`                                     // The actual name
	Type         *CodeableConcept                  `json:"type,omitempty" bson:"type,omitempty"`                 // Name type e.g. 'systematic',  'scientific, 'brand'
	Status       *CodeableConcept                  `json:"status,omitempty" bson:"status,omitempty"`             // The status of the name e.g. 'current', 'proposed'
	Preferred    bool                              `json:"preferred,omitempty" bson:"preferred,omitempty"`       // If this is the preferred name for this substance
	Language     []CodeableConcept                 `json:"language,omitempty" bson:"language,omitempty"`         // Human language that the name is written in
	Domain       []CodeableConcept                 `json:"domain,omitempty" bson:"domain,omitempty"`             // The use context of this name e.g. as an active ingredient or as a food colour additive
	Jurisdiction []CodeableConcept                 `json:"jurisdiction,omitempty" bson:"jurisdiction,omitempty"` // The jurisdiction where this name applies
	Synonym      []SubstanceDefinitionName         `json:"synonym,omitempty" bson:"synonym,omitempty"`           // A synonym of this particular name, by which the substance is also known
	Translation  []SubstanceDefinitionName         `json:"translation,omitempty" bson:"translation,omitempty"`   // A translation for this name into another human language
	Official     []SubstanceDefinitionNameOfficial `json:"official,omitempty" bson:"official,omitempty"`         // Details of the official nature of this name
	Source       []Reference                       `json:"source,omitempty" bson:"source,omitempty"`             // Supporting literature
}

func (r *SubstanceDefinitionName) Validate() error {
	var emptyString string
	if r.Name == emptyString {
		return fmt.Errorf("field 'Name' is required")
	}
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	if r.Status != nil {
		if err := r.Status.Validate(); err != nil {
			return fmt.Errorf("Status: %w", err)
		}
	}
	for i, item := range r.Language {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Language[%d]: %w", i, err)
		}
	}
	for i, item := range r.Domain {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Domain[%d]: %w", i, err)
		}
	}
	for i, item := range r.Jurisdiction {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Jurisdiction[%d]: %w", i, err)
		}
	}
	for i, item := range r.Synonym {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Synonym[%d]: %w", i, err)
		}
	}
	for i, item := range r.Translation {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Translation[%d]: %w", i, err)
		}
	}
	for i, item := range r.Official {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Official[%d]: %w", i, err)
		}
	}
	for i, item := range r.Source {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Source[%d]: %w", i, err)
		}
	}
	return nil
}

type SubstanceDefinitionNameOfficial struct {
	Id        *string          `json:"id,omitempty" bson:"id,omitempty"`               // Unique id for inter-element referencing
	Authority *CodeableConcept `json:"authority,omitempty" bson:"authority,omitempty"` // Which authority uses this official name
	Status    *CodeableConcept `json:"status,omitempty" bson:"status,omitempty"`       // The status of the official name, for example 'draft', 'active'
	Date      *string          `json:"date,omitempty" bson:"date,omitempty"`           // Date of official name change
}

func (r *SubstanceDefinitionNameOfficial) Validate() error {
	if r.Authority != nil {
		if err := r.Authority.Validate(); err != nil {
			return fmt.Errorf("Authority: %w", err)
		}
	}
	if r.Status != nil {
		if err := r.Status.Validate(); err != nil {
			return fmt.Errorf("Status: %w", err)
		}
	}
	return nil
}

type SubstanceDefinitionSourceMaterial struct {
	Id              *string           `json:"id,omitempty" bson:"id,omitempty"`                             // Unique id for inter-element referencing
	Type            *CodeableConcept  `json:"type,omitempty" bson:"type,omitempty"`                         // Classification of the origin of the raw material. e.g. cat hair is an Animal source type
	Genus           *CodeableConcept  `json:"genus,omitempty" bson:"genus,omitempty"`                       // The genus of an organism e.g. the Latin epithet of the plant/animal scientific name
	Species         *CodeableConcept  `json:"species,omitempty" bson:"species,omitempty"`                   // The species of an organism e.g. the Latin epithet of the species of the plant/animal
	Part            *CodeableConcept  `json:"part,omitempty" bson:"part,omitempty"`                         // An anatomical origin of the source material within an organism
	CountryOfOrigin []CodeableConcept `json:"countryOfOrigin,omitempty" bson:"country_of_origin,omitempty"` // The country or countries where the material is harvested
}

func (r *SubstanceDefinitionSourceMaterial) Validate() error {
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	if r.Genus != nil {
		if err := r.Genus.Validate(); err != nil {
			return fmt.Errorf("Genus: %w", err)
		}
	}
	if r.Species != nil {
		if err := r.Species.Validate(); err != nil {
			return fmt.Errorf("Species: %w", err)
		}
	}
	if r.Part != nil {
		if err := r.Part.Validate(); err != nil {
			return fmt.Errorf("Part: %w", err)
		}
	}
	for i, item := range r.CountryOfOrigin {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("CountryOfOrigin[%d]: %w", i, err)
		}
	}
	return nil
}

type SubstanceDefinitionMoiety struct {
	Id               *string          `json:"id,omitempty" bson:"id,omitempty"`                              // Unique id for inter-element referencing
	Role             *CodeableConcept `json:"role,omitempty" bson:"role,omitempty"`                          // Role that the moiety is playing
	Identifier       *Identifier      `json:"identifier,omitempty" bson:"identifier,omitempty"`              // Identifier by which this moiety substance is known
	Name             *string          `json:"name,omitempty" bson:"name,omitempty"`                          // Textual name for this moiety substance
	Stereochemistry  *CodeableConcept `json:"stereochemistry,omitempty" bson:"stereochemistry,omitempty"`    // Stereochemistry type
	OpticalActivity  *CodeableConcept `json:"opticalActivity,omitempty" bson:"optical_activity,omitempty"`   // Optical activity type
	MolecularFormula *string          `json:"molecularFormula,omitempty" bson:"molecular_formula,omitempty"` // Molecular formula for this moiety (e.g. with the Hill system)
	AmountQuantity   *Quantity        `json:"amountQuantity,omitempty" bson:"amount_quantity,omitempty"`     // Quantitative value for this moiety
	AmountString     *string          `json:"amountString,omitempty" bson:"amount_string,omitempty"`         // Quantitative value for this moiety
	MeasurementType  *CodeableConcept `json:"measurementType,omitempty" bson:"measurement_type,omitempty"`   // The measurement type of the quantitative value
}

func (r *SubstanceDefinitionMoiety) Validate() error {
	if r.Role != nil {
		if err := r.Role.Validate(); err != nil {
			return fmt.Errorf("Role: %w", err)
		}
	}
	if r.Identifier != nil {
		if err := r.Identifier.Validate(); err != nil {
			return fmt.Errorf("Identifier: %w", err)
		}
	}
	if r.Stereochemistry != nil {
		if err := r.Stereochemistry.Validate(); err != nil {
			return fmt.Errorf("Stereochemistry: %w", err)
		}
	}
	if r.OpticalActivity != nil {
		if err := r.OpticalActivity.Validate(); err != nil {
			return fmt.Errorf("OpticalActivity: %w", err)
		}
	}
	if r.AmountQuantity != nil {
		if err := r.AmountQuantity.Validate(); err != nil {
			return fmt.Errorf("AmountQuantity: %w", err)
		}
	}
	if r.MeasurementType != nil {
		if err := r.MeasurementType.Validate(); err != nil {
			return fmt.Errorf("MeasurementType: %w", err)
		}
	}
	return nil
}

type SubstanceDefinitionCharacterization struct {
	Id          *string          `json:"id,omitempty" bson:"id,omitempty"`                   // Unique id for inter-element referencing
	Technique   *CodeableConcept `json:"technique,omitempty" bson:"technique,omitempty"`     // The method used to find the characterization e.g. HPLC
	Form        *CodeableConcept `json:"form,omitempty" bson:"form,omitempty"`               // Describes the nature of the chemical entity and explains, for instance, whether this is a base or a salt form
	Description *string          `json:"description,omitempty" bson:"description,omitempty"` // The description or justification in support of the interpretation of the data file
	File        []Attachment     `json:"file,omitempty" bson:"file,omitempty"`               // The data produced by the analytical instrument or a pictorial representation of that data. Examples: a JCAMP, JDX, or ADX file, or a chromatogram or spectrum analysis
}

func (r *SubstanceDefinitionCharacterization) Validate() error {
	if r.Technique != nil {
		if err := r.Technique.Validate(); err != nil {
			return fmt.Errorf("Technique: %w", err)
		}
	}
	if r.Form != nil {
		if err := r.Form.Validate(); err != nil {
			return fmt.Errorf("Form: %w", err)
		}
	}
	for i, item := range r.File {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("File[%d]: %w", i, err)
		}
	}
	return nil
}

type SubstanceDefinitionMolecularWeight struct {
	Id     *string          `json:"id,omitempty" bson:"id,omitempty"`         // Unique id for inter-element referencing
	Method *CodeableConcept `json:"method,omitempty" bson:"method,omitempty"` // The method by which the weight was determined
	Type   *CodeableConcept `json:"type,omitempty" bson:"type,omitempty"`     // Type of molecular weight e.g. exact, average, weight average
	Amount *Quantity        `json:"amount" bson:"amount"`                     // Used to capture quantitative values for a variety of elements
}

func (r *SubstanceDefinitionMolecularWeight) Validate() error {
	if r.Method != nil {
		if err := r.Method.Validate(); err != nil {
			return fmt.Errorf("Method: %w", err)
		}
	}
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	if r.Amount == nil {
		return fmt.Errorf("field 'Amount' is required")
	}
	if r.Amount != nil {
		if err := r.Amount.Validate(); err != nil {
			return fmt.Errorf("Amount: %w", err)
		}
	}
	return nil
}

type SubstanceDefinitionStructure struct {
	Id                       *string                                      `json:"id,omitempty" bson:"id,omitempty"`                                                // Unique id for inter-element referencing
	Stereochemistry          *CodeableConcept                             `json:"stereochemistry,omitempty" bson:"stereochemistry,omitempty"`                      // Stereochemistry type
	OpticalActivity          *CodeableConcept                             `json:"opticalActivity,omitempty" bson:"optical_activity,omitempty"`                     // Optical activity type
	MolecularFormula         *string                                      `json:"molecularFormula,omitempty" bson:"molecular_formula,omitempty"`                   // An expression which states the number and type of atoms present in a molecule of a substance
	MolecularFormulaByMoiety *string                                      `json:"molecularFormulaByMoiety,omitempty" bson:"molecular_formula_by_moiety,omitempty"` // Specified per moiety according to the Hill system
	MolecularWeight          *SubstanceDefinitionMolecularWeight          `json:"molecularWeight,omitempty" bson:"molecular_weight,omitempty"`                     // The molecular weight or weight range
	Technique                []CodeableConcept                            `json:"technique,omitempty" bson:"technique,omitempty"`                                  // The method used to find the structure e.g. X-ray, NMR
	SourceDocument           []Reference                                  `json:"sourceDocument,omitempty" bson:"source_document,omitempty"`                       // Source of information for the structure
	Representation           []SubstanceDefinitionStructureRepresentation `json:"representation,omitempty" bson:"representation,omitempty"`                        // A depiction of the structure of the substance
}

func (r *SubstanceDefinitionStructure) Validate() error {
	if r.Stereochemistry != nil {
		if err := r.Stereochemistry.Validate(); err != nil {
			return fmt.Errorf("Stereochemistry: %w", err)
		}
	}
	if r.OpticalActivity != nil {
		if err := r.OpticalActivity.Validate(); err != nil {
			return fmt.Errorf("OpticalActivity: %w", err)
		}
	}
	if r.MolecularWeight != nil {
		if err := r.MolecularWeight.Validate(); err != nil {
			return fmt.Errorf("MolecularWeight: %w", err)
		}
	}
	for i, item := range r.Technique {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Technique[%d]: %w", i, err)
		}
	}
	for i, item := range r.SourceDocument {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("SourceDocument[%d]: %w", i, err)
		}
	}
	for i, item := range r.Representation {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Representation[%d]: %w", i, err)
		}
	}
	return nil
}

type SubstanceDefinitionCode struct {
	Id         *string          `json:"id,omitempty" bson:"id,omitempty"`                  // Unique id for inter-element referencing
	Code       *CodeableConcept `json:"code,omitempty" bson:"code,omitempty"`              // The specific code
	Status     *CodeableConcept `json:"status,omitempty" bson:"status,omitempty"`          // Status of the code assignment, for example 'provisional', 'approved'
	StatusDate *string          `json:"statusDate,omitempty" bson:"status_date,omitempty"` // The date at which the code status was changed
	Note       []Annotation     `json:"note,omitempty" bson:"note,omitempty"`              // Any comment can be provided in this field
	Source     []Reference      `json:"source,omitempty" bson:"source,omitempty"`          // Supporting literature
}

func (r *SubstanceDefinitionCode) Validate() error {
	if r.Code != nil {
		if err := r.Code.Validate(); err != nil {
			return fmt.Errorf("Code: %w", err)
		}
	}
	if r.Status != nil {
		if err := r.Status.Validate(); err != nil {
			return fmt.Errorf("Status: %w", err)
		}
	}
	for i, item := range r.Note {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Note[%d]: %w", i, err)
		}
	}
	for i, item := range r.Source {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Source[%d]: %w", i, err)
		}
	}
	return nil
}
