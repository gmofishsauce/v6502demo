**INCOMPLETE DRAFT OF RECOVERED WIKI PAGE**

# 6502 Unsupported Opcodes - VisualChips


	

	
	


## 6502 Unsupported Opcodes


	

		


#### From VisualChips


		

		

		

The 6502 is famous for doing interesting and sometimes useful things when the program includes invalid (or unspecified) opcodes.


For a list of all opcodes and some explanation of what they do, see 
[6502 all 256 Opcodes](index.php?title=6502_all_256_Opcodes).


The visual6502 simulator can help when investigating what these opcodes do, and why - see below for a few cases and pointers for exploration.



###  
 examples 


-  
[LAX](http://visual6502.org/JSSim/expert.html?graphics=f&steps=16&a=5555&d=44&a=0&d=af5555ea&loglevel=2&logmore=dpc3_SBX,dpc23_SBAC,plaOutputs,DPControl) will load both A and X - notice signals SBX and SBAC which control the writes to X and to A.

-  
[KIL](http://visual6502.org/JSSim/expert.html?graphics=f&steps=26&a=0&d=ea42eaea&loglevel=2) will put the T-state counter into an unrecoverable state

-  
[XAA #$5A](http://visual6502.org/JSSim/expert.html?graphics=f&steps=26&a=0&d=a9ffa2f08b5aeaea&loglevel=2&logmore=DPControl) (also known as ANE) with A=FF

-  and 
[with A=00](http://visual6502.org/JSSim/expert.html?graphics=f&steps=26&a=0&d=a900a2f08b5aeaea&loglevel=2&logmore=DPControl) shows A is OR with 00 before AND with X and the immediate value

-  for more detail see the explanation page: 
[6502 Opcode 8B (XAA, ANE)](index.php?title=6502_Opcode_8B_(XAA,_ANE))

###  
 some background 


Beware: different revisions of 6502 and versions from different manufacturers may have different behaviours.


For some of these opcodes, the chip does something logically predictable and our model has the same behaviour. But there may be opcodes which are not logically predictable, because they cause marginal voltages on the chip as different drivers fight one another, or a node which is undriven is sampled at a later time. In those cases, our visual6502 simulator, which is just a switch-level simulator with a couple of coarse heuristics for modelling contention and charge storage, won't do the same as a chip.


In fact, as some opcodes produce results which vary from one chip to another, no deterministic simulator could be 'accurate'.  (A simulator could let you know that something is amiss)


But note that the underlying circuit data which we now have includes transistor strengths and an approximation of capacitative load: it could easily be extended for resistance and more accurate capacitance. So a more refined (lower level) simulation might shed more light on these undocumented opcodes. In fact, 
[the FPGA model](https://github.com/pmonta/FPGA-netlist-tools) works differently - it moves charge from one node to another - and it might be more accurate for the difficult cases.



###  
 resources 


-  back to parent page 
[6502Observations](index.php?title=6502Observations)
-  
[Michael Steil's presentation at 27C3](http://www.youtube.com/watch?v=N9DYmlprCKA#t=5m20) youtube link direct to section on illegal opcodes

-  
[How MOS 6502 Illegal Opcodes really work](http://www.pagetable.com/?p=39) on Michael Steil's blog

-  
[64doc.txt](http://www.viceteam.org/plain/64doc.txt) by VICE team

-  
[Extra Instructions Of The 65XX Series CPU](http://www.ffd2.com/fridge/docs/6502-NMOS.extra.opcodes) by Adam Vardy

-  
[6502 Undocumented Opcodes](http://members.chello.nl/taf.offenga/illopc31.txt) by Freddy Offenga

-  
[6502/6510/8500/8502 Opcode matrix](http://www.oxyron.de/html/opcodes02.html) by "Graham"

-  
[Full 6502 Opcode List Including Undocumented Opcodes](http://bbc.nvg.org/doc/6502OpList.txt) by J.G.Harston

-  
[Michael Steil's presentation at 27C3](http://www.pagetable.com/?p=517) (pagetable.com links to 6 sections on youtube)

-  
[Vice BUGS document](http://www.viceteam.org/plain/BUGS) mentions XAA being used in a Mastertronic loader

-  
[An examination of an early tape loader](http://www.atlantis-prophecy.org/recollection/?load=online&issue=1&sub=article&id=4) by Fungus/Nostalgia/Onslaught


![Attribution-NonCommercial-ShareAlike 3.0 Unported](http://i.creativecommons.org/l/by-nc-sa/3.0/88x31.png)

