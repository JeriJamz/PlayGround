#see now I might have to import some files
import time,sys,curses


def qprint(z):
    try:
        for s in z:
        sys.stdout.write(s)
        sys.stdout.flush()
        time.sleep(.05)
    except:
        curses.cbreak("Space")
        print(z)#I think this work. I refuse to import keyboard


def rpgprint(z):
    try:    
        for s in z:
            sys.stdout.write(s)
            sys.stdout.flush()
            time.sleep(.573)
        except:
            curses.cbreak("Space")
            print(z)
