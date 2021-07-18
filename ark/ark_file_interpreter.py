def iptt(inn):
  fie = inn
def ipt():
  fie = input('File path: ')
f = open(fie, 'r')
l = str(f.readlines()).strip("['']")
print(l)
from time import perf_counter
import random as r
from time import sleep
start_time = perf_counter()
end_time = 0
stack = []
i = 0
ipt = ''
print('$~ compiling...')
while i != len(l):
  if l[i] == ':': end_time = perf_counter(); es = end_time - start_time + 2; ll = len(l) + 2; RBS = es / ll; RBS *= 1000; print('\n$~ done'); print('$~ ' + str(RBS)); exit()   
  elif l[i] == '!': i += 1; print(l[i]); i += 1
  elif l[i] == '>': i += 1; stack.append(l[i])   
  elif l[i] == '^': print(str(stack).strip('[]').replace("'", '').replace(', ', ''))
  elif l[i] == '<': ipt = input('$~ input: '); stack.append(ipt)
  elif l[i] == '+': stack.append(int(stack[0]) + int(stack[1]))
  elif l[i] == '-': stack.append(int(stack[0]) - int(stack[1]))
  elif l[i] == '*': stack.append(int(stack[0]) - int(stack[1]))
  elif l[i] == ')': stack.append(int(stack[1]) - int(stack[0]))
  elif l[i] == '/': stack.append(int(stack[0]) / int(stack[1]))
  elif l[i] == '(': stack.append(int(stack[1]) / int(stack[0]))
  elif l[i] == '#': stack = []
  elif l[i] == '$': stack.append(r.randint(int(stack[0]), int(stack[1])))
  elif l[i] == '%': stack.append(r.randint(int(stack[1]), int(stack[0])))
  i += 1
