package solana_messages

import (
	"encoding/hex"
	"encoding/json"

	"github.com/akamensky/base58"
)

func (info *DexInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"ProgramAddress": base58.Encode(info.ProgramAddress),
		"ProtocolName":   info.ProtocolName,
		"ProtocolFamily": info.ProtocolFamily,
	})
}

func (market *DexMarket) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"MarketAddress": base58.Encode(market.MarketAddress),
		"BaseCurrency":  market.BaseCurrency,
		"QuoteCurrency": market.QuoteCurrency,
	})
}

func (currency *Currency) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"Name":                 currency.Name,
		"Decimals":             currency.Decimals,
		"Uri":                  currency.Uri,
		"Symbol":               currency.Symbol,
		"Native":               currency.Native,
		"Wrapped":              currency.Wrapped,
		"Fungible":             currency.Fungible,
		"Key":                  currency.Key,
		"SellerFeeBasisPoints": currency.SellerFeeBasisPoints,
		"EditionNonce":         currency.EditionNonce,
		"TokenStandard":        currency.TokenStandard,
		"ProgramAddress":       base58.Encode(currency.ProgramAddress),
		"MintAddress":          base58.Encode(currency.MintAddress),
		"MetadataAddress":      base58.Encode(currency.MetadataAddress),
		"UpdateAuthority":      base58.Encode(currency.UpdateAuthority),
		"CollectionAddress":    base58.Encode(currency.CollectionAddress),
		"VerifiedCollection":   currency.VerifiedCollection,
		"PrimarySaleHappened":  currency.PrimarySaleHappened,
		"IsMutable":            currency.IsMutable,
		"TokenCreators":        currency.TokenCreators,
	})
}

func (creator *TokenCreator) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"Address":  base58.Encode(creator.Address),
		"Verified": creator.Verified,
		"Share":    creator.Share,
	})
}

func (account *Account) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"Address":    base58.Encode(account.Address),
		"IsSigner":   account.IsSigner,
		"IsWritable": account.IsWritable,
		"Token":      account.Token,
	})
}

func (token *TokenInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"Mint":      base58.Encode(token.Mint),
		"Owner":     base58.Encode(token.Owner),
		"Decimals":  token.Decimals,
		"ProgramId": base58.Encode(token.ProgramId),
	})
}

func (instruction *ParsedIdlInstruction) MarshalJSON() ([]byte, error) {
	var data string
	if instruction.Data != nil {
		data = hex.EncodeToString(instruction.Data)
	}
	return json.Marshal(map[string]any{
		"Index":             instruction.Index,
		"Depth":             instruction.Depth,
		"CallPath":          instruction.CallPath,
		"CallerIndex":       instruction.CallerIndex,
		"AncestorIndexes":   instruction.AncestorIndexes,
		"ExternalSeqNumber": instruction.ExternalSeqNumber,
		"InternalSeqNumber": instruction.InternalSeqNumber,
		"Program":           instruction.Program,
		"Accounts":          instruction.Accounts,
		"Logs":              instruction.Logs,
		"Data":              data,
	})
}

func (program *Program) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"Address":      base58.Encode(program.Address),
		"Parsed":       program.Parsed,
		"Name":         program.Name,
		"Method":       program.Method,
		"Arguments":    program.Arguments,
		"AccountNames": program.AccountNames,
	})
}

func (arg *ParsedArgument) MarshalJSON() ([]byte, error) {
	var val any

	switch arg.Value.(type) {
	case *ParsedArgument_String_:
		val = arg.Value.(*ParsedArgument_String_).String_
	case *ParsedArgument_UInt:
		val = arg.Value.(*ParsedArgument_UInt).UInt
	case *ParsedArgument_Int:
		val = arg.Value.(*ParsedArgument_Int).Int
	case *ParsedArgument_Bool:
		val = arg.Value.(*ParsedArgument_Bool).Bool
	case *ParsedArgument_Float:
		val = arg.Value.(*ParsedArgument_Float).Float
	case *ParsedArgument_Json:
		val = arg.Value.(*ParsedArgument_Json).Json
	case *ParsedArgument_Address:
		val = base58.Encode(arg.Value.(*ParsedArgument_Address).Address)
	}

	return json.Marshal(map[string]any{
		"Name":  arg.Name,
		"Type":  arg.Type,
		"Value": val,
	})
}

func (order *DexOrder) MarshalJSON() ([]byte, error) {
	var id string
	if order.OrderId != nil {
		id = hex.EncodeToString(order.OrderId)
	}
	return json.Marshal(map[string]any{
		"OrderId": id,
		"Account": base58.Encode(order.Account),
		"Owner":   base58.Encode(order.Owner),
		"Payer":   base58.Encode(order.Payer),
		"Mint":    base58.Encode(order.Mint),
	})
}
