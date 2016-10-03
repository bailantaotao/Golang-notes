package crypto

import (
	"fmt"

	"encoding/hex"
	"golang.org/x/crypto/scrypt"
)

func ExampleScrypt() {
	inputSlice := []string{
		"020000004c1271c211717198227392b029a64a7971931d351b387bb80db027f270411e398a07046f7d4a08dd815412a8712f874a7ebf0507e3878bd24e20a3b73fd750a667d2f451eac7471b00de6659",
		"0200000011503ee6a855e900c00cfdd98f5f55fffeaee9b6bf55bea9b852d9de2ce35828e204eef76acfd36949ae56d1fbe81c1ac9c0209e6331ad56414f9072506a77f8c6faf551eac7471b00389d01",
	}
	for _, input := range inputSlice {
		h, _ := hex.DecodeString(input)
		dk, _ := scrypt.Key(h, h, 1024, 1, 1, 32)
		s := hex.EncodeToString(dk)
		fmt.Println(s)
	}

	// Output:
	// 00000000002bef4107f882f6115e0b01f348d21195dacd3582aa2dabd7985806
	// 00000000003a0d11bdd5eb634e08b7feddcfbbf228ed35d250daf19f1c88fc94
}
