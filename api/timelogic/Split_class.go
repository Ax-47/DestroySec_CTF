package timelogic

import (
	"fmt"
	"strconv"
	"strings"
)

func (x timeDaTa) spilt_time() {
	a := strings.Split(x.Timenow, " ")
	x.Timenow_split.Date = a[0]

	ax := strings.Split(a[1], ":")
	var ee error
	x.Timenow_split.Hours, ee = strconv.Atoi(ax[0])
	x.Timenow_split.Min, ee = strconv.Atoi(ax[1])
	x.Timenow_split.Sec, ee = strconv.Atoi(ax[2])
	if ee != nil {
		fmt.Println(ee)
	}

}

func (x timeDaTa) process_min60() bool {
	timenow := x.Timenow_split
	timestamp := x.Timestamp_split

	if timenow.Min >= 56 || timenow.Min <= 60 {
		timestamp.Data_min = timestamp.Min + 5 - 60 //61-60=1
		//1 2 3 4                        //-57 -56 -55 -54
		return timestamp.Data_min <= timenow.Min*-1

	} else if timenow.Min <= 4 {

		timestamp.Data_min = timestamp.Min + 5 - 60
		return timestamp.Data_min <= timenow.Min
	} else {
		return false
	}
}
func (x timeDaTa) process() bool {
	timestamp := x.Timestamp_split
	timenow := x.Timenow_split

	if timestamp.Date == timenow.Date {
		if timestamp.Hours == timenow.Hours ||
			timestamp.Hours+1 == timenow.Hours {

			if timestamp.Min > 55 || timestamp.Min < 61 {

				return x.process_min60()
			} else {
				if timestamp.Min+5 >= timenow.Min {
					return true
				} else {
					return false
				}

			}

		} else {
			return false
		}

	} else {

		return false
	}

}
