package generator

import (
	. "github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
)

func Handler(name string, action string) error {
	f := NewFile("handlers")
	f.Type().Id(strcase.ToCamel(name) + "Handler").Interface(
		Id(strcase.ToCamel(action)).Params(
			Id("c").Id("interfaces.Context"),
		).Params(
			Id("err").Id("error"),
		),
	)
	err := f.Save("interfaces/handlers/" + strcase.ToSnake(name) + ".go")
	return err
}
