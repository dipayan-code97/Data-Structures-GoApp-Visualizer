package containers

// JSONSerializer provides JSON serialization
type JSONSerializer interface {
	// ToJSON outputs the JSON representation of containers's elements.
	ToJSON() ([]byte, error)
}

// JSONDeserializer provides JSON deserialization
type JSONDeserializer interface {
	// FromJSON populates containers's elements from the input JSON representation.
	FromJSON([]byte) error
}
