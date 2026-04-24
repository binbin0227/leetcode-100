nums=[1,2,3,4,5]
n = len(nums)
sum_list=[]
for index,num in enumerate(nums):
    if index == 0:
        sum_list.append(num)
    else:
        sum_list.append(sum_list[index-1]+num)
print(sum_list)