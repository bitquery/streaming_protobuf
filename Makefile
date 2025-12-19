all: generate_evm generate_market generate_offchain generate_solana generate_ton generate_tron generate_utxo

generate_evm:
	protoc \
	-I=. \
	--go_out=. \
	--go_opt="Mevm/block_message.proto=evm/messages;evm_messages" \
	--go_opt="Mevm/dex_block_message.proto=evm/messages;evm_messages" \
	--go_opt="Mevm/dex_pool_block_message.proto=evm/messages;evm_messages" \
	--go_opt="Mevm/parsed_abi_block_message.proto=evm/messages;evm_messages" \
	--go_opt="Mevm/token_block_message.proto=evm/messages;evm_messages" \
	$(shell find ./evm -type f -name '*.proto')
	protoc \
    	-I=. \
		--python_out="evm/python" \
    	$(shell find ./evm -type f -name '*.proto')

generate_market:
	protoc \
	-I=. \
	--go_out=. \
	--go_opt="Mmarket/marketdata.proto=market/messages;marketdata_messages" \
	--go_opt="Mmarket/trades.proto=market/messages;marketdata_messages" \
	--go_opt="Mmarket/price_index.proto=market/messages;marketdata_messages" \
	$(shell find ./market -type f -name '*.proto')
	protoc \
    	-I=. \
		--python_out="market/python" \
    	$(shell find ./market -type f -name '*.proto')

generate_offchain:
	protoc \
	-I=. \
	--go_out=. \
	--go_opt="Moffchain/entities.proto=offchain/messages;entities_messages" \
	$(shell find ./offchain -type f -name '*.proto')
	protoc \
    	-I=. \
		--python_out="offchain/python" \
    	$(shell find ./offchain -type f -name '*.proto')

generate_solana:
	protoc \
	-I=. \
	--go_out=. \
	--go-grpc_out=. \
	--go_opt=paths=source_relative \
	--go-grpc_opt=paths=source_relative \
	--go_opt="Msolana/block_message.proto=github.com/bitquery/streaming_protobuf/v2/solana/messages;solana_messages" \
	--go_opt="Msolana/dex_block_message.proto=github.com/bitquery/streaming_protobuf/v2/solana/messages;solana_messages" \
	--go_opt="Msolana/parsed_idl_block_message.proto=github.com/bitquery/streaming_protobuf/v2/solana/messages;solana_messages" \
	--go_opt="Msolana/token_block_message.proto=github.com/bitquery/streaming_protobuf/v2/solana/messages;solana_messages" \
	--go_opt="Msolana/corecast/corecast.proto=solana/corecast/stream;solana_corecast" \
	--go_opt="Msolana/corecast/request.proto=solana/corecast/stream;solana_corecast" \
	--go_opt="Msolana/corecast/stream_message.proto=solana/corecast/stream;solana_corecast" \
	--go-grpc_opt="Msolana/block_message.proto=github.com/bitquery/streaming_protobuf/v2/solana/messages;solana_messages" \
	--go-grpc_opt="Msolana/dex_block_message.proto=github.com/bitquery/streaming_protobuf/v2/solana/messages;solana_messages" \
	--go-grpc_opt="Msolana/parsed_idl_block_message.proto=github.com/bitquery/streaming_protobuf/v2/solana/messages;solana_messages" \
	--go-grpc_opt="Msolana/token_block_message.proto=github.com/bitquery/streaming_protobuf/v2/solana/messages;solana_messages" \
	--go-grpc_opt="Msolana/corecast/corecast.proto=solana/corecast/stream;solana_corecast" \
	--go-grpc_opt="Msolana/corecast/request.proto=solana/corecast/stream;solana_corecast" \
	--go-grpc_opt="Msolana/corecast/stream_message.proto=solana/corecast/stream;solana_corecast" \
	$(shell find ./solana -type f -name '*.proto' ! -name 'dex_stream.proto' ! -name 'transactions_stream.proto')
	protoc \
    	-I=. \
		--python_out="solana/python" \
    	$(shell find ./solana -type f -name '*.proto' ! -name 'dex_stream.proto' ! -name 'transactions_stream.proto')
	@echo "-- reorganizing generated files into target folders --"
	@mkdir -p solana/messages solana/corecast/stream
	@for f in block_message dex_block_message parsed_idl_block_message token_block_message; do \
	  if [ -f solana/$$f.pb.go ]; then mv -f solana/$$f.pb.go solana/messages/; fi; \
	  if [ -f solana/$$f_grpc.pb.go ]; then mv -f solana/$$f_grpc.pb.go solana/messages/; fi; \
	done
	# New unified CoreCast API: move generated files (corecast, request, stream_message)
	@for f in corecast request stream_message; do \
	  if [ -f solana/corecast/$$f.pb.go ]; then mv -f solana/corecast/$$f.pb.go solana/corecast/stream/; fi; \
	  if [ -f solana/corecast/$$f*_grpc.pb.go ]; then mv -f solana/corecast/$$f*_grpc.pb.go solana/corecast/stream/; fi; \
	done
	
 generate_ton:
	protoc \
	-I=. \
	--go_out=. \
	--go_opt="Mton/block_message.proto=ton/messages;ton_messages" \
	--go_opt="Mton/dex_block_message.proto=ton/messages;ton_messages" \
	--go_opt="Mton/event_message.proto=ton/messages;ton_messages" \
	--go_opt="Mton/trace_message.proto=ton/messages;ton_messages" \
	--go_opt="Mton/token_block_message.proto=ton/messages;ton_messages" \
	$(shell find ./ton -type f -name '*.proto')
	protoc \
		-I=. \
		--python_out="ton/python" \
		$(shell find ./ton -type f -name '*.proto')

 generate_tron:
	protoc \
	-I=. \
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
	protoc \
		-I=. \
		--python_out="tron/python" \
		$(shell find ./tron -type f -name '*.proto')

generate_utxo:
	protoc \
	-I=. \
	--go_out=. \
	--go_opt="Mutxo/block_message.proto=utxo/messages;utxo_messages" \
	--go_opt="Mutxo/parsed_block_message.proto=utxo/messages;utxo_messages" \
	$(shell find ./utxo -type f -name '*.proto')
	protoc \
		-I=. \
		--python_out="utxo/python" \
		$(shell find ./utxo -type f -name '*.proto')