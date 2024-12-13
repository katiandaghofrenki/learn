def InGrid(i, j, n):
    return (0 <= i < n) and (0 <= j < n)

def CountFree(i, j, plot, grid, n, dd):
    ans = 0
    for di, dj in dd:
        ii, jj = i + di, j + dj  # Corrected here
        if not InGrid(ii, jj, n):
            ans += 1
        elif grid[ii][jj] != grid[i][j]:
            ans += 1
    return ans

def Perimeter(plot, grid, n, dd):
    ans = 0
    for i, j in plot:
        ans += CountFree(i, j, plot, grid, n, dd)
    return ans

def CalculatePerimeter(input_string):
    grid = input_string.strip().split("\n")
    n = len(grid)
    dd = [[1, 0], [0, 1], [-1, 0], [0, -1]]

    seen = set()
    plots = []

    for i in range(n):
        for j in range(n):
            if (i, j) in seen:
                continue

            stack = [(i, j)]
            plots.append([grid[i][j], []])

            while len(stack) > 0:
                ci, cj = stack.pop()
                if (ci, cj) in seen:
                    continue
                if not InGrid(ci, cj, n):
                    continue
                if grid[ci][cj] != grid[i][j]:
                    continue
                seen.add((ci, cj))

                plots[-1][1].append((ci, cj))
                for di, dj in dd:
                    ii, jj = ci + di, cj + dj
                    stack.append((ii, jj))

    ans = 0
    for c, plot in plots:
        ans += Perimeter(plot, grid, n, dd) * len(plot)
    return ans


# usage on main.py
# from learn.perimeters import CalculatePerimeter
# input := '''copy input to here'''
# result := CalculatePerimeter(input)
# print(result)