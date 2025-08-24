package main

import (
	"encoding/json/jsontext"
	"encoding/json/v2"
	"fmt"
	"log"
	"math"
	"time"
)

func main() {
	value := struct {
		BytesBase64     []byte         `json:",format:base64"`
		BytesHex        [8]byte        `json:",format:hex"`
		BytesArray      []byte         `json:",format:array"`
		FloatNonFinite  float64        `json:",format:nonfinite"`
		MapEmitNull     map[string]any `json:",format:emitnull"`
		SliceEmitNull   []any          `json:",format:emitnull"`
		TimeDateOnly    time.Time      `json:",format:'2006-01-02'"`
		TimeUnixSec     time.Time      `json:",format:unix"`
		DurationSecs    time.Duration  `json:",format:sec"`
		DurationNanos   time.Duration  `json:",format:nano"`
		DurationISO8601 time.Duration  `json:",format:iso8601"`
	}{
		BytesBase64:     []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef},
		BytesHex:        [8]byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef},
		BytesArray:      []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef},
		FloatNonFinite:  math.NaN(),
		MapEmitNull:     nil,
		SliceEmitNull:   nil,
		TimeDateOnly:    time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
		TimeUnixSec:     time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
		DurationSecs:    12*time.Hour + 34*time.Minute + 56*time.Second + 7*time.Millisecond + 8*time.Microsecond + 9*time.Nanosecond,
		DurationNanos:   12*time.Hour + 34*time.Minute + 56*time.Second + 7*time.Millisecond + 8*time.Microsecond + 9*time.Nanosecond,
		DurationISO8601: 12*time.Hour + 34*time.Minute + 56*time.Second + 7*time.Millisecond + 8*time.Microsecond + 9*time.Nanosecond,
	}

	b, err := json.Marshal(&value)
	if err != nil {
		log.Fatal(err)
	}
	(*jsontext.Value)(&b).Indent() // indent for readability
	fmt.Println(string(b))

	// CaseSensitivity
	// JSON input using various naming conventions.
	const input = `[
		{"firstname": true},
		{"firstName": true},
		{"FirstName": true},
		{"FIRSTNAME": true},
		{"first_name": true},
		{"FIRST_NAME": true},
		{"first-name": true},
		{"FIRST-NAME": true},
		{"unknown": true}
	]`

	// Without "case:ignore", Unmarshal looks for an exact match.
	var caseStrict []struct {
		X bool `json:"firstName"`
	}
	if err := json.Unmarshal([]byte(input), &caseStrict); err != nil {
		log.Fatal(err)
	}
	fmt.Println(caseStrict) // exactly 1 match found

	// With "case:ignore", Unmarshal looks first for an exact match,
	// then for a case-insensitive match if none found.
	var caseIgnore []struct {
		X bool `json:"firstName,case:ignore"`
	}
	if err := json.Unmarshal([]byte(input), &caseIgnore); err != nil {
		log.Fatal(err)
	}
	fmt.Println(caseIgnore) // 8 matches found

	// FieldNames
	var fieldsamplevalue struct {
		// This field is explicitly ignored with the special "-" name.
		Ignored any `json:"-"`
		// No JSON name is not provided, so the Go field name is used.
		GoName any
		// A JSON name is provided without any special characters.
		JSONName any `json:"jsonName"`
		// No JSON name is not provided, so the Go field name is used.
		Option any `json:",case:ignore"`
		// An empty JSON name specified using an single-quoted string literal.
		Empty any `json:"''"`
		// A dash JSON name specified using an single-quoted string literal.
		Dash any `json:"'-'"`
		// A comma JSON name specified using an single-quoted string literal.
		Comma any `json:"','"`
		// JSON name with quotes specified using a single-quoted string literal.
		Quote any `json:"'\"\\''"`
		// An unexported field is always ignored.
		unexported any
	}

	b, err = json.Marshal(fieldsamplevalue)
	if err != nil {
		log.Fatal(err)
	}
	(*jsontext.Value)(&b).Indent() // indent for readability
	fmt.Println(string(b))

}
