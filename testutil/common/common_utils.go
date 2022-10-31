package common

import (
	"context"

	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const paramIndicator = "###"

type TestContext interface {
	GetContext() sdk.Context
	GetWrappedContext() context.Context
}

func UnmarshalJsonFileWithParams(file string, v any, params map[string]string) {
	jsonFileMinter, _ := os.Open(file)
	byteValueMinter, _ := ioutil.ReadAll(jsonFileMinter)
	jsonData := string(byteValueMinter)
	for pKey, pVal := range params {
		jsonData = strings.ReplaceAll(jsonData, paramIndicator+pKey+paramIndicator, pVal)
	}
	byteValueMinter = []byte(jsonData)
	json.Unmarshal(byteValueMinter, v)
}

func UnmarshalJsonFile(file string, v any) {
	UnmarshalJsonFileWithParams(file, v, nil)
}

func CheckInvariant(t *testing.T, ctx sdk.Context, invariant sdk.Invariant, failed bool, message string) {
	msg, wasFailed := invariant(ctx)
	require.EqualValues(t, failed, wasFailed)
	require.EqualValues(t, message, msg)
}