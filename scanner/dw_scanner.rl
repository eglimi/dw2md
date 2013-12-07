// This file is automatically generated by ragel. Do not change.
// A tested version of ragel is 6.8.
// Build with `ragel -o scanner/dw_scanner.go -Z -G0 scanner/dw_scanner.rl`

package scanner

var (
	ts, te, act int
)

func parseDoc(data []byte, c chan item) {
	
	cs, p, pe, eof := 0, 0, len(data), len(data)

%%{
	machine dw2md;

	code_start = '<code'(space [a-zA-Z]+)?'>';
	file_start = '<file'(space [a-zA-Z.]+){0,2}'>';

	main := |*
	'='{6}     => { c <-item{ts, p, "======", "#",       itemHeading1}  };
	'='{5}     => { c <-item{ts, p, "=====", "##",       itemHeading2}  };
	'='{4}     => { c <-item{ts, p, "====", "###",       itemHeading3}  };
	'='{3}     => { c <-item{ts, p, "===", "####",       itemHeading4}  };
	'='{2}     => { c <-item{ts, p, "==", "#####",       itemHeading5}  };
	'**'       => { c <-item{ts, p, "**", "**",          itemBold}      };
	'//'       => { c <-item{ts, p, "//", "*",           itemItalic}    };
	'__'       => { c <-item{ts, p, "__", "",            itemUnderline} };
	'\'\''     => { c <-item{ts, p, "''", "`",           itemMonospace} };
	#'[['       => { c <-item{ts, p, "[[", "",            itemStartLink} };
	#']]'       => { c <-item{ts, p, "]]", "",            itemEndLink}   };
	#'{{'       => { c <-item{ts, p, "{{", "",            itemStartPic}  };
	#'}}'       => { c <-item{ts, p, "}}", "",            itemEndPic}    };
	code_start => { c <-item{ts, p, "<code>", "```",     itemStartCode} };
	'</code>'  => { c <-item{ts, p, "</code>", "```",    itemEndCode}   };
	file_start => { c <-item{ts, p, "<file>", "```",     itemStartFile} };
	'</file>'  => { c <-item{ts, p, "</file>", "```",    itemEndFile}   };
	any        => { c <-item{ts, p, string(data[p]), "", itemText}      };
	*|;

	write data;
	write init;
	write exec;

}%%

	c <-item{0,0,"","",itemEOF}
}