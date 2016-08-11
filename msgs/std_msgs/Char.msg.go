// Code generated by ros-gen-go.
// source: Char.msg
// DO NOT EDIT!
package std_msgs

import (
	"io"

	"github.com/ppg/rosgo/ros"
)

type _MsgChar struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgChar) Text() string {
	return t.text
}

func (t *_MsgChar) Name() string {
	return t.name
}

func (t *_MsgChar) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgChar) NewMessage() ros.Message {
	m := new(Char)

	return m
}

var (
	MsgChar = &_MsgChar{
		`char data`,
		"std_msgs/Char",
		"1bf77f25acecdedba0e224b162199717",
	}
)

type Char struct {
	Data uint8
}

func (m *Char) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "char", &m.Data); err != nil {
		return err
	}

	return
}

func (m *Char) Deserialize(r io.Reader) (err error) {
	// Data
	if err = ros.DeserializeMessageField(r, "char", &m.Data); err != nil {
		return err
	}

	return
}