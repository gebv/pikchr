# MD to pic

**expereminal version and not sufficiently covered by tests**

cli tools to integrate into your development process.

mac
```
brew install gebv/tap/md2pikchrs
```

```bash
md2pikchrs -out ./_out -in ./_tmp/*.md
# 2021/01/31 10:58:54 md2pikchrs version: 1.0.2#f6608d1b842dfc76cb16c6de44703b12fccd95bd
# 2021/01/31 10:58:54 ./_tmp/demo.md total 4 code blocks
# 2021/01/31 10:58:54 ./_tmp/demo.md 4 interesting code blocks
# 2021/01/31 10:58:54 ./_tmp/demo.md 	 fil1.svg rendering...
# 2021/01/31 10:58:54 ./_tmp/demo.md 	 fil1.svg - OK
# 2021/01/31 10:58:54 ./_tmp/demo.md 	 foo_bar rendering...
# 2021/01/31 10:58:54 ./_tmp/demo.md 	 foo_bar.svg - OK
# 2021/01/31 10:58:54 ./_tmp/demo.md 	 foo_bar.svg rendering...
# 2021/01/31 10:58:54 ./_tmp/demo.md 	 foo_bar.svg - OK
# 2021/01/31 10:58:54 ./_tmp/demo.md 	 foo_bar.svg rendering...
# 2021/01/31 10:58:54 ./_tmp/demo.md 	 foo_bar.svg - OK
```

NOTE: Files with the same name are overwritten

**How it works?** `md2pikchrs` parse your markdown files and generated SVG files for each found code block with magic text. What is `magic text`? Follows example below

It is your markdown file

<pre class="language-text">
 # Some header

 some text

 more text...

 ```json
 {"foo":"bar"}
 ```

 ```bash
 echo "some scripts"
 ```

 end more more code blocks

 but if the block language matters __pikchr__ and after it there will be more text

 ```pikchr some text
 arrow right 200% "Markdown" "Source"
 box rad 10px "Markdown" "Formatter" "(markdown.c)" fit
 arrow right 200% "HTML+SVG" "Output"
 arrow <- down from last box.s
 box same "Pikchr" "Formatter" "(pikchr.c)" fit
 ```

 will be generated file `some_text.svg` with SVG diagram
</pre>

# How to use

NOTE: in the code block, the first line with the language name is called `string info`. First word it is language name.

- create code block with language `pikchr`
- after the language name, everything that will be specified will be the name of the generated file (before line break).
- manually add img in your markdown text with a previously known path to the file of interest

Nothing will be generated - not specified file name
<pre>
```pikchr
</pre>

Nothing will be generated - must be `pikchr` language name
<pre>
```json
</pre>

Will be generated svg file with name `foo_bar.svg`
<pre>
```pikchr foo bar
</pre>

... and the some name `foo_bar.svg`
<pre>
```pikchr foo_bar.svg
</pre>

... and the same name `foo_bar.svg`
<pre>
```pikchr foo bar.svg
</pre>

