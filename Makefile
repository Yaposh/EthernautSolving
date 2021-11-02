CONTRACT = fallback
OUT = api/$(CONTRACT)

all: sol abi

abi:
	abigen --abi=$(OUT)/$(CONTRACT).abi --bin=$(OUT)/$(CONTRACT).bin --pkg=$(CONTRACT) --out=$(OUT)/$(CONTRACT).go

sol:
	solc --overwrite --allow-paths=$(GOPATH)/src/Ethernaut/internal/lib --bin contracts/source/$(CONTRACT).sol -o $(OUT)
	solc --overwrite --allow-paths=$(GOPATH)/src/Ethernaut/internal/lib --abi contracts/source/$(CONTRACT).sol -o $(OUT)

gen:
	abigen --sol=contracts/source/$(CONTRACT).sol -pkg=$(CONTRACT) --out=$(OUT)/$(CONTRACT).go