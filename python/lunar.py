#!/usr/bin/env python3
import math
import datetime

def moon_phase(year, month, day):
    # Astronomical algorithm for moon phase (rough but accurate enough)
    if month < 3:
        year -= 1
        month += 12
    a = year // 100
    b = a // 4
    c = 2 - a + b
    e = int(365.25 * (year + 4716))
    f = int(30.6001 * (month + 1))
    jd = c + day + e + f - 1524.5

    days_since_new = jd - 2451549.5
    new_moons = days_since_new / 29.53
    phase = (new_moons - int(new_moons)) * 29.53

    return phase

def phase_name(phase):
    if phase < 1.845:
        return "🌑 New Moon"
    elif phase < 5.535:
        return "🌒 Waxing Crescent"
    elif phase < 9.225:
        return "🌓 First Quarter"
    elif phase < 12.915:
        return "🌔 Waxing Gibbous"
    elif phase < 16.605:
        return "🌕 Full Moon"
    elif phase < 20.295:
        return "🌖 Waning Gibbous"
    elif phase < 23.985:
        return "🌗 Last Quarter"
    else:
        return "🌘 Waning Crescent"

def main():
    choice = input("Enter date (YYYY-MM-DD) or press Enter for today: ").strip()
    if choice:
        try:
            y, m, d = map(int, choice.split('-'))
            date = datetime.date(y, m, d)
        except:
            print("Invalid format. Using today.")
            date = datetime.date.today()
    else:
        date = datetime.date.today()

    phase = moon_phase(date.year, date.month, date.day)
    name = phase_name(phase)
    illumination = 50 - 50 * math.cos(2 * math.pi * phase / 29.53)
    print(f"\n📅 {date}")
    print(f"{name} ({illumination:.1f}% illuminated)")

if __name__ == "__main__":
    main()
