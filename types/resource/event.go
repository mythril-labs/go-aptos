package resource

type Event struct {
	Counter string `json:"counter"`
	GUID    struct {
		ID struct {
			Addr        string `json:"addr"`
			CreationNum string `json:"creation_num"`
		} `json:"id"`
	} `json:"guid"`
}
