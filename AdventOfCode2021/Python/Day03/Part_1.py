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
        inputs.append(content.replace('\n',''))

    return inputs


def main():
    """
    main function
    """
    inputs = get_inputs()

    gamma_rate = ''
    epsilon_rate = ''

    bit_len = len(inputs)

    # loop through inputs recording position chanages
    for i in range(0,len(inputs[0])):
        total = 0
        for input_ in inputs:
            total += int(input_[i])

        if total >= bit_len/2:
            gamma_rate += '1'
            epsilon_rate += '0'

        if total < bit_len/2:
            gamma_rate += '0'
            epsilon_rate += '1'

    ans = int(gamma_rate,2)*int(epsilon_rate,2)
    print(f'Answer = {ans}')
    return ans


if __name__ == "__main__":
    main()