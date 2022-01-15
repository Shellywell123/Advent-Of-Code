def get_inputs():
    """
    function to read in list of ints from a txt file
    """

    # read in file
    with open("inputs.txt") as f:
        contents = f.read()

    data = []

    # format inputs and outputs
    for content in contents.split('\n'):
        a,b = content.split(' | ')
        data.append([a.split(' '),b.split(' ')])

    return data


def decode(inputs,outputs):
    """
    inputs to generate keys that will decode the outputs
    """

    assert(len(inputs)==10)
    assert(len(outputs)==4)

    inputs = sorted(inputs,key=len)

    numbers = {}
    letters = {}

    ##########################################
    # decode numbers and letters
    ##########################################

    # find f
    if 'f' not in letters.keys():
        for char in 'abcdefg':
            if ''.join(inputs).count(char) ==9:
                letters['f'] = char
                break

    # find 1
    if 1 not in numbers.keys():
        for input_ in inputs:
            if len(input_) == 2:
                numbers[1] = input_
                inputs.remove(input_)
                break

    # find 4
    if 4 not in numbers.keys():
        for input_ in inputs:
            if len(input_) == 4:
                numbers[4] = input_
                inputs.remove(input_)
                break        

    # find 7
    if 7 not in numbers.keys():
        for input_ in inputs:
            if len(input_) == 3:
                numbers[7] = input_
                inputs.remove(input_)
                break
    # find 3
    if 3 not in numbers.keys() and 7 in numbers.keys():
        for input_ in inputs:
            if len(input_) == 5 and all(char in input_ for char in numbers[7]):
                numbers[3] = input_
                inputs.remove(input_)
                break

    # find 8
    if 8 not in numbers.keys():
        for input_ in inputs:
            if len(input_) == 7:
                numbers[8] = input_
                inputs.remove(input_)
                break

    # find 9
    if 9 not in numbers.keys() and 4 in numbers.keys():
        for input_ in inputs:
            if len(input_) == 6 and all(char in input_ for char in numbers[4]):
                numbers[9] = input_
                inputs.remove(input_)
                break

    # find 2
    if 2 not in numbers.keys() and 'f' in letters:
        for input_ in inputs:
            if letters['f'] not in input_:
                numbers[2] = input_
                inputs.remove(input_)
                break

    # find a
    if 'a' not in letters.keys() and 7 in numbers.keys() and 1 in numbers.keys():
        for char in numbers[7]:
            if char not in numbers[1]:
                letters['a'] = char
                break

    # find e
    if 'e' not in letters.keys() and 8 in numbers.keys() and 9 in numbers.keys():
        for char in numbers[8]:
            if char not in numbers[9]:
                letters['e'] = char
                break

    # find 5
    if 5 not in numbers.keys() and 'a' in letters.keys() and 'e' in letters.keys(): # can add a check for dups
        for input_ in inputs:    
            if len(input_) == 5 and (letters['a'] not in numbers) and (letters['e'] not in numbers):
                numbers[5] = input_
                inputs.remove(input_)
                break

    # find c
    if 'c' not in letters.keys() and 1 in numbers.keys() and 'f' in letters.keys():
        for char in numbers[1]:
            if char != letters['f']:
                letters['c'] = char
                break

    # find d
    if 'd' not in letters.keys() and 2 in numbers.keys() and  4 in numbers.keys() and 'c' in letters.keys():
        for char in numbers[4]:
            if char in numbers[2] and char != letters['c']:
                letters['d'] = char
                break

    # find 0
    if 0 not in numbers.keys() and 'd' in letters.keys():
        for input_ in inputs:
            if len(input_) == 6 and (letters['d'] not in input_):
                numbers[0] = input_
                inputs.remove(input_)
                break

    # find 6
    if 6 not in numbers.keys() and len(inputs)==1: # can add a check for dups
        numbers[6] = inputs[0]

                
    assert(len(numbers)==10)

    #########################################
    # decode outputs
    #########################################

    decoded_output = ''
    while len(decoded_output) < 4:
        for output in outputs:
            for number in [0,1,2,3,4,5,6,7,8,9]:
                a = sorted(list(output))
                b = sorted(list(numbers[number]))

                if a == b:
                    decoded_output += str(number)

    return decoded_output


def main():
    """
    main function
    """ 
    data = get_inputs()

    # calc decoded outputs
    decoded_outputs = []
    for i in range(0,len(data)):
        decoded_outputs.append(decode(data[i][0],data[i][1]))

    # sum decoded outputs
    ans = sum([int(i) for i in decoded_outputs])
    print(f'Answer = {ans}')
    return ans


if __name__ == "__main__":
    main()