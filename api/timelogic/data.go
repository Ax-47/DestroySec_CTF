package timelogic

type timeDaTa struct {
	Timestamp       string
	Timenow         string
	Timestamp_split struct {
		Data_min int
		Date     string
		Hours    int
		Min      int
		Sec      int
	}
	Timenow_split struct {
		Data_min int
		Date     string
		Hours    int
		Min      int
		Sec      int
	}
	Timeout bool
}
