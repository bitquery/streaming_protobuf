# Protobuf Schemas for Streaming

Protobuf schemas for the data streamed and stored in S3

Contains the schemas by the type of blockchain


Example to use utility to decode

```
brew install protobuf-c
lz4cat FILE.block.lz4 | protoc-c --decode evm_messages.BlockMessage evm/block_message.proto > output.txt
```
