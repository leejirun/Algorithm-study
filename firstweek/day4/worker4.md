## 💡문제 분석 요약

1. 카드 게임으로 2명이 각 10장씩 카드를 뽑는다.
2. 강한 족보
10땡: 두 패가 모두 10
9땡: 두 패가 모두 9
...
2땡: 두 패가 모두 2
1땡: 두 패가 모두 1
9끗: 땡이 아니고, 두 패를 더했을 때 일의 자리의 수가 9
8끗: 땡이 아니고, 두 패를 더했을 때 일의 자리의 수가 8
...
1끗: 땡이 아니고, 두 패를 더했을 때 일의 자리의 수가 1
0끗: 땡이 아니고, 두 패를 더했을 때 일의 자리의 수가 0
3.영학이가 이길 확률을 구한다.
4.영학이가 이길 확률을 소수점 이하 셋 째 자리까지 반올림해서 출력(소수점 3자리까지가 0이더라도 다 작성)

## 💡알고리즘 설계

1. 카드 리스트 생성: 1-10까지, 각 2장씩 총 20징
2. 족보 판단 함수 필요: 점수를 반환하도록 하기

```jsx
족보는 강한 순서대로 다음과 같다.
10 띵 (10, 10)
9 띵 (9, 9)
2 띵 (2, 2)
1 띵 (1, 1)
9 끗 (두 수의 합이 9)
8 끗 (두 수의 합이 8)
1 끗 (두 수의 합이 1)
0 끗 (두 수의 합이 10)
```

1. 모든 경우의 수 탐색
- 상대가 가져갈 수 있는 두 장의 카드 조합 확인
- 현재 내가 가진 카드(A,B)를 제외한 나머지 카드들로 상대의 모든 조합을 만들어서 결과를 계산
- 상대가 가져갈 수 있는 조합의 가짓수를 고려하여 승률 계산
1. 확률 계산
- 이기는 경우의 수 / 전체 가능한 상대 카드 조합 수
- 소수점 셋째 자리까지 출력(toFixed 사용)

## 💡코드

```jsx
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
```

로직 순서 진행

1️⃣**카드 덱 만들기**

→ 1~10 숫자가 2장씩 총 20장

2️⃣**내가 가진 카드 제거**

3️⃣**각 카드 조합의 족보 판별**

4️⃣**상대 카드 조합을 만들고 승리 횟수 계산**

5️⃣**승률 계산 및 소수점 3자리 출력**

## 💡개념

핵심 개념:

- 게임 규칙과 족보 시스템을 정확히 이해해야 함.
- 상대가 가질 수 있는 모든 카드 조합을 고려하여 확률을 계산해야 함

- 가장 적절한 접근법:

완전 탐색

- 상대의 모든 카드 조합을 확인하면서 승리 횟수를 세는 방식이 가장 직관적이고 구현이 쉬움.

그 이외 방법

- 완전 탐색 (브루트포스): 가장 직관적이고 쉬운 방법.
- 확률론적 접근: 경우의 수를 직접 계산하지만 비효율적.

## 💡 틀린 이유 & 틀린 부분 수정

- 시간 초과
- 모든 경우의 수 고려하지 않음

기존 코드

```jsx
for (let i = 0; i < deck.length; i++) {
    for (let j = 0; j < deck.length; j++) {  // ❌ 중복 계산 발생
        if (i === j) continue;  // ❌ 중복 제거 시 실수
        let opponentRank = getRank(deck[i], deck[j]);
        if (myRank > opponentRank) winCount++;
        totalCount++;
    }
}
```

수정 코드

```jsx
for (let i = 0; i < deck.length; i++) {
    for (let j = i + 1; j < deck.length; j++) {  // ✅ 중복 제거
        let opponentRank = getRank(deck[i], deck[j]);
        if (myRank > opponentRank) winCount++;
        totalCount++;
    }
}

```

- 확률 계산 시 0으로 나누는 실수

기존 코드

```jsx
return (winCount / totalCount).toFixed(3);  // ❌ totalCount가 0일 수도 있음

```

수정 코드

```jsx
return totalCount === 0 ? "0.000" : (winCount / totalCount).toFixed(3);

```

- getRank() 함수에서 숫자 비교 실수: 카드 조합에 따라 **강한 족보일수록 높은 점수를 반환해야함**

기존 코드

```jsx
function getRank(x, y) {
    if (x === 10 && y === 10) return 10;
    if (x === 9 && y === 9) return 9;
    if (x === 2 && y === 2) return 2;
    if (x === 1 && y === 1) return 1;
    let sum = (x + y) % 10;
    if (sum === 9) return 4;
    if (sum === 8) return 3;
    if (sum === 1) return 2;
    if (sum === 0) return 1;
    return 0;
}

```

수정 코드

```jsx
function getRank(x, y) {
    if (x === 10 && y === 10) return 8;
    if (x === 9 && y === 9) return 7;
    if (x === 2 && y === 2) return 6;
    if (x === 1 && y === 1) return 5;
    let sum = (x + y) % 10;
    if (sum === 9) return 4;
    if (sum === 8) return 3;
    if (sum === 1) return 2;
    if (sum === 0) return 1;
    return 0;
}

```

## 💡느낀점

오늘 시간이 급박하여 풀다보니 잘 안풀렸고... 계속 집중이 안됐다.

결국 시간 초과 + 답을 보고 나서야 어느정도 감이 왔다.

다음부터 충분히 여유를 두고 풀어야 할듯하다.
내일 반드시 한 번 더 풀어보아야겠다..