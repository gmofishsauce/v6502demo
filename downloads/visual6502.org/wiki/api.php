<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN">
<html>
<head>
	<title>MediaWiki API</title>
</head>
<body>
<pre>
<span style="color:blue;">&lt;?xml version=&quot;1.0&quot;?&gt;</span>
<span style="color:blue;">&lt;api&gt;</span>
  <span style="color:blue;">&lt;error code=&quot;help&quot; info=&quot;&quot; xml:space=&quot;preserve&quot;&gt;</span>
  
  
  <b>******************************************************************</b>
  <b>**                                                              **</b>
  <b>**  This is an auto-generated MediaWiki API documentation page  **</b>
  <b>**                                                              **</b>
  <b>**                  Documentation and Examples:                 **</b>
  **               <a href="http://www.mediawiki.org/wiki/API">http://www.mediawiki.org/wiki/API</a>              **
  <b>**                                                              **</b>
  <b>******************************************************************</b>
  
  Status:          All features shown on this page should be working, but the API
                   is still in active development, and  may change at any time.
                   Make sure to monitor our mailing list for any updates.
  
  Documentation:   <a href="http://www.mediawiki.org/wiki/API">http://www.mediawiki.org/wiki/API</a>
  Mailing list:    <a href="http://lists.wikimedia.org/mailman/listinfo/mediawiki-api">http://lists.wikimedia.org/mailman/listinfo/mediawiki-api</a>
  Bugs &amp; Requests: <a href="http://bugzilla.wikimedia.org/buglist.cgi?component=API&amp;bug_status=NEW&amp;bug_status=ASSIGNED&amp;bug_status=REOPENED&amp;order=bugs.delta_ts">http://bugzilla.wikimedia.org/buglist.cgi?component=API&amp;bug_status=NEW&amp;bug_status=ASSIGNED&amp;bug_status=REOPENED&amp;order=bugs.delta_ts</a>
  
  
  
  
  
Parameters:
  format         - The format of the output
                   One value: json, jsonfm, php, phpfm, wddx, wddxfm, xml, xmlfm, yaml, yamlfm, rawfm, txt, txtfm, dbg, dbgfm
                   Default: xmlfm
  action         - What action you would like to perform
                   One value: login, logout, query, expandtemplates, parse, opensearch, feedwatchlist, help, paraminfo, purge, rollback, delete, undelete, protect, block, unblock, move, edit, upload, emailuser, watch, patrol, import, userrights
                   Default: help
  version        - When showing help, include version for each module
  maxlag         - Maximum lag
  smaxage        - Set the s-maxage header to this many seconds. Errors are never cached
                   Default: 0
  maxage         - Set the max-age header to this many seconds. Errors are never cached
                   Default: 0
  requestid      - Request ID to distinguish requests. This will just be output back to you


<b>*** *** *** *** *** *** *** *** *** ***  Modules  *** *** *** *** *** *** *** *** *** ***</b> 

<b>* action=login (lg) *</b>
  This module is used to login and get the authentication tokens. 
  In the event of a successful log-in, a cookie will be attached
  to your session. In the event of a failed log-in, you will not 
  be able to attempt another log-in through this method for 5 seconds.
  This is to prevent password guessing by automated password crackers.

This module only accepts POST requests.
Parameters:
  lgname         - User Name
  lgpassword     - Password
  lgdomain       - Domain (optional)
  lgtoken        - Login token obtained in first request
Example:
  <a href="api.php?action=login&amp;lgname=user&amp;lgpassword=password">api.php?action=login&amp;lgname=user&amp;lgpassword=password</a>

<b>* action=logout *</b>
  This module is used to logout and clear session data
Example:
  <a href="api.php?action=logout">api.php?action=logout</a>

<b>* action=query *</b>
  Query API module allows applications to get needed pieces of data from the MediaWiki databases,
  and is loosely based on the old query.php interface.
  All data modifications will first have to use query to acquire a token to prevent abuse from malicious sites.

This module requires read rights.
Parameters:
  titles         - A list of titles to work on
  pageids        - A list of page IDs to work on
  revids         - A list of revision IDs to work on
  prop           - Which properties to get for the titles/revisions/pageids
                   Values (separate with '|'): info, revisions, links, langlinks, images, imageinfo, templates, categories, extlinks, categoryinfo, duplicatefiles
  list           - Which lists to get
                   Values (separate with '|'): allimages, allpages, alllinks, allcategories, allusers, backlinks, blocks, categorymembers, deletedrevs, embeddedin, imageusage, logevents, recentchanges, search, tags, usercontribs, watchlist, watchlistraw, exturlusage, users, random, protectedtitles
  meta           - Which meta data to get about the site
                   Values (separate with '|'): siteinfo, userinfo, allmessages
  generator      - Use the output of a list as the input for other prop/list/meta items
                   NOTE: generator parameter names must be prefixed with a 'g', see examples.
                   One value: links, images, templates, categories, duplicatefiles, allimages, allpages, alllinks, allcategories, backlinks, categorymembers, embeddedin, imageusage, search, watchlist, watchlistraw, exturlusage, random, protectedtitles
  redirects      - Automatically resolve redirects
  indexpageids   - Include an additional pageids section listing all returned page IDs.
  export         - Export the current revisions of all given or generated pages
  exportnowrap   - Return the export XML without wrapping it in an XML result (same format as Special:Export). Can only be used with export
Examples:
  <a href="api.php?action=query&amp;prop=revisions&amp;meta=siteinfo&amp;titles=Main%20Page&amp;rvprop=user|comment">api.php?action=query&amp;prop=revisions&amp;meta=siteinfo&amp;titles=Main%20Page&amp;rvprop=user|comment</a>
  <a href="api.php?action=query&amp;generator=allpages&amp;gapprefix=API/&amp;prop=revisions">api.php?action=query&amp;generator=allpages&amp;gapprefix=API/&amp;prop=revisions</a>

--- --- --- --- --- --- --- ---  Query: Prop  --- --- --- --- --- --- --- --- 

<b>* prop=info (in) *</b>
  Get basic page information such as namespace, title, last touched date, ...

This module requires read rights.
Parameters:
  inprop         - Which additional properties to get:
                    protection   - List the protection level of each page
                    talkid       - The page ID of the talk page for each non-talk page
                    watched      - List the watched status of each page
                    subjectid    - The page ID of the parent page for each talk page
                    url          - Gives a full URL to the page, and also an edit URL
                    readable     - Whether the user can read this page
                    preload      - Gives the text returned by EditFormPreloadText
                   Values (separate with '|'): protection, talkid, watched, subjectid, url, readable, preload
  intoken        - Request a token to perform a data-modifying action on a page
                   Values (separate with '|'): edit, delete, protect, move, block, unblock, email, import
  incontinue     - When more results are available, use this to continue
Examples:
  <a href="api.php?action=query&amp;prop=info&amp;titles=Main%20Page">api.php?action=query&amp;prop=info&amp;titles=Main%20Page</a>
  <a href="api.php?action=query&amp;prop=info&amp;inprop=protection&amp;titles=Main%20Page">api.php?action=query&amp;prop=info&amp;inprop=protection&amp;titles=Main%20Page</a>

<b>* prop=revisions (rv) *</b>
  Get revision information.
  This module may be used in several ways:
   1) Get data about a set of pages (last revision), by setting titles or pageids parameter.
   2) Get revisions for one given page, by using titles/pageids with start/end/limit params.
   3) Get data about a set of revisions by setting their IDs with revids parameter.
  All parameters marked as (enum) may only be used with a single page (#2).

This module requires read rights.
Parameters:
  rvprop         - Which properties to get for each revision.
                   Values (separate with '|'): ids, flags, timestamp, user, size, comment, parsedcomment, content, tags
                   Default: ids|timestamp|flags|comment|user
  rvlimit        - Limit how many revisions will be returned (enum)
                   No more than 500 (5000 for bots) allowed.
  rvstartid      - From which revision id to start enumeration (enum)
  rvendid        - Stop revision enumeration on this revid (enum)
  rvstart        - From which revision timestamp to start enumeration (enum)
  rvend          - Enumerate up to this timestamp (enum)
  rvdir          - Direction of enumeration - towards &quot;newer&quot; or &quot;older&quot; revisions (enum)
                   One value: newer, older
                   Default: older
  rvuser         - Only include revisions made by user
  rvexcludeuser  - Exclude revisions made by user
  rvtag          - Only list revisions tagged with this tag
  rvexpandtemplates - Expand templates in revision content
  rvgeneratexml  - Generate XML parse tree for revision content
  rvsection      - Only retrieve the content of this section
  rvtoken        - Which tokens to obtain for each revision
                   Values (separate with '|'): rollback
  rvcontinue     - When more results are available, use this to continue
  rvdiffto       - Revision ID to diff each revision to.
                   Use &quot;prev&quot;, &quot;next&quot; and &quot;cur&quot; for the previous, next and current revision respectively.
  rvdifftotext   - Text to diff each revision to. Only diffs a limited number of revisions.
                   Overrides diffto. If rvsection is set, only that section will be diffed against this text.
Examples:
  Get data with content for the last revision of titles &quot;API&quot; and &quot;Main Page&quot;:
    <a href="api.php?action=query&amp;prop=revisions&amp;titles=API|Main%20Page&amp;rvprop=timestamp|user|comment|content">api.php?action=query&amp;prop=revisions&amp;titles=API|Main%20Page&amp;rvprop=timestamp|user|comment|content</a>
  Get last 5 revisions of the &quot;Main Page&quot;:
    <a href="api.php?action=query&amp;prop=revisions&amp;titles=Main%20Page&amp;rvlimit=5&amp;rvprop=timestamp|user|comment">api.php?action=query&amp;prop=revisions&amp;titles=Main%20Page&amp;rvlimit=5&amp;rvprop=timestamp|user|comment</a>
  Get first 5 revisions of the &quot;Main Page&quot;:
    <a href="api.php?action=query&amp;prop=revisions&amp;titles=Main%20Page&amp;rvlimit=5&amp;rvprop=timestamp|user|comment&amp;rvdir=newer">api.php?action=query&amp;prop=revisions&amp;titles=Main%20Page&amp;rvlimit=5&amp;rvprop=timestamp|user|comment&amp;rvdir=newer</a>
  Get first 5 revisions of the &quot;Main Page&quot; made after 2006-05-01:
    <a href="api.php?action=query&amp;prop=revisions&amp;titles=Main%20Page&amp;rvlimit=5&amp;rvprop=timestamp|user|comment&amp;rvdir=newer&amp;rvstart=20060501000000">api.php?action=query&amp;prop=revisions&amp;titles=Main%20Page&amp;rvlimit=5&amp;rvprop=timestamp|user|comment&amp;rvdir=newer&amp;rvstart=20060501000000</a>
  Get first 5 revisions of the &quot;Main Page&quot; that were not made made by anonymous user &quot;127.0.0.1&quot;
    <a href="api.php?action=query&amp;prop=revisions&amp;titles=Main%20Page&amp;rvlimit=5&amp;rvprop=timestamp|user|comment&amp;rvexcludeuser=127.0.0.1">api.php?action=query&amp;prop=revisions&amp;titles=Main%20Page&amp;rvlimit=5&amp;rvprop=timestamp|user|comment&amp;rvexcludeuser=127.0.0.1</a>
  Get first 5 revisions of the &quot;Main Page&quot; that were made by the user &quot;MediaWiki default&quot;
    <a href="api.php?action=query&amp;prop=revisions&amp;titles=Main%20Page&amp;rvlimit=5&amp;rvprop=timestamp|user|comment&amp;rvuser=MediaWiki%20default">api.php?action=query&amp;prop=revisions&amp;titles=Main%20Page&amp;rvlimit=5&amp;rvprop=timestamp|user|comment&amp;rvuser=MediaWiki%20default</a>

<b>* prop=links (pl) *</b>
  Returns all links from the given page(s)

This module requires read rights.
Parameters:
  plnamespace    - Show links in this namespace(s) only
                   Values (separate with '|'): 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15
  pllimit        - How many links to return
                   No more than 500 (5000 for bots) allowed.
                   Default: 10
  plcontinue     - When more results are available, use this to continue
Examples:
  Get links from the [[Main Page]]:
    <a href="api.php?action=query&amp;prop=links&amp;titles=Main%20Page">api.php?action=query&amp;prop=links&amp;titles=Main%20Page</a>
  Get information about the link pages in the [[Main Page]]:
    <a href="api.php?action=query&amp;generator=links&amp;titles=Main%20Page&amp;prop=info">api.php?action=query&amp;generator=links&amp;titles=Main%20Page&amp;prop=info</a>
  Get links from the Main Page in the User and Template namespaces:
    <a href="api.php?action=query&amp;prop=links&amp;titles=Main%20Page&amp;plnamespace=2|10">api.php?action=query&amp;prop=links&amp;titles=Main%20Page&amp;plnamespace=2|10</a>
Generator:
  This module may be used as a generator

<b>* prop=langlinks (ll) *</b>
  Returns all interlanguage links from the given page(s)

This module requires read rights.
Parameters:
  lllimit        - How many langlinks to return
                   No more than 500 (5000 for bots) allowed.
                   Default: 10
  llcontinue     - When more results are available, use this to continue
Examples:
  Get interlanguage links from the [[Main Page]]:
    <a href="api.php?action=query&amp;prop=langlinks&amp;titles=Main%20Page&amp;redirects">api.php?action=query&amp;prop=langlinks&amp;titles=Main%20Page&amp;redirects</a>

<b>* prop=images (im) *</b>
  Returns all images contained on the given page(s)

This module requires read rights.
Parameters:
  imlimit        - How many images to return
                   No more than 500 (5000 for bots) allowed.
                   Default: 10
  imcontinue     - When more results are available, use this to continue
Examples:
  Get a list of images used in the [[Main Page]]:
    <a href="api.php?action=query&amp;prop=images&amp;titles=Main%20Page">api.php?action=query&amp;prop=images&amp;titles=Main%20Page</a>
  Get information about all images used in the [[Main Page]]:
    <a href="api.php?action=query&amp;generator=images&amp;titles=Main%20Page&amp;prop=info">api.php?action=query&amp;generator=images&amp;titles=Main%20Page&amp;prop=info</a>
Generator:
  This module may be used as a generator

<b>* prop=imageinfo (ii) *</b>
  Returns image information and upload history

This module requires read rights.
Parameters:
  iiprop         - What image information to get.
                   Values (separate with '|'): timestamp, user, comment, url, size, dimensions, sha1, mime, metadata, archivename, bitdepth
                   Default: timestamp|user
  iilimit        - How many image revisions to return
                   No more than 500 (5000 for bots) allowed.
                   Default: 1
  iistart        - Timestamp to start listing from
  iiend          - Timestamp to stop listing at
  iiurlwidth     - If iiprop=url is set, a URL to an image scaled to this width will be returned.
                   Only the current version of the image can be scaled.
                   Default: -1
  iiurlheight    - Similar to iiurlwidth. Cannot be used without iiurlwidth
                   Default: -1
  iicontinue     - When more results are available, use this to continue
Examples:
  <a href="api.php?action=query&amp;titles=File:Albert%20Einstein%20Head.jpg&amp;prop=imageinfo">api.php?action=query&amp;titles=File:Albert%20Einstein%20Head.jpg&amp;prop=imageinfo</a>
  <a href="api.php?action=query&amp;titles=File:Test.jpg&amp;prop=imageinfo&amp;iilimit=50&amp;iiend=20071231235959&amp;iiprop=timestamp|user|url">api.php?action=query&amp;titles=File:Test.jpg&amp;prop=imageinfo&amp;iilimit=50&amp;iiend=20071231235959&amp;iiprop=timestamp|user|url</a>

<b>* prop=templates (tl) *</b>
  Returns all templates from the given page(s)

This module requires read rights.
Parameters:
  tlnamespace    - Show templates in this namespace(s) only
                   Values (separate with '|'): 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15
  tllimit        - How many templates to return
                   No more than 500 (5000 for bots) allowed.
                   Default: 10
  tlcontinue     - When more results are available, use this to continue
Examples:
  Get templates from the [[Main Page]]:
    <a href="api.php?action=query&amp;prop=templates&amp;titles=Main%20Page">api.php?action=query&amp;prop=templates&amp;titles=Main%20Page</a>
  Get information about the template pages in the [[Main Page]]:
    <a href="api.php?action=query&amp;generator=templates&amp;titles=Main%20Page&amp;prop=info">api.php?action=query&amp;generator=templates&amp;titles=Main%20Page&amp;prop=info</a>
  Get templates from the Main Page in the User and Template namespaces:
    <a href="api.php?action=query&amp;prop=templates&amp;titles=Main%20Page&amp;tlnamespace=2|10">api.php?action=query&amp;prop=templates&amp;titles=Main%20Page&amp;tlnamespace=2|10</a>
Generator:
  This module may be used as a generator

<b>* prop=categories (cl) *</b>
  List all categories the page(s) belong to

This module requires read rights.
Parameters:
  clprop         - Which additional properties to get for each category.
                   Values (separate with '|'): sortkey, timestamp, hidden
  clshow         - Which kind of categories to show
                   Values (separate with '|'): hidden, !hidden
  cllimit        - How many categories to return
                   No more than 500 (5000 for bots) allowed.
                   Default: 10
  clcontinue     - When more results are available, use this to continue
  clcategories   - Only list these categories. Useful for checking whether a certain page is in a certain category
Examples:
  Get a list of categories [[Albert Einstein]] belongs to:
    <a href="api.php?action=query&amp;prop=categories&amp;titles=Albert%20Einstein">api.php?action=query&amp;prop=categories&amp;titles=Albert%20Einstein</a>
  Get information about all categories used in the [[Albert Einstein]]:
    <a href="api.php?action=query&amp;generator=categories&amp;titles=Albert%20Einstein&amp;prop=info">api.php?action=query&amp;generator=categories&amp;titles=Albert%20Einstein&amp;prop=info</a>
Generator:
  This module may be used as a generator

<b>* prop=extlinks (el) *</b>
  Returns all external urls (not interwikies) from the given page(s)

This module requires read rights.
Parameters:
  ellimit        - How many links to return
                   No more than 500 (5000 for bots) allowed.
                   Default: 10
  eloffset       - When more results are available, use this to continue
Examples:
  Get a list of external links on the [[Main Page]]:
    <a href="api.php?action=query&amp;prop=extlinks&amp;titles=Main%20Page">api.php?action=query&amp;prop=extlinks&amp;titles=Main%20Page</a>

<b>* prop=categoryinfo (ci) *</b>
  Returns information about the given categories

This module requires read rights.
Parameters:
  cicontinue     - When more results are available, use this to continue
Example:
  <a href="api.php?action=query&amp;prop=categoryinfo&amp;titles=Category:Foo|Category:Bar">api.php?action=query&amp;prop=categoryinfo&amp;titles=Category:Foo|Category:Bar</a>

<b>* prop=duplicatefiles (df) *</b>
  List all files that are duplicates of the given file(s).

This module requires read rights.
Parameters:
  dflimit        - How many files to return
                   No more than 500 (5000 for bots) allowed.
                   Default: 10
  dfcontinue     - When more results are available, use this to continue
Examples:
  <a href="api.php?action=query&amp;titles=File:Albert_Einstein_Head.jpg&amp;prop=duplicatefiles">api.php?action=query&amp;titles=File:Albert_Einstein_Head.jpg&amp;prop=duplicatefiles</a>
  <a href="api.php?action=query&amp;generator=allimages&amp;prop=duplicatefiles">api.php?action=query&amp;generator=allimages&amp;prop=duplicatefiles</a>
Generator:
  This module may be used as a generator

--- --- --- --- --- --- --- ---  Query: List  --- --- --- --- --- --- --- --- 

<b>* list=allimages (ai) *</b>
  Enumerate all images sequentially

This module requires read rights.
Parameters:
  aifrom         - The image title to start enumerating from.
  aiprefix       - Search for all image titles that begin with this value.
  aiminsize      - Limit to images with at least this many bytes
  aimaxsize      - Limit to images with at most this many bytes
  ailimit        - How many total images to return.
                   No more than 500 (5000 for bots) allowed.
                   Default: 10
  aidir          - The direction in which to list
                   One value: ascending, descending
                   Default: ascending
  aisha1         - SHA1 hash of image
  aisha1base36   - SHA1 hash of image in base 36 (used in MediaWiki)
  aiprop         - Which properties to get
                   Values (separate with '|'): timestamp, user, comment, url, size, dimensions, sha1, mime, metadata, archivename, bitdepth
                   Default: timestamp|url
Examples:
  Simple Use
   Show a list of images starting at the letter &quot;B&quot;
    <a href="api.php?action=query&amp;list=allimages&amp;aifrom=B">api.php?action=query&amp;list=allimages&amp;aifrom=B</a>
  Using as Generator
   Show info about 4 images starting at the letter &quot;T&quot;
    <a href="api.php?action=query&amp;generator=allimages&amp;gailimit=4&amp;gaifrom=T&amp;prop=imageinfo">api.php?action=query&amp;generator=allimages&amp;gailimit=4&amp;gaifrom=T&amp;prop=imageinfo</a>
Generator:
  This module may be used as a generator

<b>* list=allpages (ap) *</b>
  Enumerate all pages sequentially in a given namespace

This module requires read rights.
Parameters:
  apfrom         - The page title to start enumerating from.
  apprefix       - Search for all page titles that begin with this value.
  apnamespace    - The namespace to enumerate.
                   One value: 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15
                   Default: 0
  apfilterredir  - Which pages to list.
                   One value: all, redirects, nonredirects
                   Default: all
  apminsize      - Limit to pages with at least this many bytes
  apmaxsize      - Limit to pages with at most this many bytes
  apprtype       - Limit to protected pages only
                   Values (separate with '|'): edit, move
  apprlevel      - The protection level (must be used with apprtype= parameter)
                   Can be empty, or Values (separate with '|'): autoconfirmed, sysop
  apprfiltercascade - Filter protections based on cascadingness (ignored when apprtype isn't set)
                   One value: cascading, noncascading, all
                   Default: all
  aplimit        - How many total pages to return.
                   No more than 500 (5000 for bots) allowed.
                   Default: 10
  apdir          - The direction in which to list
                   One value: ascending, descending
                   Default: ascending
  apfilterlanglinks - Filter based on whether a page has langlinks
                   One value: withlanglinks, withoutlanglinks, all
                   Default: all
Examples:
  Simple Use
   Show a list of pages starting at the letter &quot;B&quot;
    <a href="api.php?action=query&amp;list=allpages&amp;apfrom=B">api.php?action=query&amp;list=allpages&amp;apfrom=B</a>
  Using as Generator
   Show info about 4 pages starting at the letter &quot;T&quot;
    <a href="api.php?action=query&amp;generator=allpages&amp;gaplimit=4&amp;gapfrom=T&amp;prop=info">api.php?action=query&amp;generator=allpages&amp;gaplimit=4&amp;gapfrom=T&amp;prop=info</a>
   Show content of first 2 non-redirect pages begining at &quot;Re&quot;
    <a href="api.php?action=query&amp;generator=allpages&amp;gaplimit=2&amp;gapfilterredir=nonredirects&amp;gapfrom=Re&amp;prop=revisions&amp;rvprop=content">api.php?action=query&amp;generator=allpages&amp;gaplimit=2&amp;gapfilterredir=nonredirects&amp;gapfrom=Re&amp;prop=revisions&amp;rvprop=content</a>
Generator:
  This module may be used as a generator

<b>* list=alllinks (al) *</b>
  Enumerate all links that point to a given namespace

This module requires read rights.
Parameters:
  alcontinue     - When more results are available, use this to continue.
  alfrom         - The page title to start enumerating from.
  alprefix       - Search for all page titles that begin with this value.
  alunique       - Only show unique links. Cannot be used with generator or prop=ids
  alprop         - What pieces of information to include
                   Values (separate with '|'): ids, title
                   Default: title
  alnamespace    - The namespace to enumerate.
                   One value: 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15
                   Default: 0
  allimit        - How many total links to return.
                   No more than 500 (5000 for bots) allowed.
                   Default: 10
Example:
  <a href="api.php?action=query&amp;list=alllinks&amp;alunique&amp;alfrom=B">api.php?action=query&amp;list=alllinks&amp;alunique&amp;alfrom=B</a>
Generator:
  This module may be used as a generator

<b>* list=allcategories (ac) *</b>
  Enumerate all categories

This module requires read rights.
Parameters:
  acfrom         - The category to start enumerating from.
  acprefix       - Search for all category titles that begin with this value.
  acdir          - Direction to sort in.
                   One value: ascending, descending
                   Default: ascending
  aclimit        - How many categories to return.
                   No more than 500 (5000 for bots) allowed.
                   Default: 10
  acprop         - Which properties to get
                   Values (separate with '|'): size, hidden
                   Default: 
Examples:
  <a href="api.php?action=query&amp;list=allcategories&amp;acprop=size">api.php?action=query&amp;list=allcategories&amp;acprop=size</a>
  <a href="api.php?action=query&amp;generator=allcategories&amp;gacprefix=List&amp;prop=info">api.php?action=query&amp;generator=allcategories&amp;gacprefix=List&amp;prop=info</a>
Generator:
  This module may be used as a generator

<b>* list=allusers (au) *</b>
  Enumerate all registered users

This module requires read rights.
Parameters:
  aufrom         - The user name to start enumerating from.
  auprefix       - Search for all page titles that begin with this value.
  augroup        - Limit users to a given group name
                   One value: bot, sysop, bureaucrat
  auprop         - What pieces of information to include.
                   `groups` property uses more server resources and may return fewer results than the limit.
                   Values (separate with '|'): blockinfo, groups, editcount, registration
  aulimit        - How many total user names to return.
                   No more than 500 (5000 for bots) allowed.
                   Default: 10
  auwitheditsonly - Only list users who have made edits
Example:
  <a href="api.php?action=query&amp;list=allusers&amp;aufrom=Y">api.php?action=query&amp;list=allusers&amp;aufrom=Y</a>

<b>* list=backlinks (bl) *</b>
  Find all pages that link to the given page

This module requires read rights.
Parameters:
  bltitle        - Title to search.
  blcontinue     - When more results are available, use this to continue.
  blnamespace    - The namespace to enumerate.
                   Values (separate with '|'): 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15
  blfilterredir  - How to filter for redirects. If set to nonredirects when blredirect is enabled, this is only applied to the second level
                   One value: all, redirects, nonredirects
                   Default: all
  bllimit        - How many total pages to return. If blredirect is enabled, limit applies to each level separately (which means you may get up to 2 * limit results).
                   No more than 500 (5000 for bots) allowed.
                   Default: 10
  blredirect     - If linking page is a redirect, find all pages that link to that redirect as well. Maximum limit is halved.
Examples:
  <a href="api.php?action=query&amp;list=backlinks&amp;bltitle=Main%20Page">api.php?action=query&amp;list=backlinks&amp;bltitle=Main%20Page</a>
  <a href="api.php?action=query&amp;generator=backlinks&amp;gbltitle=Main%20Page&amp;prop=info">api.php?action=query&amp;generator=backlinks&amp;gbltitle=Main%20Page&amp;prop=info</a>
Generator:
  This module may be used as a generator

<b>* list=blocks (bk) *</b>
  List all blocked users and IP addresses.

This module requires read rights.
Parameters:
  bkstart        - The timestamp to start enumerating from
  bkend          - The timestamp to stop enumerating at
  bkdir          - The direction in which to enumerate
                   One value: newer, older
                   Default: older
  bkids          - Pipe-separated list of block IDs to list (optional)
  bkusers        - Pipe-separated list of users to search for (optional)
  bkip           - Get all blocks applying to this IP or CIDR range, including range blocks.
                   Cannot be used together with bkusers. CIDR ranges broader than /16 are not accepted.
  bklimit        - The maximum amount of blocks to list
                   No more than 500 (5000 for bots) allowed.
                   Default: 10
  bkprop         - Which properties to get
                   Values (separate with '|'): id, user, by, timestamp, expiry, reason, range, flags
                   Default: id|user|by|timestamp|expiry|reason|flags
Examples:
  <a href="api.php?action=query&amp;list=blocks">api.php?action=query&amp;list=blocks</a>
  <a href="api.php?action=query&amp;list=blocks&amp;bkusers=Alice|Bob">api.php?action=query&amp;list=blocks&amp;bkusers=Alice|Bob</a>

<b>* list=categorymembers (cm) *</b>
  List all pages in a given category

This module requires read rights.
Parameters:
  cmtitle        - Which category to enumerate (required). Must include Category: prefix
  cmprop         - What pieces of information to include
                   Values (separate with '|'): ids, title, sortkey, timestamp
                   Default: ids|title
  cmnamespace    - Only include pages in these namespaces
                   Values (separate with '|'): 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15
  cmcontinue     - For large categories, give the value retured from previous query
  cmlimit        - The maximum number of pages to return.
                   No more than 500 (5000 for bots) allowed.
                   Default: 10
  cmsort         - Property to sort by
                   One value: sortkey, timestamp
                   Default: sortkey
  cmdir          - In which direction to sort
                   One value: asc, desc
                   Default: asc
  cmstart        - Timestamp to start listing from. Can only be used with cmsort=timestamp
  cmend          - Timestamp to end listing at. Can only be used with cmsort=timestamp
  cmstartsortkey - Sortkey to start listing from. Can only be used with cmsort=sortkey
  cmendsortkey   - Sortkey to end listing at. Can only be used with cmsort=sortkey
Examples:
  Get first 10 pages in [[Category:Physics]]:
    <a href="api.php?action=query&amp;list=categorymembers&amp;cmtitle=Category:Physics">api.php?action=query&amp;list=categorymembers&amp;cmtitle=Category:Physics</a>
  Get page info about first 10 pages in [[Category:Physics]]:
    <a href="api.php?action=query&amp;generator=categorymembers&amp;gcmtitle=Category:Physics&amp;prop=info">api.php?action=query&amp;generator=categorymembers&amp;gcmtitle=Category:Physics&amp;prop=info</a>
Generator:
  This module may be used as a generator

<b>* list=deletedrevs (dr) *</b>
  List deleted revisions.
  This module operates in three modes:
  1) List deleted revisions for the given title(s), sorted by timestamp
  2) List deleted contributions for the given user, sorted by timestamp (no titles specified)
  3) List all deleted revisions in the given namespace, sorted by title and timestamp (no titles specified, druser not set)
  Certain parameters only apply to some modes and are ignored in others.
  For instance, a parameter marked (1) only applies to mode 1 and is ignored in modes 2 and 3.

This module requires read rights.
Parameters:
  drstart        - The timestamp to start enumerating from. (1,2)
  drend          - The timestamp to stop enumerating at. (1,2)
  drdir          - The direction in which to enumerate. (1,2)
                   One value: newer, older
                   Default: older
  drfrom         - Start listing at this title (3)
  drcontinue     - When more results are available, use this to continue (3)
  drunique       - List only one revision for each page (3)
  druser         - Only list revisions by this user
  drexcludeuser  - Don't list revisions by this user
  drnamespace    - Only list pages in this namespace (3)
                   One value: 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15
                   Default: 0
  drlimit        - The maximum amount of revisions to list
                   No more than 500 (5000 for bots) allowed.
                   Default: 10
  drprop         - Which properties to get
                   Values (separate with '|'): revid, user, comment, parsedcomment, minor, len, content, token
                   Default: user|comment
Examples:
  List the last deleted revisions of Main Page and Talk:Main Page, with content (mode 1):
    <a href="api.php?action=query&amp;list=deletedrevs&amp;titles=Main%20Page|Talk:Main%20Page&amp;drprop=user|comment|content">api.php?action=query&amp;list=deletedrevs&amp;titles=Main%20Page|Talk:Main%20Page&amp;drprop=user|comment|content</a>
  List the last 50 deleted contributions by Bob (mode 2):
    <a href="api.php?action=query&amp;list=deletedrevs&amp;druser=Bob&amp;drlimit=50">api.php?action=query&amp;list=deletedrevs&amp;druser=Bob&amp;drlimit=50</a>
  List the first 50 deleted revisions in the main namespace (mode 3):
    <a href="api.php?action=query&amp;list=deletedrevs&amp;drdir=newer&amp;drlimit=50">api.php?action=query&amp;list=deletedrevs&amp;drdir=newer&amp;drlimit=50</a>
  List the first 50 deleted pages in the Talk namespace (mode 3):
    <a href="api.php?action=query&amp;list=deletedrevs&amp;drdir=newer&amp;drlimit=50&amp;drnamespace=1&amp;drunique">api.php?action=query&amp;list=deletedrevs&amp;drdir=newer&amp;drlimit=50&amp;drnamespace=1&amp;drunique</a>

<b>* list=embeddedin (ei) *</b>
  Find all pages that embed (transclude) the given title

This module requires read rights.
Parameters:
  eititle        - Title to search.
  eicontinue     - When more results are available, use this to continue.
  einamespace    - The namespace to enumerate.
                   Values (separate with '|'): 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15
  eifilterredir  - How to filter for redirects
                   One value: all, redirects, nonredirects
                   Default: all
  eilimit        - How many total pages to return.
                   No more than 500 (5000 for bots) allowed.
                   Default: 10
Examples:
  <a href="api.php?action=query&amp;list=embeddedin&amp;eititle=Template:Stub">api.php?action=query&amp;list=embeddedin&amp;eititle=Template:Stub</a>
  <a href="api.php?action=query&amp;generator=embeddedin&amp;geititle=Template:Stub&amp;prop=info">api.php?action=query&amp;generator=embeddedin&amp;geititle=Template:Stub&amp;prop=info</a>
Generator:
  This module may be used as a generator

<b>* list=imageusage (iu) *</b>
  Find all pages that use the given image title.

This module requires read rights.
Parameters:
  iutitle        - Title to search.
  iucontinue     - When more results are available, use this to continue.
  iunamespace    - The namespace to enumerate.
                   Values (separate with '|'): 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15
  iufilterredir  - How to filter for redirects. If set to nonredirects when iuredirect is enabled, this is only applied to the second level
                   One value: all, redirects, nonredirects
                   Default: all
  iulimit        - How many total pages to return. If iuredirect is enabled, limit applies to each level separately (which means you may get up to 2 * limit results).
                   No more than 500 (5000 for bots) allowed.
                   Default: 10
  iuredirect     - If linking page is a redirect, find all pages that link to that redirect as well. Maximum limit is halved.
Examples:
  <a href="api.php?action=query&amp;list=imageusage&amp;iutitle=File:Albert%20Einstein%20Head.jpg">api.php?action=query&amp;list=imageusage&amp;iutitle=File:Albert%20Einstein%20Head.jpg</a>
  <a href="api.php?action=query&amp;generator=imageusage&amp;giutitle=File:Albert%20Einstein%20Head.jpg&amp;prop=info">api.php?action=query&amp;generator=imageusage&amp;giutitle=File:Albert%20Einstein%20Head.jpg&amp;prop=info</a>
Generator:
  This module may be used as a generator

<b>* list=logevents (le) *</b>
  Get events from logs.

This module requires read rights.
Parameters:
  leprop         - Which properties to get
                   Values (separate with '|'): ids, title, type, user, timestamp, comment, parsedcomment, details, tags
                   Default: ids|title|type|user|timestamp|comment|details
  letype         - Filter log entries to only this type(s)
                   Can be empty, or One value: block, protect, rights, delete, upload, move, import, patrol, merge, suppress, usermerge, newusers
  lestart        - The timestamp to start enumerating from.
  leend          - The timestamp to end enumerating.
  ledir          - In which direction to enumerate.
                   One value: newer, older
                   Default: older
  leuser         - Filter entries to those made by the given user.
  letitle        - Filter entries to those related to a page.
  letag          - Only list event entries tagged with this tag.
  lelimit        - How many total event entries to return.
                   No more than 500 (5000 for bots) allowed.
                   Default: 10
Example:
  <a href="api.php?action=query&amp;list=logevents">api.php?action=query&amp;list=logevents</a>

<b>* list=recentchanges (rc) *</b>
  Enumerate recent changes

This module requires read rights.
Parameters:
  rcstart        - The timestamp to start enumerating from.
  rcend          - The timestamp to end enumerating.
  rcdir          - In which direction to enumerate.
                   One value: newer, older
                   Default: older
  rcnamespace    - Filter log entries to only this namespace(s)
                   Values (separate with '|'): 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15
  rcuser         - Only list changes by this user
  rcexcludeuser  - Don't list changes by this user
  rctag          - Only list changes tagged with this tag.
  rcprop         - Include additional pieces of information
                   Values (separate with '|'): user, comment, parsedcomment, flags, timestamp, title, ids, sizes, redirect, patrolled, loginfo, tags
                   Default: title|timestamp|ids
  rctoken        - Which tokens to obtain for each change
                   Values (separate with '|'): patrol
  rcshow         - Show only items that meet this criteria.
                   For example, to see only minor edits done by logged-in users, set show=minor|!anon
                   Values (separate with '|'): minor, !minor, bot, !bot, anon, !anon, redirect, !redirect, patrolled, !patrolled
  rclimit        - How many total changes to return.
                   No more than 500 (5000 for bots) allowed.
                   Default: 10
  rctype         - Which types of changes to show.
                   Values (separate with '|'): edit, new, log
Example:
  <a href="api.php?action=query&amp;list=recentchanges">api.php?action=query&amp;list=recentchanges</a>

<b>* list=search (sr) *</b>
  Perform a full text search

This module requires read rights.
Parameters:
  srsearch       - Search for all page titles (or content) that has this value.
  srnamespace    - The namespace(s) to enumerate.
                   Values (separate with '|'): 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15
                   Default: 0
  srwhat         - Search inside the text or titles.
                   One value: title, text
  srinfo         - What metadata to return.
                   Values (separate with '|'): totalhits, suggestion
                   Default: totalhits|suggestion
  srprop         - What properties to return.
                   Values (separate with '|'): size, wordcount, timestamp, snippet
                   Default: size|wordcount|timestamp|snippet
  srredirects    - Include redirect pages in the search.
  sroffset       - Use this value to continue paging (return by query)
                   Default: 0
  srlimit        - How many total pages to return.
                   No more than 50 (500 for bots) allowed.
                   Default: 10
Examples:
  <a href="api.php?action=query&amp;list=search&amp;srsearch=meaning">api.php?action=query&amp;list=search&amp;srsearch=meaning</a>
  <a href="api.php?action=query&amp;list=search&amp;srwhat=text&amp;srsearch=meaning">api.php?action=query&amp;list=search&amp;srwhat=text&amp;srsearch=meaning</a>
  <a href="api.php?action=query&amp;generator=search&amp;gsrsearch=meaning&amp;prop=info">api.php?action=query&amp;generator=search&amp;gsrsearch=meaning&amp;prop=info</a>
Generator:
  This module may be used as a generator

<b>* list=tags (tg) *</b>
  List change tags.

This module requires read rights.
Parameters:
  tgcontinue     - When more results are available, use this to continue
  tglimit        - The maximum number of tags to list
                   No more than 500 (5000 for bots) allowed.
                   Default: 10
  tgprop         - Which properties to get
                   Values (separate with '|'): name, displayname, description, hitcount
                   Default: name
Example:
  <a href="api.php?action=query&amp;list=tags&amp;tgprop=displayname|description|hitcount">api.php?action=query&amp;list=tags&amp;tgprop=displayname|description|hitcount</a>

<b>* list=usercontribs (uc) *</b>
  Get all edits by a user

This module requires read rights.
Parameters:
  uclimit        - The maximum number of contributions to return.
                   No more than 500 (5000 for bots) allowed.
                   Default: 10
  ucstart        - The start timestamp to return from.
  ucend          - The end timestamp to return to.
  uccontinue     - When more results are available, use this to continue.
  ucuser         - The user to retrieve contributions for.
  ucuserprefix   - Retrieve contibutions for all users whose names begin with this value. Overrides ucuser.
  ucdir          - The direction to search (older or newer).
                   One value: newer, older
                   Default: older
  ucnamespace    - Only list contributions in these namespaces
                   Values (separate with '|'): 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15
  ucprop         - Include additional pieces of information
                   Values (separate with '|'): ids, title, timestamp, comment, parsedcomment, size, flags, patrolled, tags
                   Default: ids|title|timestamp|comment|size|flags
  ucshow         - Show only items that meet this criteria, e.g. non minor edits only: show=!minor
                   NOTE: if show=patrolled or show=!patrolled is set, revisions older than $wgRCMaxAge won't be shown
                   Values (separate with '|'): minor, !minor, patrolled, !patrolled
  uctag          - Only list revisions tagged with this tag
Examples:
  <a href="api.php?action=query&amp;list=usercontribs&amp;ucuser=YurikBot">api.php?action=query&amp;list=usercontribs&amp;ucuser=YurikBot</a>
  <a href="api.php?action=query&amp;list=usercontribs&amp;ucuserprefix=217.121.114.">api.php?action=query&amp;list=usercontribs&amp;ucuserprefix=217.121.114.</a>

<b>* list=watchlist (wl) *</b>
  Get all recent changes to pages in the logged in user's watchlist

This module requires read rights.
Parameters:
  wlallrev       - Include multiple revisions of the same page within given timeframe.
  wlstart        - The timestamp to start enumerating from.
  wlend          - The timestamp to end enumerating.
  wlnamespace    - Filter changes to only the given namespace(s).
                   Values (separate with '|'): 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15
  wluser         - Only list changes by this user
  wlexcludeuser  - Don't list changes by this user
  wldir          - In which direction to enumerate pages.
                   One value: newer, older
                   Default: older
  wllimit        - How many total results to return per request.
                   No more than 500 (5000 for bots) allowed.
                   Default: 10
  wlprop         - Which additional items to get (non-generator mode only).
                   Values (separate with '|'): ids, title, flags, user, comment, parsedcomment, timestamp, patrol, sizes, notificationtimestamp
                   Default: ids|title|flags
  wlshow         - Show only items that meet this criteria.
                   For example, to see only minor edits done by logged-in users, set show=minor|!anon
                   Values (separate with '|'): minor, !minor, bot, !bot, anon, !anon, patrolled, !patrolled
  wlowner        - The name of the user whose watchlist you'd like to access
  wltoken        - Give a security token (settable in preferences) to allow access to another user's watchlist
Examples:
  <a href="api.php?action=query&amp;list=watchlist">api.php?action=query&amp;list=watchlist</a>
  <a href="api.php?action=query&amp;list=watchlist&amp;wlprop=ids|title|timestamp|user|comment">api.php?action=query&amp;list=watchlist&amp;wlprop=ids|title|timestamp|user|comment</a>
  <a href="api.php?action=query&amp;list=watchlist&amp;wlallrev&amp;wlprop=ids|title|timestamp|user|comment">api.php?action=query&amp;list=watchlist&amp;wlallrev&amp;wlprop=ids|title|timestamp|user|comment</a>
  <a href="api.php?action=query&amp;generator=watchlist&amp;prop=info">api.php?action=query&amp;generator=watchlist&amp;prop=info</a>
  <a href="api.php?action=query&amp;generator=watchlist&amp;gwlallrev&amp;prop=revisions&amp;rvprop=timestamp|user">api.php?action=query&amp;generator=watchlist&amp;gwlallrev&amp;prop=revisions&amp;rvprop=timestamp|user</a>
  <a href="api.php?action=query&amp;list=watchlist&amp;wlowner=Bob_Smith&amp;wltoken=d8d562e9725ea1512894cdab28e5ceebc7f20237">api.php?action=query&amp;list=watchlist&amp;wlowner=Bob_Smith&amp;wltoken=d8d562e9725ea1512894cdab28e5ceebc7f20237</a>
Generator:
  This module may be used as a generator

<b>* list=watchlistraw (wr) *</b>
  Get all pages on the logged in user's watchlist

This module requires read rights.
Parameters:
  wrcontinue     - When more results are available, use this to continue
  wrnamespace    - Only list pages in the given namespace(s).
                   Values (separate with '|'): 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15
  wrlimit        - How many total results to return per request.
                   No more than 500 (5000 for bots) allowed.
                   Default: 10
  wrprop         - Which additional properties to get (non-generator mode only).
                   Values (separate with '|'): changed
  wrshow         - Only list items that meet these criteria.
                   Values (separate with '|'): changed, !changed
Examples:
  <a href="api.php?action=query&amp;list=watchlistraw">api.php?action=query&amp;list=watchlistraw</a>
  <a href="api.php?action=query&amp;generator=watchlistraw&amp;gwrshow=changed&amp;prop=revisions">api.php?action=query&amp;generator=watchlistraw&amp;gwrshow=changed&amp;prop=revisions</a>
Generator:
  This module may be used as a generator

<b>* list=exturlusage (eu) *</b>
  Enumerate pages that contain a given URL

This module requires read rights.
Parameters:
  euprop         - What pieces of information to include
                   Values (separate with '|'): ids, title, url
                   Default: ids|title|url
  euoffset       - Used for paging. Use the value returned for &quot;continue&quot;
  euprotocol     - Protocol of the url. If empty and euquery set, the protocol is http.
                   Leave both this and euquery empty to list all external links
                   Can be empty, or One value: http, https, ftp, irc, gopher, telnet, nntp, worldwind, mailto, news, svn
                   Default: 
  euquery        - Search string without protocol. See [[Special:LinkSearch]]. Leave empty to list all external links
  eunamespace    - The page namespace(s) to enumerate.
                   Values (separate with '|'): 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15
  eulimit        - How many pages to return.
                   No more than 500 (5000 for bots) allowed.
                   Default: 10
Example:
  <a href="api.php?action=query&amp;list=exturlusage&amp;euquery=www.mediawiki.org">api.php?action=query&amp;list=exturlusage&amp;euquery=www.mediawiki.org</a>
Generator:
  This module may be used as a generator

<b>* list=users (us) *</b>
  Get information about a list of users

This module requires read rights.
Parameters:
  usprop         - What pieces of information to include
                     blockinfo    - tags if the user is blocked, by whom, and for what reason
                     groups       - lists all the groups the user belongs to
                     editcount    - adds the user's edit count
                     registration - adds the user's registration timestamp
                     emailable    - tags if the user can and wants to receive e-mail through [[Special:Emailuser]]
                     gender       - tags the gender of the user. Returns &quot;male&quot;, &quot;female&quot;, or &quot;unknown&quot;
                   Values (separate with '|'): blockinfo, groups, editcount, registration, emailable, gender
  ususers        - A list of users to obtain the same information for
  ustoken        - Which tokens to obtain for each user
                   Values (separate with '|'): userrights
Example:
  <a href="api.php?action=query&amp;list=users&amp;ususers=brion|TimStarling&amp;usprop=groups|editcount|gender">api.php?action=query&amp;list=users&amp;ususers=brion|TimStarling&amp;usprop=groups|editcount|gender</a>

<b>* list=random (rn) *</b>
  Get a set of random pages
  NOTE: Pages are listed in a fixed sequence, only the starting point is random. This means that if, for example, &quot;Main Page&quot; is the first 
        random page on your list, &quot;List of fictional monkeys&quot; will <b>*always*</b> be second, &quot;List of people on stamps of Vanuatu&quot; third, etc.
  NOTE: If the number of pages in the namespace is lower than rnlimit, you will get fewer pages. You will not get the same page twice.

This module requires read rights.
Parameters:
  rnnamespace    - Return pages in these namespaces only
                   Values (separate with '|'): 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15
  rnlimit        - Limit how many random pages will be returned
                   No more than 10 (20 for bots) allowed.
                   Default: 1
  rnredirect     - Load a random redirect instead of a random page
Example:
  <a href="api.php?action=query&amp;list=random&amp;rnnamespace=0&amp;rnlimit=2">api.php?action=query&amp;list=random&amp;rnnamespace=0&amp;rnlimit=2</a>
Generator:
  This module may be used as a generator

<b>* list=protectedtitles (pt) *</b>
  List all titles protected from creation

This module requires read rights.
Parameters:
  ptnamespace    - Only list titles in these namespaces
                   Values (separate with '|'): 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15
  ptlevel        - Only list titles with these protection levels
                   Values (separate with '|'): autoconfirmed, sysop
  ptlimit        - How many total pages to return.
                   No more than 500 (5000 for bots) allowed.
                   Default: 10
  ptdir          - The direction in which to list
                   One value: older, newer
                   Default: older
  ptstart        - Start listing at this protection timestamp
  ptend          - Stop listing at this protection timestamp
  ptprop         - Which properties to get
                   Values (separate with '|'): timestamp, user, comment, parsedcomment, expiry, level
                   Default: timestamp|level
Example:
  <a href="api.php?action=query&amp;list=protectedtitles">api.php?action=query&amp;list=protectedtitles</a>
Generator:
  This module may be used as a generator

--- --- --- --- --- --- --- ---  Query: Meta  --- --- --- --- --- --- --- --- 

<b>* meta=siteinfo (si) *</b>
  Return general information about the site.

This module requires read rights.
Parameters:
  siprop         - Which sysinfo properties to get:
                    general      - Overall system information
                    namespaces   - List of registered namespaces and their canonical names
                    namespacealiases - List of registered namespace aliases
                    specialpagealiases - List of special page aliases
                    magicwords   - List of magic words and their aliases
                    statistics   - Returns site statistics
                    interwikimap - Returns interwiki map (optionally filtered)
                    dbrepllag    - Returns database server with the highest replication lag
                    usergroups   - Returns user groups and the associated permissions
                    extensions   - Returns extensions installed on the wiki
                    fileextensions - Returns list of file extensions allowed to be uploaded
                    rightsinfo   - Returns wiki rights (license) information if available
                    languages    - Returns a list of languages MediaWiki supports
                   Values (separate with '|'): general, namespaces, namespacealiases, specialpagealiases, magicwords, interwikimap, dbrepllag, statistics, usergroups, extensions, fileextensions, rightsinfo, languages
                   Default: general
  sifilteriw     - Return only local or only nonlocal entries of the interwiki map
                   One value: local, !local
  sishowalldb    - List all database servers, not just the one lagging the most
  sinumberingroup - Lists the number of users in user groups
Examples:
  <a href="api.php?action=query&amp;meta=siteinfo&amp;siprop=general|namespaces|namespacealiases|statistics">api.php?action=query&amp;meta=siteinfo&amp;siprop=general|namespaces|namespacealiases|statistics</a>
  <a href="api.php?action=query&amp;meta=siteinfo&amp;siprop=interwikimap&amp;sifilteriw=local">api.php?action=query&amp;meta=siteinfo&amp;siprop=interwikimap&amp;sifilteriw=local</a>
  <a href="api.php?action=query&amp;meta=siteinfo&amp;siprop=dbrepllag&amp;sishowalldb">api.php?action=query&amp;meta=siteinfo&amp;siprop=dbrepllag&amp;sishowalldb</a>

<b>* meta=userinfo (ui) *</b>
  Get information about the current user

This module requires read rights.
Parameters:
  uiprop         - What pieces of information to include
                     blockinfo  - tags if the current user is blocked, by whom, and for what reason
                     hasmsg     - adds a tag &quot;message&quot; if the current user has pending messages
                     groups     - lists all the groups the current user belongs to
                     rights     - lists all the rights the current user has
                     changeablegroups - lists the groups the current user can add to and remove from
                     options    - lists all preferences the current user has set
                     editcount  - adds the current user's edit count
                     ratelimits - lists all rate limits applying to the current user
                   Values (separate with '|'): blockinfo, hasmsg, groups, rights, changeablegroups, options, preferencestoken, editcount, ratelimits, email
Examples:
  <a href="api.php?action=query&amp;meta=userinfo">api.php?action=query&amp;meta=userinfo</a>
  <a href="api.php?action=query&amp;meta=userinfo&amp;uiprop=blockinfo|groups|rights|hasmsg">api.php?action=query&amp;meta=userinfo&amp;uiprop=blockinfo|groups|rights|hasmsg</a>

<b>* meta=allmessages (am) *</b>
  Return messages from this site.

This module requires read rights.
Parameters:
  ammessages     - Which messages to output. &quot;*&quot; means all messages
                   Default: *
  amprop         - Which properties to get
                   Values (separate with '|'): default
  amenableparser - Set to enable parser, will preprocess the wikitext of message
                   Will substitute magic words, handle templates etc.
  amargs         - Arguments to be substituted into message
  amfilter       - Return only messages that contain this string
  amlang         - Return messages in this language
  amfrom         - Return messages starting at this message
Examples:
  <a href="api.php?action=query&amp;meta=allmessages&amp;amfilter=ipb-">api.php?action=query&amp;meta=allmessages&amp;amfilter=ipb-</a>
  <a href="api.php?action=query&amp;meta=allmessages&amp;ammessages=august|mainpage&amp;amlang=de">api.php?action=query&amp;meta=allmessages&amp;ammessages=august|mainpage&amp;amlang=de</a>


<b>*** *** *** *** *** *** *** *** *** ***  Modules: continuation  *** *** *** *** *** *** *** *** *** ***</b> 


<b>* action=expandtemplates *</b>
  This module expand all templates in wikitext

This module requires read rights.
Parameters:
  title          - Title of page
                   Default: API
  text           - Wikitext to convert
  generatexml    - Generate XML parse tree
Example:
  <a href="api.php?action=expandtemplates&amp;text={{Project:Sandbox}}">api.php?action=expandtemplates&amp;text={{Project:Sandbox}}</a>

<b>* action=parse *</b>
  This module parses wikitext and returns parser output

This module requires read rights.
Parameters:
  title          - Title of page the text belongs to
                   Default: API
  text           - Wikitext to parse
  summary        - Summary to parse
  page           - Parse the content of this page. Cannot be used together with text and title
  redirects      - If the page parameter is set to a redirect, resolve it
  oldid          - Parse the content of this revision. Overrides page
  prop           - Which pieces of information to get.
                   NOTE: Section tree is only generated if there are more than 4 sections, or if the __TOC__ keyword is present
                   Values (separate with '|'): text, langlinks, categories, links, templates, images, externallinks, sections, revid, displaytitle, headitems, headhtml
                   Default: text|langlinks|categories|links|templates|images|externallinks|sections|revid|displaytitle
  pst            - Do a pre-save transform on the input before parsing it.
                   Ignored if page or oldid is used.
  onlypst        - Do a PST on the input, but don't parse it.
                   Returns PSTed wikitext. Ignored if page or oldid is used.
Example:
  <a href="api.php?action=parse&amp;text={{Project:Sandbox}}">api.php?action=parse&amp;text={{Project:Sandbox}}</a>

<b>* action=opensearch *</b>
  This module implements OpenSearch protocol

This module requires read rights.
Parameters:
  search         - Search string
  limit          - Maximum amount of results to return
                   No more than 100 (100 for bots) allowed.
                   Default: 10
  namespace      - Namespaces to search
                   Values (separate with '|'): 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15
                   Default: 0
  suggest        - Do nothing if $wgEnableOpenSearchSuggest is false
Example:
  <a href="api.php?action=opensearch&amp;search=Te">api.php?action=opensearch&amp;search=Te</a>

<b>* action=feedwatchlist *</b>
  This module returns a watchlist feed

This module requires read rights.
Parameters:
  feedformat     - The format of the feed
                   One value: rss, atom
                   Default: rss
  hours          - List pages modified within this many hours from now
                   The value must be between 1 and 72
                   Default: 24
  allrev         - Include multiple revisions of the same page within given timeframe.
  wlowner        - The user whose watchlist you want (must be accompanied by wltoken if it's not you)
  wltoken        - Security token that requested user set in their preferences
Example:
  <a href="api.php?action=feedwatchlist">api.php?action=feedwatchlist</a>

<b>* action=help *</b>
  Display this help screen.

<b>* action=paraminfo *</b>
  Obtain information about certain API parameters
Parameters:
  modules        - List of module names (value of the action= parameter)
  querymodules   - List of query module names (value of prop=, meta= or list= parameter)
  mainmodule     - Get information about the main (top-level) module as well
  pagesetmodule  - Get information about the pageset module (providing titles= and friends) as well
Example:
  <a href="api.php?action=paraminfo&amp;modules=parse&amp;querymodules=allpages|siteinfo">api.php?action=paraminfo&amp;modules=parse&amp;querymodules=allpages|siteinfo</a>

<b>* action=purge *</b>
  Purge the cache for the given titles.

This module requires read rights.
This module requires write rights.
This module only accepts POST requests.
Parameters:
  titles         - A list of titles
Example:
  <a href="api.php?action=purge&amp;titles=Main_Page|API">api.php?action=purge&amp;titles=Main_Page|API</a>

<b>* action=rollback *</b>
  Undo the last edit to the page. If the last user who edited the page made multiple edits in a row,
  they will all be rolled back.

This module requires read rights.
This module requires write rights.
This module only accepts POST requests.
Parameters:
  title          - Title of the page you want to rollback.
  user           - Name of the user whose edits are to be rolled back. If set incorrectly, you'll get a badtoken error.
  token          - A rollback token previously retrieved through prop=revisions
  summary        - Custom edit summary. If not set, default summary will be used.
  markbot        - Mark the reverted edits and the revert as bot edits
Examples:
  <a href="api.php?action=rollback&amp;title=Main%20Page&amp;user=Catrope&amp;token=123ABC">api.php?action=rollback&amp;title=Main%20Page&amp;user=Catrope&amp;token=123ABC</a>
  <a href="api.php?action=rollback&amp;title=Main%20Page&amp;user=217.121.114.116&amp;token=123ABC&amp;summary=Reverting%20vandalism&amp;markbot=1">api.php?action=rollback&amp;title=Main%20Page&amp;user=217.121.114.116&amp;token=123ABC&amp;summary=Reverting%20vandalism&amp;markbot=1</a>

<b>* action=delete *</b>
  Delete a page.

This module requires read rights.
This module requires write rights.
This module only accepts POST requests.
Parameters:
  title          - Title of the page you want to delete. Cannot be used together with pageid
  pageid         - Page ID of the page you want to delete. Cannot be used together with title
  token          - A delete token previously retrieved through prop=info
  reason         - Reason for the deletion. If not set, an automatically generated reason will be used.
  watch          - Add the page to your watchlist
  unwatch        - Remove the page from your watchlist
  oldimage       - The name of the old image to delete as provided by iiprop=archivename
Examples:
  <a href="api.php?action=delete&amp;title=Main%20Page&amp;token=123ABC">api.php?action=delete&amp;title=Main%20Page&amp;token=123ABC</a>
  <a href="api.php?action=delete&amp;title=Main%20Page&amp;token=123ABC&amp;reason=Preparing%20for%20move">api.php?action=delete&amp;title=Main%20Page&amp;token=123ABC&amp;reason=Preparing%20for%20move</a>

<b>* action=undelete *</b>
  Restore certain revisions of a deleted page. A list of deleted revisions (including timestamps) can be
  retrieved through list=deletedrevs

This module requires read rights.
This module requires write rights.
This module only accepts POST requests.
Parameters:
  title          - Title of the page you want to restore.
  token          - An undelete token previously retrieved through list=deletedrevs
  reason         - Reason for restoring (optional)
                   Default: 
  timestamps     - Timestamps of the revisions to restore. If not set, all revisions will be restored.
Examples:
  <a href="api.php?action=undelete&amp;title=Main%20Page&amp;token=123ABC&amp;reason=Restoring%20main%20page">api.php?action=undelete&amp;title=Main%20Page&amp;token=123ABC&amp;reason=Restoring%20main%20page</a>
  <a href="api.php?action=undelete&amp;title=Main%20Page&amp;token=123ABC&amp;timestamps=20070703220045|20070702194856">api.php?action=undelete&amp;title=Main%20Page&amp;token=123ABC&amp;timestamps=20070703220045|20070702194856</a>

<b>* action=protect *</b>
  Change the protection level of a page.

This module requires read rights.
This module requires write rights.
This module only accepts POST requests.
Parameters:
  title          - Title of the page you want to (un)protect.
  token          - A protect token previously retrieved through prop=info
  protections    - Pipe-separated list of protection levels, formatted action=group (e.g. edit=sysop)
  expiry         - Expiry timestamps. If only one timestamp is set, it'll be used for all protections.
                   Use 'infinite', 'indefinite' or 'never', for a neverexpiring protection.
                   Default: infinite
  reason         - Reason for (un)protecting (optional)
                   Default: 
  cascade        - Enable cascading protection (i.e. protect pages included in this page)
                   Ignored if not all protection levels are 'sysop' or 'protect'
  watch          - If set, add the page being (un)protected to your watchlist
Examples:
  <a href="api.php?action=protect&amp;title=Main%20Page&amp;token=123ABC&amp;protections=edit=sysop|move=sysop&amp;cascade&amp;expiry=20070901163000|never">api.php?action=protect&amp;title=Main%20Page&amp;token=123ABC&amp;protections=edit=sysop|move=sysop&amp;cascade&amp;expiry=20070901163000|never</a>
  <a href="api.php?action=protect&amp;title=Main%20Page&amp;token=123ABC&amp;protections=edit=all|move=all&amp;reason=Lifting%20restrictions">api.php?action=protect&amp;title=Main%20Page&amp;token=123ABC&amp;protections=edit=all|move=all&amp;reason=Lifting%20restrictions</a>

<b>* action=block *</b>
  Block a user.

This module requires read rights.
This module requires write rights.
This module only accepts POST requests.
Parameters:
  user           - Username, IP address or IP range you want to block
  token          - A block token previously obtained through the gettoken parameter or prop=info
  gettoken       - If set, a block token will be returned, and no other action will be taken
  expiry         - Relative expiry time, e.g. '5 months' or '2 weeks'. If set to 'infinite', 'indefinite' or 'never', the block will never expire.
                   Default: never
  reason         - Reason for block (optional)
  anononly       - Block anonymous users only (i.e. disable anonymous edits for this IP)
  nocreate       - Prevent account creation
  autoblock      - Automatically block the last used IP address, and any subsequent IP addresses they try to login from
  noemail        - Prevent user from sending e-mail through the wiki. (Requires the &quot;blockemail&quot; right.)
  hidename       - Hide the username from the block log. (Requires the &quot;hideuser&quot; right.)
  allowusertalk  - Allow the user to edit their own talk page (depends on $wgBlockAllowsUTEdit)
  reblock        - If the user is already blocked, overwrite the existing block
Examples:
  <a href="api.php?action=block&amp;user=123.5.5.12&amp;expiry=3%20days&amp;reason=First%20strike">api.php?action=block&amp;user=123.5.5.12&amp;expiry=3%20days&amp;reason=First%20strike</a>
  <a href="api.php?action=block&amp;user=Vandal&amp;expiry=never&amp;reason=Vandalism&amp;nocreate&amp;autoblock&amp;noemail">api.php?action=block&amp;user=Vandal&amp;expiry=never&amp;reason=Vandalism&amp;nocreate&amp;autoblock&amp;noemail</a>

<b>* action=unblock *</b>
  Unblock a user.

This module requires read rights.
This module requires write rights.
This module only accepts POST requests.
Parameters:
  id             - ID of the block you want to unblock (obtained through list=blocks). Cannot be used together with user
  user           - Username, IP address or IP range you want to unblock. Cannot be used together with id
  token          - An unblock token previously obtained through the gettoken parameter or prop=info
  gettoken       - If set, an unblock token will be returned, and no other action will be taken
  reason         - Reason for unblock (optional)
Examples:
  <a href="api.php?action=unblock&amp;id=105">api.php?action=unblock&amp;id=105</a>
  <a href="api.php?action=unblock&amp;user=Bob&amp;reason=Sorry%20Bob">api.php?action=unblock&amp;user=Bob&amp;reason=Sorry%20Bob</a>

<b>* action=move *</b>
  Move a page.

This module requires read rights.
This module requires write rights.
This module only accepts POST requests.
Parameters:
  from           - Title of the page you want to move. Cannot be used together with fromid.
  fromid         - Page ID of the page you want to move. Cannot be used together with from.
  to             - Title you want to rename the page to.
  token          - A move token previously retrieved through prop=info
  reason         - Reason for the move (optional).
  movetalk       - Move the talk page, if it exists.
  movesubpages   - Move subpages, if applicable
  noredirect     - Don't create a redirect
  watch          - Add the page and the redirect to your watchlist
  unwatch        - Remove the page and the redirect from your watchlist
  ignorewarnings - Ignore any warnings
Example:
  <a href="api.php?action=move&amp;from=Exampel&amp;to=Example&amp;token=123ABC&amp;reason=Misspelled%20title&amp;movetalk&amp;noredirect">api.php?action=move&amp;from=Exampel&amp;to=Example&amp;token=123ABC&amp;reason=Misspelled%20title&amp;movetalk&amp;noredirect</a>

<b>* action=edit *</b>
  Create and edit pages.

This module requires read rights.
This module requires write rights.
This module only accepts POST requests.
Parameters:
  title          - Page title
  section        - Section number. 0 for the top section, 'new' for a new section
  text           - Page content
  token          - Edit token. You can get one of these through prop=info
  summary        - Edit summary. Also section title when section=new
  minor          - Minor edit
  notminor       - Non-minor edit
  bot            - Mark this edit as bot
  basetimestamp  - Timestamp of the base revision (gotten through prop=revisions&amp;rvprop=timestamp).
                   Used to detect edit conflicts; leave unset to ignore conflicts.
  starttimestamp - Timestamp when you obtained the edit token.
                   Used to detect edit conflicts; leave unset to ignore conflicts.
  recreate       - Override any errors about the article having been deleted in the meantime
  createonly     - Don't edit the page if it exists already
  nocreate       - Throw an error if the page doesn't exist
  captchaword    - Answer to the CAPTCHA
  captchaid      - CAPTCHA ID from previous request
  watch          - DEPRECATED! Add the page to your watchlist
  unwatch        - DEPRECATED! Remove the page from your watchlist
  watchlist      - Unconditionally add or remove the page from your watchlist, use preferences or do not change watch
                   One value: watch, unwatch, preferences, nochange
                   Default: preferences
  md5            - The MD5 hash of the text parameter, or the prependtext and appendtext parameters concatenated.
                   If set, the edit won't be done unless the hash is correct
  prependtext    - Add this text to the beginning of the page. Overrides text.
  appendtext     - Add this text to the end of the page. Overrides text
  undo           - Undo this revision. Overrides text, prependtext and appendtext
  undoafter      - Undo all revisions from undo to this one. If not set, just undo one revision
Examples:
  Edit a page (anonymous user):
      <a href="api.php?action=edit&amp;title=Test&amp;summary=test%20summary&amp;text=article%20content&amp;basetimestamp=20070824123454&amp;token=%2B\">api.php?action=edit&amp;title=Test&amp;summary=test%20summary&amp;text=article%20content&amp;basetimestamp=20070824123454&amp;token=%2B\</a>
  Prepend __NOTOC__ to a page (anonymous user):
      <a href="api.php?action=edit&amp;title=Test&amp;summary=NOTOC&amp;minor&amp;prependtext=__NOTOC__%0A&amp;basetimestamp=20070824123454&amp;token=%2B\">api.php?action=edit&amp;title=Test&amp;summary=NOTOC&amp;minor&amp;prependtext=__NOTOC__%0A&amp;basetimestamp=20070824123454&amp;token=%2B\</a>
  Undo r13579 through r13585 with autosummary(anonymous user):
      <a href="api.php?action=edit&amp;title=Test&amp;undo=13585&amp;undoafter=13579&amp;basetimestamp=20070824123454&amp;token=%2B\">api.php?action=edit&amp;title=Test&amp;undo=13585&amp;undoafter=13579&amp;basetimestamp=20070824123454&amp;token=%2B\</a>

<b>* action=upload *</b>
  Upload a file, or get the status of pending uploads. Several methods are available:
   * Upload file contents directly, using the &quot;file&quot; parameter
   * Have the MediaWiki server fetch a file from a URL, using the &quot;url&quot; parameter
   * Complete an earlier upload that failed due to warnings, using the &quot;sessionkey&quot; parameter
  Note that the HTTP POST must be done as a file upload (i.e. using multipart/form-data) when
  sending the &quot;file&quot;. Note also that queries using session keys must be
  done in the same login session as the query that originally returned the key (i.e. do not
  log out and then log back in). Also you must get and send an edit token before doing any upload stuff.

This module requires read rights.
This module requires write rights.
This module only accepts POST requests.
Parameters:
  filename       - Target filename
  comment        - Upload comment. Also used as the initial page text for new files if &quot;text&quot; is not specified
                   Default: 
  text           - Initial page text for new files
  token          - Edit token. You can get one of these through prop=info
  watch          - Watch the page
  ignorewarnings - Ignore any warnings
  file           - File contents
  url            - Url to fetch the file from
  sessionkey     - Session key returned by a previous upload that failed due to warnings
Examples:
  Upload from a URL:
      <a href="api.php?action=upload&amp;filename=Wiki.png&amp;url=http%3A//upload.wikimedia.org/wikipedia/en/b/bc/Wiki.png">api.php?action=upload&amp;filename=Wiki.png&amp;url=http%3A//upload.wikimedia.org/wikipedia/en/b/bc/Wiki.png</a>
  Complete an upload that failed due to warnings:
      <a href="api.php?action=upload&amp;filename=Wiki.png&amp;sessionkey=sessionkey&amp;ignorewarnings=1">api.php?action=upload&amp;filename=Wiki.png&amp;sessionkey=sessionkey&amp;ignorewarnings=1</a>

<b>* action=emailuser *</b>
  Email a user.

This module requires read rights.
This module requires write rights.
This module only accepts POST requests.
Parameters:
  target         - User to send email to
  subject        - Subject header
  text           - Mail body
  token          - A token previously acquired via prop=info
  ccme           - Send a copy of this mail to me
Example:
  <a href="api.php?action=emailuser&amp;target=WikiSysop&amp;text=Content">api.php?action=emailuser&amp;target=WikiSysop&amp;text=Content</a>

<b>* action=watch *</b>
  Add or remove a page from/to the current user's watchlist

This module requires read rights.
This module requires write rights.
Parameters:
  title          - The page to (un)watch
  unwatch        - If set the page will be unwatched rather than watched
Examples:
  <a href="api.php?action=watch&amp;title=Main_Page">api.php?action=watch&amp;title=Main_Page</a>
  <a href="api.php?action=watch&amp;title=Main_Page&amp;unwatch">api.php?action=watch&amp;title=Main_Page&amp;unwatch</a>

<b>* action=patrol *</b>
  Patrol a page or revision. 

This module requires read rights.
This module requires write rights.
Parameters:
  token          - Patrol token obtained from list=recentchanges
  rcid           - Recentchanges ID to patrol
Example:
  <a href="api.php?action=patrol&amp;token=123abc&amp;rcid=230672766">api.php?action=patrol&amp;token=123abc&amp;rcid=230672766</a>

<b>* action=import *</b>
  Import a page from another wiki, or an XML file

This module requires read rights.
This module requires write rights.
This module only accepts POST requests.
Parameters:
  token          - Import token obtained through prop=info
  summary        - Import summary
  xml            - Uploaded XML file
  interwikisource - For interwiki imports: wiki to import from
                   One value: 
  interwikipage  - For interwiki imports: page to import
  fullhistory    - For interwiki imports: import the full history, not just the current version
  templates      - For interwiki imports: import all included templates as well
  namespace      - For interwiki imports: import to this namespace
                   One value: 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15
Examples:
  Import [[meta:Help:Parserfunctions]] to namespace 100 with full history:
    <a href="api.php?action=import&amp;interwikisource=meta&amp;interwikipage=Help:ParserFunctions&amp;namespace=100&amp;fullhistory&amp;token=123ABC">api.php?action=import&amp;interwikisource=meta&amp;interwikipage=Help:ParserFunctions&amp;namespace=100&amp;fullhistory&amp;token=123ABC</a>

<b>* action=userrights *</b>
  Add/remove a user to/from groups

This module requires read rights.
This module requires write rights.
This module only accepts POST requests.
Parameters:
  user           - User name
  add            - Add the user to these groups
                   Values (separate with '|'): bot, sysop, bureaucrat
  remove         - Remove the user from these groups
                   Values (separate with '|'): bot, sysop, bureaucrat
  token          - A userrights token previously retrieved through list=users
  reason         - Reason for the change
                   Default: 
Example:
  <a href="api.php?action=userrights&amp;user=FooBot&amp;add=bot&amp;remove=sysop|bureaucrat&amp;token=123ABC">api.php?action=userrights&amp;user=FooBot&amp;add=bot&amp;remove=sysop|bureaucrat&amp;token=123ABC</a>


<b>*** *** *** *** *** *** *** *** *** ***  Permissions *** *** *** *** *** *** *** *** *** ***</b> 

<b>* writeapi *</b>
  Use of the write API
Granted to:
  user, bot
<b>* apihighlimits *</b>
  Use higher limits in API queries (Slow queries: 500 results; Fast queries: 5000 results). The limits for slow queries also apply to multivalue parameters.
Granted to:
  bot, sysop

<b>*** *** *** *** *** *** *** *** *** ***  Formats  *** *** *** *** *** *** *** *** *** ***</b> 

<b>* format=json *</b>
  Output data in JSON format

This module requires read rights.
Parameters:
  callback       - If specified, wraps the output into a given function call. For safety, all user-specific data will be restricted.
Example:
  <a href="api.php?action=query&amp;meta=siteinfo&amp;siprop=namespaces&amp;format=json">api.php?action=query&amp;meta=siteinfo&amp;siprop=namespaces&amp;format=json</a>

<b>* format=jsonfm *</b>
  Output data in JSON format (pretty-print in HTML)

This module requires read rights.
Parameters:
  callback       - If specified, wraps the output into a given function call. For safety, all user-specific data will be restricted.
Example:
  <a href="api.php?action=query&amp;meta=siteinfo&amp;siprop=namespaces&amp;format=jsonfm">api.php?action=query&amp;meta=siteinfo&amp;siprop=namespaces&amp;format=jsonfm</a>

<b>* format=php *</b>
  Output data in serialized PHP format

This module requires read rights.
Example:
  <a href="api.php?action=query&amp;meta=siteinfo&amp;siprop=namespaces&amp;format=php">api.php?action=query&amp;meta=siteinfo&amp;siprop=namespaces&amp;format=php</a>

<b>* format=phpfm *</b>
  Output data in serialized PHP format (pretty-print in HTML)

This module requires read rights.
Example:
  <a href="api.php?action=query&amp;meta=siteinfo&amp;siprop=namespaces&amp;format=phpfm">api.php?action=query&amp;meta=siteinfo&amp;siprop=namespaces&amp;format=phpfm</a>

<b>* format=wddx *</b>
  Output data in WDDX format

This module requires read rights.
Example:
  <a href="api.php?action=query&amp;meta=siteinfo&amp;siprop=namespaces&amp;format=wddx">api.php?action=query&amp;meta=siteinfo&amp;siprop=namespaces&amp;format=wddx</a>

<b>* format=wddxfm *</b>
  Output data in WDDX format (pretty-print in HTML)

This module requires read rights.
Example:
  <a href="api.php?action=query&amp;meta=siteinfo&amp;siprop=namespaces&amp;format=wddxfm">api.php?action=query&amp;meta=siteinfo&amp;siprop=namespaces&amp;format=wddxfm</a>

<b>* format=xml *</b>
  Output data in XML format

This module requires read rights.
Parameters:
  xmldoublequote - If specified, double quotes all attributes and content.
  xslt           - If specified, adds &lt;xslt&gt; as stylesheet
Example:
  <a href="api.php?action=query&amp;meta=siteinfo&amp;siprop=namespaces&amp;format=xml">api.php?action=query&amp;meta=siteinfo&amp;siprop=namespaces&amp;format=xml</a>

<b>* format=xmlfm *</b>
  Output data in XML format (pretty-print in HTML)

This module requires read rights.
Parameters:
  xmldoublequote - If specified, double quotes all attributes and content.
  xslt           - If specified, adds &lt;xslt&gt; as stylesheet
Example:
  <a href="api.php?action=query&amp;meta=siteinfo&amp;siprop=namespaces&amp;format=xmlfm">api.php?action=query&amp;meta=siteinfo&amp;siprop=namespaces&amp;format=xmlfm</a>

<b>* format=yaml *</b>
  Output data in YAML format

This module requires read rights.
Example:
  <a href="api.php?action=query&amp;meta=siteinfo&amp;siprop=namespaces&amp;format=yaml">api.php?action=query&amp;meta=siteinfo&amp;siprop=namespaces&amp;format=yaml</a>

<b>* format=yamlfm *</b>
  Output data in YAML format (pretty-print in HTML)

This module requires read rights.
Example:
  <a href="api.php?action=query&amp;meta=siteinfo&amp;siprop=namespaces&amp;format=yamlfm">api.php?action=query&amp;meta=siteinfo&amp;siprop=namespaces&amp;format=yamlfm</a>

<b>* format=rawfm *</b>
  Output data with the debuging elements in JSON format (pretty-print in HTML)

This module requires read rights.
Parameters:
  callback       - If specified, wraps the output into a given function call. For safety, all user-specific data will be restricted.
Example:
  <a href="api.php?action=query&amp;meta=siteinfo&amp;siprop=namespaces&amp;format=rawfm">api.php?action=query&amp;meta=siteinfo&amp;siprop=namespaces&amp;format=rawfm</a>

<b>* format=txt *</b>
  Output data in PHP's print_r() format

This module requires read rights.
Example:
  <a href="api.php?action=query&amp;meta=siteinfo&amp;siprop=namespaces&amp;format=txt">api.php?action=query&amp;meta=siteinfo&amp;siprop=namespaces&amp;format=txt</a>

<b>* format=txtfm *</b>
  Output data in PHP's print_r() format (pretty-print in HTML)

This module requires read rights.
Example:
  <a href="api.php?action=query&amp;meta=siteinfo&amp;siprop=namespaces&amp;format=txtfm">api.php?action=query&amp;meta=siteinfo&amp;siprop=namespaces&amp;format=txtfm</a>

<b>* format=dbg *</b>
  Output data in PHP's var_export() format

This module requires read rights.
Example:
  <a href="api.php?action=query&amp;meta=siteinfo&amp;siprop=namespaces&amp;format=dbg">api.php?action=query&amp;meta=siteinfo&amp;siprop=namespaces&amp;format=dbg</a>

<b>* format=dbgfm *</b>
  Output data in PHP's var_export() format (pretty-print in HTML)

This module requires read rights.
Example:
  <a href="api.php?action=query&amp;meta=siteinfo&amp;siprop=namespaces&amp;format=dbgfm">api.php?action=query&amp;meta=siteinfo&amp;siprop=namespaces&amp;format=dbgfm</a>


<b>*** Credits: ***</b>
   API developers:
       Roan Kattouw &lt;Firstname&gt;.&lt;Lastname&gt;@home.nl (lead developer Sep 2007-present)
       Victor Vasiliev - vasilvv at gee mail dot com
       Bryan Tong Minh - bryan . tongminh @ gmail . com
       Sam Reed - sam @ reedyboy . net
       Yuri Astrakhan &lt;Firstname&gt;&lt;Lastname&gt;@gmail.com (creator, lead developer Sep 2006-Sep 2007)
   
   Please send your comments, suggestions and questions to mediawiki-api@lists.wikimedia.org
   or file a bug report at <a href="http://bugzilla.wikimedia.org/">http://bugzilla.wikimedia.org/</a>
<span style="color:blue;">&lt;/error&gt;</span>
<span style="color:blue;">&lt;/api&gt;</span>
</pre>
</body>
</html>
<!-- Served in 1.227 secs. -->