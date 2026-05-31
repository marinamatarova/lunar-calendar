package main

import (
    "fmt"
    "math"
    "time"
)

func moonPhase(year, month, day int) float64 {
    y, m := year, month
    if m < 3 {
        y--
        m += 12
    }
    a := y / 100
    b := a / 4
    c := 2 - a + b
    e := int(365.25 * float64(y+4716))
    f := int(30.6001 * float64(m+1))
    jd := float64(c) + float64(day) + float64(e) + float64(f) - 1524.5
    daysSinceNew := jd - 2451549.5
    newMoons := daysSinceNew / 29.53
    phase := (newMoons - math.Floor(newMoons)) * 29.53
    return phase
}

func phaseName(phase float64) (string, string) {
    switch {
    case phase < 1.845:
        return "🌑 New Moon", "New Moon"
    case phase < 5.535:
        return "🌒 Waxing Crescent", "Waxing Crescent"
    case phase < 9.225:
        return "🌓 First Quarter", "First Quarter"
    case phase < 12.915:
        return "🌔 Waxing Gibbous", "Waxing Gibbous"
    case phase < 16.605:
        return "🌕 Full Moon", "Full Moon"
    case phase < 20.295:
        return "🌖 Waning Gibbous", "Waning Gibbous"
    case phase < 23.985:
        return "🌗 Last Quarter", "Last Quarter"
    default:
        return "🌘 Waning Crescent", "Waning Crescent"
    }
}

func illumination(phase float64) float64 {
    return 50 - 50*math.Cos(2*math.Pi*phase/29.53)
}

func main() {
    var year, month, day int
    fmt.Print("Enter date (YYYY MM DD) or 0 for today: ")
    fmt.Scan(&year)
    if year == 0 {
        now := time.Now()
        year, month, day = now.Year(), int(now.Month()), now.Day()
    } else {
        fmt.Scan(&month, &day)
    }
    phase := moonPhase(year, month, day)
    emoji, name := phaseName(phase)
    illum := illumination(phase)
    fmt.Printf("\n📅 %d-%02d-%02d\n", year, month, day)
    fmt.Printf("%s %s (%.1f%% illuminated)\n", emoji, name, illum)
}
