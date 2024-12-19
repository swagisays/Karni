package liberrors

import (
	"fmt"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// HandleMongoValidationError handles MongoDB validation errors with detailed parsing.
func HandleMongoValidationError(err error) *KarniError {
	// Handle duplicate key error (E11000)
	if mongo.IsDuplicateKeyError(err) {
		return WrapError(1002, parseDuplicateKeyError(err), err)
	}

	// Check if the error is a WriteException
	writeErr, ok := err.(mongo.WriteException)
	if !ok {
		return WrapError(1003, "An unexpected error occurred.", err)
	}

	for _, we := range writeErr.WriteErrors {
		if we.Code == 121 {
			return WrapError(1004, parseBsonRawError(we.Details), err)
		}
	}

	return WrapError(1005, "A database error occurred.", err)
}

// parseBsonRawError extracts information from bson.Raw details.
func parseBsonRawError(details bson.Raw) string {
	var parsed struct {
		Details struct {
			SchemaRulesNotSatisfied []struct {
				OperatorName           string   `bson:"operatorName"`
				MissingProperties      []string `bson:"missingProperties"`
				PropertiesNotSatisfied []struct {
					PropertyName string `bson:"propertyName"`
					Details      []struct {
						Reason         string `bson:"reason"`
						ConsideredType string `bson:"consideredType"`
						SpecifiedAs    struct {
							BsonType string `bson:"bsonType"`
						} `bson:"specifiedAs"`
						ConsideredValue interface{} `bson:"consideredValue"`
					} `bson:"details"`
				} `bson:"propertiesNotSatisfied"`
			} `bson:"schemaRulesNotSatisfied"`
		} `bson:"details"`
	}

	err := bson.Unmarshal(details, &parsed)
	if err != nil {
		return "Error parsing validation details."
	}

	// Handle missing required fields
	for _, rule := range parsed.Details.SchemaRulesNotSatisfied {
		if rule.OperatorName == "required" && len(rule.MissingProperties) > 0 {
			return fmt.Sprintf(
				"Validation Error: Missing required field(s): %v.",
				rule.MissingProperties,
			)
		}

		// Handle type mismatch errors
		for _, prop := range rule.PropertiesNotSatisfied {
			for _, detail := range prop.Details {
				if detail.Reason == "type did not match" {
					return fmt.Sprintf(
						"Validation Error: Field '%s' has a type mismatch. Expected type '%s', but got value '%v' of type '%s'.",
						prop.PropertyName, detail.SpecifiedAs.BsonType, detail.ConsideredValue, detail.ConsideredType,
					)
				}
			}
		}
	}

	return "Validation Error: Invalid data format."
}

// parseDuplicateKeyError extracts and returns details from a duplicate key error message.
func parseDuplicateKeyError(err error) string {
	if strings.Contains(err.Error(), "duplicate key error") {
		// Extract the field name and value from the error message
		start := strings.Index(err.Error(), "{")
		end := strings.LastIndex(err.Error(), "}")
		if start != -1 && end != -1 {
			dupKeyDetails := err.Error()[start : end+1]
			return fmt.Sprintf("Duplicate Key Error: %s already exists.", dupKeyDetails)
		}
		return "Duplicate Key Error: A unique field already exists in the database."
	}
	return "Duplicate Key Error."
}
