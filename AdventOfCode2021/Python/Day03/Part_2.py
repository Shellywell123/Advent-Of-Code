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

def get_rate(rate,inputs):
    """
    calculate relevant rate given inputs
    """

    bit_len = len(inputs[0])

    # loop through bit inputs comparing bits to calculate rate
    for i in range(0,bit_len):

        # sum of bits at this index
        total = 0
        for input_ in inputs:
            total += int(input_[i])

        total_comp = (len(inputs))/2

        # determine most/least common
        if total >= total_comp:
            if rate == 'O2':
                rate_bit = '1'
            if rate == 'CO2':
                rate_bit = '0'

        if total < total_comp:
            if rate == 'O2':
                rate_bit = '0'
            if rate == 'CO2':
                rate_bit = '1'

        # cull inputs 
        inputs_keep = []
        for j in range(0,len(inputs)):
            if inputs[j][i] == rate_bit:
                inputs_keep.append(inputs[j])

        inputs = inputs_keep

        # set rate
        if len(inputs) == 1:
            rate = inputs[0]
            break

    return rate


def main():
    """
    main function
    """
    inputs = get_inputs()

    O2_rate = get_rate('O2',inputs)
    CO2_rate = get_rate('CO2',inputs)

    ans = int(O2_rate,2)*int(CO2_rate,2)
    print(f'Answer = {ans}')
    return ans


if __name__ == "__main__":
    main()