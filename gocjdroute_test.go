package main

import (
	"github.com/SashaCrofter/cjdngo"
	"testing"
)

func TestPass(t *testing.T) {
	t.Log("Untagged password: " + GenPass(""))
	t.Log("Tagged 'test' password: " + GenPass("test"))
}

func TestAuthorize(t *testing.T) {
	conf, err := cjdngo.ReadConf("./example.conf")
	if err != nil {
		t.Fatal(err)
	}
	p := Authorize(conf, "test", "Maryland", "")
	t.Log("Authorized node of 'test' with password: " + p)
	p = Authorize(conf, "test_two", "Maryland", "")
	t.Log("Authorized node of 'test_two' with password: " + p)

	ConnectTo(conf, "an.ip.v4.addr:port", "password", "publickey.k", "testConnection", "Maryland", "")
	ConnectTo(conf, "an.oth.er.addr:port", "password", "publickey.k", "testConnection", "Maryland", "")

	print(ListConnectTo(conf, SearchConnectTo(conf, ""), false))

	RemoveConnectTo(conf, "an.ip.v4.addr:port")
	print("removed an.ip.v4.addr:port\n")

	print(ListConnectTo(conf, SearchConnectTo(conf, ""), false))

	err = cjdngo.WriteConf("./temp.conf", *conf)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Wrote output config file to temp.conf")
}
