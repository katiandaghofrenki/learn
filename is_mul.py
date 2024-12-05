import re

def is_mul(s):
    total1 = 0  # for part1
    total2 = 0  # for part2
    do_seen = True
    re_mul = re.compile(r'mul\((\d{1,3}),(\d{1,3})\)')
    re_do = re.compile(r'do\(\)')
    re_dont = re.compile(r"don't\(\)")

    # Find all matches and their submatches
    found = re_mul.findall(s)

    for match in found:
        # Extract the numbers from the match
        x = int(match[0])
        y = int(match[1])

        # Calculate (x * y) and add to total
        total1 += x * y

    segments = re.findall(r'do\(\)|mul\(\d{1,3},\d{1,3}\)|don\'t\(\)', s)

    for segment in segments:
        if re_dont.match(segment):
            # Reset if 'don't()' is found
            do_seen = False
        elif re_do.match(segment):
            # Enable counting after 'do()' is found
            do_seen = True
        elif do_seen and re_mul.match(segment):
            match = re_mul.match(segment)
            x = int(match.group(1))
            y = int(match.group(2))
            total2 += x * y

    return total1, total2
