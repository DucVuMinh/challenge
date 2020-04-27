import math
import os


# Complete the angryChildren function below.
# Complete the angryChildren function below.
def angryChildren(k, packets):
    # sorting
    packets.sort()
    # cal sub
    sub = []
    for i in range(1, len(packets)):
        sub.append(packets[i] - packets[i - 1])
    # cal unfairness sum temp
    res_temp = []
    # n-1 items
    for i in range(k - 2, len(sub)):
        temp = 0
        k_val = []
        begin = i - k + 2
        end = i+1
        for j in range(begin, end):
            k_val.append(sub[j])
         # k items
        for j in range(k-1):
            mul = (j+1) * (k-1-j)
            temp += mul* k_val[j]
        res_temp.append(temp)
    # find the best
    res = res_temp[0]
    for i in res_temp:
        if res > i:
            res = i
    return res
# Bad resuls = k*n + nlogn

if __name__ == '__main__':
    print(angryChildren(3, [10,100,300,200,1000,20,30]))
