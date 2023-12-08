**Recovered visual6502.org wiki - beta release**

# 6502Observations - VisualChips

## 6502Observations

#### From VisualChips

We've found some interesting things on the 6502, from the layout level, up through circuit level to the programmer visible level.

## Programmer Visible

Notes here on bugs and undocumented behaviour.

- [BRK, the B bit](index.php-title-6502_BRK_and_B_bit.md), and other interrupts
- [Timing of Interrupt Handling](index.php-title-6502_Timing_of_Interrupt_Handling.md) noting that a taken branch delays interrupt handling, also that CLI/PLP allow one further instruction to execute, unlike RTI.
- [Unsupported or undocumented opcodes](index.php-title-6502_Unsupported_Opcodes.md) such as SAX and XAA
- [The ROR bug](index.php-title-6502_ROR_bug.md) which is found only in rare early devices

See also [our catalogue of 6502 test programs](index.php-title-6502TestPrograms.md), useful to verify simulators or emulators.

## Circuit and Logic

Notes here on timing fixes and non-digital circuit techniques, and departures from NMOS design style orthodoxy.

- [Signs of a fix](index.php-title-6502_datapath_control_timing_fix.md) to datapath control timing

## Layout

Notes here on the traces of bug fixes, and remnants of the original 6501 layout.

- [Traces in the layout](index.php-title-6502_traces_of_6501.md) of the original 6501 part which was withdrawn after legal wrangling

Retrieved from "[http://visual6502.org/wiki/index.php?title=6502Observations](index.php-title-6502Observations.md)"

