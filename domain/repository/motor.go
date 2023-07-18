package repository

import "fmt"

type PresetPose string

const (
	LOOK_UP   PresetPose = "look-up"
	LOOK_DOWN PresetPose = "look-down"
	THINK     PresetPose = "think"
	TRACK     PresetPose = "track"
)

type PoseBehavior struct {
	DoTime int64      `json:"do_time"`
	Pose   PresetPose `json:"pose"`
}

type IMotorRepository interface {
	SetPosture(poses []PoseBehavior) error
}

func ParsePresetPose(poseStr string) (PresetPose, error) {
	switch poseStr {
	case string(LOOK_UP):
		return LOOK_UP, nil
	case string(LOOK_DOWN):
		return LOOK_DOWN, nil
	case string(THINK):
		return THINK, nil
	case string(TRACK):
		return TRACK, nil
	default:
		return "", fmt.Errorf("invalid PresetPose: %s", poseStr)
	}
}
