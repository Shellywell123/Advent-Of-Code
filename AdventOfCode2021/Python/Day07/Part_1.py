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

    fuel_usage = []

    # calc fuel useages
    for hor_pos in range(min(inputs),max(inputs)):
        fuel = 0
        for crab in inputs:
            fuel += abs(hor_pos-crab)
        fuel_usage.append(fuel)

    # find min fuel usage
    ans = min(fuel_usage)
    print(f'Answer = {ans}')
    return ans


if __name__ == "__main__":
    main()