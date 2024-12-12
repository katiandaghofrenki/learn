from collections import defaultdict

class Node:
    def __init__(self, freq, row, col):
        self.freq = freq
        self.row = row
        self.col = col

def Frequency(input):
    lines = input.split("\n")
    nodes = defaultdict(list)
    for row in range(len(lines)):
        for col in range(len(lines[0])):
            if lines[row][col] != ".":
                nodes[lines[row][col]].append(Node(lines[row][col], row, col))

    antinodes = set()
    for freq in nodes:
        for node in nodes[freq]:
            for node2 in nodes[freq]:
                if node == node2:
                    continue
                dx = node2.col - node.col
                dy = node2.row - node.row
                destrow = node.row - dy
                destcol = node.col - dx
                if 0 <= destrow < len(lines) and 0 <= destcol < len(lines[0]):
                    antinodes.add((destrow, destcol))  # Adding as a tuple

    return len(antinodes)