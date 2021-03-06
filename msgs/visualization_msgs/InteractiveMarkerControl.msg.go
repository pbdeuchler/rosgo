// Code generated by ros-gen-go.
// source: InteractiveMarkerControl.msg
// DO NOT EDIT!
package visualization_msgs

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/ppg/rosgo/msgs/geometry_msgs"
	"github.com/ppg/rosgo/ros"
)

type _MsgInteractiveMarkerControl struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgInteractiveMarkerControl) Text() string {
	return t.text
}

func (t *_MsgInteractiveMarkerControl) Name() string {
	return t.name
}

func (t *_MsgInteractiveMarkerControl) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgInteractiveMarkerControl) NewMessage() ros.Message {
	m := new(InteractiveMarkerControl)

	return m
}

var (
	MsgInteractiveMarkerControl = &_MsgInteractiveMarkerControl{
		`# Represents a control that is to be displayed together with an interactive marker

# Identifying string for this control.
# You need to assign a unique value to this to receive feedback from the GUI
# on what actions the user performs on this control (e.g. a button click).
string name


# Defines the local coordinate frame (relative to the pose of the parent
# interactive marker) in which is being rotated and translated.
# Default: Identity
geometry_msgs/Quaternion orientation


# Orientation mode: controls how orientation changes.
# INHERIT: Follow orientation of interactive marker
# FIXED: Keep orientation fixed at initial state
# VIEW_FACING: Align y-z plane with screen (x: forward, y:left, z:up).
uint8 INHERIT = 0 
uint8 FIXED = 1
uint8 VIEW_FACING = 2

uint8 orientation_mode

# Interaction mode for this control
# 
# NONE: This control is only meant for visualization; no context menu.
# MENU: Like NONE, but right-click menu is active.
# BUTTON: Element can be left-clicked.
# MOVE_AXIS: Translate along local x-axis.
# MOVE_PLANE: Translate in local y-z plane.
# ROTATE_AXIS: Rotate around local x-axis.
# MOVE_ROTATE: Combines MOVE_PLANE and ROTATE_AXIS.
uint8 NONE = 0 
uint8 MENU = 1
uint8 BUTTON = 2
uint8 MOVE_AXIS = 3 
uint8 MOVE_PLANE = 4
uint8 ROTATE_AXIS = 5
uint8 MOVE_ROTATE = 6
# "3D" interaction modes work with the mouse+SHIFT+CTRL or with 3D cursors.
# MOVE_3D: Translate freely in 3D space.
# ROTATE_3D: Rotate freely in 3D space about the origin of parent frame.
# MOVE_ROTATE_3D: Full 6-DOF freedom of translation and rotation about the cursor origin.
uint8 MOVE_3D = 7
uint8 ROTATE_3D = 8
uint8 MOVE_ROTATE_3D = 9

uint8 interaction_mode


# If true, the contained markers will also be visible
# when the gui is not in interactive mode.
bool always_visible


# Markers to be displayed as custom visual representation.
# Leave this empty to use the default control handles.
#
# Note: 
# - The markers can be defined in an arbitrary coordinate frame,
#   but will be transformed into the local frame of the interactive marker.
# - If the header of a marker is empty, its pose will be interpreted as 
#   relative to the pose of the parent interactive marker.
Marker[] markers


# In VIEW_FACING mode, set this to true if you don't want the markers
# to be aligned with the camera view point. The markers will show up
# as in INHERIT mode.
bool independent_marker_orientation


# Short description (< 40 characters) of what this control does,
# e.g. "Move the robot". 
# Default: A generic description based on the interaction mode
string description
`,
		"visualization_msgs/InteractiveMarkerControl",
		"7111ea3d3e9c0ded6ccecb4dc9532391",
	}
)

type InteractiveMarkerControl struct {
	Name                         string
	Orientation                  geometry_msgs.Quaternion
	OrientationMode              uint8
	InteractionMode              uint8
	AlwaysVisible                bool
	Markers                      []Marker
	IndependentMarkerOrientation bool
	Description                  string
}

func (m *InteractiveMarkerControl) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "string", &m.Name); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "geometry_msgs/Quaternion", &m.Orientation); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "uint8", &m.OrientationMode); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "uint8", &m.InteractionMode); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "bool", &m.AlwaysVisible); err != nil {
		return err
	}

	// Write size little endian
	err = binary.Write(w, binary.LittleEndian, uint32(len(m.Markers)))
	if err != nil {
		return fmt.Errorf("could not write array length: %s", err)
	}
	for _, elem := range m.Markers {
		if err = ros.SerializeMessageField(w, "Marker", &elem); err != nil {
			return err
		}
	}

	if err = ros.SerializeMessageField(w, "bool", &m.IndependentMarkerOrientation); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "string", &m.Description); err != nil {
		return err
	}

	return
}

func (m *InteractiveMarkerControl) Deserialize(r io.Reader) (err error) {
	// Name
	if err = ros.DeserializeMessageField(r, "string", &m.Name); err != nil {
		return err
	}

	// Orientation
	if err = ros.DeserializeMessageField(r, "geometry_msgs/Quaternion", &m.Orientation); err != nil {
		return err
	}

	// OrientationMode
	if err = ros.DeserializeMessageField(r, "uint8", &m.OrientationMode); err != nil {
		return err
	}

	// InteractionMode
	if err = ros.DeserializeMessageField(r, "uint8", &m.InteractionMode); err != nil {
		return err
	}

	// AlwaysVisible
	if err = ros.DeserializeMessageField(r, "bool", &m.AlwaysVisible); err != nil {
		return err
	}

	// Markers
	{
		// Read size little endian
		var size uint32
		if err = binary.Read(r, binary.LittleEndian, &size); err != nil {
			return fmt.Errorf("cannot read array size for Markers: %s", err)
		}
		m.Markers = make([]Marker, int(size))
		for i := 0; i < int(size); i++ {
			if err = ros.DeserializeMessageField(r, "Marker", &m.Markers[i]); err != nil {
				return err
			}
		}
	}

	// IndependentMarkerOrientation
	if err = ros.DeserializeMessageField(r, "bool", &m.IndependentMarkerOrientation); err != nil {
		return err
	}

	// Description
	if err = ros.DeserializeMessageField(r, "string", &m.Description); err != nil {
		return err
	}

	return
}
