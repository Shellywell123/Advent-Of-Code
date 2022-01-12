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


def days_til_next_creation(array_of_latern_timers):
    """
    will return minimum nonzero number in list
    """

    for i in range(1,9):
        if array_of_latern_timers[i] > 0:
            return i


def breed(lanterns,n):
    """
    calculate number of lantern fish over 
    set period of days n
    """

    lantern_timer_numbers = [0,1,2,3,4,5,6,7,8,9]
    array_of_latern_timers = [0,0,0,0,0,0,0,0,0,0]

    # format lanterns into array
    for num in lantern_timer_numbers:
        array_of_latern_timers[num] = lanterns.count(num)

    day_no = -1
    dtnc = 0

    # jump through days until past limit
    while day_no <= n-1:

        # return once while loop ends
        prev_len = sum(array_of_latern_timers[0:8])

        # calcualte dtnc
        dtnc = days_til_next_creation(array_of_latern_timers)

        # subtract dtnc from every lantern counter
        for day in range(0,dtnc):

            to_be_moved = array_of_latern_timers[0]
            del array_of_latern_timers[0]
            array_of_latern_timers.append(to_be_moved)
            
        # reset counters and add new ones
        num_to_birth = array_of_latern_timers[0]

        # reset mother cycle
        array_of_latern_timers[0] = 0
        array_of_latern_timers[7] += num_to_birth

        # start baby cycle
        array_of_latern_timers[9] += num_to_birth

        # iterate days
        day_no+=dtnc

    ret = 0
    if day_no == n:
        ret = sum(array_of_latern_timers[0:8])
    if day_no > n:
        ret = prev_len
    if day_no < n:
        ret = sum(array_of_latern_timers)

    return ret


def main():
    """
    main function
    """ 

    inputs = get_inputs()

    num_of_days = 256

    ans = breed(inputs,num_of_days)

    print(f'Answer = {ans}')
    return ans


if __name__ == "__main__":
    main()  