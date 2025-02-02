package main

import (
	"bytes"
	"embed"
	"fmt"
	"go/format"
	"log"
	"os"
	"path"
	"reflect"
	"runtime"
	"strings"
	"text/template"

	"github.com/cloudquery/cloudquery/plugins/source/vercel/codegen/recipes"
	"github.com/cloudquery/cloudquery/plugins/source/vercel/internal/vercel"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

//go:embed templates/*.go.tpl
var templatesFS embed.FS

func main() {
	var resources []*recipes.Resource
	resources = append(resources, recipes.DomainResources()...)
	resources = append(resources, recipes.TeamResources()...)
	resources = append(resources, recipes.ProjectResources()...)
	resources = append(resources, recipes.DeploymentResources()...)

	for _, r := range resources {
		r.Infer()
	}
	if err := recipes.SetParentChildRelationships(resources); err != nil {
		log.Fatal(err)
	}
	for _, r := range resources {
		r.GenerateNames()

		_, filename, _, ok := runtime.Caller(0)
		if !ok {
			log.Fatal("Failed to get caller information")
		}
		codegenDir := path.Join(path.Dir(filename), "..", "resources", "services")

		generateTable(codegenDir, *r)
	}
}

func milliTimeTransformer(field reflect.StructField) (schema.ValueType, error) {
	milliTime := vercel.MilliTime{}
	switch field.Type {
	case reflect.TypeOf(milliTime), reflect.TypeOf(&milliTime):
		return schema.TypeTimestamp, nil
	default:
		return schema.TypeInvalid, nil
	}
}

func generateTable(basedir string, r recipes.Resource) {
	var err error

	r.TableName = "vercel_" + r.TableName

	log.Println("Generating table", r.TableName)
	opts := []codegen.TableOption{
		codegen.WithSkipFields(r.SkipFields),
		codegen.WithExtraColumns(r.ExtraColumns),
		codegen.WithPKColumns(r.PKColumns...),
		codegen.WithTypeTransformer(milliTimeTransformer),
	}
	if r.UnwrapEmbeddedStructs {
		opts = append(opts, codegen.WithUnwrapAllEmbeddedStructs())
	}
	r.Table, err = codegen.NewTableFromStruct(r.TableName, r.DataStruct, opts...)

	if err != nil {
		log.Fatal(err)
	}

	r.Table.Resolver = r.ResolverFuncName
	r.Table.Multiplex = r.Multiplex
	r.ImportClient = strings.HasPrefix(r.Multiplex, "client.")
	r.Table.Relations = r.Relations
	r.Table.PreResourceResolver = r.PreResourceResolver
	r.Table.PostResourceResolver = r.PostResourceResolver

	for _, c := range r.Table.Columns {
		if strings.HasPrefix(c.Resolver, "client.") {
			r.ImportClient = true
		}
	}

	mainTemplate := r.Template + ".go.tpl"
	tpl, err := template.New(mainTemplate).ParseFS(templatesFS, "templates/"+mainTemplate)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to parse vercel templates: %w", err))
	}
	tpl, err = tpl.ParseFS(codegen.TemplatesFS, "templates/*.go.tpl")
	if err != nil {
		log.Fatal(fmt.Errorf("failed to parse recipes template: %w", err))
	}
	var buff bytes.Buffer
	if err := tpl.Execute(&buff, r); err != nil {
		log.Fatal(fmt.Errorf("failed to execute template: %w", err))
	}

	pkgPath := path.Join(basedir, r.Service)
	if err := os.Mkdir(pkgPath, 0755); err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}

	filePath := path.Join(pkgPath, r.Filename)
	content, err := format.Source(buff.Bytes())
	if err != nil {
		fmt.Println(buff.String())
		log.Fatal(fmt.Errorf("failed to format code for %s: %w", filePath, err))
	}
	if err := os.WriteFile(filePath, content, 0644); err != nil {
		log.Fatal(fmt.Errorf("failed to write file %s: %w", filePath, err))
	}
}
