package indexeddb

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"errors"

	"github.com/seddonm1/cdproto/runtime"
	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// DatabaseWithObjectStores database with an array of object stores.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB#type-DatabaseWithObjectStores
type DatabaseWithObjectStores struct {
	Name         string         `json:"name"`         // Database name.
	Version      float64        `json:"version"`      // Database version (type is not 'integer', as the standard requires the version number to be 'unsigned long long')
	ObjectStores []*ObjectStore `json:"objectStores"` // Object stores in this database.
}

// ObjectStore object store.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB#type-ObjectStore
type ObjectStore struct {
	Name          string              `json:"name"`          // Object store name.
	KeyPath       *KeyPath            `json:"keyPath"`       // Object store key path.
	AutoIncrement bool                `json:"autoIncrement"` // If true, object store has auto increment flag set.
	Indexes       []*ObjectStoreIndex `json:"indexes"`       // Indexes in this object store.
}

// ObjectStoreIndex object store index.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB#type-ObjectStoreIndex
type ObjectStoreIndex struct {
	Name       string   `json:"name"`       // Index name.
	KeyPath    *KeyPath `json:"keyPath"`    // Index key path.
	Unique     bool     `json:"unique"`     // If true, index is unique.
	MultiEntry bool     `json:"multiEntry"` // If true, index allows multiple entries for a key.
}

// Key Key.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB#type-Key
type Key struct {
	Type   KeyType `json:"type"`             // Key type.
	Number float64 `json:"number,omitempty"` // Number value.
	String string  `json:"string,omitempty"` // String value.
	Date   float64 `json:"date,omitempty"`   // Date value.
	Array  []*Key  `json:"array,omitempty"`  // Array value.
}

// KeyRange key range.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB#type-KeyRange
type KeyRange struct {
	Lower     *Key `json:"lower,omitempty"` // Lower bound.
	Upper     *Key `json:"upper,omitempty"` // Upper bound.
	LowerOpen bool `json:"lowerOpen"`       // If true lower bound is open.
	UpperOpen bool `json:"upperOpen"`       // If true upper bound is open.
}

// DataEntry data entry.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB#type-DataEntry
type DataEntry struct {
	Key        *runtime.RemoteObject `json:"key"`        // Key object.
	PrimaryKey *runtime.RemoteObject `json:"primaryKey"` // Primary key object.
	Value      *runtime.RemoteObject `json:"value"`      // Value object.
}

// KeyPath key path.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB#type-KeyPath
type KeyPath struct {
	Type   KeyPathType `json:"type"`             // Key path type.
	String string      `json:"string,omitempty"` // String value.
	Array  []string    `json:"array,omitempty"`  // Array value.
}

// KeyType key type.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB#type-type
type KeyType string

// String returns the KeyType as string value.
func (t KeyType) String() string {
	return string(t)
}

// KeyType values.
const (
	KeyTypeNumber KeyType = "number"
	KeyTypeString KeyType = "string"
	KeyTypeDate   KeyType = "date"
	KeyTypeArray  KeyType = "array"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t KeyType) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t KeyType) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *KeyType) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch KeyType(in.String()) {
	case KeyTypeNumber:
		*t = KeyTypeNumber
	case KeyTypeString:
		*t = KeyTypeString
	case KeyTypeDate:
		*t = KeyTypeDate
	case KeyTypeArray:
		*t = KeyTypeArray

	default:
		in.AddError(errors.New("unknown KeyType value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *KeyType) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// KeyPathType key path type.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/IndexedDB#type-type
type KeyPathType string

// String returns the KeyPathType as string value.
func (t KeyPathType) String() string {
	return string(t)
}

// KeyPathType values.
const (
	KeyPathTypeNull   KeyPathType = "null"
	KeyPathTypeString KeyPathType = "string"
	KeyPathTypeArray  KeyPathType = "array"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t KeyPathType) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t KeyPathType) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *KeyPathType) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch KeyPathType(in.String()) {
	case KeyPathTypeNull:
		*t = KeyPathTypeNull
	case KeyPathTypeString:
		*t = KeyPathTypeString
	case KeyPathTypeArray:
		*t = KeyPathTypeArray

	default:
		in.AddError(errors.New("unknown KeyPathType value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *KeyPathType) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}
