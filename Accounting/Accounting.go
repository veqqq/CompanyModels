package Accounting

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

// Custom Types

// Custom loat64 to handle nulls etc. when unmarshaling

type OwnFloat64 struct {
	Value float64
	Valid bool
}

func (of *OwnFloat64) UnmarshalJSON(data []byte) error {
	var rawValue interface{}
	if err := json.Unmarshal(data, &rawValue); err != nil {
		return err
	}

	switch v := rawValue.(type) {
	case float64:
		of.Value = v
		of.Valid = true
	case string:
		if v == "None" {
			of.Valid = false
		} else {
			value, err := strconv.ParseFloat(v, 64)
			if err != nil {
				return err
			}
			of.Value = value
			of.Valid = true
		}
	default:
		return fmt.Errorf("unexpected value type for OwnFloat64: %T", v)
	}

	return nil
}

func (of OwnFloat64) MarshalJSON() ([]byte, error) {
	if !of.Valid {
		return json.Marshal(0)
	}
	return json.Marshal(fmt.Sprint(int64(of.Value)))
}

///////////////////////
// Stock Ticker Data:

// TIME_SERIES_DAILY
type DailyOHLCVs struct {
	MetaData   DailyOHLCVMetaData    `json:"Meta Data"`
	TimeSeries map[string]DailyOHLCV `json:"Time Series (Daily)"`
}

// reformat stuff for this?
type DailyOHLCVMetaData struct {
	Information   string `json:"1. Information"`
	Symbol        string `json:"2. Symbol"`
	LastRefreshed string `json:"3. Last Refreshed"`
	OutputSize    string `json:"4. Output Size"`
	TimeZone      string `json:"5. Time Zone"`
}

type DailyOHLCV struct {
	Open   float64 `json:"1. open,string"`
	High   float64 `json:"2. high,string"`
	Low    float64 `json:"3. low,string"`
	Close  float64 `json:"4. close,string"`
	Volume int64   `json:"5. volume,string"`
}

// TIME_SERIES_INTRADAY
type IntradayOHLCVs struct {
	MetaData       IntradayOHLCVMetaData    `json:"Meta Data"`
	TimeSeries1min map[string]IntradayOHLCV `json:"Time Series (1min)"`
}

type IntradayOHLCVMetaData struct {
	Information   string `json:"1. Information"`
	Symbol        string `json:"2. Symbol"`
	LastRefreshed string `json:"3. Last Refreshed"`
	Interval      string `json:"4. Interval"`
	OutputSize    string `json:"5. Output Size"`
	TimeZone      string `json:"6. Time Zone"`
}

type IntradayOHLCV struct {
	Open   float64 `json:"1. open,string"`
	High   float64 `json:"2. high,string"`
	Low    float64 `json:"3. low,string"`
	Close  float64 `json:"4. close,string"`
	Volume int64   `json:"5. volume,string"`
}

//////////////////
// Stock Fundementals:
// The 4 financial statements + "overview"

// overview
type StockOverview struct {
	Symbol                     string     `json:"Symbol"`
	AssetType                  string     `json:"AssetType"`
	Name                       string     `json:"Name"`
	Description                string     `json:"Description"`
	CIK                        string     `json:"CIK"`
	Exchange                   string     `json:"Exchange"`
	Currency                   string     `json:"Currency"`
	Country                    string     `json:"Country"`
	Sector                     string     `json:"Sector"`
	Industry                   string     `json:"Industry"`
	Address                    string     `json:"Address"`
	FiscalYearEnd              string     `json:"FiscalYearEnd"`
	LatestQuarter              string     `json:"LatestQuarter"`
	MarketCapitalization       OwnFloat64 `json:"MarketCapitalization"`
	EBITDA                     OwnFloat64 `json:"EBITDA"`
	PERatio                    OwnFloat64 `json:"PERatio"`
	PEGRatio                   OwnFloat64 `json:"PEGRatio"`
	BookValue                  OwnFloat64 `json:"BookValue"`
	DividendPerShare           OwnFloat64 `json:"DividendPerShare"`
	DividendYield              OwnFloat64 `json:"DividendYield"`
	EPS                        OwnFloat64 `json:"EPS"`
	RevenuePerShareTTM         OwnFloat64 `json:"RevenuePerShareTTM"`
	ProfitMargin               OwnFloat64 `json:"ProfitMargin"`
	OperatingMarginTTM         OwnFloat64 `json:"OperatingMarginTTM"`
	ReturnOnAssetsTTM          OwnFloat64 `json:"ReturnOnAssetsTTM"`
	ReturnOnEquityTTM          OwnFloat64 `json:"ReturnOnEquityTTM"`
	RevenueTTM                 OwnFloat64 `json:"RevenueTTM"`
	GrossProfitTTM             OwnFloat64 `json:"GrossProfitTTM"`
	DilutedEPSTTM              OwnFloat64 `json:"DilutedEPSTTM"`
	QuarterlyEarningsGrowthYOY OwnFloat64 `json:"QuarterlyEarningsGrowthYOY"`
	QuarterlyRevenueGrowthYOY  OwnFloat64 `json:"QuarterlyRevenueGrowthYOY"`
	AnalystTargetPrice         OwnFloat64 `json:"AnalystTargetPrice"`
	TrailingPE                 OwnFloat64 `json:"TrailingPE"`
	ForwardPE                  OwnFloat64 `json:"ForwardPE"`
	PriceToSalesRatioTTM       OwnFloat64 `json:"PriceToSalesRatioTTM"`
	PriceToBookRatio           OwnFloat64 `json:"PriceToBookRatio"`
	EVToRevenue                OwnFloat64 `json:"EVToRevenue"`
	EVToEBITDA                 OwnFloat64 `json:"EVToEBITDA"`
	Beta                       OwnFloat64 `json:"Beta"`
	WeekHigh                   OwnFloat64 `json:"52WeekHigh"`
	WeekLow                    OwnFloat64 `json:"52WeekLow"`
	DayMovingAverage50         OwnFloat64 `json:"50DayMovingAverage"`
	DayMovingAverage200        OwnFloat64 `json:"200DayMovingAverage"`
	SharesOutstanding          OwnFloat64 `json:"SharesOutstanding"`
	DividendDate               string     `json:"DividendDate"`
	ExDividendDate             string     `json:"ExDividendDate"`
}

// income statement
// json: invalid use of ,string struct tag, trying to unmarshal "None" into float64
// unsure why, so unmarshaling into map any

type IncomeStatements struct {
	Symbol           string            `json:"symbol"`
	AnnualReports    []IncomeStatement `json:"annualReports"`
	QuarterlyReports []IncomeStatement `json:"quarterlyReports"`
}

type IncomeStatement struct {
	FiscalDateEnding                  string     `json:"fiscalDateEnding"`
	ReportedCurrency                  string     `json:"reportedCurrency"`
	GrossProfit                       OwnFloat64 `json:"grossProfit"`
	TotalRevenue                      OwnFloat64 `json:"totalRevenue"`
	CostOfRevenue                     OwnFloat64 `json:"costOfRevenue"`
	CostofGoodsAndServicesSold        OwnFloat64 `json:"costofGoodsAndServicesSold"`
	OperatingIncome                   OwnFloat64 `json:"operatingIncome"`
	SellingGeneralAndAdministrative   OwnFloat64 `json:"sellingGeneralAndAdministrative"`
	ResearchAndDevelopment            OwnFloat64 `json:"researchAndDevelopment"`
	OperatingExpenses                 OwnFloat64 `json:"operatingExpenses"`
	InvestmentIncomeNet               OwnFloat64 `json:"investmentIncomeNet"`
	NetInterestIncome                 OwnFloat64 `json:"netInterestIncome"`
	InterestIncome                    OwnFloat64 `json:"interestIncome"`
	InterestExpense                   OwnFloat64 `json:"interestExpense"`
	NonInterestIncome                 OwnFloat64 `json:"nonInterestIncome"`
	OtherNonOperatingIncome           OwnFloat64 `json:"otherNonOperatingIncome"`
	Depreciation                      OwnFloat64 `json:"depreciation"`
	DepreciationAndAmortization       OwnFloat64 `json:"depreciationAndAmortization"`
	IncomeBeforeTax                   OwnFloat64 `json:"incomeBeforeTax"`
	IncomeTaxExpense                  OwnFloat64 `json:"incomeTaxExpense"`
	InterestAndDebtExpense            OwnFloat64 `json:"interestAndDebtExpense"`
	NetIncomeFromContinuingOperations OwnFloat64 `json:"netIncomeFromContinuingOperations"`
	ComprehensiveIncomeNetOfTax       OwnFloat64 `json:"comprehensiveIncomeNetOfTax"`
	EBIT                              OwnFloat64 `json:"ebit"`
	EBITDA                            OwnFloat64 `json:"ebitda"`
	NetIncome                         OwnFloat64 `json:"netIncome"`
}

// https://github.com/veqqq/StockDataSDK/issues/6
// implement UnmarshalJSON to avoid this error:
// json: invalid use of ,string struct tag, trying to unmarshal "None" into float64
// should add every "None" possible field to aux and give it a similar if statement below
//

// balance sheet

type BalanceSheets struct {
	Symbol           string         `json:"symbol"`
	AnnualReports    []BalanceSheet `json:"annualReports"`
	QuarterlyReports []BalanceSheet `json:"quarterlyReports"`
}

type BalanceSheet struct {
	FiscalDateEnding                       string     `json:"fiscalDateEnding"`
	ReportedCurrency                       string     `json:"reportedCurrency"`
	TotalAssets                            OwnFloat64 `json:"totalAssets"`
	TotalCurrentAssets                     OwnFloat64 `json:"totalCurrentAssets"`
	CashAndCashEquivalentsAtCarryingValue  OwnFloat64 `json:"cashAndCashEquivalentsAtCarryingValue"`
	CashAndShortTermInvestments            OwnFloat64 `json:"cashAndShortTermInvestments"`
	Inventory                              OwnFloat64 `json:"inventory"`
	CurrentNetReceivables                  OwnFloat64 `json:"currentNetReceivables"`
	TotalNonCurrentAssets                  OwnFloat64 `json:"totalNonCurrentAssets"`
	PropertyPlantEquipment                 OwnFloat64 `json:"propertyPlantEquipment"`
	AccumulatedDepreciationAmortizationPPE OwnFloat64 `json:"accumulatedDepreciationAmortizationPPE"`
	IntangibleAssets                       OwnFloat64 `json:"intangibleAssets"`
	IntangibleAssetsExcludingGoodwill      OwnFloat64 `json:"intangibleAssetsExcludingGoodwill"`
	Goodwill                               OwnFloat64 `json:"goodwill"`
	Investments                            OwnFloat64 `json:"investments"`
	LongTermInvestments                    OwnFloat64 `json:"longTermInvestments"`
	ShortTermInvestments                   OwnFloat64 `json:"shortTermInvestments"`
	OtherCurrentAssets                     OwnFloat64 `json:"otherCurrentAssets"`
	OtherNonCurrentAssets                  OwnFloat64 `json:"otherNonCurrentAssets"`
	TotalLiabilities                       OwnFloat64 `json:"totalLiabilities"`
	TotalCurrentLiabilities                OwnFloat64 `json:"totalCurrentLiabilities"`
	CurrentAccountsPayable                 OwnFloat64 `json:"currentAccountsPayable"`
	DeferredRevenue                        OwnFloat64 `json:"deferredRevenue"`
	CurrentDebt                            OwnFloat64 `json:"currentDebt"`
	ShortTermDebt                          OwnFloat64 `json:"shortTermDebt"`
	TotalNonCurrentLiabilities             OwnFloat64 `json:"totalNonCurrentLiabilities"`
	CapitalLeaseObligations                OwnFloat64 `json:"capitalLeaseObligations"`
	LongTermDebt                           OwnFloat64 `json:"longTermDebt"`
	CurrentLongTermDebt                    OwnFloat64 `json:"currentLongTermDebt"`
	LongTermDebtNoncurrent                 OwnFloat64 `json:"longTermDebtNoncurrent"`
	ShortLongTermDebtTotal                 OwnFloat64 `json:"shortLongTermDebtTotal"`
	OtherCurrentLiabilities                OwnFloat64 `json:"otherCurrentLiabilities"`
	OtherNonCurrentLiabilities             OwnFloat64 `json:"otherNonCurrentLiabilities"`
	TotalShareholderEquity                 OwnFloat64 `json:"totalShareholderEquity"`
	TreasuryStock                          OwnFloat64 `json:"treasuryStock"`
	RetainedEarnings                       OwnFloat64 `json:"retainedEarnings"`
	CommonStock                            OwnFloat64 `json:"commonStock"`
	CommonStockSharesOutstanding           OwnFloat64 `json:"commonStockSharesOutstanding"`
}

// Cash flow

type CashFlowStatements struct {
	Symbol           string              `json:"symbol"`
	AnnualReports    []CashFlowStatement `json:"annualReports"`
	QuarterlyReports []CashFlowStatement `json:"quarterlyReports"`
}

type CashFlowStatement struct {
	FiscalDateEnding                                          string     `json:"fiscalDateEnding"`
	ReportedCurrency                                          string     `json:"reportedCurrency"`
	OperatingCashflow                                         OwnFloat64 `json:"operatingCashflow"`
	PaymentsForOperatingActivities                            OwnFloat64 `json:"paymentsForOperatingActivities"`
	ProceedsFromOperatingActivities                           OwnFloat64 `json:"proceedsFromOperatingActivities"`
	ChangeInOperatingLiabilities                              OwnFloat64 `json:"changeInOperatingLiabilities"`
	ChangeInOperatingAssets                                   OwnFloat64 `json:"changeInOperatingAssets"`
	DepreciationDepletionAndAmortization                      OwnFloat64 `json:"depreciationDepletionAndAmortization"`
	CapitalExpenditures                                       OwnFloat64 `json:"capitalExpenditures"`
	ChangeInReceivables                                       OwnFloat64 `json:"changeInReceivables"`
	ChangeInInventory                                         OwnFloat64 `json:"changeInInventory"`
	ProfitLoss                                                OwnFloat64 `json:"profitLoss"`
	CashflowFromInvestment                                    OwnFloat64 `json:"cashflowFromInvestment"`
	CashflowFromFinancing                                     OwnFloat64 `json:"cashflowFromFinancing"`
	ProceedsFromRepaymentsOfShortTermDebt                     OwnFloat64 `json:"proceedsFromRepaymentsOfShortTermDebt"`
	PaymentsForRepurchaseOfCommonStock                        OwnFloat64 `json:"paymentsForRepurchaseOfCommonStock"`
	PaymentsForRepurchaseOfEquity                             OwnFloat64 `json:"paymentsForRepurchaseOfEquity"`
	PaymentsForRepurchaseOfPreferredStock                     OwnFloat64 `json:"paymentsForRepurchaseOfPreferredStock"`
	DividendPayout                                            OwnFloat64 `json:"dividendPayout"`
	DividendPayoutCommonStock                                 OwnFloat64 `json:"dividendPayoutCommonStock"`
	DividendPayoutPreferredStock                              OwnFloat64 `json:"dividendPayoutPreferredStock"`
	ProceedsFromIssuanceOfCommonStock                         OwnFloat64 `json:"proceedsFromIssuanceOfCommonStock"`
	ProceedsFromIssuanceOfLongTermDebtAndCapitalSecuritiesNet OwnFloat64 `json:"proceedsFromIssuanceOfLongTermDebtAndCapitalSecuritiesNet"`
	ProceedsFromIssuanceOfPreferredStock                      OwnFloat64 `json:"proceedsFromIssuanceOfPreferredStock"`
	ProceedsFromRepurchaseOfEquity                            OwnFloat64 `json:"proceedsFromRepurchaseOfEquity"`
	ProceedsFromSaleOfTreasuryStock                           OwnFloat64 `json:"proceedsFromSaleOfTreasuryStock"`
	ChangeInCashAndCashEquivalents                            OwnFloat64 `json:"changeInCashAndCashEquivalents"`
	ChangeInExchangeRate                                      OwnFloat64 `json:"changeInExchangeRate"`
	NetIncome                                                 OwnFloat64 `json:"netIncome"`
}

// Earnings
type EarningsData struct {
	Symbol            string              `json:"symbol"`
	AnnualEarnings    []AnnualEarnings    `json:"annualEarnings"`
	QuarterlyEarnings []QuarterlyEarnings `json:"quarterlyEarnings"`
}

type AnnualEarnings struct { // lol pointless
	FiscalDateEnding string     `json:"fiscalDateEnding"`
	ReportedEPS      OwnFloat64 `json:"reportedEPS"`
}

type QuarterlyEarnings struct {
	FiscalDateEnding   string     `json:"fiscalDateEnding"`
	ReportedDate       string     `json:"reportedDate"`
	ReportedEPS        OwnFloat64 `json:"reportedEPS"`
	EstimatedEPS       OwnFloat64 `json:"estimatedEPS"`
	Surprise           OwnFloat64 `json:"surprise"`
	SurprisePercentage OwnFloat64 `json:"surprisePercentage"`
}

///////

///////
// Commodities and Economic Indicators - use same structure

// WTI, BRENT, nat gas, COPPER, ALUMINUM, WHEAT, CORN, COTTON, SUGAR, COFFEE

// REAL_GDP in billions of dollars - same structure as commodities...
// real gdp per cap - in "chained 2012 dollars"
// fed funds rate - in percent - daily? monthly?
// cpi - "index 1982-1984=100" - monthly
// inflation - in percent, only annual
// retail sales - in millions, only monthly
// durable goods orders - in millions, only monthly
// unemployment - in percent, only monthly
// nonfarm payroll - in thousands of people, only monthly
// treasury yield - in percent, monthly or daily?
//          // maturities: 3month, 2year, 5year, 7year, 10year

type CommodityPrices struct { // rename
	Name     string           `json:"name"`     // e.g. global price of copper or henry hub...
	Interval string           `json:"interval"` // should be daily always
	Unit     string           `json:"unit"`
	Data     []CommodityPrice `json:"data"`
}

type CommodityPrice struct { // rename
	Date  string  `json:"date,omitempty"`
	Value float64 `json:"value,string,omitempty"`
}

// CommodityPrice satisfies Unmarshaler
func (c *CommodityPrice) UnmarshalJSON(data []byte) error {
	var aux struct {
		Date  string `json:"date"`
		Value string `json:"value"`
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	// If val or date are useless, ignore struct
	// some dates have "." val, which normal unmarshaler doesn't handle, hence writing this
	if aux.Value == "null" || aux.Value == "" || aux.Value == "." || aux.Value == "0" || aux.Value == "0.0" || aux.Date == "" || aux.Date == "null" || aux.Date == "." {
		return nil
	}
	value, err := strconv.ParseFloat(aux.Value, 64)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}
	c.Date = aux.Date
	c.Value = value
	return nil
}

/////////
//// Other

// exchange rates, daily

type ForexPrices struct {
	MetaData     ForexMetaData         `json:"Meta Data"`
	TimeSeriesFX map[string]ForexPrice `json:"Time Series FX (Daily)"` // monthly is the same, just uses (Monthly here)
}

type ForexMetaData struct {
	Information   string `json:"1. Information"` // "1. Information": "Forex Daily Prices (open, high, low, close)",
	FromSymbol    string `json:"2. From Symbol"` // "2. From Symbol": "EUR",
	ToSymbol      string `json:"3. To Symbol"`   // "3. To Symbol": "USD",
	OutputSize    string `json:"4. Output Size"`
	LastRefreshed string `json:"5. Last Refreshed"`
	TimeZone      string `json:"6. Time Zone"`
}

type ForexPrice struct {
	Open  float64 `json:"1. open,string"` // "1. open": "1.10020",
	High  float64 `json:"2. high,string"`
	Low   float64 `json:"3. low,string"`
	Close float64 `json:"4. close,string"`
}

// 20 TOP_GAINERS_LOSERS  and most actively traded

type TGLATs struct {
	Metadata     string  `json:"metadata"`
	LastUpdated  string  `json:"last_updated"`
	TopGainers   []TGLAT `json:"top_gainers"`
	TopLosers    []TGLAT `json:"top_losers"`
	MostActively []TGLAT `json:"most_actively_traded"`
}

type TGLAT struct {
	Ticker           string  `json:"ticker"`
	Price            float64 `json:"price,string"`
	ChangeAmount     float64 `json:"change_amount,string"`
	ChangePercentage string  `json:"change_percentage"` // has %, later implement json unmarshaler for this
	Volume           int64   `json:"volume,string"`
}

// News sentiment - complicated beast

type SentimentData struct {
	Items                    string     `json:"items"`
	SentimentScoreDefinition string     `json:"sentiment_score_definition"`
	RelevanceScoreDefinition string     `json:"relevance_score_definition"`
	Feed                     []FeedData `json:"feed"`
}

type FeedData struct {
	Title                 string                `json:"title"`
	URL                   string                `json:"url"`
	TimePublished         string                `json:"time_published"` // "20230805T122000",
	Authors               []string              `json:"authors"`
	Summary               string                `json:"summary"`
	BannerImage           string                `json:"banner_image"`           // lol, for a spam farm?
	Source                string                `json:"source"`                 // "CNBC",
	CategoryWithinSource  string                `json:"category_within_source"` // "Top News",
	SourceDomain          string                `json:"source_domain"`          // "www.cnbc.com",
	Topics                []TopicData           `json:"topics"`
	OverallSentimentScore float64               `json:"overall_sentiment_score"` //	0.072673,
	OverallSentimentLabel string                `json:"overall_sentiment_label"` //	"Neutral",
	TickerSentiment       []TickerSentimentData `json:"ticker_sentiment"`
}

type TopicData struct {
	Topic          string  `json:"topic"`                  // "Technology",
	RelevanceScore float64 `json:"relevance_score,string"` // "1.0"
}

type TickerSentimentData struct {
	Ticker               string  `json:"ticker"`
	RelevanceScore       float64 `json:"relevance_score,string"`        // "0.699089",
	TickerSentimentScore float64 `json:"ticker_sentiment_score,string"` // "0.116531",
	TickerSentimentLabel string  `json:"ticker_sentiment_label"`        //  "Neutral"
}

// var sentimentData SentimentData
// if err == json.Unmarshal([]byte(jsonData), &sentimentData); err != nil {}
// 	fmt.Println("Error unmarshaling JSON:", err)
// 	return
// }
