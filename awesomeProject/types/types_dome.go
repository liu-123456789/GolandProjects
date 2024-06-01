package main

type Lists interface {
	Add(index int, val any) error
	Update(index int, val any) error
	Delect(index int, val any) error
}
