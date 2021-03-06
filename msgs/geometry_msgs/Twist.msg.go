// Code generated by ros-gen-go.
// source: Twist.msg
// DO NOT EDIT!
package geometry_msgs

import (
	"io"

	"github.com/ppg/rosgo/ros"
)

type _MsgTwist struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgTwist) Text() string {
	return t.text
}

func (t *_MsgTwist) Name() string {
	return t.name
}

func (t *_MsgTwist) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgTwist) NewMessage() ros.Message {
	m := new(Twist)

	return m
}

var (
	MsgTwist = &_MsgTwist{
		`# This expresses velocity in free space broken into its linear and angular parts.
Vector3  linear
Vector3  angular
`,
		"geometry_msgs/Twist",
		"7b067cfe31b410bffd4e416af2c10eb0",
	}
)

type Twist struct {
	Linear  Vector3
	Angular Vector3
}

func (m *Twist) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "Vector3", &m.Linear); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "Vector3", &m.Angular); err != nil {
		return err
	}

	return
}

func (m *Twist) Deserialize(r io.Reader) (err error) {
	// Linear
	if err = ros.DeserializeMessageField(r, "Vector3", &m.Linear); err != nil {
		return err
	}

	// Angular
	if err = ros.DeserializeMessageField(r, "Vector3", &m.Angular); err != nil {
		return err
	}

	return
}
