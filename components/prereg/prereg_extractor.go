package prereg

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/itimofeev/hustledb/components/util"
	"gopkg.in/mgo.v2/bson"
	"strings"
	"time"
)

type PreregComp struct {
	ID              bson.ObjectId       `json:"id"bson:"_id,omitempty"`
	FCompetitionId  int64               `bson:"f_competition_id" json:"f_competition_id"`
	Nominations     []*PreregNomination `json:"nominations"`
	FCompetitionUrl string              `json:"f_competition_url"`
	UpdateDate      time.Time           `json:"update_date" bson:"update_date"`
}

type PreregNomination struct {
	Title   string          `json:"title"`
	Records []*PreregRecord `json:"records"`
}

type PreregRecord struct {
	Index   int           `json:"index"`
	Dancer1 *PreregDancer `json:"dancer_1" bson:"dancer_1,omitempty"`
	Dancer2 *PreregDancer `json:"dancer_2" bson:"dancer_2,omitempty"`
}

type PreregDancer struct {
	CodeASH     *string       `json:"code_ash"`
	DancerClass string        `json:"dancer_class"`
	Title       string        `json:"title"`
	Clubs       []*PreregClub `json:"clubs"`
}

type PreregClub struct {
	Title string `json:"title" db:"title"`
}

func (c *PreregComp) findEmptyNomination() *PreregNomination {
	for _, nom := range c.Nominations {
		if len(nom.Records) == 0 {
			return nom
		}
	}
	return nil
}

const mainUrl = "http://www.liveindance.ru/"
const regUrl = "http://www.liveindance.ru/contest/reg.php?id=%d"
const listUrl = "http://www.liveindance.ru/contest/registration/list.php?id=%d"
const forumDir = "/Users/ilyatimofee/prog/axxonsoft/src/github.com/itimofeev/hustledb/tools/prereg/"

func ParseAllPreregLinks() []string {
	data, err := util.GetUrlContent(mainUrl)
	util.CheckErr(err)
	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(data))
	util.CheckErr(err)

	var listLinks []string

	doc.
		Find(".panel-body a").
		FilterFunction(func(i int, s *goquery.Selection) bool {
			link := s.AttrOr("href", "")
			return strings.Contains(link, "http://www.liveindance.ru/contest/registration/list.php?id=")
		}).
		Each(func(i int, s *goquery.Selection) {
			link := s.AttrOr("href", "")
			listLinks = append(listLinks, link)
		})
	return listLinks
}

func ParsePreregId(listLink string) int {
	preregId := strings.Replace(listLink, "http://www.liveindance.ru/contest/registration/list.php?id=", "", 1)

	return util.Atoi(preregId)
}

func GetForumCompetitionId(preregId int) string {
	regUrlFull := fmt.Sprintf(regUrl, preregId)
	data, err := util.GetUrlContent(regUrlFull)
	util.CheckErr(err)

	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(data))
	util.CheckErr(err)

	forumLinkA := doc.
		Find("table table table tr a").
		FilterFunction(func(i int, s *goquery.Selection) bool {
			return strings.Contains(s.AttrOr("href", ""), "http://hustle-sa.ru/forum/index.php?showtopic=")
		}).First()

	return forumLinkA.AttrOr("href", "")
}

func ParsePreregCompetition(preregId int, fCompUrl string) *PreregComp {
	listUrlFull := fmt.Sprintf(listUrl, preregId)
	//data := util.DownloadUrlToFileIfNotExists(listUrlFull, fmt.Sprintf("%s%d.html", forumDir, preregId))
	data, err := util.GetUrlContent(listUrlFull)
	util.CheckErr(err)

	preregComp := PreregComp{FCompetitionUrl: fCompUrl}

	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(data))
	util.CheckErr(err)

	h2AndTables := doc.Find("form>table>tbody>tr").Eq(1).Children().Children().FilterFunction(func(i int, s *goquery.Selection) bool {
		return s.Get(0).Data == "table" || s.Get(0).Data == "h2"
	})

	h2AndTables.Each(func(i int, s *goquery.Selection) {
		if s.Get(0).Data == "h2" {
			preregComp.Nominations = append(preregComp.Nominations, &PreregNomination{
				Title: strings.TrimSpace(s.Text()),
			})
		} else {
			nom := preregComp.findEmptyNomination()
			records := parseRecordsFromTable(s)
			nom.Records = records
		}
	})

	return &preregComp
}

func parseRecordsFromTable(sTable *goquery.Selection) []*PreregRecord {
	var records []*PreregRecord
	isWomanNom := false

	sTable.Find("tr").Each(func(iTr int, sTr *goquery.Selection) {
		if _, ok := sTr.Attr("id"); !ok {
			secondColumnTitle := strings.ToLower(sTr.Children().Eq(1).Text())
			isWomanNom = strings.Contains(secondColumnTitle, "партнёрши") || strings.Contains(secondColumnTitle, "партнерши")
			return
		}

		record := &PreregRecord{Dancer1: &PreregDancer{}, Dancer2: &PreregDancer{}}
		records = append(records, record)
		record.Index = iTr

		tds := sTr.Find("td")
		tds.Each(func(i int, s *goquery.Selection) {
			if isWomanNom {
				switch i {
				case 1:
					record.Dancer2.Title = s.Text()
				case 2:
					record.Dancer2.DancerClass = s.Text()
				case 3:
					{
						code := s.Text()
						record.Dancer2.CodeASH = &code
					}
				case 4:
					record.Dancer2.Clubs = parseClubs(s.Text())
				}
			} else {
				switch i {
				case 1:
					record.Dancer1.Clubs = parseClubs(s.Text())
				case 2:
					{
						code := s.Text()
						record.Dancer1.CodeASH = &code
					}
				case 3:
					record.Dancer1.DancerClass = s.Text()
				case 4:
					record.Dancer1.Title = s.Text()
				case 5:
					record.Dancer2.Title = s.Text()
				case 6:
					record.Dancer2.DancerClass = s.Text()
				case 7:
					{
						code := s.Text()
						record.Dancer2.CodeASH = &code
					}
				case 8:
					record.Dancer2.Clubs = parseClubs(s.Text())
				}
			}
		})

		if len(record.Dancer1.Title) == 0 {
			record.Dancer1 = nil
		}

		if len(record.Dancer2.Title) == 0 {
			record.Dancer2 = nil
		}
	})
	return records
}

func parseClubs(clubNames string) []*PreregClub {
	clubNames = strings.TrimSpace(clubNames)
	if len(clubNames) == 0 || clubNames == "-" {
		return nil
	}

	clubs := make([]*PreregClub, 0, len(clubNames))
	for _, clubName := range strings.Split(clubNames, ",") {
		clubs = append(clubs, &PreregClub{Title: strings.TrimSpace(clubName)})
	}
	return clubs
}
