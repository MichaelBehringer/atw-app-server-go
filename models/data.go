package models

type SearchParam struct {
	PersNo int `json:"persNo"`
}

type SearchResult struct {
	Key             int    `json:"key"`
	City            string `json:"city"`
	DateWork        string `json:"dateWork"`
	TimeWork        int    `json:"timeWork"`
	FlaschenFuellen int    `json:"flaschenFuellen"`
	FlaschenTuev    int    `json:"flaschenTUEV"`
	MaskenReinigen  int    `json:"maskenReinigen"`
	MaskenPruefen   int    `json:"maskenPruefen"`
	LaReinigen      int    `json:"laReinigen"`
	LaPruefen       int    `json:"laPruefen"`
	GeraetePruefen  int    `json:"gereatPruefen"`
	GeraeteReinigen int    `json:"gereatReinigen"`
	Bemerkung       string `json:"bemerkung"`
}
