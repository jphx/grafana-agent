---
aliases:
- /docs/agent/latest/flow/configuration-language/standard-library/json_decode
title: json_decode
---

# `json_decode` Function

`json_decode` decodes a string representing JSON into a River value.
`json_decode` will fail if the string argument provided cannot be parsed as
JSON.

A common use case of `json_decode` is to decode the output of a
[`local.file`][] component to a River value.

> Remember to escape double quotes when passing JSON string literals to `json_decode`.
>
> For example, the JSON value `{"key": "value"}` is properly represented by the
> string `"{\"key\": \"value\"}"`.

## Examples

```
> json_decode("15")
15

> json_decode("[1, 2, 3]")
[1, 2, 3]

> json_decode("null")
null

> json_decode("{\"key\": \"value\"}")
{
  key = "value",
}

> json_decode(local.file.some_file.contents)
"Hello, world!"
```

[`local.file`]: {{< relref "../../components/local.file.md" >}}
