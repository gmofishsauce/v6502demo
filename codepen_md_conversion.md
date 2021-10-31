


window.addEventListener('DOMContentLoaded',function(){var v=archive_analytics.values;v.service='wb';v.server_name='wwwb-app224.us.archive.org';v.server_ms=344;archive_analytics.send_pageview({});});



  __wm.init("https://web.archive.org/web");
  __wm.wombat("http://visual6502.org/wiki/index.php?title=6502_Timing_of_Interrupt_Handling","20210405071552","https://web.archive.org/","web","/_static/",
      "1617606952");





6502 Timing of Interrupt Handling - VisualChips



















var skin="monobook",
stylepath="/wiki/skins",
wgUrlProtocols="http\\:\\/\\/|https\\:\\/\\/|ftp\\:\\/\\/|irc\\:\\/\\/|gopher\\:\\/\\/|telnet\\:\\/\\/|nntp\\:\\/\\/|worldwind\\:\\/\\/|mailto\\:|news\\:|svn\\:\\/\\/",
wgArticlePath="/wiki/index.php?title=$1",
wgScriptPath="/wiki",
wgScriptExtension=".php",
wgScript="/wiki/index.php",
wgVariantArticlePath=false,
wgActionPaths={},
wgServer="https://web.archive.org/web/20210405071552/http://visual6502.org",
wgCanonicalNamespace="",
wgCanonicalSpecialPageName=false,
wgNamespaceNumber=0,
wgPageName="6502_Timing_of_Interrupt_Handling",
wgTitle="6502 Timing of Interrupt Handling",
wgAction="view",
wgArticleId=61,
wgIsArticle=true,
wgUserName=null,
wgUserGroups=null,
wgUserLanguage="en",
wgContentLanguage="en",
wgBreakFrames=false,
wgCurRevisionId=598,
wgVersion="1.16.0",
wgEnableAPI=true,
wgEnableWriteAPI=true,
wgSeparatorTransformTable=["", ""],
wgDigitTransformTable=["", ""],
wgMainPageTitle="Main Page",
wgFormattedNamespaces={"-2": "Media", "-1": "Special", "0": "", "1": "Talk", "2": "User", "3": "User talk", "4": "VisualChips", "5": "VisualChips talk", "6": "File", "7": "File talk", "8": "MediaWiki", "9": "MediaWiki talk", "10": "Template", "11": "Template talk", "12": "Help", "13": "Help talk", "14": "Category", "15": "Category talk"},
wgNamespaceIds={"media": -2, "special": -1, "": 0, "talk": 1, "user": 2, "user_talk": 3, "visualchips": 4, "visualchips_talk": 5, "file": 6, "file_talk": 7, "mediawiki": 8, "mediawiki_talk": 9, "template": 10, "template_talk": 11, "help": 12, "help_talk": 13, "category": 14, "category_talk": 15, "image": 6, "image_talk": 7},
wgSiteName="VisualChips",
wgCategories=[],
wgRestrictionEdit=["autoconfirmed"],
wgRestrictionMove=["autoconfirmed"];







body {
  margin-top:0 !important;
  padding-top:0 !important;
  /*min-width:800px !important;*/
}

__wm.rw(0);



  
    
      [|IMG|](/web/)
    
    
      
                  success
          fail
          <a id="wm-save-snapshot-open" href="#"
     title="Share via My Web Archive" >
            
          
          <a href="https://archive.org/account/login.php"
             title="Sign In"
             id="wm-sign-in"
          >
            
          
          
        [](http://faq.web.archive.org/)
[](#close)
      
      
          <a href="/web/20210405071552/http://web.archive.org/screenshot/http://visual6502.org/wiki/index.php?title=6502_Timing_of_Interrupt_Handling"
             id="wm-screenshot"
             title="screenshot">
            
          
          <a href="#"
            id="wm-video"
            title="video">
            
          
[](#)
[](#)
      
    
    
      

  
    
  
  
    
      


  [**Mar**](https://web.archive.org/web/20190319194650/http://visual6502.org/wiki/index.php?title=6502_Timing_of_Interrupt_Handling)
  APR
  [**Jul**](https://web.archive.org/web/20210724003540/http://visual6502.org/wiki/index.php?title=6502_Timing_of_Interrupt_Handling)



  [|IMG|](https://web.archive.org/web/20190319194650/http://visual6502.org/wiki/index.php?title=6502_Timing_of_Interrupt_Handling)
  05
  [|IMG|](https://web.archive.org/web/20210724003540/http://visual6502.org/wiki/index.php?title=6502_Timing_of_Interrupt_Handling)



  [**2019**](https://web.archive.org/web/20190319194650/http://visual6502.org/wiki/index.php?title=6502_Timing_of_Interrupt_Handling)
  2021
  2022

      
    
  


  
        
            [10 captures](/web/20210405071552*/http://visual6502.org/wiki/index.php?title=6502_Timing_of_Interrupt_Handling)
      20 Sep 2014 - 24 Jul 2021
      
  
  
    [
      

      
    ]()
  

      
    
    
      [ About this capture](#expand)
    
  
    
                    
    COLLECTED BY
    
      

Collection: [Save Page Now](https://archive.org/details/save-page-now)
      
    
    
    
    TIMESTAMPS
    
      
      ![Wayback Machine](/_static/images/toolbar/wayback-toolbar-logo-200.png)
    
    
  The Wayback Machine - https://web.archive.org/web/20210405071552/http://visual6502.org/wiki/index.php?title=6502_Timing_of_Interrupt_Handling

  
    <iframe id="donato-if" src="https://archive.org/includes/donate.php?as_page=1&amp;platform=wb&amp;referer=https%3A//web.archive.org/web/20210405071552/http%3A//visual6502.org/wiki/index.php%3Ftitle%3D6502_Timing_of_Interrupt_Handling"
    scrolling="no" frameborder="0" style="width:100%; height:100%">
    
  

__wm.bt(650,27,25,2,"web","http://visual6502.org/wiki/index.php?title=6502_Timing_of_Interrupt_Handling","20210405071552",1996,"/_static/",["/_static/css/banner-styles.css?v=omkqRugM","/_static/css/iconochive.css?v=qtvMKcIJ"], "False");
  __wm.rw(1);




[](null)

6502 Timing of Interrupt Handling

From VisualChips

Jump to: [navigation](#column-one), [search](#searchInput)

This page contains some as-yet unpolished extracts from postings by user Hydrophilic on commodore128.org, material used by permission.
This page contains work in progress and unanswered questions, which should be answered by reference to visual6502 simulation URLs.

## Contents

<li class="toclevel-1 tocsection-1"><a href="#Interrupt_handling_sequence"><span class="tocnumber">1</span> <span class="toctext">Interrupt handling sequence</span></a>
<li class="toclevel-1 tocsection-2"><a href="#Interrupts_colliding"><span class="tocnumber">2</span> <span class="toctext">Interrupts colliding</span></a>
<li class="toclevel-1 tocsection-3"><a href="#Interrupts_and_changes_to_I_mask_bit"><span class="tocnumber">3</span> <span class="toctext">Interrupts and changes to I mask bit</span></a>
<li class="toclevel-1 tocsection-4"><a href="#Interrupts_during_branches"><span class="tocnumber">4</span> <span class="toctext">Interrupts during branches</span></a>
<li class="toclevel-1 tocsection-5"><a href="#Some_simulations_for_discussion"><span class="tocnumber">5</span> <span class="toctext">Some simulations for discussion</span></a>
<li class="toclevel-1 tocsection-6"><a href="#Resources"><span class="tocnumber">6</span> <span class="toctext">Resources</span></a>

if (window.showTocToggle) { var tocShowText = "show"; var tocHideText = "hide"; showTocToggle(); } 
###   Interrupt handling sequence 
The 6502 performs an interrupt as a 7-cycle instruction sequence which starts with an instruction fetch. (Is this true when the interrupted instruction is a branch?)  The fetched instruction is substituted by a BRK in the IR.

1.  <a href="https://web.archive.org/web/20210405071552/http://visual6502.org/JSSim/expert.html?graphics=f&amp;loglevel=2&amp;steps=50&amp;a=0011&amp;d=58&amp;a=fffe&amp;d=2000&amp;a=0020&amp;d=e840&amp;r=0010&amp;irq0=15&amp;irq1=30&amp;logmore=irq&amp;a=0014&amp;d=78" class="external text" rel="nofollow">IRQ during INC</a> showing the latest point at which an IRQ will affect the next instruction.
1.  <a href="https://web.archive.org/web/20210405071552/http://visual6502.org/JSSim/expert.html?graphics=f&amp;loglevel=2&amp;steps=50&amp;a=0011&amp;d=58&amp;a=fffe&amp;d=2000&amp;a=0020&amp;d=e840&amp;r=0010&amp;irq0=19&amp;irq1=100&amp;logmore=irq&amp;a=0014&amp;d=78" class="external text" rel="nofollow">IRQ during SEI</a> which does take effect before the I bit is set
1.  <a href="https://web.archive.org/web/20210405071552/http://visual6502.org/JSSim/expert.html?graphics=f&amp;loglevel=2&amp;steps=50&amp;a=0011&amp;d=58&amp;a=fffe&amp;d=2000&amp;a=0020&amp;d=e840&amp;r=0010&amp;irq0=20&amp;irq1=100&amp;logmore=irq&amp;a=0014&amp;d=78" class="external text" rel="nofollow">IRQ one half-cycle later during SEI</a> which does not take effect: I has been set and masks the interrupt

###   Interrupts colliding 
*  this needs to be studied and verified. Observations by Hydrophilic follow

[This simulation](https://web.archive.org/web/20210405071552/http://visual6502.org/JSSim/expert.html?graphics=f&loglevel=1&logmore=Execute,nmi,~NMIP,irq,480,629,INTG&steps=88&a=0011&d=58&a=fffe&d=2000&a=0020&d=e840&r=0010&nmi0=26&nmi1=31&irq0=12&irq1=74&a=0014&d=78) shows a lost NMI. NMI is brought low when doing an IRQ acknowledge.  Specifically, 1/2 cycle before fetching the IRQ vector (cycle 13 phase 2).  NMI remains low for 2.5 cycles.  NMI returns high on cycle 16 phase 1.  
The NMI is never serviced.  This *might* be due #NMIP being automatically cleared after fetching PC high during any interrupt response...

###   Interrupts and changes to I mask bit 
Instructions such as SEI and CLI affect the status register during the following instruction, due the the 6502 pipelining. Therefore the masking of the interrupt does not take place until the following instruction is already underway.  However, RTI restores the status register early, and so the restored I mask bit already is in effect for the next instruction.

1.  <a href="https://web.archive.org/web/20210405071552/http://visual6502.org/JSSim/expert.html?graphics=f&amp;loglevel=2&amp;steps=50&amp;a=0011&amp;d=58&amp;a=fffe&amp;d=2000&amp;a=0020&amp;d=e840&amp;r=0010&amp;irq0=3&amp;irq1=20&amp;logmore=irq" class="external text" rel="nofollow">IRQ during CLI</a> (IRQ has no effect on following instruction, interrupt occurs on the one after that.) Described as "effect of CLI is delayed by one opcode"

###   Interrupts during branches 
###   Some simulations for discussion 
1.  <a href="https://web.archive.org/web/20210405071552/http://visual6502.org/JSSim/expert.html?graphics=f&amp;loglevel=1&amp;logmore=Execute,nmi,~NMIP,irq,480,629,INTG&amp;steps=88&amp;a=0011&amp;d=58&amp;a=fffe&amp;d=2000&amp;a=0020&amp;d=e840&amp;r=0010&amp;nmi0=25&amp;nmi1=27&amp;irq0=12&amp;irq1=64&amp;a=0014&amp;d=78" class="external text" rel="nofollow">late NMI</a> (NMI during IRQ handling, causing NMI vector to be fetched and followed.)
1.  <a href="https://web.archive.org/web/20210405071552/http://visual6502.org/JSSim/expert.html?graphics=f&amp;loglevel=1&amp;logmore=Execute,nmi,~NMIP,irq,480,629,INTG&amp;steps=88&amp;a=0011&amp;d=58&amp;a=fffe&amp;d=2000&amp;a=0020&amp;d=e840&amp;r=0010&amp;nmi0=26&amp;irq0=12&amp;irq1=74&amp;a=0014&amp;d=78" class="external text" rel="nofollow">later NMI</a> (NMI during IRQ handling, just prior to vector fetch, too late to usurp the IRQ, does not even interrupt the first instruction of the IRQ handler, but the second one.)

 Perhaps this is because 'doIRQ' (line 480) is not set during the last vector fetch; that is, during cycle 15 when fetching $FFFF, 'doIRQ' is still false.  I don't know why, considering both INTG and 'normal' (line 629) have 'standard' values and there is NMI pending.
 After starting to work on INX, 'doIRQ' finally gets set correctly so that 'normal' and INTG get triggered at the end of its execution (cycle 18).  And then finally (cycle 19) the NMI is processed

1.  <a href="https://web.archive.org/web/20210405071552/http://visual6502.org/JSSim/expert.html?graphics=f&amp;loglevel=1&amp;logmore=Execute,nmi,~NMIP,irq,480,629,INTG&amp;steps=88&amp;a=0011&amp;d=58&amp;a=fffe&amp;d=2000&amp;a=0020&amp;d=e840&amp;r=0010&amp;nmi0=26&amp;nmi1=31&amp;irq0=12&amp;irq1=74&amp;a=0014&amp;d=78" class="external text" rel="nofollow">lost NMI</a> (NMI during IRQ handling, showing that 2.5 cycles is too short for an NMI to be serviced if it falls during this critical time of fetching the IRQ vector) 
1.  <a href="https://web.archive.org/web/20210405071552/http://visual6502.org/JSSim/expert.html?graphics=f&amp;loglevel=1&amp;logmore=Execute,nmi,%23NMIP,irq,480,629,INTG&amp;steps=88&amp;a=0010&amp;d=58e8&amp;a=fffe&amp;d=2000&amp;a=0020&amp;d=e840&amp;r=0010&amp;nmi0=26&amp;nmi1=31&amp;irq0=11&amp;irq1=74&amp;a=0012&amp;d=d0fe&amp;a=0014&amp;d=78" class="external text" rel="nofollow">IRQ delayed by branch</a>. As can be seen from previous simulations, the CPU will normally examine 'do IRQ' on the last cycle of an instruction and if it is set, will clear 'normal' (line 629) and set INTG.

However, in the last cycle of BNE (branch taken, no page cross), although the IRQ has been asserted and 'do IRQ' has been set true (cycle 6), the 'normal' and INTG lines are not updated.  So the CPU continues with the next instruction, also BNE. Again 'do IRQ' is examined and the 'normal' and INTG are updated, but not during the last cycle of the instruction (as per MOS Tech specs) but actually during the next-to-last cycle (see cycle 13).  Note if the branch were not taken, it would be the last cycle of the instruction.
Perhaps the designers considered cycle 2 of any branch instruction to be the 'natural' end and check 'do IRQ' there... and only there... unless a page boundary is crossed.
Now that I think of it, the fact that INTG is set on the second cycle of any branch instruction (regardless if branch is taken or not), means this line is set 1 cycles earlier than normal if the branch does get taken, and 2 cycles earlier than normal if the branch crosses a page boundary.
(The above commentary, as a pastebomb, should be revisited and tidied up and reconciled with other sources. If we fail to explain, we can remove our explanations and leave only our evidence.)

###   Resources 
*  back to parent page <a href="/web/20210405071552/http://visual6502.org/wiki/index.php?title=6502Observations" title="6502Observations">6502Observations</a>
*  <a href="https://web.archive.org/web/20210405071552/http://forum.6502.org/viewtopic.php?t=1634" class="external text" rel="nofollow">A taken branch delays interrupt handling by one instruction</a> forum thread on 6502.org
*  <a href="https://web.archive.org/web/20210405071552/http://www.atariage.com/forums/topic/168550-a-taken-branch-before-nmi-delays-nmi-execution-by-one-full-instruction/" class="external text" rel="nofollow">A Taken branch before NMI delays NMI execution by one full instruction</a> forum thread on AtariAge
*  <a href="https://web.archive.org/web/20210405071552/http://www.commodore128.org/index.php?topic=3863" class="external text" rel="nofollow">CLI - My 8502 is defective&nbsp;???</a> forum thread on commodore128.org
*  <a href="https://web.archive.org/web/20210405071552/http://forum.6502.org/viewtopic.php?t=1817" class="external text" rel="nofollow">Effects of SEI and CLI delayed by one opcode?</a> forum thread on 6502.org
*  <a href="https://web.archive.org/web/20210405071552/http://www.atariage.com/forums/topic/148595-how-can-pokey-irq-timers-mess-up-nmi-timing/page__st__100__p__1816157#entry1816157" class="external text" rel="nofollow">How can POKEY IRQ Timers mess up NMI timing?</a> forum thread on AtariAge (missed NMI)


<!-- 
NewPP limit report
Preprocessor node count: 23/1000000
Post-expand include size: 0/2097152 bytes
Template argument size: 0/2097152 bytes
Expensive parser function count: 0/100
-->



Retrieved from "[http://visual6502.org/wiki/index.php?title=6502_Timing_of_Interrupt_Handling](https://web.archive.org/web/20210405071552/http://visual6502.org/wiki/index.php?title=6502_Timing_of_Interrupt_Handling)"






##### Views


 <li id="ca-nstab-main" class="selected"><a href="/web/20210405071552/http://visual6502.org/wiki/index.php?title=6502_Timing_of_Interrupt_Handling" title="View the content page [c]" accesskey="c">Page</a>
 <li id="ca-talk" class="new"><a href="/web/20210405071552/http://visual6502.org/wiki/index.php?title=Talk:6502_Timing_of_Interrupt_Handling&amp;action=edit&amp;redlink=1" title="Discussion about the content page [t]" accesskey="t">Discussion</a>
 <li id="ca-viewsource"><a href="/web/20210405071552/http://visual6502.org/wiki/index.php?title=6502_Timing_of_Interrupt_Handling&amp;action=edit" title="This page is protected.
You can view its source [e]" accesskey="e">View source</a>
 <li id="ca-history"><a href="/web/20210405071552/http://visual6502.org/wiki/index.php?title=6502_Timing_of_Interrupt_Handling&amp;action=history" title="Past revisions of this page [h]" accesskey="h">History</a>




##### Personal tools


<li id="pt-login"><a href="/web/20210405071552/http://visual6502.org/wiki/index.php?title=Special:UserLogin&amp;returnto=6502_Timing_of_Interrupt_Handling" title="You are encouraged to log in; however, it is not mandatory [o]" accesskey="o">Log in</a>




[](/web/20210405071552/http://visual6502.org/wiki/index.php?title=Main_Page)

 if (window.isMSIE55) fixalpha(); 

##### Navigation


<li id="n-mainpage-description"><a href="/web/20210405071552/http://visual6502.org/wiki/index.php?title=Main_Page" title="Visit the main page [z]" accesskey="z">Main page</a>
<li id="n-portal"><a href="/web/20210405071552/http://visual6502.org/wiki/index.php?title=VisualChips:Community_portal" title="About the project, what you can do, where to find things">Community portal</a>
<li id="n-currentevents"><a href="/web/20210405071552/http://visual6502.org/wiki/index.php?title=VisualChips:Current_events" title="Find background information on current events">Current events</a>
<li id="n-recentchanges"><a href="/web/20210405071552/http://visual6502.org/wiki/index.php?title=Special:RecentChanges" title="The list of recent changes in the wiki [r]" accesskey="r">Recent changes</a>
<li id="n-randompage"><a href="/web/20210405071552/http://visual6502.org/wiki/index.php?title=Special:Random" title="Load a random page [x]" accesskey="x">Random page</a>
<li id="n-help"><a href="/web/20210405071552/http://visual6502.org/wiki/index.php?title=Help:Contents" title="The place to find out">Help</a>




##### Search




&nbsp;





##### Toolbox


<li id="t-whatlinkshere"><a href="/web/20210405071552/http://visual6502.org/wiki/index.php?title=Special:WhatLinksHere/6502_Timing_of_Interrupt_Handling" title="List of all wiki pages that link here [j]" accesskey="j">What links here</a>
<li id="t-recentchangeslinked"><a href="/web/20210405071552/http://visual6502.org/wiki/index.php?title=Special:RecentChangesLinked/6502_Timing_of_Interrupt_Handling" title="Recent changes in pages linked from this page [k]" accesskey="k">Related changes</a>
<li id="t-specialpages"><a href="/web/20210405071552/http://visual6502.org/wiki/index.php?title=Special:SpecialPages" title="List of all special pages [q]" accesskey="q">Special pages</a>
<li id="t-print"><a href="/web/20210405071552/http://visual6502.org/wiki/index.php?title=6502_Timing_of_Interrupt_Handling&amp;printable=yes" rel="alternate" title="Printable version of this page [p]" accesskey="p">Printable version</a><li id="t-permalink"><a href="/web/20210405071552/http://visual6502.org/wiki/index.php?title=6502_Timing_of_Interrupt_Handling&amp;oldid=598" title="Permanent link to this revision of the page">Permanent link</a>





[|IMG|](https://web.archive.org/web/20210405071552/http://www.mediawiki.org/)
[|IMG|](https://web.archive.org/web/20210405071552/http://creativecommons.org/licenses/by-nc-sa/3.0/)
<ul id="f-list">
<li id="lastmod"> This page was last modified on 17 May 2011, at 19:48.
<li id="viewcount">This page has been accessed 92,588 times.
<li id="copyright">Content is available under <a href="https://web.archive.org/web/20210405071552/http://creativecommons.org/licenses/by-nc-sa/3.0/" class="external ">Attribution-NonCommercial-ShareAlike 3.0 Unported</a>.
<li id="privacy"><a href="/web/20210405071552/http://visual6502.org/wiki/index.php?title=VisualChips:Privacy_policy" title="VisualChips:Privacy policy">Privacy policy</a>
<li id="about"><a href="/web/20210405071552/http://visual6502.org/wiki/index.php?title=VisualChips:About" title="VisualChips:About">About VisualChips</a>
<li id="disclaimer"><a href="/web/20210405071552/http://visual6502.org/wiki/index.php?title=VisualChips:General_disclaimer" title="VisualChips:General disclaimer">Disclaimers</a>




if (window.runOnloadHook) runOnloadHook();

<!--
     FILE ARCHIVED ON 07:15:52 Apr 05, 2021 AND RETRIEVED FROM THE
     INTERNET ARCHIVE ON 02:52:17 Oct 31, 2021.
     JAVASCRIPT APPENDED BY WAYBACK MACHINE, COPYRIGHT INTERNET ARCHIVE.

     ALL OTHER CONTENT MAY ALSO BE PROTECTED BY COPYRIGHT (17 U.S.C.
     SECTION 108(a)(3)).
-->
<!--
playback timings (ms):
  captures_list: 204.026
  exclusion.robots: 0.22
  exclusion.robots.policy: 0.213
  RedisCDXSource: 20.614
  esindex: 0.006
  LoadShardBlock: 160.323 (3)
  PetaboxLoader3.datanode: 191.278 (4)
  CDXLines.iter: 19.71 (3)
  load_resource: 132.031
  PetaboxLoader3.resolve: 68.3
-->


