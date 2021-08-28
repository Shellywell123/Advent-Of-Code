def import_data(filename):
    """
    imports list of numbers delimited by new lines
    """

    array = []

    with open(filename, 'r') as f:
        for line in f.readlines():
            policy_code,password = line.split(': ')

            #remove newline delimter
            password = password[:-1]

            entry = {'policy_code' : policy_code, 'password' : password}
            array.append(entry)
    return array

def countX(lst, x):
    count = 0
    for ele in lst:
        if (ele == x):
            count = count + 1
    return count

def count_instances(data):
    """
    """
    valid_count = 0

    for entry in data:
        lims,letter = entry['policy_code'].split(' ')
        low_lim,high_lim = lims.split('-')

        # get count
        count = countX(list(entry['password']),letter)

        if int(low_lim) <= count <= int(high_lim):
            valid_count = valid_count + 1

    print(valid_count)

def check_index(data):
    """
    """
    valid_count = 0

    for entry in data:
        indexs,letter = entry['policy_code'].split(' ')
        low_ind,high_ind = indexs.split('-')

        low_ind = int(low_ind)-1
        high_ind = int(high_ind)-1

        if (entry['password'][low_ind] == letter) and (entry['password'][high_ind] != letter):
            valid_count = valid_count+1

        if (entry['password'][low_ind] != letter) and (entry['password'][high_ind] == letter):
            valid_count = valid_count+1

    print(valid_count)



def main():
    """
    """
    data = import_data('datainput.txt')
    #print(data)
    check_index(data)

    

if __name__ == "__main__":
    main()
