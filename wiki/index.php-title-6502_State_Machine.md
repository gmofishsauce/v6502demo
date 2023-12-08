**Recovered visual6502.org wiki - beta release**

# 6502 State Machine - VisualChips

## 6502 State Machine

#### From VisualChips

This article contributed by Segher Boessenkool

There are many timing bits/flags/states in the 6502.

The ones I will discuss here are all valid for whole cycles,
i.e. a phi1 followed by a phi2.  There of course are others
that are latched on phi2, to link everything together, but
we'll ignore those.

One thing that complicates matters is the RDY input, with
which you can stall the CPU for as many (non-write) cycles
as you want.  RDY is not handled in a uniform way at all.

Another complication is the conditional branch instructions,
which take only two cycles if not taken, three if no page
crossing, and four otherwise; to get such efficiency for
these important instructions, evil shortcuts are taken.

And there is an assortment of optimisations and implementation
tricks that we'll see when we get there.

Okay, so what states are there?  The PLA has T0,T1,T2,T3,T4,T5
as inputs.  There is another signal that is more like a "real"
T1 as we'll see, which I'll call T1F.  This is output on the
SYNC pin.  There is also what I call T01 (node 1357 / not node
223 in jssim), but we can optimise that away as we'll see.

The BRK instruction needs a T6; it gets a VEC1 instead (and a
VEC0).

Some load instructions have the last cycle skipped if there is
no page crossing in its addressing; the corresponding store
instructions don't do that.  And the read-modify-write (RMW)
instructions have two extra cycles tucked on: SD1 and SD2
(store instructions store in the last cycle; RMW instructions
read in the last "normal" cycle, and store in both SD1 and SD2).

Every instructions starts with the fetching of its first two
bytes (the PC isn't incremented for the second fetch for most
one-byte opcodes; exceptions are BRK RTI RTS).  During this
time, the PLA still decodes the previous instruction, and most
of the datapath still handles that as well (even the cycle
after this, the datapath is a bit behind).  So the T0 and T1
inputs to the PLA actually come behind everything else; e.g.
you have T2-T3-T0-T1 for a four cycle normal instruction, or
T2-T3-T4-SD1-SD2-T0-T1 for a seven cycle RMW instruction.
None of the decoding (except the predecode) sees the instruction
during the two cycles that come before this; instead, T01 and
T1F are used.

Here's a state diagram to scare you.  "T01,T0" means both
T01 and T0 are active; horizontal arrows are transitions
when not RDY; vertical arrows are transitions when RDY.
If no arrow for not RDY is given, it means the state is kept.

```
T01,T0  ------------------------->  T0
             |                                |
             |                                |
             v                                v
        T01,T1F,T1  ---->  T01,T1F  <----  T1F,T1
             |                |               |
             +-------------+  |  +------------+
                           |  |  |
                           v  v  v
                              T2
     (or T01,T0 if the next insn is a twocycle insn)
```

This does not show the various entries into this state diagram.
The usual entry point is T01,T0.  The exceptions are non-taken
branches (which do T2 - T01,T1F) and taken branches that do not
cross page (T2 - T3 - T01,T1F).  RESET will behave quite oddly,
but will end up at T01,T0 after a few cycles clocking (with RDY
asserted).

Let's simplify this.  We know that this state diagram is always
entered somewhere with T01 asserted.  T01 does two things: it
clears T2,T3,T4,T5, and it asserts T0 if T1F isn't already
asserted.  So we do not actually need to consider it, since we
know it is always on when the state diagram is entered, it won't
do anything more that we don't have in there already.  So we get:

```
T0
             |
             |
             v
          T1F,T1  -------->  T1F
             |                |
             +-------------+  |
                           |  |
                           v  v
                            T2
     (or T01,T0 if the next insn is a twocycle insn)
```

This shows that T1F is more the "real" T1.  T1 is only asserted
for one cycle, and then forgotten, with no successor, even when
not RDY.  It is the last cycle for normal instructions, and it
cannot do anything external anymore anyway (the bus is used to
play fetch), and it doesn't do anything spectacularly interesting
(only register writeback, in fact), so it can just as well do
it at once and forget about it.

The instructions with bitpatterns xxx010x1, 1xx000x0 and
xxxx10x0 except 0xx01000 are two cycle instructions: resp.
A immediate, X/Y immediate, and all the rest (push/pull is
the excluded pattern).  These instructions do not switch to
T2 but immediately back to T0.

Push instructions (0x001000) switch to T0 after T2; pull
instructions (0x101000) after T3; direct jump (01001100)
after T2; indirect jump (01101100) after T4; jsr (00100000)
after T5; rts and rti (01x00000) after T5; and brk (including
all other interrupts) (00000000) after VEC1 ("T6").

Conditional branch instructions switch to T0 after T3.
But they can be shortcut, we'll get to that in a minute.

Everything else is a "memory" instruction, which can have
various addressing modes.  "Immediate" is already handled
in the two cycle stuff.  For the rest we define a "Tmem"
that is one of T2 to T5, depending on addressing mode:

```
-- T2 for xxx001xx, zero page
-- T3 for xxx011xx, zero page indexed
-- T3 for xxx101xx, absolute
-- T4 for xxx11xxx, absolute indexed
-- T5 for xxxx00xx, indirect indexed / indexed indirect
```

For an instruction that does not store to memory, Tmem is
followed by T0.  Except, the absolute indexed and indexed
indirect start T0 a cycle earlier if there was no page
crossing.  In Tmem the CPU issued the read to the system
bus; in T0 and T1 it will do whatever it needs to do with
it for the current instruction.  So T2,T3,etc. are for the
memory addressing part, and T0,T1 are for the actual operation.

For pure stores, it works the same, except there is never
a shortcut.

For RMW instructions, after Tmem there are SD1 and SD2 cycles,
and after that T0 and T1 as usual.  During Tmem the original
value is read; during SD1 it is written back, and the modified
result computed; during SD2 the result is written.

For the "brk" instruction, there is this extra VEC1 cycle.
VEC0 is active when T5 and RDY and the current instruction
is a brk instruction; during VEC0, the low byte of the address
of the interrupt routine is read.  VEC0 is immediately followed
by VEC1, and then (surprise!) the high byte is read.  The cycle
after VEC1 various interrupt bookkeeping tasks are done (the
IRQ/NMI/RES request flop is cleared, the I (interrupt prohibit)
bit is set, that kind of thing).

All interrupts are implemented by forcing a brk instruction
into the instruction stream.  This is done by clearing all bits
in the predecode reg during T1F, if the previous cycle was T0,
or it was T2 and the instruction is a branch (and of course an
interrupt is pending).  (All bits in the predecode reg are
cleared whenever T1F is deasserted: that way, neither the
two-cycle or one-byte signals will trigger at the wrong time!)

So, branches.  Our journey is almost at an end!

We have seen a branch will take four cycles by the normal
mechanism.  If the branch stays within the current page, this
can be cut to three cycles; and if the branch is not taken,
it can be cut to only two.

The timing diagram for branches in the MOS hardware manual
is wrong (it's the last half page of this ~180 pages excellent
manual).  The correct sequence is:

```
Tn  address bus     data bus         comments
--------------------------------------------------------
T0  PC              branch opcode    fetch opcode
T1  PC + 1          offset           fetch offset
T2  PC + 2          next opcode      fetch for branch not taken
T3  PC + 2 + off    next opcode      fetch for branch taken, same page
    (w/o carry)
T0  PC + 2 + off    next opcode      fetch for branch taken, other page
    (with carry)

(T3/T4 or just T4 are left away if branch not taken or no
page crossing).
```

But that's not quite the whole story: the next instruction
will start at T1, not T0: it has its first (opcode) byte
fetched already.  For either of the short versions, only T1F
will be active; only the four-cycle branch has T1 as well.
(In all cases T01 will be active, but you're supposed to
have forgotten about that by now :-) )

So for the three cases, you get respectively:

```
[ fetch our two bytes, T0/T1, yadda yadda ]
T2      PC + 2                                      next opcode
T1F     PC + 3                                      2nd byte of next
```
```
[ fetch our two bytes, T0/T1, yadda yadda ]
T2      PC + 2                                      useless read
T3      PC + 2 + off (w/o carry)                    next opcode
T1F     PC + 3 + off                                2nd byte of next
```
```
[ fetch our two bytes, T0/T1, yadda yadda ]
T2      PC + 2                                      useless read
T3      PC + 2 + off (w/o carry)                    useless read
T0      PC + 2 + off (with carry)                   next opcode
T1F,T1  PC + 2 + off                                2nd byte of next
```

For the "branch taken, same page" case there is an oddity
with interrupts.  In this case, T1F is preceded by T3 (not
T0 or T2), so no interrupt can happen on the next instruction!
You can mask NMIs this way even (but not reset, it messes up
the timing directly).

Retrieved from "[http://visual6502.org/wiki/index.php?title=6502\_State\_Machine](index.php-title-6502_State_Machine.md)"

