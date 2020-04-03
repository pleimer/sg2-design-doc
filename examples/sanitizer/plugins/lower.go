package main

import (
    "strings"
)

type lower struct{}

func (l lower) Sanitize(raw string) string {
    return strings.ToLower(raw)
}

var Lower lower
