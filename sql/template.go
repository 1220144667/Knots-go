package sql

const structTpl = `type {{.StructName | TableColumnCase}} struct {
{{range .Columns}} {{$length := len.Name .Comment}} {{if gt $length 0}} // {{.Comment}} {{else}} // {{.Name}} {{end}} {{$typeLen := len .Type}} {{if gt $typeLen 0}} {{.Type}} {{else}} {{.Name}} {{end}}
{{end}}
}
`

func ({{.TableName}}){{range $index, $field :=.Fields}}{{if $index}}
{{end}}{{.Name}}() {{.Type}} {
	return {{.Name}}
}
{{end}}