test:
	@p=$(shell pwd); \
	string="aptos,bitcoin,elrond,ethereum,flow,helium,near,oasis,polkadot,stacks,sui,tron,zil,zksync,avax,cosmos,eos,filecoin,harmony,kaspa,nervos,oracle,solana,starknet,tezos,waves,zkspace"; \
	array=$$(echo $$string | tr ',' ' '); \
	for var in $$array; do \
		echo "Testing $$var"; \
		cd $$p/coins/$$var && go mod tidy && go test -v && echo "Test $$var success.\n"; \
	done
.PHONY: test