**INCOMPLETE DRAFT OF RECOVERED WIKI PAGE**

# JssimUserHelp - VisualChips


	

	
	


## JssimUserHelp


	

		


#### From VisualChips


		

		

		

Welcome to the JSSim javascript simulator which powers the Visual6502 and Visual6800.


Please have a look around this wiki for 
[more information](index.php?title=Main_Page) about our reverse engineering of various chips and the 
[things we've found out](index.php?title=6502Observations) about them.


For a quick selection of examples of 6502 behaviour and layout, have a look at the links in the 
[URL interface section](#url-interface) on this page. Sorry, we don't yet have specific help on 6800 features.


The visual6502 simulator has two entry pages:


-  The 
[simple mode](http://visual6502.org/JSSim/), also known as kiosk mode, which is the default

-  The 
[advanced mode](http://visual6502.org/JSSim/expert.html)
and this page starts with the basics and works up. There's also


-  The 
[visual6800](http://visual6502.org/JSSim/expert-6800.html) (advanced mode only)

For help on reading the layout, interpreting transistor circuits, and more about digital design, please see 
[the Visual Circuit Tutorial](http://wiki.nesdev.com/w/index.php/Visual_circuit_tutorial) over at NESdev.



### Contents


- [1 Help for simple mode](#help-for-simple-mode)
- [1.1 Graphics help (basic)](#graphics-help-.28basic.29)
- [1.2 Running the program](#running-the-program)
- [1.3 Modifying the program](#modifying-the-program)
 [2 Help for advanced mode](#help-for-advanced-mode)
- [2.1 Graphics help (advanced)](#graphics-help-.28advanced.29)
- [2.2 Running the program](#running-the-program-2)
- [2.3 Interacting with the program](#interacting-with-the-program)
- [2.4 Modifying the program](#modifying-the-program-2)
- [2.5 Tracing machine state](#tracing-machine-state)
- [2.6 Busses and signals of interest](#busses-and-signals-of-interest)
- [2.7 URL interface](#url-interface)
 [3 See also](#see-also)

###  
 Help for simple mode 


In simple mode, you see the chip graphics on the left, the control buttons and chip status at top right, and the memory table below that. There's a link to the advanced page, and the overall layout is fixed: there are no draggable bars between the sections.



#####  
 Graphics help (basic) 


You can pan and zoom the chip graphics using


-  '>' on the keyboard to zoom in

-  '<' on the keyboard to zoom out

-  click and drag to pan

Click in the graphics area to highlight any shape on the chip: all the connected shapes will be highlighted and the name of the node, if any, will be displayed in the chip status area at top right.


For example, if you click on the square shape at top left of the chip, you'll see text like


```
node: 1297 nmi

```

which tells you that this is the NMI pad - in a real chip, it would be connected to the NMI pin of the package with a gold wire.


The node number is useful only as a unique reference number. If you're interested in the workings of the simulator you'll read the source files and see these numbers used to label all the polygons and transistors which are electrically connected and which therefore are at the same voltage - and therefore represent the same logical signal.



#####  
 Running the program 


Towards the top right you see a set of buttons:


-  run (or stop) - start the simulation, run for as long as you like, then stop it.

-  reset

-  back

-  forward

As the simulation runs you can see the yellow box in the memory area (bottom right) indicating which memory location is being read or written.  You may also see the contents of memory changing: perhaps the location just to the right of {{{0040:}}} will count up.



#####  
 Modifying the program 


You can't presently modify the program in the simple page: you need the Advanced page for that.



###  
 Help for advanced mode 


In advanced mode, there's an additional area at bottom right which can tabulate the state of the machine, and any signals of interest, phase by phase or instruction by instruction.  There's also a console for programs which perform I/O: it's possible to interact with a BASIC interpreter for example.


The layout in advanced mode has a couple of draggable boundaries so you can adjust according to what you're doing.


The chip graphics area has some additional controls, and can be hidden altogether.


Finally, in expert mode you can control the simulator and the graphics using additional URL parameters.



#####  
 Graphics help (advanced) 


All the graphics controls are immediately below the chip graphics area.


Use the keyboard to zoom in and out (z and x or < and >) and once you've zoomed in you can use the mouse to pan around, by clicking and dragging.


If you click on any shape in the chip, the status display (top right) will give you the coordinates and the node number, and the node name if it has one. If you click on a transistor gate, it'll give you the transistor number.  If you shift-click on a node, it'll highlight all the channel-connected nodes - so you can see which busses are connected by pass gates in the present clock cycle, which gates are writing a bus and which latches are reading it.


The second line of graphics controls allow you to select layer visibility. Chips are designed and fabricated as a series of thin layers in and on the silicon. With these controls you can get a clearer idea of the geometry in each layer.  For example, most of the long-distance connections are in metal, with some polysilicon. The logic gates themselves use diffusion and polysilicon only, in general.


The next line of graphics controls allow you to find the geometry corresponding to a node name, or number, or a collection of them.  The Clear Highlighting button also clears the logic level indication: all the diffusion becomes yellow, all the polysilicon purple, and the metal translucent.


The 'Animate during simulation' checkbox allows you to run the simulation faster, by skipping the logic level highlighting of the chip graphics.


'Hide Chip Layout' rearranges the page layout so you can concentrate on the right hand side panels - if you're looking at the logical behaviour and not the chip layout.


Finally, 'Link to this location' provides an URL which you can share, corresponding to the current pan and zoom, so you can discuss whichever interesting layout features you find. For example


-  
[instruction predecode](http://visual6502.org/JSSim/expert.html?panx=528.8&pany=182.8&zoom=7.4)
-  
[lower nibble of the ALU](http://visual6502.org/JSSim/expert.html?panx=233.5&pany=387.7&zoom=6.2)
-  
[instruction register and part of the PLA](http://visual6502.org/JSSim/expert.html?panx=486.7&pany=101.5&zoom=6.2)

#####  
 Running the program 


At the top of the right hand pane there is a series of buttons:


-  Run/Stop will free-run the simulation. It runs faster if the chip layout is not presently visible.

-  Reset will reset the chip and set time to zero.

-  Back and Forward step by a single clock phase.  Note the tabulation of machine state in the lower right pane.  When switching direction, you may wish to use the Clear Log button to avoid confusion.

-  Step will advance to the next write or the next instruction fetch.

-  Fast Forward will run for a configurable number of clock phases without updating the display or the internal trace buffer.  (The URL parameter headlesssteps defaults to 1000.)  This is useful for benchmarking.  Also, if headlesssteps is negative, the machine will free run but poll for input at that interval, which is useful for running interactive programs.


#####  
 Interacting with the program 


Any write to location $000f will cause output to the text box at the top of the lower right pane.


Location $D010 acts as a status port and $D011 acts as a data port for reading keyboard input. The status port reads 0 until a character is available.



#####  
 Modifying the program 


-  interactively, you can click on the memory map and change the content of each location. The memory map is in the top right-hand pane, below the table showing the machine state and the buttons controlling the simulator.

-  in the URL, you can use a= and d= to patch any addresses with data. So ...&a=400&d=eaea&... would put a couple of NOPs at $0400. Use the same technique if you wish to adjust the vectors at the top of memory. You can load quite long programs this way (in several sections, usually.)


#####  
 Tracing machine state 


The lower right pane offers some means of tracing signals of interest:


-  Trace More will add some pre-defined sets of signals to the tabulation

-  Trace Less removes those sets in the same order

-  Both the above can also be accessed using the loglevel URL parameter

  Trace these too: allows for a list of signals to be added to the tabulation

-  You may wish to Trace Less and then add back a different subset, or use a different order

-  You may with to explore the chip layout and find other signals to probe

  Log Up/Down allows for the tabulation to act in reverse order if you prefer not to keep scrolling to the bottom.

  Clear Log acts in the obvious way

-  It can be useful to use the single-step forward and backward in combination with log up/down to explore a few cycles of interest, because adding more signals will clear the tabulation but does not reset the cycle counter and does not clear the trace buffer: all signal activity is still stored.


#####  
 Busses and signals of interest 


You can use the 'Trace these too' box to list more signals which you're interested in - you might even use 'Trace less' several times to make this the exact list of signals to tabulate.


Some names are handled specially and are not individual signals or busses:


-  cycle is the count of cycles since reset

-  pc is the program counter, combining pcl and pch

-  p is the status register, combining p0 to p7 (there is no p5, and p1 and p2 are probed away from the other P signals)

-  tcstate collects the timing control bits, which are labelled 'clock1','clock2','t2','t3','t4','t5'

-  State is a more readable version of tcstate, which names the low bits, from T0 to T5. Note that T0 and T1 are not what they seem (link needed)

-  Fetch is blank unless the 'sync' pin is active, in which case it is a description of the 6502 opcode on the data bus (db).

-  Execute is a description of the 6502 opcode in the ir (Instruction Register)

-  plaOutputs is a list of all the active outputs from the instruction decode PLA

-  DPControl is a list of all the active control signals into the datapath

If you need to tabulate a signal which is in negative sense, use a leading minus. For example, '-pd' is the instruction predecode register.


The signal names are mostly taken from 
[Donald Hanson's block diagram](http://www.pagetable.com/?p=39)


#####  
 URL interface 


There's a variety of parameters which can be passed on the URL, to make it easy to share examples and discoveries as direct links into the simulator. In all cases these are passed like this:


```
 
[http://www.visual6502.org/JSSim?name1=value1&name2=value2](http://www.visual6502.org/JSSim/expert.html?panx=239.4&pany=352.7&zoom=10.7&steps=10)
```
 positioning the graphics window
 
[panx=240&pany=350&zoom=10](http://www.visual6502.org/JSSim/expert.html?panx=240&pany=350&zoom=10) select a larger canvas for improved graphical detail (uses more RAM)
 
[canvas=2000](http://www.visual6502.org/JSSim/expert.html?panx=240&pany=350&zoom=10&canvas=2000) suppress the simulation, for faster startup of a purely graphical session
 
[nosim=true](http://www.visual6502.org/JSSim/expert.html?panx=480&pany=100&zoom=6&nosim=t) suppressing graphics (same as the Hide Chip Layout button)
 
[graphics=false](http://www.visual6502.org/JSSim/expert.html?graphics=f) running for a fixed number of clock phases
 
[steps=10](http://www.visual6502.org/JSSim/expert.html?graphics=f&steps=10) see more groups of interesting signals in the tabulation
 
[loglevel=4](http://www.visual6502.org/JSSim/expert.html?graphics=f&steps=10&loglevel=4) add specific signals to the tabulation
 
[logmore=Execute,State,plaOutputs](http://www.visual6502.org/JSSim/expert.html?graphics=f&steps=10&logmore=Execute,State,plaOutputs) set the fastforward step count (for benchmarking, or interactive programs)
 
[headlesssteps=250](http://www.visual6502.org/JSSim/expert.html?graphics=f&steps=10&headlesssteps=250) load a test program, or patch memory contents
 
[a=0000&d=a2d0de2143](http://www.visual6502.org/JSSim/expert.html?graphics=f&loglevel=2&steps=50&a=0000&d=a2d0de2143de34236cff20eaea) adjust the reset vector
 
[r=0002](http://www.visual6502.org/JSSim/expert.html?graphics=f&steps=10&r=0002) set up some input pin transitions (Reset, IRQ, NMI, RDY)
 
[reset0=12&reset1=13](http://www.visual6502.org/JSSim/expert.html?graphics=f&loglevel=2&a=0000&d=a2559ae8e8e8&logmore=res,plaOutputs,DPControl&reset0=12&reset1=13&steps=40) 
[nmi0=4&nmi1=8](http://visual6502.org/JSSim/expert.html?graphics=f&loglevel=2&steps=50&a=fffa&d=2000&a=0020&d=e840&r=0010&nmi0=4&nmi1=8&logmore=nmi) 
[irq0=3&irq1=20](http://visual6502.org/JSSim/expert.html?graphics=f&loglevel=2&steps=50&a=0011&d=58&a=fffe&d=2000&a=0020&d=e840&r=0010&irq0=3&irq1=20&logmore=irq) 
[[1]](http://visual6502.org/JSSim/expert.html?graphics=f&a=0&d=a955a2338b0feaea&steps=17&loglevel=2&logmore=ir,rdy,idl,idb,dasb,sb,dpc2_XSB,dpc23_SBAC,dpc24_ACSB,dpc25_SBDB,dpc43_DL/DB&rdy0=10&rdy1=14&time=12&databus=c4) (RDY not yet in the released version)
 Note on timings of input transitions
 Since 2013-06-27, the displayed data for input transitions is shown one phase earlier than previously. This fixes a bug. When re-using a URL with the new version, all the chip behaviour is unchanged, but shifting the display of the input signals has made everything appear correct too.
 check every signal value against a golden checksum (for checking simulator code changes)
 
[steps=99&checksum=0fa98aab](http://visual6502.org/JSSim/expert.html?graphics=f&loglevel=-1&steps=99&checksum=0fa98aab) label some points of interest
 
[Annotated floorplan](http://visual6502.org/JSSim/expert.html?nosim=t&label=PLA,100,1169,2328,8393,934&label=Datapath,100,2143,8820,7676,5689&label=Control,100,3333,4083) of the 
[6502](index.php?title=MOS_6502) 
[Some functional blocks](http://visual6502.org/JSSim/expert.html?nosim=t&label=A,50,5040,8820,5328,5689&label=ALU,50,2814,8820,4525,5689&label=DAdj,40,4525,8820,5040,5689&label=ID,50,7365,8820,7676,5689&label=IR,50,8432,2332,9124,984&label=PC,50,5559,8820,6819,5689&label=PD,50,8424,3536,9256,2464&label=PLA,100,1169,2328,8393,934&label=S,50,2490,8820,2814,5689&label=TimC,40,600,1926,1174,604&label=X,50,2317,8820,2490,5689&label=Y,50,2143,8820,2317,5689) of the 
[6502](index.php?title=MOS_6502)
The final reference on the URL capabilities is 
[the source code](https://github.com/trebonian/visual6502/blob/master/expertWires.js#LC134).



###  
 See also 


See also the 
[ChipSim Simulator](index.php?title=The_ChipSim_Simulator) page which goes into some detail on the implementation and history of the simulator, and have a look at 
[the source code](https://github.com/trebonian/visual6502).


See also the guide to interpreting NMOS layout at the 
[Visual circuit tutorial](http://wiki.nesdev.com/w/index.php/Visual_circuit_tutorial) on the NesDev wiki.



![Attribution-NonCommercial-ShareAlike 3.0 Unported](http://i.creativecommons.org/l/by-nc-sa/3.0/88x31.png)

