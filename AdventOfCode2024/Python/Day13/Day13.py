import numpy as np

class machine:
    def __init__(self):
        self.ax = 0
        self.ay = 0
        self.bx = 0
        self.by = 0
        self.px = 0
        self.py = 0

with open("inputs.txt") as f:
        data =((f.read()).replace(" ","")).split("\n")

machines = []
newMachine = machine()
for line in data:
    if line == "":
        machines.append(newMachine)
        newMachine = machine()
        continue

    if "ButtonA" in line:
        newMachine.ax, newMachine.ay = (int(line.split("+")[1].split(",")[0]), int(line.split("+")[-1]))
    
    if "ButtonB" in line:
        newMachine.bx, newMachine.by = (int(line.split("+")[1].split(",")[0]), int(line.split("+")[-1]))
    
    if "Prize" in line:
        newMachine.px, newMachine.py = (int(line.split("=")[1].split(",")[0]), int(line.split("=")[-1]))
    
machines.append(newMachine)

def Solve(constant):

    ans = 0
    for m in machines:

        px = m.px + constant
        py = m.py + constant

        a = np.array([[m.ax, m.bx], [m.ay, m.by]])
        b = np.array([px, py])
        s = np.linalg.solve(a, b)
        aPresses, bPresses = np.round(s[0]), np.round(s[1])

        if aPresses < 0 or bPresses < 0:
            continue

        if ((aPresses*m.ax)+(bPresses*m.bx) != px) or ((aPresses*m.ay)+(bPresses*m.by) != py):
            continue
        
        ans += 3*aPresses + bPresses

    return ans

print("Advent-Of-Code 2024 - Day13\n")
print(f'Inputs: Answer to Part 1 = {Solve(0)}')
print(f'Inputs: Answer to Part 2 = {Solve(10000000000000)}')
