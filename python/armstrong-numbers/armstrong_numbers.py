def is_armstrong_number(number):
    string = str(number)
    return number == sum([int(i)**len(string) for i in string])
