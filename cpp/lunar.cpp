#include <iostream>
#include <cmath>
#include <ctime>

using namespace std;

double moonPhase(int year, int month, int day) {
    if (month < 3) {
        year--;
        month += 12;
    }
    int a = year / 100;
    int b = a / 4;
    int c = 2 - a + b;
    int e = int(365.25 * (year + 4716));
    int f = int(30.6001 * (month + 1));
    double jd = c + day + e + f - 1524.5;
    double daysSinceNew = jd - 2451549.5;
    double newMoons = daysSinceNew / 29.53;
    double phase = (newMoons - floor(newMoons)) * 29.53;
    return phase;
}

const char* phaseName(double phase) {
    if (phase < 1.845) return "🌑 New Moon";
    if (phase < 5.535) return "🌒 Waxing Crescent";
    if (phase < 9.225) return "🌓 First Quarter";
    if (phase < 12.915) return "🌔 Waxing Gibbous";
    if (phase < 16.605) return "🌕 Full Moon";
    if (phase < 20.295) return "🌖 Waning Gibbous";
    if (phase < 23.985) return "🌗 Last Quarter";
    return "🌘 Waning Crescent";
}

double illumination(double phase) {
    return 50 - 50 * cos(2 * M_PI * phase / 29.53);
}

int main() {
    int year, month, day;
    cout << "Enter date (YYYY MM DD) or 0 for today: ";
    cin >> year;
    if (year == 0) {
        time_t t = time(nullptr);
        tm* now = localtime(&t);
        year = now->tm_year + 1900;
        month = now->tm_mon + 1;
        day = now->tm_mday;
    } else {
        cin >> month >> day;
    }
    double phase = moonPhase(year, month, day);
    cout << "\n📅 " << year << "-" << month << "-" << day << endl;
    cout << phaseName(phase) << " (" << illumination(phase) << "% illuminated)" << endl;
    return 0;
}
