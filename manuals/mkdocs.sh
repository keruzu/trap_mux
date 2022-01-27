#!/bin/bash

for source in `ls -1 *.asciidoc` ; do
    manpage=`echo $source | sed -e 's/\.asciidoc//'`
    echo "$source -> $manpage"
    echo "    a2x -L -f manpage $source"

    pdf=`echo $manpage | sed -e 's/\.man//'`
    echo "    ... $manpage -> $pdf"
    echo "    groff -man $manpage | ps2pdf - $pdf"

    echo
done

