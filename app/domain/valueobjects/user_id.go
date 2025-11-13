package valueobjects

import (
	"database/sql/driver"
	"errors"
)

// エラー定義
var ErrInvalidUserIDRange = errors.New("UserIDは3000-9999の範囲である必要があります")

type UserID struct {
	value uint
}

// Go → DB（書き込み時）
func (u UserID) Value() (driver.Value, error) {
    return int64(u.value), nil
}

// DB → Go（読み取り時）
func (u *UserID) Scan(src interface{}) error {
    switch v := src.(type) {
    case int64:
        u.value = uint(v)
        return nil
    case uint:
        u.value = v
        return nil
    default:
        return errors.New("cannot scan UserID")
    }
}

// ゲッター
func (u UserID) Uint() uint {
    return u.value
}

func NewUserID(v uint) (*UserID, error) {
	if v < 3000 || v > 9999 {
        return nil, ErrInvalidUserIDRange
    }
    return &UserID{value: v}, nil
}
