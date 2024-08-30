from Printr import * 
#****************COMPILE(...or run) AT YOUR OWN RISK***************************
#I just cant compile it but it should be good. Ima dawg just check my spelling
import time,sys,pyautogui as pag, curses

#refresh()<-update screen after input, noutrefresh()<-change background(just text),doupdate()<- update display

#tweak'd need a class

class KernBootWindow:

    def __init__(name,boot):
        self. name = name
        self.boot = termin
        name = The Screen
    
    def Boots(self):#see if I can make this the main window wit a boot up and file director
        
        termin = curses.initscr()
    
        termin.print("High-Level Kernal.OS")#This might not work
        termin.noecho()
        termin.keypad(True)
        termin.clear()#The Boot Uplol wheres The ASM windows?
        termin.addstr()#I should start getting user input... Let me make a menu
    
    def oltra(self,End):
        
        self.end = end
    
        termin.nobreak()
        termin.keypad(False)
        termin.clear()
        termin.echo()
        
    def Term(self,exit):
        self.exit = exit
    
        termin.endwin()

boot()
