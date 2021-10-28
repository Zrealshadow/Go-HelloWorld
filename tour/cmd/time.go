package cmd

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/go-programming-tour-book/tour/internal/timer"
	"github.com/spf13/cobra"
)

var calculateTime string
var duration string

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "time format process",
	Long:  "time format process",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var nowTimeCmd = &cobra.Command{
	Use:   "now",
	Short: "Get current time",
	Long:  "Get current time",
	Run: func(cmd *cobra.Command, args []string) {
		nowTime := timer.GetNowTime()
		log.Printf("CurrentTime: %s, %d", nowTime.Format("2006-01-02 14:04:05"), nowTime.Unix())
	},
}

var calculateTimeCmd = &cobra.Command{
	Use:   "calc",
	Short: "calculate time",
	Long:  "calculate time",
	Run: func(cmd *cobra.Command, args []string) {
		var currentTimer time.Time
		var layout = "2006-01-02 15:04:05"

		if calculateTime == "" {
			currentTimer = timer.GetNowTime()
		} else {
			var err error

			space := strings.Count(calculateTime, " ")
			// fmt.Printf("Space %d\n", space)
			if space == 0 {
				layout = "2006-01-02"
			}

			if space == 1 {
				layout = "2006-01-02 15:04:05"
			}

			currentTimer, err = time.Parse(layout, calculateTime)

			if err != nil {
				t, _ := strconv.Atoi(calculateTime)
				currentTimer = time.Unix(int64(t), 0)
				// fmt.Printf("Current Timer %v\n", currentTimer)
			}
		}
		// fmt.Printf("Current Timer %v\n", currentTimer)
		t, err := timer.GetCalculateTime(currentTimer, duration)
		if err != nil {
			log.Fatalf("timer.GetCalculateTime err: %v", err)
		}

		log.Printf("output: %s, %d", t.Format(layout), t.Unix())

	},
}

func init() {
	timeCmd.AddCommand(nowTimeCmd)
	timeCmd.AddCommand(calculateTimeCmd)

	calculateTimeCmd.Flags().StringVarP(&calculateTime, "calculate", "c", "", "Please Input time")
	calculateTimeCmd.Flags().StringVarP(&duration, "duration", "d", "", "Keep time duration")
}
