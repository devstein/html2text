package html2text_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/devstein/html2text"
)

func TestHTML2Text(t *testing.T) {
	Convey("html2text.HTML2Text should work", t, func() {

		Convey("Links", func() {
			So(html2text.HTML2Text(`<div></div>`), ShouldEqual, "")
			So(html2text.HTML2Text(`<div>simple text</div>`), ShouldEqual, "simple text")
			So(html2text.HTML2Text(`click <a href="test">here</a>`), ShouldEqual, "click here <test>")
			So(html2text.HTML2Text(`click <a class="x" href="test">here</a>`), ShouldEqual, "click here <test>")
			So(html2text.HTML2Text(`click <a href="ents/&apos;x&apos;">here</a>`), ShouldEqual, "click here <ents/'x'>")
			So(html2text.HTML2Text(`click <a href="javascript:void(0)">here</a>`), ShouldEqual, "click here")
			So(html2text.HTML2Text(`click <a href="test"><span>here</span> or here</a>`), ShouldEqual, "click here or here <test>")
			So(html2text.HTML2Text(`click <a href="http://bit.ly/2n4wXRs">news</a>`), ShouldEqual, "click news <http://bit.ly/2n4wXRs>")
			So(html2text.HTML2Text(`<a rel="mw:WikiLink" href="/wiki/yet#English" title="yet">yet</a>, <a rel="mw:WikiLink" href="/wiki/not_yet#English" title="not yet">not yet</a>`), ShouldEqual, "yet </wiki/yet#English>, not yet </wiki/not_yet#English>")
			So(html2text.HTML2Text(`click <a href="one">here<a href="two"> or</a><span> here</span></a>`), ShouldEqual, "click here or <one> here <two>")
		})

		Convey("Inlines", func() {
			So(html2text.HTML2Text(`strong <strong>text</strong>`), ShouldEqual, "strong text")
			So(html2text.HTML2Text(`some <div id="a" class="b">div</div>`), ShouldEqual, "some div")
		})

		Convey("Line breaks and spaces", func() {
			So(html2text.HTML2Text("should    ignore more spaces"), ShouldEqual, "should ignore more spaces")
			So(html2text.HTML2Text("should \nignore \r\nnew lines"), ShouldEqual, "should ignore new lines")
			So(html2text.HTML2Text("a\nb\nc"), ShouldEqual, "a b c")
			So(html2text.HTML2Text(`two<br>line<br/>breaks`), ShouldEqual, "two\r\nline\r\nbreaks")
			So(html2text.HTML2Text(`<p>two</p><p>paragraphs</p>`), ShouldEqual, "two\r\n\r\nparagraphs")
		})

		Convey("Headings", func() {
			So(html2text.HTML2Text("<h1>First</h1>main text"), ShouldEqual, "First\r\n\r\nmain text")
			So(html2text.HTML2Text("First<h2>Second</h2>next section"), ShouldEqual, "First\r\n\r\nSecond\r\n\r\nnext section")
			So(html2text.HTML2Text("<h2>Second</h2>next section"), ShouldEqual, "Second\r\n\r\nnext section")
			So(html2text.HTML2Text("Second<h3>Third</h3>next section"), ShouldEqual, "Second\r\n\r\nThird\r\n\r\nnext section")
			So(html2text.HTML2Text("<h3>Third</h3>next section"), ShouldEqual, "Third\r\n\r\nnext section")
			So(html2text.HTML2Text("Third<h4>Fourth</h4>next section"), ShouldEqual, "Third\r\n\r\nFourth\r\n\r\nnext section")
			So(html2text.HTML2Text("<h4>Fourth</h4>next section"), ShouldEqual, "Fourth\r\n\r\nnext section")
			So(html2text.HTML2Text("Fourth<h5>Fifth</h5>next section"), ShouldEqual, "Fourth\r\n\r\nFifth\r\n\r\nnext section")
			So(html2text.HTML2Text("<h5>Fifth</h5>next section"), ShouldEqual, "Fifth\r\n\r\nnext section")
			So(html2text.HTML2Text("Fifth<h6>Sixth</h6>next section"), ShouldEqual, "Fifth\r\n\r\nSixth\r\n\r\nnext section")
			So(html2text.HTML2Text("<h6>Sixth</h6>next section"), ShouldEqual, "Sixth\r\n\r\nnext section")
			So(html2text.HTML2Text("<h7>Not Header</h7>next section"), ShouldEqual, "Not Headernext section")
		})

		Convey("HTML entities", func() {
			So(html2text.HTML2Text(`two&nbsp;&nbsp;spaces`), ShouldEqual, "two  spaces")
			So(html2text.HTML2Text(`&copy; 2017 K3A`), ShouldEqual, "© 2017 K3A")
			So(html2text.HTML2Text("&lt;printtag&gt;"), ShouldEqual, "<printtag>")
			So(html2text.HTML2Text(`would you pay in &cent;, &pound;, &yen; or &euro;?`),
				ShouldEqual, "would you pay in ¢, £, ¥ or €?")
			So(html2text.HTML2Text(`Tom & Jerry is not an entity`), ShouldEqual, "Tom & Jerry is not an entity")
			So(html2text.HTML2Text(`this &neither; as you see`), ShouldEqual, "this &neither; as you see")
			So(html2text.HTML2Text(`list of items<ul><li>One</li><li>Two</li><li>Three</li></ul>`), ShouldEqual, "list of items\r\nOne\r\nTwo\r\nThree\r\n")
			So(html2text.HTML2Text(`fish &amp; chips`), ShouldEqual, "fish & chips")
			So(html2text.HTML2Text(`&quot;I'm sorry, Dave. I'm afraid I can't do that.&quot; – HAL, 2001: A Space Odyssey`), ShouldEqual, "\"I'm sorry, Dave. I'm afraid I can't do that.\" – HAL, 2001: A Space Odyssey")
			So(html2text.HTML2Text(`Google &reg;`), ShouldEqual, "Google ®")
			So(html2text.HTML2Text(`&#8268; decimal and hex entities supported &#x204D;`), ShouldEqual, "⁌ decimal and hex entities supported ⁍")
		})

		Convey("Large Entity", func() {
			So(html2text.HTMLEntitiesToText("&abcdefghij;"), ShouldEqual, "&abcdefghij;")
		})

		Convey("Numeric HTML Entities", func() {
			So(html2text.HTMLEntitiesToText("&#39;single quotes&#39; and &#52765;"), ShouldEqual, "'single quotes' and 츝")
		})

		Convey("Full HTML structure", func() {
			So(html2text.HTML2Text(``), ShouldEqual, "")
			So(html2text.HTML2Text(`<html><head><title>Good</title></head><body>x</body>`), ShouldEqual, "x")
			So(html2text.HTML2Text(`we are not <script type="javascript"></script>interested in scripts`),
				ShouldEqual, "we are not interested in scripts")
		})

		Convey("Switching Unix and Windows line breaks", func() {
			html2text.SetUnixLbr(true)
			So(html2text.HTML2Text(`two<br>line<br/>breaks`), ShouldEqual, "two\nline\nbreaks")
			So(html2text.HTML2Text(`<p>two</p><p>paragraphs</p>`), ShouldEqual, "two\n\nparagraphs")
			html2text.SetUnixLbr(false)
			So(html2text.HTML2Text(`two<br>line<br/>breaks`), ShouldEqual, "two\r\nline\r\nbreaks")
			So(html2text.HTML2Text(`<p>two</p><p>paragraphs</p>`), ShouldEqual, "two\r\n\r\nparagraphs")
		})

		Convey("Custom HTML Tags", func() {
			So(html2text.HTML2Text(`<aa>hello</aa>`), ShouldEqual, "hello")
			So(html2text.HTML2Text(`<aa >hello</aa>`), ShouldEqual, "hello")
			So(html2text.HTML2Text(`<aa x="1">hello</aa>`), ShouldEqual, "hello")
		})

	})
}
