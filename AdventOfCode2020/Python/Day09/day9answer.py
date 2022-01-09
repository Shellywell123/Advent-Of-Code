from typing import List
from collections import namedtuple

instruction = namedtuple("instruction",["action", "value"])
instructions = List[instruction]


def import_data(filename):
    """
    imports list of numbers delimited by new lines
    """

    with open(filename, 'r') as f:
        content = f.read().split('\n')

    return content   

def check_num(target_num,ammo_nums):
    """
    """
    for first_num in ammo_nums:
        for second_num in ammo_nums:
         #   print(first_num,second_num)
            if first_num == second_num:
                continue
            if int(first_num) + int(second_num) == int(target_num):
                return True
            
    return False


def main():
    """
    """
        
    data = import_data('inputdata.txt')

    # for d in data:
    #     print(d)


    # prob 1
    preamble = 25

    for i in range(preamble,len(data)):
        target_num = data[i]
        ammo_nums  = data[i-preamble:i]

        
        if check_num(target_num,ammo_nums) == False:
            print(target_num)
            ans1 = target_num
            break

    print('#'*10)

    # prob 2
    for preamble in range(3,100):
        for i in range(0,len(data)):
            target_num = int(ans1)
            ammo_nums = data[i-preamble:i]

            for i in range(0,len(ammo_nums)):
                ammo_nums[i] = int(ammo_nums[i])

            if sum(ammo_nums) == target_num:
                print(min(ammo_nums)+max(ammo_nums))
                break







        
 

if __name__ == "__main__":
    main()