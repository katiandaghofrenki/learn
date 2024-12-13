# Define a function to check if a point (i, j) is within the grid
def InGrid(i, j, n):
    # Return True if the point is within the grid boundaries (0 <= i, j < n), False otherwise
    return (0 <= i < n) and (0 <= j < n)

# Define a function to count the number of free edges for a point (i, j) in the grid
def CountFree(i, j, plot, grid, n, dd):
    # Initialize the count of free edges to 0
    ans = 0
    # Iterate over all possible directions (up, down, left, right)
    for di, dj in dd:
        # Calculate the coordinates of the neighboring point
        ii, jj = i + di, j + dj
        # If the neighboring point is outside the grid, increment the count of free edges
        if not InGrid(ii, jj, n):
            ans += 1
        # If the neighboring point is within the grid but has a different value, increment the count of free edges
        elif grid[ii][jj]!= grid[i][j]:
            ans += 1
    # Return the total count of free edges
    return ans

# Define a function to calculate the perimeter of a plot (a connected component in the grid)
def Perimeter(plot, grid, n, dd):
    # Initialize the perimeter to 0
    ans = 0
    # Iterate over all points in the plot
    for i, j in plot:
        # Add the count of free edges for the current point to the perimeter
        ans += CountFree(i, j, plot, grid, n, dd)
    # Return the total perimeter
    return ans

# Define a function to calculate the number of sides of a plot (a connected component in the grid)
def Sides(plot, grid, n, dd):
    # Initialize sets to store the points on the top, bottom, left, and right edges of the plot
    up, down, left, right = (set() for _ in range(4))
    # Iterate over all points in the plot
    for i, j in plot:
        # If the point above is not in the plot, add the current point to the top edge set
        if (i - 1, j) not in plot:
            up.add((i, j))
        # If the point below is not in the plot, add the current point to the bottom edge set
        if (i + 1, j) not in plot:
            down.add((i, j))
        # If the point to the left is not in the plot, add the current point to the left edge set
        if (i, j - 1) not in plot:
            left.add((i, j))
        # If the point to the right is not in the plot, add the current point to the right edge set
        if (i, j + 1) not in plot:
            right.add((i, j))

    # Initialize the count of sides to 0
    count = 0
    # Iterate over the points on the top edge
    for i, j in up:
        # If the current point is also on the left edge, increment the count of sides
        if (i, j) in left:
            count += 1
        # If the current point is also on the right edge, increment the count of sides
        if (i, j) in right:
            count += 1
        # If the point above and to the left is on the right edge, increment the count of sides
        if (i - 1, j - 1) in right and (i, j) not in left:
            count += 1
        # If the point above and to the right is on the left edge, increment the count of sides
        if (i - 1, j + 1) in left and (i, j) not in right:
            count += 1

        # Iterate over the points on the bottom edge
    for i, j in down:
        # If the current point is also on the left edge, increment the count of sides
        if (i, j) in left:
            count += 1
        # If the current point is also on the right edge, increment the count of sides
        if (i, j) in right:
            count += 1
        # If the point above and to the left is on the right edge, increment the count of sides
        if (i + 1, j - 1) in right and (i, j) not in left:
            count += 1
        # If the point above and to the right is on the left edge, increment the count of sides
        if (i + 1, j + 1) in left and (i, j) not in right:
            count += 1
    # Return the total count of sides
    return count

# Define a function to calculate the perimeter and number of sides of each connected component in the grid
def CalculatePerimeter(input_string):
    # Split the input string into a list of rows
    grid = input_string.strip().split("\n")
    # Get the number of rows in the grid
    n = len(grid)
    # Define the possible directions (up, down, left, right)
    dd = [[1, 0], [0, 1], [-1, 0], [0, -1]]

    # Initialize a set to store the points that have been visited
    seen = set()
    # Initialize a list to store the connected components in the grid
    plots = []

    # Iterate over all points in the grid
    for i in range(n):
        for j in range(n):
            # If the point has already been visited, skip it
            if (i, j) in seen:
                continue
            # Initialize a stack to store the points to be visited
            stack = [(i, j)]
            # Initialize a list to store the current connected component
            plots.append([grid[i][j], []])

            # Iterate over the points in the stack
            while len(stack) > 0:
                # Get the next point from the stack
                ci, cj = stack.pop()
                # If the point has already been visited, skip it
                if (ci, cj) in seen:
                    continue
                # If the point is outside the grid, skip it
                if not InGrid(ci, cj, n):
                    continue
                # If the point has a different value than the current component, skip it
                if grid[ci][cj]!= grid[i][j]:
                    continue
                # Mark the point as visited
                seen.add((ci, cj))

                # Add the point to the current connected component
                plots[-1][1].append((ci, cj))
                # Iterate over the neighboring points
                for di, dj in dd:
                    # Calculate the coordinates of the neighboring point
                    ii, jj = ci + di, cj + dj
                    # Add the neighboring point to the stack
                    stack.append((ii, jj))

    # Calculate the sum of the perimeters of all connected components
    part1 = sum(len(plot[1]) * Perimeter(plot[1], grid, n, dd) for plot in plots)
    # Calculate the sum of the number of sides of all connected components
    part2 = sum(len(plot[1]) * Sides(plot[1], grid, n, dd) for plot in plots)

    # Return the results
    return part1, part2

# usage
# input_string = """copy input to here"""
# part1, part2 = CalculatePerimeter(input_string)
# print(f"Part 1: {part1}")
# print(f"Part 2: {part2}")