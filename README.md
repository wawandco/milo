![](https://github.com/wawandco/milo/workflows/Test/badge.svg)

# Milo

This is a linter for HTML written in Go. The goal is to provide a single binary that can lint HTML in the context of a CI server without installing other tools.

## Initial Considerations

- Milo considers html partials and validates the rules that apply to these.
- Milo considers erb and plush as part of the HTML and works around these.
- Milo will start simple and write its output only to be compatible with github, maybe later we add/other formats.

## Installation

We recommend using [gobinaries.com](gobinaries.com) to get Milo.

```sh
curl -sf https://gobinaries.com/wawandco/milo | sh
```

You can also download Milo binaries from our releases folder.

## Usage

```
milo [folder or file]
```

Example:

```
milo templates
milo templates/file.html
```

### Configuration

Referees to run can get configured by creating a file named `.milo.yml` in the root of the folder to analize. An example of the .milo.yml file that can be used as a starting point is:

```
output: github # TODO!
reviewers:
  - doctype/present 
```

If Milo does not find this file in your folder it will run All the linters, the same if the reviewers list is empty.

## Reviewers

Milo checks the following referees:

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
- [TODO] attr-no-unnecessary-whitespace: No spaces between attribute names and values.
- [0016] attr-unsafe-chars: Attribute values cannot contain unsafe chars.
- [TODO] attr-value-double-quotes: Attribute values must be in double quotes.
- [0011] attr-value-not-empty: All attributes must have values.
- [0012] alt-require: The alt attribute of an element must be present and alt attribute of area[href] and input[type=image] must have a value.
- [0014] id-unique: The value of id attributes must be unique.

### Inline

- [0003] Inline css is not allowed p.e: style="background-color: red;".
- [0009] inline-style-disabled: Inline style cannot be use.
- [TODO] inline-script-disabled: Inline script cannot be use.

## Copyright

Milo is Copyright © 2020 Wawandco SAS. It is free software, and may be redistributed under the terms specified in the LICENSE file.


