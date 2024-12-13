def Trailhead(input):
    lines = input.split("\n")
    dirs = [[0,1], [1,0], [0,-1], [-1,0]]
    trailheads = []
    for row in range(len(lines)):
        for col in range(len(lines[0])):
            if lines[row][col] == "0":
                trailheads.append( (row, col))
    total = 0
    for th in trailheads:
        v = []
        summits = set()
        currentrow = th[0]
        currentcol = th[1]
        q = [ th ]
        while len(q) >0:
            currentrow, currentcol = q.pop(0)
            v.append((currentrow, currentcol))
            if lines[currentrow][currentcol] == "9":
                summits.add((currentrow, currentcol))
            else:
                currentheight = int(lines[currentrow][currentcol])
                for dir in dirs:
                    testrow = currentrow + dir[0]
                    testcol = currentcol + dir[1]
                    if not (0<=testrow<len(lines) and 0<=testcol<len(lines[0])):
                        continue
                    testheight = int(lines[testrow][testcol])
                    if testheight == currentheight + 1:
                        if (testrow, testcol) not in q:
                            if (testrow, testcol) not in v:
                                q.append((testrow,testcol))
        # return len(summits)
        total += len(summits)
    return total

def Trailhead2(input):
    lines = input.split("\n")
    dirs = [[0,1], [1,0], [0,-1], [-1,0]]
    trailheads = []
    for row in range(len(lines)):
        for col in range(len(lines[0])):
            if lines[row][col] == "0":
                trailheads.append( (row, col))
    total = 0
    for th in trailheads:
        summits = set()
        currentrow = th[0]
        currentcol = th[1]
        q = [ (th,[]) ]
        while len(q) >0:
            ((currentrow, currentcol),path) = q.pop(0)
            if lines[currentrow][currentcol] == "9":
                summits.add(((currentrow, currentcol),tuple(path)))
            else:
                currentheight = int(lines[currentrow][currentcol])
                for dir in dirs:
                    testrow = currentrow + dir[0]
                    testcol = currentcol + dir[1]
                    if not (0<=testrow<len(lines) and 0<=testcol<len(lines[0])):
                        continue
                    testheight = int(lines[testrow][testcol])
                    if testheight == currentheight + 1:
                        if (testrow, testcol) not in q:
                                q.append(((testrow,testcol),path+[(currentrow,currentcol)]))
        total += len(summits)
    return total