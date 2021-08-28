import sys
sys.setrecursionlimit(30000)

def prod(myList):
    """
    """
    result = 1
    for x in myList:
         result = result * x
    return result

def find_entries_that_sum_to(entries,target_sum):
    """
    basic way
    """
    for num_1 in entries:
        remaining_entries = entries[:]
        remaining_entries.remove(num_1)

        for num_2 in remaining_entries:
            remaining_entries2 = remaining_entries[:]
            remaining_entries2.remove(num_2)

            for num_3 in remaining_entries2:

                if sum([num_1,num_2,num_3]) ==target_sum:
                    print(num_1,num_2,num_3,num_1*num_2*num_3)



def find_entries_that_sum_to_recur(entries,N,target_sum):
    """
    searches through a list cecking if any two elements sum to a target number
    entries = list of numbers
    N = number entries to sum
    target_sum = number to sum up to
    """

    def reset(to_sum):
        """
        """

        if len(to_sum) == N:
            for i in range(0,N):
                to_sum[i] = 0

        else:
            for i in range(0,N):
                to_sum.append(0)

    def sum_and_check(to_sum,target_sum):
        """
        """
        if sum(to_sum) == target_sum:
            print('\nnew combination found\nnums {}'.format(to_sum))
            print('sum {}'.format(sum(to_sum)))
            print('mult {}'.format(prod(to_sum)))
            match = [to_sum,sum(to_sum),prod(to_sum)]
            return match
        else:
            return False

    def recurr(entries,to_sum,counter):
        """
        recursive
        """
        counter = counter +1

        assert counter <= N
        
        for num in entries:

          #  print(to_sum,'before change',counter)

            
            to_sum[counter-1] = num

         #   print(to_sum,'after change')
            
            if counter == N and sum(to_sum) == target_sum:

                print(to_sum,counter,sum(to_sum),sum_and_check(to_sum,target_sum))
                
                if sum_and_check(to_sum,target_sum) == False:
                    reset(to_sum)
                    print(to_sum,'after reset')
                    recurr(entries,to_sum,0)


            if counter < N:
                remaining_entries = entries[:]
                remaining_entries.remove(num)
                res = recurr(remaining_entries,to_sum,counter)

                if res:
                    return res

    matches = []
    to_sum  = []
    counter = 0
    reset(to_sum)
    recurr(entries,to_sum,counter)
            
    return matches

#####################################################

def import_data(filename):
    """
    imports list of numbers delimited by new lines
    """

    array = []

    with open(filename, 'r') as f:
        for line in f.readlines():
            array.append(int(line.split('\n')[0]))
    return array

#####################################################

def main():
    
    inputs = import_data('input_data.txt')
    #find_entries_that_sum_to(inputs,2020)
    find_entries_that_sum_to_recur(inputs,3,2020)

if __name__ == "__main__":
    main()
