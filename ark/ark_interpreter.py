from time import perf_counter
import random as r
start_time = 0
end_time = 0
stack = []
i = 0
func = []
sfunc = ''
ipt = ''
mul = []
smul = ''
dbg = False
dbgd = {}
dbgo = input("$~ Debug on/off?\n([on] or [off])")
print('$~ compiling...')
if dbgo != "on" and dbgo != "off":
    dbg = False
if dbgo == "on":
    dbg = True
if dbgo == "off":
    dbg = False
print("$~ show debug data: " + str(dbg))
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
  pop = []
  spop = ''
  l = x
  def innerexec(ip):
    astack = []
    sprint = []
    ai = 0
    while ai != len(ip):
      if ip[ai] == ':': pass
      elif ip[ai] == '@': start_time = perf_counter()
      elif ip[ai] == '!': ai += 1; sprint.append(ip[ai]); ai += 1; dbgd[ai] = str("fn: printed " + ip[ai-1])
      elif ip[ai] == '>': ai += 1; astack.append(ip[ai]); dbgd[ai] = str("fn: pushed " + ip[ai])
      elif ip[ai] == '^': print(str(astack).strip('[]').replace("'", '').replace(', ', ' ')); dbgd[ai] = "fn: printed the stack (" + str(astack) + ")"
      elif ip[ai] == '<': ipt = input('$~ input: '); astack.append(ipt); dbgd[ai] = str("fn: asked for input and got " + ipt)
      elif ip[ai] == '+': astack.append(int(astack[0]) + int(stack[1])); dbgd[ai] = "fn: did maths"
      elif ip[ai] == '-': astack.append(int(astack[0]) - int(astack[1])); dbgd[ai] = "fn: did maths"
      elif ip[ai] == '*': astack.append(int(astack[0]) * int(astack[1])); dbgd[ai] = "fn: did maths"
      elif ip[ai] == '/': astack.append(int(astack[0]) / int(astack[1])); dbgd[ai] = "fn: did maths"
      elif ip[ai] == '#': astack = []; dbgd[ai] = "fn: cleared the stack"
      elif ip[ai] == '~': print(innerexec(sfunc)); dbgd[ai] = "fn: executed the function"
      elif ip[ai] == '[':
          while ip[ai] != ']':
              mul.append(ip[ai])
              i += 1
              smul = str(mul).strip('[').strip(']').replace(', ', '').replace('[', '').replace(']', '').replace("'", '').replace('"', "")
          astack.append(int(smul))
          dbgd[i] = str("pushed " + smul)
      ai += 1
    for x in astack:
      stack.append(x)
    return str(sprint).strip('[]').replace("'", '').replace(', ', '')


  while i != len(l):
    if l[i] == ':':
        end_time = perf_counter()
        es = end_time - start_time + 2
        ll = len(l) + 2
        RBS = es / ll
        RBS *= 1000
        print('\n$~ done')
        print('$~ finished with a regal byte speed of: ' + str(RBS))
        if dbg:
            print("Debug: " + str(dbgd))
            exit()
    elif l[i] == '@': start_time = perf_counter(); dbgd[i] = "started clock and program"
    elif l[i] == '!': i += 1; print(l[i]); i += 1; dbgd[i] = str("printed " + l[i-1])
    elif l[i] == '>': i += 1; stack.append(l[i]); dbgd[i] = str("pushed " + l[i])
    elif l[i] == '^': print(str(stack).strip('[]').replace("'", '').replace(', ', ' ')); dbgd[i] = "printed the stack (" + str(stack) + ")"
    elif l[i] == '<': ipt = input('$~ input: '); stack.append(ipt); dbgd[i] = str("asked for input and got " + ipt)
    elif l[i] == '+': stack.append(int(stack[0]) + int(stack[1])); dbgd[i] = "did maths"
    elif l[i] == '-': stack.append(int(stack[0]) - int(stack[1])); dbgd[i] = "did maths"
    elif l[i] == '*': stack.append(int(stack[0]) * int(stack[1])); dbgd[i] = "did maths"
    elif l[i] == '/': stack.append(int(stack[0]) / int(stack[1])); dbgd[i] = "did maths"
    elif l[i] == '#': stack = []; dbgd[i] = "cleared the stack"
    elif l[i] == '{':
        while l[i] != '}':
            func.append(l[i])
            i += 1
            sfunc = str(func).strip('[').strip(']').replace(', ', '').replace('[', '').replace(']', '').replace("'", '').replace('"', "")
        dbgd[i] = str("declared a function " + sfunc)
    elif l[i] == '~': print(innerexec(sfunc)); dbgd[i] = "executed the function"
    elif l[i] == '[':
        while l[i] != ']':
            mul.append(l[i])
            i += 1
            smul = str(mul).strip('[').strip(']').replace(', ', '').replace('[', '').replace(']', '').replace("'", '').replace('"', '')
        stack.append(int(smul))
        dbgd[i] = str("pushed " + smul)
    i += 1



while True:
  run(input('$~ input: '))
