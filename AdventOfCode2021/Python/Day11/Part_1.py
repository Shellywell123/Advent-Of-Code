def get_inputs():
    """
    function to read in list of ints from a txt file
    """

    # read in file
    with open("tests.txt") as f:
        contents = f.read()

    # format data
    data = []
    for content in contents.split('\n'):
        data.append([int(x) for x in list(content)])



    return data


def check_adj(data,flashes,sel_row,sel_col,flashed=[]):
    """
    checks if adjacent nums are 1 in array
    """
    new_data = data[:]

    for row_check in [-1,0,+1]:
        for col_check in [-1,0,+1]:
            if not row_check==col_check==9:
                try:
                    col = sel_col+col_check
                    row = sel_row+row_check
                    if [row,col] not in flashed:
                 #       data[row][col] +=1
                        if data[row][col] > 9:
                            data[row][col] = 0
                            flashed.append([row,col])
                            flashes +=1
                            #data,flashes = check_adj(data,flashes,row,col,flashed=flashed)
                except:
                    pass

    return data,flashes


def main():
    """
    main function
    """ 
    data = get_inputs()

    flashes = 0

    print('\n')
    print(f'After step 0, flases {flashes}')
    for d in data:
        stri = ''
        for num in d:
            stri+=str(num)
        print(stri)

    step = 0
    steps = 2

    while(step <= steps):
        print(step)
        # update for day
        data = [[number+1 for number in d] for d in data]
        flashes = 0
        for row_ind in range(1,len(data)-1):
            for col_ind in  range(1,len(data[0])-1):

                if data[row_ind][col_ind] >= 9:
                    data,flashes_ = check_adj(data,flashes,row_ind,col_ind)
                    flashes = flashes + flashes_
        if flashes == 0:
            step = step + 1

        print('\n')
        print(f'After step {step}, flases {flashes}')
        for d in data:
            stri = ''
            for num in d:
                stri+=str(num)
            print(stri)

        

    ans = 0
    print(f'Answer = {ans}')
    return ans


if __name__ == "__main__":
    main()