def get_inputs():
    """
    function to read in list of ints from a txt file
    """

    # read in file
    with open("inputs.txt") as f:
        contents = f.read()

    # format data
    data = []
    for content in contents.split('\n'):
        data.append(content)

    return data


def check_line(line):
    """
    check line for corruption and calculate points
    """

    # reduce line down
    change=True
    while change==True:
        line_prev = line
        line = line.replace('()','')
        line = line.replace('[]','')
        line = line.replace('{}','')
        line = line.replace('<>','')

        if line==line_prev:
            change=False
        if line!=line_prev:
            change=True


    # check if corrupt line
    check1 = ')' in line
    check2 = ']' in line
    check3 = '}' in line
    check4 = '>' in line
    check_if_corrupt = (check1,check2,check3,check4)

    # skip line if corrupt
    if any(check_if_corrupt):
        return 0

    # complete line
    missings = ''
    for i in range(1,len(line)+1):
        opener = line[-i]
        if opener == '(':
            missings+=')'
        if opener == '[':
            missings+=']'
        if opener == '{':
            missings+='}'
        if opener == '<':
            missings+='>'

    # calculate points
    points = 0
    for missing in missings:
        if missing == ')':
            points = points*5 + 1
        if missing == ']':
            points = points*5 + 2
        if missing == '}':
            points = points*5 + 3
        if missing == '>':
            points = points*5 + 4

    return points


def main():
    """
    main function
    """ 
    data = get_inputs()
 
    # process lines
    points = []
    for line in data:
        point = check_line(line)
        if point > 0:
            points.append(point)

    # find middle points
    points = sorted(points)
    ans = points[int(len(points)/2)]
    print(f'Answer = {ans}')
    return ans


if __name__ == "__main__":
    main()