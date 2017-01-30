package controllers

import (
	"strconv"

	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Result(start int, inputInterest int, interestFrequency int, inputInvestment int, investmentFrequency int, period int) revel.Result {

	revel.WARN.Println("計算開始")

	resultdata := resultData{}
	periodManth := (period * 12) + 1
	invested := float64(start)
	simpleInvested := int(start)
	interested := float64(0)
	investmentSpan := getSpanManthFromNum(investmentFrequency)
	interestSpan := getSpanManthFromNum(interestFrequency)

	resultdata.InterestManth = make([]float64, periodManth, periodManth)
	resultdata.InvestmentResultManth = make([]float64, periodManth, periodManth)
	resultdata.InvestmentNoInterestManth = make([]int, periodManth, periodManth)
	resultdata.InterestSumManth = make([]float64, periodManth, periodManth)

	resultdata.InterestManth[0] = 0
	resultdata.InvestmentNoInterestManth[0] = start
	resultdata.InvestmentResultManth[0] = invested

	revel.WARN.Println("start" + strconv.Itoa(start))
	revel.WARN.Println("select" + strconv.Itoa(interestFrequency))
	revel.WARN.Println("span" + strconv.Itoa(interestSpan))
	revel.WARN.Println("interest" + strconv.Itoa(inputInterest))
	revel.WARN.Println("invest" + strconv.Itoa(inputInvestment))

	for manth := 1; manth < periodManth; manth++ {

		//利息計算
		if interestSpan == 1 || manth%interestSpan == 0 {

			invested, interested = addInterest(invested, interested, inputInterest)
			resultdata.InterestSumManth[manth] = resultdata.InterestSumManth[manth-1] + interested

		} else {

			resultdata.InterestSumManth[manth] = resultdata.InterestSumManth[manth-1]

		}

		//追加投資計算
		if investmentSpan == 1 || manth%investmentSpan == 0 {

			simpleInvested = int(addInvest(float64(simpleInvested), inputInvestment, manth, investmentSpan))
			invested = addInvest(invested, inputInvestment, manth, investmentSpan)

		}

		resultdata.InterestManth[manth] = interested
		resultdata.InvestmentNoInterestManth[manth] = simpleInvested
		resultdata.InvestmentResultManth[manth] = invested

	}

	resultdata.InterestSpan = interestSpan

	return c.RenderJson(resultdata)
}

func addInterest(invested float64, interested float64, interest int) (resultInvest float64, resultInterest float64) {
	resultInterest = float64(invested) * (float64(interest) * 0.0100)
	resultInvest = invested + resultInterest
	return
}

func addInvest(invested float64, invest int, manth int, spanManth int) (result float64) {
	result = invested + float64(invest)
	return
}

func getSpanManthFromNum(selectNum int) (span int) {
	switch selectNum {
	case 1:
		span = 12
	case 2:
		span = 1
	}
	return
}

func clculateAsset() (assets []int) {
	return
}

type resultData struct {
	InterestSumManth          []float64 //利子の合計
	InterestManth             []float64 //毎月の利子
	InvestmentNoInterestManth []int     //利回りなしの毎月の合計
	InvestmentResultManth     []float64 //すべての合計毎月
	InterestSpan              int       //何ヶ月で利子発生か
}
