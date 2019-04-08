package generator

import (
	. "github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
)

func Usecase(name string, action string) error {
	f := NewFile("usecase")
	f.Type().Id(strcase.ToCamel(name) + "Usecase").Interface(
		Id(action).Params(
			Id("ipt").Id("input." + strcase.ToCamel(name) + strcase.ToCamel(action)),
		).Params(
			Id("opt").Id("output." + strcase.ToCamel(name) + strcase.ToCamel(action)),
		),
	)
	err := f.Save("applications/usecases/" + strcase.ToSnake(name) + ".go")
	return err
}
