package tasks

type TaskData struct {
	ID       string `json:"ID"`
	Name     string `json:"Name"`
	Priority int    `json:"Priority"`
	Status   string `json:"Status"`
	Result   string `json:"result"`
}
