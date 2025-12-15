package kurdical

// julianDayNumber calculates the Julian Day Number for a Gregorian date.
func julianDayNumber(year, month, day int) int {
	a := (14 - month) / 12
	y := year + 4800 - a
	m := month + 12*a - 3
	return day + (153*m+2)/5 + 365*y + y/4 - y/100 + y/400 - 32045
}

// jdnToGregorian converts Julian Day Number to Gregorian date.
func jdnToGregorian(jdn int) (year, month, day int) {
	l := jdn + 68569
	n := (4 * l) / 146097
	l = l - (146097*n+3)/4
	i := (4000 * (l + 1)) / 1461001
	l = l - (1461*i)/4 + 31
	j := (80 * l) / 2447
	day = l - (2447*j)/80
	l = j / 11
	month = j + 2 - 12*l
	year = 100*(n-49) + i + l
	return
}

// isSolarHijriLeap determines if a Solar Hijri year is leap.
func isSolarHijriLeap(year int) bool {
	return (year+1)%4 == 0
}

// gregorianToSolarHijri converts Gregorian date to Solar Hijri.
func gregorianToSolarHijri(gYear, gMonth, gDay int) (sYear, sMonth, sDay int) {
	gJDN := julianDayNumber(gYear, gMonth, gDay)
	sJDN := gJDN - 1948087
	sYear = 1
	for {
		daysInYear := 365
		if isSolarHijriLeap(sYear) {
			daysInYear = 366
		}
		if sJDN <= daysInYear {
			break
		}
		sJDN -= daysInYear
		sYear++
	}
	monthDays := []int{31, 31, 31, 31, 31, 31, 30, 30, 30, 30, 30, 29}
	if isSolarHijriLeap(sYear) {
		monthDays[11] = 30
	}
	sMonth = 1
	for sMonth <= 12 {
		if sJDN <= monthDays[sMonth-1] {
			sDay = sJDN
			break
		}
		sJDN -= monthDays[sMonth-1]
		sMonth++
	}
	return
}

// solarHijriToGregorian converts Solar Hijri date to Gregorian.
func solarHijriToGregorian(sYear, sMonth, sDay int) (gYear, gMonth, gDay int) {
	days := 0
	for y := 1; y < sYear; y++ {
		days += 365
		if isSolarHijriLeap(y) {
			days++
		}
	}
	monthDays := []int{31, 31, 31, 31, 31, 31, 30, 30, 30, 30, 30, 29}
	if isSolarHijriLeap(sYear) {
		monthDays[11] = 30
	}
	for m := 1; m < sMonth; m++ {
		days += monthDays[m-1]
	}
	days += sDay - 1
	jdn := 1948087 + days
	gYear, gMonth, gDay = jdnToGregorian(jdn)
	return
}
