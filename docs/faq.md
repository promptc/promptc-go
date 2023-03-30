# FAQ

## Whats the difference between `prompt` & `promptc` file?

`prompt` file is a file contains only a prompt block. `prompt` file can
contains a prompt block, in the block, it can contains implicit defined
vars (use `{}`) and some JavaScript scripts with `{%%}`. It do not need
use `{%Q%}` to represent `'''`. If you do not use `{%%}` and `{%Q%}`,
then `prompt` file should be a valid language chain text.

`promptc` file is a file contains full structured defined prompts. It
can include several prompt blocks (or 1), vars, var constraints, file
infos and even some JavaScripts.

`prompt` file provides a simple way to define a prompt. It should works
fine in most senarios. But if you want to define a complex prompt, for
example, you need var constraints or you need more than 1 prompt block,
then you should use `promptc` file.