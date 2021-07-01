package csharp

const fileTpl = `// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: {{ .File.InputPath }}
using {{ csharpPackage .File }};
using FluentValidation;

namespace {{ csharpPackage .File }}
{
{{- range .AllMessages }}
{{- template "msg" . }}
{{- end }}
}
`