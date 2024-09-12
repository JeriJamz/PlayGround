from printr import * 
from  file import *
from notepad import *
#****************COMPILE(...or run) AT YOUR OWN RISK***************************
#I just cant compile it but it should be good. Ima dawg just check my spelling
import time,sys,pyautogui as pag, curses

#refresh()<-update screen after input, noutrefresh()<-change background(just text),doupdate()<- update display
#Just thinking if I have the user input in all the other files do I need to have user input added here. 
#Ima just make this a screen
class KernBootWindow:

    def __init__(self,boot,put, usr,coline):
        self.boot = termi
        self.put = termin.getkey()
        self.coline = pag.size()#maybe not just trying to autosize the window, think this just for mouses
        self.usr = curses.newpad(coline.)#I think this can make a box inside a box so I can write in the box. M Night Shyamalan
         #Would be proud   
    def Boots(self):#see if I can make this the main window wit a boot up and file director
         #****** FOOOD TIME... plus I gott study*******
        termin = curses.initscr()

        termin.def_shell_mode() #yes all that commiting for one line
        termin.print("High-Level Kernal.OS")#This might not work
        termin.noecho()
        termin.keypad(True)
        termin.clear()#The Boot Uplol wheres The ASM windows?
        termin.addstr()#I should start getting user input... Let me make a menu

        menu()
    
    def oltra(self,End):
        
        self.end = end
    
        termin.nobreak()
        termin.keypad(False)
        termin.clear()
        termin.echo()
        
    def Term(self,exit):
        self.exit = exit
    
        termin.endwin()

    def Boxput(self, userBox):
        termin.clear()
        termin.echo()
        termin.keypad(True)
        termin.getch()
        termin.clear()
        if termin.getkey()// its a command prompt line but I got hungry
        
Boots()
