def get_inputs():
    """
    function to read in list of ints from a txt file
    """

    # read in file
    with open("inputs.txt") as f:
        contents = f.read()

    inputs = []
    outputs = []

    # format inputs and outputs
    for content in contents.split('\n'):
        a,b = content.split(' | ')
        for input_ in a.split(' '):
            inputs.append(input_)
        for output in b.split(' '):
            outputs.append(output)

    return inputs,outputs


def main():
    """
    main function
    """ 
    len_counts = [0,0,0,0,0,0,0,0]
    inputs,outputs = get_inputs()

    # count output comb lengths
    for i in range(0,len(outputs)):
        len_num = len(outputs[i])
        len_counts[len_num] += 1

    # sum chosen lengths
    ans = len_counts[2]+len_counts[3]+len_counts[4]+len_counts[7]
    print(f'Answer = {ans}')
    return ans


if __name__ == "__main__":
    main()