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

    # skip line if incomplete
    if not any(check_if_corrupt):
        return 0

    # find 1st illegal char    
    checks = [')',']','}','>']
    for char in line:
        if char in checks:
            illegal = char
            break

    # allocate point
    if illegal == ')':
        point = 3
    if illegal == ']':
        point = 57
    if illegal == '}':
        point = 1197
    if illegal == '>':
        point = 25137

    return point


def main():
    """
    main function
    """ 
    data = get_inputs()
 
    points = 0
    for line in data:
        points += check_line(line)

    ans = points
    print(f'Answer = {ans}')
    return ans


if __name__ == "__main__":
    main()