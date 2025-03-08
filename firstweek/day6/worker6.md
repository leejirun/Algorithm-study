## 💡문제 분석 요약
1.N*M 크기의 직사각형이 주어진다. (각 칸에는 한 자리 숫자가 들어 있다.)
2.4개의 꼭짓점에 같은 숫자가 있는 가장 큰 정사각형을 찾아야 한다.
3.정사각형은 반드시 행과 열에 평행해야 한다,
4.출력값: 찾은 정사각형의 넒이(한변의 길이의 제곱)

## 💡알고리즘 설계
1.입력값 받기
-N,M을 받는다.
-N개의 줄을 배열에 저장한다.
2.정사각형 크기 1부터 점점 키우면서 확인
-모든 (i,j) 좌표에서 가능한 최대 크기의 정사각형을 탐색
-size를 0부터 시작해 늘려가며 확인
-4개의 꼭짓점 숫자가 같은지 비교
-가능한 가장 큰 정사각형의 크기를 저장
3.출력하기

## 💡코드
`
const readline = require("readline");

// 입력을 받기 위한 설정
const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout,
});

let input = [];

rl.on("line", (line) => {
    input.push(line);
}).on("close", () => {
    // 1️⃣ 입력 처리
    const [N, M] = input[0].split(" ").map(Number); // N: 행 개수, M: 열 개수
    const grid = input.slice(1); // 숫자가 들어있는 보드 배열

    let maxSize = 1; // 찾을 수 있는 가장 큰 정사각형 크기 (최소 1)

    // 2️⃣ 가능한 정사각형 크기 찾기
    for (let size = Math.min(N, M); size > 0; size--) { // 가장 큰 변 크기부터 탐색
        for (let i = 0; i <= N - size; i++) { // 시작 행
            for (let j = 0; j <= M - size; j++) { // 시작 열
                // 네 꼭짓점 숫자 확인
                let num = grid[i][j]; // 좌상단 숫자
                if (
                    grid[i][j + size - 1] === num && // 우상단
                    grid[i + size - 1][j] === num && // 좌하단
                    grid[i + size - 1][j + size - 1] === num // 우하단
                ) {
                    maxSize = size; // 최대 정사각형 크기 업데이트
                    console.log(size * size); // 정사각형 넓이 출력 후 종료
                    process.exit(0);
                }
            }
        }
    }

    // 3️⃣ 결과 출력
    console.log(maxSize * maxSize);
    process.exit(0);
});
`

## 💡개념
1.행과 열에 평행한 정사각형 조건을 기억하자
`
if (
    board[i][j] === board[i][j + size] &&
    board[i][j] === board[i + size][j] &&
    board[i][j] === board[i + size][j + size]
) {
    maxSize = Math.max(maxSize, (size + 1) ** 2);
}
`
- 꼭짓점이 같은 숫자인지 확인할 때, 대각선까지 확인하지 않도록 주의한다.

## 💡 틀린 이유 & 틀린 부분 수정
1.정사각형 크기 계산 실수
`
let area = size ** 2; //size가 1부터 시작하는 걸 놓침!ㅠ
`
- size가 0부터 시작하므로, 정사각형 크기는 (size+1) * (size+1)이여야 한다.

`
let area = (size + 1) ** 2; // => (size+1) 크기로 계산한다.
`

## 💡느낀점