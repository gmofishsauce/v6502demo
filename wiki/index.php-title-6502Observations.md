**INCOMPLETE DRAFT OF RECOVERED WIKI PAGE**

# 6502Observations - VisualChips


	

	
	


## 6502Observations


	

		


#### From VisualChips


		

		

		

We've found some interesting things on the 6502, from the layout level, up through circuit level to the programmer visible level.



##  
 Programmer Visible 


Notes here on bugs and undocumented behaviour.


-  
[BRK, the B bit](index.php?title=6502_BRK_and_B_bit), and other interrupts

-  
[Timing of Interrupt Handling](index.php?title=6502_Timing_of_Interrupt_Handling) noting that a taken branch delays interrupt handling, also that CLI/PLP allow one further instruction to execute, unlike RTI.

-  
[Unsupported or undocumented opcodes](index.php?title=6502_Unsupported_Opcodes) such as SAX and XAA

-  
[The ROR bug](index.php?title=6502_ROR_bug) which is found only in rare early devices

See also 
[our catalogue of 6502 test programs](index.php?title=6502TestPrograms), useful to verify simulators or emulators.



##  
 Circuit and Logic 


Notes here on timing fixes and non-digital circuit techniques, and departures from NMOS design style orthodoxy.


-  
[Signs of a fix](index.php?title=6502_datapath_control_timing_fix) to datapath control timing


##  
 Layout 


Notes here on the traces of bug fixes, and remnants of the original 6501 layout.


-  
[Traces in the layout](index.php?title=6502_traces_of_6501) of the original 6501 part which was withdrawn after legal wrangling


![Attribution-NonCommercial-ShareAlike 3.0 Unported](http://i.creativecommons.org/l/by-nc-sa/3.0/88x31.png)

