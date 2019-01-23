package kucoin

import (
	"encoding/json"
	"testing"
)

func TestApiService_CreateDepositAddress(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.CreateDepositAddress("KCS")
	if err != nil {
		t.Fatal(err)
	}
	a := &DepositAddressModel{}
	if err := rsp.ReadData(a); err != nil {
		t.Fatal(err)
	}

	b, _ := json.Marshal(a)
	t.Log(string(b))

	switch {
	case a.Address == "":
		t.Error("Empty key 'address'")
	case a.Memo == "":
		t.Error("Empty key 'memo'")
	}
}

func TestApiService_DepositAddresses(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.DepositAddresses("BTC")
	if err != nil {
		t.Fatal(err)
	}
	as := DepositAddressesModel{}
	if err := rsp.ReadData(&as); err != nil {
		t.Fatal(err)
	}

	for _, a := range as {
		b, _ := json.Marshal(a)
		t.Log(string(b))
		switch {
		case a.Address == "":
			t.Error("Empty key 'address'")
		case a.Memo == "":
			t.Error("Empty key 'memo'")
		}
	}
}

func TestApiService_Deposits(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.Deposits("", "", 0, 0)
	if err != nil {
		t.Fatal(err)
	}
	ds := DepositsModel{}
	doPaginationTest(t, rsp, &ds)

	for _, d := range ds {
		b, _ := json.Marshal(d)
		t.Log(string(b))
		switch {
		case d.Address == "":
			t.Error("Empty key 'address'")
		case d.Memo == "":
			t.Error("Empty key 'memo'")
		case d.Amount == 0:
			t.Error("Empty key 'amount'")
		case d.Currency == "":
			t.Error("Empty key 'currency'")
		case d.WalletTxId == "":
			t.Error("Empty key 'walletTxId'")
		case d.Status == "":
			t.Error("Empty key 'status'")
		case d.CreatedAt == 0:
			t.Error("Empty key 'createdAt'")
		case d.UpdatedAt == 0:
			t.Error("Empty key 'updatedAt'")
		}
	}
}
