package prompt

// ParseFile @deprecated
// Deprecated: FunctionName is deprecated. Use ParsePromptC instead.
func ParseFile(content string) *PromptC {
	return ParsePromptC(content)
}

// ParseUnstructuredFile @deprecated
// Deprecated: FunctionName is deprecated. Use ParseBasicPrompt instead.
func ParseUnstructuredFile(content string) *PromptC {
	return ParseBasicPrompt(content)
}
