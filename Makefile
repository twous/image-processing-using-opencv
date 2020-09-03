.PHONY: all
all: compile bindings

.PHONY: compile
compile:
	# compile vendor management contract
	solc \
	--optimize \
	--optimize