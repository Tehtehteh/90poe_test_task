package ports

type Port struct {
	ID          string
	Name        string
	City        string
	Country     string
	Alias       []string
	Coordinates [2]float32
	Province    string
	Timezone    string
	Unlocs      []string
	Code        string
}
