package utils

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
)

// VersionInfoFields 版本信息字段
type VersionInfoFields struct {
	AppName   string //应用名称
	Branch    string //版本分支
	Version   string //版本号
	Revision  string //commit id
	BuildDate string //编译时间
	GoVersion string //golang版本
}

const versionInfoTmpl = `
{{.program}}, version: {{.version}} 
(branch: {{.branch}}; revision: {{.revision}})
  build date:   {{.buildDate}}
  go version:   {{.goVersion}}
`

func PrintVersion(fields VersionInfoFields) {
	m := map[string]string{
		"program":   fields.AppName,
		"version":   fields.Version,
		"branch":    fields.Branch,
		"revision":  fields.Revision,
		"buildDate": fields.BuildDate,
		"goVersion": fields.GoVersion,
	}

	tmpl, err := template.Must(template.New("version"), nil).Parse(versionInfoTmpl)
	if err != nil {
		panic(err)
	}

	buf := bytes.Buffer{}
	if exeErr := tmpl.Execute(&buf, m); exeErr != nil {
		panic(exeErr)
	}

	fmt.Println(strings.TrimSpace(buf.String()))
}
