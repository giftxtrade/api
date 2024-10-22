//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import "errors"

type CurrencyType string

const (
	CurrencyType_Usd CurrencyType = "USD"
	CurrencyType_Cad CurrencyType = "CAD"
)

func (e *CurrencyType) Scan(value interface{}) error {
	var enumValue string
	switch val := value.(type) {
	case string:
		enumValue = val
	case []byte:
		enumValue = string(val)
	default:
		return errors.New("jet: Invalid scan value for AllTypesEnum enum. Enum value has to be of type string or []byte")
	}

	switch enumValue {
	case "USD":
		*e = CurrencyType_Usd
	case "CAD":
		*e = CurrencyType_Cad
	default:
		return errors.New("jet: Invalid scan value '" + enumValue + "' for CurrencyType enum")
	}

	return nil
}

func (e CurrencyType) String() string {
	return string(e)
}