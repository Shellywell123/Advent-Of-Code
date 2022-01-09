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


def format_data(adapters):
    """
    """

    for i in range(0,len(adapters)):
        adapters[i] = int(adapters[i])

    adapters.sort(reverse=True)
    return adapters

def count_diffs(chain):
    """
    """

    diff_1 = []
    diff_2 = []
    diff_3 = []

    for i in range(0,len(chain)-1):
        current_adapter = chain[i]
        next_adapter = chain[i+1]

        diff = next_adapter - current_adapter

        if diff == 1:
            diff_1.append(current_adapter)

        if diff == 2:
            diff_2.append(current_adapter)

        if diff == 3:
            diff_3.append(current_adapter)

    return len(diff_1)+1,len(diff_2),len(diff_3)+1


def q1():
    """
    """
    adapters = format_data(import_data('testdata.txt'))
    adapters = format_data(import_data('inputdata.txt'))

 #   print(adapters)

    def recur_(chain,adapters_choices):
        """
        """
        for check in [1,2,3]:       
            for adapter_chain_link in adapter_choices:

                diff = (int(adapter_chain_link) - int(chain[-1]))

            
                if diff == check:
               #     print('{} current num {}, next num {}'.format(len(chain),chain[-1],adapter_chain_link))
                    adapters_choices.remove(adapter_chain_link)
                    chain.append(adapter_chain_link)
                    recur_(chain,adapters_choices)


        return chain,adapters_choices

    chain = [min(adapters)]
    adapter_choices = adapters[:-1]
    chain,adapter_choices = recur_(chain,adapter_choices)

    assert len(chain) == len(adapters)

   # print(chain)
        # make function to count jumps
    d1,d2,d3 = count_diffs(chain)
    print(d1*d3)

def q2():
    """
    """
    adapters = format_data(import_data('testdata.txt'))
    adapters = format_data(import_data('inputdata.txt'))

    import random 

    def recur_(chain,adapters_choices):
        """
        """
        checks =  [1,2,3]

        random.shuffle(checks)
     #   print(checks)
        for check in checks:       
            for adapter_chain_link in adapter_choices:

                diff = (int(adapter_chain_link) - int(chain[-1]))

            
                if diff == check:
                 #   print('{} current num {}, next num {}'.format(len(chain),chain[-1],adapter_chain_link))
                    adapters_choices.remove(adapter_chain_link)
                    chain.append(adapter_chain_link)
                    recur_(chain,adapters_choices)

        return chain,adapters_choices


    # remeber rating is last num +3

    print('starting q2')

    rating = max(adapters) + 3

    chains = []
    if 0 not in adapters:
        adapters.append(0)
    adapters.append(rating)
    adapters.sort(reverse=True)
    
    print(adapters)
    

    for i in range(0,1000000000000):
        if i%10000 == 0:
            print('{} checked, {} found.'.format(i,len(chains)))
        chain = [min(adapters)]
        adapter_choices = adapters[:-1]
        chain,adapter_choices = recur_(chain,adapter_choices)
       # print(chain[-1])
        if chain[-1] == rating and chain not in chains:
            chains.append(chain)

    print(len(chains))
    print(chains)


def q2_():
    """
    """

    adapters = format_data(import_data('inputdata.txt'))

    import random
    rating = max(adapters) + 3

    if 0 not in adapters:
        adapters.append(0)
    adapters.append(rating)
    adapters.sort()

    main_chain = adapters[:]

   # print(adapters)

    def check_chain(chain):
        """
        """
        for i in range(0,len(chain)-1):
            cur = chain[i]
            next_ = chain[i+1]

            if not 1<=next_-cur<=3:
                return False
        return True

    chains = []

    # pick a number of adapters to remove from list
    for num_rems in range(0,len(main_chain)-1):
        print(' - combs found {}'.format(len(chains)))
        print('testing {} removals'.format(num_rems))
        

        # pick locations to remove from (complex)
        for j in range(0,10000000):
            chain=main_chain[:]
            # loop through number of adapters to remove
            for i in range(0,num_rems):

                chain.remove(chain[int(random.random()*len(chain))])

            if check_chain(chain) == True and chain not in chains:
                chains.append(chain)

    print(len(chains))






def main():
    """
    """
    q1()
    q2_()
        
if __name__ == "__main__":
    main()