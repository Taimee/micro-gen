package generator

import (
	. "github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
)

func Interactor(name string, action string) error {
	f := NewFile("interactors")
	iname := strcase.ToCamel(name) + "Interactor"
	f.Type().Id(iname).Struct()
	f.Func().Params(
		Id("i").Id(iname),
	).Id(strcase.ToCamel(action)).Params(
		Id("ipt").Id("input." + strcase.ToCamel(name) + strcase.ToCamel(action)),
	).Params(
		Id("opt").Id("output." + strcase.ToCamel(name) + strcase.ToCamel(action)),
	).Block()
	err := f.Save("applications/interactors/" + strcase.ToSnake(name) + ".go")
	return err
}
