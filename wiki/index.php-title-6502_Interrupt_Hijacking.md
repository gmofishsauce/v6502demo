**INCOMPLETE DRAFT OF RECOVERED WIKI PAGE**

# 6502 Interrupt Hijacking - VisualChips

## 6502 Interrupt Hijacking

#### From VisualChips

6502 Interrupt Hijacking

The following is based upon drawing the node and transistor networks out on paper from visual6502 data, and conducting experiments with the simulator. In explaining the various behaviors, references are made to 6502 clock states and stages of interrupt recognition that are described in [6502 Timing States](index.php-title-6502_Timing_States) and [6502 Interrupt Recognition Stages and Tolerances](index.php-title-6502_Interrupt_Recognition_Stages_and_Tolerances), which may be used as primers for this exposition.

**Contents**

- [Introduction](#introduction)
- [NMI Hijacking IRQ/BRK](#nmi-hijacking-irq.2fbrk)
- [RES Hijacking NMI and IRQ/BRK](#res-hijacking-nmi-and-irq.2fbrk)
- [BRK Protected from NMI vs. Unprotected from RES](#brk-protected-from-nmi-vs.-unprotected-from-res)
- [Demonstrations](#demonstrations)
- [External References](#external-references)

### Introduction

Systems typically engineered around the 6502 generate interrupt signals that stay low for very appreciable lengths of time in terms of the applied clock rate. Manual resets may keep the RES line down at least a few tenths of a second when a reset button is quickly struck and released, and resets invoked by hardware when power is first applied may last for 0.1 to 0.3 seconds. Devices invoking NMI and IRQ will keep the respective lines down until the 6502 programmatically answers the interrupt with a finishing access to the device, which may be in the hundreds of microseconds to the 1+ milliseconds level. All of the above durations range from hundreds to hundreds of thousands of clock transitions.

Some obscure niches of behavior may be observed instead when the interrupt lines are held low for only a handful of clock transitions or less.

### NMI Hijacking IRQ/BRK

First introduced to this wiki by the link, "late NMI", in [6502 Timing of Interrupt Handling](index.php-title-6502_Timing_of_Interrupt_Handling), it is possible for higher-priority interrupts to hijack lower priority interrupts during the BRK instruction that is serving them.

The "late NMI" is an example of a full hijack, where a higher priority interrupt arrives after the BRK instruction for a lower priority interrupt has already started. The higher priority interrupt changes the addresses used when the BRK instruction actually fetches the vector during clock states T5 and T6: the vector selection is not fixed and "remembered" when the BRK instruction starts, or at any time before. Vector selection is independently parameterized by the NMI and RES interrupts, instead.

For NMI to hijack an IRQ (or a soft BRK instruction), stage 1 of its recognition may appear as late as T5 phase 1 during the BRK instruction's execution. This requires the NMI line to go down no later than the end of clock cycle T4 (up to just before clocking in T5 phase 1). With the node ~NMIG low during T5 and T6, the low and high bytes of the vector for indirect jump to the NMI handler are fetched together.

This is perfectly allowable if an IRQ caused the hijacked BRK, as NMI is higher priority than IRQ, and is serviced first. The IRQ will get its chance for attention after the RTI from the NMI handler has finished.

To recover from hijacking a soft BRK, the NMI handler must check the Break bit in the processor status register written to the stack. If it is true, then the handler must finish by performing a JMP ($FFFE) instead of an RTI. If it does an RTI for a soft BRK, then the soft BRK ends up being ignored: execution will resume with code after the BRK instruction instead of being processed by the system's soft BRK handler. The jump indirect to the IRQ/BRK handler will cause the soft BRK to be processed instead of missed. This special end for NMI would not have been necessary without a hijacking phenomenon.

Back to the subject of actual hijackings: If NMI could appear one cycle later, it would affect only the fetch of the high byte of the vector, while the low byte would have already been fetched for IRQ/BRK: it would be a half-hijack.

Unfortunately for this scenario (or fortunately, from the designer's P.O.V.), the 6502 has some explicit engineering to prevent an NMI half-hijacking IRQ and BRK. The node chain ~VEC, pipe~VEC, 1578, 1368 is the secret sauce. ~VEC is low during T5 and T6 of a BRK instruction, which are the cycles that internally command the low and high byte fetches, respectively, of the jump vector (each internal command takes one cycle to appear at the address pins and Read/Write pin externally). The node pipe~VEC is connected to ~VEC only during phase 2. pipe~VEC grounds 1578 and it grounds 1368 in turn. The therapeutic effect is that node 1368 is kept grounded from T5 phase 2 through T0 phase 1. That prevents NMI low from being passed through to cause the NMI to be recognized at stage 1 and affect the vector fetches. NMI stage 1 is not allowed to be recognized through all of T6 and T0 at the tail end of BRK execution. As long as the NMI line stays down, the NMI will finally be stage-1-recognized at T1 phase 1. That will allow the first instruction of the IRQ/BRK handler to run before the BRK for the NMI is started.

If NMI is released under the above circumstances before T1 phase 1 is clocked in, then the NMI is entirely missed. See also the "lost NMI" condition noted in [6502 Timing of Interrupt Handling](index.php-title-6502_Timing_of_Interrupt_Handling).

### RES Hijacking NMI and IRQ/BRK

How about RES hijacking NMI and IRQ/BRK? It turns out that full hijack can't quite happen, but half-hijack can.

To (attempt to) invoke a RES full hijack, the RES line must be brought low strictly during clock cycle T4 of the running BRK instruction (that's from just after clocking in T4 phase 1 to just before clocking in T5 phase 1). It must be released strictly during T5 (apply analogous boundaries).

The RES interrupt's effect upon the 6502 clock ends up getting in the way of the full hijack situation. RES stage 2 recognition is in control of selecting the vector at T5, so it fetches the low byte correctly in the next cycle. Things start to go awry on that next cycle (T6) because RES recognition stage 1 is also affecting the clock by that time, forcing it to also reset to T0. T6 still also remains in effect, though, because the BRK instruction uses an extension of the clock that RES reset cannot affect, and T6 is a part of that extension. The combined T0 T6 clock state causes actions in the same cycle that normally happen in separate cycles, and the T0 part also advances the schedule to the opcode fetch cycle (T1) by one cycle too early for fetching the vector high byte. Instead of having the full intact RES vector to address the next instruction with, it has the low byte of the RES vector used as the high byte of the opcode address, and the low byte of the address of the high byte of the RES vector appears as the low byte of the opcode address.

Restated: <RES low>FD is fetched instead of <RES high><RES low> for the jumped-to opcode. The 'FD' comes from FFFD addressing the high byte of the RES vector. Due to being both T0 and T6, the T6 phase 2 part of the state was still controlling the low byte of the address at fetch phase 1, and the T0 phase 2 part of the state was controlling the high byte of the address at fetch phase 1. Normally, the high byte of the vector is what was just read and is used for the high byte of the fetch address, but the low byte was what was just read and available instead.

The foiled full hijack behavior is actually a subset of what happens when RES is released too early after having been put down during T4 of a BRK instruction. Releasing RES one cycle later than for this full hijack case results in a different nonsensical address for the jumped-to opcode. Tipping the hat, the effective address is <<RES low>FD><RES low>. In this case, the 6502 has had the extra cycle needed to separately fetch the high byte of the jump vector (the cycle it was starved of in the other case), but fetched from <RES low>FD (like the opcode in the other case) instead of FFFD. The extra cycle also allows the low byte of the RES vector to appear in its proper position as the low byte of the opcode address.

Releasing RES two cycles later results in starting a new BRK instruction that jumps normally to the RES handler. This is because RES is held down long enough to prevent RES recognition stage 2 from being shut off by the already-running BRK instruction, and stage 2 still being alive will cause the next fetch cycle to start a new BRK instruction. The full description of what happens is covered by the worst-case RES invocation section under "Tolerances" in [6502 Interrupt Recognition Stages and Tolerances](index.php-title-6502_Interrupt_Recognition_Stages_and_Tolerances). Demonstrations of all cases are present there.

For RES half-hijacking, the RES line must be put down a full cycle later than for the full hijack case, strictly during T5. It must also be released a cycle later than for full, strictly during T6.

RES stage 2 will be recognized when T6 phase 2 is clocked in and control the fetch of the high byte of the RES vector after the lower-priority interrupt has already controlled the fetch of the low byte of its vector in T5. T0 will arrive separately (when it normally does) and the combined vector of <RES high><NMI or IRQ/BRK low> will send execution to a likely nonsensical location for code (but not as nonsensical as the results of the 2 failed full hijacks).

If RES is released later than strictly during T6, the same behavior happens as for the case of releasing two cycles later than T5 on full hijack: a new BRK instruction runs and jumps normally to the RES handler.

### BRK Protected from NMI vs. Unprotected from RES

So why is BRK protected against NMI half-hijack and not from RES? In the case of NMI, why is it left vulnerable to missing NMI entirely?

Hijacks (both successful half and failed full) by RES depend only upon RES coming up one to four transitions after recognition. In a healthy real system, it will not come up so soon, so real systems are safe from all RES hijacks. As a secondary matter, protecting BRK from RES hijack goes against the purpose of RES. It needs to be all-powerful to reliably redirect the processor to abandon its work in progress. Protection denies it the necessary power. Instead of complicating the design for a tiny niche of execution whose issue is already resolved another way, it is left simpler by being unprotected. Protection from RES is neither advisable nor necessary.

In contrast with RES, NMI half hijack is not dependent upon the NMI line coming up at all. A hijack will happen whether the surrounding system's NMI line comes up quickly or not, and the latter is normal hardware behavior. That demands protection. Without it, half-hijacks would inevitably occur during normal operation, making the system built around 6502s unreliable. The work-around of putting a JMP to NMI handling at the mixed address in ROM is just not acceptable. It would be like a fourth interrupt and would have to be documented.

The issue of potentially missing NMI (the side effect of protection) is prevented by the same real system behavior that prevents RES hijacking. A device exerting NMI will let the line up only after being answered by software action (a long time later).

It is only us hackers that can drive the simulator with transient interrupt signals that act like flaky hardware.

### Demonstrations

In all demonstrations, the NMI handler is set to F810, RES handler set to F920, and IRQ/BRK handler set to FA30.

[Just in time full hijack of an IRQ by NMI.](http://visual6502.org/JSSim/expert.html?graphics=f&steps=42&logmore=Execute,nmi,res,irq,~NMIG,RESP,IRQP,p2,INTG,RESG,State,tcstate,TState,Phi&a=0200&d=EAEA&a=20FD&d=4C20F9&a=F810&d=40&a=F910&d=4C20F9&a=F920&d=584C0002&a=F930&d=4C20F9&a=FA30&d=486840&a=FFFA&d=10F820F930FA&irq0=11&irq1=12&nmi0=21&nmi1=22)

The BRK instruction started by an IRQ is co-opted by an NMI and runs the NMI handler instead of the IRQ handler.

Schedule of the interrupts (Halfcycle numbers are 0-based):

```
Halfcycle 11 IRQ0 during T1 phase 2 of 1st NOP
Halfcycle 12 IRQ1 during T0 T2 phase 1 of 1st NOP
Halfcycle 21 NMI0 during T4 phase 2 of IRQ BRK
Halfcycle 22 NMI1 during T5 phase 1 of IRQ BRK
```

[Thwarted half-hijack of an IRQ by NMI](http://visual6502.org/JSSim/expert.html?graphics=f&steps=50&logmore=Execute,nmi,res,irq,~NMIG,RESP,IRQP,p2,INTG,RESG,State,tcstate,TState,Phi&a=0200&d=EAEA&a=20FD&d=4C20F9&a=F810&d=40&a=F910&d=4C20F9&a=F920&d=584C0002&a=F930&d=4C20F9&a=FA30&d=486840&a=FFFA&d=10F820F930FA&irq0=11&irq1=12&nmi0=23&nmi1=28) (and NMI line held long enough for the NMI to not be missed. See also "lost NMI" in [6502 Timing of Interrupt Handling](index.php-title-6502_Timing_of_Interrupt_Handling)).

The BRK instruction started by an IRQ is not successfully co-opted by a late NMI and runs the first instruction of the IRQ handler, then interrupted by the NMI and runs the NMI handler.

Schedule of the interrupts (Halfcycle numbers are 0-based):

```
Halfcycle 11 IRQ0 during T1 phase 2 of 1st NOP
Halfcycle 12 IRQ1 during T0 T2 phase 1 of 1st NOP
Halfcycle 23 NMI0 during T5 phase 2 of IRQ BRK
Halfcycle 28 NMI1 during T1 phase 1 of PHA (IRQ handler)
```

[Thwarted full hijack of an IRQ by RES.](http://visual6502.org/JSSim/expert.html?graphics=f&steps=34&logmore=Execute,nmi,res,irq,~NMIG,RESP,IRQP,p2,INTG,RESG,State,tcstate,TState,Phi&a=0200&d=EAEA&a=20FD&d=4C20F9&a=F810&d=40&a=F910&d=4C20F9&a=F920&d=584C0002&a=F930&d=4C20F9&a=FA30&d=486840&a=FFFA&d=10F820F930FA&irq0=11&irq1=12&reset0=21&reset1=22)

The BRK instruction started by an IRQ is co-opted by a RES that then ties its own shoes together and stumbles to an address of 20FD in the middle of nowhere. A JMP instruction placed there redirects to the RES handler anyway.

Schedule of the interrupts (Halfcycle numbers are 0-based):

```
Halfcycle 11 IRQ0 during T1 phase 2 of 1st NOP
Halfcycle 12 IRQ1 during T0 T2 phase 1 of 1st NOP
Halfcycle 21 RES0 during T4 phase 2 of IRQ BRK
Halfcycle 22 RES1 during T5 phase 1 of IRQ BRK
```

[Successful half-hijack of an IRQ by RES.](http://visual6502.org/JSSim/expert.html?graphics=f&steps=36&logmore=Execute,nmi,res,irq,~NMIG,RESP,IRQP,p2,INTG,RESG,State,tcstate,TState,Phi&a=0200&d=EAEA&a=20FD&d=4C20F9&a=F810&d=40&a=F910&d=4C20F9&a=F920&d=584C0002&a=F930&d=4C20F9&a=FA30&d=486840&a=FFFA&d=10F820F930FA&irq0=11&irq1=12&reset0=23&reset1=24)

The BRK instruction started by an IRQ is co-opted late by RES, mixing high and low bytes of the different vectors to a hybrid nonsense address of F930 (F9 of RES and 30 of IRQ). A JMP instruction placed there redirects to the RES handler.

Schedule of the interrupts (Halfcycle numbers are 0-based):

```
Halfcycle 11 IRQ0 during T1 phase 2 of 1st NOP
Halfcycle 12 IRQ1 during T0 T2 phase 1 of 1st NOP
Halfcycle 23 RES0 during T5 phase 2 of IRQ BRK
Halfcycle 24 RES1 during T6 phase 1 of IRQ BRK
```

[Successful half-hijack of an NMI by RES.](http://visual6502.org/JSSim/expert.html?graphics=f&steps=36&logmore=Execute,nmi,res,irq,~NMIG,RESP,IRQP,p2,INTG,RESG,State,tcstate,TState,Phi&a=0200&d=EAEA&a=20FD&d=4C20F9&a=F810&d=40&a=F910&d=4C20F9&a=F920&d=584C0002&a=F930&d=4C20F9&a=FA30&d=486840&a=FFFA&d=10F820F930FA&nmi0=11&nmi1=12&reset0=23&reset1=24)

The BRK instruction started by an NMI is co-opted late by RES, mixing high and low bytes of the different vectors to a hybrid nonsense address of F910 (F9 of RES and 10 of NMI). A JMP instruction placed there redirects to the RES handler.

Schedule of the interrupts (Halfcycle numbers are 0-based):

```
Halfcycle 11 NMI0 during T1 phase 2 of 1st NOP
Halfcycle 12 NMI1 during T0 T2 phase 1 of 1st NOP
Halfcycle 23 RES0 during T5 phase 2 of NMI BRK
Halfcycle 24 RES1 during T6 phase 1 of NMI BRK
```

Coding of the program used in all of the demonstrations:

```
;                  Interrupted user code
0200 NOP
0201 NOP           ; IRQ BRK or NMI BRK instead of running this instruction
;
;                  Intercept botched RES full-hijack of anything
20FD JMP F920      ; Jump to RES handler
;
;                  NMI handler
F810 RTI
;
;                  Intercept RES half-hijack of NMI
F910 JMP F920      ; Jump to RES handler
;
;                  RES handler (where visual6502 starts running code when finished starting up)
F920 CLI           ; Enable IRQs
F921 JMP 0200      ; Jump to user code
;
;                  Intercept RES half-hijack of IRQ/BRK
F930 JMP F920      ; Jump to RES handler
;
;                  IRQ/BRK handler
FA30 PHA           ; Save accumulator (something different than NMI's handler)
FA31 PLA           ; Pull back
FA32 RTI
```

### External References

"late NMI" and "lost NMI" in [6502 Timing of Interrupt Handling](index.php-title-6502_Timing_of_Interrupt_Handling)

[6502 Timing States](index.php-title-6502_Timing_States)

[6502 Interrupt Recognition Stages and Tolerances](index.php-title-6502_Interrupt_Recognition_Stages_and_Tolerances)

Retrieved from "[http://visual6502.org/wiki/index.php?title=6502\_Interrupt\_Hijacking](index.php-title-6502_Interrupt_Hijacking)"

