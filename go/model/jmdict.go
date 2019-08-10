package model

type JMDict struct {
	Entries []Entry `xml:"entry"`
}

type Entry struct {
	KanjiElements   []KanjiElement   `xml:"k_ele"`
	ReadingElements []ReadingElement `xml:"r_ele"`
	SenseElements   []Sense          `xml:"sense"`
}

type KanjiElement struct {
	Keb                  string   `xml:"keb"`
	KanjiElementInfo     []string `xml:"ke_inf"`
	KanjiElementPriority []string `xml:"ke_pri"`
}

type ReadingElement struct {
	KReb                   string   `xml:"reb"`
	NoKanji                string   `xml:"re_nokanji"`
	RestrictedTo           []string `xml:"re_restr"`
	ReadingElementInfo     []string `xml:"re_inf"`
	ReadingElementPriority []string `xml:"re_pri"`
}
type Sense struct {
	RestrictedToKanji   []string         `xml:"stagk"`
	RestrictedToReading []string         `xml:"stagr"`
	PartOfSpeech        []string         `xml:"pos"`
	CrossReference      []string         `xml:"xref"`
	Antonym             []string         `xml:"ant"`
	Field               []string         `xml:"field"`
	MiscInfo            []string         `xml:"misc"`
	SenseInfo           []string         `xml:"s_inf"`
	LanguageSource      []LanguageSource `xml:"lsource"`
	Dialect             []string         `xml:"dial"`
	Glossary            []string         `xml:"gloss"`
}

type LanguageSource struct {
	Language string `xml:"lang,attr"`
	Word     string `xml:",chardata"`
}
