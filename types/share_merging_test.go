package types

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var block825320 = [][]byte{
	{0, 0, 19, 55, 190, 239, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 19, 55, 190, 239, 0, 0, 51, 95, 81, 91, 201, 18, 0, 26, 102, 8, 2, 18, 32, 109, 84, 62, 245, 35, 11, 224, 40, 216, 216, 200, 52, 44, 108, 38, 229, 109, 25, 161, 3, 174, 152, 149, 8, 145, 189, 18, 224, 248, 226, 54, 164, 26, 64, 123, 241, 11, 51, 99, 128, 85, 240, 93, 253, 222, 148, 179, 123, 223, 14, 70, 242, 148, 108, 208, 29, 53, 165, 240, 72, 175, 7, 133, 241, 84, 51, 125, 2, 193, 121, 104, 247, 240, 107, 108, 149, 90, 192, 33, 49, 126, 62, 240, 217, 46, 144, 154, 232, 84, 191, 68, 183, 41, 13, 152, 135, 150, 8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 19, 55, 190, 239, 0, 0, 128, 4, 10, 248, 1, 10, 2, 8, 11, 18, 8, 0, 0, 19, 55, 190, 239, 0, 0, 24, 3, 32, 195, 165, 223, 146, 6, 42, 32, 109, 84, 62, 245, 35, 11, 224, 40, 216, 216, 200, 52, 44, 108, 38, 229, 109, 25, 161, 3, 174, 152, 149, 8, 145, 189, 18, 224, 248, 226, 54, 164, 50, 32, 194, 179, 27, 71, 189, 78, 94, 144, 37, 72, 136, 24, 173, 86, 60, 246, 124, 110, 92, 129, 232, 1, 170, 223, 129, 61, 93, 203, 214, 123, 240, 63, 58, 32, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 66, 32, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 74, 32, 181, 30, 241, 62, 153, 1, 2, 167, 191, 43, 205, 63, 72, 133, 65, 180, 57, 184, 220, 42, 53, 57, 159, 185, 232, 156, 147, 162, 78, 39, 75, 114, 82, 32, 227, 176, 196, 66, 152, 252, 28, 20, 154, 251, 244, 200, 153, 111, 185, 36, 39, 174, 65, 228, 100, 155, 147, 76, 164, 149, 153, 27, 120, 82, 184, 85, 90, 20, 95, 213, 143, 182, 222, 221, 86, 187, 46, 102, 147, 209, 254, 203, 235},
}

func TestParseMsgs(t *testing.T) {
	tests := []struct {
		name    string
		args    [][]byte
		want    Messages
		wantErr assert.ErrorAssertionFunc
	}{
		{"devnet_block_825320", block825320, Messages{}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseMsgs(tt.args)
			if !tt.wantErr(t, err, fmt.Sprintf("ParseMsgs(%v)", tt.args)) {
				return
			}
			assert.Equalf(t, tt.want, got, "ParseMsgs(%v)", tt.args)
		})
	}
}
