
def import_data(filename):
    """
    imports list of numbers delimited by new lines
    """

    with open(filename, 'r') as f:
        content = f.read().split('\n')

    return content   

def format_bags(bags_data):
    """
    turns bags into dict objects

    {'name' : 'twergw','contains';[list of other bag names]}
    """

    #can optimize this as a dictionary of dictionaryies

    bag_dicts = {}

    for bag in bags_data:
        bag_name = bag.split(' contain ')[0][:-1]
        bag_contents = []
        bag_cont_data = bag[:-1].split(' contain ')[1].split(', ')

        for data in bag_cont_data:
            if bag_cont_data[0].split(' ')[0] == 'no':
                quantity =0
            else:
                quantity = int(data.split(' ')[0])

                if quantity > 1:
                    sub_bag_name = data.split(str(quantity)+' ')[1][:-1]
                else:
                    sub_bag_name = data.split(str(quantity)+' ')[1]

            for i in range(0,quantity):
                bag_contents.append(sub_bag_name)

        bag_dict = {'name': bag_name, 'contains':bag_contents}
        bag_dicts[bag_name] = bag_dict

    return bag_dicts

def main():
    """
    """
        
    bags_data = import_data('datainput.txt')
  #  bags_data = import_data('testdata.txt')

    bags = format_bags(bags_data)
   # print(bags)


    def reccur_opt(bag_dicts,bag):

        return bag['name'] == 'shiny gold bag' or any(
            reccur_opt(bag_dicts,bags[sub_bag_name]) for sub_bag_name in bag['contains'])            


    def reccur_(bag_dicts,bag):

        if bag['name'] == 'shiny gold bag':
            print(' - shiny_gold_bag found')
            return True
            
     #   print(bag['contains'])
        for sub_bag_name in bag['contains']:
          #  
       #    print(bags[sub_bag_name])
         #   if sub_bag_name == bag['name']:
         #       continue

            #find sub_bag dict info and recurr
            if reccur_(bag_dicts,bags[sub_bag_name]):
                print('found')
                return True

        return False

    shiny_gold_bag_count = 0

    list_of_bag_names = [bag for bag in list(bags.keys())]

    for bag_name in list_of_bag_names:
        bag = bags[bag_name]

        print('\nchecking "{}" [{}/{}]'.format(bag_name,list_of_bag_names.index(bag_name)+1,len(list_of_bag_names)))

        if reccur_opt(bags,bag) == True:
            shiny_gold_bag_count = shiny_gold_bag_count + 1

    print(shiny_gold_bag_count-1) # -1 as you look in gold bag initially

if __name__ == "__main__":
    main()
