# Makefile for trapmux
#

# If go is not in your path, you can put the full pathname here
A2X = $(shell which a2x)

MAN_SRC = trapmux.man.asciidoc
MANPAGE = trapmux.8
MANPDF = trapmux.man.pdf

all: pdf

man: $(MAN_SRC)
	$(A2X) -L -f manpage $(MAN_SRC) 

pdf: man
	(groff -man $(MANPAGE) | ps2pdf - $(MANPDF)) || true

clean:
	/bin/rm -f $(MANPAGE) $(MANPDF)


###EOF###
