import sys

def main():

    inputFile = ""
    if(len(sys.argv)) <= 1:
        print("Too few command line arguements")
        sys.exit(1)
    if(len(sys.argv)) > 2:
        print("Too many command line arguements")
        sys.exit(1)
    while(inputFile == ""):
        inputFile = sys.argv[1]
        if (".py" not in inputFile):
            print("Not a python File")
            sys.exit(1)
        else:
            print("Opening File:")
            break

    with open (inputFile, "r") as file:
        contents = file.readlines()
        numlines = 0
        for i in range(len(contents)):
            if(contents[i][0] == "#"):
                continue
            if(contents[i].isspace() == True):
                continue
            #For indented comments
            indented = contents[i].strip()
            if(indented[0] == "#"):
                continue
            else:
                numlines += 1

        print(numlines)

main()
