**INCOMPLETE DRAFT OF RECOVERED WIKI PAGE**

# 6502TestPrograms - VisualChips

## 6502TestPrograms

#### From VisualChips

There are a number of test suites for 6502, each with their own intentions and peculiarities. (Instead of using a dedicated test suite, it can also be useful to run a monitor or BASIC interpreter, although the test coverage isn't very high and the run time is.  Favourites are: Apple 1 monitor, Apple 1 integer BASIC, C64 BASIC, BBC OS and BASIC.)

For the most part these test programs aim to test emulators, which are subject to different bugs than CPU implementations, and therefore the effective coverage may not be as good as expected.  Such tests are generally self-checking - there is no golden results file of bus activity - and generally assume some specific platform's I/O facilities.

Self-testing (6502 ROMs and programs):

- [Klaus Dormann's test suite](https://github.com/Klaus2m5/6502_65C02_functional_tests) includes decimal mode, is standalone and can be assembled to a single image around 16k.
- [Wolfgang Lorenz' C64 suite](http://www.modelb.bbcmicro.com/testsuite-2.15.tar.gz) exhaustive, excluding decimal mode, uses C64 facilities to chain each program ([stubbing instructions here](http://www.softwolves.com/arkiv/cbm-hackers/7/7114.html) by Christer Palm) (testsuite on Tom Seddon's site) (11kbyte total) ([some description](http://plus4world.powweb.com/software/Test_Suite)) ([another version with sources](http://www.baisoku.org/pc64test.zip))
- Ruud Baltissen's 8k test ROM from his [VHDL 6502 core](http://www.baltissen.org/zip/rb65-11.zip) (includes source, but only a subset of files found in the [previous version](http://www.baltissen.org/zip/rb65-10.zip))
- [NES test](http://www.qmtpro.com/~nes/misc/nestest.txt)[rom](http://nickmass.com/images/nestest.nes) by Kevin Horton (24kbyte) (haven't found source for this, he says he hasn't got clean source to release)
- [AllSuiteA.asm](http://code.google.com/p/hmc-6502/source/browse/trunk/emu/testvectors/AllSuiteA.asm) from the hcm-6502 (verilog) project. ROM available. Load at and reset to 0xf000 and set irq vector to 0xf5a4.
- [Decimal mode tests by Bruce Clark](http://www.6502.org/tutorials/decimal_mode.html) ADC/SBC (exhaustive, tests all four affected flags.) Some specific [Decimal tests here](index.php-title-6502DecimalMode).
- Test code supplied with [Rob Finch's 6502 core](http://web.archive.org/web/20070707064155/http://www.birdcomputer.ca/Projects/Prj6502/bc6502_page.html) (archive.org) (1500 bytes)
- [Acid800](http://www.virtualdub.org/beta/Acid800-0.81.7z) by Avery Lee for 8-bit Atari emulators includes some 6502 tests. See [Altirra](http://www.virtualdub.org/altirra.html) page.
- [ASAP tests](http://asap.git.sourceforge.net/git/gitweb.cgi?p=asap/asap;a=tree;f=test) by Piotr Fusik includes an exhaustive test for ADC, SBC and 0x6B as well as a few tests for other undocumented opcodes
- [64doc](http://www.zimmers.net/anonftp/pub/cbm/documents/chipdata/64doc) contains an exhaustive test for BCD mode, by Marko Mäkelä. The document was originally created by Jouko Valta and/or John West.
- Tim C. Schröder's Neskell project has [a collation of test suites](https://github.com/blitzcode/neskell#test-suite) including [a pair by Blargg](http://slack.net/~ant/misc/) which might not already be mentioned here.
- The VICE project has [a collection of test suites](https://sourceforge.net/p/vice-emu/code/HEAD/tree/testprogs/CPU/).

Test harnesses:

- [py65 tests](https://github.com/mnaberez/py65/tree/master/py65/tests/devices) by Mike Naberezny (python)

References:

- [Nesdev forum topic: "req: nestest.asm"](http://nesdev.parodius.com/bbs/viewtopic.php?p=28348)
- [NES Emulator tests](http://wiki.nesdev.com/w/index.php/Emulator_Tests) wiki page
- [6502.org topic "New 6502 core"](http://forum.6502.org/viewtopic.php?t=1660) For Ruud's announcement
- [6502.org topic "Running test6502.a65 on Py65"](http://forum.6502.org/viewtopic.php?t=1439)
- [6502.org topic "who knows a full test code for 6502?"](http://forum.6502.org/viewtopic.php?t=1436)
- [6502.org topic "Looking for test program"](http://forum.6502.org/viewtopic.php?t=1566)
- [6502.org topic "Op-code testing"](http://forum.6502.org/viewtopic.php?t=547)
- [6502.org topic "Functional Test for the NMOS 6502 - request for verification"](http://forum.6502.org/viewtopic.php?f=2&t=2241)

Retrieved from "[http://visual6502.org/wiki/index.php?title=6502TestPrograms](index.php-title-6502TestPrograms)"

