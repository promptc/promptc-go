# `promptc-go`

`promptc-go` is a go implementation of `promptc`. It uses
`promptc` specification to generate prompts.

## Syntax

### Variable

#### Type

```ts
// declare a variable
myName : string { minLen: 3, maxLen: 10, default: "John" }
// a var named `myName` of type `string`
// with default value "John"
// min length 3, max length 10

myAge : int { min: 18, max: 100, default: 18 }
// a var named `myAge` of type `int`
// with default value 18
// min value 18, max value 100

thisPrice : float { min: 0.01, default: 0.01 }
// a var named `thisPrice` of type `float`
// with default value 0.01
// min value 0.01, and unlimited max value
```

Current `promptc-go` supports `string`, `int`, `float` types.

#### Constraint

- `string`
  - `minLen`
  - `maxLen`
- `int`
  - `min`
  - `max`
- `float`
  - `min`
  - `max`

### Prompt

```py
xx{x} {{x}} {%
    if (x > 12) {
        result = "good";
    } else {
        result = "bad";
    }
%}
```

Anything in `{}` will be variable, e.g. `{x}` in previous example  
Anything in `{%%}` will be js scripts  
If you want to show `{` or `}`, use `{{` or `}}` instead

#### JavaScript

Promptc supports js scripts in `{%%}`. And it contains 2 modes:
- Standard
- Easy

In standard mode, after run the js script, the promptc will get the result from `result` variable.
    
```py
You Entered: {x}
Prompt Compiled: {%
	if (x == "1") {
		result = "Hello"
	} else {
		result = "Word!";
	}
%}
```

If entered x = 1, the result will be:

```
You Entered: 1
Prompt Compiled: Hello
```

In easy mode, the promptc will get the result from return value of js script. And it will
add an `E` at the beginning of the prompt.

```py
You Entered: {x}
Prompt Compiled: {%E
	if (x == "1") {
		return "Hello"
	} else {
		return "Word!";
	}
%}
```

If entered x = 1, the result will be:

```
You Entered: 1
Prompt Compiled: Hello
```

In easy mode, the script will be wrapped in a function in order to enable `return` statement.  
i.e. this is the actual script that will be run:

```js
result = (function(){
  if (x == "1") {
    return "Hello"
  } else {
    return "Word!";
  }  
}()
```