![](https://github.com/wawandco/milo/blob/master/milo-logo.png)

![](https://github.com/wawandco/milo/workflows/Test/badge.svg)
# Milo

This is a linter for HTML written in Go. The goal is to provide a single binary that can review HTML in the context of a CI server without installing other tools.

## Initial Considerations

- Milo considers html partials and validates the rules that apply to these.
- Milo considers erb and plush as part of the HTML and works around these.

## Installation

We recommend using [gobinaries.com](gobinaries.com) to get Milo.

```sh
curl -sf https://gobinaries.com/wawandco/milo/cmd/milo | sh
```

You can also download Milo binaries from our releases folder.

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

Once you have installed Milo you can run 
```
milo init
```

And it will generate a `.milo.yml` file that looks something like:

```
output: github # could be `github` or `silent`
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
- github.com/andybalholm/cascadia
- github.com/PuerkitoBio/goquery

We copied those repos in our source because we needed to make some modifications to it. Our goal long term goal is to contribute back as much as possible.

## Copyright

Milo is Copyright © 2020 Wawandco SAS. It is free software, and may be redistributed under the terms specified in the LICENSE file.


