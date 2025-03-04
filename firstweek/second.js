function findSevenDwarfs(heights) {
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


const heights = [20, 7, 23, 19, 10, 15, 25, 8, 13];
console.log(findSevenDwarfs(heights));
