**INCOMPLETE DRAFT OF RECOVERED WIKI PAGE**

# 6502 BRK and B bit - VisualChips

## 6502 BRK and B bit

#### From VisualChips

The 6502 has 4 sources of interrupt-like behaviour: BRK, RESET, IRQ and NMI.

Much has been said about these - it's common to find confusion about the behaviour of the B bit in the pushed status word - and we can say a little more, with reference to our in-browser simulation of the NMOS 6502.

**Contents**

- [the B flag and the various mechanisms](#the-b-flag-and-the-various-mechanisms)
- [IRQ preceding a BRK instruction](#irq-preceding-a-brk-instruction)
- [late NMI will not half-modify vector reads](#late-nmi-will-not-halfmodify-vector-reads)
- [NMI preceding a BRK](#nmi-preceding-a-brk)
- [NMI masked by BRK](#nmi-masked-by-brk)
- [masking of the stack writes during RESET](#masking-of-the-stack-writes-during-reset)
- [Resources](#resources)

### the B flag and the various mechanisms

First technical point: the B flag position in the status register is not a bit in the status register: it is unaffected by PLP and RTI. However, the 6502 does push the register with either a 1 or 0. The intention is to distinguish a BRK from an IRQ, which is needed because these two share the same vector.  Brad Taylor says:

- software instructions BRK & PHP will push the B flag as being 1.
- hardware interrupts IRQ & NMI will push the B flag as being 0.

As it happens, there are bugs such that this description isn't strictly true in all situations, and the root cause is that the machinery for

- recording a pending hardware interrupt  (using a control signal called D1x1)
- forcing zero into the IR so the PLA performs the interrupt actions (uses D1x1, but at a different time to saving B)
- saving a value in the B position (distinguishing BRK/PHP from a pending hardware interrupt)
- forcing the appropriate values on the address bus to fetch the vector destination

are separate and independent.

(Note that the visual6502 sim reports the P register as if B was a storage element: in fact it is observing the node which conditionally drives the data bus during a push of P. See [here.](http://visual6502.org/JSSim/expert.html?nosim=t&find=p4&panx=431.8&pany=310.8&zoom=10.7)This node is the output of an inverter and is a doubly-inverted D1x1.)

### IRQ preceding a BRK instruction

(D1x1 was named by Balazs Beregnyei in his [giant schematic](http://www.downloads.reactivemicro.com/Public/Electronics/CPU/6502%20Schematic.pdf).  By all means refer to the schematic but note that it is a description of Rockwell's version of the 6502)

[Here's an URL](http://visual6502.org/JSSim/expert.html?graphics=f&a=0&d=58eaeaea&irq0=5&irq1=6&steps=36&loglevel=3&logmore=irq,D1x1,DPControl) which uses CLI and sets off a very short IRQ pulse.

You'll see that the D1x1 signal latches the pending interrupt, causes the pushed B to be zero, and is then cleared during the vector pull. This same signal is gated by 'Fetch' to produce 'ClearIR' (which jams zero into the IR)

Note also that the address pushed is for the instruction after the BRK. The BRK has masked the IRQ, because the IRQ handler will inspect the saved P and process the BRK.

### late NMI will not half-modify vector reads

This is a (necessary) feature: if an NMI occurs during the read of high and low vectors, it must not modify only the second read: modifying neither or both will determine whether the interrupt acts like an NMI or like the BRK/IRQ which is in progress.

We note it here because we can point out the mechanism: [transistor t970](http://visual6502.org/JSSim/expert.html?nosim=t&find=t970&panx=52.2&pany=123.3&zoom=12.4), where the logic is:

```
1368 := NMIP and not NMIL and not <VEC>.phi2
 NMIG := <1368>.phi1 or (<NMIG>.phi2 and not <brk-done>.phi1)
```

(Note that not all these signal names are presently known to the visual6502 netlist. Will be fixed.)

### NMI preceding a BRK

- [Here's an example](http://visual6502.org/JSSim/expert.html?graphics=f&a=fffa&d=4040&a=4040&d=40&a=0&d=58ea00eaea&nmi0=7&steps=36&loglevel=3&logmore=irq,nmi,res,D1x1) showing the RTI and resumption - note that the BRK has never been executed.

Because the B bit is stored as a 1, even though the NMI vector has been followed, in this case, the NMI handler could inspect the saved P register, in case a BRK was interrupted. It would then have to adjust the saved PC. This all takes time, and yet NMI is usually for rapid interrupt servicing.

As the NMI handler would not normally inspect P, this is a case of NMI masking BRK. If BRK is an OS call, it would not be made, and so you can't do that on a system using NMI.

### NMI masked by BRK

We thought we'd seen a late NMI during a BRK being ignored. Watch this space: we might have to retract that.

- [late NMI during BRK](http://visual6502.org/JSSim/expert.html?graphics=f&a=0&d=58ea00eaea&nmi0=17&steps=36&loglevel=3&logmore=irq,nmi,brk-done,D1x1,INTG,264,202,629,967,646,480)

### masking of the stack writes during RESET

This is a feature of the NMOS 6502 but not all other versions.

- [Here's](http://visual6502.org/JSSim/expert.html?graphics=f&a=0&d=58ea00eaea&reset0=4&reset1=8&steps=36&loglevel=3&logmore=irq,nmi,res,brk-done) a reset, showing that the 3 stack writes happen as reads.

The logic which causes these writes to be suppressed is as follows:

```
WR := op-T-mem-store
   or op-T2-php/pha
   or op-T4-brk
   or SD1 or SD2
   or (PCH/DB) or (PCL/DB)
 (R/#W) := not (<WR>.phi2 and not RESG and RDY)
 R/#W := <(R/#W)>.phi1
```

with the writes during RESET suppressed by [transistor t3455](http://visual6502.org/JSSim/expert.html?nosim=t&find=t3455&panx=392.1&pany=199.6&zoom=12.4)

### Resources

- back to parent page [6502Observations](index.php-title-6502Observations.md)
- [Internals of BRK/IRQ/NMI/RESET on a MOS 6502](http://www.pagetable.com/?p=410) by Michael Steil on pagetable.com
- [Interrupts in 65xx processors](http://en.wikipedia.org/wiki/Interrupts_in_65xx_processors) (wikipedia)
- [The B flag](http://nesdev.parodius.com/the%20%27B%27%20flag%20&%20BRK%20instruction.txt) by Brad Taylor
- [B flag discussion](http://forum.6502.org/viewtopic.php?p=13036#13036) on 6502.org
- [Investigating Interrupts](http://www.6502.org/tutorials/interrupts.html) tutorial by Garth Wilson
- [Register Preservation Using The Stack (and a BRK handler)](http://www.6502.org/tutorials/register_preservation.html) tutorial by Bruce Clark
- IRQ handler in the [BBC micro OS1.20](http://mdfs.net/Docs/Comp/BBC/OS1-20/DC1C)
- [CPU Interrupts](http://wiki.nesdev.com/w/index.php/CPU_interrupts) on Nesdev wiki

Retrieved from "[http://visual6502.org/wiki/index.php?title=6502\_BRK\_and\_B\_bit](index.php-title-6502_BRK_and_B_bit.md)"

