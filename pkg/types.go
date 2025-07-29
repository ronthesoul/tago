package pkg

type Task struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Command  string `json:"command"`
	Time     string `json:"time"`
	Desc     string `json:"desc"`
	Complete bool   `json:"complete"`
}
