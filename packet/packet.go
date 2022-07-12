package packet

import (
  "fmt"
	"bytes"
)

const (
  CommandConn    = iota + 0x01
  CommandSubmit
)

const (
  CommandConnAck   = iota + 0x81
  CommandSubmitAck
)

type Packet interface {
  Decode([]byte) error     // []byte -> struct
  Encode() ([]byte, error) // struct -> []byte
}

type Conn struct {
  ID string
  Payload []byte
}


func (s *Conn) Decode(pktBody []byte) error {
  s.ID = string(pktBody[:8])
  s.Payload = pktBody[8:]
  return nil
}

func (s *Conn) Encode() ([]byte, error) {
  return bytes.Join([][]byte{[]byte(s.ID[:8]), s.Payload}, nil), nil
}

type ConnAck struct {
  ID string
  Result uint8
}

func (s *ConnAck) Decode(pktBody []byte) error {
  s.ID = string(pktBody[0:8])
  s.Result = uint8(pktBody[8])
  return nil
}

func (s *ConnAck) Encode() ([]byte, error) {
  return bytes.Join([][]byte{[]byte(s.ID[:8]),[]byte{s.Result}}, nil), nil
}

type Submit struct {
  ID string
  Payload []byte
}


func (s *Submit) Decode(pktBody []byte) error {
  s.ID = string(pktBody[:8])
  s.Payload = pktBody[8:]
  return nil
}

func (s *Submit) Encode() ([]byte, error) {
  return bytes.Join([][]byte{[]byte(s.ID[:8]), s.Payload}, nil), nil
}

type SubmitAck struct {
  ID string
  Result uint8
}

func (s *SubmitAck) Decode(pktBody []byte) error {
  s.ID = string(pktBody[0:8])
  s.Result = uint8(pktBody[8])
  return nil
}

func (s *SubmitAck) Encode() ([]byte, error) {
  return bytes.Join([][]byte{[]byte(s.ID[:8]),[]byte{s.Result}}, nil), nil
}

func Decode(packet []byte) (Packet, error) {
  commandID := packet[0]
  pktBody := packet[1:]

  switch commandID {
  case CommandConn:
    c := Conn{}
    err := c.Decode(pktBody)
    if err != nil {
      return nil, err
    }
    return &c, nil
  case CommandConnAck:
    c := ConnAck{}
    err := c.Decode(pktBody)
    if err != nil {
      return nil, err
    }
    return &c, nil
  case CommandSubmit:
    s := Submit{}
    err := s.Decode(pktBody)
    if err != nil {
      return nil, err
    }
    return &s, nil
  case CommandSubmitAck:
    s := SubmitAck{}
    err := s.Decode(pktBody)
    if err != nil {
      return nil, err
    }
    return &s, nil
  default:
    return nil, fmt.Errorf("unknown commandID [%d]", commandID)
  }
}

func Encode(p Packet) ([]byte, error) {
  var commandID uint8
  var pktBody []byte
  var err error

  switch t := p.(type) {
  case *Submit:
    commandID = CommandSubmit
    pktBody, err = p.Encode()
    if err != nil {
      return nil, err
    }
  case *SubmitAck:
    commandID = CommandSubmitAck
    pktBody, err = p.Encode()
    if err != nil {
      return nil, err
    }
  case *Conn:
    commandID = CommandConn
    pktBody, err = p.Encode()
    if err != nil {
      return nil, err
    }
  case *ConnAck:
    commandID = CommandConnAck
    pktBody, err = p.Encode()
    if err != nil {
      return nil, err
    }
  default:
    return nil, fmt.Errorf("unknown type [%s]", t)
  }
  return bytes.Join([][]byte{[]byte{commandID}, pktBody}, nil), nil
}
