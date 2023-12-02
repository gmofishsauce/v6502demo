**INCOMPLETE DRAFT OF RECOVERED WIKI PAGE**

# NMOS Depletion Mode Transistors - VisualChips

## NMOS Depletion Mode Transistors

#### From VisualChips

The usual circuit design of a logic gate in NMOS technology is a network of pull-down transistors and a single pull-up. The pull-up will be a depletion mode device, and the gate will be connected to the logic gate's output. The depletion implant adjusts the transistor threshold to below zero volts, with the effects that such a pull-up transistor

- is always on
- can pull all the way to the positive rail
- conducts better at a given voltage, for an improved rising edge.

These pull-ups account for the higher power consumption of NMOS compared to the later CMOS technology.

Our problem in reverse-engineering NMOS chips is that the implant cannot be seen in our photographs.  (There may be staining techniques which will help but we haven't yet tried them.)

So, having identified all the transistors on a chip such as the 6502, we have to engage in some deduction and guesswork to infer which transistors are depletion mode.

First, the easy cases:

- all pull-downs are enhancement mode
- all pass gates are enhancement mode
- all transistors with gate connected to source are depletion mode

Next, there are three cases we can be fairly sure of

- clock drivers must have depletion mode pullups because they need logic 1 to reach the rail
- the huge pull-ups at the output pads will be enhancement mode, which means a logic 0 will pull down to ground, and a logic 1 will be a volt or so less than the positive rail.
- a conventional super buffer circuit with a single pullup driven by the inverse of the single pulldown will have a depletion mode pullup

Here are some cases found on the 6502 which require some judgement and explanation:

- t481, t1169, t1344, t1035 (all in the clock generation)
- t2544, t11  (inverse of phi1, used to pull down some of the datapath control signals)
- t1477 (timing generator)
- Datapath control line drivers (t1527, t439, ..., t2326)
- t2066 the tristate driver
- t367, the driver for RDY.
- t397 SYNC, t2208/t441 R/#W, t578/t3122 D1, and probably all the other first stage pin drivers.
- t2523, t76, t2179, t362 in the clock again, but second stage drivers this time.  Huge huge pulldowns.
- t1322, the dead PLA line.

(All of these are depletion on the Rockwell and Atari 6507+6532 schematics; could not verify SYNC and R/#W and the external clock outputs on those).

(Need to convert the above into links: [t3353](http://visual6502.org/JSSim/expert.html?nosim=t&find=t3353) ADL bus precharge

The precharges for all four internal busses are enhancement mode.

Retrieved from "[http://visual6502.org/wiki/index.php?title=NMOS\_Depletion\_Mode\_Transistors](index.php-title-NMOS_Depletion_Mode_Transistors.md)"

