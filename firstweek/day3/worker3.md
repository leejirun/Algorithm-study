#first
문제 풀이 날짜: 2025-03-05(wed)
문제 유형: 브루트포스
문제 레벨: 브론즈2 
문제명: 유레카 이론
문제 링크: https://www.acmicpc.net/problem/10448
정답 여부: 실패

💡문제 분석 요약
- 표준입력 필요, 종료시 값 출력
- 테스트케이스: 중보고 허용한 자연수 n개,  1000 이하
- 삼각수 : Tn = n(n+1) /2
- 출력: 테스트케이스 3개의 삼각수의 합 표현 - 0, 표현X - 1

💡알고리즘 설계
1.삼각수 계산을 통해 배열을 만든다: Tn = n(n+1) /2 
2.유레카 판별식 사용한다 : a+b+c===k
3.입력받아서 결과 출력만든다.

💡코드
`
const fs = require("fs");

triangular_num = []
for (let i=1; ;i++) {
    let t = (i*(i+1)) /2;
    if (t > 1000) break;
    triangular_num.push(t);
}

function eureka(k) {
    for (let a of triangular_num) {
        for (let b of triangular_num) {
            for (let c of triangular_num) {
                if (a + b+ c=== k) {
                    return 1;
                } else {
                    return 0;
                }
            }
        }
    }
}

const readline = require("readline");
const rl = readline.createInterface({input: process.stdin, output: process.stdout});

let input = [];
rl.on("line", (line) => {
    input.push(line.trim());
}).on("close", () => {
    let T = parseInt(input[0]);
    for (let i = 1; i <= T; i++) {
        let K = parseInt(input[i]); 
        console.log(eureka(K));
    }
    process.exit();
});
`

💡개념
[삼각수]
- 삼각수의 개념만 이해하면 사실 쉬운 문제이다!
- 삼각수: 첫번째와 두번째 수의 합을 나타내는 자연수이다.
- ex) 1+ 3+ 6 = 10
- 1000 이하인 자연수 3개의 숫자로 합을 구하면 된다.


💡 틀린 이유 & 틀린 부분 수정 
function eureka(k) {
    for (let a of triangular_num) {
        for (let b of triangular_num) {
            for (let c of triangular_num) {
                if (a + b + c === k) {
                    return 1;  // 세 개의 삼각수 합으로 표현 가능하면 1 반환
                }
            }
        }
    }
    return 0;  // 모든 경우를 확인한 후에도 없으면 0 반환
}

- return 값의 위치를 잘못 씀,  else { return 0; } 때문에 첫 번째 삼각수 조합만 검사하고, 나머지를 확인하지 않고 바로 종료됨

💡 다른 풀이 or 기억할정보 
[다른 풀이]

`
// 1. 1000 이하의 삼각수를 미리 구하고 Set에 저장
const tNumbers = [];
const triSet = new Set();
for (let i = 1; i < 50; i++) {
    let t = (i * (i + 1)) / 2;
    if (t > 1000) break;
    tNumbers.push(t);
    triSet.add(t); // Set에 저장
}

// 2. 유레카 수 판별 함수 (Set 활용)
function isEureka(k) {
    for (let a of tNumbers) {
        for (let b of tNumbers) {
            if (triSet.has(k - a - b)) return 1; // 나머지가 삼각수인지 확인
        }
    }
    return 0;
}
`
[정리]
- 반복문이 아닌 set()을 사용한 코드이다.


💡느낀점 or 기억할 정보
삼각수의 개념만 이해하면 풀기 쉬운 문제같다.
하지만 나는 다른 블로그들의 내용을 보고 삼중 for문을 해야하구나 하고 깨달았다. 다음에 삼각수 문제는 오늘과 같이 실수만 하지 않는다면 쉽게 풀 것 같다.