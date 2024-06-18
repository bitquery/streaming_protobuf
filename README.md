# Protobuf Schemas for Streaming

Protobuf schemas for the data streamed and stored in S3

Contains the schemas by the type of blockchain


Example to use utility to decode

```
brew install protobuf-c
lz4cat FILE.block.lz4 | protoc-c --decode evm_messages.BlockMessage evm/block_message.proto > output.txt
```
examples

```
lz4cat 000025207397_0x72e7402c5fc28ef31e0528c5f25a95469c124705952ce89d5a10d60f334c4057_b417e0808dc463173e34c088a028ed2152ef38adbfb6033d3ea2943039c7b463.block.lz4 | protoc-c --decode evm_messages.BlockMessage evm/block_message.proto > out2.txt

lz4cat 000120696577_0x299f8fcfdeea511802fa48aa4f18a916b27a0de0aedff2332d306c2244b89284_499f95c2aa01a45db19a635652810740cc412b286fcd7a8660839a2a1cbd8668.block.lz4 | protoc-c --decode evm_messages.BlockMessage evm/block_message.proto > out6.txt
```