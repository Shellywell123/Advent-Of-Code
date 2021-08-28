def import_data(filename):
    """
    imports list of numbers delimited by new lines
    """

    with open(filename, 'r') as f:
        content = f.read()

    return content

def main():
    """
    """
    data = import_data('datainput.txt')
    #print(data)

    passports = data.split('\n\n')
    import ast

    # reformat data into dicts
    for i in range(0,len(passports)):
        passport = "{'"+passports[i].replace("\n"," ").replace(" ","', '").replace(":","':'")+"'}"
        passports[i] = ast.literal_eval(passport)

    #print(passports)

    valid = 0

    checks = ['byr','iyr','eyr','hgt','hcl','ecl','pid'] #cid

    for passport in passports:

        keys = list(passport.keys())
     #   print(keys)


        #check if passport has all fields
        if not all(x in keys for x in checks) :
            continue

        if not 1920 <= int(passport['byr']) <= 2002:
            continue

        if not 2010 <= int(passport['iyr']) <= 2020:
            continue

        if not 2020 <= int(passport['eyr']) <= 2030:
            continue

     #   print(passport['hcl'])

        if (len( passport['hcl']) !=7) and passport['hcl'][0] != '#':
            for val in list(passport['hcl'][0:]):
                if val in ['g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z']:
                    continue
      #      print(passport['hcl'])
            continue

        if len(passport['hgt']) <= 2:
            continue

        if not (passport['hgt'][-2:] == 'cm') and (150 <= int(passport['hgt'][:-2]) <= 193):
            continue

        if not (passport['hgt'][-2:] == 'in') and (59 <= int(passport['hgt'][:-2]) <= 76):
            continue

      #  print(passport['hgt'])

        if not 2020 <= int(passport['eyr']) <= 2030:
            continue

        if passport['ecl'] not in ['amb', 'blu', 'brn', 'gry', 'grn', 'hzl', 'oth']:
            continue

        if (len(passport['pid']) != 9):
            continue

        valid = valid+1
        print(passport)

    print(valid,len(passports))

if __name__ == "__main__":
    main()
