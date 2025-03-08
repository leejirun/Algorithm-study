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
