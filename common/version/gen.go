package version

import (
	"fmt"
	"strconv"
	"strings"
)

type VERSION_UPGRADE_TYPE int

const (
	VERSION_UPGRADE_TYPE_MAJOR VERSION_UPGRADE_TYPE = 1
	VERSION_UPGRADE_TYPE_MINOR VERSION_UPGRADE_TYPE = 2
	VERSION_UPGRADE_TYPE_PATCH VERSION_UPGRADE_TYPE = 3
)

// GenNext 生成下一个版本号
func GenNext(version string) string {
	// 去掉前缀 "v" 并分割成三个部分
	parts := strings.Split(version[1:], ".")
	major, _ := strconv.Atoi(parts[0])
	minor, _ := strconv.Atoi(parts[1])
	patch, _ := strconv.Atoi(parts[2])

	// 根据当前版本号的大小，更新各个部分的值
	if patch < 9 {
		patch++
	} else {
		patch = 0
		if minor < 9 {
			minor++
		} else {
			minor = 0
			major++
		}
	}
	// 生成新的版本号
	nextVersion := fmt.Sprintf("v%d.%d.%d", major, minor, patch)
	return nextVersion
}

// GenUpgrade 生成版本号
func GenUpgrade(version string, upgradeType VERSION_UPGRADE_TYPE) string {
	// 去掉前缀 "v" 并分割成三个部分
	parts := strings.Split(version[1:], ".")
	major, _ := strconv.Atoi(parts[0])
	minor, _ := strconv.Atoi(parts[1])
	patch, _ := strconv.Atoi(parts[2])

	// 根据指定的升级类型，更新相应的版本部分
	switch upgradeType {
	case VERSION_UPGRADE_TYPE_MAJOR:
		major++
		minor = 0
		patch = 0
	case VERSION_UPGRADE_TYPE_MINOR:
		minor++
		patch = 0
	case VERSION_UPGRADE_TYPE_PATCH:
		patch++
	}
	// 生成新的版本号
	nextVersion := fmt.Sprintf("v%d.%d.%d", major, minor, patch)
	return nextVersion
}
