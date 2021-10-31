[![Wayback
Machine](/_static/images/toolbar/wayback-toolbar-logo-200.png)](/web/ "Wayback Machine home page")

success

fail

[](# "Share via My Web Archive")
[](https://archive.org/account/login.php "Sign In")
[](http://faq.web.archive.org/ "Get some help using the Wayback Machine")
[](#close "Close the toolbar")

[](/web/20210405071552/http://web.archive.org/screenshot/http://visual6502.org/wiki/index.php?title=6502_Timing_of_Interrupt_Handling "screenshot")
[](# "video") [](# "Share on Facebook") [](# "Share on Twitter")

[**Mar**](https://web.archive.org/web/20190319194650/http://visual6502.org/wiki/index.php?title=6502_Timing_of_Interrupt_Handling "19 Mar 2019")

APR

[**Jul**](https://web.archive.org/web/20210724003540/http://visual6502.org/wiki/index.php?title=6502_Timing_of_Interrupt_Handling "24 Jul 2021")

[![Previous
capture](/_static/images/toolbar/wm_tb_prv_on.png)](https://web.archive.org/web/20190319194650/http://visual6502.org/wiki/index.php?title=6502_Timing_of_Interrupt_Handling "19:46:50 Mar 19, 2019")

05

[![Next
capture](/_static/images/toolbar/wm_tb_nxt_on.png)](https://web.archive.org/web/20210724003540/http://visual6502.org/wiki/index.php?title=6502_Timing_of_Interrupt_Handling "00:35:40 Jul 24, 2021")

[**2019**](https://web.archive.org/web/20190319194650/http://visual6502.org/wiki/index.php?title=6502_Timing_of_Interrupt_Handling "19 Mar 2019")

2021

2022

[10
captures](/web/20210405071552*/http://visual6502.org/wiki/index.php?title=6502_Timing_of_Interrupt_Handling "See a list of every capture for this URL")

20 Sep 2014 - 24 Jul 2021

[]()

[About this capture](#expand)

COLLECTED BY

Collection: [Save Page Now](https://archive.org/details/save-page-now)

TIMESTAMPS

![loading](/_static/images/loading.gif)

The Wayback Machine -
https://web.archive.org/web/20210405071552/http://visual6502.org/wiki/index.php?title=6502\_Timing\_of\_Interrupt\_Handling

6502 Timing of Interrupt Handling {#firstHeading .firstHeading}
=================================

### From VisualChips {#siteSub}

Jump to: [navigation](#column-one), [search](#searchInput)

This page contains some as-yet unpolished extracts from postings by user
Hydrophilic on commodore128.org, material used by permission.

This page contains work in progress and unanswered questions, which
should be answered by reference to visual6502 simulation URLs.

+--------------------------------------------------------------------------+
| Contents                                                                 |
| --------                                                                 |
|                                                                          |
| -   [1 Interrupt handling sequence](#Interrupt_handling_sequence)        |
| -   [2 Interrupts colliding](#Interrupts_colliding)                      |
| -   [3 Interrupts and changes to I mask                                  |
|     bit](#Interrupts_and_changes_to_I_mask_bit)                          |
| -   [4 Interrupts during branches](#Interrupts_during_branches)          |
| -   [5 Some simulations for                                              |
|     discussion](#Some_simulations_for_discussion)                        |
| -   [6 Resources](#Resources)                                            |
+--------------------------------------------------------------------------+

### Interrupt handling sequence

The 6502 performs an interrupt as a 7-cycle instruction sequence which
starts with an instruction fetch. (Is this true when the interrupted
instruction is a branch?) The fetched instruction is substituted by a
BRK in the IR.

1.  [IRQ during
    INC](https://web.archive.org/web/20210405071552/http://visual6502.org/JSSim/expert.html?graphics=f&loglevel=2&steps=50&a=0011&d=58&a=fffe&d=2000&a=0020&d=e840&r=0010&irq0=15&irq1=30&logmore=irq&a=0014&d=78)
    showing the latest point at which an IRQ will affect the next
    instruction.
2.  [IRQ during
    SEI](https://web.archive.org/web/20210405071552/http://visual6502.org/JSSim/expert.html?graphics=f&loglevel=2&steps=50&a=0011&d=58&a=fffe&d=2000&a=0020&d=e840&r=0010&irq0=19&irq1=100&logmore=irq&a=0014&d=78)
    which does take effect before the I bit is set
3.  [IRQ one half-cycle later during
    SEI](https://web.archive.org/web/20210405071552/http://visual6502.org/JSSim/expert.html?graphics=f&loglevel=2&steps=50&a=0011&d=58&a=fffe&d=2000&a=0020&d=e840&r=0010&irq0=20&irq1=100&logmore=irq&a=0014&d=78)
    which does not take effect: I has been set and masks the interrupt

### Interrupts colliding

-   this needs to be studied and verified. Observations by Hydrophilic
    follow

[This
simulation](https://web.archive.org/web/20210405071552/http://visual6502.org/JSSim/expert.html?graphics=f&loglevel=1&logmore=Execute,nmi,~NMIP,irq,480,629,INTG&steps=88&a=0011&d=58&a=fffe&d=2000&a=0020&d=e840&r=0010&nmi0=26&nmi1=31&irq0=12&irq1=74&a=0014&d=78)
shows a lost NMI. NMI is brought low when doing an IRQ acknowledge.
Specifically, 1/2 cycle before fetching the IRQ vector (cycle 13 phase
2). NMI remains low for 2.5 cycles. NMI returns high on cycle 16 phase
1.

The NMI is never serviced. This \*might\* be due \#NMIP being
automatically cleared after fetching PC high during any interrupt
response...

### Interrupts and changes to I mask bit

Instructions such as SEI and CLI affect the status register during the
following instruction, due the the 6502 pipelining. Therefore the
masking of the interrupt does not take place until the following
instruction is already underway. However, RTI restores the status
register early, and so the restored I mask bit already is in effect for
the next instruction.

1.  [IRQ during
    CLI](https://web.archive.org/web/20210405071552/http://visual6502.org/JSSim/expert.html?graphics=f&loglevel=2&steps=50&a=0011&d=58&a=fffe&d=2000&a=0020&d=e840&r=0010&irq0=3&irq1=20&logmore=irq)
    (IRQ has no effect on following instruction, interrupt occurs on the
    one after that.) Described as "effect of CLI is delayed by one
    opcode"

### Interrupts during branches

### Some simulations for discussion

1.  [late
    NMI](https://web.archive.org/web/20210405071552/http://visual6502.org/JSSim/expert.html?graphics=f&loglevel=1&logmore=Execute,nmi,~NMIP,irq,480,629,INTG&steps=88&a=0011&d=58&a=fffe&d=2000&a=0020&d=e840&r=0010&nmi0=25&nmi1=27&irq0=12&irq1=64&a=0014&d=78)
    (NMI during IRQ handling, causing NMI vector to be fetched and
    followed.)
2.  [later
    NMI](https://web.archive.org/web/20210405071552/http://visual6502.org/JSSim/expert.html?graphics=f&loglevel=1&logmore=Execute,nmi,~NMIP,irq,480,629,INTG&steps=88&a=0011&d=58&a=fffe&d=2000&a=0020&d=e840&r=0010&nmi0=26&irq0=12&irq1=74&a=0014&d=78)
    (NMI during IRQ handling, just prior to vector fetch, too late to
    usurp the IRQ, does not even interrupt the first instruction of the
    IRQ handler, but the second one.)

<!-- -->

     Perhaps this is because 'doIRQ' (line 480) is not set during the last vector fetch; that is, during cycle 15 when fetching $FFFF, 'doIRQ' is still false.  I don't know why, considering both INTG and 'normal' (line 629) have 'standard' values and there is NMI pending.
     After starting to work on INX, 'doIRQ' finally gets set correctly so that 'normal' and INTG get triggered at the end of its execution (cycle 18).  And then finally (cycle 19) the NMI is processed

1.  [lost
    NMI](https://web.archive.org/web/20210405071552/http://visual6502.org/JSSim/expert.html?graphics=f&loglevel=1&logmore=Execute,nmi,~NMIP,irq,480,629,INTG&steps=88&a=0011&d=58&a=fffe&d=2000&a=0020&d=e840&r=0010&nmi0=26&nmi1=31&irq0=12&irq1=74&a=0014&d=78)
    (NMI during IRQ handling, showing that 2.5 cycles is too short for
    an NMI to be serviced if it falls during this critical time of
    fetching the IRQ vector)
2.  [IRQ delayed by
    branch](https://web.archive.org/web/20210405071552/http://visual6502.org/JSSim/expert.html?graphics=f&loglevel=1&logmore=Execute,nmi,%23NMIP,irq,480,629,INTG&steps=88&a=0010&d=58e8&a=fffe&d=2000&a=0020&d=e840&r=0010&nmi0=26&nmi1=31&irq0=11&irq1=74&a=0012&d=d0fe&a=0014&d=78).
    As can be seen from previous simulations, the CPU will normally
    examine 'do IRQ' on the last cycle of an instruction and if it is
    set, will clear 'normal' (line 629) and set INTG.

However, in the last cycle of BNE (branch taken, no page cross),
although the IRQ has been asserted and 'do IRQ' has been set true (cycle
6), the 'normal' and INTG lines are not updated. So the CPU continues
with the next instruction, also BNE. Again 'do IRQ' is examined and the
'normal' and INTG are updated, but not during the last cycle of the
instruction (as per MOS Tech specs) but actually during the next-to-last
cycle (see cycle 13). Note if the branch were not taken, it would be the
last cycle of the instruction.

Perhaps the designers considered cycle 2 of any branch instruction to be
the 'natural' end and check 'do IRQ' there... and only there... unless a
page boundary is crossed.

Now that I think of it, the fact that INTG is set on the second cycle of
any branch instruction (regardless if branch is taken or not), means
this line is set 1 cycles earlier than normal if the branch does get
taken, and 2 cycles earlier than normal if the branch crosses a page
boundary.

(The above commentary, as a pastebomb, should be revisited and tidied up
and reconciled with other sources. If we fail to explain, we can remove
our explanations and leave only our evidence.)

### Resources

-   back to parent page
    [6502Observations](/web/20210405071552/http://visual6502.org/wiki/index.php?title=6502Observations "6502Observations")
-   [A taken branch delays interrupt handling by one
    instruction](https://web.archive.org/web/20210405071552/http://forum.6502.org/viewtopic.php?t=1634)
    forum thread on 6502.org
-   [A Taken branch before NMI delays NMI execution by one full
    instruction](https://web.archive.org/web/20210405071552/http://www.atariage.com/forums/topic/168550-a-taken-branch-before-nmi-delays-nmi-execution-by-one-full-instruction/)
    forum thread on AtariAge
-   [CLI - My 8502 is
    defective ???](https://web.archive.org/web/20210405071552/http://www.commodore128.org/index.php?topic=3863)
    forum thread on commodore128.org
-   [Effects of SEI and CLI delayed by one
    opcode?](https://web.archive.org/web/20210405071552/http://forum.6502.org/viewtopic.php?t=1817)
    forum thread on 6502.org
-   [How can POKEY IRQ Timers mess up NMI
    timing?](https://web.archive.org/web/20210405071552/http://www.atariage.com/forums/topic/148595-how-can-pokey-irq-timers-mess-up-nmi-timing/page__st__100__p__1816157#entry1816157)
    forum thread on AtariAge (missed NMI)

Retrieved from
"[http://visual6502.org/wiki/index.php?title=6502\_Timing\_of\_Interrupt\_Handling](https://web.archive.org/web/20210405071552/http://visual6502.org/wiki/index.php?title=6502_Timing_of_Interrupt_Handling)"

##### Views

-   [Page](/web/20210405071552/http://visual6502.org/wiki/index.php?title=6502_Timing_of_Interrupt_Handling "View the content page [c]")
-   [Discussion](/web/20210405071552/http://visual6502.org/wiki/index.php?title=Talk:6502_Timing_of_Interrupt_Handling&action=edit&redlink=1 "Discussion about the content page [t]")
-   [View
    source](/web/20210405071552/http://visual6502.org/wiki/index.php?title=6502_Timing_of_Interrupt_Handling&action=edit "This page is protected.
    You can view its source [e]")
-   [History](/web/20210405071552/http://visual6502.org/wiki/index.php?title=6502_Timing_of_Interrupt_Handling&action=history "Past revisions of this page [h]")

##### Personal tools

-   [Log
    in](/web/20210405071552/http://visual6502.org/wiki/index.php?title=Special:UserLogin&returnto=6502_Timing_of_Interrupt_Handling "You are encouraged to log in; however, it is not mandatory [o]")

[](/web/20210405071552/http://visual6502.org/wiki/index.php?title=Main_Page "Visit the main page")

##### Navigation

-   [Main
    page](/web/20210405071552/http://visual6502.org/wiki/index.php?title=Main_Page "Visit the main page [z]")
-   [Community
    portal](/web/20210405071552/http://visual6502.org/wiki/index.php?title=VisualChips:Community_portal "About the project, what you can do, where to find things")
-   [Current
    events](/web/20210405071552/http://visual6502.org/wiki/index.php?title=VisualChips:Current_events "Find background information on current events")
-   [Recent
    changes](/web/20210405071552/http://visual6502.org/wiki/index.php?title=Special:RecentChanges "The list of recent changes in the wiki [r]")
-   [Random
    page](/web/20210405071552/http://visual6502.org/wiki/index.php?title=Special:Random "Load a random page [x]")
-   [Help](/web/20210405071552/http://visual6502.org/wiki/index.php?title=Help:Contents "The place to find out")

##### Search

 

##### Toolbox

-   [What links
    here](/web/20210405071552/http://visual6502.org/wiki/index.php?title=Special:WhatLinksHere/6502_Timing_of_Interrupt_Handling "List of all wiki pages that link here [j]")
-   [Related
    changes](/web/20210405071552/http://visual6502.org/wiki/index.php?title=Special:RecentChangesLinked/6502_Timing_of_Interrupt_Handling "Recent changes in pages linked from this page [k]")
-   [Special
    pages](/web/20210405071552/http://visual6502.org/wiki/index.php?title=Special:SpecialPages "List of all special pages [q]")
-   [Printable
    version](/web/20210405071552/http://visual6502.org/wiki/index.php?title=6502_Timing_of_Interrupt_Handling&printable=yes "Printable version of this page [p]")
-   [Permanent
    link](/web/20210405071552/http://visual6502.org/wiki/index.php?title=6502_Timing_of_Interrupt_Handling&oldid=598 "Permanent link to this revision of the page")

[![Powered by
MediaWiki](/web/20210405071552im_/http://visual6502.org/wiki/skins/common/images/poweredby_mediawiki_88x31.png)](https://web.archive.org/web/20210405071552/http://www.mediawiki.org/)

[![Attribution-NonCommercial-ShareAlike 3.0
Unported](https://web.archive.org/web/20210405071552im_/http://i.creativecommons.org/l/by-nc-sa/3.0/88x31.png)](https://web.archive.org/web/20210405071552/http://creativecommons.org/licenses/by-nc-sa/3.0/)

-   This page was last modified on 17 May 2011, at 19:48.
-   This page has been accessed 92,588 times.
-   Content is available under [Attribution-NonCommercial-ShareAlike 3.0
    Unported](https://web.archive.org/web/20210405071552/http://creativecommons.org/licenses/by-nc-sa/3.0/).
-   [Privacy
    policy](/web/20210405071552/http://visual6502.org/wiki/index.php?title=VisualChips:Privacy_policy "VisualChips:Privacy policy")
-   [About
    VisualChips](/web/20210405071552/http://visual6502.org/wiki/index.php?title=VisualChips:About "VisualChips:About")
-   [Disclaimers](/web/20210405071552/http://visual6502.org/wiki/index.php?title=VisualChips:General_disclaimer "VisualChips:General disclaimer")

