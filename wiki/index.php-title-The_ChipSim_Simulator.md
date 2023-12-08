**Recovered visual6502.org wiki - beta release**

# The ChipSim Simulator - VisualChips

## The ChipSim Simulator

#### From VisualChips

The visual6502 project offers an in-browser simulation of the 6502 and 6800 microprocessors. The engine for the simulation, written in JavaScript, is variously called JSSim and ChipSim. This page offers some background on how the simulator works and how it came to be written.

The source code is copyright with an open source license and [available on github](https://github.com/trebonian/visual6502).

**Contents**

- [Algorithm](#algorithm)
- [Data Files](#data-files)
- [History](#history)
- [Resources](#resources)

#### Algorithm

TBC

#### Data Files

To simulate each chip's behaviour, ChipSim needs to know at least the list of transistors and the information as to how they are connected. To be useful it also needs the names of the pins and ideally the names of some internal nets (or nodes.) To display the chip layout and animate it as the logic values change it needs the geometric layout data. All this information is found in three files. Each file populates a single Javascript data structure.

transdefs.js defines an array transdefs[], with a typical element ['t1', 1608, 657, 349, [5424, 5629, 548, 922],[1007, 1079, 17, 5, 17477] ], defining the name, the gate node, the source and drain nodes, the bounding box, and some geometrical information. See [this source file](https://github.com/trebonian/visual6502/blob/master/chip-6800/transdefs.js) for more detail. nodenames.js defines an object nodenames{}, with a typical element nodenames['ab7']=1493. This structure allows us to give several names to each node. We define a function nodeName()to provide an inverse mapping. segdefs.js defines an array segdefs[], with a typical element [4,'+',1,4351,8360,4351,8334,4317,8334,4317,8360], giving the node number, the pullup status, the layer index and a list of coordinate pairs for a polygon. There is one element for each polygon on the chip, and therefore generally several elements for each node. The pullup status can be '+' or '-' and should be consistent for all of a node's entries - it's an historical anomaly that this information is in segdefs. Not all chip layers or polygons appear in segdefs, but enough layers appear for an appealing and educational display.

(These formats are subject to revision, as we model new chips and find information which is worth capturing and modelling. For the 6800, we added a weak/strong indicator to transdefs, which was needed because this chip uses only enhancement mode transistors and had some circuit configurations we hadn't previously encountered.)

Note that the coordinates found in these files are transformed on the way to the JavaScript canvas, so the coordinates in the source files are not the same as those reported in the user interface or used in the URLs.

#### History

TBC

#### Resources

- [source code](https://github.com/trebonian/visual6502) on github
- [Intel 4004 anniversary project](http://www.4004.com/) includes Lajos Kintli's simulator

Retrieved from "[http://visual6502.org/wiki/index.php?title=The\_ChipSim\_Simulator](index.php-title-The_ChipSim_Simulator.md)"

