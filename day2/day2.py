import re 

sum_power = 0
sum_id = 0

max = {"red":12 , "green":13, "blue":14}
pattern = re.compile(r'(?P<quantity>\d+)\s+(?P<color>\w+)')
with open('input.txt', 'r') as f:
  id = 1
  for line in f.readlines():
    possible = True
    power = {"red":0 , "green":0, "blue":0}
    matches = pattern.findall(line)

    for match in matches:
      power[match[1]] = int(match[0]) if int(match[0]) > power[match[1]] else power[match[1]]
      possible = False if power[match[1]] > max[match[1]] else possible

    sum_power += power["red"] * power["green"] * power["blue"] 
    if possible:
      sum_id += id
    id+=1

print(sum_id)
print(sum_power)
