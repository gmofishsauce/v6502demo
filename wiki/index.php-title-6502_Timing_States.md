**Recovered visual6502.org wiki - beta release**

# 6502 Timing States - VisualChips

## 6502 Timing States

#### From VisualChips
**Contents**

- [Introduction](#introduction)
- [Time States Seen Around Instruction Execution](#time-states-seen-around-instruction-execution)
- [Branch Instructions Timing States](#branch-instructions-timing-states)
- [BRK Instruction Timing States](#brk-instruction-timing-states)
- [The RMW Instructions' Timing States](#the-rmw-instructions-timing-states)
- [Forever Instructions](#forever-instructions)
- [Time States That Do Not Occur During Normal Instruction Execution](#time-states-that-do-not-occur-during-normal-instruction-execution)
- [Demonstration of All Time States](#demonstration-of-all-time-states)
- [External References](#external-references)

### Introduction

There are two things that are critical for correct instruction execution in the 6502 (indeed, for any complex CPU chip): the pattern of bits in the instruction register AND the pattern of "time code" bits from the timing control block of circuits.

Both sets of bits (IR and time code), in combination, control the output bits of the Programmable Logic Array (PLA) block of circuits. PLA outputs, in turn, affect the Random Control Logic (RCL) block of circuits, and their control outputs directly operate the connections among the registers, busses, and ALU on the other end of the chip die to actually get the instructions' work done.

In some situations, time code and PLA interaction merely starts a parallel chain of timing signals through the RCL part of the chip network that manage some actions themselves, instead of burdening the PLA with controlling them. The full state of timing control presented here includes such parallel processes.

The, "some situations", mentioned are for the Read-Modify-Write (RMW) instructions.

With all the combinations of parallel independent control and PLA-based control, the 6502's timing control has 24 states. A total of eleven nodes ("bits") comprise these states. Six of them are the time code applied to the PLA, the remaining five are part of the non-PLA control. The latter five are further divided up into three and two: three are internal-state members of timing generation, and the last two are in the RCL block.

Labeling of the time states derives from the "TIMING GENERATION LOGIC" block and the "CONTROL FLIP-FLOPS" sub-block inside the "RANDOM CONTROL LOGIC" block in [Dr. Donald Hanson's block diagram](http://www.witwright.com/DonPub/6502-Block-Diagram.pdf), and from discussion of timing states in [6502 State Machine](index.php-title-6502_State_Machine.md).

The notation developed for trace/debug output, and the notation presented hereafter in this document, uses eight fields. It lists the six explicit (PLA controlling) output nodes in numeric order first, followed by square brackets around the non-explicit internal state of timing generation (three nodes), followed by the state of the two RCL nodes.

Wherever one of the explicit nodes is inactive, a blank placeholder of ".." is present for it. Similar logic also applies to the last two fields: if *all* the nodes that they correspond to are inactive, the blank placeholder appears in the respective field (".." within square brackets for the seventh field, "..." for the eighth field). Only one node at a time is active for each of the last two fields.

The six PLA controlling nodes are T0, T+, T2, T3, T4, and T5. The second one, called T+, corresponds to T1X in the Hanson diagram. It is presented as T+ instead to correspond to the names of the PLA nodes that it affects in the visual6502 project. The PLA controllers are active when their logic states are low.

The other five nodes (squeezed into fields seven and eight) are active when their logic states are high. The seventh field, in square brackets, shows three states (due to three nodes) called T1, V0 (short for VEC0), and T6 (a synonym
for VEC1). The eighth field shows two states, due to two nodes, called SD1 and SD2.

When T1 is displayed inside the seventh (bracketed) field, the external SYNC pin is also being driven high (by the node tested for T1) to indicate that the current memory read operation is for an instruction opcode.

Here are all the combinations of states (labelled, but otherwise without context).

```
T0 .. T2 .. .. .. [..] ...    ; 1st cycle of 2 cycle opcodes
.. .. T2 .. .. .. [..] ...    ; 1st cycle of 3+ cycle opcodes
.. .. .. T3 .. .. [..] ...    ; 2nd cycle of 4+ cycle opcodes
.. .. .. T3 .. .. [..] SD1    ; 2nd cycle of RMW zp
.. .. .. .. T4 .. [..] ...    ; 3rd cycle of 5+ cycle opcodes
.. .. .. .. T4 .. [..] SD1    ; 3rd cycle of RMW zp,X & abs
.. .. .. .. T4 .. [..] SD2    ; 3rd cycle of RMW zp
.. .. .. .. .. T5 [..] ...    ; 4th cycle of 6+ cycle opcodes
.. .. .. .. .. T5 [..] SD1    ; 4th cycle of RMW abs,X & abs,Y (latter illegal)
.. .. .. .. .. T5 [..] SD2    ; 4th cycle of RMW zp,X & abs
.. .. .. .. .. T5 [V0] ...    ; 4th cycle of BRK
.. .. .. .. .. .. [T6] ...    ; 5th cycle of BRK
.. .. .. .. .. .. [..] SD1    ; 5th cycle of RMW (zp,X) & (zp),Y (illegal)
.. .. .. .. .. .. [..] SD2    ; 5th of RMW abs,X&Y, 6th of RMW (zp,X) & (zp),Y
.. .. .. .. .. .. [..] ...    ; Terminal state of opcodes that run forever
T0 .. .. .. .. .. [..] ...    ; 2nd-to-last cycle of all(*) opcodes
.. T+ .. .. .. .. [T1] ...    ; Last cycle of all(*) opcodes
.. .. .. .. .. .. [T1] ...    ; Last cycle of branches not taken & no page cross
T0 T+ .. .. .. .. [..] ...    ; Clock long-term hold by RES
T0 .. .. .. .. .. [T6] ...    ; 1 cycle after RES of BRK during T5 [V0] cycle
T0 .. .. .. .. .. [..] SD1    ; 1 cycle after RES of RMW during SD cycle
T0 .. .. .. .. .. [..] SD2    ; 1 cycle after RES of RMW just after SD cycle
T0 T+ .. .. .. .. [..] SD2    ; 2 cycles after RES (still held) of RMW during SD cycle
.. T+ .. .. .. .. [T1] SD2    ; 2 cycles after RES (released) of RMW during SD cycle

* All opcodes except branches not taken & branches taken with no page cross
```

As seen above, there are two states where three nodes are active at the same time (three-hot), twelve states where two nodes are active at the same time (two-hot), nine that are one-hot, and a single none-hot state.

The PLA subset itself has two two-hot states, six one-hot states, and one none-hot.

The low-profile "blank" notation of ".." assists visual examination of trace/debug output by keeping consistent placeholders for nodes when they are inactive, with minimized visual clutter. Aligning everything in fixed positions
contributes to rapid recognition of changes.

Eight fields is a compromise between a full eleven fields of one node each, and a minimum of three fields to show states that have three nodes active at the same time. This format shows the segregation of timing's external PLA-affecting nodes from its internal nodes and from the RCL nodes. An alternative compact compromise could be to reserve two fields for the external nodes (for its two-hot states) and one field each for the internal and RCL nodes.

### Time States Seen Around Instruction Execution

During normal operation (when the clock is not under the influence of the RES interrupt line being active), only 18 of the 24 possible states actually arises.

The convention for presenting time states for instruction execution here is a a little different from that supplied in the usual programming literature in two ways: by time numbering and by first through last cycles listed.

The time numbering issue applies to matching up the time codes used in Appendix A of, ["MCS6500 Microcomputer Family Hardware Manual"](http://archive.6502.org/books/mcs6500_family_hardware_manual.pdf), with the time states documented here. "T0" in the manual matches with the states that have [T1] in them here (most often T+ [T1]). The rest of the time codes in the hardware manual listings match up with those here after being incremented by one. The second-to-last time code in each hardware manual listing corresponds to the T0-containing states here, and the last code in each listing is the [T1] state again.

The convention here for "first" and "last" cycle time states is focused upon when each opcode is actually occupying the instruction register (IR). The emphasis is upon the span of states in which control signals are initiated, and not upon the span over which operations are completed. The typical narrative has the opcode fetch cycle as the first one, listing [T1] through T0. Here, all instructions are listed as beginning with T2 in their time state and ending with [T1] in their time state. Although [T1] reads the opcode from memory, the instruction register (IR) is only set to the new opcode during the first half of T2 (T2 phase 1). That is the first moment of actual execution of the opcode, when it can start sending signals by affecting the states of PLA nodes and the RCL nodes further across the chip. The opcode remains undisturbed inside the IR all the way to the end of the next [T1] clock state ([T1] phase 2). This allows an instruction to do its last signal origination even during the fetching of its successor instruction by other circuits on the chip. These propagated last signals can perform the final operations of an instruction even one cycle later (T2 again) when the next instruction is in the IR. Not all instructions have any work to do by [T1], so it is sometimes an idle cycle in terms of instruction work. The strictly two-cycle instructions are always busy during [T1]. [T1] is always busy for fetching instructions.

After having gone through so much justification to start listing instruction timing states with T2, it must also be noted that there is some opcode identification that happens in the second half of [T1]. This is the predecode register's function that identifies two-cycle instructions and one-byte instructions by binary digital interference pattern. Two-cycle identification determines which of two possible T2 states will start the instruction execution. One-byte instruction identification determines whether the program counter (PC) is auto-incremented after reading the byte after the opcode during T2 (auto-increment is inhibited for one-byte instructions).

Time states with T0 in them have a special significance. All instructions must, at some late point in their execution, generate a signal that propagates around to set the clock to T0. That initiates the fetch for the next instruction at [T1] one cycle later. Instructions that don't do this run forever (but with two exceptions). There are twelve opcodes in the baseline 6502 that have this problem, and they can only be terminated by the RES interrupt. RES is effective because it sets the clock to T0 on its own, among other effects.

By contrast, the IRQ and NMI interrupts depend upon the currently running instruction to set the clock to T0 the normal way: the T0 state lets a pending IRQ or NMI signal propagate to where it causes a BRK instruction to be started at the next T2 state, beginning the response to the interrupt. This is the implementation of the IRQ/NMI feature that allows the currently running instruction to finish before starting the interrupt response. This is also its vulnerability to the "forever" instructions that don't set the clock back to T0: IRQ and NMI interrupts are locked out of having any effect while the non-terminating instruction runs. It is correct behavior that goes on for too long.

Generically, instructions run through time states of

```
.. .. T2 .. .. .. [..] ...
.. .. .. T3 .. .. [..] ...
.. .. .. .. T4 .. [..] ...
.. .. .. .. .. T5 [..] ...
T0 .. .. .. .. .. [..] ...
.. T+ .. .. .. .. [T1] ...
```

The above example is for an instruction that runs for six cycles. Instructions that run for fewer cycles have T0 and [T1] occur sooner:

```
Five cycles          |         Four cycles          |         Three cycles
.. .. T2 .. .. .. [..] ...  |  .. .. T2 .. .. .. [..] ...  |  .. .. T2 .. .. .. [..] ...
.. .. .. T3 .. .. [..] ...  |  .. .. .. T3 .. .. [..] ...  |  T0 .. .. .. .. .. [..] ...
.. .. .. .. T4 .. [..] ...  |  T0 .. .. .. .. .. [..] ...  |  .. T+ .. .. .. .. [T1] ...
T0 .. .. .. .. .. [..] ...  |  .. T+ .. .. .. .. [T1] ...  |
.. T+ .. .. .. .. [T1] ...  |                              |
```

All of the above examples also apply to instructions that have variable execution times due to using indexed addressing modes: requiring one more cycle when a page crossing is required to access the correct memory address. Variable execution times for memory access applies only to instructions that only perform a read of memory.

The instructions that run in only two cycles have T0 and T2 combined into the same cycle. This is courtesy of the effects of predecode upon the clock, where it identifies two-cycle instructions.

```
T0 .. T2 .. .. .. [..] ...
.. T+ .. .. .. .. [T1] ...
```

### Branch Instructions Timing States

The mentioned two exceptions to requiring instructions to set the clock to T0 are the conditional branch instructions. For review, branch instruction execution has three cases of behavior:

```
* Branch not taken. Runs for two cycles.
* Branch taken without crossing a memory page. Runs for three cycles.
* Branch taken and it crosses to another memory page. Runs for four cycles.
```

The first two cases do not set the clock to T0. Instead, they bypass T0, going directly to [T1] for opcode fetch. The third case acts like a normal instruction.

The [T1] state caused by T0-bypassing is qualitatively different:

```
.. .. .. .. .. .. [T1] ...
```

The T+ (aka T1X) node is inactive. It only appears if the previous clock state had T0 active.

Comparison of all the branch cases:

```
No branch           |       Branch, no cross       |      Branch with cross
.. .. T2 .. .. .. [..] ...  |  .. .. T2 .. .. .. [..] ...  |  .. .. T2 .. .. .. [..] ...
.. .. .. .. .. .. [T1] ...  |  .. .. .. T3 .. .. [..] ...  |  .. .. .. T3 .. .. [..] ...
                            |  .. .. .. .. .. .. [T1] ...  |  T0 .. .. .. .. .. [..] ...
                            |                              |  .. T+ .. .. .. .. [T1] ...
```

There is an issue with the lack of a T0 state for the first two branch instruction cases. NMI and IRQ depend upon T0 happening, so they would be locked out for the brief extra time it takes to get to a non-branch instruction (or a branch instruction of the third case). They would be completely locked out by an infinite loop of non-crossing branch instructions.

This problem is remediated for all the branch instructions by allowing the T2 state to also let a pending NMI or IRQ signal propagate further to cause the interrupt response, just like T0. The T2 window is ONLY available for the
branch instructions. This window for a branch catching IRQ and NMI for the first two cases is admittedly a little narrow (2 cycles wide: [T1] followed by T2), so if it is missed, one will still have to wait until the end of the instruction after the branch (or T2 if that next instruction is a branch) to catch them. The saving grace is the prevention of IRQ and NMI lock-out by the infinite loop scenario.

In all branch cases where IRQ and NMI are caught early at T2, the branch instruction does finish. Confirming the reader's suspicions, the page-crossing case gives IRQ and NMI TWO opportunities to be caught and processed.

### BRK Instruction Timing States

The BRK instruction has a special extension of the clock for itself. The extensions originate the control signals needed to clear the interrupt state(s) of the 6502. BRK is the instruction for the processor's response to an interrupt. It is also available as a readable opcode, providing a software interrupt for debugging. It is helpful, and most general, to think of there being four ways to run a BRK instruction: pulling the RES line, pulling the NMI line, pulling the IRQ line, and reading a BRK opcode. The internal difference between the interrupt line invocation and the opcode read invocation occurs when clock state T2 phase 1 is clocked in. For the interrupt cases, the IR is cleared to all zero bits (set to the BRK opcode) instead of being set to whatever was read from memory during [T1]. For the opcode case, the same thing happens as with all other opcodes: IR is set to the bitwise NOT of the predecode register (the predecode register is itself the bitwise NOT of what was read from memory).

BRK has a different T5 state from all other instructions, and a unique T6 state of its own.

```
.. .. T2 .. .. .. [..] ...
.. .. .. T3 .. .. [..] ...
.. .. .. .. T4 .. [..] ...
.. .. .. .. .. T5 [V0] ...
.. .. .. .. .. .. [T6] ...
T0 .. .. .. .. .. [..] ...
.. T+ .. .. .. .. [T1] ...
```

The "V0" in the T5 state is short for VEC0. It is a node that is activated by a PLA node when the clock enters the T5 state. The signal carried independently by VEC0 propagates one cycle later to a node called VEC1: this is the T6 state. The signal continues back to the main clock one cycle later to set it to T0.

### The RMW Instructions' Timing States

These are the remaining normal operation timing states: the ones that use the independent timing signal paths in the RCL. Nodes in the PLA initiate the SD1 and SD2 states, and that signal propagation ends with setting the clock to T0.

The alignment of SD1 and SD2 with the clock states shown so far depends upon the addressing mode used by the RMWs. There are six RMW instructions: ASL, DEC, INC, LSR, ROL, and ROR. They use 4 addressing modes that reach into memory.

```
Zero Page           |    ZP Indexed & Absolute     |       Absolute Indexed
.. .. T2 .. .. .. [..] ...  |  .. .. T2 .. .. .. [..] ...  |  .. .. T2 .. .. .. [..] ...
.. .. .. T3 .. .. [..] SD1  |  .. .. .. T3 .. .. [..] ...  |  .. .. .. T3 .. .. [..] ...
.. .. .. .. T4 .. [..] SD2  |  .. .. .. .. T4 .. [..] SD1  |  .. .. .. .. T4 .. [..] ...
T0 .. .. .. .. .. [..] ...  |  .. .. .. .. .. T5 [..] SD2  |  .. .. .. .. .. T5 [..] SD1
.. T+ .. .. .. .. [T1] ...  |  T0 .. .. .. .. .. [..] ...  |  .. .. .. .. .. .. [..] SD2
                            |  .. T+ .. .. .. .. [T1] ...  |  T0 .. .. .. .. .. [..] ...
                            |                              |  .. T+ .. .. .. .. [T1] ...
```

In all of the examples above, the time state before the SD1 state is the one that initiates the separate SD1 SD2 signal chain. For zero page addressing mode, the signal is initiated at T2 by a PLA node, and it clock phase propagates to the SD1 node by one cycle later. Contrast this against V0 being activated immediately at T5 by a PLA node for the BRK instruction.

There are still more instructions that use SD1 and SD2. They happen to be RMWs among the illegal/undocumented opcodes. They use all of the addressing modes above (plus Absolute Indexed Y that the official RMWs do not use). They also use the two indirect indexed addressing modes, (zp,X) and (zp),Y. That takes eight cycles, and those addressing modes initiate the signal chain at T5. The mnemonics for the opcodes, according to [6502 all 256 Opcodes](index.php-title-6502_all_256_Opcodes.md) are SLO, RLA, SRE, RRA, DCP, and ISC.

```
.. .. T2 .. .. .. [..] ...
.. .. .. T3 .. .. [..] ...
.. .. .. .. T4 .. [..] ...
.. .. .. .. .. T5 [..] ...
.. .. .. .. .. .. [..] SD1  <-- Unique to illegal RMW (zp,X) and (zp),Y
.. .. .. .. .. .. [..] SD2
T0 .. .. .. .. .. [..] ...
.. T+ .. .. .. .. [T1] ...
```

The unique timing state added to the collection that doesn't happen with any of the official instructions is the all blank state with SD1.

### Forever Instructions

There's one more time state from normal operation to be introduced.

```
.. .. T2 .. .. .. [..] ...
.. .. .. T3 .. .. [..] ...
.. .. .. .. T4 .. [..] ...
.. .. .. .. .. T5 [..] ...
.. .. .. .. .. .. [..] ...
.. .. .. .. .. .. [..] ...  <-- etc, never changes
```

The above is from any of the twelve illegal opcodes that runs forever. The all blank time state shows that there is no signal in progress that will set the clock back to T0 to initiate the fetch of a new instruction.

### Time States That Do Not Occur During Normal Instruction Execution
```
T0 .. .. .. .. .. [T6] ...    ; 1 cycle after RES of BRK during T5 [V0] cycle
T0 T+ .. .. .. .. [..] ...    ; Clock long-term hold by RES
T0 .. .. .. .. .. [..] SD1    ; 1 cycle after RES of RMW during SD cycle
.. T+ .. .. .. .. [T1] SD2    ; 2 cycles after RES (released) of RMW during SD cycle
T0 T+ .. .. .. .. [..] SD2    ; 2 cycles after RES (still held) of RMW during SD cycle
T0 .. .. .. .. .. [..] SD2    ; 1 cycle after RES of RMW just after SD cycle
```

These are all due to the RES interrupt aborting various instructions at specific times during their operation. Some background on when RES is sensed and its short term effects on the clock is necessary.

The RES interrupt line being down (asserted) is first sensed when clock phase 2 is changed to clock phase 1 (when phase 1 is "clocked in"). If RES is asserted after the processor is already in phase 1, it will not be sensed until the next time that phase 1 is clocked in.

The clock cycle that begins with this phase 2 to phase 1 transition shall be called the "sense cycle". Its clock time state is not affected. The time state for the next clock cycle IS affected: it is changed to T0 by RES having been sensed.

```
Tn-1 Phase 1
    Tn-1 Phase 2

RES asserted (low)
    Tn Phase 1              <==+
RES de-asserted (high)         |- RES sense cycle
    Tn Phase 2              <==+

    T0 Phase 1              <==\\_ T0 cycle caused by RES
    T0 Phase 2              <==/
```

The newly invented term, "sense cycle", shall be used shortly.

The time state:

```
T0 .. .. .. .. .. [T6] ...
```

is caused by RES interrupting a BRK instruction at a sense cycle of T5 [V0].

The time state:

```
T0 T+ .. .. .. .. [..] ...
```

arises when RES is down when a T0 phase 1 clock state is clocked in. This can be either the T0 that is usually scheduled by an instruction's imminent termination, or the T0 caused by instruction abort shown above. Holding down RES long enough causes the T0 T+ state to arise.

```
Instruction's own T0   |     T0 caused by RES
--------------------------|--------------------------
      Tn Phase 2          |      Tn-1 Phase 2
                          |
  RES asserted (low)      |  RES asserted (low)
      T0 Phase 1          |      Tn Phase 1
  RES de-asserted (high)  |      Tn Phase 2
      T0 Phase 2          |
                          |      T0 Phase 1
      T0 T+ Phase 1       |  RES de-asserted (high)
      ...etc.             |      T0 Phase 2
                          |
                          |      T0 T+ Phase 1
                          |      ...etc.
```

The sense cycle for T0 T+ can be any clock time for all instructions other than the RMWs. For the RMW instructions, the sense cycle must not be its SD chain initiation cycle. That cycle is T2 for zero page mode, T3 for zero page indexed and absolute (unindexed) modes, T4 for absolute indexed modes (X and Y), and T5 for the indirect and indexed modes (illegal RMWs).

The rest of the time states occur only with the RMW instructions.

The time state:

```
T0 .. .. .. .. .. [..] SD1
```

occurs if the sense cycle for RES is the SD chain initiation cycle of an RMW. If RES rises back up before this (T0) cycle's phase 1 is clocked in, then the next time state will be:

```
.. T+ .. .. .. .. [T1] SD2
```

If, however, RES stays down until after the T0 SD1 cycle's phase 1 is clocked in, then the next time state will instead be:

```
T0 T+ .. .. .. .. [..] SD2
```

This is exactly the same situation as the T0 T+ state without SD2: the SD signal just happens to be propagating independently in parallel with T0 T+.

An extra phenomenon to note with T0 SD1 followed by T+ [T1] SD2: the T0 SD1 state would have caused the opcode read state (T+ [T1] SD2) to be in write mode (!) were it not for RES having already forced the 6502 to be in read-only mode. Elaborating further: the 6502 will have been in read-only mode since during the T0 SD1 cycle, which is the cycle after the sense cycle. The effect of SD1 commanding a write operation shows up at the external read-write pin one cycle later, when we are in the T+ [T1] SD2 state, but RES has already overridden that to be a read (and maintaining consistency with SYNC being high to indicate opcode read).

The final time state to present,

```
T0 .. .. .. .. .. [..] SD2
```

occurs if the sense cycle for RES is the SD1 cycle for an RMW (the cycle after the SD chain initiating cycle).

### Demonstration of All Time States

[This link](http://visual6502.org/JSSim/expert.html?graphics=f&steps=292&logmore=Execute,res,State,tcstate,TState,Phi&r=2EC&a=2EC&d=A90B8DFCFFA9038DFDFFA10018B00090009001FFE600EE0000FE0000130002A9128DFCFFE600A9198DFCFFE600A9208DFCFFE600A9278DFCFF0001EA&reset0=123&reset1=126&reset0=157&reset1=158&reset0=193&reset1=196&reset0=229&reset1=230&reset0=267&reset1=272) runs the expert version of visual6502 with a minimal 6502 program, including five RES interrupts, that causes the appearance of all the time code states documented above.

The human readable coding of the program:

```
;                  Test 1
02EC LDA #0B       ; Change RES vector to 030B for test 2
02EE STA FFFC
02F1 LDA #03
02F3 STA FFFD
02F6 LDA (00,X)    ; Show all common T states (a 6 cycle instruction)
02F8 CLC           ; Show T0 T2 (a 2-cycle instruction)
02F9 BCS 02FB      ; Show [T1] at end (AKA next opcode fetch), after this one's T2 (no branch)
02FB BCC 02FD      ; Show [T1] at end (AKA next opcode fetch), after this one's T3 (branch, no page cross)
02FD BCC 0300      ; Show branch with page cross
02FF FF            ; FF padding
0300 INC 00        ; Show T3 SD1 and T4 SD2
0302 INC 0000      ; Show T4 SD1 and T5 SD2
0305 INC 0000,X    ; Show T5 SD1 and blank SD2
0308 SLO (00),Y    ; Show blank SD1 and blank SD2
030A KIL           ; Show all blank
Phase 123, RES0    ; RES down during second T blank Phase 2
Phase 126, RES1    ; RES up after clocking in T0 Phase 1
                   ; Shows T0 T+ state before invoking RES BRK
                   ; The RES invoked BRK shall also show the T5 [V0] and the [T6] states
;
;                  Test 2
030B LDA #12       ; Change RES vector to 0312 for test 3
030D STA FFFC
0310 INC 00
Phase 157, RES0    ; RES down during T+ [T1] Phase 2
Phase 158, RES1    ; RES up after clocking in T2 Phase 1
                   ; Shows T0 SD1 and T+ [T1] SD2 states before RES BRK
                   ; Also hiccups a second T0 and T+ [T1] before RES BRK
;
;                  Test 3
0312 LDA #19       ; Change RES vector to 0319 for test 4
0314 STA FFFC
0317 INC 00
Phase 193, RES0    ; RES down during T+ [T1] Phase 2
Phase 196, RES1    ; RES up after clocking in T0 Phase 1
                   ; Shows T0 T+ SD2 state before RES BRK
;
;                  Test 4
0319 LDA #20       ; Change RES vector to 0320 for test 5
031B STA FFFC
031E INC 00
Phase 229, RES0    ; RES down during T2 Phase 2
Phase 230, RES1    ; RES up after clocking in T3 Phase 1
                   ; Shows T0 SD2 before RES BRK
;
;                  Test 5
0320 LDA #27       ; Change RES vector to 0327 for end (not absolutely necessary)
0322 STA FFFC
0325 BRK #01
Phase 267, RES0    ; RES down during T4 Phase 2
Phase 272, RES1    ; RES up after clocking in first T0 T+ Phase 1
                   ; Shows T0 [T6] before RES BRK
                   ; Only needs RES down for one clock pulse to cause T0 [T6], but needs RES down for five clock pulses
                   ; to outlast the signal from T6 that turns off the interrupt state, allowing us to successfully
                   ; invoke a RES BRK
                   ; That's preferable to the unpredictable PC after a failed invocation of RES BRK (0000, 00FD seen)
0327 NOP           ; Confirm end
```

Some extra effects are noticed from the interaction of RES and the interrupted instruction. The RMW instructions can "hiccup" an extra pair of T0 and T+ [T1] states after the T0 and T+ [T1] states caused directly by RES. This happens when the sense cycle for RES is the SD initiating cycle for the instruction's addressing mode, and RES is let back up before the T0 caused by RES is clocked in (before T0 phase 1). The second T0 is caused by the SD signal chain propagating back to the clock, re-setting it to T0 when it would have gone to the T2 state.

The RMW hiccup phenomenon delays the reset response by an extra two cycles. The delay can be reduced to only one extra cycle by holding RES down long enough to clock in T0 phase 1, and then let RES up before clocking in the next phase 1. After T0 SD1, that will cause the T0 T+ SD2 state to arise (instead of T+ [T1] SD2), and then T+ [T1] after that. The T+ [T1] state will be unaffected by the arrival of the clock resetting signal, unlike states T2 and beyond.

The other extra effect is the requirement for RES to be held down for longer than the minimum needed to be sensed when the instruction being interrupted is a BRK, and it is interrupted with a sense cycle of T5 [V0] or [T6]. The signal chain through the clock extension for BRK will normally turn off the interrupt state (nodes RESG and INTG in visual6502) in phase 1 of the cycle after the \_T6\_ cycle ("\_T6\_" meaning a time state that has [T6] active in it, of which there are two: plain [T6] and T0 [T6]). Sustaining RESG through that cycle after \_T6\_ requires RES to be held down until after phase 1 of that cycle has been clocked in. That's a minimum of five clock pulses for sense at T5 [V0], and three clock pulses for sense at [T6]. Holding RES down so long keeps another node (RESP) on for both phases of that cycle-after-\_T6\_. RESP keeps the RESG node from being grounded by the shut-off signal. The shut-off signal is fully extinct at phase 1 of the second cycle after \_T6\_.

When RES is not held down for long enough, one can see RESG turned off late at phase 2 of the cycle after \_T6\_. When RESP changes state in phase 1, its effect upon RESG is not immediate: there's a further node that is connected to RESP only during phase 2, and that further node (named pipephi2Reset0x) has direct influence upon RESG. Only when phase 2 arrives is that further node connected to RESP again and changed to the same state.

In our case of the further node being driven false by phase 2 connection to a false RESP, it allows RESG to be grounded by the shut-off signal. The phase 2 delay also causes the opposite behavior when RESP goes true: phase 2 of the sense cycle sees RESG ungrounded and become true.

The phase 2 shut-off of RESG happens when RES is held for 4 and 3 clock pulses after sense at T5 [V0], and held for less than three pulses after sense at [T6]. RES hold for 2 and 1 clock pulses after sense at T5 [V0] results in the normal RESG shut-off time.

The normal RESG shut-off (phase 1) corresponds to RESP false during \_T6\_ and the cycle after \_T6\_. The late RESG shut-off at phase 2 corresponds to RESP true during \_T6\_ and false during the cycle after \_T6\_. Prevention of RESG shut-off corresponds to RESP true during \_T6\_ and the cycle after \_T6\_.

We can conclude this last tangent with what happens to BRK execution when RES is not held long enough. Sense cycle of T5 [V0] with all RES holds less than five pulses messes up the PC and fetches the next opcode from an unpredictable address (0000 and 00FD have been witnessed). Sense cycle of [T6] with RES holds of two and one pulses merely substitutes the RES vector for the BRK/IRQ vector and puts control into the RES handler. The original BRK instruction otherwise
finishes normally without really being interrupted.

### External References

[Dr. Donald Hanson's block diagram](http://www.witwright.com/DonPub/6502-Block-Diagram.pdf)

[6502 State Machine](index.php-title-6502_State_Machine.md)

["MCS6500 Microcomputer Family Hardware Manual"](http://archive.6502.org/books/mcs6500_family_hardware_manual.pdf)

[6502 all 256 Opcodes](index.php-title-6502_all_256_Opcodes.md)

["A taken branch delays interrupt handling by one instruction"](http://forum.6502.org/viewtopic.php?f=4&t=1634) forum thread

Retrieved from "[http://visual6502.org/wiki/index.php?title=6502\_Timing\_States](index.php-title-6502_Timing_States.md)"

