// Code generated by ros-gen-go.
// source: AccelWithCovariance.msg
// DO NOT EDIT!
package geometry_msgs

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/ppg/rosgo/ros"
)

type _MsgAccelWithCovariance struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgAccelWithCovariance) Text() string {
	return t.text
}

func (t *_MsgAccelWithCovariance) Name() string {
	return t.name
}

func (t *_MsgAccelWithCovariance) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgAccelWithCovariance) NewMessage() ros.Message {
	m := new(AccelWithCovariance)

	return m
}

var (
	MsgAccelWithCovariance = &_MsgAccelWithCovariance{
		`# This expresses acceleration in free space with uncertainty.

Accel accel

# Row-major representation of the 6x6 covariance matrix
# The orientation parameters use a fixed-axis representation.
# In order, the parameters are:
# (x, y, z, rotation about X axis, rotation about Y axis, rotation about Z axis)
float64[36] covariance
`,
		"geometry_msgs/AccelWithCovariance",
		"c48b3713c98db26549cd6f7fa86d1e4d",
	}
)

type AccelWithCovariance struct {
	Accel      Accel
	Covariance [36]float64
}

func (m *AccelWithCovariance) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "Accel", &m.Accel); err != nil {
		return err
	}

	// Write size little endian
	err = binary.Write(w, binary.LittleEndian, uint32(len(m.Covariance)))
	if err != nil {
		return fmt.Errorf("could not write array length: %s", err)
	}
	for _, elem := range m.Covariance {
		if err = ros.SerializeMessageField(w, "float64", &elem); err != nil {
			return err
		}
	}

	return
}

func (m *AccelWithCovariance) Deserialize(r io.Reader) (err error) {
	// Accel
	if err = ros.DeserializeMessageField(r, "Accel", &m.Accel); err != nil {
		return err
	}

	// Covariance
	{
		// Read size little endian
		var size uint32
		if err = binary.Read(r, binary.LittleEndian, &size); err != nil {
			return fmt.Errorf("cannot read array size for Covariance: %s", err)
		}
		if size > 36 {
			return fmt.Errorf("array size for Covariance too large: expected=36, got=%d", size)
		}
		for i := 0; i < int(size); i++ {
			if err = ros.DeserializeMessageField(r, "float64", &m.Covariance[i]); err != nil {
				return err
			}
		}
	}

	return
}