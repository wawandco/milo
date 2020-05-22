# Milo

This is a linter for html written in Go. The goal is to provide a single binary that can lint HTML in the context of a CI server without installing other tools.

## Installation

You can pull Milo's binary from Github's [releases](#) folder:

Mac
```
```

Linux
```
```

Windows
```
```


## Usage

Milo checks the following rules:

-  Doctype must be declared.
-  Doctype must be valid.
-  The `<script>` tag cannot be used in header.
-  `<style>` tags cannot be used.
-  `<title>` must be present in tag.

## Copyright

Milo is Copyright © 2008-2015 Wawandco SAS. It is free software, and may be redistributed under the terms specified in the LICENSE file.


