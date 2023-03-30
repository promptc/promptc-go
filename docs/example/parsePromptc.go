package example

import (
	"fmt"
	"github.com/promptc/promptc-go/prompt"
)

func ParsePromptc(promptc string) {
	var file *prompt.File
	file = prompt.ParseFile(promptc)

	// get prompt info
	info := file.FileInfo
	fmt.Println(info)

	// get prompts
	prompts := file.Prompts
	fmt.Println(prompts)

	// get vars
	vars := file.Vars
	fmt.Println(vars)
	
	// get var constraints
	varConstraints := file.VarConstraint
	fmt.Println(varConstraints)
	// get parsing exceptions
	exceptions := file.Exceptions
	fmt.Println(exceptions)

}
