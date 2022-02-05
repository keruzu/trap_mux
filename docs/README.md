# Man Pages
Until we decide on a better way, the man pages are generated from
the `asciidoc` sources in this directory.

We use the "a2x" command to generate the initial nroff-formated version,
which may need to be manually edited to clean it up a bit.  The "a2x" command
is:

    a2x -f manpage  trapmux.man.asciidoc

In some cases, you may need to add the `-L` or `--no-xmllint` argument
    
This creates the `trapmux.8` man page.  However, depending on the `a2x` and/or
`asciidoc` configuration on your system, you may have to edit the "trapmux.8"
file directly to remove the '[FIXME: source/manual]' string embedded within.
At present, we simply remove them.  There may also be places where you want
items on successive line without intervening lines (i.e. the `AUTHORS` section
of the generated man page).  In those cases, simply change the `.sp`
between those lines to `.br` commands.

For creating HTML versions of the man pages, simply use the `-f xhtml`
option to the `a2x` command:

    a2x -f xhtml trapmux.man.asciidoc

## Requires
To install `asciidoc`, perform:

    yum install -y asciidoc
