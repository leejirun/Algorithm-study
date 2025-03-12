
## 💡문제 분석 요약
- 초기 현금 (1000 이하의 양의 정수)
- 1월 1일부터 1월 14일까지의 주식 가격(공백으로 구분된 14개의 양의 정수, 1000 이하)
- BNP (Buy and Pray, 준현): 살 수 있으면 최대한 매수하고, 절대 팔지 않음
- TIMING (33 매매법, 성민): 
3일 연속 상승 시 다음 날 전량 매도
3일 연속 하락 시 다음 날 전량 매수
모든 거래는 전량 매수 또는 전량 매도로 진행
빚을 내거나 부분 매수/매도 불가능
-1월 14일 기준 준현의 자산이 크면 "BNP"/성민의 자산이 크면 "TIMING"/같으면 "SAMESAME"

## 💡알고리즘 설계
1.BNP 
매일 주식을 살 수 있는지 확인 → 최대한 매수
1월 14일 자산 계산

2.TIMING 전략 시뮬레이션
연속 상승/하락 여부 확인하여 매매 수행
1월 14일 자산 계산

3.비교

## 💡코드
package main

import (
	"fmt"
)

func main() {
	// 입력 받기
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


## 💡개념
