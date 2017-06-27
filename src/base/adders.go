package base

import (
	
)


func AddContext() *cxContext {
	return &cxContext{}
}

func (cxt *cxContext) AddModule() *cxContext {
	mod := &cxModule{}
	cxt.CurrentModule = mod
	cxt.Modules =
		append(cxt.Modules,
		mod)
	return cxt
}

func (cxt *cxContext) AddFunction() *cxContext {
	mod := cxt.CurrentModule
	fn := &cxFunction{}
	mod.CurrentFunction = fn
	mod.Functions =
		append(mod.Functions,
		fn)

	return cxt
}

func (cxt *cxContext) AddInput() *cxContext {
	mod := cxt.CurrentModule
	fn := mod.CurrentFunction
	
	fn.Input =
		append(mod.Functions,
		fn)

	return cxt
}

func (cxt *cxContext) Add() *cxContext {
	return cxt
}



























// Adders should not have any paramaters
// Information


// We should also have something like "selectors" in order to select a module, a function, a struct to work on
// A selector would modify "CurrentXXX"

func (cxt *cxContext) AddDefinition() *cxContext {
	currentModule := cxt.CurrentModule
	currentModule.Definitions =
		append(currentModule.Definitions,
		&cxDefinition{})
	// A definition is added to a MODULE
	// So we need to go Context -> Module and add it there
	// After determining the currentModule
	
	return cxt
}

func (mod *cxModule) AddFunction(fn *cxFunction) *cxModule {
	//mod.Functions[fn.Name] = fn
	mod.Functions = append(mod.Functions, fn)
	return mod
}

func (mod *cxModule) AddStruct(strct *cxStruct) *cxModule {
	//mod.Structs[strct.Name] = strct
	mod.Structs = append(mod.Structs, strct)
	return mod
}

func (mod *cxModule) AddImport(imp *cxModule) *cxModule {
	//mod.Imports[imp.Name] = imp
	mod.Imports = append(mod.Imports, imp)
	return mod
}

func (strct *cxStruct) AddField(field *cxField) *cxStruct {
	strct.Fields = append(strct.Fields, field)
	return strct
}

func (fn *cxFunction) AddExpression (expr *cxExpression) *cxFunction {
	fn.Expressions = append(fn.Expressions, expr)
	return fn
}

func (fn *cxFunction) AddInput(param *cxParameter) *cxFunction {
	fn.Inputs = append(fn.Inputs, param)
	return fn
}

func (fn *cxFunction) AddOutput(param *cxParameter) *cxFunction {
	fn.Output = param
	return fn
}

func (expr *cxExpression) AddArgument(arg *cxArgument) *cxExpression {
	expr.Arguments = append(expr.Arguments, arg)
	return expr
}


// Adders for Affordances

func (param *cxParameter) AddName (name *[]byte) *cxParameter {
	param.Name = name
	return param
}

func (param *cxParameter) AddType (name *[]byte) *cxParameter {
	param.Name = name
	return param
}
