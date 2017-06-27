package base

type cxAction string
const (
	acAddArgument cxAction = "AddArgument"
	acAddDefinition cxAction = "AddDefinition"
	acAddExpression cxAction = "AddExpression"
	acAddInput cxAction = "AddInput"
	acAddOutput cxAction = "AddOutput"
)

type cxType struct {
	Name *[]byte
	Affordances []*cxAffordance
}

type cxField struct {
	Name *[]byte
	Typ *cxType
	Affordances []*cxAffordance
}

type cxStruct struct {
	Name *[]byte
	Fields []*cxField
	Affordances []*cxAffordance
}

/*
  Context
*/

// Context should also have information about the last added structures. We can manage them as stacks, where makers consume them
type cxContext struct {
	Modules []*cxModule

	// Execution information
	CurrentModule *cxModule
	CurrentFunction *cxFunction
	Scopes [][]*cxDefinition
	Affordances []*cxAffordance
}

/*
  Functions
*/

type cxParameter struct {
	Name *[]byte
	Typ *cxType
	Affordances []*cxAffordance
}

type cxArgument struct {
	Typ *cxType
	Value *[]byte
	Affordances []*cxAffordance
}

type cxExpression struct {
	Fn *cxFunction
	Arguments []*cxArgument
	Affordances []*cxAffordance
}

type cxFunction struct {
	Name *[]byte
	Inputs []*cxParameter
	Output *cxParameter
	Expressions []*cxExpression
	Affordances []*cxAffordance
}

/*
  Modules
*/

type cxDefinition struct {
	Name *[]byte
	Typ *cxType
	Value *[]byte
	Affordances []*cxAffordance
}

type cxModule struct {
	Name *[]byte
	Imports []*cxModule
	Functions []*cxFunction
	Structs []*cxStruct
	Definitions []*cxDefinition

	CurrentFunction *cxFunction
	
	Affordances []*cxAffordance
}

/*
  Affordances
*/

type cxAffordance struct {
	Action cxAction
	Input interface{}
}
