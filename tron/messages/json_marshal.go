package tron_messages

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"github.com/akamensky/base58"
	evm_messages "github.com/bitquery/streaming_protobuf/v2/evm/messages"
	"github.com/bitquery/streaming_protobuf/v2/pkg/encoder"
)

func (chain *Chain) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"Chain": string(chain.ChainId),
	})
}

func (header *BlockHeader) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"Number":            int(header.Number),
		"Hash":              hex.EncodeToString(header.Hash),
		"Timestamp":         header.Timestamp,
		"ParentHash":        hex.EncodeToString(header.ParentHash),
		"ParentNumber":      int(header.ParentNumber),
		"Version":           int(header.Version),
		"TxTrieRoot":        hex.EncodeToString(header.TxTrieRoot),
		"AccountStateRoot":  hex.EncodeToString(header.AccountStateRoot),
		"TransactionsCount": int(header.TransactionsCount),
	})
}

func (witness *Witness) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"Address":   encodeTronAddress(witness.Address),
		"Id":        witness.Id,
		"Signature": hex.EncodeToString(witness.Signature),
	})
}

func (header *TransactionHeader) MarshalJSON() ([]byte, error) {
	signatures := make([]string, len(header.Signatures))
	for i, s := range header.Signatures {
		signatures[i] = hex.EncodeToString(s)
	}
	return json.Marshal(map[string]any{
		"Hash":       hex.EncodeToString(header.Hash),
		"Fee":        header.Fee,
		"Index":      header.Index,
		"Expiration": header.Expiration,
		"Data":       hex.EncodeToString(header.Data),
		"FeeLimit":   header.FeeLimit,
		"Timestamp":  header.Timestamp,
		"Signatures": signatures,
		"Time":       header.Time,
		"FeePayer":   encodeTronAddress(header.FeePayer),
	})
}

func CallJsonStructure(call *evm_messages.ParsedAbiCall) map[string]any {
	return map[string]any{
		"Header":       callHeader(call.Header),
		"Arguments":    arguments(call.Arguments),
		"ReturnValues": arguments(call.ReturnValues),
		"Logs":         logs(call.Logs),
	}
}

func (tx *ParsedAbiTransaction) ToJSONStruct() map[string]any {
	calls := make([]map[string]any, len(tx.Calls))
	for i, call := range tx.Calls {
		calls[i] = CallJsonStructure(call)
	}
	return map[string]any{
		"TransactionHeader": tx.TransactionHeader,
		"Receipt":           tx.Receipt,
		"TransactionStatus": tx.TransactionStatus,
		"TransactionResult": tx.TransactionResult,
		"ContractInfo":      tx.ContractInfo,
		"Calls":             calls,
		"RewardWithdraw":    tx.RewardWithdraw,
	}
}

func (tx *ParsedAbiTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(tx.ToJSONStruct())
}

func LogJsonStructure(log *evm_messages.ParsedAbiLog) map[string]any {
	return map[string]any{
		"Header":    logHeader(log.Header),
		"Parsed":    parsedLogHeader(log.Parsed),
		"Topics":    topics(log.Topics),
		"Arguments": arguments(log.Arguments),
	}
}

func logs(logs []*evm_messages.ParsedAbiLog) []map[string]any {
	if len(logs) == 0 {
		return nil
	}
	result := make([]map[string]any, len(logs))
	for i, log := range logs {
		result[i] = LogJsonStructure(log)
	}
	return result
}

func topics(topics []*evm_messages.Topic) []string {
	if len(topics) == 0 {
		return nil
	}
	result := make([]string, len(topics))
	for i, topic := range topics {
		result[i] = hex.EncodeToString(topic.Hash)
	}
	return result
}

func logHeader(header *evm_messages.LogHeader) map[string]any {
	return map[string]any{
		"Index":    header.Index,
		"Address":  encodeTronAddress(header.Address),
		"Data":     hex.EncodeToString(header.Data),
		"Removed":  header.Removed,
		"IsSystem": boolValue(header.IsSystem),
	}
}

func parsedLogHeader(header *evm_messages.ParsedAbiLogHeader) map[string]any {
	return map[string]any{
		"Index":         header.Index,
		"SmartContract": encodeTronAddress(header.SmartContract),
		"Signature":     header.Signature,
	}
}

func callHeader(header *evm_messages.ParsedAbiCallHeader) map[string]any {
	return map[string]any{
		"Index":         header.Index,
		"Depth":         header.Depth,
		"CallerIndex":   header.CallerIndex,
		"InternalCalls": header.InternalCalls,
		"CallPath":      header.CallPath,
		"From":          encodeTronAddress(header.From),
		"To":            encodeTronAddress(header.To),
		"Create":        header.Create,
		"Input":         hex.EncodeToString(header.Input),
		"Gas":           header.Gas,
		"Value":         encoder.BytesToBigInt(header.Value, false).Text(10),
		"Output":        hex.EncodeToString(header.Output),
		"GasUsed":       header.GasUsed,
		"Error":         header.Error,
		"Opcode":        header.Opcode,
		"SelfDestruct":  header.SelfDestruct,
		"Delegated":     header.Delegated,
		"Success":       header.Success,
		"Reverted":      header.Reverted,
		"Signature":     header.Signature,
	}
}

func arguments(args []*evm_messages.ParsedArgument) []map[string]any {
	if len(args) == 0 {
		return nil
	}
	result := make([]map[string]any, len(args))
	for i, arg := range args {
		result[i] = map[string]any{
			"Name":    arg.Name,
			"Indexed": arg.Indexed,
			"Value":   argValue(arg.Value, ""),
		}
	}
	return result
}

func argValue(value *evm_messages.ArgumentValue, name string) map[string]any {
	var val any

	switch value.Value.(type) {
	case *evm_messages.ArgumentValue_Bytes:
		if value.Type == "address" {
			val = encodeTronAddress(value.Value.(*evm_messages.ArgumentValue_Bytes).Bytes)
		} else {
			val = hex.EncodeToString(value.Value.(*evm_messages.ArgumentValue_Bytes).Bytes)
		}
	case *evm_messages.ArgumentValue_UInt:
		val = value.Value.(*evm_messages.ArgumentValue_UInt).UInt
	case *evm_messages.ArgumentValue_Int:
		val = value.Value.(*evm_messages.ArgumentValue_Int).Int
	case *evm_messages.ArgumentValue_Bool:
		val = value.Value.(*evm_messages.ArgumentValue_Bool).Bool
	case *evm_messages.ArgumentValue_Array:
		array := make([]any, len(value.Value.(*evm_messages.ArgumentValue_Array).Array.Elements))
		for i, x := range value.Value.(*evm_messages.ArgumentValue_Array).Array.Elements {
			array[i] = argValue(x, "")
		}
	case *evm_messages.ArgumentValue_Tuple:
		tuple := value.Value.(*evm_messages.ArgumentValue_Tuple).Tuple
		array := make([]any, len(tuple.Elements))
		for i, x := range tuple.Elements {
			array[i] = argValue(x, tuple.Name)
		}
	case *evm_messages.ArgumentValue_String_:
		val = value.Value.(*evm_messages.ArgumentValue_String_).String_
	}

	return map[string]any{
		"Type":     value.Type,
		"IsHashed": value.IsHashed,
		"Value":    val,
		"Name":     name,
	}
}

func (contract *ContractInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"Address": encodeTronAddress(contract.Address),
		"Type":    contract.Type,
		"TypeUrl": contract.TypeUrl,
	})
}

func (withdraw *RewardWithdraw) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"Receiver": encodeTronAddress(withdraw.Receiver),
		"Amount":   withdraw.Amount,
	})
}

func encodeTronAddress(byteAddr []byte) string {
	if len(byteAddr) < 5 {
		return ""
	}
	prefix := []byte{'A'}
	byteAddr = append(prefix, byteAddr...)

	h1 := sha256.New()
	h1.Write(byteAddr)

	h2 := sha256.New()
	h2.Write(h1.Sum(nil))
	hashBytes := h2.Sum(nil)

	return base58.Encode(append(byteAddr, hashBytes[0:4]...))
}

func boolValue(v *bool) bool {
	if v == nil {
		return false
	}
	return *v
}

func TradeJsonStructure(trade *evm_messages.DexTrade) map[string]any {
	fees := make([]map[string]any, len(trade.Fees))
	for i, fee := range trade.Fees {
		fees[i] = FeeJsonStructure(fee)
	}
	return map[string]any{
		"Dex":              DexJsonStructure(trade.Dex),
		"Buy":              TradeSideJsonStructure(trade.Buy),
		"Sell":             TradeSideJsonStructure(trade.Sell),
		"Sender":           encodeTronAddress(trade.Sender),
		"Success":          trade.Success,
		"Fees":             fees,
		"Index":            trade.Index,
		"TransactionIndex": trade.TransactionIndex,
		"CallIndex":        trade.CallIndex,
		"LogIndex":         trade.LogIndex,
		"HasLog":           trade.HasLog,
	}
}

func TradeSideJsonStructure(side *evm_messages.TradeSide) any {
	assets := make([]map[string]any, len(side.Assets))
	for i, asset := range side.Assets {
		assets[i] = AssetJsonStructure(asset)
	}
	return map[string]any{
		"Seller":  encodeTronAddress(side.Seller),
		"Buyer":   encodeTronAddress(side.Buyer),
		"OrderId": hex.EncodeToString(side.OrderId),
		"Assets":  assets,
	}
}

func AssetJsonStructure(asset *evm_messages.TradeAsset) map[string]any {
	return map[string]any{
		"Currency": CurrencyJsonStructure(asset.Currency),
		"Amount":   encoder.BytesToBigInt(asset.Amount, false).Text(10),
		"Id":       encoder.BytesToBigInt(asset.Id, false).Text(10),
		"URI":      asset.URI,
		"Data":     hex.EncodeToString(asset.Data),
	}
}

func CurrencyJsonStructure(currency *evm_messages.TokenInfo) map[string]any {
	if currency == nil {
		return nil
	}
	return map[string]any{
		"SmartContract": encodeTronAddress(currency.SmartContract),
		"Delegated":     currency.Delegated,
		"DelegatedTo":   encodeTronAddress(currency.DelegatedTo),
		"ProtocolName":  currency.ProtocolName,
		"Name":          currency.Name,
		"Symbol":        currency.Symbol,
		"Decimals":      currency.Decimals,
		"HasURI":        currency.HasURI,
		"Fungible":      currency.Fungible,
		"AssetId":       currency.AssetId,
	}
}

func DexJsonStructure(dex *evm_messages.DexInfo) any {
	currencies := make([]map[string]any, len(dex.Currencies))
	for i, c := range dex.Currencies {
		currencies[i] = CurrencyJsonStructure(c)
	}
	ucurrencies := make([]map[string]any, len(dex.UnderlyingCurrencies))
	for i, c := range dex.UnderlyingCurrencies {
		ucurrencies[i] = CurrencyJsonStructure(c)
	}
	return map[string]any{
		"SmartContract":        encodeTronAddress(dex.SmartContract),
		"Delegated":            dex.Delegated,
		"DelegatedTo":          encodeTronAddress(dex.DelegatedTo),
		"OwnerAddress":         encodeTronAddress(dex.OwnerAddress),
		"FeeRecipient":         encodeTronAddress(dex.FeeRecipient),
		"ProtocolName":         dex.ProtocolName,
		"ProtocolFamily":       dex.ProtocolFamily,
		"ProtocolVersion":      dex.ProtocolVersion,
		"Pair":                 CurrencyJsonStructure(dex.Pair),
		"Currencies":           currencies,
		"UnderlyingCurrencies": ucurrencies,
	}
}

func FeeJsonStructure(fee *evm_messages.TradeFee) map[string]any {
	return map[string]any{
		"Currency":  CurrencyJsonStructure(fee.Currency),
		"Amount":    encoder.BytesToBigInt(fee.Amount, false).Text(10),
		"Id":        encoder.BytesToBigInt(fee.Id, false).Text(10),
		"Payer":     encodeTronAddress(fee.Payer),
		"Recipient": encodeTronAddress(fee.Recipient),
	}
}
