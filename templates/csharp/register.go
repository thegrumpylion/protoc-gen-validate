package csharp

import (
	"strings"
	"text/template"

	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
)

func Register(tpl *template.Template, params pgs.Parameters) {
	fns := csharpFuncs{pgsgo.InitContext(params)}

	tpl.Funcs(map[string]interface{}{
		"msgTyp":        fns.msgTyp,
		"pkg":           fns.PackageName,
		"csharpPackage": csharpPackage,
	})
	template.Must(tpl.Parse(fileTpl))
	template.Must(tpl.New("msg").Parse(msgTpl))
}

type csharpFuncs struct {
	pgsgo.Context
}

func (fns csharpFuncs) msgTyp(message pgs.Message) pgsgo.TypeName {
	return pgsgo.TypeName(fns.Name(message))
}

func FilePath(f pgs.File, ctx pgsgo.Context, tpl *template.Template) *pgs.FilePath {
	out := ctx.OutputPath(f)
	out = out.SetExt(".validate.cs")
	return &out
}

func csharpPackage(file pgs.File) string {
	options := file.Descriptor().GetOptions()
	if options != nil && options.CsharpNamespace != nil {
		return options.GetCsharpNamespace()
	}
	parts := []string{}
	for _, s := range strings.Split(file.Package().ProtoName().String(), ".") {
		parts = append(parts, strings.Title(s))
	}
	return strings.Join(parts, ".")
}
