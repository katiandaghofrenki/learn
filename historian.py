def Historian(input_str):
    lines = input_str.strip().split("\n")
    
    grid = [list(line.strip()) for line in lines]

    char = '^'
    x, y, found = Found(grid, char)

    for _ in range(len(grid) * len(grid[0])):  # Ensure enough iterations
        # Initial movement loop
        while found:
            if x > 0 and grid[x-1][y] != '#':  # Ensure boundary check first
                grid[x][y] = 'X'
                x = x - 1  # Move to the new position
                grid[x][y] = char
            elif x > 0 and grid[x-1][y] == '#':
                grid[x][y] = '>'  # Mark the final position before hitting '#'
                char = '>'
                x, y, found = Found(grid, char)  # Re-find position
                break
            elif x <= 0:  # Check boundary to avoid out-of-bounds access
                grid[x][y] = 'X'
                found = False  # Stop the loop
                break
        
        # Continue moving right
                # Continue moving right
        while found:
            if y + 1 < len(grid[0]) and grid[x][y + 1] != '#':  # Ensure boundary check first
                grid[x][y] = 'X'
                y = y + 1  # Move to the new position
                grid[x][y] = char
            elif y + 1 < len(grid[0]) and grid[x][y + 1] == '#':
                grid[x][y] = 'v'  # Mark the final position before hitting '#'
                char = 'v'
                x, y, found = Found(grid, char)  # Re-find position
                break
            elif y + 1 >= len(grid[0]):  # Check boundary to avoid out-of-bounds access
                grid[x][y] = 'X'
                found = False  # Stop the loop
                break

        # Continue moving down
        while found:
            if x + 1 < len(grid) and grid[x + 1][y] != '#':  # Ensure boundary check first
                grid[x][y] = 'X'
                x = x + 1  # Move to the new position
                grid[x][y] = char
            elif x + 1 < len(grid) and grid[x + 1][y] == '#':
                grid[x][y] = '<'  # Mark the final position before hitting '#'
                char = '<'
                x, y, found = Found(grid, char)  # Re-find position
                break
            elif x + 1 >= len(grid):  # Check boundary to avoid out-of-bounds access
                grid[x][y] = 'X'
                found = False  # Stop the loop
                break

        # Continue moving left
        while found:
            if y > 0 and grid[x][y - 1] != '#':  # Ensure boundary check first
                grid[x][y] = 'X'
                y = y - 1  # Move to the new position
                grid[x][y] = char
            elif y > 0 and grid[x][y - 1] == '#':
                grid[x][y] = '^'  # Mark the final position before hitting '#'
                char = '^'
                x, y, found = Found(grid, char)  # Re-find position
                break
            elif y <= 0:  # Check boundary to avoid out-of-bounds access
                grid[x][y] = 'X'
                found = False  # Stop the loop
                break

    found_x = Found_X_Count(grid, 'X')
    return GridToString(grid), found_x

def Found(grid, char):
    for i, row in enumerate(grid):
        for j, col in enumerate(row):
            if col == char:
                return i, j, True
    return -1, -1, False

def Found_X_Count(grid, char):
    return sum(row.count(char) for row in grid)


def GridToString(grid):
    return "\n".join("".join(row) for row in grid)


def Historian2(input_str):
    # Convert input string into a 2D list (grid) of characters
    grid = [list(line) for line in input_str.strip().split("\n")]

    n = len(grid)  # Number of rows in the grid
    m = len(grid[0])  # Number of columns in the grid

    # Find the initial position of the character '^'
    found = False
    for i in range(n):
        for j in range(m):
            if grid[i][j] == "^":
                found = True
                break
        if found:
            break

    ii = i  # Initial row position of '^'
    jj = j  # Initial column position of '^'

    # Possible movement directions: up, right, down, left
    dd = [[-1, 0], [0, 1], [1, 0], [0, -1]]

    # Assess possible starting locations
    dir = 0  # Initial direction (up)
    og_seen = set()  # Set to store seen positions
    while True:
        og_seen.add((i, j))  # Add current position to the set

        # Calculate next position based on current direction
        next_i = i + dd[dir][0]
        next_j = j + dd[dir][1]

        # Check if the next position is within grid boundaries
        if not (0 <= next_i < n and 0 <= next_j < m):
            break

        # Change direction if the next position is a wall '#'
        if grid[next_i][next_j] == "#":
            dir = (dir + 1) % 4  # Rotate direction clockwise
        else:
            i, j = next_i, next_j  # Move to the next position

    def Will_Loop(oi, oj):
        if grid[oi][oj] == "#":
            return False  # Cannot place obstacle on a wall '#'
        
        grid[oi][oj] = "#"  # Temporarily place obstacle
        i, j = ii, jj  # Start from initial position of '^'

        dir = 0  # Initial direction (up)
        seen = set()  # Set to store seen positions with directions
        while True:
            if (i, j, dir) in seen:
                grid[oi][oj] = "."  # Reset the temporary obstacle
                return True  # Infinite loop detected
            seen.add((i, j, dir))  # Add current position and direction to the set

            # Calculate next position based on current direction
            next_i = i + dd[dir][0]
            next_j = j + dd[dir][1]

            # Check if the next position is within grid boundaries
            if not (0 <= next_i < n and 0 <= next_j < m):
                grid[oi][oj] = "."  # Reset the temporary obstacle
                return False  # Not an infinite loop

            # Change direction if the next position is a wall '#'
            if grid[next_i][next_j] == "#":
                dir = (dir + 1) % 4  # Rotate direction clockwise
            else:
                i, j = next_i, next_j  # Move to the next position

    ans = 0
    for oi, oj in og_seen:
        # Cannot place obstacle where the guard currently is
        if oi == ii and oj == jj:
            continue
        loop = Will_Loop(oi, oj)  # Check if placing obstacle creates an infinite loop
        ans += loop  # Increment count if an infinite loop is detected

    return ans