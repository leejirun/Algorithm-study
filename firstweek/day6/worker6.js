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
