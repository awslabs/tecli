import itertools
import lorem
import random

# oList=["description","enabled","mode","new-name","organization-token","team-token","user-token","config","log","log-file-path","organization","profile","verbosity"]

# original list
oList=[
"--description=string",
"--enabled=bool",
"--mode=string",
"--new-name=string",
"--organization-token=string",
"--team-token=string",
"--user-token=string",
"--config=string",
"--log=string",
"--log-file-path=string",
"--organization=string",
"--profile=string",
"--verbosity=string",
]

# new list 
nList=['']

for i in range(len(oList)):
    item = oList[i]
    s=item.split("=")
    print(s)
    if s[0] == "--description":
        l = '{}="{}"'.format(s[0], lorem.sentence())
        nList.append(l)
    if s[0] == "--enabled":
        bit = random.getrandbits(1)
        b = bool(bit)
        l = '{}={}'.format(s[0], b)
        nList.append(l)
    if s[0] == "--mode":
        bit = random.getrandbits(1)
        b = bool(bit)
        if b == True:
            l = '{}={}'.format(s[0], "interactive")
        else:
            l = '{}={}'.format(s[0], "non-interactive")
        nList.append(l)
print(nList)

# for i in range(len(oList)):
#     # item = oList[i]
#     combinations = itertools.combinations(oList, i)
#     for result in combinations:
#         print(result)