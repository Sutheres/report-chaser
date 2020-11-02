package cmd

import (
	"encoding/xml"
	"fmt"
	"github.com/Sutheres/report-chaser/internal/edgar"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(workerCmd)
}

var workerCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts Report Chaser server",
	Run:   startServer,
}

func startServer(cmd *cobra.Command, args []string) {

	_ = edgar.NewClient("https://www.sec.gov/Archives/edgar/daily-index")


	//reports, err := e.GetDailyReports()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//for _, r := range reports {
	//	content, err := ioutil.ReadFile(r.Name)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	var forms []models.FormFiling
	//	// Convert []byte to string and print to screen
	//	text := string(content)
	//	lines := strings.Split(text, "\n")
	//	//fmt.Println(lines)
	//	for _, line := range lines {
	//		parsedLine := strings.Split(line, "|")
	//		//fmt.Println(parsedLine)
	//		if len(parsedLine) == 5 {
	//			f := models.FormFiling{
	//				CIK:         parsedLine[0],
	//				CompanyName: parsedLine[1],
	//				Type:        models.FormType(parsedLine[2]),
	//				DateFiled:   parsedLine[3],
	//				FileName:    parsedLine[4],
	//			}
	//			forms = append(forms, f)
	//		}
	//	}
	//
	//	for _, form := range forms {
	//		fmt.Println(form.CompanyName)
	//		fmt.Println("form type:", form.Type)
	//		fmt.Println(fmt.Sprintf("%s/%s", "https://www.sec.gov/Archives/", form.FileName))
	//		fmt.Println(" ")
	//	}
	//}
	//
	//fmt.Println("No new reports today...")

	blob := []byte(`<xbrli:xbrl>
						<link:schemaRef xlink:href="trne-20191231.xsd" xlink:type="simple"/>
				
						<xbrli:context id="From2019-01-01to2019-12-31">
							<xbrli:entity>
								<xbrli:identifier scheme="http://www.sec.gov/CIK">0001754820</xbrli:identifier>
							</xbrli:entity>
							<xbrli:period>
								<xbrli:startDate>2019-01-01</xbrli:startDate>
								<xbrli:endDate>2019-12-31</xbrli:endDate>
						</xbrli:period>
						</xbrli:context>


						<xbrli:context id="From2019-01-01to2019-12-31_us-gaap_CommonClassAMember">
							<xbrli:entity>
							<xbrli:identifier scheme="http://www.sec.gov/CIK">0001754820</xbrli:identifier>
								<xbrli:segment>
									<xbrldi:explicitMember dimension="us-gaap:StatementClassOfStockAxis">us-gaap:CommonClassAMember</xbrldi:explicitMember>
								</xbrli:segment>
							</xbrli:entity>
							<xbrli:period>
								<xbrli:startDate>2019-01-01</xbrli:startDate>
								<xbrli:endDate>2019-12-31</xbrli:endDate>
							</xbrli:period>
						</xbrli:context>

						
						<xbrli:context id="AsOf2018-12-31_us-gaap_CommonClassAMember">
							<xbrli:entity>
								<xbrli:identifier scheme="http://www.sec.gov/CIK">0001754820</xbrli:identifier>
								<xbrli:segment>
									<xbrldi:explicitMember dimension="us-gaap:StatementClassOfStockAxis">us-gaap:CommonClassAMember</xbrldi:explicitMember>
								</xbrli:segment>
							</xbrli:entity>
								<xbrli:period>
									<xbrli:instant>2018-12-31</xbrli:instant>
								</xbrli:period>
						</xbrli:context>

					</xbrli:xbrl>`)

	fmt.Println(string(blob))

	var p XBRLParser
	if err := xml.Unmarshal(blob, &p); err != nil {
		panic(err)
	}

	//reader := bytes.NewReader(blob)
	//decoder := xml.NewDecoder(reader)
	//decoder.CharsetReader = charset.NewReaderLabel
	//err = decoder.Decode(&p)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(p)
	for _, ctx := range p.Ctx {
		fmt.Println(ctx.ID)
		fmt.Println(ctx.Entity.Identifier.Scheme)
	}


}

type XBRLParser struct {
	XMLName xml.Name        `xml:"xbrli\:xbrl"`
	Ctx     []ReportContext `xml:"context"`
	Origin  []string        `xml:"origin"`
}

type ReportContext struct {
	XMLName xml.Name `xml:"xbrli\:context,omitempty"`
	ID      string   `xml:"id,attr"`
	Entity  Entity   `xml:"entity,omitempty"`
	Period  Period   `xml:"period,omitempty"`
}

type Entity struct {
	XMLName    xml.Name   `xml:"xbrli\:entity,omitempty"`
	Identifier Identifier `xml:"identifier,omitempty"`
	Segment    Segment    `xml:"segment,omitempty"`
}

type Period struct {
	XMLName   xml.Name `xml:"xbrli\:period,omitempty"`
	StartDate string   `xml:"startDate,omitempty"`
	EndDate   string   `xml:"endDate,omitempty"`
	Instant string `xml:"instant,omitempty"`
}

type Identifier struct {
	XMLName xml.Name `xml:"xbrli\:identifier,omitempty"`
	Scheme  string   `xml:"scheme,attr,omitempty"`
	Value   string   `xml:",chardata"`
}

type Segment struct {
	XMLName        xml.Name       `xml:"xbrli\:segment"`
	ExplicitMember ExplicitMember `xml:"explicitMember"`
}

type ExplicitMember struct {
	XMLName   xml.Name `xml:"xbrldi\:explicitMember"`
	Dimension string   `xml:"dimension,attr"`
	Value     string   `xml:",chardata"`
}

type Instant struct {
	XMLName xml.Name `xml:"xbrli\:instant"`
	Value string `xml:",chardata"`
}
