package main

import (
	"fmt"
)

func main() {
	var cash int
	fmt.Scan(&cash)

	prices := make([]int, 14)
	for i := 0; i < 14; i++ {
		fmt.Scan(&prices[i])
	}

	// BNP 전략 실행
	bnpCash, bnpStocks := cash, 0
	for _, price := range prices {
		if bnpCash >= price {
			buy := bnpCash / price
			bnpStocks += buy
			bnpCash -= buy * price
		}
	}
	//(현금 + (보유 주식 * 1월 14일 주가))
	bnpAssets := bnpCash + bnpStocks*prices[13]

	// TIMING 전략 실행
	timingCash, timingStocks := cash, 0
	upCnt, downCnt := 0, 0

	for i := 1; i < 14; i++ {
		if prices[i] > prices[i-1] {
			upCnt++
			downCnt = 0
		} else if prices[i] < prices[i-1] {
			downCnt++
			upCnt = 0
		} else {
			upCnt, downCnt = 0, 0
		}

		// 3일 연속 상승 → 전량 매도
		if upCnt == 3 && timingStocks > 0 {
			timingCash += timingStocks * prices[i]
			timingStocks = 0
		}

		// 3일 연속 하락 → 전량 매수
		if downCnt == 3 && timingCash >= prices[i] {
			timingStocks = timingCash / prices[i]
			timingCash -= timingStocks * prices[i]
		}
	}
	//(현금 + (보유 주식 * 1월 14일 주가)).
	timingAssets := timingCash + timingStocks*prices[13]

	// 결과 출력
	if bnpAssets > timingAssets {
		fmt.Println("BNP")
	} else if timingAssets > bnpAssets {
		fmt.Println("TIMING")
	} else {
		fmt.Println("SAMESAME")
	}
}
