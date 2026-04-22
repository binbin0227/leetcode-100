p="aab123"
p_dict = {}
for char in p:
    if char not in p_dict:
        p_dict[char]=1
    else:
        p_dict[char]+=1
print(p_dict)