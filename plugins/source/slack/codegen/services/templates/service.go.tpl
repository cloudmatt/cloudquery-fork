// Code generated by codegen; DO NOT EDIT.
package services

import (
    "context"
    "io"
    "{{ .Import }}"
)

//go:generate mockgen -package=mocks -destination=../mocks/{{.PackageName}}.go -source={{.PackageName}}.go {{.ClientName}}
type {{.ClientName}} interface {
    {{- range $sig := .Signatures }}
    {{ $sig }}
    {{- end }}
}