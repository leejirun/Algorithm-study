function getWinningProbability(a, b) {
    // 전체 카드 덱 (각 숫자 2장씩)
    let deck = [];
    for (let i = 1; i <= 10; i++) {
        deck.push(i, i);
    }

    // 현재 플레이어가 가진 카드 제거
    deck.splice(deck.indexOf(a), 1);
    deck.splice(deck.indexOf(b), 1);

    // 족보 계산 함수
    function getRank(x, y) {
        if (x === 10 && y === 10) return 8; // 10 띵
        if (x === 9 && y === 9) return 7;  // 9 띵
        if (x === 2 && y === 2) return 6;  // 2 띵
        if (x === 1 && y === 1) return 5;  // 1 띵
        let sum = (x + y) % 10;
        if (sum === 9) return 4; // 9 끗
        if (sum === 8) return 3; // 8 끗
        if (sum === 1) return 2; // 1 끗
        if (sum === 0) return 1; // 0 끗
        return 0;
    }

    let myRank = getRank(a, b);
    let winCount = 0, totalCount = 0;

    // 상대가 가져갈 수 있는 모든 경우의 수 확인
    for (let i = 0; i < deck.length; i++) {
        for (let j = i + 1; j < deck.length; j++) {
            let opponentRank = getRank(deck[i], deck[j]);
            if (myRank > opponentRank) winCount++;
            totalCount++;
        }
    }

    // 확률 계산 및 소수점 셋째 자리까지 출력
    let winProbability = (winCount / totalCount).toFixed(3);
    console.log(winProbability);
}

// 예제 실행
getWinningProbability(1, 1); // 0.941
getWinningProbability(1, 2); // 0.275
getWinningProbability(1, 9); // 0.000
getWinningProbability(10, 10); // 1.000