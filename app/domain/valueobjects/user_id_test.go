package valueobjects

import (
	"testing"
	"errors"
)

func TestNewUserID(t *testing.T) {
	tests := []struct {
  		name    string
  		value   uint
  		wantErr error
  	}{
  		{name: "正常系: 最小値", value: 3000, wantErr: nil},
  		{name: "正常系: 最大値", value: 9999, wantErr: nil},
  		{name: "正常系: 中間値", value: 5000, wantErr: nil},
  		{name: "異常系: 最小値未満", value: 2999, wantErr: ErrInvalidUserIDRange},
  		{name: "異常系: 最大値超過", value: 10000, wantErr: ErrInvalidUserIDRange},
  		{name: "異常系: ゼロ", value: 0, wantErr: ErrInvalidUserIDRange},
  	}

  	for _, tt := range tests {
  		t.Run(tt.name, func(t *testing.T) {
  			uid, err := NewUserID(tt.value)

  			// エラーチェック
  			if tt.wantErr != nil {
  				if !errors.Is(err, tt.wantErr) {
  					t.Errorf("NewUserID(%d) error = %v, want %v", tt.value, err, tt.wantErr)
  				}
  			} else {
  				if err != nil {
  					t.Errorf("NewUserID(%d) unexpected error: %v", tt.value, err)
  					return
  				}
  				if uid.Uint() != tt.value {
  					t.Errorf("NewUserID(%d) = %d, want %d", tt.value, uid.Uint(), tt.value)
  				}
  			}
  		})
  	}
}
