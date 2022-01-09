def get_inputs():
    """
    function to read in list of ints from a txt file
    """

    # read in file
    with open("inputs.txt") as f:
        contents = f.readlines()

    # format inputs
    inputs = []
    for content in contents:
        inputs.append(int(content.replace('\n','')))

    return inputs


def main():
    """
    main function
    """
    inputs = get_inputs()
    count = -1
    input_prev = 0

    # loop through inputs and count increases in values
    for input_ in inputs:
        if input_ > input_prev:
            count += 1
        input_prev = input_

    ans = count
    print(f'Answer = {ans}')
    return ans


if __name__ == "__main__":
    main()