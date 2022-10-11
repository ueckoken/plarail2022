package train2IP

type Name2Id struct {
	Trains []Trains `yaml:"trains"`
}
type Trains struct {
	Name string `yaml:"name"`
	IP   string `yaml:"ip"`
}

func GetTable() Name2Id {
	n2i := Name2Id{Trains: []Trains{{Name: "TAKAO", IP: "192.168.100.150:1111"}}}
	return n2i
}
func (t *Name2Id) SearchIp(name string) (ip string) {
	for _, i := range t.Trains {
		if name == i.Name {
			return i.IP
		}
	}
	return ""
}
