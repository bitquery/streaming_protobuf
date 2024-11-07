package solana_messages

import (
	"encoding/json"
	"github.com/mr-tron/base58"
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
