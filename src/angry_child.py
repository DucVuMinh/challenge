import math
import os


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
    for i in range(k - 2, len(sub)):
        temp = 0
        k_val = []
        begin = i - k + 2
        end = i+1
        for j in range(begin, end):
            k_val.append(sub[j])
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


if __name__ == '__main__':
    print(angryChildren(5, [4504,1520,5857,4094,4157,3902,822,6643,2422,7288,8245,9948,2822,1784,7802,3142,9739,5629,5413,7232]))
