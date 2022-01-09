def get_inputs():
    """
    function to read in list of ints from a txt file
    """

    # read in file
    with open("inputs.txt") as f:
        contents = f.readlines()

    # format inputs
    numbers =[]
    tables = []    
    table = []

    for i in range(0,len(contents)):
        if i == 0:
            for num in contents[i].split(','):
                if num != ' ' or num != '\n':
                    numbers.append(int(num))
            continue

        if len(contents[i]) < 2:
            tables.append(table)
            table = []
        else:
            row = []
            for num in contents[i].replace('\n','').split(' '):
                if num != '':
                    row.append(int(num))
            table.append(row)
    tables.append(table)

    del tables[0]
    inputs = [numbers,tables]
    return inputs


def bingo(numbers_called,table):
    """
    check if a table has bingo
    """

    def check(line,numbers_called):
        """
        checks a row/col for bingo
        """
        numbers_called_ = numbers_called[:]
        count = 0

        for num in line:
            if num in numbers_called_:
                count +=1 
                numbers_called_.remove(num)

        if count == len(line):
            return True

        return False


    # check minimum amount of nums has been called
    if len(numbers_called) >= len(table[0]):

        # check rows
        for row in table:
            if check(row,numbers_called):
                return True

        # check cols
        for i in range(0,len(table[0])):
            col = []
            for row in table:
                col.append(row[i])

            if check(col,numbers_called):
                return True


def main():
    """
    main function
    """ 
    numbers,tables = get_inputs()

    # play bingo
    for i in range(0,len(numbers)):
        numbers_called = numbers[:i+1]
        for table in tables:
            if bingo(numbers_called,table) == True:
                break
        if bingo(numbers_called,table) == True:
            break

    # find unmarked number
    unmarked_numbers = []
    for row in table:
        for num in row:
            if num not in numbers_called:
                unmarked_numbers.append(num)

    # sum unmarked numbers
    ans = sum(unmarked_numbers) * numbers_called[-1]
    print(f'Answer = {ans}')
    return ans


if __name__ == "__main__":
    main()