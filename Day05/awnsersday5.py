import numpy as np

def import_data(filename):
    """
    imports list of numbers delimited by new lines
    """

    with open(filename, 'r') as f:
        content = f.read().split('\n')

    return content

def upper_half(low_lim,high_lim):
    """
    """
    if low_lim+1 == high_lim:
      #  print('ca')
        return high_lim
    
    new_low_lim = int(np.ceil(np.mean([low_lim,high_lim])))

    return new_low_lim,high_lim


def lower_half(low_lim,high_lim):
    """
    """
    if low_lim+1 ==high_lim:
      #  print('ca')
        return low_lim
    
    new_high_lim = int(np.floor(np.mean([low_lim,high_lim])))

    return low_lim,new_high_lim

def main():
    """
    """
    boarding_passes = import_data('datainput.txt')
    print(len(boarding_passes))

    start_row_low_lim  = 0
    start_row_high_lim = 127

    start_col_low_lim = 0
    start_col_high_lim = 7

    taken_seats = []

    for boarding_pass in boarding_passes:
      #  print(boarding_pass)
   #     boarding_pass = 'FBFBBFFRLR'

        low_lim = start_row_low_lim
        high_lim = start_row_high_lim

        # first seven characters to get row
        for i in range(0,7):
            character = boarding_pass[i]

         #   print(boarding_pass[i],i,low_lim,high_lim,'->',)


            if i == 6:

                if character == 'F':
                    row_number = lower_half(low_lim,high_lim)

                if character == 'B':
                    row_number = upper_half(low_lim,high_lim)
                break

            if character == 'F':
                low_lim,high_lim = lower_half(low_lim,high_lim)

            if character == 'B':
                low_lim,high_lim = upper_half(low_lim,high_lim)

         #   print(low_lim,high_lim)


      #  print('gh',boarding_pass[6],row_number)

        # last 3 characters to get seat

        low_lim = start_col_low_lim
        high_lim = start_col_high_lim

        for i in range(6,10):
            character = boarding_pass[i]

           # print(boarding_pass[i],i,low_lim,high_lim,'->',)

            if i == 9:

                if character == 'L':
                    col_number =  lower_half(low_lim,high_lim)

                if character == 'R':
                    col_number = upper_half(low_lim,high_lim)

                break

            if character == 'L':
                low_lim,high_lim =  lower_half(low_lim,high_lim)

            if character == 'R':
                low_lim,high_lim = upper_half(low_lim,high_lim)

          #  print(low_lim,high_lim)

        assert type(row_number) == int
        assert type(col_number) == int

        print(row_number,'-',col_number)

        seati_id = row_number*8 + col_number
        taken_seats.append(seati_id)

    print(taken_seats, len(taken_seats))

    print(min(taken_seats),max(taken_seats))

    for i in range(min(taken_seats),max(taken_seats)):
        if i not in taken_seats:
            print(i)


if __name__ == "__main__":
    main()
