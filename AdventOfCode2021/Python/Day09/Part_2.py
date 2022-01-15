def get_inputs():
    """
    function to read in list of ints from a txt file
    """

    # read in file
    with open("inputs.txt") as f:
        contents = f.read()

    contents = contents.split('\n')

    # format data
    data = [list('0'*(len(contents[0])+2))]
    for content in contents:

        # format all 9->0 and all other ints->1
        for i in range(0,9):
            content = content.replace(str(i),'1')
        content='0'+content.replace('9','0')+'0'

        data.append(list(content))
    data.append(list('0'*(len(contents[0])+2)))

    # convert all to int
    for i in range(0,len(data)):
        data[i] = [int(x) for x in data[i]]

    return data


def check_adj(data,sel_row,sel_col,basin):
    """
    checks if adjacent nums are 1 in array
    """

    # check r
    if data[sel_row][sel_col+1] == 1:
        basin += 1
        # remove counted basin from data to prevent double counting
        data[sel_row][sel_col+1] = 0
        basin = check_adj(data,sel_row,sel_col+1,basin)

    # check l
    if data[sel_row][sel_col-1] == 1:
        basin += 1
        # remove counted basin from data to prevent double counting
        data[sel_row][sel_col-1] = 0
        basin = check_adj(data,sel_row,sel_col-1,basin)

    # check u
    if data[sel_row+1][sel_col] == 1:
        basin += 1
        # remove counted basin from data to prevent double counting
        data[sel_row+1][sel_col] = 0
        basin = check_adj(data,sel_row+1,sel_col,basin)

    # check d
    if data[sel_row-1][sel_col] == 1:
        basin += 1
        # remove counted basin from data to prevent double counting
        data[sel_row-1][sel_col] = 0
        basin = check_adj(data,sel_row-1,sel_col,basin)

    return basin


def main():
    """
    main function
    """ 
    data = get_inputs()

    # find groups of 1 connected by adjacent moves
    basins = []
    for row_ind in range(1,len(data)-1):
        for col_ind in  range(1,len(data[0])-1):
            basin = 0
            if int(data[row_ind][col_ind]) == 1:
                basin += 1
                # remove counted basin from data to prevent double counting
                data[row_ind][col_ind] = 0
                basins.append(check_adj(data,row_ind,col_ind,basin))
    
    # check none have been missed
    assert(sum([sum(d) for d in data])==0)

    largest_basins = []

    for i in range(0,3):
        largest_basins.append(max(basins))
        basins.remove(max(basins))

    ans = largest_basins[0]*largest_basins[1]*largest_basins[2]
    print(f'Answer = {ans}')
    return ans


if __name__ == "__main__":
    main()