# promptc-cli

> The .NET command-line interface (CLI) is a cross-platform toolchain for developing, building, running, and publishing .NET applications.  -- [.NET Documentation](https://learn.microsoft.com/en-us/dotnet/core/tools/)

The promptc command-line interface (CLI) is a cross-platform toolchain for developing, building, running promptc files.

## Installation

? I cannot install it! I don't know how to install it!

## Config API Keys

```sh
> promptc show
{
  openai_token: ''
}
> promptc set openai_token <your_token>
{
  openai_token: <your_token>
}
```

## Analyse Promptc

```sh
> promptc analyse <prompt_file>
```

![](img/analyse.png)

The colour in Tokens sections shows parsed token type.

- `text/literal` is the literal text in prompt file. It shows in grey.
- `var` is the variable in prompt file. It shows in blue.
- `script` is the js script in prompt file. It shows in yellow.
- `reserved` is the reserved value in prompt file. It shows in white. (`'''` in last line)


## Compile Promptc

```sh
> promptc compile <prompt_file> <var_file>
```

It will show compile details and the compiled prompt.

`var_file` is a file which contains variables for promptc. It is a key-value pair file (or ini?).

```ini
> cat varfile
x=111
var1=This is a variable
```

## Run Promptc

```sh
> promptc run <prompt_file> <var_file>
```

Not ready