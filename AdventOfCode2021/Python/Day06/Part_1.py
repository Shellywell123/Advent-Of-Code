def get_inputs():
    """
    function to read in list of ints from a txt file
    """

    # read in file
    with open("inputs.txt") as f:
        contents = f.read()

    # format inputs
    inputs = []
    for num in contents.split(','):
        inputs.append(int(num))
    return inputs


def main():
    """
    main function
    """ 
    inputs = get_inputs()
    inputs = [4]

    num_of_days = 80
    for day in range(0,num_of_days-1):
        for i in range(0,len(inputs)):
            if inputs[i] == 0:
                inputs[i] = 7
            inputs[i] -= 1
            if inputs[i] == 0:
                inputs.append(9)

    ans = len(inputs)
    print(f'Answer = {ans}')
    return ans


if __name__ == "__main__":
    main()