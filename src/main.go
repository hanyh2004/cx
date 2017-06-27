// Tasks
// Create a cxStruct             x
// Create a Module Struct        x
// Create a Function struct      x
// Create a DataType struct      x
// Create a Context struct       x
// Create an Expression struct   x
// Create basic functions        

// Need: Mechanism to execute functions/expressions               x
// Need: Mechanism to constraint data types
// Need: Mechanism for explicitly casting types                   

// Need: Mechanism for reflection                                 
// Need: Mechanism to determine "affordances" and "restrictions"  
// Need: Start with a null object and call affordances            

// Need: Mechanism to serialize structs                           
// Need: Mechanism to hash and pull data from program             

package main

import (
	"fmt"
	"log"
	"os"
	
	//"github.com/skycoin/skycoin/src/cipher/encoder"
	. "github.com/skycoin/cx/src/base"
)

func init() {
	log.SetOutput(os.Stdout)
}

func allClear(e []error) bool {
	for i := 1; i < len(e); i++ {
		if e[i] != nil {
			return false
		}
	}
	return true
}

func dbg(elt string) {
	fmt.Println(elt)
}

func main() {
	fmt.Println(AddContext())
	fmt.Println(AddContext().AddModule())
	fmt.Println(AddContext().AddModule().AddFunction())
}
