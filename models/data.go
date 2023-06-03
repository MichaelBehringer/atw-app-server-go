package models

type SearchParam struct {
	PersNo int `json:"persNo"`
}

type SearchResult struct {
	DataNo          int    `json:"key"`
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

type EntryObj struct {
	DataNo            int    `json:"dataNo"`
	City              int    `json:"city"`
	User              int    `json:"user"`
	DateWork          string `json:"dateWork"`
	TimeWork          int    `json:"arbeitszeit"`
	FlaschenFuellen   int    `json:"flaschenFuellen"`
	FlaschenTuev      int    `json:"flaschenTUEV"`
	MaskenReinigen    int    `json:"maskenReinigen"`
	MaskenPruefen     int    `json:"maskenPruefen"`
	LaReinigen        int    `json:"laReinigen"`
	LaPruefen         int    `json:"laPruefen"`
	GeraetePruefen    int    `json:"geraetePruefen"`
	GeraeteReinigen   int    `json:"geraeteReinigen"`
	Bemerkung         string `json:"bemerkung"`
	FlaschenFuellenNr string `json:"flaschenFuellenNr"`
	FlaschenTuevNr    string `json:"flaschenTUEVNr"`
	MaskenReinigenNr  string `json:"maskenReinigenNr"`
	MaskenPruefenNr   string `json:"maskenPruefenNr"`
	LaReinigenNr      string `json:"laReinigenNr"`
	LaPruefenNr       string `json:"laPruefenNr"`
	GeraetePruefenNr  string `json:"geraetePruefenNr"`
	GeraeteReinigenNr string `json:"geraeteReinigenNr"`
}
