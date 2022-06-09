DATA_DIR=docs/paper

.PHONY: paper.pdf
paper.pdf:
	MERMAID_BIN=mmdc pandoc \
		--from=markdown \
		--citeproc \
		--bibliography=${DATA_DIR}/bib.bib \
		--filter pandoc-mermaid \
		${DATA_DIR}/header.tex \
		${DATA_DIR}/metadata.yml \
		${DATA_DIR}/toc.tex \
		${DATA_DIR}/intro.md \
		${DATA_DIR}/arch.md \
		${DATA_DIR}/impl.md \
		${DATA_DIR}/eval.md \
		${DATA_DIR}/conclusions.md \
		--output=paper.pdf 