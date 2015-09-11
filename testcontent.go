package htmlparser

const blogPost string = `
<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<title>Attention Cannibalization | M-Shaped Brain</title>
<link rel="profile" href="http://gmpg.org/xfn/11">
<link rel="pingback" href="http://blog.calbucci.com/xmlrpc.php">
<link rel="alternate" type="application/rss+xml" title="M-Shaped Brain &raquo; Feed" href="http://blog.calbucci.com/feed/" />
<link rel="alternate" type="application/rss+xml" title="M-Shaped Brain &raquo; Comments Feed" href="http://blog.calbucci.com/comments/feed/" />
<link rel="alternate" type="application/rss+xml" title="M-Shaped Brain &raquo; Attention Cannibalization Comments Feed" href="http://blog.calbucci.com/2015/01/27/attention-cannibalization/feed/" />
<script type="text/javascript">
/* <![CDATA[ */
function addLoadEvent(func){var oldonload=window.onload;if(typeof window.onload!='function'){window.onload=func;}else{window.onload=function(){oldonload();func();}}}
/* ]]> */
</script>
<link rel='stylesheet' id='all-css-0' href='https://s1.wp.com/wp-content/mu-plugins/smileyproject/smileyproject.css?m=1412325662g' type='text/css' media='all' />
<link rel='stylesheet' id='open-sans-css'  href='//fonts.googleapis.com/css?family=Open+Sans%3A300italic%2C400italic%2C600italic%2C300%2C400%2C600&#038;subset=latin%2Clatin-ext&#038;ver=4.1' type='text/css' media='all' />
<link rel='stylesheet' id='all-css-2' href='https://s1.wp.com/_static/??-eJx9kFkOwjAMRC9EsAoIqR+IsySpoW6zqU5Ybo9biSJAyt/EfmOPA/ekKFhXOmSwzNBp7snGwFtPYSuVDfwiupOWMnr6QsSTMWTwRSVXrhQYuBi2E6VMMg8u0bl4r/F9vOGkTDHGySrOT4f/CQaGLAlHEx+rqA11NM4uzEnbUS2vGr4eJ30bvfocu5gIQszL/6yiNs3GCaXuk84z4bEjjQ69YDWbT8e3a5a9JF/WnP2pOex2+7Zpm+PwAm1jpSc=' type='text/css' media='all' />
<link rel='stylesheet' id='writr-montserrat-css'  href='http://fonts.googleapis.com/css?family=Montserrat:400,700' type='text/css' media='all' />
<link rel='stylesheet' id='all-css-4' href='https://s0.wp.com/_static/??-eJyNkNFuwjAMRX9onmnXFx7QviUNJjVK4shxV+3vVyjSOhBlL46vcq7sa5wKeMlG2TCNUOIYOFcMlEl5/njSvvta33DltYESVSxjj5OyKVb7jvQamwEMSpT/h058JH2NcvbLAjAVL+nBsEpKSc681C0si80T3DFxht4pfLVbdCCBKN4ZS/4j4BQdPwZYWZX6KGFuwzXwr9wymVRQKqIGJ9F0r7esRaotW2EdnF7m3t5nZx46DFF6Fy/AZzo0Xdvs9h9d055/ACRB3SM=' type='text/css' media='all' />
<link rel='stylesheet' id='print-css-4' href='https://s0.wp.com/wp-content/mu-plugins/global-print/global-print.css?m=1387483371g' type='text/css' media='print' />
<script type='text/javascript' src='https://s2.wp.com/_static/??-eJx9kNsOwjAIQH/Irl62hz0Yv2WrbGG2pRa6Rb/ePswYl2pCQggnwEEvQaE3Nl2B9ZTjniA+1lRNvNP/AOVwjJ1A5dC/YUNewIt2SQWbRvSsOfVsIgZBytVA1tKyxfPoQCwOmLsRCl1HPVpQiSFmwEu+aqAC922AfkYoLduIgoTO3FQExieUvDl8FH8/bKUu7nyo93V7atrmOL0ASSmFOw=='></script>
<link rel='stylesheet' id='all-css-0' href='https://s2.wp.com/wp-content/mu-plugins/highlander-comments/style.css?m=1377793621g' type='text/css' media='all' />
<!--[if lt IE 8]>
<link rel='stylesheet' id='highlander-comments-ie7-css'  href='https://s2.wp.com/wp-content/mu-plugins/highlander-comments/style-ie7.css?m=1351637563g&#038;ver=20110606' type='text/css' media='all' />
<![endif]-->
<link rel="EditURI" type="application/rsd+xml" title="RSD" href="https://calbucci.wordpress.com/xmlrpc.php?rsd" />
<link rel="wlwmanifest" type="application/wlwmanifest+xml" href="https://s1.wp.com/wp-includes/wlwmanifest.xml" />
<link rel='prev' title='Are you the user or the product? The tale of Parents Night Out vs. Kids Night&nbsp;Out.' href='http://blog.calbucci.com/2015/01/16/are-you-the-user-or-the-product-the-tale-of-parents-night-out-vs-kids-night-out/' />
<link rel='next' title='Open Letter to Larry Page: Let&#8217;s spin-off&nbsp;Blogger!' href='http://blog.calbucci.com/2015/02/01/open-letter-to-larry-page-lets-spin-off-blogger/' />
<meta name="generator" content="WordPress.com" />
<link rel='canonical' href='http://blog.calbucci.com/2015/01/27/attention-cannibalization/' />
<link rel='shortlink' href='http://wp.me/p5rzcW-b9' />
<link rel="alternate" type="application/json+oembed" href="https://public-api.wordpress.com/oembed/1.0/?format=json&amp;url=http%3A%2F%2Fblog.calbucci.com%2F2015%2F01%2F27%2Fattention-cannibalization%2F&amp;for=wpcom-auto-discovery" /><link rel="alternate" type="application/xml+oembed" href="https://public-api.wordpress.com/oembed/1.0/?format=xml&amp;url=http%3A%2F%2Fblog.calbucci.com%2F2015%2F01%2F27%2Fattention-cannibalization%2F&amp;for=wpcom-auto-discovery" />
<!-- Jetpack Open Graph Tags -->
<meta property="og:type" content="article" />
<meta property="og:title" content="Attention Cannibalization" />
<meta property="og:url" content="http://blog.calbucci.com/2015/01/27/attention-cannibalization/" />
<meta property="og:description" content="If you come from a business or marketing background, revenue cannibalization is something you are familiar with. It’s the concept of a business launching a new product line that is creating revenue..." />
<meta property="article:published_time" content="2015-01-27T19:55:00+00:00" />
<meta property="article:modified_time" content="2015-01-27T20:11:37+00:00" />
<meta property="article:author" content="https://www.facebook.com/profile.php?id=562681776" />
<meta property="og:site_name" content="M-Shaped Brain" />
<meta property="og:image" content="https://c1.staticflickr.com/5/4053/4556605516_db311a3724_z.jpg" />
<meta name="twitter:site" content="@calbucci" />
<meta name="twitter:image" content="https://c1.staticflickr.com/5/4053/4556605516_db311a3724_z.jpg?w=240" />
<meta name="twitter:card" content="summary" />
<meta name="twitter:creator" content="@calbucci" />
<meta property="article:publisher" content="https://www.facebook.com/WordPresscom" />
<link rel="shortcut icon" type="image/x-icon" href="http://s2.wp.com/i/favicon.ico?m=1405977958g" sizes="16x16 24x24 32x32 48x48" />
<link rel="icon" type="image/x-icon" href="http://s2.wp.com/i/favicon.ico?m=1405977958g" sizes="16x16 24x24 32x32 48x48" />
<link rel="apple-touch-icon-precomposed" href="http://s0.wp.com/i/webclip.png?m=1391188133g" />
<link rel='openid.server' href='http://calbucci.wordpress.com/?openidserver=1' />
<link rel='openid.delegate' href='http://calbucci.wordpress.com/' />
<link rel="search" type="application/opensearchdescription+xml" href="http://blog.calbucci.com/osd.xml" title="M-Shaped Brain" />
<link rel="search" type="application/opensearchdescription+xml" href="http://wordpress.com/opensearch.xml" title="WordPress.com" />
<style type="text/css">
.recentcomments a {
display: inline !important;
padding: 0 !important;
margin: 0 !important;
}
table.recentcommentsavatartop img.avatar, table.recentcommentsavatarend img.avatar {
border: 0px;
margin: 0;
}
table.recentcommentsavatartop a, table.recentcommentsavatarend a {
border: 0px !important;
background-color: transparent !important;
}
td.recentcommentsavatarend, td.recentcommentsavatartop {
padding: 0px 0px 1px 0px;
margin: 0px;
}
td.recentcommentstextend {
border: none !important;
padding: 0px 0px 2px 10px;
}
.rtl td.recentcommentstextend {
padding: 0px 10px 2px 0px;
}
td.recentcommentstexttop {
border: none;
padding: 0px 0px 0px 10px;
}
.rtl td.recentcommentstexttop {
padding: 0px 10px 0px 0px;
}
</style>
<style type="text/css" media="print">#wpadminbar { display:none; }</style>
<style type="text/css" media="screen">
html { margin-top: 32px !important; }
* html body { margin-top: 32px !important; }
@media screen and ( max-width: 782px ) {
html { margin-top: 46px !important; }
* html body { margin-top: 46px !important; }
}
</style>
<meta name="application-name" content="M-Shaped Brain" /><meta name="msapplication-window" content="width=device-width;height=device-height" /><meta name="msapplication-tooltip" content="Marcelo Calbucci thoughts on startups, technology, entrepreneurship, education, health, fitness, parenting and more. @Calbucci" /><meta name="msapplication-task" content="name=Edit post;action-uri=https://wordpress.com/post/80451878/691;icon-uri=http://s2.wp.com/i/icons/post.ico" /><meta name="msapplication-task" content="name=Write a post;action-uri=https://calbucci.wordpress.com/wp-admin/post-new.php;icon-uri=http://s2.wp.com/i/icons/post.ico" /><meta name="msapplication-task" content="name=Moderate comments;action-uri=https://calbucci.wordpress.com/wp-admin/edit-comments.php?comment_status=moderated;icon-uri=http://s0.wp.com/i/icons/comment.ico" /><meta name="msapplication-task" content="name=Upload new media;action-uri=https://calbucci.wordpress.com/wp-admin/media-new.php;icon-uri=http://s2.wp.com/i/icons/media.ico" /><meta name="msapplication-task" content="name=Blog stats;action-uri=https://calbucci.wordpress.com/wp-admin/index.php?page=stats;icon-uri=http://s1.wp.com/i/icons/stats.ico" /><meta name="title" content="Attention Cannibalization | M-Shaped Brain on WordPress.com" />
<meta name="description" content="If you come from a business or marketing background, revenue cannibalization is something you are familiar with. It’s the concept of a business launching a new product line that is creating revenue by taking revenue away from another product line in the same business but not growing the revenue/profit pie. If you are an engineer&hellip;" />
<style type="text/css" id="syntaxhighlighteranchor"></style>
<style type="text/css" id="custom-colors-css">
span.infinite-loader .spinner div div {
background: #303030 !important;
}
body { background-color: #303030;}
.jetpack-recipe-title { border-color: #1abc9c;}
.jetpack-recipe { border-color: #1abc9c;}
.jetpack-recipe { border-color: rgba( 26, 188, 156, 0.15 );}
cite,
.widget-area button,
.widget-area html input[type="button"],
.widget-area input[type="reset"],
.widget-area input[type="submit"],
a,
.social-links a,
.entry-title a:hover,
.entry-title a:focus,
.entry-title a:active,
.entry-meta a:hover,
.entry-meta a:focus,
.entry-meta a:active,
#content [class*="navigation"] a:hover,
#content [class*="navigation"] a:focus,
#content [class*="navigation"] a:active,
.comment-metadata a:hover,
.comment-metadata a:focus,
.comment-metadata a:active,
.widget_tag_cloud a:hover,
.widget_tag_cloud a:focus,
.widget_tag_cloud a:active,
.content-area .widget_categories ul a:hover,
.content-area .widget_recent_entries ul a:hover,
.widget-grofile .grofile-full-link,
.widget-grofile .grofile-full-link:hover,
.widget-grofile .grofile-full-link:focus,
.widget-grofile .grofile-full-link:active,
.custom-colors .wp_widget_tag_cloud a:hover,
.custom-colors .wp_widget_tag_cloud a:focus,
.custom-colors .wp_widget_tag_cloud a:active { color: #1ABC9C;}
mark,
ins,
button,
html input[type="button"],
input[type="reset"],
input[type="submit"],
body:after,
.site-header,
.main-navigation ul ul,
.more-link,
.wp-audio-shortcode .mejs-controls .mejs-time-rail .mejs-time-current,
.wp-audio-shortcode .mejs-controls .mejs-horizontal-volume-slider .mejs-horizontal-volume-current,
.content-area .widget_tag_cloud a,
#infinite-handle span { background-color: #1ABC9C;}
a:hover,
a:focus,
a:active,
.site-title a:hover,
.site-info a:hover,
.widget_archive ul a:hover,
.widget_categories ul a:hover,
.widget_links ul a:hover,
.widget_nav_menu ul a:hover,
.widget_meta ul a:hover,
.widget_pages ul a:hover,
.widget_recent_comments ul a:hover,
.widget_recent_entries ul a:hover,
.widget_rss ul a:hover,
.widget_rss_links ul a:hover,
.custom-colors .widget_top-posts ul a:hover { color: #16A085;}
button:hover,
html input[type="button"]:hover,
input[type="reset"]:hover,
input[type="submit"]:hover,
button:focus,
html input[type="button"]:focus,
input[type="reset"]:focus,
input[type="submit"]:focus,
button:active,
html input[type="button"]:active,
input[type="reset"]:active,
input[type="submit"]:active,
#sidebar-toggle,
.main-navigation a:hover,
.menu-toggle,
.dropdown-icon,
.social-links a,
.more-link:hover,
.more-link:focus,
.more-link:active,
.widget th,
.widget tfoot tr,
.widget_tag_cloud a,
.content-area .widget_tag_cloud a:hover,
.content-area .widget_tag_cloud a:focus,
.content-area .widget_tag_cloud a:active,
#infinite-handle span:hover,
.custom-colors .wp_widget_tag_cloud a { background-color: #16A085;}
.widget-area input[type="text"],
.widget-area input[type="url"],
.widget-area input[type="email"],
.widget-area input[type="password"],
.widget-area input[type="search"],
.widget-area textarea { border-color: #16A085;}
#sidebar-toggle:before,
#sidebar-toggle:after { border-top-color: #16A085;}
.site-branding,
.main-navigation:after,
#social-links,
.widget,
.widget th,
.widget td { border-bottom-color: #16A085;}
</style>
<script type="text/javascript" id="custom-fonts-js">
(function(doc) {

var config = {"kitId":"sbc4lod","scriptTimeout":3000},
h=doc.getElementsByTagName("html")[0];h.className+=" wf-loading";var t=setTimeout(function(){h.className=h.className.replace(/(\s|^)wf-loading(\s|$)/g," ");h.className+="wf-inactive"},config.scriptTimeout);var tk=doc.createElement("script"),d=false;tk.src='//use.typekit.net/'+config.kitId+'.js';tk.type="text/javascript";tk.async="true";tk.onload=tk.onreadystatechange=function(){var a=this.readyState;if(d||a&&a!="complete"&&a!="loaded")return;d=true;clearTimeout(t);try{Typekit.load(config)}catch(b){}};var s=doc.getElementsByTagName("script")[0];s.parentNode.insertBefore(tk,s);
})(document);
</script>
<style type="text/css" id="custom-fonts-css">/* Body Text */ .wf-loading body,.wf-loading button,.wf-loading input,.wf-loading select,.wf-loading textarea { visibility: hidden; } .wf-active body,.wf-active button,.wf-active input,.wf-active select,.wf-active textarea { font-family: museo-sans-1,museo-sans-2,Montserrat,sans-serif; font-variant: normal; font-size: 17px; } .wf-loading pre { visibility: hidden; } .wf-active pre { font-size: 16px; } .wf-loading code,.wf-loading kbd,.wf-loading tt,.wf-loading var { visibility: hidden; } .wf-active code,.wf-active kbd,.wf-active tt,.wf-active var { font-size: 16px; } .wf-loading sub,.wf-loading sup { visibility: hidden; } .wf-active sub,.wf-active sup { font-size: 80%; } .wf-loading small { visibility: hidden; } .wf-active small { font-size: 80%; } .wf-loading big { visibility: hidden; } .wf-active big { font-size: 133%; } .wf-loading button,.wf-loading input,.wf-loading select,.wf-loading textarea { visibility: hidden; } .wf-active button,.wf-active input,.wf-active select,.wf-active textarea { font-size: 106%; } .wf-loading button,html.wf-loading input[type="button"],.wf-loading input[type="reset"],.wf-loading input[type="submit"] { visibility: hidden; } .wf-active button,html.wf-active input[type="button"],.wf-active input[type="reset"],.wf-active input[type="submit"] { font-size: 15px; } .wf-loading .site-info { visibility: hidden; } .wf-active .site-info { font-size: 13px; } .wf-loading .entry-meta { visibility: hidden; } .wf-active .entry-meta { font-size: 13px; } .wf-loading .entry-meta .genericon { visibility: hidden; } .wf-active .entry-meta .genericon { font-size: 13px; } .wf-loading .more-link { visibility: hidden; } .wf-active .more-link { font-size: 15px; } .wf-loading .format-status .entry-content { visibility: hidden; } .wf-active .format-status .entry-content { font-size: 27px; } .wf-loading .wp-caption-text { visibility: hidden; } .wf-active .wp-caption-text { font-size: 13px; } .wf-loading #content [class*="navigation"] { visibility: hidden; } .wf-active #content [class*="navigation"] { font-size: 13px; } .wf-loading #content [class*="navigation"] .genericon { visibility: hidden; } .wf-active #content [class*="navigation"] .genericon { font-size: 6px; } .wf-loading #comments #respond { visibility: hidden; } .wf-active #comments #respond { font-family: museo-sans-1,museo-sans-2,Montserrat,sans-serif; font-variant: normal; } .wf-loading .comment-metadata { visibility: hidden; } .wf-active .comment-metadata { font-size: 13px; } .wf-loading .comment-metadata .genericon { visibility: hidden; } .wf-active .comment-metadata .genericon { font-size: 13px; } .wf-loading .form-allowed-tags { visibility: hidden; } .wf-active .form-allowed-tags { font-size: 13px; } .wf-loading .form-allowed-tags code { visibility: hidden; } .wf-active .form-allowed-tags code { font-size: 12px; } .wf-loading .widget { visibility: hidden; } .wf-active .widget { font-size: 15px; } .wf-loading .widget_rss .rssSummary { visibility: hidden; } .wf-active .widget_rss .rssSummary { font-size: 13px; } .wf-loading .widget_rss cite { visibility: hidden; } .wf-active .widget_rss cite { font-size: 13px; } .wf-loading .widget_rss cite:before { visibility: hidden; } .wf-active .widget_rss cite:before { font-size: 13px; } .wf-loading .widget_tag_cloud a { visibility: hidden; } .wf-active .widget_tag_cloud a { font-size: 13px; } .wf-loading .widget_recent_entries .post-date,.wf-loading .widget_rss .rss-date { visibility: hidden; } .wf-active .widget_recent_entries .post-date,.wf-active .widget_rss .rss-date { font-size: 13px; } .wf-loading .widget_recent_entries .post-date:before,.wf-loading .widget_rss .rss-date:before { visibility: hidden; } .wf-active .widget_recent_entries .post-date:before,.wf-active .widget_rss .rss-date:before { font-size: 13px; } .wf-loading #infinite-handle span { visibility: hidden; } .wf-active #infinite-handle span { font-size: 15px; } .wf-loading #page .entry-content .rating-msg,.wf-loading #page .entry-content div.sharedaddy h3,.wf-loading #page .entry-content h3.sd-title,.wf-loading #page .entry-summary .rating-msg,.wf-loading #page .entry-summary div.sharedaddy h3,.wf-loading #page .entry-summary h3.sd-title,.wf-loading #primary div.sharedaddy .jp-relatedposts-headline em { visibility: hidden; } .wf-active #page .entry-content .rating-msg,.wf-active #page .entry-content div.sharedaddy h3,.wf-active #page .entry-content h3.sd-title,.wf-active #page .entry-summary .rating-msg,.wf-active #page .entry-summary div.sharedaddy h3,.wf-active #page .entry-summary h3.sd-title,.wf-active #primary div.sharedaddy .jp-relatedposts-headline em { font-family: museo-sans-1,museo-sans-2,Montserrat,sans-serif; font-variant: normal; font-size: 12px; } .wf-loading .comment-content .pd-rating { visibility: hidden; } .wf-active .comment-content .pd-rating { font-size: 12px; } .wf-loading .widget-grofile .grofile-meta h4 { visibility: hidden; } .wf-active .widget-grofile .grofile-meta h4 { font-family: museo-sans-1,museo-sans-2,Montserrat,sans-serif; font-variant: normal; } .wf-loading .widget-grofile .grofile-full-link { visibility: hidden; } .wf-active .widget-grofile .grofile-full-link { font-size: 15px; }
</style>
<link rel="stylesheet" id="custom-css-css" type="text/css" href="https://s2.wp.com/?custom-css=1&#038;csblog=5rzcW&#038;cscache=6&#038;csrev=12" />
</head>
<body class="single single-post postid-691 single-format-standard logged-in admin-bar no-customize-support custom-background mp6 typekit-enabled customizer-styles-applied color-scheme-green sidebar-closed highlander-enabled highlander-light custom-colors">
<div id="page" class="hfeed site">

<header id="masthead" class="site-header" role="banner">
<a class="site-logo"  href="http://blog.calbucci.com/" title="M-Shaped Brain" rel="home">
<img src="http://gravatar.com/avatar/002201b3a1df4df7adc77ab91d675568/?s=120&#038;d=identicon" width="120" height="120" alt="" class="no-grav header-image" />
</a>
<div class="site-branding">
<h1 class="site-title"><a href="http://blog.calbucci.com/" title="M-Shaped Brain" rel="home">M-Shaped Brain</a></h1>
<h2 class="site-description">Marcelo Calbucci thoughts on startups, technology, entrepreneurship, education, health, fitness, parenting and more. @Calbucci</h2>
</div>
</header><!-- #masthead -->

<div id="sidebar" class="sidebar-area">
<a id="sidebar-toggle" href="#" title="Sidebar"><span class="genericon genericon-close"></span><span class="screen-reader-text">Sidebar</span></a>
<div id="secondary" class="widget-area" role="complementary">
<aside id="text-3" class="widget widget_text">			<div class="textwidget"><a href='http://cloud.feedly.com/#subscription%2Ffeed%2Fhttp%3A%2F%2Ffeeds.feedburner.com%2FMarceloCalbucci%20' target='blank'><img id='feedlyFollow' src='http://s3.feedly.com/img/follows/feedly-follow-rectangle-flat-big_2x.png' alt='follow us in feedly' width='131' height='56'></a></div>
</aside>		<aside id="recent-posts-3" class="widget widget_recent_entries">		<h1 class="widget-title">Recent Posts</h1>		<ul>
<li>
<a href="http://blog.calbucci.com/2015/02/03/the-unmet-unspoken-needs-of-customers/">The Unmet &amp; Unspoken Needs of&nbsp;Customers</a>
</li>
<li>
<a href="http://blog.calbucci.com/2015/02/01/open-letter-to-larry-page-lets-spin-off-blogger/">Open Letter to Larry Page: Let&#8217;s spin-off&nbsp;Blogger!</a>
</li>
<li>
<a href="http://blog.calbucci.com/2015/01/27/attention-cannibalization/">Attention Cannibalization</a>
</li>
<li>
<a href="http://blog.calbucci.com/2015/01/16/are-you-the-user-or-the-product-the-tale-of-parents-night-out-vs-kids-night-out/">Are you the user or the product? The tale of Parents Night Out vs. Kids Night&nbsp;Out.</a>
</li>
<li>
<a href="http://blog.calbucci.com/2015/01/09/helping-more-women-men-in-2015-open-office-hours/">Helping More Women (&amp; Men) in 2015: Open Office&nbsp;Hours</a>
</li>
</ul>
</aside><aside id="archives-3" class="widget widget_archive"><h1 class="widget-title">Archives</h1>		<ul>
<li><a href='http://blog.calbucci.com/2015/02/'>February 2015</a></li>
<li><a href='http://blog.calbucci.com/2015/01/'>January 2015</a></li>
<li><a href='http://blog.calbucci.com/2014/07/'>July 2014</a></li>
<li><a href='http://blog.calbucci.com/2014/05/'>May 2014</a></li>
<li><a href='http://blog.calbucci.com/2014/03/'>March 2014</a></li>
<li><a href='http://blog.calbucci.com/2014/01/'>January 2014</a></li>
<li><a href='http://blog.calbucci.com/2013/12/'>December 2013</a></li>
<li><a href='http://blog.calbucci.com/2013/10/'>October 2013</a></li>
<li><a href='http://blog.calbucci.com/2013/09/'>September 2013</a></li>
<li><a href='http://blog.calbucci.com/2013/08/'>August 2013</a></li>
<li><a href='http://blog.calbucci.com/2013/04/'>April 2013</a></li>
<li><a href='http://blog.calbucci.com/2013/03/'>March 2013</a></li>
<li><a href='http://blog.calbucci.com/2013/02/'>February 2013</a></li>
<li><a href='http://blog.calbucci.com/2013/01/'>January 2013</a></li>
<li><a href='http://blog.calbucci.com/2012/11/'>November 2012</a></li>
<li><a href='http://blog.calbucci.com/2012/10/'>October 2012</a></li>
<li><a href='http://blog.calbucci.com/2012/09/'>September 2012</a></li>
<li><a href='http://blog.calbucci.com/2012/08/'>August 2012</a></li>
<li><a href='http://blog.calbucci.com/2012/07/'>July 2012</a></li>
<li><a href='http://blog.calbucci.com/2012/05/'>May 2012</a></li>
<li><a href='http://blog.calbucci.com/2012/04/'>April 2012</a></li>
<li><a href='http://blog.calbucci.com/2012/03/'>March 2012</a></li>
<li><a href='http://blog.calbucci.com/2012/02/'>February 2012</a></li>
<li><a href='http://blog.calbucci.com/2012/01/'>January 2012</a></li>
<li><a href='http://blog.calbucci.com/2011/12/'>December 2011</a></li>
<li><a href='http://blog.calbucci.com/2011/11/'>November 2011</a></li>
<li><a href='http://blog.calbucci.com/2011/10/'>October 2011</a></li>
<li><a href='http://blog.calbucci.com/2011/09/'>September 2011</a></li>
<li><a href='http://blog.calbucci.com/2011/08/'>August 2011</a></li>
<li><a href='http://blog.calbucci.com/2011/07/'>July 2011</a></li>
<li><a href='http://blog.calbucci.com/2011/06/'>June 2011</a></li>
<li><a href='http://blog.calbucci.com/2011/05/'>May 2011</a></li>
<li><a href='http://blog.calbucci.com/2011/04/'>April 2011</a></li>
<li><a href='http://blog.calbucci.com/2011/03/'>March 2011</a></li>
<li><a href='http://blog.calbucci.com/2011/02/'>February 2011</a></li>
<li><a href='http://blog.calbucci.com/2011/01/'>January 2011</a></li>
<li><a href='http://blog.calbucci.com/2010/12/'>December 2010</a></li>
<li><a href='http://blog.calbucci.com/2010/11/'>November 2010</a></li>
<li><a href='http://blog.calbucci.com/2010/10/'>October 2010</a></li>
<li><a href='http://blog.calbucci.com/2010/09/'>September 2010</a></li>
<li><a href='http://blog.calbucci.com/2010/08/'>August 2010</a></li>
<li><a href='http://blog.calbucci.com/2010/07/'>July 2010</a></li>
<li><a href='http://blog.calbucci.com/2010/06/'>June 2010</a></li>
<li><a href='http://blog.calbucci.com/2010/05/'>May 2010</a></li>
<li><a href='http://blog.calbucci.com/2010/04/'>April 2010</a></li>
<li><a href='http://blog.calbucci.com/2010/03/'>March 2010</a></li>
<li><a href='http://blog.calbucci.com/2010/02/'>February 2010</a></li>
<li><a href='http://blog.calbucci.com/2010/01/'>January 2010</a></li>
<li><a href='http://blog.calbucci.com/2009/11/'>November 2009</a></li>
<li><a href='http://blog.calbucci.com/2009/10/'>October 2009</a></li>
<li><a href='http://blog.calbucci.com/2009/09/'>September 2009</a></li>
<li><a href='http://blog.calbucci.com/2009/08/'>August 2009</a></li>
<li><a href='http://blog.calbucci.com/2009/07/'>July 2009</a></li>
<li><a href='http://blog.calbucci.com/2009/05/'>May 2009</a></li>
<li><a href='http://blog.calbucci.com/2009/04/'>April 2009</a></li>
<li><a href='http://blog.calbucci.com/2009/03/'>March 2009</a></li>
<li><a href='http://blog.calbucci.com/2009/02/'>February 2009</a></li>
<li><a href='http://blog.calbucci.com/2009/01/'>January 2009</a></li>
<li><a href='http://blog.calbucci.com/2008/11/'>November 2008</a></li>
<li><a href='http://blog.calbucci.com/2008/10/'>October 2008</a></li>
<li><a href='http://blog.calbucci.com/2008/09/'>September 2008</a></li>
<li><a href='http://blog.calbucci.com/2008/08/'>August 2008</a></li>
<li><a href='http://blog.calbucci.com/2008/07/'>July 2008</a></li>
<li><a href='http://blog.calbucci.com/2008/06/'>June 2008</a></li>
<li><a href='http://blog.calbucci.com/2008/05/'>May 2008</a></li>
<li><a href='http://blog.calbucci.com/2008/04/'>April 2008</a></li>
<li><a href='http://blog.calbucci.com/2008/03/'>March 2008</a></li>
<li><a href='http://blog.calbucci.com/2008/02/'>February 2008</a></li>
<li><a href='http://blog.calbucci.com/2008/01/'>January 2008</a></li>
<li><a href='http://blog.calbucci.com/2007/12/'>December 2007</a></li>
<li><a href='http://blog.calbucci.com/2007/11/'>November 2007</a></li>
<li><a href='http://blog.calbucci.com/2007/10/'>October 2007</a></li>
<li><a href='http://blog.calbucci.com/2007/09/'>September 2007</a></li>
<li><a href='http://blog.calbucci.com/2007/08/'>August 2007</a></li>
<li><a href='http://blog.calbucci.com/2007/07/'>July 2007</a></li>
<li><a href='http://blog.calbucci.com/2007/06/'>June 2007</a></li>
<li><a href='http://blog.calbucci.com/2007/05/'>May 2007</a></li>
<li><a href='http://blog.calbucci.com/2007/04/'>April 2007</a></li>
<li><a href='http://blog.calbucci.com/2007/03/'>March 2007</a></li>
<li><a href='http://blog.calbucci.com/2007/02/'>February 2007</a></li>
<li><a href='http://blog.calbucci.com/2007/01/'>January 2007</a></li>
<li><a href='http://blog.calbucci.com/2006/12/'>December 2006</a></li>
<li><a href='http://blog.calbucci.com/2006/11/'>November 2006</a></li>
<li><a href='http://blog.calbucci.com/2006/10/'>October 2006</a></li>
<li><a href='http://blog.calbucci.com/2006/06/'>June 2006</a></li>
</ul>
</aside>					</div><!-- #secondary -->
</div><!-- #sidebar -->
<div id="content" class="site-content">
<div id="primary" class="content-area">
<main id="main" class="site-main" role="main">


<article id="post-691" class="post-691 post type-post status-publish format-standard hentry category-uncategorized">
<header class="entry-header">
<h1 class="entry-title">Attention Cannibalization</h1>

<span class="entry-format-badge genericon genericon-standard"><span class="screen-reader-text">Standard</span></span>
</header><!-- .entry-header -->
<div class="entry-content">
<p>If you come from a business or marketing background, revenue cannibalization is something you are familiar with. It’s the concept of a business launching a new product line that is creating revenue by taking revenue away from another product line in the same business but not growing the revenue/profit pie. If you are an engineer or designer and you haven’t heard about revenue cannibalization, you just did.</p>
<p>There are good reasons to do this. You can do it because you know one type of business is dying and a new model is becoming more prevalent (e.g., people are subscribing to Office 365 in lieu of buying a full license of MS Office). Another reason is to convert users from an older version of a product to a newer version (e.g., Gillette blades), either to prevent customers jumping ship or because the new product will present newer revenue opportunities in the future.</p>
<p><img class="alignright" src="https://c1.staticflickr.com/5/4053/4556605516_db311a3724_z.jpg" alt="" width="300" height="187" />There are bad reasons to do it, and that’s most common than not. It’s when a company is too big and lacks upper management coordination, so many units are building similar or substitutes products and competing against each other, or the same product has multiple add-ons that compete for the same dollar. In those cases no new value is being created for the company as a whole, just revenue is moved from one product to another.</p>
<h2>What if we think of the products as business, and the features as products?</h2>
<p>Let’s take this a level down. Imagine now that you are talking about a single product and instead of revenue you change the metric to attention. This is <em><strong>attention cannibalization</strong></em>. The idea that a new feature might be cannibalizing the attention of an existing feature, without creating any new attention to the product as a whole.</p>
<p>This is dangerous and I’m going to give away the conclusion first! It&#8217;s dangerous because you are led to believe what you are doing is improving the product and the overall experience and it to &#8220;validate&#8221; decisions made months ago, at the cost of not moving the product forward at all &#8212; you are actually moving it sideways.</p>
<p>Product Managers, Designers or Developers, might talk about their success in terms of usage of their feature. Maybe there was a feature that didn’t exist before, and after they launch they see that 30% of users use it right after sign up. Pop the champagne! Or, they change how a feature work, and they see that the usage for that feature went from 2 times per visit to 3 times per visit. Holy crap, a 50% improvement in usage!</p>
<p>The reality is that in a system, looking at the production value of a single component is worthless to understand the value of the whole system. But wait, it could be worse than that. It could be a new feature took attention away from a higher value feature. Maybe you added a new way for users to customize their avatar and that choice was enough to justify some folks from not upgrading to the premium version of the product. Revenue starts to dip, slowly and given the noise in the revenue signal and the fast speed in which new features are released, it can become quite harder to figure out the root cause of a drop in revenue, or a drop in overall conversion, engagement, virality or retention.</p>
<p>Another problem to pay attention to is when the attention cannibalization has a delayed impact. You add new feature X and you see a 25% usage on that feature on a daily basis and no drop in revenue or any other feature usage, so you might say to yourself that was worth it, but then a couple of weeks later you start experiencing a drop in retention. For example, when Twitter decided to auto-follow accounts for new users, and it seemed like it was a good idea because there was a jump in initial engagement with the product, but overall was a bad idea since retention was affected by it. This is the short-term gain for long-term loss strategy that you might not even know about it.</p>
<p>At this point you might be asking how to identify this problem. First and foremost, you must look at the product as whole. Did conversion, engagement, virality and retention went up or down for the product as a whole? Second, you must look at cohorts of data and see what&#8217;s the longitudinal impact of each feature launch? Third &#8212; and this might go against all the new trends in Lean Startup &#8212; don&#8217;t ship too fast. Making dramatic changes to your product without understanding the impact of previous changes is the sure way to lose track of what&#8217;s working and what&#8217;s not &#8212; clearly this is much easier said than done in this day and age.</p>
<div id="jp-post-flair" class="sharedaddy sd-like-enabled sd-sharing-enabled"><div class="sharedaddy sd-sharing-enabled"><div class="robots-nocontent sd-block sd-social sd-social-icon sd-sharing"><div class="sd-content"><ul><li class="share-twitter"><a rel="nofollow" class="share-twitter sd-button share-icon no-text" href="http://blog.calbucci.com/2015/01/27/attention-cannibalization/?share=twitter" target="_blank" title="Click to share on Twitter" id="sharing-twitter-691"><span></span><span class="sharing-screen-reader-text">Click to share on Twitter (Opens in new window)</span></a></li><li class="share-facebook"><a rel="nofollow" class="share-facebook sd-button share-icon no-text" href="http://blog.calbucci.com/2015/01/27/attention-cannibalization/?share=facebook" target="_blank" title="Share on Facebook" id="sharing-facebook-691"><span></span><span class="sharing-screen-reader-text">Share on Facebook (Opens in new window)</span></a></li><li class="share-linkedin"><a rel="nofollow" class="share-linkedin sd-button share-icon no-text" href="http://blog.calbucci.com/2015/01/27/attention-cannibalization/?share=linkedin" target="_blank" title="Click to share on LinkedIn" id="sharing-linkedin-691"><span></span><span class="sharing-screen-reader-text">Click to share on LinkedIn (Opens in new window)</span></a></li><li class="share-reddit"><a rel="nofollow" class="share-reddit sd-button share-icon no-text" href="http://blog.calbucci.com/2015/01/27/attention-cannibalization/?share=reddit" target="_blank" title="Click to share on Reddit"><span></span><span class="sharing-screen-reader-text">Click to share on Reddit (Opens in new window)</span></a></li><li class="share-google-plus-1"><a rel="nofollow" class="share-google-plus-1 sd-button share-icon no-text" href="http://blog.calbucci.com/2015/01/27/attention-cannibalization/?share=google-plus-1" target="_blank" title="Click to share on Google+" id="sharing-google-691"><span></span><span class="sharing-screen-reader-text">Click to share on Google+ (Opens in new window)</span></a></li><li class="share-end"></li></ul></div></div></div></div>			</div><!-- .entry-content -->
<footer class="entry-meta">
<ul class="clear">

<li class="date-meta">
<div class="genericon genericon-month"></div>
<span class="screen-reader-text">Date</span>
<a href="http://blog.calbucci.com/2015/01/27/attention-cannibalization/" rel="bookmark" title="11:55 AM">January 27, 2015</a>
</li>

<li class="comment-meta">
<div class="genericon genericon-comment"></div>
<span class="screen-reader-text">Comments</span>
<a href="http://blog.calbucci.com/2015/01/27/attention-cannibalization/#comments" title="Comment on Attention Cannibalization">1 Comment</a>		</li>
<li class="edit-link"><div class="genericon genericon-edit"></div><a class="post-edit-link" href="https://wordpress.com/post/80451878/691">Edit</a></li>		</ul>
</footer><!-- .entry-meta -->
</article><!-- #post-## -->
<nav role="navigation" id="nav-below" class="post-navigation">
<h1 class="screen-reader-text">Post navigation</h1>

<div class="nav-previous"><a href="http://blog.calbucci.com/2015/01/16/are-you-the-user-or-the-product-the-tale-of-parents-night-out-vs-kids-night-out/" rel="prev"><span class="genericon genericon-leftarrow"></span> Are you the user or the product? The tale of Parents Night Out vs. Kids Night&nbsp;Out.</a></div>		<div class="nav-next"><a href="http://blog.calbucci.com/2015/02/01/open-letter-to-larry-page-lets-spin-off-blogger/" rel="next">Open Letter to Larry Page: Let&#8217;s spin-off&nbsp;Blogger! <span class="genericon genericon-rightarrow"></span></a></div>

</nav><!-- #nav-below -->


<div id="comments" class="comments-area">

<h2 class="comments-title">
One thought on &ldquo;<span>Attention Cannibalization</span>&rdquo;		</h2>

<ol class="comment-list">

<li id="comment-207" class="pingback even thread-even depth-1 highlander-comment">
<div class="comment-body">
Pingback: <a href='http://blog.offeryour.com/?p=33427' rel='external nofollow' class='url'>1p – Attention Cannibalization &#8211; blog.offeryour.com</a> <span class="edit-link"><a class="comment-edit-link" href="https://calbucci.wordpress.com/wp-admin/comment.php?action=editcomment&amp;c=207">Edit</a></span>		</div>
</li><!-- #comment-## -->
</ol><!-- .comment-list -->



<div id="respond" class="comment-respond">
<h3 id="reply-title" class="comment-reply-title">Leave a Reply <small><a rel="nofollow" id="cancel-comment-reply-link" href="/2015/01/27/attention-cannibalization/#respond" style="display:none;">Cancel reply</a></small></h3>
<form action="http://blog.calbucci.com/wp-comments-post.php" method="post" id="commentform" class="comment-form">
<input type="hidden" id="highlander_comment_nonce" name="highlander_comment_nonce" value="f178928ac8" /><input type="hidden" name="_wp_http_referer" value="/2015/01/27/attention-cannibalization/" />
<input type="hidden" name="hc_post_as" id="hc_post_as" value="wordpress" />
<div class="comment-form-field comment-textarea">
<label for="comment">Enter your comment here...</label>
<div id="comment-form-comment"><textarea id="comment" name="comment" title="Enter your comment here..."></textarea></div>
</div>
<div id="comment-form-identity">
<div id="comment-form-nascar">
<p>Fill in your details below or click an icon to log in:</p>
<ul>
<li style="display:none;">
<a href="#comment-form-guest" id="postas-guest" title="Guest">
<span></span>
</a>
</li>
<li class="selected">
<a href="#comment-form-load-service:WordPress.com" id="postas-wordpress" title="WordPress.com">
<span></span>
</a>
</li>
<li>
<a href="#comment-form-load-service:Twitter" id="postas-twitter" title="Twitter">
<span></span>
</a>
</li>
<li>
<a href="#comment-form-load-service:Facebook" id="postas-facebook" title="Facebook">
<span></span>
</a>
</li>
<li>
</ul>
</div>
<div id="comment-form-guest" class="comment-form-service">
<div class="comment-form-padder">
<div class="comment-form-avatar">
<a href="https://gravatar.com/site/signup/" target="_blank">				<img src="http://1.gravatar.com/avatar/ad516503a11cd5ca435acc9bb6523536?s=25&amp;d=identicon&amp;forcedefault=y&amp;r=G" alt="Gravatar" width="25" class="no-grav" />
</a>			</div>
<div class="comment-form-fields">
<div class="comment-form-field comment-form-email">
<label for="email">Email <span class="required">(required)</span> <span class="nopublish">(Address never made public)</span></label>
<div class="comment-form-input"><input id="email" name="email" type="email" value="" /></div>
</div>
<div class="comment-form-field comment-form-author">
<label for="author">Name <span class="required">(required)</span></label>
<div class="comment-form-input"><input id="author" name="author" type="text" value="" /></div>
</div>
<div class="comment-form-field comment-form-url">
<label for="url">Website</label>
<div class="comment-form-input"><input id="url" name="url" type="text" value="" /></div>
</div>
</div>

</div>
</div>
<div id="comment-form-wordpress" class="comment-form-service selected">
<div class="comment-form-padder">
<div class="comment-form-avatar">
<img src="http://0.gravatar.com/avatar/002201b3a1df4df7adc77ab91d675568?s=25&amp;d=http%3A%2F%2F1.gravatar.com%2Favatar%2Fad516503a11cd5ca435acc9bb6523536%3Fs%3D25%26amp%3Bd%3Didenticon%26amp%3Bforcedefault%3Dy%26amp%3Br%3DG&amp;r=G" alt="Gravatar" width="25" class="no-grav" />
</div>
<div class="comment-form-fields">
<input type="hidden" name="wp_avatar" id="wordpress-avatar" class="comment-meta-wordpress" value="http://0.gravatar.com/avatar/002201b3a1df4df7adc77ab91d675568?s=25&#038;d=https%3A%2F%2Fs2.wp.com%2Fwp-content%2Fmu-plugins%2Fhighlander-comments%2Fimages%2Fwplogo.png&#038;r=G" />
<input type="hidden" name="wp_user_id" id="wordpress-user_id" class="comment-meta-wordpress" value="63737" />
<input type="hidden" name="wp_access_token" id="wordpress-access_token" class="comment-meta-wordpress" value="c934336089f76849cd5450f5f2333a47a88c8476" />
<p class="comment-form-posting-as pa-wordpress"><strong>Marcelo Calbucci:</strong> You are commenting using your WordPress.com account. <span class="comment-form-log-out">(&nbsp;<a href="javascript:HighlanderComments.doExternalLogout( 'wordpress' );">Log&nbsp;Out</a>&nbsp;/&nbsp;<a href="#" onclick="javascript:HighlanderComments.switchAccount();return false;">Change</a>&nbsp;)</span></p>
</div>

</div>
</div>
<div id="comment-form-twitter" class="comment-form-service">
<div class="comment-form-padder">
<div class="comment-form-avatar">
<img src="http://1.gravatar.com/avatar/ad516503a11cd5ca435acc9bb6523536?s=25&amp;d=identicon&amp;forcedefault=y&amp;r=G" alt="Twitter picture" width="25" class="no-grav" />
</div>
<div class="comment-form-fields">
<input type="hidden" name="twitter_avatar" id="twitter-avatar" class="comment-meta-twitter" value="" />
<input type="hidden" name="twitter_user_id" id="twitter-user_id" class="comment-meta-twitter" value="" />
<input type="hidden" name="twitter_access_token" id="twitter-access_token" class="comment-meta-twitter" value="" />
<p class="comment-form-posting-as pa-twitter"><strong></strong> You are commenting using your Twitter account. <span class="comment-form-log-out">(&nbsp;<a href="javascript:HighlanderComments.doExternalLogout( 'twitter' );">Log&nbsp;Out</a>&nbsp;/&nbsp;<a href="#" onclick="javascript:HighlanderComments.switchAccount();return false;">Change</a>&nbsp;)</span></p>
</div>

</div>
</div>
<div id="comment-form-facebook" class="comment-form-service">
<div class="comment-form-padder">
<div class="comment-form-avatar">
<img src="http://1.gravatar.com/avatar/ad516503a11cd5ca435acc9bb6523536?s=25&amp;d=identicon&amp;forcedefault=y&amp;r=G" alt="Facebook photo" width="25" class="no-grav" />
</div>
<div class="comment-form-fields">
<input type="hidden" name="fb_avatar" id="facebook-avatar" class="comment-meta-facebook" value="" />
<input type="hidden" name="fb_user_id" id="facebook-user_id" class="comment-meta-facebook" value="" />
<input type="hidden" name="fb_access_token" id="facebook-access_token" class="comment-meta-facebook" value="" />
<p class="comment-form-posting-as pa-facebook"><strong></strong> You are commenting using your Facebook account. <span class="comment-form-log-out">(&nbsp;<a href="javascript:HighlanderComments.doExternalLogout( 'facebook' );">Log&nbsp;Out</a>&nbsp;/&nbsp;<a href="#" onclick="javascript:HighlanderComments.switchAccount();return false;">Change</a>&nbsp;)</span></p>
</div>

</div>
</div>
<div id="comment-form-googleplus" class="comment-form-service">
<div class="comment-form-padder">
<div class="comment-form-avatar">
<img src="http://1.gravatar.com/avatar/ad516503a11cd5ca435acc9bb6523536?s=25&amp;d=identicon&amp;forcedefault=y&amp;r=G" alt="Google+ photo" width="25" class="no-grav" />
</div>
<div class="comment-form-fields">
<input type="hidden" name="googleplus_avatar" id="googleplus-avatar" class="comment-meta-googleplus" value="" />
<input type="hidden" name="googleplus_user_id" id="googleplus-user_id" class="comment-meta-googleplus" value="" />
<input type="hidden" name="googleplus_access_token" id="googleplus-access_token" class="comment-meta-googleplus" value="" />
<p class="comment-form-posting-as pa-googleplus"><strong></strong> You are commenting using your Google+ account. <span class="comment-form-log-out">(&nbsp;<a href="javascript:HighlanderComments.doExternalLogout( 'googleplus' );">Log&nbsp;Out</a>&nbsp;/&nbsp;<a href="#" onclick="javascript:HighlanderComments.switchAccount();return false;">Change</a>&nbsp;)</span></p>
</div>

</div>
</div>
<div id="comment-form-load-service" class="comment-form-service">
<div class="comment-form-posting-as-cancel"><a href="javascript:HighlanderComments.cancelExternalWindow();">Cancel</a></div>
<p>Connecting to %s</p>
</div>

</div>
<script type="text/javascript">
var highlander_expando_javascript = function(){
var input = document.createElement( 'input' ),
comment = jQuery( '#comment' );
if ( 'placeholder' in input ) {
comment.attr( 'placeholder', jQuery( '.comment-textarea label' ).remove().text() );
}
// Expando Mode: start small, then auto-resize on first click + text length
jQuery( '#comment-form-identity' ).hide();
jQuery( '#comment-form-subscribe' ).hide();
jQuery( '#commentform .form-submit' ).hide();
comment.css( { 'height':'10px' } ).one( 'focus', function() {
var timer = setInterval( HighlanderComments.resizeCallback, 10 )
jQuery( this ).animate( { 'height': HighlanderComments.initialHeight } ).delay( 100 ).queue( function(n) { clearInterval( timer ); HighlanderComments.resizeCallback(); n(); } );
jQuery( '#comment-form-identity' ).slideDown();
jQuery( '#comment-form-subscribe' ).slideDown();
jQuery( '#commentform .form-submit' ).slideDown();
});
}
jQuery(document).ready( highlander_expando_javascript );
</script>
<div id="comment-form-subscribe">
<p class="comment-subscription-form"><input type="checkbox" name="subscribe" id="subscribe" value="subscribe" style="width: auto;" tabindex="6"/> <label class="subscribe-label" id="subscribe-label" for="subscribe" style="display: inline;">Notify me of new comments via email.</label></p><p class="post-subscription-form"><input type="checkbox" name="subscribe_blog" id="subscribe_blog" value="subscribe" style="width: auto;" tabindex="7"/> <label class="subscribe-label" id="subscribe-blog-label" for="subscribe_blog"  style="display: inline;">Notify me of new posts via email.</label></p></div>
<p class="form-submit">
<input name="submit" type="submit" id="comment-submit" class="submit" value="Post Comment" />
<input type='hidden' name='comment_post_ID' value='691' id='comment_post_ID' />
<input type='hidden' name='comment_parent' id='comment_parent' value='0' />
</p>
<p style="display: none;"><input type="hidden" id="akismet_comment_nonce" name="akismet_comment_nonce" value="50a82219dc" /></p>
<input type="hidden" name="genseq" value="1423094356" />
<p style="display: none;"><input type="hidden" id="ak_js" name="ak_js" value="4"/></p>					</form>
</div><!-- #respond -->
<div style="clear: both"></div>
</div><!-- #comments -->

</main><!-- #main -->
</div><!-- #primary -->
</div><!-- #content -->
<footer id="colophon" class="site-footer" role="contentinfo">
<div id="social-links" class="clear">
<ul class="social-links clear">
<li>
<a href="http://twitter.com/calbucci" class="genericon genericon-twitter" title="Twitter" target="_blank">
<span class="screen-reader-text">Twitter</span>
</a>
</li>

<li>
<a href="https://www.facebook.com/calbucci" class="genericon genericon-facebook" title="Facebook" target="_blank">
<span class="screen-reader-text">Facebook</span>
</a>
</li>


<li>
<a href="http://www.linkedin.com/profile/view?id=732810" class="genericon genericon-linkedin-alt" title="LinkedIn" target="_blank">
<span class="screen-reader-text">LinkedIn</span>
</a>
</li>

</ul>
</div><!-- #social-links -->

<div class="site-info">
<div><a href="https://wordpress.com/?ref=footer_blog">Blog at WordPress.com</a>.</div>
<div><a href="https://wordpress.com/themes/writr/" title="Learn more about this theme">The Writr Theme</a>.</div>
</div><!-- .site-info -->
</footer><!-- #colophon -->
</div><!-- #page -->
<!-- wpcom_wp_footer -->
<script type="text/javascript">
/* <![CDATA[ */
jQuery(document).ready( function($) {
function doFollowingHover() {
$('#wp-admin-bar-follow > a').unbind( '.unfollow' );
$('#wp-admin-bar-follow > a').bind( 'mouseover.unfollow', function() {
$(this).html( "Unfollow" ).parent( 'li' ).addClass( 'unfollow' );
});
$('#wp-admin-bar-follow > a').bind( 'mouseout.unfollow', function() {
$(this).html( "Following" ).parent( 'li' ).removeClass( 'unfollow' );
});
}

$('#wp-admin-bar-follow > a').click( function( e ) {
$('#wp-admin-bar-follow > a').unbind( '.unfollow' );
e.preventDefault();
var link = $( this ), li = $( '#wp-admin-bar-follow' ), timeout = 0;
if ( li.hasClass( 'subscribed' ) ) {
li.removeClass( 'subscribed' ).removeClass( 'unfollow' );
link.html( "Follow" );
$('body').append( $( 'div.wpcom-bubble' ).removeClass( 'fadein' ) ).off( 'click.bubble' );
var action = 'ab_unsubscribe_from_blog';
} else {
li.addClass( 'subscribed' ).removeClass( 'unfollow' );
link.html( "Following" );
var left = 131 - link.width();
li.append( $( 'div.wpcom-bubble' ).css( { left: '-' + left + 'px' } ) );
$( 'div.bubble-txt', 'div.wpcom-bubble' ).html( "New posts from this blog will now appear in <a target=\"_blank\" href=\"http:\/\/wordpress.com\/\">your reader<\/a>." );

$( 'div.wpcom-bubble.action-bubble' ).addClass( 'fadein' );

setTimeout( function() {
$('body').on( 'click.bubble touchstart.bubble', function(e) {
if ( !$(e.target).hasClass('wpcom-bubble') && !$(e.target).parents( 'div.wpcom-bubble' ).length )
hideBubble();
});
setTimeout( hideBubble, 10000 );
}, 500 );

var action = 'ab_subscribe_to_blog';
$('#wp-admin-bar-follow > a').bind( 'mouseout.shift', function() {
doFollowingHover();
$(this).unbind( '.shift' );
});
}

var nonce = link.attr( 'href' ).split( '_wpnonce=' );
nonce = nonce[1];

$.post( "http:\/\/blog.calbucci.com\/wp-admin\/admin-ajax.php", {
'action': action,
'_wpnonce': nonce,
'source': 'admin_bar',
'blog_id': 80451878				});
});
});
/* ]]> */
</script>
<script type='text/javascript' src='//0.gravatar.com/js/gprofiles.js?ver=201506x'></script>
<script type='text/javascript'>
/* <![CDATA[ */
var WPGroHo = {"my_hash":"002201b3a1df4df7adc77ab91d675568"};
/* ]]> */
</script>
<script type='text/javascript' src='https://s2.wp.com/wp-content/mu-plugins/gravatar-hovercards/wpgroho.js?m=1380573781g'></script>

<script>
//initialize and attach hovercards to all gravatars
jQuery( document ).ready( function( $ ) {

if (typeof Gravatar === "undefined"){
return;
}

if ( typeof Gravatar.init !== "function" ) {
return;
}

Gravatar.profile_cb = function( hash, id ) {
WPGroHo.syncProfileData( hash, id );
};
Gravatar.my_hash = WPGroHo.my_hash;
Gravatar.init( 'body', '#wp-admin-bar-my-account' );
});
</script>

<div style="display:none">
</div>
<div id="report-form-window" style="display:none;"></div><script type='text/javascript'>
/* <![CDATA[ */
var HighlanderComments = {"loggingInText":"Logging In\u2026","submittingText":"Posting Comment\u2026","postCommentText":"Post Comment","connectingToText":"Connecting to %s","commentingAsText":"%1$s: You are commenting using your %2$s account.","logoutText":"Log Out","loginText":"Log In","connectURL":"https:\/\/calbucci.wordpress.com\/public.api\/connect\/?action=request","logoutURL":"https:\/\/calbucci.wordpress.com\/wp-login.php?action=logout&_wpnonce=0fad6596fa","homeURL":"http:\/\/blog.calbucci.com\/","postID":"691","gravDefault":"identicon","enterACommentError":"Please enter a comment","enterEmailError":"Please enter your email address here","invalidEmailError":"Invalid email address","enterAuthorError":"Please enter your name here","gravatarFromEmail":"This picture will show whenever you leave a comment. Click to customize it.","logInToExternalAccount":"Log in to use details from one of these accounts.","change":"Change","changeAccount":"Change Account","comment_registration":"","userIsLoggedIn":"1","isJetpack":"0"};
/* ]]> */
</script>
<script type='text/javascript' src='https://s0.wp.com/_static/??/wp-content/js/jquery/jquery.autoresize.js,/wp-content/mu-plugins/highlander-comments/script.js?m=1414003384j'></script>

<script type="text/javascript">
WPCOM_sharing_counts = {"http:\/\/blog.calbucci.com\/2015\/01\/27\/attention-cannibalization\/":691}	</script>
<script type="text/javascript">
jQuery(document).on( 'ready post-load', function(){
jQuery( 'a.share-twitter' ).on( 'click', function() {
window.open( jQuery(this).attr( 'href' ), 'wpcomtwitter', 'menubar=1,resizable=1,width=600,height=350' );
return false;
});
});
</script>
<script type="text/javascript">
jQuery(document).on( 'ready post-load', function(){
jQuery( 'a.share-facebook' ).on( 'click', function() {
window.open( jQuery(this).attr( 'href' ), 'wpcomfacebook', 'menubar=1,resizable=1,width=600,height=400' );
return false;
});
});
</script>
<script type="text/javascript">
jQuery(document).on( 'ready post-load', function(){
jQuery( 'a.share-linkedin' ).on( 'click', function() {
window.open( jQuery(this).attr( 'href' ), 'wpcomlinkedin', 'menubar=1,resizable=1,width=580,height=450' );
return false;
});
});
</script>
<script type="text/javascript">
jQuery(document).on( 'ready post-load', function(){
jQuery( 'a.share-google-plus-1' ).on( 'click', function() {
window.open( jQuery(this).attr( 'href' ), 'wpcomgoogle-plus-1', 'menubar=1,resizable=1,width=480,height=550' );
return false;
});
});
</script>
<script type='text/javascript' src='https://s1.wp.com/wp-content/mu-plugins/akismet-3.0/_inc/form.js?m=1404442431g'></script>
<script type='text/javascript'>
/* <![CDATA[ */
var thickboxL10n = {"next":"Next >","prev":"< Prev","image":"Image","of":"of","close":"Close","noiframes":"This feature requires inline frames. You have iframes disabled or your browser does not support them.","loadingAnimation":"https:\/\/s1.wp.com\/wp-includes\/js\/thickbox\/loadingAnimation.gif"};
var thickboxL10n = {"next":"Next >","prev":"< Prev","image":"Image","of":"of","close":"Close","noiframes":"This feature requires inline frames. You have iframes disabled or your browser does not support them.","loadingAnimation":"https:\/\/s1.wp.com\/wp-includes\/js\/thickbox\/loadingAnimation.gif"};
/* ]]> */
</script>
<script type='text/javascript' src='https://s2.wp.com/_static/??-eJxdjUsOgCAQQy8kEomyM55FB6KDMhAB9fjiwk9ctcnrS/nuGRIsSenATeC9skhs6NcyZ2lCwfMAHEVN8eJKbwjaHzf6unFCmAd3POXv28T8kkakz0/m4Cx7f03obFvVopJCykaYEwRwOks='></script>
<script type='text/javascript'>
/* <![CDATA[ */
var kissmetrics_api = {"api_key":"39e3e583aca28fab00c85ad9019b51d9f073fc6d","username":"mcalbucci","user_id":"63737"};
var kissmetrics_queries = {"events":[],"properties":[{"Browser_Type":"desktop"}]};
/* ]]> */
</script>
<script type='text/javascript' src='http://s0.wp.com/wp-content/mu-plugins/kissmetrics/kissmetrics.js?m=1417525689g&#038;ver=20141202'></script>
<script type='text/javascript'>
/* <![CDATA[ */
var JetpackEmojiSettings = {"base_url":"http:\/\/s0.wp.com\/wp-content\/mu-plugins\/emoji\/twemoji\/"};
/* ]]> */
</script>
<script type='text/javascript'>
/* <![CDATA[ */
var wpNotesArgs = {"cacheBuster":"20150109","iframeUrl":"https:\/\/widgets.wp.com\/notifications\/","iframeAppend":"2","iframeScroll":"yes"};
/* ]]> */
</script>
<script type='text/javascript'>
/* <![CDATA[ */
var wpcom_tos_report_form = {"ajaxurl":"\/wp-admin\/admin-ajax.php","isLoggedoutUser":"","post_ID":"691","current_url":"http:\/\/blog.calbucci.com\/2015\/01\/27\/attention-cannibalization","report_this_content":"Report this content"};
/* ]]> */
</script>
<script type='text/javascript'>
/* <![CDATA[ */
var sharing_js_options = {"lang":"en","counts":"1"};
/* ]]> */
</script>
<script type='text/javascript' src='https://s1.wp.com/_static/??-eJyNkd1qwzAMhV9ormi2Qm/GnsVx3EROLBlJbvr4M2OFtYykN/pB3zlCCNbiApNFMrAp5qhQag+roAkkBfJXHL0h0yHpG+zQOmNxC9LsLhyqugve7jKksNShCRpWaYiigSUeMr5k/FM8g7m6stQRSSFmTgi2PuRdPkUrPszuNfqB+ntO3zx6pn+PaeNc1XyY4sYCYou/sY1zZnLXbpf3Q1voei/bsLE6iYXF2lckP/cbysLamMWjgE5ekMZ7bqKv/Hn86LrT+/F8PqVv34vZmA=='></script>
<script type="text/javascript">
// <![CDATA[
(function() {
try{
if ( window.external &&'msIsSiteMode' in window.external) {
if (window.external.msIsSiteMode()) {
var jl = document.createElement('script');
jl.type='text/javascript';
jl.async=true;
jl.src='/wp-content/plugins/ie-sitemode/custom-jumplist.php';
var s = document.getElementsByTagName('script')[0];
s.parentNode.insertBefore(jl, s);
}
}
}catch(e){}
})();
// ]]>
</script><script src="//stats.wp.com/w.js?39" type="text/javascript" async defer></script>
<script type="text/javascript">
_tkq = window._tkq || [];
_stq = window._stq || [];
_tkq.push(['identifyUser', 63737, 'mcalbucci']);
_tkq.push(['storeContext', {'blog_id':'80451878','blog_tz':'-8','user_lang':'','blog_lang':'en'}]);
_stq.push(['view', {'blog':'80451878','v':'wpcom','tz':'-8','user':'1','user_id':'63737','post':'691','subd':'calbucci'}]);
_stq.push(['extra', {'crypt':'UE40eW5QN0p8M2Y/RE1zNDZ8S252Wis9XUQyb3YrcUVIU2R0VH5TdXRaZ1pyV0ZjN3RuNllzeFhCbHFqbnlOXUJ3aDBpTlZheml1VTV4WXBQa29MPWpjXVYzQVJvS3FLNVgsLUEvcnV6eklEdC1sOVJselomaHlMekssV35rLTdOeVsyYmp4VktKfjdUVUw4TEpwM2tXM3NfNkxINWlMeXNjcFptci1qTTlJNT1rLn45MkQzRFY0eStsenNwRUx1by5Udl8wZzdtS1hrNE1HLl1DP0tVU19mM1BDTmkyQ1ZNP2NTTUkzYldjQ3lTTEUwMWF5TTJKWFFWbnNsRGQ1RXxxdl91S1U1WjlHZ0p5ZFBzRzMwY1Nud1cufmdLJW8lYzRMdmRbfE1hdU9NWDdHeCxWcW9PRW1bTC1teS9bZU8seFlQTHBuak5QQ3djbn5mbFlf'}]);
</script>
<noscript><img src="http://pixel.wp.com/b.gif?v=noscript" style="height:0px;width:0px;overflow:hidden" alt="" /></noscript>
<script type="text/javascript">
(function() {
var request, b = document.body, c = 'className', cs = 'customize-support', rcs = new RegExp('(^|\\s+)(no-)?'+cs+'(\\s+|$)');

request = (function(){ var xhr = new XMLHttpRequest(); return ('withCredentials' in xhr); })();

b[c] = b[c].replace( rcs, ' ' );
b[c] += ( window.postMessage && request ? ' ' : ' no-' ) + cs;
}());
</script>
<div id="wpadminbar" class="nojq nojs ltr" role="navigation">
<a class="screen-reader-shortcut" href="#wp-toolbar" tabindex="1">Skip to toolbar</a>
<div class="quicklinks" id="wp-toolbar" role="navigation" aria-label="Top navigation toolbar." tabindex="0">
<ul id="wp-admin-bar-root-default" class="ab-top-menu">
<li id="wp-admin-bar-blog" class="menupop my-sites"><a class="ab-item"  aria-haspopup="true" href="https://wordpress.com/stats/day/blog.calbucci.com">My Sites</a><div class="ab-sub-wrapper"><ul id="wp-admin-bar-blog-default" class="ab-submenu">
<li id="wp-admin-bar-blog-info" class="current-site"><div class="ab-item ab-empty-item" ><div class="ab-site-icon"></div><span class="ab-site-title">M-Shaped Brain</span><span class="ab-site-description">calbucci.wordpress.com</span></div>		</li>
<li id="wp-admin-bar-switch-site" class="menupop"><a class="ab-item"  aria-haspopup="true" href="https://wordpress.com/stats">Switch Site</a><div class="ab-sub-wrapper"><ul id="wp-admin-bar-my-sites" class="ab-sub-secondary ab-submenu">
<li id="wp-admin-bar-blog-80451878" class="menupop"><a class="ab-item"  aria-haspopup="true" href="http://calbucci.wordpress.com/"><img src="http://s2.wp.com/i/wpmini-blue.png?m=1311732102g" alt="Blavatar" width="16" height="16" />M-Shaped Brain</a><div class="ab-sub-wrapper"><ul id="wp-admin-bar-blog-80451878-default" class="ab-submenu">
<li id="wp-admin-bar-blog-80451878-d"><a class="ab-item"  href="http://calbucci.wordpress.com/wp-admin/">WP Admin</a>		</li>
<li id="wp-admin-bar-blog-80451878-n"><a class="ab-item"  href="https://wordpress.com/post/80451878/new">New Post</a>		</li>
<li id="wp-admin-bar-blog-80451878-s"><a class="ab-item"  href="https://wordpress.com/stats/calbucci.wordpress.com">Stats</a>		</li>
<li id="wp-admin-bar-blog-80451878-c"><a class="ab-item"  href="http://calbucci.wordpress.com/wp-admin/edit-comments.php">Manage Comments</a>		</li>
<li id="wp-admin-bar-blog-80451878-v"><a class="ab-item"  href="http://calbucci.wordpress.com">Read Blog</a>		</li></ul></div>		</li></ul></div>		</li>
<li id="wp-admin-bar-view-site" class="mb-icon"><a class="ab-item"  href="http://blog.calbucci.com/">View Site</a>		</li>
<li id="wp-admin-bar-dashboard-top" class="mb-icon"><a class="ab-item"  href="http://calbucci.wordpress.com/wp-admin/">WP Admin</a>		</li>
<li id="wp-admin-bar-blog-stats" class="mb-icon"><a class="ab-item"  href="https://wordpress.com/stats/calbucci.wordpress.com">Stats</a>		</li>
<li id="wp-admin-bar-comments" class="mb-icon"><a class="ab-item"  href="http://calbucci.wordpress.com/wp-admin/edit-comments.php">Comments <span class="comment-moderation-count">2</span></a>		</li>
<li id="wp-admin-bar-edit" class="mb-icon"><a class="ab-item"  href="https://wordpress.com/post/80451878/691">Edit</a>		</li></ul><ul id="wp-admin-bar-publish" class="ab-submenu">
<li id="wp-admin-bar-publish-header" class="ab-submenu-header"><div class="ab-item ab-empty-item" >Publish</div>		</li>
<li id="wp-admin-bar-new-post" class="inline-action"><div class="ab-item ab-empty-item" ><a href="https://calbucci.wordpress.com/wp-admin/edit.php?post_type=post" class="ab-item primary mb-icon" id="wp-admin-bar-edit-post">Blog Posts</a><a href="https://wordpress.com/post/80451878/new" class="secondary" id="wp-admin-bar-new-post">Add</a></div>		</li>
<li id="wp-admin-bar-new-page" class="inline-action"><div class="ab-item ab-empty-item" ><a href="https://calbucci.wordpress.com/wp-admin/edit.php?post_type=page" class="ab-item primary mb-icon" id="wp-admin-bar-edit-page">Pages</a><a href="https://calbucci.wordpress.com/wp-admin/post-new.php?post_type=page" class="secondary" id="wp-admin-bar-new-page">Add</a></div>		</li>
<li id="wp-admin-bar-new-attachment" class="inline-action"><div class="ab-item ab-empty-item" ><a href="https://calbucci.wordpress.com/wp-admin/upload.php" class="ab-item primary mb-icon" id="wp-admin-bar-edit-attachment">Media</a><a href="https://calbucci.wordpress.com/wp-admin/media-new.php" class="secondary" id="wp-admin-bar-new-attachment">Add</a></div>		</li></ul><ul id="wp-admin-bar-look-and-feel" class="ab-submenu">
<li id="wp-admin-bar-look-and-feel-header" class="ab-submenu-header"><div class="ab-item ab-empty-item" >Look and Feel</div>		</li>
<li id="wp-admin-bar-themes" class="mb-icon"><a class="ab-item"  href="https://wordpress.com/themes/calbucci.wordpress.com">Themes</a>		</li>
<li id="wp-admin-bar-customize" class="hide-if-no-customize mb-icon"><a class="ab-item"  href="https://wordpress.com/customize/calbucci.wordpress.com">Customize</a>		</li>
<li id="wp-admin-bar-menus" class="mb-icon"><a class="ab-item"  href="https://calbucci.wordpress.com/wp-admin/nav-menus.php">Menus</a>		</li></ul><ul id="wp-admin-bar-configuration" class="ab-submenu">
<li id="wp-admin-bar-configuration-header" class="ab-submenu-header"><div class="ab-item ab-empty-item" >Configuration</div>		</li>
<li id="wp-admin-bar-sharing" class="mb-icon"><a class="ab-item"  href="https://calbucci.wordpress.com/wp-admin/options-general.php?page=sharing">Sharing</a>		</li>
<li id="wp-admin-bar-users" class="mb-icon"><a class="ab-item"  href="https://calbucci.wordpress.com/wp-admin/users.php">Users</a>		</li>
<li id="wp-admin-bar-blog-settings" class="mb-icon"><a class="ab-item"  href="https://wordpress.com/settings/general/calbucci.wordpress.com">Settings</a>		</li>
<li id="wp-admin-bar-upgrades" class="mb-icon"><a class="ab-item"  href="https://wordpress.com/plans/80451878">Upgrades</a>		</li></ul><ul id="wp-admin-bar-blog-secondary" class="ab-sub-secondary ab-submenu">
<li id="wp-admin-bar-shortlink"><a class="ab-item"  href="http://wp.me/p5rzcW-b9">Shortlink http://wp.me/p5rzcW-b9</a>Shortlink: <input id="adminbar-shortlink-input" value="http://wp.me/p5rzcW-b9" readonly="readonly" />		</li>
<li id="wp-admin-bar-theme-writr"><a class="ab-item"  href="http://theme.wordpress.com/themes/writr/">Theme: Writr</a>		</li>
<li id="wp-admin-bar-wpcom_report_url"><a class="ab-item"  href="http://en.wordpress.com/complaints/">Report this content</a>		</li></ul></div>		</li>
<li id="wp-admin-bar-jumptotop-button-menu"><a class="ab-item"  href="#"><div id="jumptotop"><span class="noticon noticon-top"></span></div></a>		</li>
<li id="wp-admin-bar-newdash" class="menupop"><a class="ab-item"  aria-haspopup="true" href="http://wordpress.com/">Reader</a><div class="ab-sub-wrapper"><ul id="wp-admin-bar-newdash-default" class="ab-submenu">
<li id="wp-admin-bar-following"><a class="ab-item"  href="http://wordpress.com/">Blogs I Follow</a>		</li></ul><ul id="wp-admin-bar-discover" class="ab-submenu">
<li id="wp-admin-bar-discover-header" class="ab-submenu-header"><div class="ab-item ab-empty-item" >Discover</div>		</li>
<li id="wp-admin-bar-discover-freshly-pressed"><a class="ab-item"  href="http://wordpress.com/fresh/">Freshly Pressed</a>		</li>
<li id="wp-admin-bar-discover-recommended-blogs"><a class="ab-item"  href="http://wordpress.com/recommendations/">Recommended Blogs</a>		</li>
<li id="wp-admin-bar-discover-find-friends"><a class="ab-item"  href="https://wordpress.com/me/find-friends/">Find Friends</a>		</li></ul><ul id="wp-admin-bar-my-activity" class="ab-submenu">
<li id="wp-admin-bar-my-activity-header" class="ab-submenu-header"><div class="ab-item ab-empty-item" >My Activity</div>		</li>
<li id="wp-admin-bar-my-activity-my-comments"><a class="ab-item"  href="http://wordpress.com/activities/comments/">My Comments</a>		</li>
<li id="wp-admin-bar-my-activity-my-likes"><a class="ab-item"  href="http://wordpress.com/activities/likes/">My Likes</a>		</li></ul></div>		</li>
<li id="wp-admin-bar-follow"><a class="ab-item"  href="http://wordpress.com/read?b=80451878&#038;_wpnonce=366da32a90"><span id="admin-bar-follow-link">Follow</span></a>		</li>
<li id="wp-admin-bar-stats"><a class="ab-item"  href="http://wordpress.com/stats/80451878"><div><script type='text/javascript'>var src;if(typeof(window.devicePixelRatio)=='undefined'||window.devicePixelRatio<2){src='/wp-includes/charts/admin-bar-hours-scale.php';}else{src='/wp-includes/charts/admin-bar-hours-scale-2x.php';}document.write('<img style=\'vertical-align: baseline\' src=\''+src+'\' alt=\'Stats\' title=\'Showing site views per hour for the last 48 hours. Click for full Site Stats.\' />');</script></div></a>		</li></ul><ul id="wp-admin-bar-top-secondary" class="ab-top-secondary ab-top-menu">
<li id="wp-admin-bar-notes" class="menupop"><div class="ab-item ab-empty-item" ><span id="wpnt-notes-unread-count" class="wpnt-loading wpn-unread">
<span class="noticon noticon-comment"></span>
</span></div><div id="wpnt-notes-panel2" style="display:none" lang="en" dir="ltr"><div class="wpnt-notes-panel-header"><span class="wpnt-notes-header"></span><span class="wpnt-notes-panel-link"></span></div></div>		</li>
<li id="wp-admin-bar-my-account" class="menupop with-avatar"><a class="ab-item"  aria-haspopup="true" href="https://wordpress.com/me/public-profile/"><img alt='' src='http://0.gravatar.com/avatar/002201b3a1df4df7adc77ab91d675568?s=32&#038;d=mm&#038;r=G' class='avatar avatar-32' height='32' width='32' /></a><div class="ab-sub-wrapper"><ul id="wp-admin-bar-user-actions" class="ab-submenu">
<li id="wp-admin-bar-user-info" class="user-info user-info-item"><div class="ab-item ab-empty-item" tabindex="-1"><img alt='' src='http://0.gravatar.com/avatar/002201b3a1df4df7adc77ab91d675568?s=128&#038;d=mm&#038;r=G' class='avatar avatar-128' height='128' width='128' /><span class="display-name">Marcelo Calbucci</span><span class="username"><a class="at-username" href="http://gravatar.com/mcalbucci">mcalbucci</a>&nbsp;&middot;&nbsp;<a href="https://calbucci.wordpress.com/wp-login.php?action=logout&amp;redirect_to=http%3A%2F%2Fblog.calbucci.com%2F&amp;_wpnonce=0fad6596fa">Sign Out</a></span></div>		</li>
<li id="wp-admin-bar-profile-header" class="ab-submenu-header"><div class="ab-item ab-empty-item" >Profile</div>		</li>
<li id="wp-admin-bar-account-settings"><a class="ab-item"  href="https://wordpress.com/me/public-profile/">Account Settings</a>		</li>
<li id="wp-admin-bar-trophy-case"><a class="ab-item"  href="https://wordpress.com/me/trophies">Trophy Case</a>		</li>
<li id="wp-admin-bar-billing"><a class="ab-item"  href="https://wordpress.com/me/billing">Billing</a>		</li>
<li id="wp-admin-bar-extras-header" class="ab-submenu-header"><div class="ab-item ab-empty-item" >Extras</div>		</li>
<li id="wp-admin-bar-help" class="user-info-item"><a class="ab-item"  href="http://en.support.wordpress.com/">Help</a>		</li>
<li id="wp-admin-bar-find-friends" class="user-info-item"><a class="ab-item"  href="https://wordpress.com/me/find-friends/">Find Friends</a>		</li></ul></div>		</li>
<li id="wp-admin-bar-ab-new-post"><a class="ab-item"  href="http://wordpress.com/post/80451878/new/"></a>		</li></ul>			</div>
<a class="screen-reader-shortcut" href="https://calbucci.wordpress.com/wp-login.php?action=logout&#038;_wpnonce=0fad6596fa">Log Out</a>
</div>


<script type="text/javascript">
/* <![CDATA[ */
var clickDebugLink;

jQuery(document).ready( function($) {
var single = true, w = 1000,
supe = false;

$(document.body).load(function(){ $('#wpadminbar img.grav-hashed').removeClass('grav-hashed'); }); // hack to hide the gravatar hovercard

if ( single && supe )
w = 1385;
else if ( supe )
w = 1200;

// debug bar extra
clickDebugLink = function( parent_id, obj ) {

$('#'+parent_id).children('div').hide();

document.getElementById( obj.href.substr( obj.href.indexOf( '#' ) + 1 ) ).style.display = 'block';
$( obj.href.substr( obj.href.indexOf( '#' ) ) ).show()

$(obj).parent().addClass('current').siblings('li').removeClass('current');

return false;
};

$( '#wpadminbar #wp-admin-bar-shortlink' ).click( function() {
$( '#adminbar-shortlink-input' ).select();
});

});
/* ]]> */
</script>
<div class="wpcom-bubble action-bubble">
<div class="bubble-txt"></div>
</div>
<script type="text/javascript">function hideBubble() { jQuery('body').append( jQuery( 'div.wpcom-bubble' ).removeClass( 'fadein' ) ).off( 'click.bubble touchstart.bubble' ); jQuery(document).off( 'scroll.bubble' ); };</script>
<script type="text/javascript">
jQuery( '#wp-admin-bar-jumptotop-button-menu a' ).click( function( e ) {
e.preventDefault();
jQuery( 'html, body' ).animate( { scrollTop: 0 }, 'fast' );
} );
function hideShowTbJumpToTop() {
var total_width = 0;
// Calculate total width taken by li's minus our button to see if we it'll
// overlap with other toolbar menus.
jQuery( '#wp-admin-bar-root-default > li' ).each( function() {
var self = jQuery( this );
if ( 'wp-admin-bar-jumptotop-button-menu' != self.attr( 'id' ) )
total_width += self.width();
});
if ( jQuery( '#wp-admin-bar-jumptotop-button-menu' ).position()['left'] - total_width < 10 ) {
jQuery( '#jumptotop' ).animate( { 'top': '-50px' }, 'fast' );
} else if (  jQuery( window ).scrollTop() >= 350 && parseInt( jQuery( '#jumptotop' ).css( 'top' ) ) < 0 ) {
if ( jQuery( '#wp-admin-bar-jumptotop-button-menu' ).position()['left'] - total_width < 10 )
return;
jQuery( '#jumptotop' ).animate( { 'top': 0 }, 'fast' );
} else if (  jQuery( window ).scrollTop() < 1 && parseInt( jQuery( '#jumptotop' ).css( 'top' ) ) >= 0 ) {
jQuery( '#jumptotop' ).animate( { 'top': '-50px' }, 'fast' );
}
}
// handle on page load if auto scrolling to a position
hideShowTbJumpToTop();
// bind to scrolll event
var jumpToTopTimer = null;
jQuery( window ).scroll( function() {
clearTimeout( jumpToTopTimer );
jumpToTopTimer = setTimeout( jumpToTopRefresh, 150 );
} );
var jumpToTopRefresh = function() {
hideShowTbJumpToTop();
};
</script>
<script>
if ( 'object' === typeof wpcom_mobile_user_agent_info ) {

wpcom_mobile_user_agent_info.init();
var mobileStatsQueryString = "";

if( false !== wpcom_mobile_user_agent_info.matchedPlatformName )
mobileStatsQueryString += "&x_" + 'mobile_platforms' + '=' + wpcom_mobile_user_agent_info.matchedPlatformName;

if( false !== wpcom_mobile_user_agent_info.matchedUserAgentName )
mobileStatsQueryString += "&x_" + 'mobile_devices' + '=' + wpcom_mobile_user_agent_info.matchedUserAgentName;

if( wpcom_mobile_user_agent_info.isIPad() )
mobileStatsQueryString += "&x_" + 'ipad_views' + '=' + 'views';

if( "" != mobileStatsQueryString ) {
new Image().src = document.location.protocol + '//pixel.wp.com/g.gif?v=wpcom-no-pv' + mobileStatsQueryString + '&baba=' + Math.random();
}

}
</script>
</body>
</html>
`

const googleHomepage = `

<!doctype html><html itemscope="" itemtype="http://schema.org/WebPage" lang="en"><head><meta content="Search the world's information, including webpages, images, videos and more. Google has many special features to help you find exactly what you're looking for." name="description"><meta content="noodp" name="robots"><meta content="/images/google_favicon_128.png" itemprop="image"><meta content="origin" id="mref" name="referrer"><title>Google</title>   <script>(function(){window.google={kEI:'zKjSVMXoLpK3oQSs14GoDQ',kEXPI:'3700315,4010073,4011550,4011551,4011557,4011558,4011559,4014789,4017578,4020347,4020562,4021073,4021587,4023678,4023709,4024207,4024626,4025090,4025124,4025127,4027916,4027921,4028063,4028135,4028465,4028522,4028551,4028585,4028706,4028716,4028730,4029215,8300096,8300108,8300111,8500393,10200083,10200903,10200905',authuser:0,j:{en:1,bv:21,pm:'p',u:'c9c918f0',qbp:0,rre:false},kSID:'zKjSVMXoLpK3oQSs14GoDQ'};google.kHL='en';})();(function(){google.lc=[];google.li=0;google.getEI=function(a){for(var b;a&&(!a.getAttribute||!(b=a.getAttribute("eid")));)a=a.parentNode;return b||google.kEI};google.getLEI=function(a){for(var b=null;a&&(!a.getAttribute||!(b=a.getAttribute("leid")));)a=a.parentNode;return b};google.https=function(){return"https:"==window.location.protocol};google.ml=function(){};google.time=function(){return(new Date).getTime()};google.log=function(a,b,e,f,l){var d=new Image,h=google.lc,g=google.li,c="",m=google.ls||"";d.onerror=d.onload=d.onabort=function(){delete h[g]};h[g]=d;if(!e&&-1==b.search("&ei=")){var k=google.getEI(f),c="&ei="+k;-1==b.search("&lei=")&&((f=google.getLEI(f))?c+="&lei="+f:k!=google.kEI&&(c+="&lei="+google.kEI))}a=e||"/"+(l||"gen_204")+"?atyp=i&ct="+a+"&cad="+b+c+m+"&zx="+google.time();/^http:/i.test(a)&&google.https()?(google.ml(Error("a"),!1,{src:a,glmm:1}),delete h[g]):(d.src=a,google.li=g+1)};google.y={};google.x=function(a,b){google.y[a.id]=[a,b];return!1};google.load=function(a,b,e){google.x({id:a+n++},function(){google.load(a,b,e)})};var n=0;})();google.kCSI={};
google.j.b=(!!location.hash&&!!location.hash.match('[#&]((q|fp)=|tbs=rimg|tbs=simg|tbs=sbi)'))
||(google.j.qbp==1);(function(){google.sn="webhp";google.timers={};google.startTick=function(a,b){google.timers[a]={t:{start:google.time()},bfr:!!b};window.performance&&window.performance.now&&(google.timers[a].wsrt=Math.floor(window.performance.now()))};google.tick=function(a,b,c){google.timers[a]||google.startTick(a);google.timers[a].t[b]=c||google.time()};google.startTick("load",!0);
google.aft=function(a,b,c){};google.iml=function(a,b){google.tick("iml",a.id||a.src||a.name,b)};
try{google.pt=window.chrome&&window.chrome.csi&&Math.floor(window.chrome.csi().pageT);}catch(d){};})();
(function(){'use strict';var g=this,h=Date.now||function(){return+new Date};var u=function(b,c){if(null===c)return!1;if("contains"in b&&1==c.nodeType)return b.contains(c);if("compareDocumentPosition"in b)return b==c||Boolean(b.compareDocumentPosition(c)&16);for(;c&&b!=c;)c=c.parentNode;return c==b};var w=function(b,c){return function(a){a||(a=window.event);return c.call(b,a)}},x=function(b){b=b.target||b.srcElement;!b.getAttribute&&b.parentNode&&(b=b.parentNode);return b},z="undefined"!=typeof navigator&&/Macintosh/.test(navigator.userAgent),A="undefined"!=typeof navigator&&!/Opera/.test(navigator.userAgent)&&/WebKit/.test(navigator.userAgent),E=function(b){var c=b.which||b.keyCode||b.key;A&&3==c&&(c=13);if(13!=c&&32!=c)return!1;var a=x(b),d=(a.getAttribute("role")||a.type||a.tagName).toUpperCase(),e;(e="keydown"!=b.type)||("getAttribute"in a?(e=(a.getAttribute("role")||a.type||a.tagName).toUpperCase(),e="TEXT"!=e&&"TEXTAREA"!=e&&"PASSWORD"!=e&&"SEARCH"!=e&&("COMBOBOX"!=
e||"INPUT"!=a.tagName.toUpperCase())&&!a.isContentEditable):e=!1,e=!e);if(e||b.ctrlKey||b.shiftKey||b.altKey||b.metaKey||"INPUT"==a.tagName.toUpperCase()&&a.type&&a.type.toUpperCase()in B&&32==c)return!1;if(b.originalTarget&&b.originalTarget!=a)return!0;(b=a.tagName in C)||(b=a.getAttributeNode("tabindex"),b=null!=b&&b.specified);if(!(b&&0<=a.tabIndex)||a.disabled)return!1;a="INPUT"!=a.tagName.toUpperCase()||a.type;b=!(d in D)&&13==c;return(0==D[d]%c||b)&&!!a},C={A:1,INPUT:1,TEXTAREA:1,SELECT:1,BUTTON:1},D={A:13,BUTTON:0,CHECKBOX:32,COMBOBOX:13,LINK:13,LISTBOX:13,MENU:0,MENUBAR:0,MENUITEM:0,MENUITEMCHECKBOX:0,MENUITEMRADIO:0,OPTION:13,RADIO:32,RADIOGROUP:32,RESET:0,SUBMIT:0,TAB:0,TABLIST:0,TREE:13,TREEITEM:13},B={CHECKBOX:1,OPTION:1,RADIO:1};var F=function(){this.o=this.i=null},I=function(b,c){var a=G;a.i=b;a.o=c;return a};F.prototype.k=function(){var b=this.i;this.i&&this.i!=this.o?this.i=this.i.__owner||this.i.parentNode:this.i=null;return b};var J=function(){this.p=[];this.i=0;this.o=null;this.s=!1};J.prototype.k=function(){if(this.s)return G.k();if(this.i!=this.p.length){var b=this.p[this.i];this.i++;b!=this.o&&b&&b.__owner&&(this.s=!0,I(b.__owner,this.o));return b}return null};var G=new F,K=new J;var M=function(){this.B=[];this.i=[];this.k=[];this.s={};this.o=null;this.p=[];L(this,"_custom")},N="undefined"!=typeof navigator&&/iPhone|iPad|iPod/.test(navigator.userAgent),O=/\s*;\s*/,Q=function(b,c){return function(a){var d=c;if("_custom"==d){d=a.detail;if(!d||!d._type)return;d=d._type}var e;var f=d;"click"==f&&(z&&a.metaKey||!z&&a.ctrlKey||2==a.which||null==a.which&&4==a.button||a.shiftKey)?f="clickmod":E(a)&&(f="clickkey");var k=a.srcElement||a.target,d=P(f,a,k,"",null),l,n;a.path?(K.p=a.path,K.i=0,K.o=this,K.s=!1,n=K):n=I(k,this);for(var p;p=n.k();){var q=e=p;l=f;p=q.__jsaction;if(!p){p=
{};q.__jsaction=p;var r=void 0,r=null;"getAttribute"in q&&(r=q.getAttribute("jsaction"));if(r)for(var q=r.split(O),r=0,X=q?q.length:0;r<X;r++){var t=q[r];if(t){var y=t.indexOf(":"),H=-1!=y,Y=H?t.substr(0,y).replace(/^\s+/,"").replace(/\s+$/,""):"click",t=H?t.substr(y+1).replace(/^\s+/,"").replace(/\s+$/,""):t;p[Y]=t}}}"clickkey"==l?l="click":"click"!=l||p.click||(l="clickonly");l={v:l,action:p[l]||"",event:null,D:!1};d=P(l.v,l.event||a,k,l.action||"",e,d.timeStamp);if(l.D||l.action)break}if(l&&l.action){if(k=
"clickkey"==f)k=x(a),k=(k.type||k.tagName).toUpperCase(),(k=32==(a.which||a.keyCode||a.key)&&"CHECKBOX"!=k)||(n=x(a),k=(n.getAttribute("role")||n.tagName).toUpperCase(),n=n.type,k="BUTTON"==k||!!n&&!(n.toUpperCase()in B));k&&(a.preventDefault?a.preventDefault():a.returnValue=!1);if("mouseenter"==f||"mouseleave"==f)if(k=a.relatedTarget,!("mouseover"==a.type&&"mouseenter"==f||"mouseout"==a.type&&"mouseleave"==f)||k&&(k===e||u(e,k)))d.action="",d.actionElement=null;else{var f={},m;for(m in a)"function"!==
typeof a[m]&&"srcElement"!==m&&"target"!==m&&(f[m]=a[m]);f.type="mouseover"==a.type?"mouseenter":"mouseleave";f.target=f.srcElement=e;f.bubbles=!1;d.event=f;d.targetElement=e}}else d.action="",d.actionElement=null;e=d;b.o&&(m=P(e.eventType,e.event,e.targetElement,e.action,e.actionElement,e.timeStamp),"clickonly"==m.eventType&&(m.eventType="click"),b.o(m,!0));if(e.actionElement)if("A"!=e.actionElement.tagName||"click"!=e.eventType&&"clickmod"!=e.eventType||(a.preventDefault?a.preventDefault():a.returnValue=
!1),b.o)b.o(e);else{var v;if((m=g.document)&&!m.createEvent&&m.createEventObject)try{v=m.createEventObject(a)}catch(ca){v=a}else v=a;e.event=v;b.p.push(e)}}},P=function(b,c,a,d,e,f){return{eventType:b,event:c,targetElement:a,action:d,actionElement:e,timeStamp:f||h()}},aa=function(b,c){return function(a){var d=b,e=c,f=!1;"mouseenter"==d?d="mouseover":"mouseleave"==d&&(d="mouseout");if(a.addEventListener){if("focus"==d||"blur"==d||"error"==d||"load"==d)f=!0;a.addEventListener(d,e,f)}else a.attachEvent&&("focus"==d?d="focusin":"blur"==d&&(d="focusout"),e=w(a,e),a.attachEvent("on"+d,e));return{v:d,w:e,C:f}}},L=function(b,c){if(!b.s.hasOwnProperty(c)){var a=Q(b,c),d=aa(c,a);b.s[c]=a;b.B.push(d);for(a=0;a<b.i.length;++a){var e=b.i[a];e.k.push(d.call(null,e.i))}"click"==c&&L(b,"keydown")}};M.prototype.w=function(b){return this.s[b]};var V=function(b){var c=R,a=new ba(b);n:{for(var d=0;d<c.i.length;d++)if(S(c.i[d],b)){b=!0;break n}b=!1}if(b)return c.k.push(a),a;T(c,a);c.i.push(a);U(c);return a},U=function(b){for(var c=b.k.concat(b.i),a=[],d=[],e=0;e<b.i.length;++e){var f=b.i[e];W(f,c)?(a.push(f),Z(f)):d.push(f)}for(e=0;e<b.k.length;++e)f=b.k[e],W(f,c)?a.push(f):(d.push(f),T(b,f));b.i=d;b.k=a},T=function(b,c){var a=c.i;N&&(a.style.cursor="pointer");for(a=0;a<b.B.length;++a)c.k.push(b.B[a].call(null,c.i))},ba=function(b){this.i=b;this.k=[]},S=function(b,c){for(var a=b.i,d=c;a!=d&&d.parentNode;)d=d.parentNode;return a==d},W=function(b,c){for(var a=0;a<c.length;++a)if(c[a].i!=b.i&&S(c[a],b.i))return!0;return!1},Z=function(b){for(var c=0;c<b.k.length;++c){var a=b.i,d=b.k[c];a.removeEventListener?a.removeEventListener(d.v,d.w,d.C):a.detachEvent&&a.detachEvent("on"+d.v,d.w)}b.k=[]};var R=new M;V(window.document.documentElement);L(R,"click");L(R,"focus");L(R,"focusin");L(R,"blur");L(R,"focusout");L(R,"error");L(R,"load");L(R,"change");L(R,"dblclick");L(R,"input");L(R,"keyup");L(R,"keydown");L(R,"keypress");L(R,"mousedown");L(R,"mouseenter");L(R,"mouseleave");L(R,"mouseout");L(R,"mouseover");L(R,"mouseup");L(R,"touchstart");L(R,"touchend");L(R,"touchcancel");L(R,"speech");window.google.jsad=function(b){var c=R;c.o=b;c.p&&(0<c.p.length&&b(c.p),c.p=null)};window.google.jsaac=function(b){return V(b)};window.google.jsarc=function(b){var c=R;Z(b);for(var a=!1,d=0;d<c.i.length;++d)if(c.i[d]===b){c.i.splice(d,1);a=!0;break}if(!a)for(d=0;d<c.k.length;++d)if(c.k[d]===b){c.k.splice(d,1);break}U(c)};}).call(window);(function(){'use strict';var f=this,g=function(d,e){var b=d.split("."),a=f;b[0]in a||!a.execScript||a.execScript("var "+b[0]);for(var c;b.length&&(c=b.shift());)b.length||void 0===e?a[c]?a=a[c]:a=a[c]={}:a[c]=e};var h=[];g("google.jsc.xx",h);g("google.jsc.x",function(d){h.push(d)});}).call(window);google.arwt=function(a){a.href=document.getElementById(a.id.substring(1)).href;return!0};</script><style>[dir='ltr'],[dir='rtl']{unicode-bidi:-webkit-isolate;unicode-bidi:isolate}bdo[dir='ltr'],bdo[dir='rtl']{unicode-bidi:bidi-override;unicode-bidi:-webkit-isolate-override;unicode-bidi:isolate-override}#cnt{padding-top:0;position:relative}#subform_ctrl{min-height:11px}.gb_e .gbqfi::before{left:-56px;top:-35px}.gb_E .gbqfb:focus .gbqfi{outline:1px dotted #fff}@-webkit-keyframes gb__a{0%{opacity:0}50%{opacity:1}}@keyframes gb__a{0%{opacity:0}50%{opacity:1}}#gb#gb a.gb_f,#gb#gb a.gb_g{color:#404040;text-decoration:none}#gb#gb a.gb_g:hover,#gb#gb a.gb_g:focus{color:#000;text-decoration:underline}.gb_h.gb_i{display:none;padding-left:15px;vertical-align:middle}.gb_h.gb_i:first-child{padding-left:0}.gb_j.gb_i{display:inline-block;-webkit-flex:0 1 auto;flex:0 1 auto;-webkit-flex:0 1 main-size;flex:0 1 main-size;display:-webkit-flex;display:flex}.gb_k .gb_j{display:none}.gb_h .gb_g{display:inline-block;line-height:24px;outline:none;vertical-align:middle}.gb_j .gb_g{min-width:60px;overflow:hidden;-webkit-flex:0 1 auto;flex:0 1 auto;-webkit-flex:0 1 main-size;flex:0 1 main-size;text-overflow:ellipsis}.gb_l .gb_j .gb_g{min-width:0}.gb_m .gb_j .gb_g{width:0!important}.gb_n .gb_g{font-weight:bold;text-shadow:0 1px 1px rgba(255,255,255,.9)}.gb_o .gb_g{font-weight:bold;text-shadow:0 1px 1px rgba(0,0,0,.6)}#gb#gb.gb_o a.gb_g{color:#fff}.gb_C .gb_D{background-position:-326px -52px;opacity:.55}.gb_n .gb_C .gb_D{background-position:-97px -57px;opacity:.7}.gb_o .gb_C .gb_D{background-position:-214px 0;opacity:1}.gb_Oc{left:0;min-width:1152px;position:absolute;top:0;-webkit-user-select:none;width:100%}.gb_Vb{font:13px/27px Arial,sans-serif;position:relative;height:60px;width:100%}.gb_ba .gb_Vb{height:28px}#gba{height:60px}#gba.gb_ba{height:28px}#gba.gb_Pc{height:90px}#gba.gb_Pc.gb_ba{height:58px}.gb_Vb>.gb_i{height:60px;line-height:58px;vertical-align:middle}.gb_ba .gb_Vb>.gb_i{height:28px;line-height:26px}.gb_Vb::before{background:#e5e5e5;bottom:0;content:'';display:none;height:1px;left:0;position:absolute;right:0}.gb_Vb{background:#f1f1f1}.gb_Qc .gb_Vb{background:#fff}.gb_Qc .gb_Vb::before,.gb_ba .gb_Vb::before{display:none}.gb_n .gb_Vb,.gb_o .gb_Vb,.gb_ba .gb_Vb{background:transparent}.gb_n .gb_Vb::before{background:#e1e1e1;background:rgba(0,0,0,.12)}.gb_o .gb_Vb::before{background:#333;background:rgba(255,255,255,.2)}.gb_i{display:inline-block;-webkit-flex:0 0 auto;flex:0 0 auto;-webkit-flex:0 0 main-size;flex:0 0 main-size}.gb_i.gb_Rc{float:right;-webkit-order:1;order:1}.gb_Sc{white-space:nowrap;display:-webkit-flex;display:flex;margin-left:0!important;margin-right:0!important}.gb_i{margin-left:0!important;margin-right:0!important}.gb_Ta{background-image:url('//ssl.gstatic.com/gb/images/i1_71651352.png');-webkit-background-size:356px 144px;background-size:356px 144px}@media (min-resolution:1.25dppx),(-webkit-min-device-pixel-ratio:1.25),(min-device-pixel-ratio:1.25){.gb_Ta{background-image:url('//ssl.gstatic.com/gb/images/i2_9ef0f6fa.png')}}.gb_kb{display:inline-block;padding:0 0 0 15px;vertical-align:middle}.gb_kb:first-child,#gbsfw:first-child+.gb_kb{padding-left:0}.gb_Za{position:relative}.gb_D{display:inline-block;outline:none;vertical-align:middle;-webkit-border-radius:2px;border-radius:2px;-webkit-box-sizing:border-box;box-sizing:border-box;height:30px;width:30px}#gb#gb a.gb_D{color:#404040;cursor:default;text-decoration:none}#gb#gb a.gb_D:hover,#gb#gb a.gb_D:focus{color:#000}.gb_ja{border-color:transparent;border-bottom-color:#fff;border-style:dashed dashed solid;border-width:0 8.5px 8.5px;display:none;position:absolute;left:6.5px;top:37px;z-index:1;height:0;width:0;-webkit-animation:gb__a .2s;animation:gb__a .2s}.gb_ka{border-color:transparent;border-style:dashed dashed solid;border-width:0 8.5px 8.5px;display:none;position:absolute;left:6.5px;z-index:1;height:0;width:0;-webkit-animation:gb__a .2s;animation:gb__a .2s;border-bottom-color:#ccc;border-bottom-color:rgba(0,0,0,.2);top:36px}x:-o-prefocus,div.gb_ka{border-bottom-color:#ccc}.gb_H{background:#fff;border:1px solid #ccc;border-color:rgba(0,0,0,.2);-webkit-box-shadow:0 2px 10px rgba(0,0,0,.2);box-shadow:0 2px 10px rgba(0,0,0,.2);display:none;outline:none;overflow:hidden;position:absolute;right:0;top:44px;-webkit-animation:gb__a .2s;animation:gb__a .2s;-webkit-border-radius:2px;border-radius:2px;-webkit-user-select:text}.gb_kb.gb_La .gb_ja,.gb_kb.gb_La .gb_ka,.gb_kb.gb_La .gb_H{display:block}.gb_gc{position:absolute;right:0;top:44px;z-index:-1}.gb_ba .gb_ja,.gb_ba .gb_ka,.gb_ba .gb_H{margin-top:-10px}.gb_7{-webkit-background-size:32px 32px;background-size:32px 32px;-webkit-border-radius:50%;border-radius:50%;display:block;margin:-1px;height:32px;width:32px}.gb_7:hover,.gb_7:focus{-webkit-box-shadow:0 1px 0 rgba(0,0,0,.15);box-shadow:0 1px 0 rgba(0,0,0,.15)}.gb_7:active{-webkit-box-shadow:inset 0 2px 0 rgba(0,0,0,.15);box-shadow:inset 0 2px 0 rgba(0,0,0,.15)}.gb_7:active::after{background:rgba(0,0,0,.1);-webkit-border-radius:50%;border-radius:50%;content:'';display:block;height:100%}.gb_8:not(.gb_e) .gb_7::before,.gb_8:not(.gb_e) .gb_9::before{content:none}.gb_aa{cursor:pointer;line-height:30px;min-width:30px;overflow:hidden;vertical-align:middle;width:auto;text-overflow:ellipsis}.gb_ba .gb_aa,.gb_ba .gb_ca{line-height:26px}#gb#gb.gb_ba a.gb_aa,.gb_ba .gb_ca{color:#666;font-size:11px;height:auto}#gb#gb.gb_ba a.gb_aa:hover,#gb#gb.gb_ba a.gb_aa:focus{color:#000}.gb_da{border-top:4px solid #404040;border-left:4px dashed transparent;border-right:4px dashed transparent;display:inline-block;margin-left:6px;vertical-align:middle}.gb_ba .gb_da{border-top-color:#999}.gb_ea:hover .gb_da{border-top-color:#000}.gb_n .gb_aa{font-weight:bold;text-shadow:0 1px 1px rgba(255,255,255,.9)}.gb_o .gb_aa{font-weight:bold;text-shadow:0 1px 1px rgba(0,0,0,.6)}#gb#gb.gb_o.gb_o a.gb_aa{color:#fff}.gb_o.gb_o .gb_da{border-top-color:#fff}.gb_n .gb_7,.gb_o .gb_7{-webkit-box-shadow:0 1px 2px rgba(0,0,0,.2);box-shadow:0 1px 2px rgba(0,0,0,.2)}.gb_n .gb_7:hover,.gb_o .gb_7:hover,.gb_n .gb_7:focus,.gb_o .gb_7:focus{-webkit-box-shadow:0 1px 0 rgba(0,0,0,.15),0 1px 2px rgba(0,0,0,.2);box-shadow:0 1px 0 rgba(0,0,0,.15),0 1px 2px rgba(0,0,0,.2)}.gb_fa .gb_ga,.gb_ha .gb_ga{position:absolute;right:1px}.gb_ga.gb_i,.gb_ia.gb_i,.gb_ea.gb_i{-webkit-flex:0 1 auto;flex:0 1 auto;-webkit-flex:0 1 main-size;flex:0 1 main-size}.gb_8.gb_m .gb_aa{width:30px!important}.gb_Wb{display:none!important}.gb_Z{background:#f8f8f8;border:1px solid #c6c6c6;display:inline-block;line-height:28px;padding:0 12px;-webkit-border-radius:2px;border-radius:2px}#gb a.gb_Z.gb_Z{color:#666;cursor:default;text-decoration:none}.gb_0{border:1px solid #4285f4;font-weight:bold;outline:none;background:#4285f4;background:-webkit-linear-gradient(top,#4387fd,#4683ea);background:linear-gradient(top,#4387fd,#4683ea);filter:progid:DXImageTransform.Microsoft.gradient(startColorstr=#4387fd,endColorstr=#4683ea,GradientType=0)}#gb a.gb_0.gb_0{color:#fff}.gb_0:hover{-webkit-box-shadow:0 1px 0 rgba(0,0,0,.15);box-shadow:0 1px 0 rgba(0,0,0,.15)}.gb_0:active{-webkit-box-shadow:inset 0 2px 0 rgba(0,0,0,.15);box-shadow:inset 0 2px 0 rgba(0,0,0,.15);background:#3c78dc;background:-webkit-linear-gradient(top,#3c7ae4,#3f76d3);background:linear-gradient(top,#3c7ae4,#3f76d3);filter:progid:DXImageTransform.Microsoft.gradient(startColorstr=#3c7ae4,endColorstr=#3f76d3,GradientType=0)}.gb_1{display:inline-block;line-height:normal;position:relative;z-index:987}#gbsfw{min-width:400px;overflow:visible}.gb_Ka,#gbsfw.gb_La{display:block;outline:none}#gbsfw.gb_P iframe{display:none}.gb_Ma{padding:118px 0;text-align:center}.gb_Na{background:no-repeat center 0;color:#aaa;font-size:13px;line-height:20px;padding-top:76px;background-image:-webkit-image-set(url('//ssl.gstatic.com/gb/images/a/f5cdd88b65.png') 1x,url('//ssl.gstatic.com/gb/images/a/133fc21e88.png') 2x)}.gb_Na a{color:#4285f4;text-decoration:none}.gb_Oa{min-width:127px;overflow:hidden;position:relative;z-index:987}.gb_Pa{position:absolute;padding:0 20px 0 15px}.gb_Qa .gb_Pa{right:100%;margin-right:-127px}.gb_Ra{display:inline-block;outline:none;vertical-align:middle}.gb_Sa .gb_Ra{position:relative;top:2px}.gb_Ra .gb_Ta,.gb_Ua{display:block}.gb_Va{border:none;display:block;visibility:hidden}.gb_Ra .gb_Ta{background-position:0 -105px;height:33px;width:92px}.gb_Ua{background-repeat:no-repeat}.gb_o .gb_Ra .gb_Ta{background-position:-97px -92px;margin:-3px 0 0 -10px;height:52px;width:112px}.gb_n .gb_Ra .gb_Ta{margin:-3px 0 0 -10px;height:52px;width:112px;background-position:-97px 0}@-webkit-keyframes gb__nb{0%{-webkit-transform:scale(0,0);transform:scale(0,0)}20%{-webkit-transform:scale(1.4,1.4);transform:scale(1.4,1.4)}50%{-webkit-transform:scale(.8,.8);transform:scale(.8,.8)}85%{-webkit-transform:scale(1.1,1.1);transform:scale(1.1,1.1)}to{-webkit-transform:scale(1.0,1.0);transform:scale(1.0,1.0)}}@keyframes gb__nb{0%{-webkit-transform:scale(0,0);transform:scale(0,0)}20%{-webkit-transform:scale(1.4,1.4);transform:scale(1.4,1.4)}50%{-webkit-transform:scale(.8,.8);transform:scale(.8,.8)}85%{-webkit-transform:scale(1.1,1.1);transform:scale(1.1,1.1)}to{-webkit-transform:scale(1.0,1.0);transform:scale(1.0,1.0)}}.gb_Xa .gb_Za{font-size:14px;font-weight:bold;top:0;right:0}.gb_Xa .gb_D{display:inline-block;vertical-align:middle;-webkit-box-sizing:border-box;box-sizing:border-box;height:30px;width:30px}.gb_0a{background-position:-56px 0;opacity:.55;height:100%;width:100%}.gb_D:hover .gb_0a,.gb_D:focus .gb_0a{opacity:.85}.gb_1a .gb_0a{background-position:-291px -103px}.gb_2a{background-color:#cb4437;-webkit-border-radius:8px;border-radius:8px;font:bold 11px Arial;color:#fff;line-height:16px;min-width:14px;padding:0 1px;position:absolute;right:0;text-align:center;text-shadow:0 1px 0 rgba(0,0,0,0.1);top:0;visibility:hidden;z-index:990}.gb_3a .gb_2a,.gb_3a .gb_4a,.gb_3a .gb_4a.gb_5a{visibility:visible}.gb_4a{padding:0 2px;visibility:hidden}.gb_Xa:not(.gb_6a) .gb_ka,.gb_Xa:not(.gb_6a) .gb_ja{left:3px}.gb_Xa .gb_ja{border-bottom-color:#e5e5e5}.gb_2a.gb_7a{-webkit-animation:gb__nb .6s 1s both ease-in-out;animation:gb__nb .6s 1s both ease-in-out;-webkit-perspective-origin:top right;perspective-origin:top right;-webkit-transform:scale(1,1);transform:scale(1,1);-webkit-transform-origin:top right;transform-origin:top right}.gb_7a .gb_4a{visibility:visible}.gb_8a{background-color:rgba(0,0,0,.55);color:#fff;font-size:12px;font-weight:bold;line-height:20px;margin:5px;padding:0 2px;text-align:center;-webkit-box-sizing:border-box;box-sizing:border-box;-webkit-border-radius:50%;border-radius:50%;height:20px;width:20px}.gb_8a.gb_9a{background-position:-214px -117px}.gb_8a.gb_ab{background-position:-256px -73px}.gb_D:hover .gb_8a,.gb_D:focus .gb_8a{background-color:rgba(0,0,0,.85)}#gbsfw.gb_bb{background:#e5e5e5;border-color:#ccc}.gb_n .gb_D .gb_0a{background-position:-167px -57px;opacity:.7}.gb_n .gb_1a .gb_0a{background-position:-132px -57px}.gb_n .gb_D:hover .gb_0a,.gb_n .gb_D:focus .gb_0a{opacity:.85}.gb_o .gb_D .gb_0a{background-position:-326px -87px;opacity:1}.gb_o .gb_1a .gb_0a{background-position:0 -70px}.gb_n .gb_2a,.gb_o .gb_2a{border:none;-webkit-box-shadow:-1px 1px 1px rgba(0,0,0,.2);box-shadow:-1px 1px 1px rgba(0,0,0,.2)}.gb_n .gb_8a{background-color:rgba(0,0,0,.7);-webkit-box-shadow:0 1px 2px rgba(255,255,255,.9);box-shadow:0 1px 2px rgba(255,255,255,.9)}.gb_o .gb_8a.gb_8a{background-color:#fff;color:#404040;-webkit-box-shadow:0 1px 2px rgba(0,0,0,.2);box-shadow:0 1px 2px rgba(0,0,0,.2)}.gb_o .gb_8a.gb_9a{background-position:-326px -122px}.gb_o .gb_8a.gb_ab{background-position:-214px -92px}.gb_3a .gb_8a.gb_8a{background-color:#db4437;color:#fff}.gb_3a .gb_D:hover .gb_8a,.gb_3a .gb_D:focus .gb_8a{background-color:#a52714}.gb_3a .gb_8a.gb_ab{background-position:-256px -73px}.gb_Rb .gb_D{background-position:-326px -17px;opacity:.55;height:30px;width:30px}.gb_Rb .gb_D:hover,.gb_Rb .gb_D:focus{opacity:.85}.gb_Rb .gb_ja{border-bottom-color:#f5f5f5}#gbsfw.gb_Sb{background:#f5f5f5;border-color:#ccc}.gb_o .gb_Rb .gb_D{background-position:0 -35px;opacity:1}.gb_n .gb_Rb .gb_D{background-position:-256px -103px;opacity:.7}.gb_n .gb_Rb .gb_D:hover,.gb_n .gb_Rb .gb_D:focus{opacity:.85}.gb_8{min-width:315px;padding-left:30px;padding-right:30px;position:relative;text-align:right;z-index:986;-webkit-align-items:center;align-items:center;-webkit-justify-content:flex-end;justify-content:flex-end}.gb_8.gb_i{-webkit-flex:1 1 auto;flex:1 1 auto;-webkit-flex:1 1 main-size;flex:1 1 main-size}.gb_Cc{line-height:normal;position:relative;text-align:left}.gb_Cc.gb_i,.gb_Dc.gb_i,.gb_ca.gb_i{-webkit-flex:0 1 auto;flex:0 1 auto;-webkit-flex:0 1 main-size;flex:0 1 main-size}.gb_Ec,.gb_Fc{display:inline-block;padding:0 0 0 15px;position:relative;vertical-align:middle}.gb_Dc{line-height:normal;padding-right:15px}.gb_8 .gb_Dc.gb_k{padding-right:0}.gb_ca{color:#404040;line-height:30px;min-width:30px;overflow:hidden;vertical-align:middle;text-overflow:ellipsis}#gb#gb.gb_ba .gb_9b,#gb#gb.gb_ba .gb_Cc>.gb_Fc .gb_ac{background:none;border:none;color:#36c;cursor:pointer;font-size:11px;line-height:26px;padding:0;-webkit-box-shadow:none;box-shadow:none}.gb_ba .gb_9b{text-transform:uppercase}.gb_8.gb_l{padding-left:0;padding-right:29px}.gb_8.gb_Hc{max-width:400px}.gb_Ic{background-clip:content-box;background-origin:content-box;opacity:.27;padding:22px;height:16px;width:16px}.gb_Ic.gb_i{display:none}.gb_Ic:hover,.gb_Ic:focus{opacity:.55}.gb_Jc{background-position:-35px 0}.gb_Kc{background-position:-291px -17px;padding-left:30px;padding-right:14px;position:absolute;right:0;top:0;z-index:990}.gb_fa:not(.gb_ha) .gb_Kc,.gb_l .gb_Jc{display:inline-block}.gb_fa .gb_Jc{padding-left:30px;padding-right:0;width:0}.gb_fa:not(.gb_ha) .gb_Lc{display:none}.gb_8.gb_i.gb_l,.gb_l:not(.gb_ha) .gb_Cc{-webkit-flex:0 0 auto;flex:0 0 auto;-webkit-flex:0 0 main-size;flex:0 0 main-size}.gb_Ic,.gb_l .gb_Dc,.gb_ha .gb_Cc{overflow:hidden}.gb_fa .gb_Dc{padding-right:0}.gb_l .gb_Cc{padding:1px 1px 1px 0}.gb_fa .gb_Cc{width:75px}.gb_8.gb_Mc,.gb_8.gb_Mc .gb_Jc,.gb_8.gb_Mc .gb_Jc::before,.gb_8.gb_Mc .gb_Dc,.gb_8.gb_Mc .gb_Cc{-webkit-transition:width .5s ease-in-out,min-width .5s ease-in-out,max-width .5s ease-in-out,padding .5s ease-in-out,left .5s ease-in-out;transition:width .5s ease-in-out,min-width .5s ease-in-out,max-width .5s ease-in-out,padding .5s ease-in-out,left .5s ease-in-out}.gb_Eb .gb_8{min-width:0}.gb_8.gb_m,.gb_8.gb_m .gb_Cc,.gb_8.gb_Nc,.gb_8.gb_Nc .gb_Cc{min-width:0!important}.gb_8.gb_m,.gb_8.gb_m .gb_i{-webkit-flex:0 0 auto!important;-webkit-flex:0 0 auto!important;flex:0 0 auto!important}.gb_8.gb_m .gb_ca{width:30px!important}.gb_n .gb_ca{font-weight:bold;text-shadow:0 1px 1px rgba(255,255,255,.9)}.gb_o .gb_ca{color:#fff;font-weight:bold;text-shadow:0 1px 1px rgba(0,0,0,.6)}.gb_Vb ::-webkit-scrollbar{height:15px;width:15px}.gb_Vb ::-webkit-scrollbar-button{height:0;width:0}.gb_Vb ::-webkit-scrollbar-thumb{background-clip:padding-box;background-color:rgba(0,0,0,.3);border:5px solid transparent;-webkit-border-radius:10px;border-radius:10px;min-height:20px;min-width:20px;height:5px;width:5px}.gb_Vb ::-webkit-scrollbar-thumb:hover,.gb_Vb ::-webkit-scrollbar-thumb:active{background-color:rgba(0,0,0,.4)}#gb.gb_Vc{min-width:980px}#gb.gb_Vc .gb_Db{min-width:0;position:static;width:0}.gb_Vc .gb_Vb{background:transparent;border-bottom-color:transparent}.gb_Vc .gb_Vb::before{display:none}.gb_Vc.gb_Vc .gb_h{display:inline-block}.gb_Vc.gb_8 .gb_Dc.gb_k{padding-right:15px}.gb_Vc .gb_j.gb_Sc{display:-webkit-flex;display:flex}.gb_Vc.gb_Eb #gbqf{display:block}.gb_Vc #gbq{height:0;position:absolute}.gb_Vc.gb_8{z-index:987}.gb_Pa.gb_Wc{padding-left:30px}.gb_Qa .gb_Wc{margin-right:-142px}.gbqfb,.gbqfba,.gbqfbb{cursor:default!important;display:inline-block;font-weight:bold;height:29px;line-height:29px;min-width:54px;padding:0 8px;text-align:center;text-decoration:none!important;-webkit-border-radius:2px;border-radius:2px;-webkit-user-select:none}.gbqfba:focus{border:1px solid #4d90fe;outline:none;-webkit-box-shadow:inset 0 0 0 1px #fff;box-shadow:inset 0 0 0 1px #fff}.gbqfba:hover{border-color:#c6c6c6;color:#222!important;-webkit-box-shadow:0 1px 0 rgba(0,0,0,.15);box-shadow:0 1px 0 rgba(0,0,0,.15);background:#f8f8f8;background:-webkit-linear-gradient(top,#f8f8f8,#f1f1f1);background:linear-gradient(top,#f8f8f8,#f1f1f1);filter:progid:DXImageTransform.Microsoft.gradient(startColorstr=#f8f8f8,endColorstr=#f1f1f1,GradientType=1)}.gbqfba:hover:focus{-webkit-box-shadow:0 1px 0 rgba(0,0,0,.15),inset 0 0 0 1px #fff;box-shadow:0 1px 0 rgba(0,0,0,.15),inset 0 0 0 1px #fff}.gbqfb::-moz-focus-inner{border:0}.gbqfba::-moz-focus-inner{border:0}.gbqfba{border:1px solid #dcdcdc;border-color:rgba(0,0,0,0.1);color:#444!important;font-size:11px;background:#f5f5f5;background:-webkit-linear-gradient(top,#f5f5f5,#f1f1f1);background:linear-gradient(top,#f5f5f5,#f1f1f1);filter:progid:DXImageTransform.Microsoft.gradient(startColorstr=#f5f5f5,endColorstr=#f1f1f1,GradientType=1)}.gbqfba:active{-webkit-box-shadow:inset 0 1px 2px rgba(0,0,0,0.1);box-shadow:inset 0 1px 2px rgba(0,0,0,0.1)}.gb_Db{position:relative;width:650px;z-index:986}#gbq2{padding-top:15px}.gb_Eb .gb_Db{min-width:200px;-webkit-flex:0 2 auto;flex:0 2 auto;-webkit-flex:0 2 main-size;flex:0 2 main-size}.gb_l~.gb_Db{min-width:0}.gb_Eb #gbqf{margin-right:0;display:-webkit-flex;display:flex}.gb_Eb .gbqff{min-width:0;-webkit-flex:1 1 auto;flex:1 1 auto;-webkit-flex:1 1 main-size;flex:1 1 main-size}#gbq2{display:block}#gbqf{display:block;margin:0;margin-right:60px;white-space:nowrap}.gbqff{border:none;display:inline-block;margin:0;padding:0;vertical-align:top;width:100%}.gbqfqw,#gbqfb,.gbqfwa{vertical-align:top}#gbqfaa,#gbqfab,#gbqfqwb{position:absolute}#gbqfaa{left:0}#gbqfab{right:0}.gbqfqwb,.gbqfqwc{right:0;left:0;height:100%}.gbqfqwb{padding:0 8px}#gbqfbw{display:inline-block;vertical-align:top}#gbqfb{border:1px solid transparent;border-bottom-left-radius:0;border-top-left-radius:0;height:30px;margin:0;outline:none;padding:0 0;width:60px;-webkit-box-shadow:none;box-shadow:none;-webkit-box-sizing:border-box;box-sizing:border-box;background:#4285f4;background:-webkit-linear-gradient(top,#4387fd,#4683ea);background:linear-gradient(top,#4387fd,#4683ea);filter:progid:DXImageTransform.Microsoft.gradient(startColorstr=#4387fd,endColorstr=#4683ea,GradientType=1)}#gbqfb:hover{-webkit-box-shadow:0 1px 0 rgba(0,0,0,.15);box-shadow:0 1px 0 rgba(0,0,0,.15)}#gbqfb:focus{-webkit-box-shadow:inset 0 0 0 1px #fff;box-shadow:inset 0 0 0 1px #fff}#gbqfb:hover:focus{-webkit-box-shadow:0 1px 0 rgba(0,0,0,.15),inset 0 0 0 1px #fff;box-shadow:0 1px 0 rgba(0,0,0,.15),inset 0 0 0 1px #fff}#gbqfb:active:active{border:1px solid transparent;-webkit-box-shadow:inset 0 2px 0 rgba(0,0,0,.15);box-shadow:inset 0 2px 0 rgba(0,0,0,.15);background:#3c78dc;background:-webkit-linear-gradient(top,#3c7ae4,#3f76d3);background:linear-gradient(top,#3c7ae4,#3f76d3);filter:progid:DXImageTransform.Microsoft.gradient(startColorstr=#3c7ae4,endColorstr=#3f76d3,GradientType=1)}.gbqfi{background-position:-56px -35px;display:inline-block;margin:-1px;height:30px;width:30px}.gbqfqw{background:#fff;background-clip:padding-box;border:1px solid #cdcdcd;border-color:rgba(0,0,0,.15);border-right-width:0;height:30px;-webkit-box-sizing:border-box;box-sizing:border-box}#gbfwc .gbqfqw{border-right-width:1px}#gbqfqw{position:relative}.gbqfqw.gbqfqw:hover{border-color:#a9a9a9;border-color:rgba(0,0,0,.3)}.gbqfwa{display:inline-block;width:100%}.gbqfwb{width:40%}.gbqfwc{width:60%}.gbqfwb .gbqfqw{margin-left:10px}.gbqfqw.gbqfqw:active,.gbqfqw.gbqfqwf.gbqfqwf{border-color:#4285f4}#gbqfq,#gbqfqb,#gbqfqc{background:transparent;border:none;height:20px;margin-top:4px;padding:0;vertical-align:top;width:100%}#gbqfq:focus,#gbqfqb:focus,#gbqfqc:focus{outline:none}.gbqfif,.gbqfsf{color:#222;font:16px arial,sans-serif}#gbqfbwa{display:none;text-align:center;height:0}#gbqfbwa .gbqfba{margin:16px 8px}#gbqfsa,#gbqfsb{font:bold 11px/27px Arial,sans-serif!important;vertical-align:top}.gb_n .gbqfqw.gbqfqw,.gb_o .gbqfqw.gbqfqw{border-color:rgba(255,255,255,1);-webkit-box-shadow:0 1px 2px rgba(0,0,0,.2);box-shadow:0 1px 2px rgba(0,0,0,.2)}.gb_n #gbqfb,.gb_o #gbqfb{-webkit-box-shadow:0 1px 2px rgba(0,0,0,.2);box-shadow:0 1px 2px rgba(0,0,0,.2)}.gb_n #gbqfb:hover,.gb_o #gbqfb:hover{-webkit-box-shadow:0 1px 0 rgba(0,0,0,.15),0 1px 2px rgba(0,0,0,.2);box-shadow:0 1px 0 rgba(0,0,0,.15),0 1px 2px rgba(0,0,0,.2)}.gb_n #gbqfb:active,.gb_o #gbqfb:active{-webkit-box-shadow:inset 0 2px 0 rgba(0,0,0,.15),0 1px 2px rgba(0,0,0,.2);box-shadow:inset 0 2px 0 rgba(0,0,0,.15),0 1px 2px rgba(0,0,0,.2)}sentinel{}.gbii{background-image:url(//ssl.gstatic.com/gb/images/silhouette_27.png)}.gbip{background-image:url()}.gbii::before{content:url(//ssl.gstatic.com/gb/images/silhouette_27.png);position:absolute}.gbip::before{content:url();position:absolute}@media (min-resolution:1.25dppx),(-o-min-device-pixel-ratio:5/4),(-webkit-min-device-pixel-ratio:1.25),(min-device-pixel-ratio:1.25){.gbii{background-image:url(//ssl.gstatic.com/gb/images/silhouette_27.png)}.gbii::before{content:url(//ssl.gstatic.com/gb/images/silhouette_27.png)}.gbip{background-image:url()}.gbip::before{content:url()}.gbii::before,.gbip::before{-webkit-transform:scale(.5);-moz-transform:scale(.5);-ms-transform:scale(.5);-o-transform:scale(.5);transform:scale(.5);-webkit-transform-origin:0 0;-moz-transform-origin:0 0;-ms-transform-origin:0 0;-o-transform-origin:0 0;transform-origin:0 0}}#gbq .gbgt-hvr,#gbq .gbgt:focus{background-color:transparent;background-image:none}.gbqfh#gbq1{display:none}.gbxx{display:none !important}#gbq{line-height:normal;position:relative;top:0px;white-space:nowrap}#gbq{left:0;width:100%}#gbq2{top:0px;z-index:986}#gbq4{display:inline-block;max-height:29px;overflow:hidden;position:relative}.gbqfh#gbq2{z-index:985}.gbqfh#gbq2{margin:0;margin-left:0 !important;padding-top:0;position:relative;top:310px}.gbqfh #gbqf{margin:auto;min-width:534px;padding:0 !important}.gbqfh #gbqfbw{display:none}.gbqfh #gbqfbwa{display:block}.gbqfh #gbqf{max-width:572px;min-width:572px}.gbqfh .gbqfqw{border-right-width:1px}.gsfe_a.gsfe_a{border-right-width:0;-moz-box-shadow:none;-webkit-box-shadow:none;box-shadow:none}.gsfe_b.gsfe_b{border-right-width:0;border-color:#4285f4;-moz-box-shadow:none;-webkit-box-shadow:none;box-shadow:none}.gbqfh .gsfe_a,.gbqfh .gsfe_b{border-width:1px}.gbm{background:#fff;border:1px solid #bebebe;box-shadow:0 2px 4px rgba(0,0,0,.2);-moz-box-shadow:-1px 1px 1px rgba(0,0,0,.2);-webkit-box-shadow:0 2px 4px rgba(0,0,0,.2);position:absolute;top:-999px;visibility:hidden;z-index:999}#sfcnt,#subform_ctrl{display:none}</style><style data-jiis="cc" id="gstyle">html,body{height:100%;margin:0}#viewport{min-height:100%;position:relative;width:100%}.content{padding-bottom:35px}#footer{bottom:0;font-size:10pt;height:35px;position:absolute;width:100%}#gog{padding:3px 8px 0}.gac_m td{line-height:17px}body,td,a,p,.h{font-family:arial,sans-serif}.h{color:#12c;font-size:20px}.q{color:#00c}.ts td{padding:0}.ts{border-collapse:collapse}em{font-weight:bold;font-style:normal}.ds{display:inline-block;}span.ds{margin:3px 0 4px;margin-left:4px}.ctr-p{margin:0 auto;min-width:980px}a.gb1,a.gb2,a.gb3,a.gb4{color:#11c !important}body{background:#fff;color:#222}a{color:#12c;text-decoration:none}a:hover,a:active{text-decoration:underline}.fl a{color:#12c}a:visited{color:#609}a.gb1,a.gb4{text-decoration:underline}a.gb3:hover{text-decoration:none}#ghead a.gb2:hover{color:#fff !important}.sblc{padding-top:5px}.sblc a{display:block;margin:2px 0;margin-left:13px;font-size:11px}.lsbb{height:30px;display:block}.ftl,#footer a{color:#666;margin:2px 10px 0}#footer a:active{color:#dd4b39}.lsb{border:none;color:#000;cursor:pointer;height:30px;margin:0;outline:0;vertical-align:top}.lst:focus{outline:none}body,html{font-size:small}h1,ol,ul,li{margin:0;padding:0}.nojsv{visibility:hidden}.hp #logocont.nojsv{display:none}#body,#footer{display:block}.igehp{display:none}#flci{float:left;margin-left:0;text-align:left;width:0}#fll{float:right;text-align:right;width:100%}#ftby{padding-left:0}#ftby>div,#fll>div,#footer a{display:inline-block}@media only screen and (min-width:1222px){#ftby{margin:0 44px}}.nojsb{display:none}</style><style>.kpbb,.kprb,.kpgb,.kpgrb{-webkit-border-radius:2px;border-radius:2px;color:#fff}.kpbb:hover,.kprb:hover,.kpgb:hover,.kpgrb:hover{-webkit-box-shadow:0 1px 1px rgba(0,0,0,0.1);box-shadow:0 1px 1px rgba(0,0,0,0.1);color:#fff}.kpbb:active,.kprb:active,.kpgb:active,.kpgrb:active{-webkit-box-shadow:inset 0 1px 2px rgba(0,0,0,0.3);box-shadow:inset 0 1px 2px rgba(0,0,0,0.3)}.kpbb{background-image:-webkit-gradient(linear,left top,left bottom,from(#4d90fe),to(#4787ed));background-color:#4d90fe;background-image:-webkit-linear-gradient(top,#4d90fe,#4787ed);background-image:linear-gradient(top,#4d90fe,#4787ed);border:1px solid #3079ed}.kpbb:hover{background-image:-webkit-gradient(linear,left top,left bottom,from(#4d90fe),to(#357ae8));background-color:#357ae8;background-image:-webkit-linear-gradient(top,#4d90fe,#357ae8);background-image:linear-gradient(top,#4d90fe,#357ae8);border:1px solid #2f5bb7}a.kpbb:link,a.kpbb:visited{color:#fff}.kprb{background-image:-webkit-gradient(linear,left top,left bottom,from(#dd4b39),to(#d14836));background-color:#dd4b39;background-image:-webkit-linear-gradient(top,#dd4b39,#d14836);background-image:linear-gradient(top,#dd4b39,#d14836);border:1px solid #dd4b39}.kprb:hover{background-image:-webkit-gradient(linear,left top,left bottom,from(#dd4b39),to(#c53727));background-color:#c53727;background-image:-webkit-linear-gradient(top,#dd4b39,#c53727);background-image:linear-gradient(top,#dd4b39,#c53727);border:1px solid #b0281a;border-bottom-color:#af301f}.kprb:active{background-image:-webkit-gradient(linear,left top,left bottom,from(#dd4b39),to(#b0281a));background-color:#b0281a;background-image:-webkit-linear-gradient(top,#dd4b39,#b0281a);background-image:linear-gradient(top,#dd4b39,#b0281a)}.kpgb{background-image:-webkit-gradient(linear,left top,left bottom,from(#3d9400),to(#398a00));background-color:#3d9400;background-image:-webkit-linear-gradient(top,#3d9400,#398a00);background-image:linear-gradient(top,#3d9400,#398a00);border:1px solid #29691d}.kpgb:hover{background-image:-webkit-gradient(linear,left top,left bottom,from(#3d9400),to(#368200));background-color:#368200;background-image:-webkit-linear-gradient(top,#3d9400,#368200);background-image:linear-gradient(top,#3d9400,#368200);border:1px solid #2d6200}.kpgrb{background-image:-webkit-gradient(linear,left top,left bottom,from(#f5f5f5),to(#f1f1f1));background-color:#f5f5f5;background-image:-webkit-linear-gradient(top,#f5f5f5,#f1f1f1);background-image:linear-gradient(top,#f5f5f5,#f1f1f1);border:1px solid #dcdcdc;color:#555}.kpgrb:hover{background-image:-webkit-gradient(linear,left top,left bottom,from(#f8f8f8),to(#f1f1f1));background-color:#f8f8f8;background-image:-webkit-linear-gradient(top,#f8f8f8,#f1f1f1);background-image:linear-gradient(top,#f8f8f8,#f1f1f1);border:1px solid #dcdcdc;color:#333}a.kpgrb:link,a.kpgrb:visited{color:#555}.lst-t{width:100%}input#gbqfq{padding:0 0 0 9px}#pocs{background:#fff1a8;color:#000;font-size:10pt;margin:0;padding:5px 7px 0}#pocs.sft{background:transparent;color:#777}#pocs a{color:#11c}#pocs.sft a{color:#36c}#pocs>div{margin:0;padding:0}.gl{white-space:nowrap}.big .tsf-p{padding-left:126px;padding-right:352px}.tsf-p{padding-left:126px;padding-right:46px}#fkbx-tchm{}.fkbx-chm{}.fkbx-chme{}#fkbx-chmer{}#fkbx-chmed{}.fkbx-chmt{}#fkbx-chmtr{}.chw-oc{}#chw-o{}.jfk-bubble-arrowimplafter{}#fkbx-tchm{display:none}.fkbx-chm{line-height:22px;text-align:center}.fkbx-chm a{color:#2518b5;cursor:pointer;margin:5px}._gSc{background:url(data:image/gif;base64,R0lGODlhEAAQAKIHAPzu7PfT0Oh5cfGtqONbUuBLQeBKP////yH5BAEAAAcALAAAAAAQABAAAANKeLrcfkAI8NowZtQFCCbUJmCYsAWFAQBGEVSjyhqmc2HBnDUdGQQkEOOGA5I0CkCKxMQUQjEnAMU0GUkuZTPgaRaWTEK0Sa5tGgkAOw==) no-repeat center;display:inline-block;height:16px;width:16px}#chw-o{display:none}#chw-o a{color:#4285F4;line-height:31px}.chw-oc{font-size:13px;padding:0px !important;text-align:left;width:400px}.chw-oc .jfk-bubble-arrowimplafter{border-color:#f5f5f5 transparent !important}._mSc{color:#000;font-size:16px;font-weight:bold}._kSc{color:#555}._dKb{border-radius:2px;cursor:pointer;font-size:12px;line-height:27px;margin:0;padding-left:14px;padding-right:14px}#chw-o ._dKb{float:right;margin-left:10px}._k3{background-color:#f9f9f9;border:1px solid #bdbdbd;color:#000}._k3:hover{background-color:#fcfcfc}._k3:active,._k3:hover,._k3:focus{border-color:#3e7ef8}._k3:active{background-color:#e6e6e6}._WW{background-color:#5a97ff;border:1px solid #2558b0;color:#fff}._WW:hover{background-color:#629cff}._WW:hover,._WW:focus{box-shadow:inset 0 0 1px}._WW:active,._qyd:focus,._WW:hover{border-color:#2352a2}._WW:active{background-color:#4279d8}._jSc{background-color:#f5f5f5;border-top:1px solid #e7e7e7;padding-bottom:14px;padding-left:20px;padding-top:14px}._fdc{color:#888;display:block;margin-left:20px}._fdc:hover{color:#000}._iSc{position:relative;top:-3px}._lSc{padding:20px}._hSc{display:inline-block;margin-left:-20px;margin-right:7px}</style><script>var _gjwl=location;function _gjuc(){var a=_gjwl.href.indexOf("#");return 0<=a&&(a=_gjwl.href.substring(a+1),/(^|&)q=/.test(a)&&-1==a.indexOf("#")&&!/(^|&)cad=h($|&)/.test(a))?(_gjwl.replace("/search?"+a.replace(/(^|&)fp=[^&]*/g,"")+"&cad=h"),1):0}function _gjh(){!_gjuc()&&google.x({id:"GJH"},function(){google.nav&&google.nav.gjh&&google.nav.gjh()})};window.rwt=function(a,g,h,n,o,i,c,p,j,d){return true};
(window['gbar']=window['gbar']||{})._CONFIG=[[[0,"www.gstatic.com","og.og2.en_US.Dtz1Q9QKJyk.O","com","en","1",0,[3,2,".64.40.36.36.40.36.36.","r_qf.","17259,3700315","1422848042","0"],"40400","zKjSVI7_L4HfoASqnIKIDg",0,0,"og.og2.1497lsiin6dwu.L.W.O","AItRSTMaVzgSD3JgKdV83-qMqjqnH9cN3w","AItRSTMR2xxSfm8n9CU2QuETquAnmYCFaQ","",2,0,200,"USA"],null,0,["m;/_/scs/abc-static/_/js/k=gapi.gapi.en.V4LcX4GIySc.O/m=__features__/am=AAI/rt=j/d=1/rs=AItRSTMeAjKAPwkFy_1E-NAjuYKuxXbXEw","https://apis.google.com","","","","","",1,"es_plusone_gc_20150129.0_p0","en"],["1","gci_91f30755d6a6b787dcc2a4062e6e9824.js","googleapis.client:plusone:gapi.iframes","","en"],null,null,null,[0.009999999776482582,"com","1",[null,"","w",null,1,5184000,1,0],null,[["","","",0,0,-1]],[null,0,0],0,null,null,["5061451","google\\.(com|ru|ca|by|kz|com\\.mx)$",1]],null,[0,0,0,null,"","",""],[1,0.001000000047497451,1],[1,0.1000000014901161,2,1],[0,"",null,"",0,"There was an error loading your Marketplace apps.","You have no Marketplace apps.",0,[1,"https://www.google.com/webhp?tab=ww","Search","","0 -414px",null,0],0,0,1],[1],[0,1,["lg"],1,["lat"]],[["","","","","","","","","","","","","","","","","","","","def","","","",""],[""]],null,null,null,[30,127,1,980],null,null,null,null,null,[1,0]]];(window['gbar']=window['gbar']||{})._DPG=[{'_fdm_':['_r'],'dbg':['sy10','sy4','sy5','sy8','sy9'],'def':['sy11','sy12','sy13','sy4','sy5','sy6','sy8','sy9'],'drt':['sy4','sy5','sy6','sy7'],'fg':['sy16'],'fot':['sy11','sy12','sy5','sy9'],'ig':['sy16'],'in':['_r'],'jb':['sy4','sy5','sy6','sy7'],'lat':['sy10','sy11','sy13','sy4','sy5','sy6','sy8','sy9'],'lg':['sy16'],'sg':['sy16'],'sy10':['sy4','sy5','sy8'],'sy11':['_r'],'sy12':['sy11','sy5','sy9'],'sy13':['sy11','sy4','sy5','sy6','sy8','sy9'],'sy16':['_r'],'sy4':['sy5'],'sy5':['_r'],'sy6':['sy4'],'sy7':['sy4','sy6'],'sy8':['sy4','sy5'],'sy9':['_r']}];(window['gbar']=window['gbar']||{})._LDD=["in","fot"];this.gbar_=this.gbar_||{};(function(_){var window=this;
try{
var da,ea;_.aa=_.aa||{};_.m=this;_.n=function(a){return void 0!==a};_.p=function(a,c){for(var d=a.split("."),e=c||_.m,f;f=d.shift();)if(null!=e[f])e=e[f];else return null;return e};_.ba=function(a){a.N=function(){return a.Qe?a.Qe:a.Qe=new a}};_.t=function(a){return"string"==typeof a};_.ca="closure_uid_"+(1E9*Math.random()>>>0);da=function(a,c,d){return a.call.apply(a.bind,arguments)};
ea=function(a,c,d){if(!a)throw Error();if(2<arguments.length){var e=Array.prototype.slice.call(arguments,2);return function(){var d=Array.prototype.slice.call(arguments);Array.prototype.unshift.apply(d,e);return a.apply(c,d)}}return function(){return a.apply(c,arguments)}};_.u=function(a,c,d){_.u=Function.prototype.bind&&-1!=Function.prototype.bind.toString().indexOf("native code")?da:ea;return _.u.apply(null,arguments)};_.fa=Date.now||function(){return+new Date};
_.v=function(a,c){var d=a.split("."),e=_.m;d[0]in e||!e.execScript||e.execScript("var "+d[0]);for(var f;d.length&&(f=d.shift());)!d.length&&_.n(c)?e[f]=c:e[f]?e=e[f]:e=e[f]={}};_.w=function(a,c){function d(){}d.prototype=c.prototype;a.G=c.prototype;a.prototype=new d;a.prototype.constructor=a;a.qi=function(a,d,g){for(var h=Array(arguments.length-2),l=2;l<arguments.length;l++)h[l-2]=arguments[l];return c.prototype[d].apply(a,h)}};
_.y=function(){this.ja=this.ja;this.Na=this.Na};_.y.prototype.ja=!1;_.y.prototype.Y=function(){this.ja||(this.ja=!0,this.L())};_.y.prototype.L=function(){if(this.Na)for(;this.Na.length;)this.Na.shift()()};_.ga=function(a){if(Error.captureStackTrace)Error.captureStackTrace(this,_.ga);else{var c=Error().stack;c&&(this.stack=c)}a&&(this.message=String(a))};_.w(_.ga,Error);_.ga.prototype.name="CustomError";_.ha=String.prototype.trim?function(a){return a.trim()}:function(a){return a.replace(/^[\s\xa0]+|[\s\xa0]+$/g,"")};_.ia=Array.prototype;_.ja=_.ia.indexOf?function(a,c,d){return _.ia.indexOf.call(a,c,d)}:function(a,c,d){d=null==d?0:0>d?Math.max(0,a.length+d):d;if(_.t(a))return _.t(c)&&1==c.length?a.indexOf(c,d):-1;for(;d<a.length;d++)if(d in a&&a[d]===c)return d;return-1};_.ka=_.ia.forEach?function(a,c,d){_.ia.forEach.call(a,c,d)}:function(a,c,d){for(var e=a.length,f=_.t(a)?a.split(""):a,g=0;g<e;g++)g in f&&c.call(d,f[g],g,a)};
_.la=_.ia.filter?function(a,c,d){return _.ia.filter.call(a,c,d)}:function(a,c,d){for(var e=a.length,f=[],g=0,h=_.t(a)?a.split(""):a,l=0;l<e;l++)if(l in h){var q=h[l];c.call(d,q,l,a)&&(f[g++]=q)}return f};_.ma=_.ia.map?function(a,c,d){return _.ia.map.call(a,c,d)}:function(a,c,d){for(var e=a.length,f=Array(e),g=_.t(a)?a.split(""):a,h=0;h<e;h++)h in g&&(f[h]=c.call(d,g[h],h,a));return f};
_.na=_.ia.reduce?function(a,c,d,e){e&&(c=(0,_.u)(c,e));return _.ia.reduce.call(a,c,d)}:function(a,c,d,e){var f=d;(0,_.ka)(a,function(d,h){f=c.call(e,f,d,h,a)});return f};_.oa=_.ia.some?function(a,c,d){return _.ia.some.call(a,c,d)}:function(a,c,d){for(var e=a.length,f=_.t(a)?a.split(""):a,g=0;g<e;g++)if(g in f&&c.call(d,f[g],g,a))return!0;return!1};_.pa=function(a,c){return 0<=(0,_.ja)(a,c)};
_.qa=/\uffff/.test("\uffff")?/[\\\"\x00-\x1f\x7f-\uffff]/g:/[\\\"\x00-\x1f\x7f-\xff]/g;var ra;ra="constructor hasOwnProperty isPrototypeOf propertyIsEnumerable toLocaleString toString valueOf".split(" ");_.sa=function(a,c){for(var d,e,f=1;f<arguments.length;f++){e=arguments[f];for(d in e)a[d]=e[d];for(var g=0;g<ra.length;g++)d=ra[g],Object.prototype.hasOwnProperty.call(e,d)&&(a[d]=e[d])}};
_.z=function(){};_.A=function(a,c,d,e){a.d={};c||(c=d?[d]:[]);a.A=d?String(d):void 0;a.k=0===d?-1:0;a.b=c;t:{if(a.b.length&&(c=a.b.length-1,(d=a.b[c])&&"object"==typeof d&&"number"!=typeof d.length)){a.w=c-a.k;a.o=d;break t}a.w=Number.MAX_VALUE}if(e)for(c=0;c<e.length;c++)d=e[c],d<a.w?(d+=a.k,a.b[d]=a.b[d]||[]):a.o[d]=a.o[d]||[]};_.B=function(a,c){return c<a.w?a.b[c+a.k]:a.o[c]};_.D=function(a,c,d){if(!a.d[d]){var e=_.B(a,d);e&&(a.d[d]=new c(e))}return a.d[d]};_.z.prototype.$a=function(){return this.b}; _.z.prototype.toString=function(){return this.b.toString()};
_.ta=function(a){_.A(this,a,0,null)};_.w(_.ta,_.z);var ua=function(a){_.A(this,a,0,null)};_.w(ua,_.z);var Ba;_.va=function(){this.b={};this.d={}};_.ba(_.va);_.xa=function(a,c){a.N=function(){return _.wa(_.va.N(),c)};a.Wh=function(){return _.va.N().b[c]||null}};_.E=function(a){return _.wa(_.va.N(),a)};_.za=function(a,c){var d=_.va.N();if(a in d.b){if(d.b[a]!=c)throw new ya(a);}else{d.b[a]=c;var e=d.d[a];if(e)for(var f=0,g=e.length;f<g;f++)e[f].b(d.b,a);delete d.d[a]}};_.wa=function(a,c){if(c in a.b)return a.b[c];throw new Aa(c);};Ba=function(a){_.ga.call(this);this.aa=a};_.w(Ba,_.ga); var ya=function(a){Ba.call(this,a)};_.w(ya,Ba);var Aa=function(a){Ba.call(this,a)};_.w(Aa,Ba);
_.Ca=function(a){_.A(this,a,0,null)};_.w(_.Ca,_.z);var Da=function(a){_.A(this,a,0,null)};_.w(Da,_.z);Da.prototype.yc=function(){return _.D(this,_.Ca,14)};_.F=function(a,c){return null!=a?a:!!c};_.G=function(a){var c;void 0==c&&(c="");return null!=a?a:c};_.H=function(a,c){void 0==c&&(c=0);return null!=a?a:c};var Ea;Ea=new Da(window.gbar&&window.gbar._CONFIG?window.gbar._CONFIG[0]:[[,,,,,,,[]],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[]]);_.Fa=_.F(_.B(Ea,3));_.I=function(){};_.v("gbar_._DumpException",function(a){if(_.Fa)throw a;_.I(a)});var Ga=function(){this.d=!1;this.b=[]};Ga.prototype.k=function(a){if(!this.d){this.d=!0;for(var c=0;c<this.b.length;c++)try{this.b[c]()}catch(d){a(d)}this.b=null;try{_.E("api").Ua()}catch(e){}}};_.Ha=new Ga;var Ia=function(){_.y.call(this);this.rh=Ea};_.w(Ia,_.y);_.xa(Ia,"cs");_.J=function(){return Ia.N().rh};_.Ja=function(){return _.D(_.J(),ua,1)||new ua};_.za("cs",new Ia);var Ka=function(a,c,d){this.o=a;this.d=!1;this.b=c;this.k=d};Ka.prototype.Ua=function(a){if(this.d)throw Error("e` + "`" + `"+this.b);try{a.apply(this.o,this.k),this.d=!0}catch(c){}};var La=function(a){_.y.call(this);this.k=a;this.b=[];this.d={}};_.w(La,_.y);La.prototype.o=function(a){var c=(0,_.u)(function(){this.b.push(new Ka(this.k,a,Array.prototype.slice.call(arguments)))},this);return this.d[a]=c};
La.prototype.Ua=function(){for(var a=this.b.length,c=this.b,d=[],e=0;e<a;++e){var f=c[e].b,g;t:{g=this.k;for(var h=f.split("."),l=h.length,q=0;q<l;++q)if(g[h[q]])g=g[h[q]];else{g=null;break t}g=g instanceof Function?g:null}if(g&&g!=this.d[f])try{c[e].Ua(g)}catch(r){}else d.push(c[e])}this.b=d.concat(c.slice(a))};
var Na;_.Ma="bbh bbr bbs has prm sngw so".split(" ");Na=new La(_.m);_.za("api",Na);
for(var Oa="addExtraLink addLink aomc asmc close cp.c cp.l cp.me cp.ml cp.rc cp.rel ela elc elh gpca gpcr lGC lPWF ldb mls noam paa pc pca pcm pw.clk pw.hvr qfaae qfaas qfaau qfae qfas qfau qfhi qm qs qsi rtl sa setContinueCb snaw sncw som sp spd spn spp sps tsl tst up.aeh up.aop up.dpc up.iic up.nap up.r up.sl up.spd up.tp upel upes upet".split(" ").concat(_.Ma),Pa=(0,_.u)(Na.o,Na),Qa=0;Qa<Oa.length;Qa++){var Ra="gbar."+Oa[Qa];null==_.p(Ra,window)&&_.v(Ra,Pa(Ra))}_.v("gbar.up.gpd",function(){return""});
var Sa=_.Ja(),Ta=_.D(Sa,_.ta,8)||new _.ta;_.v("gbar.bv",{n:_.H(_.B(Ta,2)),r:_.G(_.B(Ta,4)),f:_.G(_.B(Ta,3)),e:_.G(_.B(Ta,5)),m:_.H(null!=_.B(Ta,1)?_.B(Ta,1):1,1)});_.v("gbar.kn",function(){return!0});_.v("gbar.sb",function(){return!1});
}catch(e){_._DumpException(e)}
try{
_.v("gbar.elr",function(){return{es:{f:152,h:60,m:30},mo:"md",vh:window.innerHeight||0,vw:window.innerWidth||0}});
}catch(e){_._DumpException(e)}
})(this.gbar_);
// Google Inc.
</script><link href="/" rel="canonical"></head><body class="hp vasq" onload="try{if(!google.j.b){document.f&amp;&amp;document.f.q.focus();document.gbqf&amp;&amp;document.gbqf.q.focus();}}catch(e){}if(document.images)new Image().src='/images/nav_logo195.png'" alink="#dd4b39" bgcolor="#fff" id="gsr" link="#12c" text="#222" vlink="#61c"><div class="ctr-p" id="viewport"><div id="pocs" style="display:none;left:0px;white-space:nowrap;position:absolute"><div id="pocs0"><span><span>Google</span> Instant is unavailable. Press Enter to search.</span>&nbsp;<a href="/support/websearch/bin/answer.py?answer=186645&amp;form=bb&amp;hl=en">Learn more</a></div><div id="pocs1"><span>Google</span> Instant is off due to connection speed. Press Enter to search.</div><div id="pocs2">Press Enter to search.</div></div><div data-jiis="cc" id="cst"><style>.fade #center_col,.fade #rhs,.fade #leftnav,.fade #brs,.fade #footcnt{filter:alpha(opacity=33.3);opacity:0.333}.fade-hidden #center_col,.fade-hidden #rhs,.fade-hidden #leftnav{visibility:hidden}.flyr-o,.flyr-w{position:absolute;background-color:#fff;z-index:3;display:block;}.flyr-o{filter:alpha(opacity=66.6);opacity:0.666}.flyr-w{filter:alpha(opacity=20.0);opacity:0.2}.flyr-h{filter:alpha(opacity=0);opacity:0}.flyr-c{display:none}.flt,.flt u,a.fl{text-decoration:none}.flt:hover,.flt:hover u,a.fl:hover{text-decoration:underline}#knavm{color:#4273db;display:inline;font:11px arial,sans-serif !important;left:-13px;position:absolute;top:2px;z-index:2}#pnprev #knavm{bottom:1px;top:auto}#pnnext #knavm{bottom:1px;left:40px;top:auto}a.noline{outline:0}.y.yp{display:none}.y.yf,.y.ys{display:block}.yi{}.s2er{}.s2fp{}.s2fp-h{}.s2ml{}.s2ra{}.s2tb{}.s2tb-h{}.spch{}.spchc{}.spch{background:#fff;height:100%;left:0;opacity:0;overflow:hidden;position:fixed;text-align:left;top:0;visibility:hidden;width:100%;z-index:10000;transition:visibility 0s linear 0.218s,opacity 0.218s,background-color 0.218s}.s2fp.spch{opacity:1;visibility:visible;transition-delay:0s}.s2tb-h.spch{background:rgba(255,255,255,0);opacity:0;visibility:hidden}.s2tb.spch{background:rgba(255,255,255,0);opacity:1;visibility:visible;transition-delay:0s}.close-button{color:#777;cursor:pointer;font-size:26px;right:0;height:11px;line-height:15px;margin:15px;opacity:.6;padding:0;position:absolute;top:0;width:15px}.close-button:hover{opacity:.8}.close-button:active{opacity:1}.google-logo{background:url(data:image/gif;base64,R0lGODlhxABEAMIGALGzsMPFwtHT0N/h3u3v7P3//P///////yH5BAEKAAcALAAAAADEAEQAAAP+WLrc/jDKSau9OJcBuveCJo5kaZ7oJHwsMKSwQsx0bN84M+zDygY5DaEXaH0CggEhyGwKWyGnhCAoGq+ewEvK7cpaW++C+lMuZUOrUXAWu20E8Lvg62grAzUrOu+X4ixhXYAebRYcV4J+ixaEHopOdQB8GpIgjJiNcl16lCKIj5miE44dkEF6pxqgqqOupS5ckp4kgK2uo7C3MKAAQClFhrjDX4FOpcK1ycTDuk56v8zScJtBvQDL09ojzkx6ALlEVgFIuxdU5EXkSdmD1RRDVepI7aW0c3lYWe2k30Z33Ijoq9ZtQj599yyZc9KL3YyDR/g5IKBGCw2IHxbq8Df+cFKDghAoguBhqYOwFotEdrhXQGVGC3UCLHN5CU8WMxgf8XgA0kGdhC3awPLjSKI/lgzU7CoZLYKaZNd81XvX4GmEkmcsNRVTtIK/W2qQMmAqwSrPFlsZ9FygdMK3X2/nONIICxuEiheO3g0VIWraYi8dgPr7UU7cN3owRAXn84NEtUagfqDw9ayxwoUqoC2w2U3DDHohT8ZAtsEReJ0xB046msLm1F4uX1hc9YNYy7LpnKYcdCJVR4TVdjpjxDOLxw6MhCmlkfUP5ytt5gbMd8EsCRABLigu5puI0rpXz+be8rnmPb6nd+I5a1lkMShXkf/WnIFy6r4sfKPVE73+8JtlQeHFUE/0xtl0+gm4QXwUKJSeeKVEAxEbDcLGkIGgybFWBSVFAQtyDqpWXXkRxdQcbV1gJUJlsNwGgV/4eUTBNbiN+OEMItRVH0889OijEuGJd0FJL8AS3BSpkSeBJH+B9GEJHNVS0hoKdPiJhkqeo2CQNQWY2YM2UpXBYju2xBSVCsAo33EksjDCO+9JYFtIVFmowRVC4NiAijEyNiaD993JpnV2jtUlmKbUNugIU7oIE4YHIjijeQX4c2QEDLJV6ILRRdCTlbVcgVyFi1a6KXayoYiBPQ8wJQgZiSI5XV2jQjClXSaUJOKXQ7ppn5gGlarDFeQIW2Osv1L+OgJHl3KYqQKh5SUkeI9mYRBHQtKJ4JRlinZqta2NEWcFwD0QaIK8wtOmB5pImsgfxNbamK97fltvuuJCiuqIhzzrKbCLxUJCwAJX4i+0W8qK7Ly7uXZor8YeW/C9CU8h8T4G0+tAZQpP3Kq9pna6ZkSoSYrwu55SyBO2UpF2sKb+1YicX9m45Ki2+iKan1P6WHSRFRIR7IvKIeXUQbBgGJJdrTStZAisLUN5hVjMJgEkxR3ZMeqtPwjgtdeWsgPuDz/Iu4DR/3RrLrEvdnRK01iUiXZH5QxcRdonoPMP0TnuIE/PPNEguOCk3I33H+L8gIRDN+iZQw2Is4yFTI+RE77N5VO8xXiVbGPuOVfQ3FKXjJ+XzgRz5EZp+uqNR3wVsKzHXiC+zjYs++0aQLOs7bj3XjK/Wlrr+/AquH6eHcQnv5djJISl/PPQLZyhSdBXf5gI4VaffFRmA6a29pdniW6z4Mc+pbw+kF9+7Bw5qpL667NP7NVqTRh/+ULbYWn397PO9Rr865/s4hE2+rkiAQA7) no-repeat center;background-size:98px 34px;float:right;height:34px;left:255px;opacity:0;pointer-events:none;position:relative;top:6px;width:98px;transition:opacity .5s ease-in,left .5s ease-in}.s2tb .google-logo{opacity:1;left:270px;transition:opacity .5s ease-out,left .5s ease-out}.spchc{display:block;height:27px;position:absolute;pointer-events:none}.s2fp .spchc,.s2fp-h .spchc{margin:auto;margin-top:312px;max-width:572px;min-width:534px;padding:0 223px;position:relative;top:0}.s2tb .spchc,.s2tb-h .spchc{background:#fff;box-shadow:0 2px 6px rgba(0,0,0,0.2);margin:0;min-width:100%;overflow:hidden;padding:51px 0 65px 126px;position:absolute}._o3{height:100%;opacity:.1;pointer-events:none;width:100%;transition:opacity .318s ease-in}.s2tb-h ._o3,.s2tb ._o3{height:100%;width:572px;transition:opacity .318s ease-in}.s2ml ._o3,.s2ra ._o3,.s2er ._o3{opacity:1;transition:opacity 0s}.button{background-color:#fff;border:1px solid #eee;border-radius:100%;bottom:0;box-shadow:0 2px 5px rgba(0,0,0,.1);cursor:pointer;display:inline-block;left:0;opacity:0;pointer-events:none;position:absolute;right:0;top:0;transition:background-color 0.218s,border 0.218s,box-shadow 0.218s}.s2tb-h .button{left:-83px;opacity:0;pointer-events:none;position:absolute;top:-83px;transition-delay:0}.s2fp-h .button{opacity:0;pointer-events:none;position:absolute;transition-delay:0}.s2fp .button,.s2tb .button{opacity:1;pointer-events:auto;position:absolute;-webkit-transform:scale(1);transition-delay:0}.s2ra .button{background-color:#ff4444;border:0;box-shadow:none}._CMb{background-color:#dbdbdb;border-radius:100%;display:inline-block;height:301px;left:-69px;opacity:1;pointer-events:none;position:absolute;top:-69px;width:301px;-webkit-transform:scale(.01);transition:opacity 0.218s}.s2tb-h ._CMb,.s2tb ._CMb{height:151px;left:-28px;top:-28px;width:151px}._AM{float:right;pointer-events:none;position:relative;transition:-webkit-transform 0.218s,opacity 0.218s ease-in}.s2fp-h ._AM,.s2fp ._AM{height:165px;right:-70px;top:-70px;width:165px}.s2fp-h ._AM,.s2tb-h ._AM{-webkit-transform:scale(.1)}.s2fp ._AM,.s2tb ._AM{-webkit-transform:scale(1)}.s2tb-h ._AM,.s2tb ._AM{height:95px;right:-31px;top:-27px;width:95px}.s2ra .button:active{background-color:#cd0000}.button:active{background-color:#eee}._wPb{height:87px;left:43px;pointer-events:none;position:absolute;top:47px;width:42px;-webkit-transform:scale(1)}.s2tb-h ._wPb,.s2tb ._wPb{left:17px;top:7px;-webkit-transform:scale(.53)}._AUb{background-color:#999;border-radius:30px;height:46px;left:25px;pointer-events:none;position:absolute;width:24px}._Fjd{bottom:0;height:53px;left:11px;overflow:hidden;pointer-events:none;position:absolute;width:52px}._oXb{background-color:#999;bottom:14px;height:14px;left:22px;pointer-events:none;position:absolute;width:9px;z-index:1}._dWb{border:7px solid #999;border-radius:28px;bottom:27px;height:57px;pointer-events:none;position:absolute;width:38px;z-index:0}.s2ml ._AUb,.s2ml ._oXb{background-color:#f44}.s2ml ._dWb{border-color:#f44}.s2ra ._AUb,.s2ra ._oXb{background-color:#fff}.s2ra ._dWb{border-color:#fff}.spcht{}.spchta{}.spch-2l{}.spch-3l{}.spch-4l{}.spch-5l{}._gjb{pointer-events:none}.s2fp-h ._gjb,.s2fp ._gjb{position:absolute}.s2tb-h ._gjb,.s2tb ._gjb{position:relative}.spcht{font-weight:normal;line-height:1.2;opacity:0;pointer-events:none;position:absolute;text-align:left;-webkit-font-smoothing:antialiased;transition:opacity .1s ease-in,margin-left .5s ease-in,top 0s linear 0.218s}.s2fp-h .spcht{margin-left:44px}.s2tb-h .spcht{margin-left:32px}.s2fp-h .spcht,.s2fp .spcht{font-size:32px;left:-44px;top:-.2em;width:460px}.s2tb-h .spcht,.s2tb .spcht{font-size:27px;left:7px;top:.2em;width:490px}.s2fp .spcht,.s2tb .spcht{margin-left:0;opacity:1;transition:opacity .5s ease-out,margin-left .5s ease-out}.spchta{color:#1155cc;cursor:pointer;font-size:18px;font-weight:500;pointer-events:auto;text-decoration:underline}.spch-2l.spcht,.spch-3l.spcht,.spch-4l.spcht{transition:top 0.218s ease-out}.spch-2l.spcht{top:-.6em}.spch-3l.spcht{top:-1.3em}.spch-4l.spcht{top:-1.7em}.s2fp .spch-5l.spcht{top:-2.5em}.s2tb .spch-5l.spcht{font-size:24px;top:-1.7em;transition:font-size 0.218s ease-out}.s2wfp{}._ypc{margin-top:-100px;opacity:0;pointer-events:none;position:absolute;width:500px;transition:opacity 0.218s ease-in,margin-top .4s ease-in}.s2wfp ._ypc{margin-top:-300px;opacity:1;transition:opacity .5s ease-out 0.218s,margin-top 0.218s ease-out 0.218s}._zpc{box-shadow:0 1px 0px #4285F4;height:80px;left:0;margin:0;opacity:0;pointer-events:none;position:fixed;right:0;top:-80px;transition:opacity 0.218s,box-shadow 0.218s}.s2wfp ._zpc{box-shadow:0 1px 80px #4285F4;opacity:1;pointer-events:none;-webkit-animation:allow-alert .75s 0 infinite;-webkit-animation-direction:alternate;-webkit-animation-timing-function:ease-out;transition:opacity 0.218s,box-shadow 0.218s}@-webkit-keyframes allow-alert {from{opacity:1}to{opacity:.35}}</style></div> <a href="/setprefs?suggon=2&amp;prev=https://www.google.com/?gws_rd%3Dssl&amp;sig=0_GPJhogRFG5XIzq6ZErqatAJOSJE%3D" style="left:-1000em;position:absolute">Screen reader users, click here to turn off Google Instant.</a>  <textarea name="csi" id="csi" style="display:none"></textarea><script>if(google.j.b)document.body.style.visibility='hidden';</script><div data-jibp="" id="mngb">   <div class="gb_Oc gb_Vc" id="gb"><div class="gb_Vb gb_Sc"><div class="gb_8 gb_Sc gb_i gb_Rc gb_Vc"><div class="gb_Dc gb_i gb_Sc gb_Lc gb_k"><div class="gb_h gb_i gb_j gb_Sc"><a class="gb_g gb_i" href="https://plus.google.com/?gpsrc=ogpy0&amp;tab=wX" data-pid="119" data-ved="0CAIQwi4oAA">+You</a></div><div class="gb_h gb_i"><a class="gb_g" href="https://mail.google.com/mail/?tab=wm" data-pid="23" data-ved="0CAMQwi4oAQ">Gmail</a></div><div class="gb_h gb_i"><a class="gb_g" href="https://www.google.com/imghp?hl=en&amp;tab=wi&amp;ei=zKjSVI7_L4HfoASqnIKIDg&amp;ved=0CAQQqi4oAg" data-pid="2">Images</a></div></div><div class="gb_Cc gb_Sc gb_i"><div class="gb_H" id="gbsfw"></div><div class="gb_C gb_kb gb_i" id="gbwa"><div class="gb_Za"><a class="gb_D gb_Ta" href="http://www.google.com/intl/en/options/" title="Apps" aria-haspopup="true" aria-expanded="false" data-ved="0CAUQvSc"></a><div class="gb_ka"></div><div class="gb_ja"></div></div></div><div class="gb_Ec gb_i"><div class="gb_Za"><a class="gb_9b gb_0 gb_Z" id="gb_70" href="https://accounts.google.com/ServiceLogin?hl=en&amp;continue=https://www.google.com/%3Fgws_rd%3Dssl" target="_top">Sign in</a><div class="gb_ka"></div><div class="gb_ja"></div></div></div></div></div><div class="gb_Oa gb_i gbqfh gb_Sa" id="gbq1" style="max-width:127px;min-width:127px"><div class="gb_Pa"><a class="gb_Wa gb_Ra" href="/webhp?tab=ww&amp;ei=zKjSVI7_L4HfoASqnIKIDg&amp;ved=0CAYQ1S4" title="Go to Google Home"><span class="gb_Ta"></span></a></div></div><div class="gb_i gb_Db"><div id="gbq"><div class="gbt gbqfh" id="gbq2"><div id="gbqfw"><form class="gb_8b" action="/search" onsubmit="" target="" id="gbqf" name="gbqf" method="get" data-ved="0CAcQuyc"><fieldset class="gbxx"><legend class="gbxx">Hidden fields</legend><div id="gbqffd"><input name="output" value="search" type="hidden"><input name="sclient" value="psy-ab" type="hidden"></div></fieldset><fieldset class="gbqff gb_i" id="gbqff"><legend class="gbxx"></legend><div id=gbfwa class="gbqfwa "><div id=gbqfqw class=gbqfqw><div id=gbqfqwb class=gbqfqwc><input id=gbqfq class=gbqfif name=q type=text autocomplete=off value=""></div></div></div></fieldset><div class="gb_i gb_7b" id="gbqfbw"><button class="gbqfb" aria-label="Google Search" name="btnG" id="gbqfb"><span class="gbqfi gb_Ta"></span></button></div><div class="jsb" id="gbqfbwa"><button class="gbqfba" aria-label="Google Search" id="gbqfba" name="btnK"><span id="gbqfsa">Google Search</span></button><button class="gbqfba" aria-label="I'm Feeling Lucky" onclick="google.x(this,function() {google.ifl &amp;&amp; google.ifl.o();})" id="gbqfbb" name="btnI"><span id="gbqfsb">I'm Feeling Lucky</span></button></div></form></div></div></div></div></div><div id="gbw"></div></div><div id="gba"></div>  <div style="display:none" id="iflved" data-ved="0CAQQnRs"></div></div><div class="spch s2fp-h" style="display:none" id="spch"><div class="spchc" id="spchc"><div class="_o3"><div class="_AM"><span class="_CMb" id="spchl"></span><span class="button" id="spchb"><div class="_wPb"><span class="_AUb"></span><div class="_Fjd"><span class="_oXb"></span><span class="_dWb"></span></div></div></span></div><div class="_gjb"><span class="spcht" id="spchi" style="color:#777"></span><span class="spcht" id="spchf" style="color:#000"></span></div><div class="google-logo"></div></div><div class="_ypc"><div class="_zpc"></div></div></div><div class="close-button" id="spchx">&times;</div></div><div class="content" data-jiis="cc" id="main"><span class="ctr-p" data-jiis="bp" id="body"><center><div style="height:231px;margin-top:20px" id="lga"><img alt="Google" height="95" id="hplogo" src="/images/srpr/logo11w.png" style="padding-top:112px" width="269" onload="window.lol&&lol()"></div><div style="height:102px"><div id="chw-o" data-ved="0CAYQuzY"><div class="_lSc"><div class="_mSc">Say "Ok Google" to start a voice search.</div><p class="_kSc">Search without lifting a finger. When you say "Ok Google," Chrome will search for what you say next.</p><div><a href="https://support.google.com/chrome/?p=ui_hotword_search" target="_blank" onmousedown="return rwt(this,'','','','','AFQjCNHyNmNOmnFoJjpgISA7IEqNETP_hw','','0CAcQgDY','','',event)">Learn more</a><button class="_dKb _k3" href="#" id="hotword__chw-on" jsaction="chw.optInNoThanksButtonClicked" data-ved="0CAgQ_jU">No thanks</button><button class="_dKb _WW" href="#" id="hotword__chw-oe" jsaction="chw.optInEnableButtonClicked" data-ved="0CAkQ_zU">Enable "Ok Google"</button></div></div><div class="_jSc"><div class="_tyd"><label class="_fdc"><input class="_hSc" checked="true" type="checkbox" jsaction="chw.handleAudioLoggingInteraction"><span class="_iSc">Improve voice search by sending the sound of "Ok Google," and a few seconds before, to Google.</span></label></div></div></div></div><div id="prm-pt" style="margin-top:12px"><script>window.gbar&&gbar.up&&gbar.up.tp&&gbar.up.tp();</script></div></center></span><div class="ctr-p" data-jiis="bp" id="footer"> <div id="footcnt"> <style>.fmulti{}._dQc{bottom:0;left:0;position:absolute;right:0}._GR{background:#f2f2f2;left:0;right:0;-webkit-text-size-adjust:none}.fbar p{display:inline}.fbar a,#fsettl{text-decoration:none;white-space:nowrap}.fbar{margin-left:-27px}._Gs{padding-left:27px;margin:0 !important}._eA{padding:0 !important;margin:0 !important}#fbarcnt{display:block}._E2 a:hover{text-decoration:underline}._HR img{margin-right:4px}._HR a,._GR #swml a{text-decoration:none}.fmulti{text-align:center}.fmulti #fsr{display:block;float:none}.fmulti #fuser{display:block;float:none}#fuserm{line-height:25px}#fsr{float:right;white-space:nowrap}#fsl{white-space:nowrap}#fsett{background:#fff;border:1px solid #999;bottom:30px;padding:10px 0;position:absolute;box-shadow:0 2px 4px rgba(0,0,0,0.2);-webkit-box-shadow:0 2px 4px rgba(0,0,0,0.2);text-align:left;z-index:100}#fsett a{display:block;line-height:44px;padding:0 20px;text-decoration:none;white-space:nowrap}._E2 #fsettl:hover{text-decoration:underline}._E2 #fsett a:hover{text-decoration:underline}._mk{color:#777}._Nh{color:#222;font-size:14px;font-weight:normal;-webkit-tap-highlight-color:rgba(0,0,0,0)}._Nh:hover,._Nh:active{color:#444}._Mo{display:inline-block;position:relative}._ri,._Wk{height:13px;width:8px}._ri:before,._Wk:before{border:8px solid rgba(255,255,255,0);border-radius:8px;content:'';position:absolute}._ri:before{border-left:8px solid #777;left:1px}._Wk:before{border-right:8px solid #777;left:-9px}._ri:after,._Wk:after{border:12px solid rgba(255,255,255,0);content:'';position:absolute;top:-4px}._ri:after{border-left:10px solid #f6f6f6;left:-4px}._Wk:after{border-right:10px solid #f6f6f6;left:-10px}._sN ._ri:after{border-left:10px solid white}._sN ._Wk:after{border-right:10px solid white}._Nh{padding:8px 16px;margin-right:10px}._mk{margin:40px 0}._dD{margin-right:10px}._nW{margin-left:135px}#fbar{background:#f2f2f2;border-top:1px solid #e4e4e4;line-height:40px;min-width:980px}._xac{margin-left:135px}.fbar p,.fbar a,#fsettl,#fsett a{color:#777}.fbar a:hover,#fsett a:hover{color:#333}.fbar{font-size:small}#fuser{float:right}._HR{margin-left:135px;line-height:15px;}</style> <style>#fsl{margin-left:30px}#fsr{margin-right:30px}.fmulti #fsl{margin-left:0}.fmulti #fsr{margin-right:0}</style> <div class="_dQc _E2" id="fbar"> <div class="fbar"> <span id="fsr">  <a class="_Gs" href="//www.google.com/intl/en/policies/privacy/?fg=1">Privacy</a> <a class="_Gs" href="//www.google.com/intl/en/policies/terms/?fg=1">Terms</a>    <span style="display:inline-block;position:relative"> <a class="_Gs" href="https://www.google.com/preferences?hl=en" id="fsettl" aria-controls="fsett" aria-expanded="false" role="button" jsaction="foot.cst">Settings</a> <span id="fsett" style="display:none"> <a href="https://www.google.com/preferences?hl=en&amp;fg=1">Search settings</a> <span data-jibp="h" data-jiis="uc" id="advsl"> <a href="/advanced_search?hl=en&amp;fg=1">Advanced search</a> </span> <a href="/history?hl=en&amp;fg=1">  History </a> <a href="//support.google.com/websearch/?p=ws_results_help&amp;hl=en&amp;fg=1">Search Help</a> <a href="javascript:void(0)" data-bucket="websearch" id="_Yvd" target="_blank" jsaction="gf.sf" data-ved="0CAwQLg">  Send feedback </a> </span> </span>   </span> <span id="fsl"> <a class="_Gs" href="//www.google.com/intl/en/ads/?fg=1">Advertising</a> <a class="_Gs" href="//www.google.com/services/?fg=1">Business</a> <a class="_Gs" href="//www.google.com/intl/en/about.html?fg=1">About</a> </span> </div>  </div> </div>   </div></div><script>(function(){var _mstr='\74span class\75ctr-p id\75body data-jiis\75bp\76\74/span\76\74span class\75ctr-p id\75footer data-jiis\75bp\76\74/span\76\74span id\75xjsi\76\74/span\076';var commands = [];var index = 0;var gstyle = document.getElementById('gstyle');if (gstyle){commands[index++]=
{'n':'pcs','i':'gstyle','css':gstyle.innerHTML,'is':'','r':true,'sc':true};}
commands[index++]=
{'n':'pc','i':'cst','h':document.getElementById('cst').innerHTML,'is':'','r':true,'sc':true};commands[index]=
{'n':'pc','i':'main','h':_mstr,'is':'','r':true,'sc':true};google.j[1]={cmds:commands};})();</script><script data-url="/extern_chrome/7f5ab5fc6a7abe9a.js?bav=or.r_qf" id="ecs"></script><script>this.gbar_=this.gbar_||{};(function(_){var window=this;
try{
_.Xa=function(a){var c=_.Ha;c.d?a():c.b.push(a)};_.Ya=function(){};_.Za=function(a){_.Za[" "](a);return a};_.Za[" "]=_.Ya;
}catch(e){_._DumpException(e)}
try{
var Ng;Ng=function(a){if(a.classList)return a.classList;a=a.className;return _.t(a)&&a.match(/\S+/g)||[]};_.R=function(a,c){return a.classList?a.classList.contains(c):_.pa(Ng(a),c)};_.S=function(a,c){a.classList?a.classList.add(c):_.R(a,c)||(a.className+=0<a.className.length?" "+c:c)};
_.Og=function(a,c){if(a.classList)(0,_.ka)(c,function(c){_.S(a,c)});else{var d={};(0,_.ka)(Ng(a),function(a){d[a]=!0});(0,_.ka)(c,function(a){d[a]=!0});a.className="";for(var e in d)a.className+=0<a.className.length?" "+e:e}};_.T=function(a,c){a.classList?a.classList.remove(c):_.R(a,c)&&(a.className=(0,_.la)(Ng(a),function(a){return a!=c}).join(" "))};_.Pg=function(a,c){a.classList?(0,_.ka)(c,function(c){_.T(a,c)}):a.className=(0,_.la)(Ng(a),function(a){return!_.pa(c,a)}).join(" ")};

}catch(e){_._DumpException(e)}
try{
var Dj,Jj,Mj,Lj;_.yj=function(a){_.A(this,a,0,null)};_.w(_.yj,_.z);_.zj=function(){return _.D(_.J(),_.yj,11)};_.Aj=function(a){_.A(this,a,0,null)};_.w(_.Aj,_.z);_.Cj=function(){var a=_.Bj();return _.B(a,9)};Dj=function(a){_.A(this,a,0,null)};_.w(Dj,_.z);_.Ej=function(a){_.A(this,a,0,null)};_.w(_.Ej,_.z);var Fj=function(a){_.A(this,a,0,null)};_.w(Fj,_.z);_.Gj=function(a){return _.B(a,10)};_.Hj=function(a){return _.B(a,5)};_.Bj=function(){return _.D(_.J(),Dj,4)||new Dj}; _.Kj=function(a){var c="//www.google.com/gen_204?",c=c+a.d(2040-c.length);Jj(c)};Jj=function(a){var c=new window.Image,d=Lj;c.onerror=c.onload=c.onabort=function(){d in Mj&&delete Mj[d]};Mj[Lj++]=c;c.src=a};Mj=[];Lj=0;
_.Nj=function(){this.data={}};_.Nj.prototype.b=function(){window.console&&window.console.log&&window.console.log("Log data: ",this.data)};_.Nj.prototype.d=function(a){var c=[],d;for(d in this.data)c.push((0,window.encodeURIComponent)(d)+"="+(0,window.encodeURIComponent)(String(this.data[d])));return("atyp=i&zx="+(new Date).getTime()+"&"+c.join("&")).substr(0,a)};
var Oj=function(a){this.b=a};Oj.prototype.log=function(a,c){try{if(this.A(a)){var d=this.k(a,c);this.d(d)}}catch(e){}};Oj.prototype.d=function(a){this.b?a.b():_.Kj(a)};var Pj=function(a,c){this.data={};var d=_.D(a,_.ta,8)||new _.ta;this.data.ei=_.G(_.Gj(a));this.data.ogf=_.G(_.B(d,3));var e;e=window.google&&window.google.sn?/.*hp$/.test(window.google.sn)?!1:!0:_.F(_.B(a,7));this.data.ogrp=e?"1":"";this.data.ogv=_.G(_.B(d,6))+"."+_.G(_.B(d,7));this.data.ogd=_.G(_.B(a,21));this.data.ogc=_.G(_.B(a,20));this.data.ogl=_.G(_.Hj(a));c&&(this.data.oggv=c)};_.w(Pj,_.Nj);
_.Qj=function(a,c,d,e,f){Pj.call(this,a,c);_.sa(this.data,{jexpid:_.G(_.B(a,9)),srcpg:"prop="+_.G(_.B(a,6)),jsr:Math.round(1/e),emsg:d.name+":"+d.message});if(f){f._sn&&(f._sn="og."+f._sn);for(var g in f)this.data[(0,window.encodeURIComponent)(g)]=f[g]}};_.w(_.Qj,Pj);
var Rj=[1,2,3,4,5,6,9,10,11,13,14,28,29,30,34,35,37,38,39,40,41,42,43,48,49,50,51,52,53,500],Uj=function(a,c,d,e,f,g){Pj.call(this,a,c);_.sa(this.data,{oge:e,ogex:_.G(_.B(a,9)),ogp:_.G(_.B(a,6)),ogsr:Math.round(1/(Sj(e)?_.H(null!=_.B(d,3)?_.B(d,3):1):_.H(null!=_.B(d,2)?_.B(d,2):1E-4))),ogus:f});if(g){"ogw"in g&&(this.data.ogw=g.ogw,delete g.ogw);"ved"in g&&(this.data.ved=g.ved,delete g.ved);a=[];for(var h in g)0!=a.length&&a.push(","),a.push(Tj(h)),a.push("."),a.push(Tj(g[h]));g=a.join("");""!=g&& (this.data.ogad=g)}};_.w(Uj,Pj);var Tj=function(a){return(a+"").replace(".","%2E").replace(",","%2C")},Vj=null,Sj=function(a){if(!Vj){Vj={};for(var c=0;c<Rj.length;c++)Vj[Rj[c]]=!0}return!!Vj[a]};
var Wj=function(a,c,d,e,f){this.b=f;this.F=a;this.O=c;this.ja=e;this.B=_.H(null!=_.B(a,2)?_.B(a,2):1E-4,1E-4);this.w=_.H(null!=_.B(a,3)?_.B(a,3):1,1);c=Math.random();this.C=_.F(_.B(a,1))&&c<this.B;this.o=_.F(_.B(a,1))&&c<this.w;a=0;_.F(_.B(d,1))&&(a|=1);_.F(_.B(d,2))&&(a|=2);_.F(_.B(d,3))&&(a|=4);this.J=a};_.w(Wj,Oj);Wj.prototype.A=function(a){return this.b||(Sj(a)?this.o:this.C)};Wj.prototype.k=function(a,c){return new Uj(this.O,this.ja,this.F,a,this.J,c)};
var Xj=function(a,c,d,e){this.b=e;this.F=c;this.ja=d;this.w=_.H(null!=_.B(a,2)?_.B(a,2):.001,.001);this.O=_.F(_.B(a,1))&&Math.random()<this.w;this.C=_.H(null!=_.B(a,3)?_.B(a,3):1,1);this.o=0;this.B=_.F(null!=_.B(a,4)?_.B(a,4):!0,!0)};_.w(Xj,Oj);Xj.prototype.log=function(a,c){Xj.G.log.call(this,a,c);if(this.b&&this.B)throw a;};Xj.prototype.A=function(){return this.b||this.O&&this.o<this.C}; Xj.prototype.k=function(a,c){try{return _.wa(_.va.N(),"lm").Df(a,c)}catch(d){return new _.Qj(this.F,this.ja,a,this.w,c)}};Xj.prototype.d=function(a){Xj.G.d.call(this,a);this.o++};
var Yj;Yj=null;_.Zj=function(){if(!Yj){var a=_.D(_.J(),_.Ej,13)||new _.Ej,c=_.Ja();Yj=new Xj(a,c,_.Cj(),_.Fa)}return Yj};_.I=function(a,c){_.Zj().log(a,c)};_.ak=function(a,c){return function(){try{return a.apply(c,arguments)}catch(d){_.I(d)}}};var bk;bk=null;_.ck=function(){if(!bk){var a=_.D(_.J(),Fj,12)||new Fj,c=_.Ja(),d=_.zj()||new _.yj;bk=new Wj(a,c,d,_.Cj(),_.Fa)}return bk};_.V=function(a,c){_.ck().log(a,c)};_.V(8,{m:"BackCompat"==window.document.compatMode?"q":"s"});
}catch(e){_._DumpException(e)}
try{
var dk,hk,jk;dk=[3,5];_.ek=function(a){_.A(this,a,0,dk)};_.w(_.ek,_.z);var fk=function(a){_.A(this,a,0,null)};_.w(fk,_.z);_.gk=function(a){_.A(this,a,0,null)};_.w(_.gk,_.z);_.gk.prototype.Ra=function(){return _.B(this,6)};
_.ik=function(a,c,d,e,f,g){_.y.call(this);this.F=c;this.J=f;this.o=g;this.R=!1;this.k={"":!0};this.H={"":!0};this.A=[];this.w=[];this.Q=["//"+_.G(_.B(a,2)),"og/_/js","k="+_.G(_.B(a,3)),"rt=j"];this.O=""==_.G(_.B(a,14))?null:_.B(a,14);this.K=["//"+_.G(_.B(a,2)),"og/_/ss","k="+_.G(_.B(a,13))];this.B=""==_.G(_.B(a,15))?null:_.B(a,15);this.U=_.F(_.B(a,1))?"?host=www.gstatic.com&bust="+_.G(_.B(a,16)):"";this.T=_.F(_.B(a,1))?"?host=www.gstatic.com&bust="+1E11*Math.random():"";this.d=d;this.b=_.H(null!=
_.B(a,17)?_.B(a,17):1,1);a=0;for(c=e[a];a<e.length;a++,c=e[a])hk(this,c,!0)};_.w(_.ik,_.y);_.xa(_.ik,"m");hk=function(a,c,d){if(!a.k[c]&&(a.k[c]=!0,d&&a.d[c]))for(var e=0;e<a.d[c].length;e++)hk(a,a.d[c][e],d)};jk=function(a,c){for(var d=[],e=0;e<c.length;e++){var f=c[e];if(!a.k[f]){var g=a.d[f];g&&(g=jk(a,g),d=d.concat(g));d.push(f);a.k[f]=!0}}return d};_.lk=function(a,c,d){c=jk(a,c);0<c.length&&(c=a.Q.join("/")+"/"+("m="+c.join(",")),a.O&&(c+="/rs="+a.O),c=c+a.U,kk(a,c,(0,_.u)(a.P,a,d)),a.A.push(c))};
_.ik.prototype.P=function(a){_.E("api").Ua();for(var c=0;c<this.w.length;c++)this.w[c].call(null);a&&a.call(null)};
var kk=function(a,c,d,e){var f=window.document.createElement("SCRIPT");f.async=!0;f.type="text/javascript";f.charset="UTF-8";f.src=c;var g=!0,h=e||1;e=(0,_.u)(function(){g=!1;this.o.log(47,{att:h,max:a.b,url:c});h<a.b?kk(a,c,d,h+1):this.J.log(Error("U"+h+""+a.b),{url:c})},a);var l=(0,_.u)(function(){g&&(this.o.log(46,{att:h,max:a.b,url:c}),g=!1,d&&d.call(null))},a),q=function(a){"loaded"==a.readyState||"complete"==a.readyState?l():g&&window.setTimeout(function(){q(a)},100)};"undefined"!==typeof f.addEventListener?
f.onload=function(){l()}:f.onreadystatechange=function(){f.onreadystatechange=null;q(f)};f.onerror=e;a.o.log(45,{att:h,max:a.b,url:c});window.document.getElementsByTagName("HEAD")[0].appendChild(f)};_.ik.prototype.Od=function(a,c){for(var d=[],e=0,f=a[e];e<a.length;e++,f=a[e])this.H[f]||(d.push(f),this.H[f]=!0);0<d.length&&(d=this.K.join("/")+"/"+("m="+d.join(",")),this.B&&(d+="/rs="+this.B),d+=this.T,mk(d,c))};
var mk=function(a,c){var d=window.document.createElement("LINK");d.setAttribute("rel","stylesheet");d.setAttribute("type","text/css");d.setAttribute("href",a);d.onload=d.onreadystatechange=function(){d.readyState&&"loaded"!=d.readyState&&"complete"!=d.readyState||c&&c.call(null)};window.document.getElementsByTagName("HEAD")[0].appendChild(d)};
_.ik.prototype.C=function(a){this.R||(void 0!=a?window.setTimeout((0,_.u)(this.C,this,void 0),a):(_.lk(this,_.B(this.F,1),(0,_.u)(this.S,this)),this.Od(_.B(this.F,2)),this.R=!0))};_.ik.prototype.S=function(){_.v("gbar.qm",(0,_.u)(function(a){try{a()}catch(c){this.J.log(c)}},this))};
var nk=function(a,c){var d={};d._sn=["v.gas",c].join(".");_.I(a,d)};var ok=["gbq1","gbq2","gbqfbwa"],pk=function(a){var c=window.document.getElementById("gbqld");c&&(c.style.display=a?"none":"block",c=window.document.getElementById("gbql"))&&(c.style.display=a?"block":"none")};var qk=function(){};var rk=function(a,c,d){this.d=a;this.k=c;this.b=d||_.m};var sk=function(){this.b=[]};sk.prototype.A=function(a,c,d){this.F(a,c,d);this.b.push(new rk(a,c,d))};sk.prototype.F=function(a,c,d){d=d||_.m;for(var e=0,f=this.b.length;e<f;e++){var g=this.b[e];if(g.d==a&&g.k==c&&g.b==d){this.b.splice(e,1);break}}};sk.prototype.w=function(a){for(var c=0,d=this.b.length;c<d;c++){var e=this.b[c];"hrc"==e.d&&e.k.call(e.b,a)}};
var tk,vk,wk,xk,yk;tk=null;_.uk=function(){if(null!=tk)return tk;var a=window.document.body.style;if(!(a="flexGrow"in a||"webkitFlexGrow"in a))t:{if(a=window.navigator.userAgent){var c=/Trident\/(\d+)/.exec(a);if(c&&7<=Number(c[1])){a=/\bMSIE (\d+)/.exec(a);a=!a||"10"==a[1];break t}}a=!1}return tk=a};
vk=function(a,c,d){var e=window.NaN;window.getComputedStyle&&(a=window.getComputedStyle(a,null).getPropertyValue(c))&&"px"==a.substr(a.length-2)&&(e=d?(0,window.parseFloat)(a.substr(0,a.length-2)):(0,window.parseInt)(a.substr(0,a.length-2),10));return e};
wk=function(a){var c=a.offsetWidth,d=vk(a,"width");if(!(0,window.isNaN)(d))return c-d;var e=a.style.padding,f=a.style.paddingLeft,g=a.style.paddingRight;a.style.padding=a.style.paddingLeft=a.style.paddingRight=0;d=a.clientWidth;a.style.padding=e;a.style.paddingLeft=f;a.style.paddingRight=g;return c-d};
xk=function(a){var c=vk(a,"min-width");if(!(0,window.isNaN)(c))return c;var d=a.style.width,e=a.style.padding,f=a.style.paddingLeft,g=a.style.paddingRight;a.style.width=a.style.padding=a.style.paddingLeft=a.style.paddingRight=0;c=a.clientWidth;a.style.width=d;a.style.padding=e;a.style.paddingLeft=f;a.style.paddingRight=g;return c};yk=function(a,c){c||-.5!=a-Math.round(a)||(a-=.5);return Math.round(a)}; _.zk=function(a){if(a){var c=a.style.opacity;a.style.opacity=".99";_.Za(a.offsetWidth);a.style.opacity=c}};
var Ak=function(a){_.y.call(this);this.b=a;this.d=[];this.k=[]};_.w(Ak,_.y);Ak.prototype.L=function(){Ak.G.L.call(this);this.b=null;for(var a=0;a<this.d.length;a++)this.d[a].Y();for(a=0;a<this.k.length;a++)this.k[a].Y();this.d=this.k=null};
Ak.prototype.La=function(a){void 0==a&&(a=this.b.offsetWidth);for(var c=wk(this.b),d=[],e=0,f=0,g=0,h=0,l=0;l<this.d.length;l++){var q=this.d[l],r=Bk(q),x=wk(q.b);d.push({item:q,gb:r,Ch:x,uc:0});e+=r.Hc;f+=r.Wc;g+=r.Sb;h+=x}a=a-h-c-g;e=0<a?e:f;f=a;c=d;do{g=!0;h=[];for(l=q=0;l<c.length;l++){var r=c[l],x=0<f?r.gb.Hc:r.gb.Wc,C=0==e?0:x/e*f+r.uc,C=yk(C,g),g=!g;r.uc=Ck(r.item,C,r.Ch,r.gb.Sb);0<x&&C==r.uc&&(h.push(r),q+=x)}c=h;f=a-(0,_.na)(d,function(a,c){return a+c.uc},0);e=q}while(0!=f&&0!=c.length);
for(l=0;l<this.k.length;l++)this.k[l].La()};var Ek=function(a){var c={};c.items=(0,_.ma)(a.d,function(a){return Dk(a)});c.children=(0,_.ma)(a.k,function(a){return Ek(a)});return c},Fk=function(a,c){for(var d=0;d<a.d.length;d++)a.d[d].b.style.width=c.items[d];for(d=0;d<a.k.length;d++)Fk(a.k[d],c.children[d])};Ak.prototype.M=function(){return this.b};
var Gk=function(a,c,d,e){Ak.call(this,a);this.w=c;this.A=d;this.o=e};_.w(Gk,Ak);
var Bk=function(a,c){var d=a.w,e=a.A,f;if(-1==a.o){var g=c;void 0==g&&(g=wk(a.b));f=Dk(a);var h=Ek(a),l=vk(a.b,"width",!0);(0,window.isNaN)(l)&&(l=a.b.offsetWidth-g);g=Math.ceil(l);a.b.style.width=f;Fk(a,h);f=g}else f=a.o;return{Hc:d,Wc:e,Sb:f}},Ck=function(a,c,d,e){void 0==d&&(d=wk(a.b));void 0==e&&(e=Bk(a,d).Sb);c=e+c;0>c&&(c=0);a.b.style.width=c+"px";d=a.b.offsetWidth-d;a.b.style.width=d+"px";return d-e},Dk=function(a){var c=a.b.style.width;a.b.style.width="";return c};
var Hk=function(a,c,d){var e;void 0==e&&(e=-1);return{className:a,gb:{Hc:c||0,Wc:d||0,Sb:e}}},Ik={className:"gb_Vb",items:[Hk("gb_Oa"),Hk("gb_bc"),Hk("gb_Db",0,2),Hk("gb_cc"),Hk("gb_8",1,1)],eb:[{className:"gb_8",items:[Hk("gb_Dc",0,1),Hk("gb_Cc",0,1)],eb:[function(a){a=a.gb_Dc;var c;if(a)c=a.M();else{c=window.document.querySelector(".gb_Dc");if(!c)return null;a=new Ak(c)}c=c.querySelectorAll(".gb_h");for(var d=0;d<c.length;d++){var e;if(_.R(c[d],"gb_j")){e=new Gk(c[d],0,1,-1);var f=c[d].querySelector(".gb_g");
f&&(f=new Gk(f,0,1,-1),e.d.push(f),a.k.push(e))}else e=new Gk(c[d],0,0,-1);a.d.push(e)}return a},{className:"gb_Cc",items:[Hk("gb_C"),Hk("gb_Xa"),Hk("gb_Rb"),Hk("gb_ga",0,1),Hk("gb_Ec"),Hk("gb_ca",0,1),Hk("gb_Fc"),Hk("gb_ec")],eb:[{className:"gb_ga",items:[Hk("gb_ia",0,1)],eb:[{className:"gb_ia",items:[Hk("gb_ea",0,1)],eb:[]}]}]}]},{className:"gb_8b",items:[Hk("gbqff",1,1),Hk("gb_7b")],eb:[]}]},Jk=function(a,c){var d=c;if(!d){d=window.document.querySelector("."+a.className);if(!d)return null;d=new Ak(d)}for(var e=
{},f=0;f<a.items.length;f++){var g=a.items[f],h;h=g;var l=window.document.querySelector("."+h.className);if(h=l?new Gk(l,h.gb.Hc,h.gb.Wc,h.gb.Sb):null)d.d.push(h),e[g.className]=h}for(f=0;f<a.eb.length;f++){var g=a.eb[f],q;"function"==typeof g?q=g(e):q=Jk(g,e[g.className]);q&&d.k.push(q)}return d};
_.Mk=function(a){_.y.call(this);this.B=new sk;this.d=window.document.getElementById("gb");this.O=(this.b=window.document.querySelector(".gb_8"))?this.b.querySelector(".gb_Cc"):null;this.C=[];this.me=60;this.J=_.B(a,4);this.Kh=_.H(_.B(a,2),152);this.Rf=_.H(_.B(a,1),30);this.k=null;this.Re=_.F(_.B(a,3),!0);this.o=1;this.d&&this.J&&(this.d.style.minWidth=this.J+"px");_.Kk(this);this.Re&&(this.d&&(Lk(this),_.S(this.d,"gb_Eb"),_.uk()||(this.k=Jk(Ik))),this.La(),window.setTimeout((0,_.u)(this.La,this),
0));_.v("gbar.elc",(0,_.u)(this.K,this));_.v("gbar.ela",_.Ya);_.v("gbar.elh",(0,_.u)(this.R,this))};_.w(_.Mk,_.y);_.xa(_.Mk,"el");var Nk=function(){var a=_.Mk.Wh();return{es:a?{f:a.Kh,h:a.me,m:a.Rf}:{f:152,h:60,m:30},mo:"md",vh:window.innerHeight||0,vw:window.innerWidth||0}};_.Mk.prototype.L=function(){_.Mk.G.L.call(this)};_.Mk.prototype.La=function(a){a&&Lk(this);this.k&&this.k.La(Math.max(window.document.documentElement.clientWidth,xk(this.d)));_.zk(this.b)};
_.Mk.prototype.H=function(){try{var a=window.document.getElementById("gb"),c=a.querySelector(".gb_8");_.T(a,"gb_Vc");c&&_.T(c,"gb_Vc");for(var a=0,d;d=ok[a];a++)_.T(window.document.getElementById(d),"gbqfh");pk(!1)}catch(e){nk(e,"rhcc")}this.La(!0)};_.Mk.prototype.T=function(){try{var a=window.document.getElementById("gb"),c=a.querySelector(".gb_8");_.S(a,"gb_Vc");c&&_.S(c,"gb_Vc");for(var a=0,d;d=ok[a];a++)_.S(window.document.getElementById(d),"gbqfh");pk(!0)}catch(e){nk(e,"ahcc")}this.La(!0)};
_.Kk=function(a){if(a.d){var c=a.d.offsetWidth;0==a.o?900<=c&&(a.o=1,a.w(new qk)):900>c&&(a.o=0,a.w(new qk))}};_.Mk.prototype.K=function(a){this.C.push(a)};_.Mk.prototype.R=function(a){var c=Nk().es.h;this.me=c+a;for(a=0;a<this.C.length;a++)try{this.C[a](Nk())}catch(d){_.I(d)}};var Lk=function(a){if(a.b){var c;a.k&&(c=Ek(a.k));_.S(a.b,"gb_m");a.b.style.minWidth=a.b.offsetWidth-wk(a.b)+"px";a.O.style.minWidth=a.O.offsetWidth-wk(a.O)+"px";_.T(a.b,"gb_m");c&&Fk(a.k,c)}}; _.Mk.prototype.A=function(a,c,d){this.B.A(a,c,d)};_.Mk.prototype.F=function(a,c){this.B.F(a,c)};_.Mk.prototype.w=function(a){this.B.w(a)};
_.Xa(function(){var a=_.D(_.J(),fk,21)||new fk,a=new _.Mk(a);_.za("el",a);_.v("gbar.gpca",(0,_.u)(a.T,a));_.v("gbar.gpcr",(0,_.u)(a.H,a))});_.v("gbar.elr",Nk);_.Ok=function(a){this.k=_.Mk.N();this.d=a};_.Ok.prototype.b=function(a,c){0==this.k.o?(_.S(a,"gb_l"),c?(_.T(a,"gb_fa"),_.S(a,"gb_Hc")):(_.T(a,"gb_Hc"),_.S(a,"gb_fa"))):_.Pg(a,["gb_l","gb_fa","gb_Hc"])};_.v("gbar.sos",function(){return window.document.querySelectorAll(".gb_ac")});_.v("gbar.si",function(){return window.document.querySelector(".gb_9b")});
_.Xa(function(){if(_.D(_.J(),_.ek,16)){var a=window.document.querySelector(".gb_8"),c=_.D(_.J(),_.ek,16)||new _.ek,c=new _.Ok(_.F(_.B(c,1),!1));a&&c.d&&c.b(a,!1)}});
}catch(e){_._DumpException(e)}
try{
var Pk=function(){_.Ha.k(_.I)};var Qk=function(a,c){var d=_.ak(Pk);a.addEventListener?a.addEventListener(c,d):a.attachEvent("on"+c,d)};var Rk=[1,2],Sk=function(a,c){a.w.push(c)},Tk=function(a){_.A(this,a,0,Rk)};_.w(Tk,_.z);var Uk=function(){_.y.call(this);this.o=this.b=null;this.d={};this.w={};this.k={}};_.w(Uk,_.y);_.k=Uk.prototype;_.k.hf=function(a){a&&this.b&&a!=this.b&&this.b.close();this.b=a};_.k.Te=function(a){a=this.k[a]||a;return this.b==a};_.k.Oh=function(a){this.o=a};
_.k.Se=function(a){return this.o==a};_.k.kd=function(){this.b&&this.b.close();this.b=null};_.k.Cf=function(a){this.b&&this.b.getId()==a&&this.kd()};_.k.Pb=function(a,c,d){this.d[a]=this.d[a]||{};this.d[a][c]=this.d[a][c]||[];this.d[a][c].push(d)};_.k.hd=function(a,c){var d=c.getId();if(this.d[a]&&this.d[a][d])for(var e=0;e<this.d[a][d].length;e++)try{this.d[a][d][e]()}catch(f){_.I(f)}};_.k.Sh=function(a,c){this.w[a]=c};_.k.Af=function(a){return!this.w[a.getId()]};
_.k.$g=function(){return!!this.b&&this.b.S};_.k.zf=function(){return!!this.b};_.k.Ye=function(){this.b&&this.b.ka()};_.k.mf=function(a){this.k[a]&&(this.b&&this.b.getId()==a||this.k[a].open())};_.k.sh=function(a){this.k[a.getId()]=a};var Vk;window.gbar&&window.gbar._DPG?Vk=window.gbar._DPG[0]||{}:Vk={};var Wk;window.gbar&&window.gbar._LDD?Wk=window.gbar._LDD:Wk=[];var Xk=_.Ja(),Yk=new _.ik(Xk,_.D(_.J(),Tk,17)||new Tk,Vk,Wk,_.Zj(),_.ck());_.za("m",Yk); if(_.F(_.B(Xk,18),!0))Yk.C();else{var Zk=(0,_.u)(Yk.C,Yk,_.H(_.B(Xk,19),200));_.Xa(Zk)}Qk(window.document,"DOMContentLoaded");Qk(window,"load");
_.v("gbar.ldb",(0,_.u)(_.Ha.k,_.Ha));_.v("gbar.mls",function(){});var $k=function(){_.y.call(this);this.k=this.b=null;this.w=0;this.o={};this.d=!1;var a=window.navigator.userAgent;0<=a.indexOf("MSIE")&&0<=a.indexOf("Trident")&&(a=/\b(?:MSIE|rv)[: ]([^\);]+)(\)|;)/.exec(a))&&a[1]&&9>(0,window.parseFloat)(a[1])&&(this.d=!0)};_.w($k,_.y);
var al=function(a,c,d){if(!a.d)if(d instanceof Array)for(var e in d)al(a,c,d[e]);else{e=(0,_.u)(a.A,a,c);var f=a.w+d;a.w++;c.setAttribute("data-eqid",f);a.o[f]=e;c&&c.addEventListener?c.addEventListener(d,e,!1):c&&c.attachEvent?c.attachEvent("on"+d,e):_.I(Error("V` + "`" + `"+c))}};
$k.prototype.Ue=function(a,c){if(this.d)return null;if(c instanceof Array){var d=null,e;for(e in c){var f=this.Ue(a,c[e]);f&&(d=f)}return d}d=null;this.b&&this.b.type==c&&this.k==a&&(d=this.b,this.b=null);if(e=a.getAttribute("data-eqid"))a.removeAttribute("data-eqid"),(e=this.o[e])?a.removeEventListener?a.removeEventListener(c,e,!1):a.detachEvent&&a.detachEvent("on"+c,e):_.I(Error("W` + "`" + `"+a));return d};$k.prototype.A=function(a,c){this.b=c;this.k=a;c.preventDefault?c.preventDefault():c.returnValue=!1};
_.za("eq",new $k);var bl=function(){_.y.call(this);this.le=[];this.ld=[]};_.w(bl,_.y);bl.prototype.b=function(a,c){this.le.push({vc:a,options:c})};bl.prototype.init=function(){window.gapi={};var a=_.Bj(),c=window.___jsl={};c.h=_.G(_.B(a,1));c.ms=_.G(_.B(a,2));c.m=_.G(_.B(a,3));c.l=[];a=_.D(_.J(),_.Aj,5)||new _.Aj;_.B(a,1)&&(a=_.B(a,3))&&this.ld.push(a);a=_.D(_.J(),_.gk,6)||new _.gk;_.B(a,1)&&(a=_.B(a,2))&&this.ld.push(a);_.v("gapi.load",(0,_.u)(this.b,this));return this};
var cl=_.Bj();window.__PVT=_.G(_.B(cl,7));_.za("gs",(new bl).init());(function(){for(var a=function(a){return function(){_.V(44,{n:a})}},c=0;c<_.Ma.length;c++){var d="gbar."+_.Ma[c];_.v(d,a(d))}var e=_.va.N();_.wa(e,"api").Ua();Sk(_.wa(e,"m"),function(){_.wa(e,"api").Ua()})})();var dl=function(a){_.Xa(function(){var c=window.document.querySelector("."+a);c&&(c=c.querySelector(".gb_D"))&&al(_.E("eq"),c,"click")})};var el=window.document.querySelector(".gb_C"),fl=/(\s+|^)gb_6b(\s+|$)/;el&&!fl.test(el.className)&&dl("gb_C");var gl=new Uk;_.za("dd",gl);_.v("gbar.close",(0,_.u)(gl.kd,gl));_.v("gbar.cls",(0,_.u)(gl.Cf,gl));_.v("gbar.abh",(0,_.u)(gl.Pb,gl,0));_.v("gbar.adh",(0,_.u)(gl.Pb,gl,1));_.v("gbar.ach",(0,_.u)(gl.Pb,gl,2));_.v("gbar.aeh",(0,_.u)(gl.Sh,gl));_.v("gbar.bsy",(0,_.u)(gl.$g,gl));_.v("gbar.op",(0,_.u)(gl.zf,gl));
dl("gb_ga");_.Xa(function(){var a=window.document.querySelector(".gb_Ra");a&&al(_.E("eq"),a,"click")});dl("gb_Xa");_.v("gbar.qfgw",(0,_.u)(window.document.getElementById,window.document,"gbqfqw"));_.v("gbar.qfgq",(0,_.u)(window.document.getElementById,window.document,"gbqfq"));_.v("gbar.qfgf",(0,_.u)(window.document.getElementById,window.document,"gbqf"));_.v("gbar.qfsb",(0,_.u)(window.document.getElementById,window.document,"gbqfb"));
dl("gb_Rb");dl("gb_ec");
}catch(e){_._DumpException(e)}
})(this.gbar_);
// Google Inc.
</script><div class="gb_Wb"><div class="gb_Xb gb_H gb_Q" aria-label="Apps" role="region" aria-hidden="true" tabindex="0"><ul class="gb_J gb_A" aria-dropeffect="move"><li class="gb_p" aria-grabbed="false"><a class="gb_f" id="gb119" href="https://plus.google.com/?gpsrc=ogpy0&amp;tab=wX" data-pid="119" data-ved="0CAgQwS4oAA"><span class="gb_t" style="background-position:0 -207px"></span><span class="gb_u">+You</span></a></li><li class="gb_p" aria-grabbed="false"><a class="gb_f" id="gb1" href="https://www.google.com/webhp?tab=ww&amp;ei=zKjSVI7_L4HfoASqnIKIDg&amp;ved=0CAkQqS4oAQ" data-pid="1"><span class="gb_t" style="background-position:0 -414px"></span><span class="gb_u">Search</span></a></li><li class="gb_p" aria-grabbed="false"><a class="gb_f" id="gb36" href="https://www.youtube.com/" data-pid="36" data-ved="0CAoQwS4oAg"><span class="gb_t" style="background-position:0 -1449px"></span><span class="gb_u">YouTube</span></a></li><li class="gb_p" aria-grabbed="false"><a class="gb_f" id="gb8" href="https://maps.google.com/maps?hl=en&amp;tab=wl" data-pid="8" data-ved="0CAsQwS4oAw"><span class="gb_t" style="background-position:0 -828px"></span><span class="gb_u">Maps</span></a></li><li class="gb_p" aria-grabbed="false"><a class="gb_f" id="gb78" href="https://play.google.com/?hl=en&amp;tab=w8" data-pid="78" data-ved="0CAwQwS4oBA"><span class="gb_t" style="background-position:0 -1725px"></span><span class="gb_u">Play</span></a></li><li class="gb_p" aria-grabbed="false"><a class="gb_f" id="gb5" href="https://news.google.com/nwshp?hl=en&amp;tab=wn&amp;ei=zKjSVI7_L4HfoASqnIKIDg&amp;ved=0CA0QqS4oBQ" data-pid="5"><span class="gb_t" style="background-position:0 -1656px"></span><span class="gb_u">News</span></a></li><li class="gb_p" aria-grabbed="false"><a class="gb_f" id="gb23" href="https://mail.google.com/mail/?tab=wm" data-pid="23" data-ved="0CA4QwS4oBg"><span class="gb_t" style="background-position:0 -690px"></span><span class="gb_u">Gmail</span></a></li><li class="gb_p" aria-grabbed="false"><a class="gb_f" id="gb49" href="https://drive.google.com/?tab=wo" data-pid="49" data-ved="0CA8QwS4oBw"><span class="gb_t" style="background-position:0 -1242px"></span><span class="gb_u">Drive</span></a></li><li class="gb_p" aria-grabbed="false"><a class="gb_f" id="gb24" href="https://www.google.com/calendar?tab=wc" data-pid="24" data-ved="0CBAQwS4oCA"><span class="gb_t" style="background-position:0 -1518px"></span><span class="gb_u">Calendar</span></a></li></ul><a class="gb_F gb_5b" href="http://www.google.com/intl/en/options/">More</a><span class="gb_K"></span><ul class="gb_J gb_B" aria-dropeffect="move"><li class="gb_p" aria-grabbed="false"><a class="gb_f" id="gb51" href="https://translate.google.com/?hl=en&amp;tab=wT" data-pid="51" data-ved="0CBEQwS4oCQ"><span class="gb_t" style="background-position:0 -1587px"></span><span class="gb_u">Translate</span></a></li><li class="gb_p" aria-grabbed="false"><a class="gb_f" id="gb10" href="https://books.google.com/bkshp?hl=en&amp;tab=wp&amp;ei=zKjSVI7_L4HfoASqnIKIDg&amp;ved=0CBIQqS4oCg" data-pid="10"><span class="gb_t" style="background-position:0 -138px"></span><span class="gb_u">Books</span></a></li><li class="gb_p" aria-grabbed="false"><a class="gb_f" id="gb212" href="https://wallet.google.com/manage/?tab=wa" data-pid="212" data-ved="0CBMQwS4oCw"><span class="gb_t" style="background-position:0 -621px"></span><span class="gb_u">Wallet</span></a></li><li class="gb_p" aria-grabbed="false"><a class="gb_f" id="gb6" href="http://www.google.com/shopping?hl=en&amp;tab=wf&amp;ei=zKjSVI7_L4HfoASqnIKIDg&amp;ved=0CBQQqS4oDA" data-pid="6"><span class="gb_t" style="background-position:0 -1380px"></span><span class="gb_u">Shopping</span></a></li><li class="gb_p" aria-grabbed="false"><a class="gb_f" id="gb30" href="https://www.blogger.com/?tab=wj" data-pid="30" data-ved="0CBUQwS4oDQ"><span class="gb_t" style="background-position:0 -1311px"></span><span class="gb_u">Blogger</span></a></li><li class="gb_p" aria-grabbed="false"><a class="gb_f" id="gb27" href="https://www.google.com/finance?tab=we" data-pid="27" data-ved="0CBYQwS4oDg"><span class="gb_t" style="background-position:0 -1035px"></span><span class="gb_u">Finance</span></a></li><li class="gb_p" aria-grabbed="false"><a class="gb_f" id="gb31" href="https://plus.google.com/photos?tab=wq" data-pid="31" data-ved="0CBcQwS4oDw"><span class="gb_t" style="background-position:0 -966px"></span><span class="gb_u">Photos</span></a></li><li class="gb_p" aria-grabbed="false"><a class="gb_f" id="gb25" href="https://docs.google.com/document/?usp=docs_alc" data-pid="25" data-ved="0CBgQwS4oEA"><span class="gb_t" style="background-position:0 -483px"></span><span class="gb_u">Docs</span></a></li></ul><a class="gb_K gb_Zb" href="http://www.google.com/intl/en/options/">Even more from Google</a></div></div><style>.gb_e .gbqfi::before{left:-56px;top:-35px}.gb_E .gbqfb:focus .gbqfi{outline:1px dotted #fff}@-webkit-keyframes gb__a{0%{opacity:0}50%{opacity:1}}@keyframes gb__a{0%{opacity:0}50%{opacity:1}}#gb#gb a.gb_f,#gb#gb a.gb_g{color:#404040;text-decoration:none}#gb#gb a.gb_g:hover,#gb#gb a.gb_g:focus{color:#000;text-decoration:underline}.gb_h.gb_i{display:none;padding-left:15px;vertical-align:middle}.gb_h.gb_i:first-child{padding-left:0}.gb_j.gb_i{display:inline-block;-webkit-flex:0 1 auto;flex:0 1 auto;-webkit-flex:0 1 main-size;flex:0 1 main-size;display:-webkit-flex;display:flex}.gb_k .gb_j{display:none}.gb_h .gb_g{display:inline-block;line-height:24px;outline:none;vertical-align:middle}.gb_j .gb_g{min-width:60px;overflow:hidden;-webkit-flex:0 1 auto;flex:0 1 auto;-webkit-flex:0 1 main-size;flex:0 1 main-size;text-overflow:ellipsis}.gb_l .gb_j .gb_g{min-width:0}.gb_m .gb_j .gb_g{width:0!important}.gb_n .gb_g{font-weight:bold;text-shadow:0 1px 1px rgba(255,255,255,.9)}.gb_o .gb_g{font-weight:bold;text-shadow:0 1px 1px rgba(0,0,0,.6)}#gb#gb.gb_o a.gb_g{color:#fff}.gb_C .gb_D{background-position:-326px -52px;opacity:.55}.gb_n .gb_C .gb_D{background-position:-97px -57px;opacity:.7}.gb_o .gb_C .gb_D{background-position:-214px 0;opacity:1}.gb_Oc{left:0;min-width:1152px;position:absolute;top:0;-webkit-user-select:none;width:100%}.gb_Vb{font:13px/27px Arial,sans-serif;position:relative;height:60px;width:100%}.gb_ba .gb_Vb{height:28px}#gba{height:60px}#gba.gb_ba{height:28px}#gba.gb_Pc{height:90px}#gba.gb_Pc.gb_ba{height:58px}.gb_Vb>.gb_i{height:60px;line-height:58px;vertical-align:middle}.gb_ba .gb_Vb>.gb_i{height:28px;line-height:26px}.gb_Vb::before{background:#e5e5e5;bottom:0;content:'';display:none;height:1px;left:0;position:absolute;right:0}.gb_Vb{background:#f1f1f1}.gb_Qc .gb_Vb{background:#fff}.gb_Qc .gb_Vb::before,.gb_ba .gb_Vb::before{display:none}.gb_n .gb_Vb,.gb_o .gb_Vb,.gb_ba .gb_Vb{background:transparent}.gb_n .gb_Vb::before{background:#e1e1e1;background:rgba(0,0,0,.12)}.gb_o .gb_Vb::before{background:#333;background:rgba(255,255,255,.2)}.gb_i{display:inline-block;-webkit-flex:0 0 auto;flex:0 0 auto;-webkit-flex:0 0 main-size;flex:0 0 main-size}.gb_i.gb_Rc{float:right;-webkit-order:1;order:1}.gb_Sc{white-space:nowrap;display:-webkit-flex;display:flex;margin-left:0!important;margin-right:0!important}.gb_i{margin-left:0!important;margin-right:0!important}.gb_Ta{background-image:url('//ssl.gstatic.com/gb/images/i1_71651352.png');-webkit-background-size:356px 144px;background-size:356px 144px}@media (min-resolution:1.25dppx),(-webkit-min-device-pixel-ratio:1.25),(min-device-pixel-ratio:1.25){.gb_Ta{background-image:url('//ssl.gstatic.com/gb/images/i2_9ef0f6fa.png')}}.gb_kb{display:inline-block;padding:0 0 0 15px;vertical-align:middle}.gb_kb:first-child,#gbsfw:first-child+.gb_kb{padding-left:0}.gb_Za{position:relative}.gb_D{display:inline-block;outline:none;vertical-align:middle;-webkit-border-radius:2px;border-radius:2px;-webkit-box-sizing:border-box;box-sizing:border-box;height:30px;width:30px}#gb#gb a.gb_D{color:#404040;cursor:default;text-decoration:none}#gb#gb a.gb_D:hover,#gb#gb a.gb_D:focus{color:#000}.gb_ja{border-color:transparent;border-bottom-color:#fff;border-style:dashed dashed solid;border-width:0 8.5px 8.5px;display:none;position:absolute;left:6.5px;top:37px;z-index:1;height:0;width:0;-webkit-animation:gb__a .2s;animation:gb__a .2s}.gb_ka{border-color:transparent;border-style:dashed dashed solid;border-width:0 8.5px 8.5px;display:none;position:absolute;left:6.5px;z-index:1;height:0;width:0;-webkit-animation:gb__a .2s;animation:gb__a .2s;border-bottom-color:#ccc;border-bottom-color:rgba(0,0,0,.2);top:36px}x:-o-prefocus,div.gb_ka{border-bottom-color:#ccc}.gb_H{background:#fff;border:1px solid #ccc;border-color:rgba(0,0,0,.2);-webkit-box-shadow:0 2px 10px rgba(0,0,0,.2);box-shadow:0 2px 10px rgba(0,0,0,.2);display:none;outline:none;overflow:hidden;position:absolute;right:0;top:44px;-webkit-animation:gb__a .2s;animation:gb__a .2s;-webkit-border-radius:2px;border-radius:2px;-webkit-user-select:text}.gb_kb.gb_La .gb_ja,.gb_kb.gb_La .gb_ka,.gb_kb.gb_La .gb_H{display:block}.gb_gc{position:absolute;right:0;top:44px;z-index:-1}.gb_ba .gb_ja,.gb_ba .gb_ka,.gb_ba .gb_H{margin-top:-10px}.gb_7{-webkit-background-size:32px 32px;background-size:32px 32px;-webkit-border-radius:50%;border-radius:50%;display:block;margin:-1px;height:32px;width:32px}.gb_7:hover,.gb_7:focus{-webkit-box-shadow:0 1px 0 rgba(0,0,0,.15);box-shadow:0 1px 0 rgba(0,0,0,.15)}.gb_7:active{-webkit-box-shadow:inset 0 2px 0 rgba(0,0,0,.15);box-shadow:inset 0 2px 0 rgba(0,0,0,.15)}.gb_7:active::after{background:rgba(0,0,0,.1);-webkit-border-radius:50%;border-radius:50%;content:'';display:block;height:100%}.gb_8:not(.gb_e) .gb_7::before,.gb_8:not(.gb_e) .gb_9::before{content:none}.gb_aa{cursor:pointer;line-height:30px;min-width:30px;overflow:hidden;vertical-align:middle;width:auto;text-overflow:ellipsis}.gb_ba .gb_aa,.gb_ba .gb_ca{line-height:26px}#gb#gb.gb_ba a.gb_aa,.gb_ba .gb_ca{color:#666;font-size:11px;height:auto}#gb#gb.gb_ba a.gb_aa:hover,#gb#gb.gb_ba a.gb_aa:focus{color:#000}.gb_da{border-top:4px solid #404040;border-left:4px dashed transparent;border-right:4px dashed transparent;display:inline-block;margin-left:6px;vertical-align:middle}.gb_ba .gb_da{border-top-color:#999}.gb_ea:hover .gb_da{border-top-color:#000}.gb_n .gb_aa{font-weight:bold;text-shadow:0 1px 1px rgba(255,255,255,.9)}.gb_o .gb_aa{font-weight:bold;text-shadow:0 1px 1px rgba(0,0,0,.6)}#gb#gb.gb_o.gb_o a.gb_aa{color:#fff}.gb_o.gb_o .gb_da{border-top-color:#fff}.gb_n .gb_7,.gb_o .gb_7{-webkit-box-shadow:0 1px 2px rgba(0,0,0,.2);box-shadow:0 1px 2px rgba(0,0,0,.2)}.gb_n .gb_7:hover,.gb_o .gb_7:hover,.gb_n .gb_7:focus,.gb_o .gb_7:focus{-webkit-box-shadow:0 1px 0 rgba(0,0,0,.15),0 1px 2px rgba(0,0,0,.2);box-shadow:0 1px 0 rgba(0,0,0,.15),0 1px 2px rgba(0,0,0,.2)}.gb_fa .gb_ga,.gb_ha .gb_ga{position:absolute;right:1px}.gb_ga.gb_i,.gb_ia.gb_i,.gb_ea.gb_i{-webkit-flex:0 1 auto;flex:0 1 auto;-webkit-flex:0 1 main-size;flex:0 1 main-size}.gb_8.gb_m .gb_aa{width:30px!important}.gb_Wb{display:none!important}.gb_e .gb_C .gb_D::before{left:-326px;top:-52px}.gb_e.gb_n .gb_C .gb_D::before{left:-97px;top:-57px}.gb_e.gb_o .gb_C .gb_D::before{left:-214px;top:0}.gb_E .gb_F{position:relative}.gb_e .gb_Ra .gb_Ta::before{left:0;top:-105px}.gb_e.gb_o .gb_Ra .gb_Ta::before{left:-97px;top:-92px}.gb_e.gb_n .gb_Ra .gb_Ta::before{left:-97px;top:0}.gb_e .gb_Ua{background-image:none!important}.gb_e .gb_Va{visibility:visible}.gb_E .gb_Wa span{background:transparent}.gb_e .gb_0a::before{left:-56px;top:0}.gb_e .gb_1a .gb_0a::before{left:-291px;top:-103px}.gb_e.gb_n .gb_D .gb_0a::before{left:-167px;top:-57px}.gb_e.gb_n .gb_1a .gb_0a::before{left:-132px;top:-57px}.gb_e.gb_o .gb_D .gb_0a::before{left:-326px;top:-87px}.gb_e.gb_o .gb_1a .gb_0a::before{left:0;top:-70px}.gb_E .gb_8a{border:1px solid #fff;color:#fff}.gb_E.gb_n .gb_8a{border-color:#000;color:#000}.gb_e .gb_8a.gb_9a::before,.gb_E.gb_e.gb_o .gb_8a.gb_9a::before{left:-214px;top:-117px}.gb_e .gb_8a.gb_ab::before,.gb_E.gb_e.gb_o .gb_8a.gb_ab::before{left:-256px;top:-73px}.gb_e.gb_o .gb_8a.gb_9a::before,.gb_E.gb_e.gb_n .gb_8a.gb_9a::before{left:-326px;top:-122px}.gb_e.gb_o .gb_8a.gb_ab::before,.gb_E.gb_e.gb_n .gb_8a.gb_ab::before{left:-214px;top:-92px}.gb_db{display:none;margin:28px;margin-bottom:-12px;outline:none;position:relative;width:264px;z-index:1;-webkit-border-radius:2px;border-radius:2px;-webkit-box-shadow:0 1px 2px rgba(0,0,0,0.1),0 0 1px rgba(0,0,0,0.1);box-shadow:0 1px 2px rgba(0,0,0,0.1),0 0 1px rgba(0,0,0,0.1)}.gb_db.gb_La{display:block}.gb_cb{-webkit-background-size:64px 64px;background-size:64px 64px;display:inline-block;margin:12px;vertical-align:top;height:64px;width:64px}.gb_fb{display:inline-block;padding:16px 16px 16px 0;vertical-align:top;white-space:normal}.gb_cb~.gb_fb{margin-right:88px}.gb_fb:first-child{padding-left:16px}.gb_gb{color:#262626;font:16px/24px Arial,sans-serif}.gb_hb{color:#737373;font:13px/18px Arial,sans-serif}#gb#gb .gb_db .gb_ib{color:#427fed;text-decoration:none}#gb#gb .gb_db .gb_ib:hover{text-decoration:underline}.gb_db .gb_eb{background-position:-256px 0;cursor:pointer;opacity:.27;outline:none;position:absolute;right:4px;top:4px;height:12px;width:12px}.gb_db .gb_eb:hover{opacity:.55}.gb_kb.gb_lb{padding:0}.gb_lb .gb_H{padding:26px 26px 22px;background:#ffffff}.gb_mb.gb_lb .gb_H{background:#4d90fe}a.gb_nb{color:#666666!important;font-size:22px;height:9px;opacity:.8;position:absolute;right:14px;top:4px;text-decoration:none!important;width:9px}.gb_mb a.gb_nb{color:#c1d1f4!important}a.gb_nb:hover,a.gb_nb:active{opacity:1}.gb_ob{padding:0;width:258px;white-space:normal}.gb_mb .gb_ob{width:200px}.gb_pb{color:#333333;font-size:16px;line-height:20px;margin:0;margin-bottom:16px}.gb_mb .gb_pb{color:#ffffff}.gb_qb{color:#666666;line-height:17px;margin:0;margin-bottom:5px}.gb_mb .gb_qb{color:#ffffff}.gb_rb{position:absolute;background:transparent;top:-999px;z-index:-1;visibility:hidden;margin-top:1px;margin-left:1px}#gb .gb_lb{margin:0}.gb_lb .gb_Z{background:#4d90fe;border-color:#3079ed;margin-top:15px}#gb .gb_lb a.gb_Z.gb_Z{color:#ffffff}.gb_lb .gb_Z:hover{background:#357ae8;border-color:#2f5bb7}.gb_sb .gb_Za .gb_ja{border-bottom-color:#ffffff;display:block}.gb_tb .gb_Za .gb_ja{border-bottom-color:#4d90fe;display:block}.gb_sb .gb_Za .gb_ka,.gb_tb .gb_Za .gb_ka{display:block}.gb_ub{margin-bottom:32px;font-size:small}.gb_ub .gb_vb{margin-right:5px}.gb_ub .gb_wb{color:red}.gb_xb{color:#ffffff;font-size:13px;font-weight:bold;height:25px;line-height:19px;padding-top:5px;padding-left:12px;position:relative;background-color:#4d90fe}.gb_xb .gb_eb{color:#ffffff;cursor:default;font-size:22px;font-weight:normal;position:absolute;right:12px;top:5px}.gb_xb .gb_ib,.gb_xb .gb_yb{color:#ffffff;display:inline-block;font-size:11px;margin-left:16px;padding:0 8px;white-space:nowrap}.gb_zb{background:none;background-image:-webkit-gradient(linear,left top,left bottom,from(rgba(0,0,0,0.16)),to(rgba(0,0,0,0.2)));background-image:-webkit-linear-gradient(top,rgba(0,0,0,0.16),rgba(0,0,0,0.2));background-image:linear-gradient(top,rgba(0,0,0,0.16),rgba(0,0,0,0.2));background-image:-webkit-linear-gradient(top,rgba(0,0,0,0.16),rgba(0,0,0,0.2));border-radius:2px;border:1px solid #dcdcdc;border:1px solid rgba(0,0,0,0.1);cursor:default!important;filter:progid:DXImageTransform.Microsoft.gradient(startColorstr=#160000ff,endColorstr=#220000ff);text-decoration:none!important;-webkit-border-radius:2px}.gb_zb:hover{background:none;background-image:-webkit-gradient(linear,left top,left bottom,from(rgba(0,0,0,0.14)),to(rgba(0,0,0,0.2)));background-image:-webkit-linear-gradient(top,rgba(0,0,0,0.14),rgba(0,0,0,0.2));background-image:linear-gradient(top,rgba(0,0,0,0.14),rgba(0,0,0,0.2));background-image:-webkit-linear-gradient(top,rgba(0,0,0,0.14),rgba(0,0,0,0.2));border:1px solid rgba(0,0,0,0.2);box-shadow:0 1px 1px rgba(0,0,0,0.1);-webkit-box-shadow:0 1px 1px rgba(0,0,0,0.1);filter:progid:DXImageTransform.Microsoft.gradient(startColorstr=#14000000,endColorstr=#22000000)}.gb_zb:active{box-shadow:inset 0 1px 2px rgba(0,0,0,0.3);-webkit-box-shadow:inset 0 1px 2px rgba(0,0,0,0.3)}.gb_jb{display:none}.gb_jb.gb_La{display:block}.gb_e .gb_cb{background-image:none!important}.gb_e .gb_cb::before{display:inline-block;-webkit-transform:scale(.5);transform:scale(.5);-webkit-transform-origin:0 0;transform-origin:0 0}.gb_e .gb_db .gb_eb{position:absolute}.gb_e .gb_db .gb_eb::before{left:-256px;top:0}.gb_e .gb_Rb .gb_D::before{left:-326px;top:-17px}.gb_e.gb_n .gb_Rb .gb_D::before{left:-256px;top:-103px}.gb_e.gb_o .gb_Rb .gb_D::before{left:0;top:-35px}.gb_E .gb_ka{border:0;border-left:1px solid rgba(0,0,0,.2);border-top:1px solid rgba(0,0,0,.2);height:14px;width:14px;-webkit-transform:rotate(45deg);transform:rotate(45deg)}.gb_E .gb_ja{border:0;border-left:1px solid rgba(0,0,0,.2);border-top:1px solid rgba(0,0,0,.2);height:14px;width:14px;-webkit-transform:rotate(45deg);transform:rotate(45deg);border-color:#fff;background:#fff}.gb_e .gb_Jc::before{clip:rect(-0 51px 16px 35px);left:-13px;top:22px}.gb_e .gb_Ta.gb_Kc{position:absolute}.gb_e .gb_Kc::before{clip:rect(17px 307px 33px 291px);left:-261px;top:5px}.gb_e .gb_fa .gb_Jc::before{left:-5px}@media (min-resolution:1.25dppx),(-webkit-min-device-pixel-ratio:1.25),(min-device-pixel-ratio:1.25){.gb_e .gb_Jc::before{clip:rect(-0 102px 32px 70px)}.gb_e .gb_Kc::before{clip:rect(34px 614px 66px 582px)}}.gb_e .gb_Ta,.gb_e .gbii,.gb_e .gbip{background-image:none;overflow:hidden;position:relative}.gb_e .gb_Ta::before{content:url('//ssl.gstatic.com/gb/images/i1_71651352.png');position:absolute}@media (min-resolution:1.25dppx),(-webkit-min-device-pixel-ratio:1.25),(min-device-pixel-ratio:1.25){.gb_e .gb_Ta::before{content:url('//ssl.gstatic.com/gb/images/i2_9ef0f6fa.png');-webkit-transform:scale(.5);transform:scale(.5);-webkit-transform-origin:0 0;transform-origin:0 0}}.gb_E a:focus{outline:1px dotted #fff!important}sentinel{}</style><div id="xjsd"></div><div id="xjsi" data-jiis="bp"><script>(function(){function c(b){window.setTimeout(function(){var a=document.createElement("script");a.src=b;document.getElementById("xjsd").appendChild(a)},0)}google.dljp=function(b,a){google.xjsu=b;c(a)};google.dlj=c;})();(function(){window.google.xjsrm=[];})();if(google.y)google.y.first=[];if(!google.xjs){window._=window._||{};window._._DumpException=function(e){throw e};if(google.timers&&google.timers.load.t){google.timers.load.t.xjsls=new Date().getTime();}google.dljp('/xjs/_/js/k\x3dxjs.s.en_US.pa3Vcej1AaE.O/m\x3dsx,c,sb,cdos,cr,elog,jsa,r,hsm,j,p,d,csi/am\x3dtEj9FoOhLNQJRB0/rt\x3dj/d\x3d1/t\x3dzcms/rs\x3dACT90oFZ5cL6iK6m-9VNpsnbwx7NUExpXA','/xjs/_/js/k\x3dxjs.s.en_US.pa3Vcej1AaE.O/m\x3dsx,c,sb,cdos,cr,elog,jsa,r,hsm,j,p,d,csi/am\x3dtEj9FoOhLNQJRB0/rt\x3dj/d\x3d1/t\x3dzcms/rs\x3dACT90oFZ5cL6iK6m-9VNpsnbwx7NUExpXA');google.xjs=1;}google.pmc={"sx":{},"c":{"mcr":5},"sb":{"agen":false,"cgen":true,"client":"hp","dh":true,"ds":"","exp":"msedr","fl":true,"host":"google.com","jam":1,"msgs":{"che":"Not listening. Something went wrong.","ched":"Help","cher":"Restart listening","chh":"Say \"Ok Google\"","cht":"Hotword detection is off.","chtr":"Start listening for \"Ok Google\"","chtt":"Listening for \"Ok Google\"","cibl":"Clear Search","dym":"Did you mean:","lcky":"I\u0026#39;m Feeling Lucky","lml":"Learn more","oskt":"Input tools","psrc":"This search was removed from your \u003Ca href=\"/history\"\u003EWeb History\u003C/a\u003E","psrl":"Remove","sbit":"Search by image","srch":"Google Search","srtt":"Search by voice"},"ovr":{},"pq":"","psy":"p","refoq":true,"scd":10,"sce":4,"sre":true,"stok":"ENYNYoJCJJ33PbVhODiCsFEB9eg"},"abd":{"abd":false,"dabp":false,"deb":false,"der":false,"det":false,"psa":false,"sup":false},"aldd":{},"async":{},"cdos":{},"cr":{"eup":false,"qir":false,"rctj":true,"ref":false,"uff":false},"ddls":{},"elog":{},"erh":{},"foot":{"pf":true,"po":false,"qe":false},"fpe":{"js":true},"gf":{"pid":196},"hv":{},"idck":{},"ifl":{"opts":[{"href":"/url?url=/doodles/alan-turings-100th-birthday","id":"doodley","msg":"I'm Feeling Doodley"},{"href":"/url?url=http://www.googleartproject.com/collection/uffizi-gallery/\u0026sa=t\u0026usg=AFQjCNHrqY5KQKLuYFHyqcpEEimlUCs0wA","id":"artistic","msg":"I'm Feeling Artistic"},{"href":"/url?url=/search?gws_rd%3Dssl%26q%3Drestaurants","id":"hungry","msg":"I'm Feeling Hungry"},{"href":"/url?url=http://www.agoogleaday.com/%23date%3D04-22-2011\u0026sa=t\u0026usg=AFQjCNFRPF1FMy8HKec5BWfHfJNjXo8J2g","id":"puzzled","msg":"I'm Feeling Puzzled"},{"href":"/url?url=/trends/hottrends","id":"trendy","msg":"I'm Feeling Trendy"},{"href":"/url?url=/earth/explore/showcase/hubble20th.html%23tab%3Dngc-1300","id":"stellar","msg":"I'm Feeling Stellar"},{"href":"/url?url=/doodles/30th-anniversary-of-pac-man","id":"playful","msg":"I'm Feeling Playful"},{"href":"/url?url=/culturalinstitute/entity/%252Fm%252F0jz631_?hl%3Den%26projectId%3Dworld-wonders","id":"wonderful","msg":"I'm Feeling Wonderful"},{"href":"/url?url=/onetoday/","id":"generous","msg":"I'm Feeling Generous"}]},"jsa":{},"jsaleg":{},"lc":{},"llc":{},"lu":{},"m":{"ab":{"on":true},"ajax":{"gl":"us","hl":"en","q":""},"css":{"showTopNav":true},"exp":{"lrb":true,"tnav":true},"msgs":{"details":"Result details","hPers":"Hide private results","hPersD":"Currently hiding private results","loading":"Still loading...","mute":"Mute","noPreview":"Preview not available","sPers":"Show all results","sPersD":"Currently showing private results","unmute":"Unmute"},"time":{"hUnit":1500}},"r":{},"rk":{"bl":"Feedback","db":"Reported","di":"Thank you.","dl":"Report another problem","efe":true,"rb":"Wrong?","ri":"Please report the problem.","rl":"Cancel"},"rkab":{},"rmcl":{"bl":"Feedback","db":"Reported","di":"Thank you.","dl":"Report another problem","rb":"Wrong?","ri":"Please report the problem.","rl":"Cancel"},"sf":{},"spch":{"ae":"Please check your microphone.  \u003Ca href=\"https://support.google.com/chrome/?p=ui_voice_search\" target=\"_blank\"\u003ELearn more\u003C/a\u003E","hen":true,"hl":"en-US","htt":"Listening for \"Ok Google\"","im":"Click \u003Cb\u003EAllow\u003C/b\u003E to start voice search","iw":"Waiting...","lm":"Listening...","lu":"%1$s voice search not available","ne":"No Internet connection","nt":"Didn't get that. \u003Cspan\u003ETry again\u003C/span\u003E","nv":"Please check your microphone and audio levels.  \u003Ca href=\"https://support.google.com/chrome/?p=ui_voice_search\" target=\"_blank\"\u003ELearn more\u003C/a\u003E","pe":"Voice search has been turned off.  \u003Ca href=\"https://support.google.com/chrome/?p=ui_voice_search\" target=\"_blank\"\u003EDetails\u003C/a\u003E","rm":"Speak now"},"st":{},"vm":{"bv":85142067,"d":"cGU","tc":true,"te":true,"tk":true,"ts":true},"hsm":{},"j":{"ajrp":true,"cmt":true,"ftwd":200,"icmt":false,"lbtfdr":10000,"lcuwl":true,"mcr":5,"miml":true,"scmt":true,"sirs":"clone","tct":" \\u3000?","tlh":true,"ufl":true,"witu":false},"p":{"ae":true,"avgTtfc":2000,"brba":false,"dlen":24,"dper":3,"eae":true,"fbdc":500,"fbdu":-1,"fbh":true,"fd":1000000,"focus":true,"gpsj":true,"hiue":true,"hpt":310,"iavgTtfc":2000,"knrt":true,"maxCbt":1500,"mds":"prc,sp,mbl_he,mbl_hs,mbl_re,mbl_rs,mbl_sv","msg":{"dym":"Did you mean:","gs":"Google Search","kntt":"Use the up and down arrow keys to select each result. Press Enter to go to the selection.","pcnt":"New Tab","sif":"Search instead for","srf":"Showing results for"},"nprr":1,"ohpt":false,"ophe":true,"pmt":250,"pq":true,"rpt":15,"sc":"psy-ab","tdur":50},"d":{},"csi":{"acsi":true},"oAaJvQ":{"oq":""},"TG8rFw":{"sd":"1"},"hLaaFQ":{"ed":"Please enter a description.","eu":"Please enter a valid URL."},"SpiLtA":{},"q1cupA":{},"8AF24Q":{},"bnhGTQ":{},"4RZUyg":{},"/nNC3A":{},"ITl3wQ":{},"l6UIRQ":{},"FmbnUA":{},"c+PT4g":{},"/1S6iw":{},"GqeGtQ":{},"+idT0Q":{},"NpA8BQ":{},"ADSBcg":{},"BwDLOw":{},"8aqNqA":{},"A/Ucpg":{},"63G1zA":{},"cm4D8w":{},"GfrcvQ":{}};google.y.first.push(function(){google.loadAll(['abd','async','erh','foot','fpe','hv','idck','ifl','lc','lu','m','sf','spch','vm'].concat(google.xjsrm||[]));if(google.med){google.med('init');google.initHistory();google.med('history');}});if(google.j&&google.j.en&&google.j.xi){window.setTimeout(google.j.xi,0);}
</script></div><script>(function(){if(google.timers&&google.timers.load.t){var f=function(){google.timers.load.t&&(google.timers.load.t.ol=(new Date).getTime(),google.timers.load.t.iml=b,google.kCSI.imc=c,google.kCSI.imn=d,google.kCSI.imp=e,google.stt&&(google.kCSI.stt=google.stt),google.csiReport&&google.csiReport())},k=function(a){b=(new Date).getTime();++c;a=a||window.event;a=a.target||a.srcElement;h(a,k);google.iml(a,b);},h=function(a,
g){a.removeEventListener?(a.removeEventListener("load",g,!1),a.removeEventListener("error",g,!1)):(a.detachEvent("onload",g),a.detachEvent("onerror",g))},d,c,e,b;var l=document.getElementsByTagName("img");d=l.length;for(var m=c=0,n;m<d;++m){n=l[m];var p="string"!=typeof n.src||!n.src,q=p||n.complete;p&&n.getAttribute("data-bsrc")&&(q=!1);q?++c:n.addEventListener?(n.addEventListener("load",k,!1),n.addEventListener("error",k,!1)):(n.attachEvent("onload",k),n.attachEvent("onerror",k))}e=d-c;window.addEventListener?window.addEventListener("load",f,!1):window.attachEvent&&
window.attachEvent("onload",f);google.timers.load.t.prt=b=google.time()};})();
</script></div></body></html>
`
