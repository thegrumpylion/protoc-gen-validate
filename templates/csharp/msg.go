package csharp

const msgTpl = `
    public class {{ msgTyp . }}Validator : AbstractValidator<{{ msgTyp . }}>
    {
        public {{ msgTyp . }}Validator()
        {
            {{- range .Fields }}{{ with (context .) }}{{ $f := .Field }}
            RuleFor(msg => msg.{{ $f.Name }}).NotEmpty();
            {{- end }}{{ end }}
        }
        {{ range .NonOneOfFields }}
			{{ render (context .) }}
		{{ end }}
    }
`
