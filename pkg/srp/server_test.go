package srp_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/Venafi/venafi-go-srp/srp"
)

func TestSRP_Marshal_Unmarshal_Binary(t *testing.T) {
	params, err := srp.GetParams(4096)
	require.NoError(t, err)
	I := []byte("some")
	P := []byte("randomdude123")
	s := srp.BytesFromHexString("beb25379d1a8581eb5a727673a2441ee")
	a := srp.BytesFromHexString("60975527035cf2ad1989806f0407210bc81edc04e2762a56afd529ddda2d4393")
	b := srp.BytesFromHexString("e487cb59d31ac550471e81f00f6928e01dda08e974a004f49e61f5d105284d20")

	verifier := srp.ComputeVerifier(params, s, I, P)
	client := srp.NewClient(params, s, I, P, a)
	A := client.ComputeA()

	expectedServer := srp.NewServer(params, verifier, b)
	expectedServer.SetA(A)

	expectedServerBytes, err := expectedServer.MarshalBinary()
	require.NoError(t, err)

	actualServer := &srp.SRPServer{}
	err = actualServer.UnmarshalBinary(expectedServerBytes)
	require.NoError(t, err)

	require.Equal(t, expectedServer.Params, actualServer.Params)
	require.Equal(t, expectedServer.B, actualServer.B)
	require.Equal(t, expectedServer.K, actualServer.K)
	require.Equal(t, expectedServer.Verifier, actualServer.Verifier)
	require.Equal(t, expectedServer.Secret2, actualServer.Secret2)
	require.Equal(t, expectedServer.M2, actualServer.M2)
	require.Equal(t, expectedServer.M1, actualServer.M1)

}
