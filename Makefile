
all: generate_evm generate_market generate_offchain generate_solana generate_ton generate_tron generate_utxo

generate_evm:
	protoc \
	-I=. \
	--experimental_allow_proto3_optional \
	--go_out=. \
	--go_opt="Mevm/block_message.proto=evm/messages;evm_messages" \
	--go_opt="Mevm/dex_block_message.proto=evm/messages;evm_messages" \
	--go_opt="Mevm/parsed_abi_block_message.proto=evm/messages;evm_messages" \
	--go_opt="Mevm/token_block_message.proto=evm/messages;evm_messages" \
	$(shell find ./evm -type f -name '*.proto')

generate_market:
	protoc \
	-I=. \
	--experimental_allow_proto3_optional \
	--go_out=. \
	--go_opt="Mmarket/marketdata.proto=market/messages;marketdata_messages" \
	--go_opt="Mmarket/trades.proto=market/messages;marketdata_messages" \
	--go_opt="Mmarket/price_index.proto=market/messages;marketdata_messages" \
	$(shell find ./market -type f -name '*.proto')

generate_offchain:
	protoc \
	-I=. \
	--experimental_allow_proto3_optional \
	--go_out=. \
	--go_opt="Moffchain/entities.proto=offchain/messages;entities_messages" \
	$(shell find ./offchain -type f -name '*.proto')

 generate_solana:
	protoc \
	-I=. \
	--experimental_allow_proto3_optional \
	--go_out=. \
	--go_opt="Msolana/block_message.proto=solana/messages;solana_messages" \
	--go_opt="Msolana/dex_block_message.proto=solana/messages;solana_messages" \
	--go_opt="Msolana/parsed_idl_block_message.proto=solana/messages;solana_messages" \
	--go_opt="Msolana/token_block_message.proto=solana/messages;solana_messages" \
	$(shell find ./solana -type f -name '*.proto')
	
 generate_ton:
	protoc \
	-I=. \
	--experimental_allow_proto3_optional \
	--go_out=. \
	--go_opt="Mton/block_message.proto=ton/messages;ton_messages" \
	--go_opt="Mton/dex_block_message.proto=ton/messages;ton_messages" \
	--go_opt="Mton/event_message.proto=ton/messages;ton_messages" \
	--go_opt="Mton/trace_message.proto=ton/messages;ton_messages" \
	--go_opt="Mton/token_block_message.proto=ton/messages;ton_messages" \
	$(shell find ./ton -type f -name '*.proto')	
	
 generate_tron:
	protoc \
	-I=. \
	--experimental_allow_proto3_optional \
	--go_out=. \
	--go_opt="Mevm/block_message.proto=github.com/bitquery/streaming_protobuf/v2/evm/messages;evm_messages" \
	--go_opt="Mevm/parsed_abi_block_message.proto=github.com/bitquery/streaming_protobuf/v2/evm/messages;evm_messages" \
	--go_opt="Mevm/token_block_message.proto=github.com/bitquery/streaming_protobuf/v2/evm/messages;evm_messages" \
	--go_opt="Mevm/dex_block_message.proto=github.com/bitquery/streaming_protobuf/v2/evm/messages;evm_messages" \
	--go_opt="Mtron/block_message.proto=tron/messages;tron_messages" \
	--go_opt="Mtron/dex_block_message.proto=tron/messages;tron_messages" \
	--go_opt="Mtron/parsed_abi_block_message.proto=tron/messages;tron_messages" \
	--go_opt="Mtron/token_block_message.proto=tron/messages;tron_messages" \
	$(shell find ./tron -type f -name '*.proto')

generate_utxo:
	protoc \
	-I=. \
	--experimental_allow_proto3_optional \
	--go_out=. \
	--go_opt="Mutxo/block_message.proto=utxo/messages;utxo_messages" \
	--go_opt="Mutxo/parsed_block_message.proto=utxo/messages;utxo_messages" \
	$(shell find ./utxo -type f -name '*.proto')