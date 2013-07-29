from math import sqrt
import sys

#Recursive function to check if the sample is a double square. if it is print the number of ways it is if not print 0
def calcDblSq(i):
    result = 0
    for x in range(int(sqrt(int(i)))+1):
        y = sqrt(int(i) - x*x)
        if y == int(y):
            if x*x == y:
                result += 2
            else:
                result +=1
    print str(result/2)


#read the input data samples
inputData = open(sys.argv[1], 'r')
lineNum = 0
rows = 0
samples = inputData.readlines()
samples.pop(0)
for i in samples:
        calcDblSq(i)

inputData.close()