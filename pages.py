#Hopefully I can make a window you can write on
from notepad import *

import curses,os

class papyrus:
    
    def __init__(self):
        
        self.tab = 
    
    def scribe(self):

        wind = curses.initscr()
        wind.cbreak(False)
        wind.echo(False)
        wind.keypad(True)
        wind.getmouse()

    def save(self):

        if wind:
            wind.putwin(os.path('edition','x+',encoding = 'utf-8'))#its not right but you get the idea, ima add the function to notepad soon
#FoodTime
