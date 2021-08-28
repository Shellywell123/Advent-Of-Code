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

def format_data(data) -> instructions:
    """
    """
    instruction_list = []
    executed = []

    for d in data:
        inst,val = d.split(' ')
        x =instruction(inst,int(val))
        executed.append(False)
        instruction_list.append(x)

    return instruction_list,executed

def read_inst(insts,inst_ind,acc,exes):
    """
    """
    insts_ = insts[:]
    inst = insts_[inst_ind]
    exes[inst_ind] = True

    if inst.action == 'jmp':
        return acc,inst_ind+inst.value,exes,insts_

    if inst.action == 'nop':
        return acc,inst_ind+1,exes,insts_

    if inst.action == 'acc':
        return acc+inst.value,inst_ind+1,exes,insts_

def fix_inst(insts,inst_ind):
    """
    """

    inst = insts[inst_ind]

    if inst.action == 'jmp':
        insts[inst_ind] = instruction('nop',inst.value)


    if inst.action == 'nop':
        insts[inst_ind] = instruction('jmp',inst.value)

    return insts


def main():
    """
    """
        
    insts,exes = format_data(import_data('datainput.txt'))

    acc_init = 0

    def goto_(acc,insts,inst_ind,exes):
        """
        """
        acc,inst_ind,exes,insts_ = read_inst(insts,inst_ind,acc,exes)

        print(acc,insts[inst_ind],exes[inst_ind])

        if exes[inst_ind] == False:
            acc,inst_ind = goto_(acc,insts,inst_ind,exes)

        return acc, inst_ind

    #part 1
    acc, inst_ind = goto_(acc_init,insts,0,exes)
    print(acc)

    print('#'*10)
    #part_2


    for inst_ind in range(0,len(insts)):
        insts_,exes = format_data(import_data('datainput.txt'))

        inst = insts_[inst_ind]
        if inst.action == ('jmp' or 'nop'):

            if inst.action == 'jmp':
                insts_[inst_ind] = instruction('nop',inst.value)

            if inst.action == 'nop':
                insts_[inst_ind] = instruction('jmp',inst.value)

            acc, inst_ind = goto_(0,insts_,0,exes)
            if acc != 1563:
                print(acc)
    # if this returns index out of range the first number on the last line printed is the ans                   


if __name__ == "__main__":
    main()