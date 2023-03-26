<h1 align="center">⚙️ promptc-go</h1>
<p align="center">
  [ <a href="README.md">English</a> | <strong>简体中文</strong> ]
</p>

`promptc-go` 是 `promptc` 的 Go 语言实现，其使用
`promptc` 的标准以生成，解析，编译和执行 `promptc` 文件。

## 样例 promptc 文件

```ts
// 定义变量约束
vars: {
    x: int
    // var x with int type
    y: int {min: 0, max: 10}
    z: int {min: 0, max: 10, default: '5'}
}
// 狂野地定义变量约束
a: int {min: 0, max: 10}

// 定义 prompts
prompts: [
    // role: 'user' is meta info for ChatGPT
    // to make it empty, use {}
    '''role: 'user'
    You Entered: {x}
    Prompt Compiled: {%
        if (x == "1") {
            result = "Hello";
        } else {
            result = "Word!";
        }
    %}
    {%Q%}
    '''
]
```

## 语法

### 变量

#### 类型

目前 `promptc-go` 支持 `string`, `int`, `float` 类型

```ts
// 定义一个变量
myName : string { minLen: 3, maxLen: 10, default: "John" }
// 一个名为 `myName` 的 `string` 类型变量
// 其默认值为 "John"
// 最小长度为 3, 最大长度为 10

myAge : int { min: 18, max: 100, default: '18' }
// 一个名为`myAge` 的 `int` 类型变量
// 其缺省值为 18
// 最小值 18, 最大值 100

thisPrice : float { min: 0.01, default: '0.01' }
// 一个名为 `thisPrice` 的 `float` 类型变量
// 其缺省值为 0.01
// 最小值为 0.01 并且有无限大的最大值
```


#### 约束

- `string`
  - `minLen`: int
  - `maxLen`: int
- `int`
  - `min`: int64
  - `max`: int64
- `float`
  - `min`: float64
  - `max`: float64
- Shared
  - `default`: string

### Prompt

```py
{role: 'user'}
xx{x} {{x}} {%
    if (x > 12) {
        result = "good";
    } else {
        result = "bad";
    }
%}
```

任何位于 `{}` 内的内容为变量，例如在上面的例子中的 `{x}`  
任何位于 `{%%}` 内的将为 JavaScript 脚本
如果你想输出 `{` 或者 `}`，请使用 `{{` or `}}` 替代

prompt 的第一行是非常独特的，其为 prompt 提供了额外的信息。  
例如 ChatGPT 的 role 信息。例如：

```
role: 'user'
Show me more about {x}
```

如果你想提供空的额外信息，请使用 `{}` 作为第一行。  
虽然它不是必须的，因为一旦 hjson 解析失败，`promptc` 将会将第一行添加到你的 prompt 中，但这可能会导致大量未定义的行为。

#### 保留值

`promptc` 保留了一些值，以便在 prompt 中使用。

我们保留了 `{%Q%}` 用于 `'''`，这在 hjson 的多行文本语法中很难做到。

例如

```py
This is reserved {%Q%} {{%Q%}}
```

将会被编译为


```py
This is reserved ''' {%Q%}
```

#### Prompt 中的 JavaScript

> **Note**  
> 在 prompt 中使用 JavaScript 是一个实验性功能。  
> `promptc-go` 使用 [otto](https://github.com/robertkrimen/otto) 作为其 JavaScript 运行时

> **Warning**  
> 在 prompt 中使用 JavaScript 可能会使 prompt 变得脆弱并导致潜在的安全问题。  
> `promptc-go` **不会**对此负责。

`promptc` 支持内嵌 JavaScript 脚本使用 `{%%}` 语法。其支持 2 个模式：

- 标准模式（Standard）
- 简易模式（Easy）

##### 标准模式/Standard Mode

在标准模式中，执行完 js 脚本后，`promptc` 将从 `result` 变量中获取结果。

```py
You Entered: {x}
Prompt Compiled: {%
    if (x == "1") {
        result = "Hello";
    } else {
        result = "Word!";
    }
%}
```

如果输入 `x = 1`，结果将是：

```
You Entered: 1
Prompt Compiled: Hello
```

##### 简易模式/Easy Mode

在简易模式中，`promptc` 将从 js 脚本的返回值中获取结果。
为使用简易模式，需要在 prompt 的脚本开头添加一个 `E`。(`{%E /*script here*/ %}`)

```py
You Entered: {x}
Prompt Compiled: {%E
    if (x == "1") {
        return "Hello";
    } else {
        return "Word!";
    }
%}
```

如果输入 `x = 1`，结果将是：

```
You Entered: 1
Prompt Compiled: Hello
```

中简易模式中，`promptc` 将会将脚本包装在一个函数中，以便使用 `return` 语句。  
例如上面的例子实际将会被编译为：

```js
result = (function(){
  if (x == "1") {
    return "Hello"
  } else {
    return "Word!";
  }  
}()
```
