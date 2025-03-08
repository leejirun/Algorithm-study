## 💡문제 분석 요약
- 주어진 M x N 크기의 보드에서 8 x 8 크기의 체스판을 만들려고 함
- 체스판은 검은색(B)과 흰색(W)이 번갈아가며 배치되어 있어야 함
- 보드에서 가능한 모든 8 x 8 부분을 검사하여, 최소한의 색칠 변경으로 체스판을 만들 수 있는 경우를 찾는 것이 문제의 핵심
- 체스판은 두 가지 패턴으로 시작
왼쪽 위 칸이 W일 때 → (W 시작 패턴)
왼쪽 위 칸이 B일 때 → (B 시작 패턴)
-8*8 영역을 두 패턴과 비교해서 다시 칠해야 하는 칸 수의 최솟값을 찾는 것이 목표임

## 💡알고리즘 설계
1.입력값 받기 -> 저장
2.모든 가능한 8*8 영역을 검사
3.선택한 8*8 영역을 두 가지 패턴과 비교
4.모든 8*8 경우 중 최소값 출력

## 💡코드
`
function minRecolor(N, M, board) {
    let minCount = Infinity; // 최소 칠해야 할 개수 저장

    // W로 시작하는 정답 체스판과 B로 시작하는 정답 체스판 패턴
    const wChess = [
        "WBWBWBWB",
        "BWBWBWBW",
        "WBWBWBWB",
        "BWBWBWBW",
        "WBWBWBWB",
        "BWBWBWBW",
        "WBWBWBWB",
        "BWBWBWBW"
    ];
    const bChess = [
        "BWBWBWBW",
        "WBWBWBWB",
        "BWBWBWBW",
        "WBWBWBWB",
        "BWBWBWBW",
        "WBWBWBWB",
        "BWBWBWBW",
        "WBWBWBWB"
    ];

    // 가능한 모든 8x8 영역을 순회
    for (let i = 0; i <= N - 8; i++) {
        for (let j = 0; j <= M - 8; j++) {
            let wCount = 0; // W 시작 패턴과 비교할 때 칠해야 할 개수
            let bCount = 0; // B 시작 패턴과 비교할 때 칠해야 할 개수

            // 8x8 영역 검사
            for (let x = 0; x < 8; x++) {
                for (let y = 0; y < 8; y++) {
                    if (board[i + x][j + y] !== wChess[x][y]) wCount++; // W 패턴과 다르면 색칠 필요
                    if (board[i + x][j + y] !== bChess[x][y]) bCount++; // B 패턴과 다르면 색칠 필요
                }
            }

            // 최소값 갱신
            minCount = Math.min(minCount, wCount, bCount);
        }
    }

    console.log(minCount);
}

// 예제 입력 실행
const input = [
    "10 13",
    "BBBBBBBBWBWBW",
    "BBBBBBBBBWBWB",
    "BBBBBBBBWBWBW",
    "BBBBBBBBBWBWB",
    "BBBBBBBBWBWBW",
    "BBBBBBBBBWBWB",
    "BBBBBBBBWBWBW",
    "BBBBBBBBBWBWB",
    "WWWWWWWWWWBWB",
    "WWWWWWWWWWBWB"
];

// 입력값 처리
const [N, M] = input[0].split(" ").map(Number);
const board = input.slice(1);
minRecolor(N, M, board);
`

## 💡개념
1) 2차원 배열
`
function minRecolor(N, M, board) {
    let minCount = Infinity; // 최소 칠해야 할 개수 저장

    // W로 시작하는 정답 체스판과 B로 시작하는 정답 체스판 패턴
    const wChess = [
        "WBWBWBWB",
        "BWBWBWBW",
        "WBWBWBWB",
        "BWBWBWBW",
        "WBWBWBWB",
        "BWBWBWBW",
        "WBWBWBWB",
        "BWBWBWBW"
    ];
    const bChess = [
        "BWBWBWBW",
        "WBWBWBWB",
        "BWBWBWBW",
        "WBWBWBWB",
        "BWBWBWBW",
        "WBWBWBWB",
        "BWBWBWBW",
        "WBWBWBWB"
    ];

    // 가능한 모든 8x8 영역을 순회
    for (let i = 0; i <= N - 8; i++) {
        for (let j = 0; j <= M - 8; j++) {
            let wCount = 0; // W 시작 패턴과 비교할 때 칠해야 할 개수
            let bCount = 0; // B 시작 패턴과 비교할 때 칠해야 할 개수

            // 8x8 영역 검사
            for (let x = 0; x < 8; x++) {
                for (let y = 0; y < 8; y++) {
                    if (board[i + x][j + y] !== wChess[x][y]) wCount++; // W 패턴과 다르면 색칠 필요
                    if (board[i + x][j + y] !== bChess[x][y]) bCount++; // B 패턴과 다르면 색칠 필요
                }
            }

            // 최소값 갱신
            minCount = Math.min(minCount, wCount, bCount);
        }
    }

    console.log(minCount);
}

// 예제 입력 실행
const input = [
    "10 13",
    "BBBBBBBBWBWBW",
    "BBBBBBBBBWBWB",
    "BBBBBBBBWBWBW",
    "BBBBBBBBBWBWB",
    "BBBBBBBBWBWBW",
    "BBBBBBBBBWBWB",
    "BBBBBBBBWBWBW",
    "BBBBBBBBBWBWB",
    "WWWWWWWWWWBWB",
    "WWWWWWWWWWBWB"
];

// 입력값 처리
const [N, M] = input[0].split(" ").map(Number);
const board = input.slice(1);
minRecolor(N, M, board);
`
- 문자열 배열은 2차원 배열처럼 사용할 수 있다.
- 각 행은 문자열로 저장되므로 board[i][j] 형태로 인덱싱이 가능하다.
- 문자열은 불변이므로 값을 변경할 수 없다. ex) board[i][j] = 'w'; -> 불가능함

2) Math.min() 활용
- 개념: 여러 개의 값 중 최소값 찾기
`
let minValue = Math.min(10, 5, 8, 3);
console.log(minValue); // 3
`
- Meth.min(a,b,c, ...) 형식으로 사용

## 💡 틀린 이유 & 틀린 부분 수정
x

## 💡느낀점