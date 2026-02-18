package main

import (
	"encoding/json"
	"fmt"
)

type Span struct {
	ID        string            `json:"id"`
	Name      string            `json:"name"`
	Duration  int               `json:"duration"`
	Timestamp int64             `json:"timestamp"`
	Tags      map[string]string `json:"tags"`

	// Stores all unknown JSON fields
	Extra map[string]json.RawMessage `json:"-"`
}

func (s *Span) UnmarshalJSON(data []byte) error {
	type Alias Span // This is because we need to define a type that doesn't define UnmarshalJSON. Otherwise if we do json.Unmarshal() it will call this method that will call json.Unmarshal() that will call.... oops infinite recursion
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(s),
	}

	// Let's the standard JSON decoder read ID, Name, ...
	// As we pass Aux and Aux is backed by Span the result is directly written into
	// span struct. But aliased. So fields of Span are updated and others are ignored.
	if err := json.Unmarshal(data, &aux); err != nil {
		return fmt.Errorf("failed to unmarshal known fields: %w", err)
	}

	// Now decode again using raw message
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return fmt.Errorf("failed to unmarshal raw map: %w", err)
	}

	// Remove known fields
	delete(raw, "id")
	delete(raw, "name")
	delete(raw, "duration")
	delete(raw, "timestamp")
	delete(raw, "tags")

	if len(raw) > 0 {
		s.Extra = raw
	}

	return nil
}

func (s Span) Print() {
	fmt.Printf("ID       : %s\n", s.ID)
	fmt.Printf("Name     : %s\n", s.Name)
	fmt.Printf("Duration : %d\n", s.Duration)
	fmt.Printf("Timestamp: %d\n", s.Timestamp)
	fmt.Printf("Tags\n")
	for key, raw := range s.Tags {
		fmt.Printf("    %-20s: %s\n", key, string(raw))
	}

	if len(s.Extra) == 0 {
		fmt.Println("No extra fields")
	} else {
		fmt.Println("Extra fields:")
		for key, raw := range s.Extra {
			fmt.Printf("  %s: %s\n", key, string(raw))
		}
	}
}
