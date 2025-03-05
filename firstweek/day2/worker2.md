#first
문제 풀이 날짜: 2025-03-04(tue)
문제 유형: 브루트포스
문제 레벨: 브론즈2 
문제명: 일곱난쟁이
문제 링크: https://www.acmicpc.net/problem/2309
정답 여부: 성공

💡문제 분석 요약
-총 9명의 난쟁이가 주어짐
-이 중 7명의 키 합이 100이 되도록 찾아야 함
-가능한 정답이 여러 개일 수 있으며, 하나만 출력하면 됨
-출력은 7명의 키를 오름차순으로 정렬해야 함

💡알고리즘 설계
1.입력 받기
- 9명의 난쟁이 키를 배열로 저장한다.

2.전체 키 합 구하기
- 9명의 난쟁이 키의 총합을 구한다.
- 목표는 이 총합에서 2명의 키를 제외했을 때 100이 되는 경우를 찾는 것

3.제외할 두 난쟁이 찾기
-9명 중 2명을 선택해 총합에서 빼서 100이 되는지 확인한다.
-이중 루프 (이중 for문) 을 이용하면 쉽게 찾을 수 있음.
-즉, totalSum - (난쟁이A + 난쟁이B) === 100 을 만족하는 두 난쟁이를 찾는다.

4.결과 출력
-제외한 두 난쟁이를 빼고 남은 7명의 난쟁이만 남긴다.
-오름차순 정렬 후 출력한다.

💡코드

`function findSevenDwarfs(heights) {
    let totalSum = heights.reduce((acc, cur) => acc + cur, 0);

    for (let i = 0; i < 9; i++) {
        for (let j = i + 1; j < 9; j++) {
            if (totalSum - (heights[i] + heights[j]) === 100) {
                let result = heights.filter((_, idx) => idx !== i && idx !== j);
                return result.sort((a, b) => a - b);
            }
        }
    }
}
`

💡개념
1.반복문
- 이중 반복문 예제
`
for (let i = 0; i < 9; i++) {
    for (let j = i + 1; j < 9; j++) {
        console.log(i, j);  // i와 j의 조합 확인
    }
}
`
- j=i+1로 설정해야하는 이유: 문제에서 두 개의 난쟁이를 선택하는 모든 조합을
고려해야하는데, j가 0부터 시작하면 (i, j) 순서가 바뀌고, 중복된 조합이 발생한다. 
[잘못된 경우 예시]
`
for (let i = 0; i < 9; i++) {
    for (let j = 0; j < 9; j++) { 
        if (i !== j) { // 같은 난쟁이는 제외하기
            console.log(i, j);
        }
    }
}
`
위와 같이 하면 결과는 아래와 같다.
`
(0,1) (0,2) (0,3) ... (1,0) (1,2) (1,3) ... (2,0) (2,1) (2,3) ...
`
중복된 조합이 발생하고 있는 걸 확인할 수 있다.

이 부분에서 처음에 j=i만 하다가 답이 나오지 않아 j=i+1로 수정해서 진행했다.

[올바른 경우 예시]
`
for (let i = 0; i < 9; i++) {
    for (let j = i + 1; j < 9; j++) { 
        console.log(i, j);
    }
}
`
결과는
`
(0,1) (0,2) (0,3) (0,4) ... 
(1,2) (1,3) (1,4) ...
(2,3) (2,4) ...
`
이렇게 중복되지 않는다.


2.배열의 특정 요소 제거: filter()
`
const arr = [1, 2, 3, 4, 5];
const newArr = arr.filter((num) => num !== 2 && num !== 4);
console.log(newArr); // [1, 3, 5]
`

3. 배열 정렬: sort()
`
const numbers = [20, 7, 23, 19, 10, 15, 25, 8, 13];
numbers.sort((a, b) => a - b);
console.log(numbers);
`
- sort((a, b) => a - b)를 사용하면 오름차순 정렬이 가능!


💡 틀린 이유 & 틀린 부분 수정 
x

💡 다른 풀이 or 기억할정보 
function findSevenDwarfs(heights) {
    heights.sort((a, b) => a - b); // 정렬
    let totalSum = heights.reduce((acc, cur) => acc + cur, 0);
    
    let left = 0, right = 8; // 투 포인터 설정
    
    while (left < right) {
        let sum = totalSum - (heights[left] + heights[right]);
        if (sum === 100) {
            return heights.filter((_, idx) => idx !== left && idx !== right);
        } else if (sum < 100) {
            right--; // 합이 작으면 right 줄이기
        } else {
            left++; // 합이 크면 left 증가
        }
    }
}

-gpt에게 물어본 다른 풀이 방식이다.
정렬 후 투 포인터를 활용해 불필요한 연산을 줄이는 방법 중 하나라는데 솔직히 용어부터 이해가 잘 안간다.

시간 복잡도: O(N log N) (정렬) + O(N) (투 포인터) = O(N log N)

풀이 설명:
1.먼저 배열을 정렬 (O(N log N))
2.left와 right 포인터를 사용해 두 개의 난쟁이를 찾음 (O(N))
3.총합에서 두 개를 뺐을 때 100이 되는지 검사
4.맞으면 두 개를 제외하고 반환


💡느낀점 or 기억할 정보
첫 날 만큼의 어려움은 없었다. 어제 개념을 조금 이해가 되어 풀만 했다.
하지만, 그렇다고 시간이 단축되거나 하진 않았다.
이중 for문을 작성하다가 꼬여서 몇 번 지웠다가 다시 작성하느라 고생을 했다.