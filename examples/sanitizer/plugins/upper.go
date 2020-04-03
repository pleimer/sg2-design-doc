package main

import (
    "strings"
)

type upper struct{}

func (u upper) Sanitize(raw string) string {
    return strings.ToUpper(raw)
}

var Upper upper
