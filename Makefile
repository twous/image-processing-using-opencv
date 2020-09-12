.PHONY: all
all: compile bindings

.PHONY: compile
compile:
	# compile vendor management contract
	solc \
	--optimize \
	--optimize-runs 200 \
	--bin \
	--abi \
	--overwrite \
	-o bin/vendorManagement \
	contracts/VendorManagement.sol
	# compile vendor factory
	solc \
	--optimize \
	--optimize-runs 200 \
	--bin \
	--abi \
	--overwrite \
	-o bin/vendorFactory \
	contracts/VendorFactory.sol
	# compile vending machine
	solc \
	--optimize \
	--optimize-runs 200 \
	--bin \
	--abi \
	--overwrite \
	-o bin/vendingMachine \
	contracts/VendingMachine.sol

.PHONY: bindings
bindings:
	# build bindings for vendor management
	abigen \
	--abi bin/vendorManagement/VendorManagement.abi \
	--bin bin/vendorManagement/VendorManagement.bin \
	--pkg vendormanagement \
	--out bindings/vendorManag