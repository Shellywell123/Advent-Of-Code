def import_data(filename):
    """
    imports list of numbers delimited by new lines
    """

    with open(filename, 'r') as f:
        content = f.read().split('\n\n')

    return content

def part1():
        """
        """
        groups = import_data('datainput.txt')

        awnser_count = 0

     #   print(len(groups),groups)

        for group in groups:
            checks = ['a','b','c','d','e','f','g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z']
                        
            group_entries = group.split('\n')
       #     print('----',len(checks))

         #   print(len(group_entries))
            for entry in group_entries:
            #    print(' ---',len(checks))
             #   print(entry)
                for awnser in entry:
                    if awnser in checks:
                    #    print('  --',awnser)
                        checks.remove(awnser)
                      #  print(len(checks))
                        awnser_count = awnser_count + 1

        print(awnser_count)

def part2():
        """
        """

        groups = import_data('datainput.txt')
        
        awnser_count_2 = []

        # groups = ['abc','a\nb\nc','ab\nac','a\na\na\na\na','b']
        # groups = ['mgbwxcvkrl\ntzoauif\nqpjntdoysueh']

        for group in groups:  
            group_entries = group.split('\n')
            print('\n')

         #   print(len(group_entries))

           # checks = list(group_entries[0])
            checks = list(min(group_entries, key=len)) #this needs to not overwrtite itself
                    

            for entry in group_entries:

               # print('----',len(checks),checks)

             #   print('\n ---',entry,len(checks),checks)

                def recur_(checks,entry):
                    """
                    """
                    for check in checks:
                #            print(checks.index(check),len(checks))
                         #   print(check)

                        if check not in list(entry):
                            print(check,' not in ', entry)
                            
                            checks.remove(check)
                            recur_(checks,entry)

                recur_(checks,entry)

                 #      print('  --',entry,len(checks))
                          #  print(len(checks))
               # print('checks done next entry')
            #    print(' ---',entry,len(checks),checks)

            awnser_count_2.append(len(checks))
            print('----',len(checks),checks)

        print(sum(awnser_count_2))
        print(len(groups),len(awnser_count_2))

def main():
    """
    """

    part1()
    part2()

if __name__ == "__main__":
    main()
