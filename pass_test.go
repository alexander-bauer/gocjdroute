package main

import (
	"testing"
	"cjdngo"
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
	err = cjdngo.WriteConf("./temp.conf", *conf)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Wrote output config file to temp.conf")
}
