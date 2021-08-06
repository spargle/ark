from time import perf_counter
import random as r
from time import sleep
start_time = 0
end_time = 0
stack = []
i = 0
func = []
sfunc = ''
ipt = ''
mul = []
smul = ''
print('$~ compiling...')

def run(x):
  start_time = 0
  end_time = 0
  stack = []
  i = 0
  func = []
  sfunc = ''
  ipt = ''
  mul = []
  smul = ''
  l = x
  def innerexec(ip):
    astack = []
    sprint = []
    ai = 0
    while ai != len(ip):
      if ip[ai] == ':': ai = ai
      elif ip[ai] == '@': start_time = perf_counter()
      elif ip[ai] == '!': ai += 1; sprint.append(ip[ai]); ai += 1
      elif ip[ai] == '>': ai += 1; astack.append(ip[ai])   
      elif ip[ai] == '^': print(str(astack).strip('[]').replace("'", '').replace(', ', ' '))
      elif ip[ai] == '<': ipt = input('$~ input: '); astack.append(ipt)
      elif ip[ai] == '+': astack.append(int(astack[0]) + int(stack[1]))
      elif ip[ai] == '-': astack.append(int(astack[0]) - int(astack[1]))
      elif ip[ai] == '*': astack.append(int(astack[0]) - int(astack[1]))
      elif ip[ai] == ')': astack.append(int(astack[1]) * int(astack[0]))
      elif ip[ai] == '/': astack.append(int(astack[0]) / int(astack[1]))
      elif ip[ai] == '(': astack.append(int(astack[1]) / int(astack[0]))
      elif ip[ai] == '#': astack = []
      elif ip[ai] == '$': astack.append(r.randint(int(astack[0]), int(astack[1])))
      elif ip[ai] == '%': astack.append(r.randint(int(astack[1]), int(astack[0])))
      elif ip[ai] == '~': print(innerexec(sfunc))
      elif ip[ai] == '[': exec("""while l[i] != ']':\n mul.append(l[i])\n i += 1\n global smul\n smul = str(mul).strip('[').strip(']').replace('\\'', '').replace(', ', '').replace('[', '').replace(']', '')\n"""); stack.append(int(smul))
      ai += 1
      
      return str(sprint).strip('[]').replace("'", '').replace(', ', '')


  while i != len(l):
    if l[i] == ':': end_time = perf_counter(); es = end_time - start_time + 2; ll = len(l) + 2; RBS = es / ll; RBS *= 1000; print('\n$~ done'); print('$~ ' + str(RBS)); exit()
    elif l[i] == '@': start_time = perf_counter()
    elif l[i] == '!': i += 1; print(l[i]); i += 1
    elif l[i] == '>': i += 1; stack.append(l[i])   
    elif l[i] == '^': print(str(stack).strip('[]').replace("'", '').replace(', ', ' '))
    elif l[i] == '<': ipt = input('$~ input: '); stack.append(ipt)
    elif l[i] == '+': stack.append(int(stack[0]) + int(stack[1]))
    elif l[i] == '-': stack.append(int(stack[0]) - int(stack[1]))
    elif l[i] == '*': stack.append(int(stack[0]) - int(stack[1]))
    elif l[i] == ')': stack.append(int(stack[1]) * int(stack[0]))
    elif l[i] == '/': stack.append(int(stack[0]) / int(stack[1]))
    elif l[i] == '(': stack.append(int(stack[1]) / int(stack[0]))
    elif l[i] == '#': stack = []
    elif l[i] == '$': stack.append(r.randint(int(stack[0]), int(stack[1])))
    elif l[i] == '%': stack.append(r.randint(int(stack[1]), int(stack[0])))
    elif l[i] == '{': exec("""\nwhile l[i] != '}':\n  func.append(l[i])\n  i += 1\n  sfunc = str(func).strip('[').strip(']').replace('\\'', '').replace(', ', '').replace('[', '').replace(']', '')\n""")
    elif l[i] == '~': print(innerexec(sfunc))
    elif l[i] == '[': exec("""while l[i] != ']':\n mul.append(l[i])\n i += 1\n global smul\n smul = str(mul).strip('[').strip(']').replace('\\'', '').replace(', ', '').replace('[', '').replace(']', '')\n"""); stack.append(int(smul))
    i += 1
    
def le():
  l = input('$~ input: ')
  run(l)
le()
