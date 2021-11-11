package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

const minSecretKeySize = 32

type Maker interface {
	CreateToken(userName string, duration time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
}

type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewPayload(userName string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        tokenID,
		Username:  userName,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}

	return payload, nil
}

var (
    ErrInvalidToken = errors.New("token is invalid")
    ErrExpiredToken = errors.New("token has expired")
)

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}

	return nil
}
