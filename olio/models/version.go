package models

type Version struct {
	SkeletonVersion string `json:"serviceFrameworkVersion"`
	AppVersion      string `json:"appVersion"`
}
