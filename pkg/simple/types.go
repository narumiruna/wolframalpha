package simple

type QueryOptions struct {
	Width    int    `json:"width"`
	Fontsize int    `json:"fontsize"`
	Units    string `json:"units"`
	Timeout  int    `json:"timeout"`
}
