// Copyright 2019 Huawei Technologies Co.,Ltd.
// Licensed under the Apache License, Version 2.0 (the "License"); you may not use
// this file except in compliance with the License.  You may obtain a copy of the
// License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed
// under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
// CONDITIONS OF ANY KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations under the License.

package obs

var mimeTypes = map[string]string{
	"001":     "application/x-001",
	"301":     "application/x-301",
	"323":     "text/h323",
	"7z":      "application/x-7z-compressed",
	"906":     "application/x-906",
	"907":     "drawing/907",
	"IVF":     "video/x-ivf",
	"a11":     "application/x-a11",
	"aac":     "audio/x-aac",
	"acp":     "audio/x-mei-aac",
	"ai":      "application/postscript",
	"aif":     "audio/aiff",
	"aifc":    "audio/aiff",
	"aiff":    "audio/aiff",
	"anv":     "application/x-anv",
	"apk":     "application/vnd.android.package-archive",
	"asa":     "text/asa",
	"asf":     "video/x-ms-asf",
	"asp":     "text/asp",
	"asx":     "video/x-ms-asf",
	"atom":    "application/atom+xml",
	"au":      "audio/basic",
	"avi":     "video/avi",
	"awf":     "application/vnd.adobe.workflow",
	"biz":     "text/xml",
	"bmp":     "application/x-bmp",
	"bot":     "application/x-bot",
	"bz2":     "application/x-bzip2",
	"c4t":     "application/x-c4t",
	"c90":     "application/x-c90",
	"cal":     "application/x-cals",
	"cat":     "application/vnd.ms-pki.seccat",
	"cdf":     "application/x-netcdf",
	"cdr":     "application/x-cdr",
	"cel":     "application/x-cel",
	"cer":     "application/x-x509-ca-cert",
	"cg4":     "application/x-g4",
	"cgm":     "application/x-cgm",
	"cit":     "application/x-cit",
	"class":   "java/*",
	"cml":     "text/xml",
	"cmp":     "application/x-cmp",
	"cmx":     "application/x-cmx",
	"cot":     "application/x-cot",
	"crl":     "application/pkix-crl",
	"crt":     "application/x-x509-ca-cert",
	"csi":     "application/x-csi",
	"css":     "text/css",
	"csv":     "text/csv",
	"cu":      "application/cu-seeme",
	"cut":     "application/x-cut",
	"dbf":     "application/x-dbf",
	"dbm":     "application/x-dbm",
	"dbx":     "application/x-dbx",
	"dcd":     "text/xml",
	"dcx":     "application/x-dcx",
	"deb":     "application/x-debian-package",
	"der":     "application/x-x509-ca-cert",
	"dgn":     "application/x-dgn",
	"dib":     "application/x-dib",
	"dll":     "application/x-msdownload",
	"doc":     "application/msword",
	"docx":    "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
	"dot":     "application/msword",
	"drw":     "application/x-drw",
	"dtd":     "text/xml",
	"dvi":     "application/x-dvi",
	"dwf":     "application/x-dwf",
	"dwg":     "application/x-dwg",
	"dxb":     "application/x-dxb",
	"dxf":     "application/x-dxf",
	"edn":     "application/vnd.adobe.edn",
	"emf":     "application/x-emf",
	"eml":     "message/rfc822",
	"ent":     "text/xml",
	"eot":     "application/vnd.ms-fontobject",
	"epi":     "application/x-epi",
	"eps":     "application/postscript",
	"epub":    "application/epub+zip",
	"etd":     "application/x-ebx",
	"etx":     "text/x-setext",
	"exe":     "application/x-msdownload",
	"fax":     "image/fax",
	"fdf":     "application/vnd.fdf",
	"fif":     "application/fractals",
	"flac":    "audio/flac",
	"flv":     "video/x-flv",
	"fo":      "text/xml",
	"frm":     "application/x-frm",
	"g4":      "application/x-g4",
	"gbr":     "application/x-gbr",
	"gif":     "image/gif",
	"gl2":     "application/x-gl2",
	"gp4":     "application/x-gp4",
	"gz":      "application/gzip",
	"hgl":     "application/x-hgl",
	"hmr":     "application/x-hmr",
	"hpg":     "application/x-hpgl",
	"hpl":     "application/x-hpl",
	"hqx":     "application/mac-binhex40",
	"hrf":     "application/x-hrf",
	"hta":     "application/hta",
	"htc":     "text/x-component",
	"htm":     "text/html",
	"html":    "text/html",
	"htt":     "text/webviewhtml",
	"htx":     "text/html",
	"icb":     "application/x-icb",
	"ico":     "application/x-ico",
	"ics":     "text/calendar",
	"iff":     "application/x-iff",
	"ig4":     "application/x-g4",
	"igs":     "application/x-igs",
	"iii":     "application/x-iphone",
	"img":     "application/x-img",
	"ini":     "text/plain",
	"ins":     "application/x-internet-signup",
	"ipa":     "application/vnd.iphone",
	"iso":     "application/x-iso9660-image",
	"isp":     "application/x-internet-signup",
	"jar":     "application/java-archive",
	"java":    "java/*",
	"jfif":    "image/jpeg",
	"jpe":     "image/jpeg",
	"jpeg":    "image/jpeg",
	"jpg":     "image/jpeg",
	"js":      "application/x-javascript",
	"json":    "application/json",
	"jsp":     "text/html",
	"la1":     "audio/x-liquid-file",
	"lar":     "application/x-laplayer-reg",
	"latex":   "application/x-latex",
	"lavs":    "audio/x-liquid-secure",
	"lbm":     "application/x-lbm",
	"lmsff":   "audio/x-la-lms",
	"log":     "text/plain",
	"ls":      "application/x-javascript",
	"ltr":     "application/x-ltr",
	"m1v":     "video/x-mpeg",
	"m2v":     "video/x-mpeg",
	"m3u":     "audio/mpegurl",
	"m4a":     "audio/mp4",
	"m4e":     "video/mpeg4",
	"m4v":     "video/mp4",
	"mac":     "application/x-mac",
	"man":     "application/x-troff-man",
	"math":    "text/xml",
	"mdb":     "application/msaccess",
	"mfp":     "application/x-shockwave-flash",
	"mht":     "message/rfc822",
	"mhtml":   "message/rfc822",
	"mi":      "application/x-mi",
	"mid":     "audio/mid",
	"midi":    "audio/mid",
	"mil":     "application/x-mil",
	"mml":     "text/xml",
	"mnd":     "audio/x-musicnet-download",
	"mns":     "audio/x-musicnet-stream",
	"mocha":   "application/x-javascript",
	"mov":     "video/quicktime",
	"movie":   "video/x-sgi-movie",
	"mp1":     "audio/mp1",
	"mp2":     "audio/mp2",
	"mp2v":    "video/mpeg",
	"mp3":     "audio/mp3",
	"mp4":     "video/mp4",
	"mp4a":    "audio/mp4",
	"mp4v":    "video/mp4",
	"mpa":     "video/x-mpg",
	"mpd":     "application/vnd.ms-project",
	"mpe":     "video/mpeg",
	"mpeg":    "video/mpeg",
	"mpg":     "video/mpeg",
	"mpg4":    "video/mp4",
	"mpga":    "audio/rn-mpeg",
	"mpp":     "application/vnd.ms-project",
	"mps":     "video/x-mpeg",
	"mpt":     "application/vnd.ms-project",
	"mpv":     "video/mpg",
	"mpv2":    "video/mpeg",
	"mpw":     "application/vnd.ms-project",
	"mpx":     "application/vnd.ms-project",
	"mtx":     "text/xml",
	"mxp":     "application/x-mmxp",
	"net":     "image/pnetvue",
	"nrf":     "application/x-nrf",
	"nws":     "message/rfc822",
	"odc":     "text/x-ms-odc",
	"oga":     "audio/ogg",
	"ogg":     "audio/ogg",
	"ogv":     "video/ogg",
	"ogx":     "application/ogg",
	"out":     "application/x-out",
	"p10":     "application/pkcs10",
	"p12":     "application/x-pkcs12",
	"p7b":     "application/x-pkcs7-certificates",
	"p7c":     "application/pkcs7-mime",
	"p7m":     "application/pkcs7-mime",
	"p7r":     "application/x-pkcs7-certreqresp",
	"p7s":     "application/pkcs7-signature",
	"pbm":     "image/x-portable-bitmap",
	"pc5":     "application/x-pc5",
	"pci":     "application/x-pci",
	"pcl":     "application/x-pcl",
	"pcx":     "application/x-pcx",
	"pdf":     "application/pdf",
	"pdx":     "application/vnd.adobe.pdx",
	"pfx":     "application/x-pkcs12",
	"pgl":     "application/x-pgl",
	"pgm":     "image/x-portable-graymap",
	"pic":     "application/x-pic",
	"pko":     "application/vnd.ms-pki.pko",
	"pl":      "application/x-perl",
	"plg":     "text/html",
	"pls":     "audio/scpls",
	"plt":     "application/x-plt",
	"png":     "image/png",
	"pnm":     "image/x-portable-anymap",
	"pot":     "application/vnd.ms-powerpoint",
	"ppa":     "application/vnd.ms-powerpoint",
	"ppm":     "application/x-ppm",
	"pps":     "application/vnd.ms-powerpoint",
	"ppt":     "application/vnd.ms-powerpoint",
	"pptx":    "application/vnd.openxmlformats-officedocument.presentationml.presentation",
	"pr":      "application/x-pr",
	"prf":     "application/pics-rules",
	"prn":     "application/x-prn",
	"prt":     "application/x-prt",
	"ps":      "application/postscript",
	"ptn":     "application/x-ptn",
	"pwz":     "application/vnd.ms-powerpoint",
	"qt":      "video/quicktime",
	"r3t":     "text/vnd.rn-realtext3d",
	"ra":      "audio/vnd.rn-realaudio",
	"ram":     "audio/x-pn-realaudio",
	"rar":     "application/x-rar-compressed",
	"ras":     "application/x-ras",
	"rat":     "application/rat-file",
	"rdf":     "text/xml",
	"rec":     "application/vnd.rn-recording",
	"red":     "application/x-red",
	"rgb":     "application/x-rgb",
	"rjs":     "application/vnd.rn-realsystem-rjs",
	"rjt":     "application/vnd.rn-realsystem-rjt",
	"rlc":     "application/x-rlc",
	"rle":     "application/x-rle",
	"rm":      "application/vnd.rn-realmedia",
	"rmf":     "application/vnd.adobe.rmf",
	"rmi":     "audio/mid",
	"rmj":     "application/vnd.rn-realsystem-rmj",
	"rmm":     "audio/x-pn-realaudio",
	"rmp":     "application/vnd.rn-rn_music_package",
	"rms":     "application/vnd.rn-realmedia-secure",
	"rmvb":    "application/vnd.rn-realmedia-vbr",
	"rmx":     "application/vnd.rn-realsystem-rmx",
	"rnx":     "application/vnd.rn-realplayer",
	"rp":      "image/vnd.rn-realpix",
	"rpm":     "audio/x-pn-realaudio-plugin",
	"rsml":    "application/vnd.rn-rsml",
	"rss":     "application/rss+xml",
	"rt":      "text/vnd.rn-realtext",
	"rtf":     "application/x-rtf",
	"rv":      "video/vnd.rn-realvideo",
	"sam":     "application/x-sam",
	"sat":     "application/x-sat",
	"sdp":     "application/sdp",
	"sdw":     "application/x-sdw",
	"sgm":     "text/sgml",
	"sgml":    "text/sgml",
	"sis":     "application/vnd.symbian.install",
	"sisx":    "application/vnd.symbian.install",
	"sit":     "application/x-stuffit",
	"slb":     "application/x-slb",
	"sld":     "application/x-sld",
	"slk":     "drawing/x-slk",
	"smi":     "application/smil",
	"smil":    "application/smil",
	"smk":     "application/x-smk",
	"snd":     "audio/basic",
	"sol":     "text/plain",
	"sor":     "text/plain",
	"spc":     "application/x-pkcs7-certificates",
	"spl":     "application/futuresplash",
	"spp":     "text/xml",
	"ssm":     "application/streamingmedia",
	"sst":     "application/vnd.ms-pki.certstore",
	"stl":     "application/vnd.ms-pki.stl",
	"stm":     "text/html",
	"sty":     "application/x-sty",
	"svg":     "image/svg+xml",
	"swf":     "application/x-shockwave-flash",
	"tar":     "application/x-tar",
	"tdf":     "application/x-tdf",
	"tg4":     "application/x-tg4",
	"tga":     "application/x-tga",
	"tif":     "image/tiff",
	"tiff":    "image/tiff",
	"tld":     "text/xml",
	"top":     "drawing/x-top",
	"torrent": "application/x-bittorrent",
	"tsd":     "text/xml",
	"ttf":     "application/x-font-ttf",
	"txt":     "text/plain",
	"uin":     "application/x-icq",
	"uls":     "text/iuls",
	"vcf":     "text/x-vcard",
	"vda":     "application/x-vda",
	"vdx":     "application/vnd.visio",
	"vml":     "text/xml",
	"vpg":     "application/x-vpeg005",
	"vsd":     "application/vnd.visio",
	"vss":     "application/vnd.visio",
	"vst":     "application/x-vst",
	"vsw":     "application/vnd.visio",
	"vsx":     "application/vnd.visio",
	"vtx":     "application/vnd.visio",
	"vxml":    "text/xml",
	"wav":     "audio/wav",
	"wax":     "audio/x-ms-wax",
	"wb1":     "application/x-wb1",
	"wb2":     "application/x-wb2",
	"wb3":     "application/x-wb3",
	"wbmp":    "image/vnd.wap.wbmp",
	"webm":    "video/webm",
	"wiz":     "application/msword",
	"wk3":     "application/x-wk3",
	"wk4":     "application/x-wk4",
	"wkq":     "application/x-wkq",
	"wks":     "application/x-wks",
	"wm":      "video/x-ms-wm",
	"wma":     "audio/x-ms-wma",
	"wmd":     "application/x-ms-wmd",
	"wmf":     "application/x-wmf",
	"wml":     "text/vnd.wap.wml",
	"wmv":     "video/x-ms-wmv",
	"wmx":     "video/x-ms-wmx",
	"wmz":     "application/x-ms-wmz",
	"woff":    "application/x-font-woff",
	"wp6":     "application/x-wp6",
	"wpd":     "application/x-wpd",
	"wpg":     "application/x-wpg",
	"wpl":     "application/vnd.ms-wpl",
	"wq1":     "application/x-wq1",
	"wr1":     "application/x-wr1",
	"wri":     "application/x-wri",
	"wrk":     "application/x-wrk",
	"ws":      "application/x-ws",
	"ws2":     "application/x-ws",
	"wsc":     "text/scriptlet",
	"wsdl":    "text/xml",
	"wvx":     "video/x-ms-wvx",
	"x_b":     "application/x-x_b",
	"x_t":     "application/x-x_t",
	"xap":     "application/x-silverlight-app",
	"xbm":     "image/x-xbitmap",
	"xdp":     "application/vnd.adobe.xdp",
	"xdr":     "text/xml",
	"xfd":     "application/vnd.adobe.xfd",
	"xfdf":    "application/vnd.adobe.xfdf",
	"xhtml":   "text/html",
	"xls":     "application/vnd.ms-excel",
	"xlsx":    "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
	"xlw":     "application/x-xlw",
	"xml":     "text/xml",
	"xpl":     "audio/scpls",
	"xpm":     "image/x-xpixmap",
	"xq":      "text/xml",
	"xql":     "text/xml",
	"xquery":  "text/xml",
	"xsd":     "text/xml",
	"xsl":     "text/xml",
	"xslt":    "text/xml",
	"xwd":     "application/x-xwd",
	"yaml":    "text/yaml",
	"yml":     "text/yaml",
	"zip":     "application/zip",
	"dotx":    "application/vnd.openxmlformats-officedocument.wordprocessingml.template",
	"wps":     "application/vnd.ms-works",
	"wpt":     "x-lml/x-gps",
	"pptm":    "application/vnd.ms-powerpoint.presentation.macroenabled.12",
	"heic":    "image/heic",
	"mkv":     "video/x-matroska",
	"raw":     "image/x-panasonic-raw",
}
