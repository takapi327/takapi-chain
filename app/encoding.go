package app

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
)

/*
EncodingConfig 与えられたアプリで使用する具体的なエンコーディングタイプを指定します。
これは、protobufとaminoの実装間の互換性のために提供されます。
InterfaceRegistry
Protobufコーデックがgoogle.protobuf.Anyを使ってエンコードおよびデコードされたインターフェースを扱うために使用
Anyから安全に展開できるインターフェースと実装を登録するための機構を提供
Codec
SDK全体で使用されるデフォルトのコーデック
状態のエンコードとデコードに使用
TxConfig
クライアントがアプリケーション定義の具体的なトランザクションタイプを生成するために利用できるインターフェイスを定義
LegacyAmino
古いコーデック(削除予定)
*/
type EncodingConfig struct {
	InterfaceRegistry types.InterfaceRegistry
	Codec             codec.Codec
	TxConfig          client.TxConfig
	LegacyAmino       *codec.LegacyAmino
}

func makeEncodingConfig() EncodingConfig {
	amino := codec.NewLegacyAmino()
	interfaceRegistry := types.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(interfaceRegistry)
	txCfg := tx.NewTxConfig(cdc, tx.DefaultSignModes)

	return EncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Codec:             cdc,
		TxConfig:          txCfg,
		LegacyAmino:       amino,
	}
}

func MakeEncodingConfig(moduleBasics module.BasicManager) EncodingConfig {
	encodingConfig := makeEncodingConfig()
	std.RegisterLegacyAminoCodec(encodingConfig.LegacyAmino)
	std.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	moduleBasics.RegisterLegacyAminoCodec(encodingConfig.LegacyAmino)
	moduleBasics.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	return encodingConfig
}
