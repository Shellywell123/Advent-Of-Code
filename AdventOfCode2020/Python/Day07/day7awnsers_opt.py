
def import_data(filename):
    """
    imports list of numbers delimited by new lines
    """

    with open(filename, 'r') as f:
        content = f.read().split('\n')

    return content   

from typing import List
from collections import namedtuple

Bags = namedtuple("Bags", ["name", "contains"])
BagDict = List[Bags]


def format_bags(bags_data) -> BagDict:
    """
    """

    bag_dicts = []

    for bag in bags_data:
        bag_name = bag.split(' contain ')[0][:-1]
        bag_contents = []
        bag_cont_data = bag[:-1].split(' contain ')[1].split(', ')

        if bag_cont_data[0].split(' ')[0] == 'no':
            continue

        for data in bag_cont_data:
            quantity = int(data.split(' ')[0])

            if quantity > 1:
                sub_bag_name = data.split(str(quantity)+' ')[1][:-1]
            else:
                sub_bag_name = data.split(str(quantity)+' ')[1]

            for i in range(0,quantity):
                bag_contents.append(sub_bag_name)

        # HERE
        bag = Bags(bag_name, bag_contents)
        bag_dicts.append(bag)

    return bag_dicts

def main():
    """
    """
        
    bags_data = import_data('datainput.txt')
   # bags_data = import_data('testdata.txt')

    bags = format_bags(bags_data)

    def reccur_(bag: Bags, bag_dicts: BagDict) -> bool:
        return bag.name == "shiny gold bag" or any(
            reccur_(b, bag_dicts) for b in bag_dicts if b.name in bag.contains
        )

        
    def reccur_2(count,bag: Bags, bag_dicts: BagDict):
        count = count + len(bag.contains)
        for sub_bag_name in bag.contains:
            for b in bag_dicts:
                if b.name == sub_bag_name:
                    count = reccur_2(count,b, bag_dicts)
        return count

    count = 0

    #part 1
    print(sum([reccur_(bag,bags)for bag in bags])-1)

    for bag in bags:
        if bag.name == 'shiny gold bag':
            break

    # part 2
    print(bag.name)
    print(reccur_2(count,bag,bags))

if __name__ == "__main__":
    main()
