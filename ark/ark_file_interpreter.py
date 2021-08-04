import ark_interpreter as arkin
def iptt(inn):
  f = open(inn, 'r')
  l = str(f.readlines()).strip("['']")
  arkin.run(l)
