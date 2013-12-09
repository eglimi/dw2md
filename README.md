# DokuWiki To Markdown Converter

This program converts a DokuWiki formatted text to markdown format.

It is by no means complete. I implemented just what was required to transform a
set of files. If you are interested in more functionality, please send me a
pull request (preferred) or request a feature with an issue.

It uses [Ragel State Machine Compiler](http://www.complang.org/ragel/) for
scanning the DokuWiki text.

dw2md is written in Go.

## Status

### Supported DokuWiki Syntax

* heading 1-5
* bold, italic, underline, monospace
* &lt;code&gt; and &lt;file&gt;
* &lt;del&gt;

### Unsupported DokuWiki Syntax

* subscript and superscript
* footnotes
* links and images
* tables
* lists
* quoting
* syntax highlighting for &lt;code&gt;
* &lt;html&gt; and &lt;php&gt; tags
* extensions from addon

## Limitations

dw2md does not check the validity of the given DokuWiki text. It is assumed
that the text is valid. If the input text is invalid, the output might be
anything.

Error reporting should be added / improved.

## Installation

	go get github.com/eglimi/dw2md

## Usage

If no options are given, dw2md assumes input from stdin, and outputs the result
to stdout. Alternatively, `-i` can be used to specify an input file, and `-o`
can be used to specify an output file.

Check `dw2md -help` for more information.

## Build Status

https://travis-ci.org/eglimi/dw2md.png
