package cmd

import (
	"context"
	"fmt"
	"github.com/Sutheres/report-chaser/internal/edgar"
	"github.com/Sutheres/report-chaser/internal/edgar/models"
	"github.com/Sutheres/report-chaser/worker"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"strings"
)

func init() {
	rootCmd.AddCommand(workerCmd)
}

var workerCmd = &cobra.Command{
	Use:   "worker",
	Short: "Starts Report Chaser worker",
	Run:   startWorker,
}

func startWorker(cmd *cobra.Command, args []string) {

	e := edgar.NewClient("https://www.sec.gov/Archives/edgar/daily-index")

	w := worker.NewWorker(
		context.Background(),
		worker.WithSEC(e),
	)

	reports, err := e.GetDailyReports()
	if err != nil {
		log.Fatal(err)
	}
	for _, r := range reports {
		content, err := ioutil.ReadFile(r.Name)
		if err != nil {
			log.Fatal(err)
		}

		var forms []models.FormFiling
		// Convert []byte to string and print to screen
		text := string(content)
		lines := strings.Split(text, "\n")
		//fmt.Println(lines)
		for _, line := range lines {
			parsedLine := strings.Split(line, "|")
			//fmt.Println(parsedLine)
			if len(parsedLine) == 5 {
				f := models.FormFiling{
					CIK:         parsedLine[0],
					CompanyName: parsedLine[1],
					Type:        models.FormType(parsedLine[2]),
					DateFiled:   parsedLine[3],
					FileName:    parsedLine[4],
				}
				forms = append(forms, f)
			}
		}

		for _, form := range forms {
			fmt.Println(form.CompanyName)
			fmt.Println("form type:", form.Type)
			fmt.Println(fmt.Sprintf("%s/%s", "https://www.sec.gov/Archives/", form.FileName))
			fmt.Println(" ")
		}
	}

	fmt.Println("No new reports today...")


	w.Start()
}