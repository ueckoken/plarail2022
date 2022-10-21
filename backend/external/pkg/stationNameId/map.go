package stationNameId

import (
	"fmt"
	"ueckoken/plarail2022-external/spec"
)

func ID2Name(id int32) (string, error) {
	name, ok := spec.Stations_StationId_name[id]
	if !ok {
		return "", fmt.Errorf("station ID `%d` not found ", id)
	}
	return name, nil
}
func Name2Id(name string) (int32, error) {
	id, ok := spec.Stations_StationId_value[name]
	if !ok {
		return 0, fmt.Errorf("station `%s` not found", name)
	}
	return id, nil
}
