package generator

import (
	. "github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
)

func Usecase(name string, action string) error {
	f := NewFile("usecase")
	f.Type().Id(strcase.ToCamel(name) + "Usecase").Interface(
		Id(strcase.ToCamel(action)).Params(
			Id("ipt").Id("input." + strcase.ToCamel(name) + strcase.ToCamel(action)),
		).Params(
			Id("opt").Id("output." + strcase.ToCamel(name) + strcase.ToCamel(action)),
		),
	)
	err := f.Save("applications/usecases/" + strcase.ToSnake(name) + ".go")
	if err != nil {
		return err
	}

	f = NewFile("input")
	f.Type().Id(strcase.ToCamel(name) + strcase.ToCamel(action)).Struct()
	err = f.Save("applications/usecases/inputs/" + strcase.ToSnake(name) + ".go")
	if err != nil {
		return err
	}

	f = NewFile("output")
	f.Type().Id(strcase.ToCamel(name) + strcase.ToCamel(action)).Struct()
	err = f.Save("applications/usecases/outputs/" + strcase.ToSnake(name) + ".go")
	if err != nil {
		return err
	}

	return nil
}
