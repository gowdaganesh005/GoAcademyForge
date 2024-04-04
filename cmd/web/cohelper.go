package main

func testRemarks(p float32) string {
	var remarks string
	if p > 90 {
		remarks = "ExtraOrdinary"
	} else if p >= 80 {
		remarks = "Excellent"
	} else if p >= 70 {
		remarks = "Good"
	} else if p >= 60 {
		remarks = "Keep it up but dont settle"
	} else if p >= 50 {
		remarks = "Keep working hard "
	} else if p >= 40 {
		remarks = "Can do better "
	} else {
		remarks = "need attention"
	}
	return remarks
}
