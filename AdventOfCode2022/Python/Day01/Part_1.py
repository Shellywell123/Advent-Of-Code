def get_inputs(file):
    """
    function to read in list of ints from a txt file
    """

    with open(file) as f:
        data =(f.read()).split("\n")

    return data 


def main():
    """
    main function
    """ 
    data = get_inputs("inputs.txt")

    totals = []

    elf = 0
    for i in range(0,len(data)):
        cal = data[i]
        if cal == "" or i == len(data)-1:
            # end of elf 
            totals.append(elf)
            # starting new elf
            elf = 0
        else:
            elf += int(cal)
            
    print(max(totals))

    return 0


if __name__ == "__main__":
    main()