package deepmind

import (
	"fmt"
	"testing"

	pbcosmos "github.com/figment-networks/proto-cosmos/pb/sf/cosmos/type/v1"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/require"
)

func TestEncode(t *testing.T) {
	str := "failed to execute message; message index: 0: failed to update grant with key \x01\x14\va\x9a\ft\x9c\xadkw@\xcf)\xb3\xf3ϭ\xf4-y\xd4\x14ۯ\xdb\xe1ӈf\xafW\xa4\xf7\xef\\\xcbM!\xa6=rC/cosmos.gov.v1.MsgVote: authorization not found"

	str2 := string([]rune(str))

	fmt.Println(len(str), str)
	fmt.Println(len(str2), str2)

	x1 := "hello world"
	x2 := string([]rune(x1))

	fmt.Println(len(x1), x1)
	fmt.Println(len(x2), x2)

	tx := &pbcosmos.TxResult{
		Result: &pbcosmos.ResponseDeliverTx{
			Log: str2,
		},
	}

	_, err := proto.Marshal(tx)
	require.NoError(t, err)
}
