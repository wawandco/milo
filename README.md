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

Milo checks the following rules:

-  Doctype must be declared.
-  Doctype must be valid.
-  [TODO] Inline css is not allowed.
-  [TODO] OL and UL should only have LI siblings.
-  [TODO] The `<script>` tag cannot be used in header.
-  [TODO] `<style>` tags cannot be used.
-  [TODO] `<title>` must be present in header.
-  [TODO] `style` attribute should not be used.

## Copyright

Milo is Copyright © 2020 Wawandco SAS. It is free software, and may be redistributed under the terms specified in the LICENSE file.


