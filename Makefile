DATA_DIR=docs/paper

.PHONY: paper.pdf
paper.pdf:
	pandoc \
		${DATA_DIR}/metadata.yml \
		${DATA_DIR}/toc.md \
		${DATA_DIR}/intro.md \
		${DATA_DIR}/arch.md \
		${DATA_DIR}/impl.md \
		${DATA_DIR}/eval.md \
		${DATA_DIR}/conclusions.md \
		--output=paper.pdf 