// statfunc
package main

func average(patients [][]string, param int) float64 {
	pLen := len(patients)
	var sumvalues float64 = 0
	for i := 0; i < pLen; i++ {
		sumvalues += float64(convint(patients[i][param]))
	}
	return sumvalues / float64(pLen)
}

func percent(patients [][]string, match string, param int) float64 {
	pLen := len(patients)
	targetvalue := 0

	var sumvalues float64 = 0
	for i := 0; i < pLen; i++ {
		if match == patients[i][param] {
			targetvalue += 1
		}
		sumvalues += float64(convint(patients[i][param]))
	}
	return float64(targetvalue) / float64(pLen) * 100
}
