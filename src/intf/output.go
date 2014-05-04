package intf

type Output interface {
	PutRecord(record Request) bool
}
