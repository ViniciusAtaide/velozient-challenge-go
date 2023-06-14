//go:build wireinject
// +build wireinject

package main

func InitializeServer() Server {
	wire.Build(NewServer)

}
