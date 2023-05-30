package controller

import (
	. "ffAPI/models"

	_ "github.com/go-sql-driver/mysql"
)

func GetSearchResult(searchParam SearchParam) []SearchResult {
	results := ExecuteSQL("select d.DATA_NO , nvl(ac.CITY_NAME,''), DATE_FORMAT(d.DATE_WORK, '%d.%m.%Y'), d.TIME_WORK , d.FLASCHEN_FUELLEN , d.FLASCHEN_TUEV , d.MASKEN_REINIGEN , d.MASKEN_PRUEFEN , d.LA_REINIGEN , d.LA_PRUEFEN , d.GERAETE_PRUEFEN , d.GERAETE_REINIGEN, d.BEMERKUNG from atemschutzpflegestelle_data d left join atemschutzpflegestelle_cities ac on d.CITY_NO=ac.CITY_NO where PERS_NO = ? order by d.DATE_WORK desc", searchParam.PersNo)
	searchResults := []SearchResult{}
	for results.Next() {
		var searchResult SearchResult
		results.Scan(&searchResult.Key, &searchResult.City, &searchResult.DateWork, &searchResult.TimeWork, &searchResult.FlaschenFuellen, &searchResult.FlaschenTuev, &searchResult.MaskenReinigen, &searchResult.MaskenPruefen, &searchResult.LaReinigen, &searchResult.LaPruefen, &searchResult.GeraetePruefen, &searchResult.GeraeteReinigen, &searchResult.Bemerkung)
		searchResults = append(searchResults, searchResult)
	}
	return searchResults
}
