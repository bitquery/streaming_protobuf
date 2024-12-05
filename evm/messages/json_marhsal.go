package evm_messages

import (
	"encoding/hex"
	"encoding/json"
	"github.com/bitquery/streaming_protobuf/v2/pkg/util"

	"strconv"
)

func (signature *ParsedAbiSignature) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"Parsed":        signature.Parsed,
		"Name":          signature.Name,
		"Abi":           signature.Abi,
		"SignatureHash": hex.EncodeToString(signature.SignatureHash),
		"Type":          signature.Type,
	})
}

func (header *TransactionHeader) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"Hash":           util.HexEncode(header.Hash),
		"Gas":            strconv.FormatUint(header.Gas, 10),
		"Value":          util.BytesToBigInt(header.Value, false).Text(10),
		"Data":           hex.EncodeToString(header.Data),
		"Nonce":          strconv.FormatUint(header.Nonce, 10),
		"ChainId":        util.BytesToBigInt(header.ChainId, false).Text(10),
		"Cost":           util.BytesToBigInt(header.Cost, false).Text(10),
		"GasFeeCap":      util.BytesToBigInt(header.GasFeeCap, false).Text(10),
		"GasPrice":       util.BytesToBigInt(header.GasPrice, false).Text(10),
		"GasTipCap":      util.BytesToBigInt(header.GasTipCap, false).Text(10),
		"Protected":      header.Protected,
		"Type":           header.Type,
		"To":             util.HexEncode(header.To),
		"From":           util.HexEncode(header.From),
		"AccessList":     header.AccessList,
		"GasL1":          header.GasL1,
		"IsSystemTx":     header.IsSystemTx,
		"EffectiveNonce": header.EffectiveNonce,
		"SourceHash":     header.SourceHash,
		"Time":           header.Time,
		"Blob":           header.Blob,
	})
}

func (tuple *AccessTuple) MarshalJSON() ([]byte, error) {
	keys := make([]string, len(tuple.StorageKeys))
	for i, k := range tuple.StorageKeys {
		keys[i] = hex.EncodeToString(k)
	}

	return json.Marshal(map[string]any{
		"Address":     util.HexEncode(tuple.Address),
		"StorageKeys": keys,
	})
}

func (signature *Signature) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"R": hex.EncodeToString(signature.R),
		"V": hex.EncodeToString(signature.V),
		"S": hex.EncodeToString(signature.S),
	})
}

func (header *ReceiptHeader) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"Bloom":             hex.EncodeToString(header.Bloom),
		"GasUsed":           strconv.FormatUint(header.GasUsed, 10),
		"Type":              header.Type,
		"ContractAddress":   util.HexEncode(header.ContractAddress),
		"CumulativeGasUsed": strconv.FormatUint(header.CumulativeGasUsed, 10),
		"PostState":         hex.EncodeToString(header.PostState),
		"Status":            header.Status,
	})
}

func (state *ParsedStateChange) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"EnterIndex":           state.EnterIndex,
		"ExitIndex":            state.ExitIndex,
		"ChangeAfterCallIndex": state.ChangeAfterCallIndex,
		"Pc":                   state.Pc,
		"Address":              util.HexEncode(state.Address),
		"Location":             hex.EncodeToString(state.Location),
		"Value":                hex.EncodeToString(state.Value),
	})
}

func (topic *Topic) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"Index": topic.Index,
		"Hash":  hex.EncodeToString(topic.Hash),
	})
}

func (header *LogHeader) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"Index":    header.Index,
		"Address":  util.HexEncode(header.Address),
		"Data":     hex.EncodeToString(header.Data),
		"Removed":  header.Removed,
		"IsSystem": header.IsSystem,
	})
}

func (header *ParsedAbiLogHeader) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"Index":             header.Index,
		"EnterIndex":        header.EnterIndex,
		"ExitIndex":         header.ExitIndex,
		"LogAfterCallIndex": header.LogAfterCallIndex,
		"Pc":                header.Pc,
		"SmartContract":     util.HexEncode(header.SmartContract),
		"Signature":         header.Signature,
	})
}

func (trace *Trace) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"Calls":        trace.Calls,
		"CaptureStart": trace.CaptureStart,
		"CaptureEnd":   trace.CaptureEnd,
		"CaptureFault": trace.CaptureFault,
	})
}

func (header *ParsedAbiCallHeader) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"Index":         header.Index,
		"Depth":         header.Depth,
		"EnterIndex":    header.EnterIndex,
		"ExitIndex":     header.ExitIndex,
		"CallerIndex":   header.CallerIndex,
		"InternalCalls": header.InternalCalls,
		"CallPath":      header.CallPath,
		"From":          util.HexEncode(header.From),
		"To":            util.HexEncode(header.To),
		"Create":        header.Create,
		"Input":         hex.EncodeToString(header.Input),
		"Gas":           header.Gas,
		"Value":         util.BytesToBigInt(header.Value, false).Text(10),
		"Output":        hex.EncodeToString(header.Output),
		"GasUsed":       header.GasUsed,
		"Error":         header.Error,
		"Opcode":        header.Opcode,
		"SelfDestruct":  header.SelfDestruct,
		"Delegated":     header.Delegated,
		"Success":       header.Success,
		"Reverted":      header.Reverted,
		"Signature":     header.Signature,
	})
}

func (start *CaptureStart) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"From":   util.HexEncode(start.From),
		"To":     util.HexEncode(start.To),
		"Create": start.Create,
		"Input":  hex.EncodeToString(start.Input),
		"Gas":    strconv.FormatUint(start.Gas, 10),
		"Value":  util.BytesToBigInt(start.Value, false).Text(10),
	})
}

func (end *CaptureEnd) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"Output":  hex.EncodeToString(end.Output),
		"GasUsed": strconv.FormatUint(end.GasUsed, 10),
		"Error":   end.Error,
	})
}

func (fault *CaptureFault) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"Pc":       strconv.FormatUint(fault.Pc, 10),
		"Opcode":   fault.Opcode,
		"Gas":      strconv.FormatUint(fault.Gas, 10),
		"Cost":     strconv.FormatUint(fault.Cost, 10),
		"Memory":   hex.EncodeToString(fault.Memory),
		"Depth":    fault.Depth,
		"Error":    fault.Error,
		"Contract": fault.Contract,
	})
}

func (contract *Contract) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"CallerAddress": util.HexEncode(contract.CallerAddress),
		"Caller":        util.HexEncode(contract.Caller),
		"Address":       util.HexEncode(contract.Address),
		"Input":         hex.EncodeToString(contract.Input),
		"Value":         util.BytesToBigInt(contract.Value, false).Text(10),
	})
}

func (header *CaptureStateHeader) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"EnterIndex": header.EnterIndex,
		"ExitIndex":  header.ExitIndex,
		"Pc":         strconv.FormatUint(header.Pc, 10),
		"Opcode":     header.Opcode,
		"Gas":        strconv.FormatUint(header.Gas, 10),
		"Cost":       strconv.FormatUint(header.Cost, 10),
		"Depth":      strconv.FormatUint(header.Cost, 10),
		"Error":      header.Error,
		"Index":      header.Index,
	})
}

func (store *Store) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"Address":  util.HexEncode(store.Address),
		"Location": hex.EncodeToString(store.Location),
		"Value":    hex.EncodeToString(store.Value),
	})
}

func (currency *DexInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"SmartContract":        util.HexEncode(currency.SmartContract),
		"Delegated":            currency.Delegated,
		"DelegatedTo":          currency.DelegatedTo,
		"OwnerAddress":         util.HexEncode(currency.OwnerAddress),
		"FeeRecipient":         util.HexEncode(currency.FeeRecipient),
		"ProtocolName":         currency.ProtocolName,
		"ProtocolFamily":       currency.ProtocolFamily,
		"ProtocolVersion":      currency.ProtocolVersion,
		"Pair":                 currency.Pair,
		"Currencies":           currency.Currencies,
		"UnderlyingCurrencies": currency.UnderlyingCurrencies,
	})
}

func (currency *TokenInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"SmartContract": util.HexEncode(currency.SmartContract),
		"Delegated":     currency.Delegated,
		"DelegatedTo":   util.HexEncode(currency.DelegatedTo),
		"ProtocolName":  currency.ProtocolName,
		"Name":          currency.Name,
		"Symbol":        currency.Symbol,
		"Decimals":      currency.Decimals,
		"HasURI":        currency.HasURI,
		"Fungible":      currency.Fungible,
		"AssetId":       currency.AssetId,
	})
}

func (trade *DexTrade) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"Dex":     trade.Dex,
		"Buy":     trade.Buy,
		"Sell":    trade.Sell,
		"Sender":  util.HexEncode(trade.Sender),
		"Success": trade.Success,
		"Fees":    trade.Fees,
		"Index":   trade.Index,
	})
}

func (fee *TradeFee) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"Currency":  fee.Currency,
		"Amount":    util.BytesToBigInt(fee.Amount, false).Text(10),
		"Id":        hex.EncodeToString(fee.Id),
		"Payer":     util.HexEncode(fee.Payer),
		"Recipient": util.HexEncode(fee.Recipient),
	})
}

func (side *TradeSide) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"Seller":  util.HexEncode(side.Seller),
		"Buyer":   util.HexEncode(side.Buyer),
		"OrderId": hex.EncodeToString(side.OrderId),
		"Assets":  side.Assets,
	})
}

func (asset *TradeAsset) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"Currency": asset.Currency,
		"Amount":   util.BytesToBigInt(asset.Amount, false).Text(10),
		"Id":       hex.EncodeToString(asset.Id),
		"URI":      asset.URI,
		"Data":     hex.EncodeToString(asset.Data),
	})
}

func (value *ArgumentValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(value.argValue(""))
}

func (value *ArgumentValue) argValue(name string) map[string]any {
	var val any

	switch value.Value.(type) {
	case *ArgumentValue_Bytes:
		if value.Type == "address" {
			val = util.HexEncode(value.Value.(*ArgumentValue_Bytes).Bytes)
		} else {
			val = hex.EncodeToString(value.Value.(*ArgumentValue_Bytes).Bytes)
		}
	case *ArgumentValue_UInt:
		val = value.Value.(*ArgumentValue_UInt).UInt
	case *ArgumentValue_Int:
		val = value.Value.(*ArgumentValue_Int).Int
	case *ArgumentValue_Bool:
		val = value.Value.(*ArgumentValue_Bool).Bool
	case *ArgumentValue_Array:
		array := make([]any, len(value.Value.(*ArgumentValue_Array).Array.Elements))
		for i, x := range value.Value.(*ArgumentValue_Array).Array.Elements {
			array[i] = x.argValue("")
		}
	case *ArgumentValue_Tuple:
		tuple := value.Value.(*ArgumentValue_Tuple).Tuple
		array := make([]any, len(tuple.Elements))
		for i, x := range tuple.Elements {
			array[i] = x.argValue(tuple.Name)
		}
	case *ArgumentValue_String_:
		val = value.Value.(*ArgumentValue_String_).String_
	}

	return map[string]any{
		"Type":     value.Type,
		"IsHashed": value.IsHashed,
		"Value":    val,
		"Name":     name,
	}
}
