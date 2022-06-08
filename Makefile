.PHONY: paper.pdf
paper.pdf:
	pandoc \
		docs/paper.md \
		docs/metadata.yml \
		--template=docs/template.txt \
		--output=paper.pdf 