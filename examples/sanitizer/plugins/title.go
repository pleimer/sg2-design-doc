package main

import (
    "strings"
)

type title struct {}

func (t title) Sanitize(raw string) string {
    return strings.Title(strings.ToLower(raw))
}

var Title title 
