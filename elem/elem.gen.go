//go:generate go run generate.go

// Package elem defines markup to create DOM elements.
//
// Generated from "HTML element reference" by Mozilla Contributors,
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element, licensed under
// CC-BY-SA 2.5.
package elem

import "github.com/gopherjs/vecty"

// Anchor (or anchor element) creates a hyperlink to other web pages, files,
// locations within the same page, email addresses, or any other URL.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/a
func Anchor(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("a", markup...)
}

// The HTML Abbreviation element (<abbr>) represents an abbreviation or
// acronym; the optional title attribute can provide an expansion or
// description for the abbreviation.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/abbr
func Abbreviation(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("abbr", markup...)
}

// Address indicates that the enclosed HTML provides contact information for a
// person or people, or for an organization.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/address
func Address(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("address", markup...)
}

// Area defines a hot-spot region on an image, and optionally associates it
// with a hypertext link. This element is used only within a <map> element.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/area
func Area(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("area", markup...)
}

// Article represents a self-contained composition in a document, page,
// application, or site, which is intended to be independently distributable or
// reusable (e.g., in syndication). Examples include: a forum post, a magazine
// or newspaper article, or a blog entry.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/article
func Article(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("article", markup...)
}

// Aside represents a portion of a document whose content is only indirectly
// related to the document's main content.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/aside
func Aside(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("aside", markup...)
}

//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/audio
func Audio(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("audio", markup...)
}

// The HTML Bring Attention To element (<b>) is used to draw the reader's
// attention to the element's contents, which are not otherwise granted special
// importance.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/b
func Bold(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("b", markup...)
}

// Base specifies the base URL to use for all relative URLs contained within a
// document. There can be only one <base> element in a document.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/base
func Base(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("base", markup...)
}

// The HTML BiDirectional Isolation element (<bdi>) is used to indicate spans
// of text which might need to be rendered in the opposite direction than the
// surrounding text.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/bdi
func BidirectionalIsolation(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("bdi", markup...)
}

// The HTML Bidirectional Text Override element (<bdo>) overrides the current
// directionality of text, so that the text within is rendered in a different
// direction.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/bdo
func BidirectionalOverride(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("bdo", markup...)
}

// BlockQuote (or HTML Block Quotation Element) indicates that the enclosed
// text is an extended quotation. Usually, this is rendered visually by
// indentation (see Notes for how to change it). A URL for the source of the
// quotation may be given using the cite attribute, while a text representation
// of the source can be given using the <cite> element.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/blockquote
func BlockQuote(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("blockquote", markup...)
}

// Body represents the content of an HTML document. There can be only one
// <body> element in a document.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/body
func Body(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("body", markup...)
}

// Break produces a line break in text (carriage-return). It is useful for
// writing a poem or an address, where the division of lines is significant.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/br
func Break(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("br", markup...)
}

// Button represents a clickable button, which can be used in forms, or
// anywhere in a document that needs simple, standard button functionality.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/button
func Button(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("button", markup...)
}

// Use the HTML <canvas> element with either the canvas scripting API or the
// WebGL API to draw graphics and animations.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/canvas
func Canvas(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("canvas", markup...)
}

// The HTML Table Caption element (<caption>) specifies the caption (or title)
// of a table, and if used is always the first child of a <table>.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/caption
func Caption(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("caption", markup...)
}

// The HTML Citation element (<cite>) is used to describe a reference to a
// cited creative work, and must include either the title or the URL of that
// work.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/cite
func Citation(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("cite", markup...)
}

// Code displays its contents styled in a fashion intended to indicate that the
// text is a short fragment of computer code.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/code
func Code(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("code", markup...)
}

// Column defines a column within a table and is used for defining common
// semantics on all common cells. It is generally found within a <colgroup>
// element.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/col
func Column(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("col", markup...)
}

// ColumnGroup defines a group of columns within a table.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/colgroup
func ColumnGroup(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("colgroup", markup...)
}

// Data links a given content with a machine-readable translation. If the
// content is time- or date-related, the <time> element must be used.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/data
func Data(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("data", markup...)
}

// DataList contains a set of <option> elements that represent the values
// available for other controls.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/datalist
func DataList(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("datalist", markup...)
}

// Description provides the details about or the definition of the preceding
// term (<dt>) in a description list (<dl>).
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/dd
func Description(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("dd", markup...)
}

// DeletedText represents a range of text that has been deleted from a
// document.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/del
func DeletedText(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("del", markup...)
}

// The HTML Details Element (<details>) creates a disclosure widget in which
// information is visible only when the widget is toggled into an "open" state.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/details
func Details(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("details", markup...)
}

// The HTML Definition element (<dfn>) is used to indicate the term being
// defined within the context of a definition phrase or sentence.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/dfn
func Definition(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("dfn", markup...)
}

// Dialog represents a dialog box or other interactive component, such as an
// inspector or window.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/dialog
func Dialog(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("dialog", markup...)
}

// The HTML Content Division element (<div>) is the generic container for flow
// content. It has no effect on the content or layout until styled using CSS.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/div
func Div(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("div", markup...)
}

// DescriptionList represents a description list. The element encloses a list
// of groups of terms (specified using the <dt> element) and descriptions
// (provided by <dd> elements). Common uses for this element are to implement a
// glossary or to display metadata (a list of key-value pairs).
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/dl
func DescriptionList(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("dl", markup...)
}

// DefinitionTerm specifies a term in a description or definition list, and as
// such must be used inside a <dl> element.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/dt
func DefinitionTerm(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("dt", markup...)
}

// Emphasis marks text that has stress emphasis. The <em> element can be
// nested, with each level of nesting indicating a greater degree of emphasis.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/em
func Emphasis(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("em", markup...)
}

// Embed embeds external content at the specified point in the document. This
// content is provided by an external application or other source of
// interactive content such as a browser plug-in.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/embed
func Embed(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("embed", markup...)
}

// FieldSet is used to group several controls as well as labels (<label>)
// within a web form.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/fieldset
func FieldSet(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("fieldset", markup...)
}

// FigureCaption represents a caption or a legend associated with a figure or
// an illustration described by the rest of the data of the <figure> element
// which is its immediate ancestor.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/figcaption
func FigureCaption(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("figcaption", markup...)
}

// Figure represents self-contained content, frequently with a caption
// (<figcaption>), and is typically referenced as a single unit.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/figure
func Figure(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("figure", markup...)
}

// Footer represents a footer for its nearest sectioning content or sectioning
// root element. A footer typically contains information about the author of
// the section, copyright data or links to related documents.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/footer
func Footer(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("footer", markup...)
}

// Form represents a document section that contains interactive controls for
// submitting information to a web server.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/form
func Form(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("form", markup...)
}

// The HTML <h1>–<h6> elements represent six levels of section headings. <h1>
// is the highest section level and <h6> is the lowest.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Heading_Elements
func Heading1(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("h1", markup...)
}

// The HTML <h1>–<h6> elements represent six levels of section headings. <h1>
// is the highest section level and <h6> is the lowest.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Heading_Elements
func Heading2(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("h2", markup...)
}

// The HTML <h1>–<h6> elements represent six levels of section headings. <h1>
// is the highest section level and <h6> is the lowest.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Heading_Elements
func Heading3(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("h3", markup...)
}

// The HTML <h1>–<h6> elements represent six levels of section headings. <h1>
// is the highest section level and <h6> is the lowest.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Heading_Elements
func Heading4(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("h4", markup...)
}

// The HTML <h1>–<h6> elements represent six levels of section headings. <h1>
// is the highest section level and <h6> is the lowest.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Heading_Elements
func Heading5(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("h5", markup...)
}

// The HTML <h1>–<h6> elements represent six levels of section headings. <h1>
// is the highest section level and <h6> is the lowest.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Heading_Elements
func Heading6(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("h6", markup...)
}

// Header represents introductory content, typically a group of introductory or
// navigational aids. It may contain some heading elements but also other
// elements like a logo, a search form, an author name, and so on.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/header
func Header(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("header", markup...)
}

// HeadingsGroup represents a multi-level heading for a section of a document.
// It groups a set of <h1>–<h6> elements.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/hgroup
func HeadingsGroup(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("hgroup", markup...)
}

// HorizontalRule represents a thematic break between paragraph-level elements
// (for example, a change of scene in a story, or a shift of topic with a
// section); historically, this has been presented as a horizontal rule or
// line.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/hr
func HorizontalRule(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("hr", markup...)
}

// Italic represents a range of text that is set off from the normal text for
// some reason. Some examples include technical terms, foreign language
// phrases, or fictional character thoughts. It is typically displayed in
// italic type.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/i
func Italic(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("i", markup...)
}

// The HTML Inline Frame element (<iframe>) represents a nested browsing
// context, effectively embedding another HTML page into the current page.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/iframe
func InlineFrame(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("iframe", markup...)
}

// Image embeds an image into the document.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/img
func Image(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("img", markup...)
}

// Input is used to create interactive controls for web-based forms in order to
// accept data from the user.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input
func Input(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("input", markup...)
}

// InsertedText represents a range of text that has been added to a document.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/ins
func InsertedText(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("ins", markup...)
}

// The HTML Keyboard Input element (<kbd>) represents a span of inline text
// denoting textual user input from a keyboard, voice input, or any other text
// entry device.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/kbd
func KeyboardInput(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("kbd", markup...)
}

// Label represents a caption for an item in a user interface.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/label
func Label(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("label", markup...)
}

// Legend represents a caption for the content of its parent <fieldset>.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/legend
func Legend(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("legend", markup...)
}

// ListItem is used to represent an item in a list. It must be contained in a
// parent element: an ordered list (<ol>), an unordered list (<ul>), or a menu
// (<menu>). In menus and unordered lists, list items are usually displayed
// using bullet points. In ordered lists, they are usually displayed with an
// ascending counter on the left, such as a number or letter.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/li
func ListItem(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("li", markup...)
}

// Link specifies relationships between the current document and an external
// resource. Possible uses for this element include defining a relational
// framework for navigation. This element is most used to link to style sheets.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/link
func Link(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("link", markup...)
}

// Main represents the dominant content of the <body> of a document, portion of
// a document or application. The main content area consists of content that is
// directly related to or expands upon the central topic of a document, or the
// central functionality of an application.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/main
func Main(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("main", markup...)
}

// Map is used with <area> elements to define an image map (a clickable link
// area).
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/map
func Map(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("map", markup...)
}

// The HTML Mark Text element (<mark>) represents text which is marked or
// highlighted for reference or notation purposes, due to the marked passage's
// relevance or importance in the enclosing context.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/mark
func Mark(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("mark", markup...)
}

// Menu represents a group of commands that a user can perform or activate.
// This includes both list menus, which might appear across the top of a
// screen, as well as context menus, such as those that might appear underneath
// a button after it has been clicked.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/menu
func Menu(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("menu", markup...)
}

// Meta represents metadata that cannot be represented by other HTML
// meta-related elements, like <base>, <link>, <script>, <style> or <title>.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/meta
func Meta(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("meta", markup...)
}

// Meter represents either a scalar value within a known range or a fractional
// value.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/meter
func Meter(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("meter", markup...)
}

// Navigation represents a section of a page whose purpose is to provide
// navigation links, either within the current document or to other documents.
// Common examples of navigation sections are menus, tables of contents, and
// indexes.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/nav
func Navigation(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("nav", markup...)
}

// NoScript defines a section of HTML to be inserted if a script type on the
// page is unsupported or if scripting is currently turned off in the browser.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/noscript
func NoScript(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("noscript", markup...)
}

// Object represents an external resource, which can be treated as an image, a
// nested browsing context, or a resource to be handled by a plugin.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/object
func Object(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("object", markup...)
}

// OrderedList represents an ordered list of items, typically rendered as a
// numbered list.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/ol
func OrderedList(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("ol", markup...)
}

// OptionsGroup creates a grouping of options within a <select> element.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/optgroup
func OptionsGroup(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("optgroup", markup...)
}

// Option is used to define an item contained in a <select>, an <optgroup>, or
// a <datalist> element. As such, <option> can represent menu items in popups
// and other lists of items in an HTML document.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/option
func Option(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("option", markup...)
}

// The HTML Output element (<output>) is a container element into which a site
// or app can inject the results of a calculation or the outcome of a user
// action.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/output
func Output(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("output", markup...)
}

// Paragraph represents a paragraph of text.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/p
func Paragraph(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("p", markup...)
}

// Parameter defines parameters for an <object> element.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/param
func Parameter(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("param", markup...)
}

// Picture serves as a container for zero or more <source> elements and one
// <img> element to provide versions of an image for different display device
// scenarios.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/picture
func Picture(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("picture", markup...)
}

// Preformatted represents preformatted text which is to be presented exactly
// as written in the HTML file.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/pre
func Preformatted(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("pre", markup...)
}

// Progress displays an indicator showing the completion progress of a task,
// typically displayed as a progress bar.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/progress
func Progress(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("progress", markup...)
}

// Quote indicates that the enclosed text is a short inline quotation. Most
// modern browsers implement this by surrounding the text in quotation marks.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/q
func Quote(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("q", markup...)
}

// The HTML Ruby Fallback Parenthesis (<rp>) element is used to provide
// fall-back parentheses for browsers that do not support display of ruby
// annotations using the <ruby> element.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/rp
func RubyParenthesis(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("rp", markup...)
}

// The HTML Ruby Text (<rt>) element specifies the ruby text component of a
// ruby annotation, which is used to provide pronunciation, translation, or
// transliteration information for East Asian typography. The <rt> element must
// always be contained within a <ruby> element.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/rt
func RubyText(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("rt", markup...)
}

// The HTML Ruby Text Container (<rtc>) element embraces semantic annotations
// of characters presented in a ruby of <rb> elements used inside of <ruby>
// element. <rb> elements can have both pronunciation (<rt>) and semantic
// (<rtc>) annotations.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/rtc
func RubyTextContainer(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("rtc", markup...)
}

// Ruby represents a ruby annotation. Ruby annotations are for showing
// pronunciation of East Asian characters.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/ruby
func Ruby(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("ruby", markup...)
}

// Strikethrough renders text with a strikethrough, or a line through it. Use
// the <s> element to represent things that are no longer relevant or no longer
// accurate. However, <s> is not appropriate when indicating document edits;
// for that, use the <del> and <ins> elements, as appropriate.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/s
func Strikethrough(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("s", markup...)
}

// The HTML Sample Element (<samp>) is used to enclose inline text which
// represents sample (or quoted) output from a computer program.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/samp
func Sample(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("samp", markup...)
}

// Script is used to embed or reference executable code; this is typically used
// to embed or refer to JavaScript code.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/script
func Script(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("script", markup...)
}

// Section represents a standalone section — which doesn't have a more
// specific semantic element to represent it — contained within an HTML
// document.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/section
func Section(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("section", markup...)
}

// Select represents a control that provides a menu of options:
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/select
func Select(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("select", markup...)
}

// Slot—part of the Web Components technology suite—is a placeholder inside
// a web component that you can fill with your own markup, which lets you
// create separate DOM trees and present them together.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/slot
func Slot(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("slot", markup...)
}

// Small makes the text font size one size smaller (for example, from large to
// medium, or from small to x-small) down to the browser's minimum font size.
// In HTML5, this element is repurposed to represent side-comments and small
// print, including copyright and legal text, independent of its styled
// presentation.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/small
func Small(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("small", markup...)
}

// Source specifies multiple media resources for the <picture>, the <audio>
// element, or the <video> element. It is an empty element. It is commonly used
// to serve the same media content in multiple formats supported by different
// browsers.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/source
func Source(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("source", markup...)
}

// Span is a generic inline container for phrasing content, which does not
// inherently represent anything. It can be used to group elements for styling
// purposes (using the class or id attributes), or because they share attribute
// values, such as lang.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/span
func Span(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("span", markup...)
}

// The HTML Strong Importance Element (<strong>) indicates that its contents
// have strong importance, seriousness, or urgency. Browsers typically render
// the contents in bold type.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/strong
func Strong(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("strong", markup...)
}

// Style contains style information for a document, or part of a document.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/style
func Style(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("style", markup...)
}

// The HTML Subscript element (<sub>) specifies inline text which should be
// displayed as subscript for solely typographical reasons.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/sub
func Subscript(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("sub", markup...)
}

// The HTML Disclosure Summary element (<summary>) element specifies a summary,
// caption, or legend for a <details> element's disclosure box.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/summary
func Summary(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("summary", markup...)
}

// The HTML Superscript element (<sup>) specifies inline text which is to be
// displayed as superscript for solely typographical reasons.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/sup
func Superscript(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("sup", markup...)
}

// Table represents tabular data — that is, information presented in a
// two-dimensional table comprised of rows and columns of cells containing
// data.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/table
func Table(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("table", markup...)
}

// The HTML Table Body element (<tbody>) encapsulates a set of table row (<tr>
// elements, indicating that they comprise the body of the table (<table>).
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/tbody
func TableBody(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("tbody", markup...)
}

// TableData defines a cell of a table that contains data. It participates in
// the table model.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/td
func TableData(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("td", markup...)
}

// The HTML Content Template (<template>) element is a mechanism for holding
// client-side content that is not to be rendered when a page is loaded but may
// subsequently be instantiated during runtime using JavaScript.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/template
func Template(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("template", markup...)
}

// TextArea represents a multi-line plain-text editing control.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/textarea
func TextArea(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("textarea", markup...)
}

// TableFoot defines a set of rows summarizing the columns of the table.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/tfoot
func TableFoot(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("tfoot", markup...)
}

// TableHeader defines a cell as header of a group of table cells. The exact
// nature of this group is defined by the scope and headers attributes.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/th
func TableHeader(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("th", markup...)
}

// TableHead defines a set of rows defining the head of the columns of the
// table.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/thead
func TableHead(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("thead", markup...)
}

// Time represents a specific period in time. It may include the datetime
// attribute to translate dates into machine-readable format, allowing for
// better search engine results or custom features such as reminders.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/time
func Time(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("time", markup...)
}

// The HTML Title element (<title>) defines the title of the document, shown in
// a browser's title bar or on the page's tab.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/title
func Title(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("title", markup...)
}

// TableRow defines a row of cells in a table. The row's cells can then be
// established using a mix of <td> (data cell) and <th> (header cell) elements.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/tr
func TableRow(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("tr", markup...)
}

// Track is used as a child of the media elements <audio> and <video>. It lets
// you specify timed text tracks (or time-based data), for example to
// automatically handle subtitles. The tracks are formatted in WebVTT format
// (.vtt files) — Web Video Text Tracks.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/track
func Track(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("track", markup...)
}

// The HTML Unarticulated Annotation element (<u>) represents a span of inline
// text which should be rendered in a way that indicates that it has a
// non-textual annotation.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/u
func Underline(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("u", markup...)
}

// UnorderedList represents an unordered list of items, typically rendered as a
// bulleted list.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/ul
func UnorderedList(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("ul", markup...)
}

// The HTML Variable element (<var>) represents the name of a variable in a
// mathematical expression or a programming context.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/var
func Variable(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("var", markup...)
}

// The HTML Video element (<video>) embeds a media player which supports video
// playback into the document.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/video
func Video(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("video", markup...)
}

// WordBreakOpportunity represents a word break opportunity—a position within
// text where the browser may optionally break a line, though its line-breaking
// rules would not otherwise create a break at that location.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/wbr
func WordBreakOpportunity(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("wbr", markup...)
}

// <input> elements of type "button" are rendered as simple push buttons, which
// can be programmed to control custom functionality anywhere on a webpage as
// required when assigned an event handler function (typically for the click
// event).
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/button
func Input type="button"(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("input type="button"", markup...)
}

// <input> elements of type checkbox are rendered by default as square boxes
// that are checked (ticked) when activated, like you might see in an official
// government paper form. They allow you to select single values for submission
// in a form (or not).
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/checkbox
func Input type="checkbox"(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("input type="checkbox"", markup...)
}

// <input> elements of type "color" provide a user interface element that lets
// a user specify a color, either by using a visual color picker interface or
// by entering the color into a text field in "#rrggbb" hexadecimal format.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/color
func Input type="color"(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("input type="color"", markup...)
}

// <input> elements of type date create input fields that let the user enter a
// date, either using a text box that automatically validates the content, or
// using a special date picker interface. The resulting value includes the
// year, month, and day, but not the time. The time and datetime-local input
// types support time and date/time inputs.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/date
func Input type="date"(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("input type="date"", markup...)
}

// <input> elements of type datetime-local create input controls that let the
// user easily enter both a date and a time, including the year, month, and day
// as well as the time in hours and minutes. The user's local time zone is
// used.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/datetime-local
func Input type="datetime-local"(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("input type="datetime-local"", markup...)
}

// <input> elements of type "email" are used to let the user enter and edit an
// email address, or, if the multiple attribute is specified, a list of email
// addresses. The input value is automatically validated to ensure that it's
// either empty or a properly-formatted email address (or list of addresses)
// before the form can be submitted.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/email
func Input type="email"(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("input type="email"", markup...)
}

// <input> elements with type="file" let the user choose one or more files from
// their device storage. Once chosen, the files can be uploaded to a server
// using form submission, or manipulated using JavaScript code and the File
// API.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/file
func Input type="file"(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("input type="file"", markup...)
}

// <input> elements of type "hidden" let web developers include data that
// cannot be seen or modified by users when a form is submitted. For example,
// the ID of the content that is currently being ordered or edited, or a unique
// security token. Hidden inputs are completely invisible in the rendered page,
// and there is no way to make it visible in the page's content.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/hidden
func Input type="hidden"(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("input type="hidden"", markup...)
}

// <input> elements of type "image" are used to create graphical submit
// buttons, i.e. submit buttons that take the form of an image rather than
// text.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/image
func Input type="image"(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("input type="image"", markup...)
}

// <input> elements of type month create input fields that let the user enter a
// month and year allowing a month and year to be easily entered. The value is
// a string whose value is in the format "YYYY-MM", where YYYY is the
// four-digit year and MM is the month number.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/month
func Input type="month"(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("input type="month"", markup...)
}

// <input> elements of type "number" are used to let the user enter a number.
// They include built-in validation to reject non-numerical entries.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/number
func Input type="number"(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("input type="number"", markup...)
}

// <input> elements of type "password" provide a way for the user to securely
// enter a password.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/password
func Input type="password"(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("input type="password"", markup...)
}

// <input> elements of type radio are generally used in radio
// groups—collections of radio buttons describing a set of related options.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/radio
func Input type="radio"(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("input type="radio"", markup...)
}

// <input> elements of type "range" let the user specify a numeric value which
// must be no less than a given value, and no more than another given value.
// The precise value, however, is not considered important. This is typically
// represented using a slider or dial control rather than a text entry box like
// the "number" input type.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/range
func Input type="range"(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("input type="range"", markup...)
}

// <input> elements of type "reset" are rendered as buttons, with a default
// click event handler that resets all of the inputs in the form to their
// initial values.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/reset
func Input type="reset"(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("input type="reset"", markup...)
}

// <input> elements of type "search" are text fields designed for the user to
// enter search queries into. These are functionally identical to text inputs,
// but may be styled differently by the user agent.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/search
func Input type="search"(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("input type="search"", markup...)
}

// <input> elements of type "submit" are rendered as buttons. When the click
// event occurs (typically because the user clicked the button), the user agent
// attempts to submit the form to the server.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/submit
func Input type="submit"(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("input type="submit"", markup...)
}

// <input> elements of type "tel" are used to let the user enter and edit a
// telephone number. Unlike <input type="email"> and <input type="url"> , the
// input value is not automatically validated to a particular format before the
// form can be submitted, because formats for telephone numbers vary so much
// around the world.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/tel
func Input type="tel"(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("input type="tel"", markup...)
}

// <input> elements of type "text" create basic single-line text fields.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/text
func Input type="text"(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("input type="text"", markup...)
}

// <input> elements of type time create input fields designed to let the user
// easily enter a time (hours and minutes, and optionally seconds).
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/time
func Input type="time"(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("input type="time"", markup...)
}

// <input> elements of type "url" are used to let the user enter and edit a
// URL. The input value is automatically validated to ensure that it's either
// empty or a properly-formatted URL before the form can be submitted.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/url
func Input type="url"(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("input type="url"", markup...)
}

// <input> elements of type "week" create input fields allowing easy entry of a
// year plus the ISO 8601 week number during that year (i.e., week 1 to 52 or
// 53).
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/week
func Input type="week"(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("input type="week"", markup...)
}
