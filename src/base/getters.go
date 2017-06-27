package base

import (
	
)

func (cxt *cxContext) GetCurrentModule() *cxModule {
	return cxt.CurrentModule
}

func (cxt *cxContext) GetCurrentFunction() *cxFunction {
	mod := cxt.CurrentModule
	return mod.CurrentFunction
}
