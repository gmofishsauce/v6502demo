**INCOMPLETE DRAFT OF RECOVERED WIKI PAGE**

# 6502 Interrupt Recognition Stages and Tolerances - VisualChips


	

	
	


## 6502 Interrupt Recognition Stages and Tolerances


	

		


#### From VisualChips


		

		

		

6502 Interrupt Recognition Stages and Tolerances


The following is based upon drawing the node and transistor networks out on paper from visual6502 data, and conducting experiments with the simulator. References are made to 6502 clock states that are described in 
[ 6502 Timing States ](index.php?title=6502_Timing_States), which may be used as a primer for this exposition.



### Contents


- [1 Stages of Interrupt Recognition](#stages-of-interrupt-recognition)
- [2 Clearing the Interrupt Stages](#clearing-the-interrupt-stages)
- [3 Tolerances](#tolerances)
- [4 Demonstrations](#demonstrations)
- [5 External References](#external-references)
- [6 Further Reading](#further-reading)

###  
 Stages of Interrupt Recognition 


There are four stages of recognition of interrupt signals within the 6502, numbered from zero to three.


Stage 0 of interrupt recognition converts the asynchronous interrupt signals into synchronous interrupt signals. It is a subnetwork of nodes and transistors that react to the change in the interrupt line only when one particular phase is in effect. The synchronous interrupt signals change with the phase changes caused by clock pulses, allowing them to mesh with the processor's cycle-structured operation and thus instruction execution. The output of the sync complexes change only during one half of a cycle, and stay stable for the other half of a cycle.


Stage 0 of interrupt recognition (or clearing of recognition) happens at phase 2. When an interrupt line changes state during phase 1, the output node of the complex will switch state when phase 2 is clocked in. When an interrupt line changes state during phase 2, the output node of the complex will switch state immediately.


The synchronization of IRQ and NMI have identical topologies and the output for a recognized interrupt is high at phase 2. RES has a variation on the sync topology that makes its output the opposite (low) at phase 2 for a recognized RES.


Stage 1 of interrupt recognition happens at phase 1, immediately after the phase 2 of stage 0 recognition. Later on we'll see an exceptional case for the immediacy of the NMI interrupt.


Stage 1 of RES is represented by the node RESP going high and starting a signal chain that will set the 6502 clock to the T0 state one cycle later, at the next phase 1.


Stage 1 of NMI is represented by node ~NMIG grounded low. In addition to being stage 1 of pending NMI status, it is also responsible for selecting the NMI vector used by BRK.


Stage 1 of IRQ is represented by node IRQP high.


Independently, both ~NMIG and IRQP connect two signal chains that ordinarily stay unconnected. The sending chain carries a signal from the clock and PLA that can be described as clock-T0/branch-T2, and the receiving chain sets NMI/IRQ stage 2 recognition. When the IRQ disable bit is high, it prevents the connection enabled by IRQP. ~NMIG low will always enable the connection (one of the ways it exerts a higher priority).


The progression from stage 1 to stage 2 differs markedly at this point between RES and NMI/IRQ together. RES always proceeds to stage 2 after stage 1 in the same cycle, at phase 2. NMI and IRQ have to wait until the 6502 clock reaches its T0 or its T2 state, depending upon what kind of instruction is executing. This is the feature that allows the currently running instruction to finish before being interrupted by IRQ or NMI. IRQ and NMI's stage 2 also happens at phase 2 during the required clock state.


Stage 2 of interrupt recognition enables the next fetch cycle to invoke the interrupt instruction BRK. It also drives the Break bit low.


Stage 2 of RES recognition is represented by node RESG going high. As stated above, it happens when phase 2 is clocked in immediately after the phase 1 that recognized stage 1 of RES. RESG is responsible for putting the 6502 into write disabled mode, starting with the next phase 1. It is also responsible for selecting the RES vector used by BRK.


Stage 2 of NMI and IRQ recognition is represented by node INTG going high. It happens when T0 phase 2 is clocked in for non-branch instructions and branches that page-cross. It also happens when T2 phase 2 is clocked in for all branch instructions.


Stage 3 of interrupt recognition arrives with the next fetch cycle (T1) that the clock advances to. RESG or INTG high causes the fetch cycle to prepare to substitute a BRK instruction into the IR in phase 2 instead of the opcode that was read from memory in phase 1. The following T2 phase 1 sees the BRK instruction loaded into the IR and issuing its earliest execution signals. Seven cycles later, the first instruction of the appropriate interrupt handler is starting.


With the RES interrupt, the RES line must be released in order to allow stage 3 actions to be taken. While RES is down, stage 1 recognition is maintained, and that sends a signal to the 6502 clock to reset to its T0 state and suppress opcode fetch. Continued assertion of clock reset while the clock tries to advance results in a time state of T0 T+ on subsequent cycles, and fetch will still be suppressed by RES stage 1. Releasing RES stops the continued resetting of the clock, allowing it to advance to load and execute the BRK instruction for stage 3.


The latency for releasing the clock after raising the RES line is the same as the latency for recognizing stage 1 after lowering RES, plus one cycle. That's a total of just under two cycles. Stage 0 recognition must clear first at phase two of the cycle in which it was released, followed by stage 1 clearing at the very next phase 1, then one more cycle for stage 1 clearance to reach the clock to stop resetting it. Re-summarized, there will be the RES raise cycle, a cycle where stage 1 is clear and the clock is still affected by the previous cycle's stage 1, then the first cycle of the clock being free and fetch permitted. That will be 3 or 4 clock transitions: 4 if RES is raised during phase 1, 3 if RES is raised during phase 2.


IRQ recognition tabulation


```
+-----------------------+-------------------------------------------------------------------------------+
| Stage 0               | Any cycle, phase 2                                                            |
|-----------------------|-------------------------------------------------------------------------------|
| Stage 1               | Any cycle + 1, phase 1                                                        |
|-----------------------|-------------------------------------------------------------------------------|
| Stage 1 signal bridge | If IRQ disable bit is low: immediately at Stage 1                             |
|                       | Else upon IRQ disable bit reset to low from high (and stage 1 still persists) |
|-----------------------|-------------------------------------------------------------------------------|
| Stage 2               | If IRQ disable bit is low (stage 1 signal bridge connected):                  |
|                       | T2 phase 2 for all branch instructions                                        |
|                       | T0 phase 2 for non-branch instructions and page-crossing branches             |
|                       | Break bit goes low                                                            |
+-----------------------+-------------------------------------------------------------------------------+

```

NMI recognition tabulation


```
+---------+-------------------------------------------------------------------+
| Stage 0 | Any cycle, phase 2                                                |
|---------|-------------------------------------------------------------------|
| Stage 1 | Any cycle + 1, phase 1                                            |
|         | Stage 1 signal bridge always happens for NMI at Stage 1           |
|         | NMI vector selection on                                           |
|---------|-------------------------------------------------------------------|
| Stage 2 | T2 phase 2 for all branch instructions                            |
|         | T0 phase 2 for non-branch instructions and page-crossing branches |
|         | Break bit goes low                                                |
+---------+-------------------------------------------------------------------+

```

RES recognition tabulation


```
+---------+-----------------------------+
| Stage 0 | Any cycle, phase 2          |
|---------|-----------------------------|
| Stage 1 | Any cycle + 1, phase 1      |
|         | Clock reset signal starts   |
|---------|-----------------------------|
| Stage 2 | Any cycle + 1, phase 2      |
|         | Write disable signal starts |
|         | RES vector selection on     |
|         | Break bit goes low          |
|---------|-----------------------------|
| *       | Any cycle + 2, phase 1      |
|         | Clock reset to T0           |
|         | Processor is write disabled |
+---------+-----------------------------+

```


###  
 Clearing the Interrupt Stages 


Clearance of the various stages of interrupt recognition is divided between merely releasing the corresponding interrupt line(s) to go high again, and the BRK instruction explicitly clearing the recognition. NMI edge detection reset depends upon both.


Tabulation of what BRK clears


```
+------------+----------------------------------------------------------------------------------------------------------+
|    When    | BRK actions                                                                                              |
|============|==========================================================================================================|
| T6 phase 2 | * NMI/IRQ stage 2 cleared (INTG low).                                                                    |
|            | * IRQ stage 1 directly temporarily interdicted for one cycle (blocks IRQ stage 1 without clearing it),   |
|            |   apparent as node Pout2 low. If there was no NMI to recognize, then this will disconnect NMI/IRQ stage  |
|            |   2 recognition from the clock-T0/branch-T2 signal chain.                                                |
|------------|----------------------------------------------------------------------------------------------------------|
| T0 phase 1 | * RES stage 2 cleared (RESG low) if RES stage 1 was cleared when T6 phase 1 was clocked in or earlier.   |
|            |   Processor is immediately enabled to perform writes to memory.                                          |
|            | * NMI stage 1 cleared (~NMIG high). This always disconnects NMI/IRQ stage 2 recognition from the         |
|            |   clock-T0/branch-T2 signal chain.                                                                       |
|            | * IRQ disable bit set (blocks IRQ stage 1 without clearing it).                                          |
|------------|----------------------------------------------------------------------------------------------------------|
| T0 phase 2 | * RES stage 2 cleared (RESG low) if RES stage 1 was cleared when T0 phase 1 was clocked in. Processor is |
|            |   enabled to perform writes to memory on the next phase 1 (next clock transition).                       |
+------------|----------------------------------------------------------------------------------------------------------+

```

The Break bit goes high again when all forms of stage 2 interrupt recognition (RESG OR INTG) have been cleared. Tabulated above, that ranges from as early as T6 phase 2 to as late as T0 phase 2.


NMI falling edge detection reset requires the action of BOTH the BRK instruction above AND releasing the NMI line to go high. They may happen in either order. The reset is not independently controlled by the NMI line rising high again by itself. Edge detection is reset by NMI already being up when the BRK instruction clears NMI stage 1 recognition, or by coming up after that.


The soonest that NMI can come down and cause NMI recognition again is just after T0 phase 1 is clocked in. That NMI will be recognized when T+ T1 phase 1 is clocked in. That will allow the first instruction of the NMI handler to run before another BRK instruction starts due to the new NMI.


Stage 1 of IRQ and RES recognition is cleared only by letting the respective line rise high again. The next phase 1 clocked in after releasing the line will clear stage 1.


Stage 0 of all the interrupts are cleared by raising the respective line, and clearance is recognized during phase 2.



###  
 Tolerances 


In real world systems, the interrupt lines stay down until the issuing device is answered, so minimal thresholds for invoking respective interrupts are not an issue. They can stay down for hundreds to hundreds of thousands of clock transitions typically. RES invoked manually can last at least a few tenths of seconds, as do power on invocations by dedicated hardware.


For us hackers of visual6502, extremely small durations of interrupt line activity are more relevant. Merely one clock transition while a line is down is sufficient. There are niche cases where more is required. Here's the guidance.


Absolute minimum action needed to invoke the respective interrupt.


IRQ:


It's a little complex for IRQ.


For non-branch instructions and page-crossing branch instructions, put the IRQ line down during phase 2 of the clock cycle immediately before T0 (depends upon instruction, and watch out for indexed memory-read instructions that terminate a cycle early when they don't page-cross), then clock in T0 phase 1, IRQ line up during T0 phase 1.


For all branch instructions, IRQ line down during phase 2 of T1 (fetch cycle), clock in T2 phase 1, IRQ line up during T2 phase 1.


The page-crossing branch instructions may be treated either way.


NMI:


NMI line down during phase 2, clock in phase 1, NMI line up during phase 1.


RES:


RES line down during phase 2, clock in phase 1, RES line up during phase 1.


In all the above cases, the interrupt line may be put down one phase earlier just after clocking in phase 1 of the same cycle, and may be let up one phase later during phase two just before clocking in phase 1 of the next cycle.


Maximal action needed to invoke the respective interrupt (such as when clock state and phase are unknown or not being monitored).


IRQ:


Varies with execution time and current state of the currently running instruction, due to IRQ's level sensitivity: Simply put the IRQ line down and wait for the current instruction to finish: release when SYNC goes high (after having been low), which is T1 phase 1 of the invoked BRK instruction.


NMI:


NMI line down and kept down for at least six (6) clock transitions. This covers the worst case of NMI coming down during T5 phase 1 of a BRK instruction, which has protection against NMI stage 1 recognition causing a mixed vector indirect jump (low byte of BRK/IRQ and high byte of NMI), and also protects against BRK failing to clear NMI stage 1 recognition.


After six transitions when put down at the worst case time, the NMI line may be safely raised during T1 phase 1, and NMI stage 1 will have been successfully recognized when that state, T1 phase 1, was clocked in. The first instruction of the IRQ/BRK handler will run and then stage 2 of NMI will be recognized, causing the next fetch cycle to start a new BRK instruction that jumps to the NMI handler.


Putting NMI down at T5 phase 1 or later and raising it back up again before T1 phase 1 will cause the "lost NMI" condition noted in 
[ 6502 Timing of Interrupt Handling ](index.php?title=6502_Timing_of_Interrupt_Handling).


RES:


RES line down and kept down for at least six (6) clock transitions. This covers the worst case of RES coming down during T4 phase 1 of a BRK instruction. The long duration is needed to keep stage 2 of RES recognition alive despite the BRK instruction's attempt to shut it off.


RES stage 1 recognition resets the 6502 clock to the T0 state in the next cycle, but the BRK instruction uses an extension of the clock that RES stage 1 cannot affect. The clock ends up with a state called T0 T6 instead of pure T0. It is the T6 part that sends the signal to clear stage 2 of RES recognition, and it is the reason that the RES line must be held down extra long. The cycle after this case's T0 T6 (and otherwise pure T6) cycle is when stage 2 is normally shut off.


Stage 1 of RES recognition must be kept true during the cycle after T0 T6 to prevent stage 2 being cleared. The RES line may be safely raised during phase 1 of the cycle after T0 T6. Stage 1 of RES recognition will then not become false until phase 1 of the 
next
 cycle is clocked in (2nd after T0 T6), after the danger has passed.


With stage 2 having survived the tail end of a BRK instruction, it will cause the next fetch cycle to start a new BRK instruction that jumps to the RES handler.


Tabular synopsis of the worst-case BRK events that affect RES stage 2, dual-labeled with clock time states for RES line high (Normal) and for RES line low (RES Altered). Alternate labeling in terms of the T0 T+ clock states is clearer than the prose terms of, "cycle after T0 T6", and provides an easy cross-reference with the earlier tabulation of what BRK clears. Recall that T0 T+ clock states are caused by the continued assertion of clock reset while the clock tries to advance, due to the extended time that RES is down.


```
+--------+-------------+-------------------------------------------------------------------------+
| Normal | RES Altered | Events                                                                  |
|========|=============|=========================================================================|
|   T4   |      T4     | RES line down during phase 1.                                           |
|        |             | Stage 0 recognized at phase 2.                                          |
|--------|-------------|-------------------------------------------------------------------------|
|   T5   |      T5     | Stage 1 recognized (RESP) at phase 1.                                   |
|        |             | Stage 2 recognized (RESG) at phase 2.                                   |
|--------|-------------|-------------------------------------------------------------------------|
|   T6   |      T0 T6  | Stage 1 resets clock at phase 1 (& later cycles).                       |
|        |             | THIS CYCLE phase 2 initiates stage 2 clear signal (node brk-done high). |
|--------|-------------|-------------------------------------------------------------------------|
|   T0   |  1st T0 T+  | Stage 2 clear effective during both phases.                             |
|        |             | Stage 1 must persist to counter it.                                     |
|        |             | RES may be released during phase 1.                                     |
|--------|-------------|-------------------------------------------------------------------------|
|  T+ T1 |  2nd T0 T+  | Stage 2 clear signal no longer exists.                                  |
|        |             | Stage 1 allowed to have gone false.                                     |
+--------+-------------+-------------------------------------------------------------------------+

```

What happens when RES is let up too early? Opcode execution jumps to one of two nonsense addresses depending upon when RES is let up. The possibilities are tabulated below, along with repeating the recommended hold time case.


```
+-----------------------------------------------------------------------------------------------------------------------+
| With RES down (stage 0 recognized) when T5 phase 1 is clocked in, then...                                             |
+=======================================================================================================================|
| RES up (stage 0 cleared) when T0 T6 phase 1 is clocked in:                                                            |
|     The BRK ends one cycle early and jumps to an address of <RES low>FD, where <RES low> is the low byte of the RES   |
| vector, and it appears as the high byte of the jumped-to address. FD for the low byte of the jumped-to address is a   |
| copy of the low byte of where the RES vector high byte is located (FFFD).                                             |
|     A new BRK instruction for the RES is NOT invoked.                                                                 |
|-----------------------------------------------------------------------------------------------------------------------|
| RES up (stage 0 cleared) when 1st T0 T+ phase 1 is clocked in:                                                        |
|     The BRK ended one cycle early and is followed by a non-fetch held-clock cycle of T0 T+, then jumps to an address  |
| of <<RES low>FD><RES low>. The address formed for the opcode fetch in the previous case is used to read the high byte |
| of the address for this case. Meanwhile, the low byte of the RES vector appears in its proper place as the low byte   |
| of the address for this case.                                                                                         |
|     A new BRK instruction for the RES is NOT invoked.                                                                 |
|-----------------------------------------------------------------------------------------------------------------------|
| RES up (stage 0 cleared) when 2nd T0 T+ phase 1 is clocked in:                                                        |
|     The recommendation for holding down RES long enough is satisfied.                                                 |
|     The BRK ended one cycle early and is followed by two non-fetch held-clock cycles of T0 T+, then invokes a new BRK |
| instruction for the RES.                                                                                              |
|     Control is transferred to the RES handler when it finishes (unless it is disturbed by RES going down again).      |
+-----------------------------------------------------------------------------------------------------------------------+

```

Holding RES down for a cycle longer beyond the minimum time merely adds another T0 T+ cycle to the end and still results in a new BRK for RES, ad infinitum.



###  
 Demonstrations 


For all the following demonstrations, the RES vector is set to F933.


[RES up before T0 T6 phase 1](http://visual6502.org/JSSim/expert.html?graphics=f&steps=26&logmore=Execute,res,RESP,RESG,State,tcstate,TState,Phi&a=0200&d=0001&a=33FD&d=4C33F9&a=4C33&d=4C33F9&a=F933&d=4C0002&r=F933&reset0=12&reset1=14)

A soft BRK instruction is interrupted by a brief RES that causes it to jump to <RES low>FD (33FD) where a JMP instruction redirects it to the RES handler again.


Schedule of the interrupts (Halfcycle numbers are 0-based):


```
Halfcycle 12 RES0 during T4 phase 1 of soft BRK (phase 2 at 13 would work just as well)
Halfcycle 14 RES1 during T5 phase 1 of soft BRK (or phase 2 at 15)

```

[RES up before 1st T0 T+ phase 1](http://visual6502.org/JSSim/expert.html?graphics=f&steps=28&logmore=Execute,res,RESP,RESG,State,tcstate,TState,Phi&a=0200&d=0001&a=33FD&d=4C33F9&a=4C33&d=4C33F9&a=F933&d=4C0002&r=F933&reset0=12&reset1=16)

A soft BRK instruction is interrupted by a one-cycle-longer RES that causes it to jump to <<RES low>FD><RES low> (<33FD>33 => 4C33) where another JMP instruction redirects it to the RES handler again.


Schedule of the interrupts (Halfcycle numbers are 0-based):


```
Halfcycle 12 RES0 during T4 phase 1 of soft BRK (or phase 2 at 13)
Halfcycle 16 RES1 during T0 T6 phase 1 of soft BRK (or phase 2 at 17)

```

[RES down long enough](http://visual6502.org/JSSim/expert.html?graphics=f&steps=38&logmore=Execute,res,RESP,RESG,State,tcstate,TState,Phi&a=0200&d=0001&a=33FD&d=4C33F9&a=4C33&d=4C33F9&a=F933&d=4C0002&r=F933&reset0=12&reset1=18)

A soft BRK instruction is interrupted by a two-cycle-longer (long enough) RES that causes a new BRK instruction to be started for RES, which jumps normally to the RES handler.


Schedule of the interrupts (Halfcycle numbers are 0-based):


```
Halfcycle 12 RES0 during T4 phase 1 of soft BRK (or phase 2 at 13)
Halfcycle 18 RES1 during 1st T0 T+ phase 1 of soft BRK (or phase 2 at 19)

```


Coding of the program used in all of the demonstrations:


```
;                User code
0200 BRK +01     ; Soft BRK interrupted by late RES
;
;                Intercept earliest-release RES jump result
33FD JMP F933    ; Redirect to RES handler from <RES low>FD
                 ; JMP opcode of 4C used as high byte of next-earliest-release jump point
;
;                Intercept next-earliest-release RES jump result
4C33 JMP F933    ; Redirect to RES handler from <<RES low>FD><RES low>
;
;                RES handler (where visual6502 starts running code when finished starting up)
F933 JMP 0200    ; Jump to user code

```


###  
 External References 


"lost NMI" in 
[ 6502 Timing of Interrupt Handling ](index.php?title=6502_Timing_of_Interrupt_Handling)

[ 6502 Timing States ](index.php?title=6502_Timing_States)


###  
 Further Reading 


[ 6502 Interrupt Hijacking ](index.php?title=6502_Interrupt_Hijacking)


![Attribution-NonCommercial-ShareAlike 3.0 Unported](http://i.creativecommons.org/l/by-nc-sa/3.0/88x31.png)

