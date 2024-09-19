#Hopefully I can make a window you can write on
from notepad import *

import curses,os

class papyrus:
    
    def __init__(self,scribe):
        
#        self.tab =  dow
        
    
    def scribe(self):

        #andd stop on notepad gone be right here. I need a setting file and way to save stuff
        wind = curses.initscr()
        wind.border(5,5,25,15,0,0,0,0) #The boarders of the notepad app. Still import most of the funcs also the menu?!? 
        wind.cbreak(False)
        wind.echo(False)
        wind.keypad(True)
        wind.getmouse()

    def save(self):
        for wind ++1:#this this function made my head hurt
            lil_win = curses.newwin(50,150)
            lil_win.addstr(25,55,"Save <Enter> Yes Or No\n\n",
            "<Enter>"))#think I got it down I'm reading docs
            response = input()
            if response.islower() == "Yes":
                wind.putwin(os.path('edition','x+',encoding = 'utf-8'))#I think this is closer but I need to act like its local and not ontop of an O.S
            else:
               curses.endwin(lil_win())#Save Func          
#FoodTime
