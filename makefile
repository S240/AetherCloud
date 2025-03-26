export ROOT_MOD=github.com/S240/AetherCloud


.PHONY:	all
all: gen-gateway  gen-uploadfile

.PHONY: gen-gateway 
gen-gateway: 
	@cd app/apigw && cwgo server --type HTTP --idl ../../idl/apigw.proto  --service apigw --module ${ROOT_MOD}/app/apigw 

.PHONY: gen-uploadfile
gen-uploadfile:
	@mkdir -p client/uploadfile
	@cd client/uploadfile && cwgo client --type RPC --idl ../../idl/uploadfile.proto --service uploadfile --module ${ROOT_MOD}/client/uploadfile -I ../../idl
	@mkdir -p app/uploadfile 
	@cd app/uploadfile && cwgo server --type RPC --idl ../../idl/uploadfile.proto --service uploadfile --module ${ROOT_MOD}/app/uploadfile -I ../../idl
