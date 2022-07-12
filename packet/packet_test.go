package packet

import (
  "testing"
)

func TestSubmitDecode(t *testing.T) {
	s, err := Decode([]byte{0x2, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, '1', 'h', 'e', 'l', 'l', 'o' })
	if err != nil {
		t.Errorf("decode fail %s", err.Error())
	}

	if ns, ok := s.(*Submit); ok {
	 t.Log(string(ns.ID))
	 t.Log(string(ns.Payload))
	}

	r, err := s.Encode()

	if err != nil {
		t.Errorf("encode fail %s", err.Error())
	}

	t.Log(r)
}

func TestSubmitEncode(t *testing.T) {
	p := &Submit{
		ID: string([]byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0 ,'1'}),
		Payload: []byte("hello"),
	}
	buf, err := Encode(p)
	if err != nil {
		t.Errorf("encode fail %s", err.Error())
	}
	t.Log(string(buf))
}


func TestSubmitAckDecode(t *testing.T) {
	sc, err := Decode([]byte{0x82, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, '1', 1})
	if err != nil {
		t.Errorf("decode fail %s", err.Error())
	}

	if nsc, ok := sc.(*SubmitAck); ok {
	 t.Log(string(nsc.ID))
	 t.Log(uint8(nsc.Result))
	}

	r, err := sc.Encode()

	if err != nil {
		t.Errorf("encode fail %s", err.Error())
	}

	t.Log(r)
}

func TestConnDecode(t *testing.T) {
	c, err := Decode([]byte{0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, '1', 'h', 'e', 'l', 'l', 'o' })
	if err != nil {
		t.Errorf("decode fail %s", err.Error())
	}

	if nc, ok := c.(*Conn); ok {
	 t.Log(string(nc.ID))
	 t.Log(string(nc.Payload))
	}

	r, err := c.Encode()

	if err != nil {
		t.Errorf("encode fail %s", err.Error())
	}

	t.Log(r)
}

func TestConnAckDecode(t *testing.T) {
	ca, err := Decode([]byte{0x81, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, '1', 1})
	if err != nil {
		t.Errorf("decode fail %s", err.Error())
	}

	if nca, ok := ca.(*ConnAck); ok {
	 t.Log(string(nca.ID))
	 t.Log(uint8(nca.Result))
	}

	r, err := ca.Encode()

	if err != nil {
		t.Errorf("encode fail %s", err.Error())
	}

	t.Log(r)
}
