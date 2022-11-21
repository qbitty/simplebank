package gapi

import (
	db "pro.qbitty/simplebank/db/sqlc"
	"pro.qbitty/simplebank/pb"
	"pro.qbitty/simplebank/token"
	"pro.qbitty/simplebank/util"
)

type Server struct {
	pb.UnimplementedSimpleBankServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	maker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, err
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: maker,
	}

	return server, nil
}
