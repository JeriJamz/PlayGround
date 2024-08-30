#see now I might have to import some files
import time,sys


def qprint(z):
    for s in z:
    sys.stdout.write(s)
    sys.stdout.flush()
    time.sleep(.05)

def rpgprint(z):
    for s in z:
        sys.stdout.write(s)
        sys.stdout.flush()
        time.sleep(.573)
