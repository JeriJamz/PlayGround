from printr import * 
from  file import *
from notepad import *
#****************COMPILE(...or run) AT YOUR OWN RISK***************************
#I just cant compile it but it should be good. Ima dawg just check my spelling
import time,sys,pyautogui as pag, curses

#refresh()<-update screen after input, noutrefresh()<-change background(just text),doupdate()<- update display

class KernBootWindow:

    def __init__(self,boot,put, usr):
        self.boot = termi
        self.put = termin.getkey()
        self.usr = curses.newwin(2560,20)#trying to make a window but im hungry
    
    def Boots(self):#see if I can make this the main window wit a boot up and file director
         
        termin = curses.initscr()

        termin.def_shell_mode() #yes all that commiting for one line
        termin.print("High-Level Kernal.OS")#This might not work
        termin.noecho()
        termin.keypad(True)
        termin.clear()#The Boot Uplol wheres The ASM windows?
        termin.addstr()#I should start getting user input... Let me make a menu

        menu()

        act = self.put
        act()#this whooping on me I gotta go read
        
    
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
        
boot()
