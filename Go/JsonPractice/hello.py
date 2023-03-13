import json

with open('C:/Users/edelk/Code/Go/StartPractis/txt.json', 'r') as f:
    data = json.load(f)

total_global_id = 0

for element in data:
    total_global_id += element['global_id']

print(total_global_id)
