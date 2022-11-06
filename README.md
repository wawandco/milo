![](https://github.com/wawandco/milo/blob/main/milo.png)

![](https://github.com/wawandco/milo/workflows/Test/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/Wawandco/milo)](https://goreportcard.com/report/github.com/Wawandco/milo)
# Milo

Milo is a linter tool that will check for both syntax correctness and style recommendations in HTML files. The end goal is to have a single binary that could be used in the context of a CI server which could check the codebase HTML before PR's get merged.

## Important Considerations

- Milo will check the HTML syntax of HTML only. ([see this](https://html.spec.whatwg.org/multipage/syntax.html))
- Milo will also check some style best practices and point those as issues to fix.
- Rules can be enabled/disabled. (See configuration)
- Milo considers html partials and validates the rules that apply to these.
- Milo considers erb and [plush](https://github.com/gobuffalo/plush) as part of the HTML.

## Installation

You can download the precompiled Milo binary from the [releases page](https://github.com/wawandco/milo/releases).

### With Go

If you have Go installed in your system you can build from source:

```sh
go install github.com/wawandco/milo/cmd/milo@latest
```
### Github Action

On github you can use Milo action to run in your workflows. For example:

```yaml
name: Main Workflow
on: [push]
jobs:
  lint:
    name: HTML Linting (Milo)
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      # Run the milo 
      - uses: wawandco/milo-action@v0.2.0
        with:
          folder: "templates"
          version: "v0.6.0"

  lint-go:
    ...
```

See more instructions on the [milo-action](https://github.com/wawandco/milo-action) repo.

## Usage

```
milo review [folder or file]
```

Example:

```
milo review templates
milo review templates/file.html
```

### Configuration

By default Milo will run all the linters it has. However, some teams will want to disable some of the linters in the list, if this is your case you can add a `.milo.yml` file in the root of your codebase.
Once you have installed Milo it can generate that file by running:
```
milo init
```

That `.milo.yml` will look like the following example:

```
output: text # could be `text`, `github` or `silent`
reviewers:
  - doctype/present 
  - ...
```

This file will be used by the Milo binary to determine which reviewers to run our files against.

## Reviewers

Milo uses the following reviewers:

### Head

- [0001] Doctype must be declared.
- [0002] Doctype must be valid.
- [0004] `<title>` must be present inside `<head>` tag.
- [0005] `<style>` must not be used.

### Tags

- [0006] All HTML element names must be in lowercase.
- [0007] `src`, `href` and `data` attributes of must have a value.
- [0008] `ol` and `ul` must only have `li` direct child tags.
- [0015] tag-pair: Tag must be paired.

### Attributes

- [0013] attr-lowercase: All attribute names must be in lowercase.
- [0010] attr-no-duplication: Elements cannot have duplicate attributes.
- [0019] attr-no-unnecessary-whitespace: No spaces between attribute names and values.
- [0016] attr-unsafe-chars: Attribute values cannot contain unsafe chars.
- [0018] attr-value-double-quotes: Attribute values must be in double quotes.
- [0011] attr-value-not-empty: All attributes must have values.
- [0012] alt-require: The alt attribute of an element must be present and alt attribute of area[href] and input[type=image] must have a value.
- [0014] id-unique: The value of id attributes must be unique.

### Inline

- [0003] Inline css is not allowed p.e: style="background-color: red;".
- [0009] inline-style-disabled: Inline style cannot be used.
- [0017] inline-script-disabled: Inline script cannot be used.

## Credits

This repo depends heavily in the following libraries that deserve all the credit for making Milo possible:

- golang.org/net/html 

We copied this in our source because we needed to make some modifications to it. Our goal long term goal is to contribute back as much as possible.

## Copyright

Milo is Copyright © 2020 Wawandco SAS. It is free software, and may be redistributed under the terms specified in the LICENSE file.


