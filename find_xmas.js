export function FindXmas(s, f) {
    // Split the input string by lines
    let lines = s.trim().split("\n");

    // Convert the lines to a 2D grid of characters
    let grid = [];
    for (let i = 0; i < lines.length; i++) {
        grid[i] = lines[i].split('');
    }

    let count = 0;
    const directions = [
        [0, 1],  // horizontal right
        [0, -1], // horizontal left
        [1, 0],  // vertical down
        [-1, 0], // vertical up
        [1, 1],  // diagonal down-right
        [-1, -1],// diagonal up-left
        [1, -1], // diagonal down-left
        [-1, 1], // diagonal up-right
    ];

    for (let i = 0; i < grid.length; i++) {
        for (let j = 0; j < grid[i].length; j++) {
            for (let dir of directions) {
                let search = true;
                for (let k = 0; k < f.length; k++) {
                    let nx = i + k * dir[0];
                    let ny = j + k * dir[1];
                    if (nx < 0 || ny < 0 || nx >= grid.length || ny >= grid[0].length || grid[nx][ny]!== f[k]) {
                        search = false;
                    }
                }
                if (search) {
                    count++;
                }
            }
        }
    }

    return `Xmas Appears : ${count} times`;
}

//======part two=====

function checkPattern(grid, word, x, y, dx, dy) {
    let length = word.length;
    for (let k = 0; k < length; k++) {
        let nx = x + k * dx;
        let ny = y + k * dy;
        if (grid[nx][ny]!= word[k]) {
            return false;
        }
    }
    return true;
}

function isWithinBounds(grid, x, y, dx, dy, length) {
    for (let k = 0; k < length; k++) {
        let nx = x + k * dx;
        let ny = y + k * dy;
        if (nx < 0 || ny < 0 || nx >= grid.length || ny >= grid[0].length) {
            return false;
        }
    }
    return true;
}

export function FindXMASPattern(s, f) {
    let lines = s.trim().split("\n");
    let grid = lines.map(line => line.split(''));
    let count = 0;
    let length = f.length;

    for (let i = 0; i < grid.length; i++) {
        for (let j = 0; j < grid[i].length; j++) {
            if (isWithinBounds(grid, i, j, 1, 1, length) && checkPattern(grid, f, i, j, 1, 1)) {
                if ((isWithinBounds(grid, i, j+length-1, 1, -1, length) && checkPattern(grid, f, i, j+length-1, 1, -1)) ||
                (isWithinBounds(grid, i+length-1, j, -1, 1, length) && checkPattern(grid, f, i+length-1, j, -1, 1))) {
                    count++;
                }
            }
            if (isWithinBounds(grid, i, j, 1, -1, length) && checkPattern(grid, f, i, j, 1, -1)) {
                if ((isWithinBounds(grid, i, j-length+1, 1, 1, length) && checkPattern(grid, f, i, j-length+1, 1, 1)) ||
                (isWithinBounds(grid, i+length-1, j, -1, -1, length) && checkPattern(grid, f, i+length-1, j, -1, -1))) {
                    count++;
                }
            }
            if (isWithinBounds(grid, i, j, -1, 1, length) && checkPattern(grid, f, i, j, -1, 1)) {
                if ((isWithinBounds(grid, i, j+length-1, -1, -1, length) && checkPattern(grid, f, i, j+length-1, -1, -1)) ||
                (isWithinBounds(grid, i-length+1, j, 1, 1, length) && checkPattern(grid, f, i-length+1, j, 1, 1))) {
                    count++;
                }
            }
            if (isWithinBounds(grid, i, j, -1, -1, length) && checkPattern(grid, f, i, j, -1, -1)) {
                if ((isWithinBounds(grid, i, j-length+1, -1, 1, length) && checkPattern(grid, f, i, j-length+1, -1, 1)) ||
                (isWithinBounds(grid, i-length+1, j, 1, -1, length) && checkPattern(grid, f, i-length+1, j, 1, -1))) {
                    count++;
                }
            }
        }
    }
    return `X-MAS appears ${count/2} times`;
}