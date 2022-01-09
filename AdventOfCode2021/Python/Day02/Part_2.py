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
        inputs.append(content.replace('\n','').split(' '))

    return inputs


def main():
    """
    main function
    """
    inputs = get_inputs()
    count = -1

    hor_pos = 0
    ver_pos = 0
    aim = 0

    # loop through inputs recording position and aim chanages
    for input_ in inputs:

        if input_[0] == 'forward':
            hor_pos += int(input_[1])
            ver_pos += int(input_[1])*aim

        if input_[0] == 'up':
            aim -= int(input_[1])

        if input_[0] == 'down':
            aim += int(input_[1])
            
    ans = ver_pos*hor_pos
    print(f'Answer = {ans}')
    return ans


if __name__ == "__main__":
    main()