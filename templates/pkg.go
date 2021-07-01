package templates

import (
	"text/template"

	"github.com/envoyproxy/protoc-gen-validate/templates/cc"
	"github.com/envoyproxy/protoc-gen-validate/templates/csharp"
	golang "github.com/envoyproxy/protoc-gen-validate/templates/go"
	"github.com/envoyproxy/protoc-gen-validate/templates/java"
	"github.com/envoyproxy/protoc-gen-validate/templates/shared"
	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
)

type RegisterFn func(tpl *template.Template, params pgs.Parameters)
type FilePathFn func(f pgs.File, ctx pgsgo.Context, tpl *template.Template) *pgs.FilePath

func makeTemplate(ext string, fn RegisterFn, params pgs.Parameters) *template.Template {
	tpl := template.New(ext)
	shared.RegisterFunctions(tpl, params)
	fn(tpl, params)
	return tpl
}

func Template(params pgs.Parameters) map[string][]*template.Template {
	return map[string][]*template.Template{
		"cc":     {makeTemplate("h", cc.RegisterHeader, params), makeTemplate("cc", cc.RegisterModule, params)},
		"go":     {makeTemplate("go", golang.Register, params)},
		"java":   {makeTemplate("java", java.Register, params)},
		"csharp": {makeTemplate("csharp", csharp.Register, params)},
	}
}

func FilePathFor(tpl *template.Template) FilePathFn {
	switch tpl.Name() {
	case "h":
		return cc.CcFilePath
	case "cc":
		return cc.CcFilePath
	case "java":
		return java.JavaFilePath
	case "csharp":
		return csharp.FilePath
	default:
		return func(f pgs.File, ctx pgsgo.Context, tpl *template.Template) *pgs.FilePath {
			out := ctx.OutputPath(f)
			out = out.SetExt(".validate." + tpl.Name())
			return &out
		}
	}
}
