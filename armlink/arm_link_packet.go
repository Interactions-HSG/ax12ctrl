package armlink

import (
	"fmt"
)

var (
	ExtendedStop  = byte(17)
	ExtendedReset = byte(64)
	ExtendedSleep = byte(96)
)

// ArmLinkPacket holds the armlink packet parameters
type ArmLinkPacket struct {
	baseRotation            uint16
	shoulderRotation        uint16
	elbowRotation           uint16
	wristAngle              uint16
	wristRotation           uint16
	gripper                 uint16
	deltaByte               byte
	buttonByte              byte
	extendedInstructionByte byte
}

// SetExtended set the extendedInstructionByte
func (alp *ArmLinkPacket) SetExtended(e byte) {
	alp.extendedInstructionByte = e
}

func (alp *ArmLinkPacket) String() string {
	return fmt.Sprintf("Base: %v, Shoulder: %v, Elbow: %v, WristAngle: %v, WristRotation: %v, Gripper: %v, Delta: %v, Button: %v, Extended: %v", alp.baseRotation, alp.shoulderRotation, alp.elbowRotation, alp.wristAngle, alp.wristRotation, alp.gripper, alp.deltaByte, alp.buttonByte, alp.extendedInstructionByte)
}

// NewArmLinkPacket creates a new ArmLinkPacket
func NewArmLinkPacket(br, sr, er, wa, wr, g uint16, d, b, e byte) *ArmLinkPacket {
	alp := &ArmLinkPacket{
		baseRotation:            br,
		shoulderRotation:        sr,
		elbowRotation:           er,
		wristAngle:              wa,
		wristRotation:           wr,
		gripper:                 g,
		deltaByte:               d,
		buttonByte:              b,
		extendedInstructionByte: e,
	}
	return alp
}

func (alp *ArmLinkPacket) Bytes() []byte {
	baseRotationHighByte := byte((alp.baseRotation / 256) % 256)
	baseRotationLowByte := byte(alp.baseRotation % 256)
	shoulderRotationHighByte := byte((alp.shoulderRotation / 256) % 256)
	shoulderRotationLowByte := byte(alp.shoulderRotation % 256)
	elbowRotationHighByte := byte((alp.elbowRotation / 256) % 256)
	elbowRotationLowByte := byte(alp.elbowRotation % 256)
	wristAngleHighByte := byte((alp.wristAngle / 256) % 256)
	wristAngleLowByte := byte(alp.wristAngle % 256)
	wristRotationHighByte := byte((alp.wristRotation / 256) % 256)
	wristRotationLowByte := byte(alp.wristRotation % 256)
	gripperHighByte := byte((alp.gripper / 256) % 256)
	gripperLowByte := byte(alp.gripper % 256)

	checksum := ^((baseRotationHighByte +
		baseRotationLowByte +
		shoulderRotationHighByte +
		shoulderRotationLowByte +
		elbowRotationHighByte +
		elbowRotationLowByte +
		wristAngleHighByte +
		wristAngleLowByte +
		wristRotationHighByte +
		wristRotationLowByte +
		gripperHighByte +
		gripperLowByte +
		alp.deltaByte +
		alp.buttonByte +
		alp.extendedInstructionByte) % 0xff)

	return []byte{
		byte(0xff), // header
		baseRotationHighByte,
		baseRotationLowByte,
		shoulderRotationHighByte,
		shoulderRotationLowByte,
		elbowRotationHighByte,
		elbowRotationLowByte,
		wristAngleHighByte,
		wristAngleLowByte,
		wristRotationHighByte,
		wristRotationLowByte,
		gripperHighByte,
		gripperLowByte,
		alp.deltaByte,
		alp.buttonByte,
		alp.extendedInstructionByte,
		checksum,
	}
}
