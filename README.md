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
        return "good";
    } else {
        return "bad";
    }
%}
```

Anything in `{}` will be variable, e.g. `{x}` in previous example  
Anything in `{%%}` will be js scripts  
If you want to show `{` or `}`, use `{{` or `}}` instead