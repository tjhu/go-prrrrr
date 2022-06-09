DATA_DIR=docs/paper

.PHONY: paper.pdf
paper.pdf:
	MERMAID_BIN=mmdc pandoc \
		--from=markdown \
		--filter pandoc-crossref \
		--citeproc \
		--bibliography=${DATA_DIR}/bib.bib \
		--csl=${DATA_DIR}/acm-sig-proceedings.csl \
		--filter mermaid-filter \
		${DATA_DIR}/header.tex \
		${DATA_DIR}/metadata.yml \
		${DATA_DIR}/toc.tex \
		${DATA_DIR}/intro.md \
		${DATA_DIR}/arch.md \
		${DATA_DIR}/impl.md \
		${DATA_DIR}/eval.md \
		${DATA_DIR}/conclusions.md \
		--output=paper.pdf 