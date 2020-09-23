package mytarget_sdk

import "time"

type StatisticsPartnersPadsFilter struct {
	DateFrom time.Time //Начальная дата
	DateTo   time.Time //Конечная дата (включительно)
	Id       string    //Список идентификаторов площадок или групп площадок
}

type StatisticsPadsWithSitesFilter struct {
	DateFrom time.Time //Начальная дата
	DateTo   time.Time //Конечная дата (включительно)
	Pads     string    //Список идентификаторов площадок
}

type StatisticsResource struct {
}

func (sr *StatisticsResource) GetPartnersPadsList(filter *StatisticsPartnersPadsFilter) {

}

func (sr *StatisticsResource) GetPadsWithSitesList(filter *StatisticsPadsWithSitesFilter) {

}
