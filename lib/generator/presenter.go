package generator

import (
	. "github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
)

func Presenter(name string, action string) error {
	f := NewFile("presenters")
	pname := strcase.ToCamel(name) + "Presenter"
	f.Type().Id(pname).Interface(
		Id(action).Params(
			Id("opt").Id("output." + strcase.ToCamel(name) + strcase.ToCamel(action)),
		),
	)

	f.Type().Id(pname + "Impl").Struct()
	f.Func().Params(
		Id("p").Id(pname),
	).Id(strcase.ToCamel(action)).Params(
		Id("opt").Id("output." + strcase.ToCamel(name) + strcase.ToCamel(action)),
	)
	err := f.Save("interfaces/presenters/" + strcase.ToSnake(name) + ".go")
	return err
}
