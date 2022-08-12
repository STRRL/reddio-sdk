package crypto

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

const privateKey = "3c1e9550e66958296d11b60f8e8e7a7ad990d07fa65d5f7652c4a6c87d4e3cc"

type SignCase struct {
	msgHash, s, r string
}

type VerifyCase struct {
	publicKey, msgHash, s, r string
	valid                    bool
}

type GetPublicKeyCase struct {
	privateKey, publicKey string
}

func TestSign(t *testing.T) {
	cases := []SignCase{
		{
			msgHash: "1",
			r:       "3162358736122783857144396205516927012128897537504463716197279730251407200037",
			s:       "1447067116407676619871126378936374427636662490882969509559888874644844560850",
		},
		{
			msgHash: "11",
			r:       "2282960348362869237018441985726545922711140064809058182483721438101695251648",
			s:       "2905868291002627709651322791912000820756370440695830310841564989426104902684",
		},

		{
			msgHash: "223",
			r:       "2851492577225522862152785068304516872062840835882746625971400995051610132955",
			s:       "2227464623243182122770469099770977514100002325017609907274766387592987135410",
		},

		{
			msgHash: "9999",
			r:       "3551214266795401081823453828727326248401688527835302880992409448142527576296",
			s:       "2580950807716503852408066180369610390914312729170066679103651110985466032285",
		},

		{
			msgHash: "387e76d1667c4454bfb835144120583af836f8e32a516765497d23eabe16b3f",
			r:       "3518448914047769356425227827389998721396724764083236823647519654917215164512",
			s:       "3042321032945513635364267149196358883053166552342928199041742035443537684462",
		},

		{
			msgHash: "3a7e76d1697c4455bfb835144120283af236f8e32a516765497d23eabe16b2",
			r:       "2261926635950780594216378185339927576862772034098248230433352748057295357217",
			s:       "2708700003762962638306717009307430364534544393269844487939098184375356178572",
		},

		{
			msgHash: "fa5f0cd1ebff93c9e6474379a213ba111f9e42f2f1cb361b0327e0737203",
			r:       "3016953906936760149710218073693613509330129567629289734816320774638425763370",
			s:       "306146275372136078470081798635201810092238376869367156373203048583896337506",
		},

		{
			msgHash: "4c1e9550e66958296d11b60f8e8e7f7ae99dd0cfa6bd5fa652c1a6c87d4e2cc",
			r:       "3562728603055564208884290243634917206833465920158600288670177317979301056463",
			s:       "1958799632261808501999574190111106370256896588537275453140683641951899459876",
		},

		{
			msgHash: "6362b40c218fb4c8a8bd42ca482145e8513b78e00faa0de76a98ba14fc37ae8",
			r:       "3485557127492692423490706790022678621438670833185864153640824729109010175518",
			s:       "897592218067946175671768586886915961592526001156186496738437723857225288280",
		},
	}

	pk, ok := new(big.Int).SetString(privateKey, 16)
	assert.True(t, ok)

	for _, c := range cases {
		hash, ok := new(big.Int).SetString(c.msgHash, 16)
		assert.True(t, ok)
		r, s, err := Sign(pk, hash, nil)
		assert.Nil(t, err)
		assert.Equal(t, c.r, r.Text(10))
		assert.Equal(t, c.s, s.Text(10))
	}
}

func TestVerify(t *testing.T) {
	cases := []VerifyCase{
		{
			publicKey: "01ef15c18599971b7beced415a40f0c7deacfd9b0d1819e03d723d8bc943cfca",
			msgHash:   "0000000000000000000000000000000000000000000000000000000000000002",
			r:         "0411494b501a98abd8262b0da1351e17899a0c4ef23dd2f96fec5ba847310b20",
			s:         "0405c3191ab3883ef2b763af35bc5f5d15b3b4e99461d70e84c654a351a7c81b",
			valid:     true,
		},
		{
			publicKey: "077a4b314db07c45076d11f62b6f9e748a39790441823307743cf00d6597ea43",
			msgHash:   "0397e76d1667c4454bfb83514e120583af836f8e32a516765497823eabe16a3f",
			r:         "0173fd03d8b008ee7432977ac27d1e9d1a1f6c98b1a2f05fa84a21c84c44e882",
			s:         "01f2c44a7798f55192f153b4c48ea5c1241fbb69e6132cc8a0da9c5b62a4286e",
			valid:     false,
		},
	}

	for _, c := range cases {
		publicKey, ok := new(big.Int).SetString(c.publicKey, 16)
		assert.True(t, ok)
		msgHash, ok := new(big.Int).SetString(c.msgHash, 16)
		assert.True(t, ok)
		r, ok := new(big.Int).SetString(c.r, 16)
		assert.True(t, ok)
		s, ok := new(big.Int).SetString(c.s, 16)
		assert.True(t, ok)

		valid, err := Verify(publicKey, msgHash, r, s)
		assert.Nil(t, err)
		assert.Equal(t, c.valid, valid)
	}
}

func TestGetPublicKey(t *testing.T) {
	cases := []GetPublicKeyCase{
		{
			privateKey: "03c1e9550e66958296d11b60f8e8e7a7ad990d07fa65d5f7652c4a6c87d4e3cc",
			publicKey:  "077a3b314db07c45076d11f62b6f9e748a39790441823307743cf00d6597ea43",
		},
		{
			privateKey: "0000000000000000000000000000000000000000000000000000000000000012",
			publicKey:  "019661066e96a8b9f06a1d136881ee924dfb6a885239caa5fd3f87a54c6b25c4",
		},
	}

	for _, c := range cases {
		privateKey, ok := new(big.Int).SetString(c.privateKey, 16)
		assert.True(t, ok)
		expectedKey, ok := new(big.Int).SetString(c.publicKey, 16)
		assert.True(t, ok)
		publicKey, err := GetPublicKey(privateKey)
		assert.Nil(t, err)
		assert.True(t, expectedKey.Cmp(publicKey) == 0)
	}
}

func TestFFIError(t *testing.T) {
	pk, ok := new(big.Int).SetString(privateKey, 16)
	assert.True(t, ok)
	invalidHex, ok := new(big.Int).SetString("fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff", 16)
	assert.True(t, ok)
	_, _, err := Sign(pk, invalidHex, nil)
	assert.NotNil(t, err)
	assert.Equal(t, "not an invalid hex number", err.Error())
}
