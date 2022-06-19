package timelogic

func Logic_time_lteq(timestamp, timenow string) bool {
	var XA timeDaTa
	XA.Timestamp = timestamp
	XA.Timenow = timenow

	XA.spilt_time()
	return XA.process()

}
