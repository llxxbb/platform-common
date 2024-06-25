package tag

import "strings"

const settingOmit = "omitempty"
const settingDrillSub = "sub"

type mapSetting struct {
	key      string
	omit     bool
	drillSub bool
}

// return key and omitEmpty
func getMapSetting(tag string) *mapSetting {
	split := strings.Split(tag, ",")
	rtn := &mapSetting{}
	rtn.key = split[0]
	if len(split) > 1 {
		if split[1] == settingOmit {
			rtn.omit = true
		}
		if split[1] == settingDrillSub {
			rtn.drillSub = true
		}
	}
	return rtn
}
