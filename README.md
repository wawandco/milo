# Milo

This is a linter for HTML written in Go. The goal is to provide a single binary that can lint HTML in the context of a CI server without installing other tools.

## Design considerations

- Milo considers html partials and validates the rules that apply to these.
- Milo considers other languages on top of html as erb and plush.

## Installation

You can pull Milo's binary from Github's [releases](https://github.com/wawandco/milo/releases) folder:

#### MacOS
```sh
$ curl -OL https://github.com/wawandco/milo/releases/latest/download/milo_Darwin_x86_64.tar.gz
$ tar -xvzf milo_Darwin_x86_64.tar.gz
$ sudo mv milo /usr/local/bin/milo
# or if you have ~/bin folder setup in the environment PATH variable
$ mv milo ~/bin/milo
```

#### Linux
```sh
$ wget https://github.com/wawandco/milo/releases/latest/download/milo_Linux_x86_64.tar.gz
$ tar -xvzf milo_Linux_x86_64.tar.gz
$ sudo mv milo /usr/local/bin/milo
```

## Usage

```
milo [folder or file]
```

Example:

```
milo templates
milo templates/file.html
```

# Rules

Milo checks the following rules (most of these come from [htmlhint](https://htmlhint.com/docs/user-guide/list-rules)):

### Head Rules

- [0001] Doctype must be declared.
- [0002] Doctype must be valid.
- [0004] `<title>` must be present inside `<head>` tag.
- [TODO] `<style>` tags cannot be used.
- [TODO] The `<script>` tag cannot be used in a `<head>` tag.

### Attributes

- [TODO] attr-lowercase: All attribute names must be in lowercase.
- [TODO] attr-no-duplication: Elements cannot have duplicate attributes.
- [TODO] attr-no-unnecessary-whitespace: No spaces between attribute names and values.
- [TODO] attr-unsafe-chars: Attribute values cannot contain unsafe chars.
- [TODO] attr-value-double-quotes: Attribute values must be in double quotes.
- [TODO] attr-value-not-empty: All attributes must have values.
- [TODO] alt-require: The alt attribute of an element must be present and alt attribute of area[href] and input[type=image] must have a value.
- [TODO] id-class-ad-disabled: The id and class attributes cannot use the ad keyword, it will be blocked by adblock software.
- [TODO] id-class-value: The id and class attribute values must meet the specified rules.
- [TODO] id-unique: The value of id attributes must be unique.

### Tags

- [TODO] tags-check: Allowing specify rules for any tag and validate that
- [TODO] tag-pair: Tag must be paired.
- [TODO] tag-self-close: Empty tags must be self closed.
- [TODO] tagname-lowercase: All html element names must be in lowercase.
- [TODO] empty-tag-not-self-closed: The empty tag should not be closed by self.
- [TODO] src-not-empty: The src attribute of an img(script,link) must have a value.
- [TODO] href-abs-or-rel: An href attribute must be either absolute or relative.
- [TODO] OL and UL should only have LI siblings.

### Inline
- [0003] Inline css is not allowed p.e: style="background-color: red;".
- [TODO] inline-style-disabled: Inline style cannot be use.
- [TODO] inline-script-disabled: Inline script cannot be use.

### Formatting

- [TODO] space-tab-mixed-disabled: Do not mix tabs and spaces for indentation.
- [TODO] spec-char-escape: Special characters must be escaped.


## Copyright

Milo is Copyright © 2020 Wawandco SAS. It is free software, and may be redistributed under the terms specified in the LICENSE file.


