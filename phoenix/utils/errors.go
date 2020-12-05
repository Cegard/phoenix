package utils

import (
    "fmt"
)


type NotNumberError struct {
    originalValue string
}


func NewNotNumberError (value string) *NotNumberError {
    
    return &NotNumberError {
        originalValue: value,
    }
}


func (err *NotNumberError) Error() string {
    
    return fmt.Sprintf("\nValue: %s, Cannot be converted to an integer\n", err.originalValue)
}