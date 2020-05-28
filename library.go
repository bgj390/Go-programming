// A home exersice for http://terokarvinen.com/2020/go-programming-course-2020-w22/#h2
// Purpose is to use at least two libraries

// This program has the mid-summer eve (juhannus) 
// and its date hardcoded at j and jA respectively
// and tells pseudo-random facts about them
// in finnish, sry

package main

import (
        "fmt"
        "time"
        s "strings"
)

func main() {

        p := fmt.Println
        j := "juhannus"

        p(s.Repeat(j,4))

        now := time.Now()

        p("Tänään on", now)
        p("Sanassa", j, "on välilyönti?", s.Contains(j, " "))
        p("Siinä on", len(j), "kirjainta")
        p("Ja se huudetaan näin:", s.ToUpper(j),"!")

        jA := time.Date(2020, 06, 19, 00, 00, 00, 00, time.UTC)

        p("Sen päivämäärä on", jA)
        p("Ja viikonpäivä on...")

        switch jA.Weekday() {
                case time.Monday:
                        p("MAANANTAI")
                case time.Tuesday:
                        p("TIISTAI")
                case time.Wednesday:
                        p("KESKIVIIKKO")
                case time.Thursday:
                        p("TORSTAI")
                case time.Friday:
                        p("PERJANTAI")
                case time.Saturday:
                        p("LAUANTAI")
                case time.Sunday:
                        p("SUNNUNTAI")
        }

        diff := jA.Sub(now)

        p("Siihen on enää", (diff.Hours()/24), "päivää,")
        p("tunneissa se on karkeasti", (diff.Hours()), "tuntia.")
}
