package config

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/breml/logstash-config/ast"
)

var g = &grammar{
	rules: []*rule{
		{
			name: "config",
			pos:  position{line: 10, col: 1, offset: 252},
			expr: &actionExpr{
				pos: position{line: 11, col: 5, offset: 265},
				run: (*parser).callonconfig1,
				expr: &seqExpr{
					pos: position{line: 11, col: 5, offset: 265},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 11, col: 5, offset: 265},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 11, col: 7, offset: 267},
							label: "ps",
							expr: &ruleRefExpr{
								pos:  position{line: 11, col: 10, offset: 270},
								name: "plugin_section",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 11, col: 25, offset: 285},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 11, col: 27, offset: 287},
							label: "pss",
							expr: &zeroOrMoreExpr{
								pos: position{line: 11, col: 31, offset: 291},
								expr: &actionExpr{
									pos: position{line: 12, col: 9, offset: 301},
									run: (*parser).callonconfig9,
									expr: &seqExpr{
										pos: position{line: 12, col: 9, offset: 301},
										exprs: []interface{}{
											&ruleRefExpr{
												pos:  position{line: 12, col: 9, offset: 301},
												name: "_",
											},
											&labeledExpr{
												pos:   position{line: 12, col: 11, offset: 303},
												label: "ps",
												expr: &ruleRefExpr{
													pos:  position{line: 12, col: 14, offset: 306},
													name: "plugin_section",
												},
											},
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 15, col: 8, offset: 369},
							name: "_",
						},
						&ruleRefExpr{
							pos:  position{line: 15, col: 10, offset: 371},
							name: "EOF",
						},
					},
				},
			},
		},
		{
			name: "comment",
			pos:  position{line: 23, col: 1, offset: 521},
			expr: &oneOrMoreExpr{
				pos: position{line: 24, col: 5, offset: 535},
				expr: &seqExpr{
					pos: position{line: 24, col: 6, offset: 536},
					exprs: []interface{}{
						&zeroOrOneExpr{
							pos: position{line: 24, col: 6, offset: 536},
							expr: &ruleRefExpr{
								pos:  position{line: 24, col: 6, offset: 536},
								name: "whitespace",
							},
						},
						&litMatcher{
							pos:        position{line: 24, col: 18, offset: 548},
							val:        "#",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 24, col: 22, offset: 552},
							expr: &charClassMatcher{
								pos:        position{line: 24, col: 22, offset: 552},
								val:        "[^\\r\\n]",
								chars:      []rune{'\r', '\n'},
								ignoreCase: false,
								inverted:   true,
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 24, col: 31, offset: 561},
							expr: &litMatcher{
								pos:        position{line: 24, col: 31, offset: 561},
								val:        "\r",
								ignoreCase: false,
							},
						},
						&litMatcher{
							pos:        position{line: 24, col: 37, offset: 567},
							val:        "\n",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "_",
			pos:  position{line: 30, col: 1, offset: 661},
			expr: &zeroOrMoreExpr{
				pos: position{line: 31, col: 5, offset: 669},
				expr: &choiceExpr{
					pos: position{line: 31, col: 6, offset: 670},
					alternatives: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 31, col: 6, offset: 670},
							name: "comment",
						},
						&ruleRefExpr{
							pos:  position{line: 31, col: 16, offset: 680},
							name: "whitespace",
						},
					},
				},
			},
		},
		{
			name: "whitespace",
			pos:  position{line: 37, col: 1, offset: 776},
			expr: &oneOrMoreExpr{
				pos: position{line: 38, col: 5, offset: 793},
				expr: &charClassMatcher{
					pos:        position{line: 38, col: 5, offset: 793},
					val:        "[ \\t\\r\\n]",
					chars:      []rune{' ', '\t', '\r', '\n'},
					ignoreCase: false,
					inverted:   false,
				},
			},
		},
		{
			name: "plugin_section",
			pos:  position{line: 47, col: 1, offset: 950},
			expr: &actionExpr{
				pos: position{line: 48, col: 5, offset: 971},
				run: (*parser).callonplugin_section1,
				expr: &seqExpr{
					pos: position{line: 48, col: 5, offset: 971},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 48, col: 5, offset: 971},
							label: "pt",
							expr: &ruleRefExpr{
								pos:  position{line: 48, col: 8, offset: 974},
								name: "plugin_type",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 48, col: 20, offset: 986},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 48, col: 22, offset: 988},
							val:        "{",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 48, col: 26, offset: 992},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 48, col: 28, offset: 994},
							label: "bops",
							expr: &zeroOrMoreExpr{
								pos: position{line: 48, col: 33, offset: 999},
								expr: &actionExpr{
									pos: position{line: 49, col: 9, offset: 1009},
									run: (*parser).callonplugin_section10,
									expr: &seqExpr{
										pos: position{line: 49, col: 9, offset: 1009},
										exprs: []interface{}{
											&labeledExpr{
												pos:   position{line: 49, col: 9, offset: 1009},
												label: "bop",
												expr: &ruleRefExpr{
													pos:  position{line: 49, col: 13, offset: 1013},
													name: "branch_or_plugin",
												},
											},
											&ruleRefExpr{
												pos:  position{line: 49, col: 30, offset: 1030},
												name: "_",
											},
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 52, col: 8, offset: 1082},
							val:        "}",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "branch_or_plugin",
			pos:  position{line: 62, col: 1, offset: 1240},
			expr: &ruleRefExpr{
				pos:  position{line: 63, col: 5, offset: 1263},
				name: "plugin",
			},
		},
		{
			name: "plugin_type",
			pos:  position{line: 69, col: 1, offset: 1342},
			expr: &choiceExpr{
				pos: position{line: 70, col: 5, offset: 1360},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 70, col: 5, offset: 1360},
						run: (*parser).callonplugin_type2,
						expr: &litMatcher{
							pos:        position{line: 70, col: 5, offset: 1360},
							val:        "input",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 72, col: 9, offset: 1408},
						run: (*parser).callonplugin_type4,
						expr: &litMatcher{
							pos:        position{line: 72, col: 9, offset: 1408},
							val:        "filter",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 74, col: 9, offset: 1458},
						run: (*parser).callonplugin_type6,
						expr: &litMatcher{
							pos:        position{line: 74, col: 9, offset: 1458},
							val:        "output",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "plugin",
			pos:  position{line: 105, col: 1, offset: 2062},
			expr: &actionExpr{
				pos: position{line: 106, col: 5, offset: 2075},
				run: (*parser).callonplugin1,
				expr: &seqExpr{
					pos: position{line: 106, col: 5, offset: 2075},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 106, col: 5, offset: 2075},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 106, col: 10, offset: 2080},
								name: "name",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 106, col: 15, offset: 2085},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 106, col: 17, offset: 2087},
							val:        "{",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 106, col: 21, offset: 2091},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 106, col: 23, offset: 2093},
							label: "attributes",
							expr: &zeroOrOneExpr{
								pos: position{line: 106, col: 34, offset: 2104},
								expr: &actionExpr{
									pos: position{line: 107, col: 9, offset: 2115},
									run: (*parser).callonplugin10,
									expr: &seqExpr{
										pos: position{line: 107, col: 9, offset: 2115},
										exprs: []interface{}{
											&labeledExpr{
												pos:   position{line: 107, col: 9, offset: 2115},
												label: "attribute",
												expr: &ruleRefExpr{
													pos:  position{line: 107, col: 19, offset: 2125},
													name: "attribute",
												},
											},
											&labeledExpr{
												pos:   position{line: 107, col: 29, offset: 2135},
												label: "attrs",
												expr: &zeroOrMoreExpr{
													pos: position{line: 107, col: 35, offset: 2141},
													expr: &actionExpr{
														pos: position{line: 108, col: 13, offset: 2155},
														run: (*parser).callonplugin16,
														expr: &seqExpr{
															pos: position{line: 108, col: 13, offset: 2155},
															exprs: []interface{}{
																&ruleRefExpr{
																	pos:  position{line: 108, col: 13, offset: 2155},
																	name: "whitespace",
																},
																&ruleRefExpr{
																	pos:  position{line: 108, col: 24, offset: 2166},
																	name: "_",
																},
																&labeledExpr{
																	pos:   position{line: 108, col: 26, offset: 2168},
																	label: "attribute",
																	expr: &ruleRefExpr{
																		pos:  position{line: 108, col: 36, offset: 2178},
																		name: "attribute",
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 114, col: 8, offset: 2320},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 114, col: 10, offset: 2322},
							val:        "}",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "name",
			pos:  position{line: 125, col: 1, offset: 2492},
			expr: &choiceExpr{
				pos: position{line: 126, col: 7, offset: 2505},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 126, col: 7, offset: 2505},
						run: (*parser).callonname2,
						expr: &oneOrMoreExpr{
							pos: position{line: 126, col: 8, offset: 2506},
							expr: &charClassMatcher{
								pos:        position{line: 126, col: 8, offset: 2506},
								val:        "[A-Za-z0-9_-]",
								chars:      []rune{'_', '-'},
								ranges:     []rune{'A', 'Z', 'a', 'z', '0', '9'},
								ignoreCase: false,
								inverted:   false,
							},
						},
					},
					&actionExpr{
						pos: position{line: 128, col: 9, offset: 2567},
						run: (*parser).callonname5,
						expr: &labeledExpr{
							pos:   position{line: 128, col: 9, offset: 2567},
							label: "value",
							expr: &ruleRefExpr{
								pos:  position{line: 128, col: 15, offset: 2573},
								name: "string_value",
							},
						},
					},
				},
			},
		},
		{
			name: "attribute",
			pos:  position{line: 137, col: 1, offset: 2722},
			expr: &actionExpr{
				pos: position{line: 138, col: 5, offset: 2738},
				run: (*parser).callonattribute1,
				expr: &seqExpr{
					pos: position{line: 138, col: 5, offset: 2738},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 138, col: 5, offset: 2738},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 138, col: 10, offset: 2743},
								name: "name",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 138, col: 15, offset: 2748},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 138, col: 17, offset: 2750},
							val:        "=>",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 138, col: 22, offset: 2755},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 138, col: 24, offset: 2757},
							label: "value",
							expr: &ruleRefExpr{
								pos:  position{line: 138, col: 30, offset: 2763},
								name: "value",
							},
						},
					},
				},
			},
		},
		{
			name: "value",
			pos:  position{line: 146, col: 1, offset: 2900},
			expr: &choiceExpr{
				pos: position{line: 147, col: 5, offset: 2912},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 147, col: 5, offset: 2912},
						name: "plugin",
					},
					&ruleRefExpr{
						pos:  position{line: 147, col: 14, offset: 2921},
						name: "bareword",
					},
					&ruleRefExpr{
						pos:  position{line: 147, col: 25, offset: 2932},
						name: "string_value",
					},
					&ruleRefExpr{
						pos:  position{line: 147, col: 40, offset: 2947},
						name: "number",
					},
					&ruleRefExpr{
						pos:  position{line: 147, col: 49, offset: 2956},
						name: "array",
					},
					&ruleRefExpr{
						pos:  position{line: 147, col: 57, offset: 2964},
						name: "hash",
					},
				},
			},
		},
		{
			name: "array_value",
			pos:  position{line: 153, col: 1, offset: 3051},
			expr: &choiceExpr{
				pos: position{line: 154, col: 5, offset: 3069},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 154, col: 5, offset: 3069},
						name: "bareword",
					},
					&ruleRefExpr{
						pos:  position{line: 154, col: 16, offset: 3080},
						name: "string_value",
					},
					&ruleRefExpr{
						pos:  position{line: 154, col: 31, offset: 3095},
						name: "number",
					},
					&ruleRefExpr{
						pos:  position{line: 154, col: 40, offset: 3104},
						name: "array",
					},
					&ruleRefExpr{
						pos:  position{line: 154, col: 48, offset: 3112},
						name: "hash",
					},
				},
			},
		},
		{
			name: "bareword",
			pos:  position{line: 161, col: 1, offset: 3219},
			expr: &actionExpr{
				pos: position{line: 162, col: 5, offset: 3234},
				run: (*parser).callonbareword1,
				expr: &seqExpr{
					pos: position{line: 162, col: 5, offset: 3234},
					exprs: []interface{}{
						&charClassMatcher{
							pos:        position{line: 162, col: 5, offset: 3234},
							val:        "[A-Za-z_]",
							chars:      []rune{'_'},
							ranges:     []rune{'A', 'Z', 'a', 'z'},
							ignoreCase: false,
							inverted:   false,
						},
						&oneOrMoreExpr{
							pos: position{line: 162, col: 15, offset: 3244},
							expr: &charClassMatcher{
								pos:        position{line: 162, col: 15, offset: 3244},
								val:        "[A-Za-z0-9_]",
								chars:      []rune{'_'},
								ranges:     []rune{'A', 'Z', 'a', 'z', '0', '9'},
								ignoreCase: false,
								inverted:   false,
							},
						},
					},
				},
			},
		},
		{
			name: "double_quoted_string",
			pos:  position{line: 170, col: 1, offset: 3454},
			expr: &actionExpr{
				pos: position{line: 171, col: 5, offset: 3481},
				run: (*parser).callondouble_quoted_string1,
				expr: &seqExpr{
					pos: position{line: 171, col: 7, offset: 3483},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 171, col: 7, offset: 3483},
							val:        "\"",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 171, col: 11, offset: 3487},
							expr: &choiceExpr{
								pos: position{line: 171, col: 13, offset: 3489},
								alternatives: []interface{}{
									&litMatcher{
										pos:        position{line: 171, col: 13, offset: 3489},
										val:        "\\\"",
										ignoreCase: false,
									},
									&seqExpr{
										pos: position{line: 171, col: 20, offset: 3496},
										exprs: []interface{}{
											&notExpr{
												pos: position{line: 171, col: 20, offset: 3496},
												expr: &litMatcher{
													pos:        position{line: 171, col: 21, offset: 3497},
													val:        "\"",
													ignoreCase: false,
												},
											},
											&anyMatcher{
												line: 171, col: 25, offset: 3501,
											},
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 171, col: 30, offset: 3506},
							val:        "\"",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "single_quoted_string",
			pos:  position{line: 179, col: 1, offset: 3667},
			expr: &actionExpr{
				pos: position{line: 180, col: 5, offset: 3694},
				run: (*parser).callonsingle_quoted_string1,
				expr: &seqExpr{
					pos: position{line: 180, col: 7, offset: 3696},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 180, col: 7, offset: 3696},
							val:        "'",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 180, col: 11, offset: 3700},
							expr: &choiceExpr{
								pos: position{line: 180, col: 13, offset: 3702},
								alternatives: []interface{}{
									&litMatcher{
										pos:        position{line: 180, col: 13, offset: 3702},
										val:        "\\'",
										ignoreCase: false,
									},
									&seqExpr{
										pos: position{line: 180, col: 20, offset: 3709},
										exprs: []interface{}{
											&notExpr{
												pos: position{line: 180, col: 20, offset: 3709},
												expr: &litMatcher{
													pos:        position{line: 180, col: 21, offset: 3710},
													val:        "'",
													ignoreCase: false,
												},
											},
											&anyMatcher{
												line: 180, col: 25, offset: 3714,
											},
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 180, col: 30, offset: 3719},
							val:        "'",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "string_value",
			pos:  position{line: 188, col: 1, offset: 3847},
			expr: &actionExpr{
				pos: position{line: 189, col: 5, offset: 3866},
				run: (*parser).callonstring_value1,
				expr: &labeledExpr{
					pos:   position{line: 189, col: 5, offset: 3866},
					label: "str",
					expr: &choiceExpr{
						pos: position{line: 189, col: 11, offset: 3872},
						alternatives: []interface{}{
							&actionExpr{
								pos: position{line: 189, col: 11, offset: 3872},
								run: (*parser).callonstring_value4,
								expr: &labeledExpr{
									pos:   position{line: 189, col: 11, offset: 3872},
									label: "str",
									expr: &ruleRefExpr{
										pos:  position{line: 189, col: 15, offset: 3876},
										name: "double_quoted_string",
									},
								},
							},
							&actionExpr{
								pos: position{line: 191, col: 9, offset: 3986},
								run: (*parser).callonstring_value7,
								expr: &labeledExpr{
									pos:   position{line: 191, col: 9, offset: 3986},
									label: "str",
									expr: &ruleRefExpr{
										pos:  position{line: 191, col: 13, offset: 3990},
										name: "single_quoted_string",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "regexp",
			pos:  position{line: 201, col: 1, offset: 4230},
			expr: &seqExpr{
				pos: position{line: 202, col: 7, offset: 4245},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 202, col: 7, offset: 4245},
						val:        "/",
						ignoreCase: false,
					},
					&zeroOrMoreExpr{
						pos: position{line: 202, col: 11, offset: 4249},
						expr: &choiceExpr{
							pos: position{line: 202, col: 13, offset: 4251},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 202, col: 13, offset: 4251},
									val:        "\\/",
									ignoreCase: false,
								},
								&seqExpr{
									pos: position{line: 202, col: 20, offset: 4258},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 202, col: 20, offset: 4258},
											expr: &litMatcher{
												pos:        position{line: 202, col: 21, offset: 4259},
												val:        "/",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 202, col: 25, offset: 4263,
										},
									},
								},
							},
						},
					},
					&litMatcher{
						pos:        position{line: 202, col: 30, offset: 4268},
						val:        "/",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "number",
			pos:  position{line: 209, col: 1, offset: 4373},
			expr: &actionExpr{
				pos: position{line: 210, col: 5, offset: 4386},
				run: (*parser).callonnumber1,
				expr: &seqExpr{
					pos: position{line: 210, col: 5, offset: 4386},
					exprs: []interface{}{
						&zeroOrOneExpr{
							pos: position{line: 210, col: 5, offset: 4386},
							expr: &litMatcher{
								pos:        position{line: 210, col: 5, offset: 4386},
								val:        "-",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 210, col: 10, offset: 4391},
							expr: &charClassMatcher{
								pos:        position{line: 210, col: 10, offset: 4391},
								val:        "[0-9]",
								ranges:     []rune{'0', '9'},
								ignoreCase: false,
								inverted:   false,
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 210, col: 17, offset: 4398},
							expr: &seqExpr{
								pos: position{line: 210, col: 18, offset: 4399},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 210, col: 18, offset: 4399},
										val:        ".",
										ignoreCase: false,
									},
									&zeroOrMoreExpr{
										pos: position{line: 210, col: 22, offset: 4403},
										expr: &charClassMatcher{
											pos:        position{line: 210, col: 22, offset: 4403},
											val:        "[0-9]",
											ranges:     []rune{'0', '9'},
											ignoreCase: false,
											inverted:   false,
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "array",
			pos:  position{line: 226, col: 1, offset: 4720},
			expr: &actionExpr{
				pos: position{line: 227, col: 5, offset: 4732},
				run: (*parser).callonarray1,
				expr: &seqExpr{
					pos: position{line: 227, col: 5, offset: 4732},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 227, col: 5, offset: 4732},
							val:        "[",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 227, col: 9, offset: 4736},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 227, col: 11, offset: 4738},
							label: "value",
							expr: &zeroOrOneExpr{
								pos: position{line: 227, col: 17, offset: 4744},
								expr: &actionExpr{
									pos: position{line: 228, col: 9, offset: 4755},
									run: (*parser).callonarray7,
									expr: &seqExpr{
										pos: position{line: 228, col: 9, offset: 4755},
										exprs: []interface{}{
											&labeledExpr{
												pos:   position{line: 228, col: 9, offset: 4755},
												label: "value",
												expr: &ruleRefExpr{
													pos:  position{line: 228, col: 15, offset: 4761},
													name: "value",
												},
											},
											&labeledExpr{
												pos:   position{line: 228, col: 21, offset: 4767},
												label: "values",
												expr: &zeroOrMoreExpr{
													pos: position{line: 228, col: 28, offset: 4774},
													expr: &actionExpr{
														pos: position{line: 229, col: 13, offset: 4788},
														run: (*parser).callonarray13,
														expr: &seqExpr{
															pos: position{line: 229, col: 13, offset: 4788},
															exprs: []interface{}{
																&ruleRefExpr{
																	pos:  position{line: 229, col: 13, offset: 4788},
																	name: "_",
																},
																&litMatcher{
																	pos:        position{line: 229, col: 15, offset: 4790},
																	val:        ",",
																	ignoreCase: false,
																},
																&ruleRefExpr{
																	pos:  position{line: 229, col: 19, offset: 4794},
																	name: "_",
																},
																&labeledExpr{
																	pos:   position{line: 229, col: 21, offset: 4796},
																	label: "value",
																	expr: &ruleRefExpr{
																		pos:  position{line: 229, col: 27, offset: 4802},
																		name: "value",
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 235, col: 8, offset: 4933},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 235, col: 10, offset: 4935},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "hash",
			pos:  position{line: 248, col: 1, offset: 5104},
			expr: &actionExpr{
				pos: position{line: 249, col: 5, offset: 5115},
				run: (*parser).callonhash1,
				expr: &seqExpr{
					pos: position{line: 249, col: 5, offset: 5115},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 249, col: 5, offset: 5115},
							val:        "{",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 249, col: 9, offset: 5119},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 249, col: 11, offset: 5121},
							label: "entries",
							expr: &zeroOrOneExpr{
								pos: position{line: 249, col: 19, offset: 5129},
								expr: &ruleRefExpr{
									pos:  position{line: 249, col: 19, offset: 5129},
									name: "hashentries",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 249, col: 32, offset: 5142},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 249, col: 34, offset: 5144},
							val:        "}",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "hashentries",
			pos:  position{line: 258, col: 1, offset: 5303},
			expr: &actionExpr{
				pos: position{line: 259, col: 5, offset: 5321},
				run: (*parser).callonhashentries1,
				expr: &seqExpr{
					pos: position{line: 259, col: 5, offset: 5321},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 259, col: 5, offset: 5321},
							label: "hashentry",
							expr: &ruleRefExpr{
								pos:  position{line: 259, col: 15, offset: 5331},
								name: "hashentry",
							},
						},
						&labeledExpr{
							pos:   position{line: 259, col: 25, offset: 5341},
							label: "hashentries1",
							expr: &zeroOrMoreExpr{
								pos: position{line: 259, col: 38, offset: 5354},
								expr: &actionExpr{
									pos: position{line: 260, col: 9, offset: 5364},
									run: (*parser).callonhashentries7,
									expr: &seqExpr{
										pos: position{line: 260, col: 9, offset: 5364},
										exprs: []interface{}{
											&ruleRefExpr{
												pos:  position{line: 260, col: 9, offset: 5364},
												name: "whitespace",
											},
											&labeledExpr{
												pos:   position{line: 260, col: 20, offset: 5375},
												label: "hashentry",
												expr: &ruleRefExpr{
													pos:  position{line: 260, col: 30, offset: 5385},
													name: "hashentry",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "hashentry",
			pos:  position{line: 272, col: 1, offset: 5637},
			expr: &actionExpr{
				pos: position{line: 273, col: 5, offset: 5653},
				run: (*parser).callonhashentry1,
				expr: &seqExpr{
					pos: position{line: 273, col: 5, offset: 5653},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 273, col: 5, offset: 5653},
							label: "name",
							expr: &choiceExpr{
								pos: position{line: 273, col: 11, offset: 5659},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 273, col: 11, offset: 5659},
										name: "number",
									},
									&ruleRefExpr{
										pos:  position{line: 273, col: 20, offset: 5668},
										name: "bareword",
									},
									&ruleRefExpr{
										pos:  position{line: 273, col: 31, offset: 5679},
										name: "string_value",
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 273, col: 45, offset: 5693},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 273, col: 47, offset: 5695},
							val:        "=>",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 273, col: 52, offset: 5700},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 273, col: 54, offset: 5702},
							label: "value",
							expr: &ruleRefExpr{
								pos:  position{line: 273, col: 60, offset: 5708},
								name: "value",
							},
						},
					},
				},
			},
		},
		{
			name: "branch",
			pos:  position{line: 284, col: 1, offset: 5875},
			expr: &seqExpr{
				pos: position{line: 285, col: 5, offset: 5888},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 285, col: 5, offset: 5888},
						name: "if_cond",
					},
					&zeroOrMoreExpr{
						pos: position{line: 285, col: 13, offset: 5896},
						expr: &seqExpr{
							pos: position{line: 285, col: 14, offset: 5897},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 285, col: 14, offset: 5897},
									name: "_",
								},
								&ruleRefExpr{
									pos:  position{line: 285, col: 16, offset: 5899},
									name: "else_if",
								},
							},
						},
					},
					&zeroOrOneExpr{
						pos: position{line: 285, col: 26, offset: 5909},
						expr: &seqExpr{
							pos: position{line: 285, col: 27, offset: 5910},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 285, col: 27, offset: 5910},
									name: "_",
								},
								&ruleRefExpr{
									pos:  position{line: 285, col: 29, offset: 5912},
									name: "else_cond",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "if_cond",
			pos:  position{line: 292, col: 1, offset: 6041},
			expr: &seqExpr{
				pos: position{line: 293, col: 5, offset: 6055},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 293, col: 5, offset: 6055},
						val:        "if",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 293, col: 10, offset: 6060},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 293, col: 12, offset: 6062},
						name: "condition",
					},
					&ruleRefExpr{
						pos:  position{line: 293, col: 22, offset: 6072},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 293, col: 24, offset: 6074},
						val:        "{",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 293, col: 28, offset: 6078},
						name: "_",
					},
					&zeroOrMoreExpr{
						pos: position{line: 293, col: 30, offset: 6080},
						expr: &seqExpr{
							pos: position{line: 293, col: 31, offset: 6081},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 293, col: 31, offset: 6081},
									name: "branch_or_plugin",
								},
								&ruleRefExpr{
									pos:  position{line: 293, col: 48, offset: 6098},
									name: "_",
								},
							},
						},
					},
					&litMatcher{
						pos:        position{line: 293, col: 52, offset: 6102},
						val:        "}",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "else_if",
			pos:  position{line: 300, col: 1, offset: 6241},
			expr: &seqExpr{
				pos: position{line: 301, col: 5, offset: 6255},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 301, col: 5, offset: 6255},
						val:        "else",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 301, col: 12, offset: 6262},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 301, col: 14, offset: 6264},
						val:        "if",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 301, col: 19, offset: 6269},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 301, col: 21, offset: 6271},
						name: "condition",
					},
					&ruleRefExpr{
						pos:  position{line: 301, col: 31, offset: 6281},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 301, col: 33, offset: 6283},
						val:        "{",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 301, col: 37, offset: 6287},
						name: "_",
					},
					&zeroOrMoreExpr{
						pos: position{line: 301, col: 39, offset: 6289},
						expr: &seqExpr{
							pos: position{line: 301, col: 41, offset: 6291},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 301, col: 41, offset: 6291},
									name: "branch_or_plugin",
								},
								&ruleRefExpr{
									pos:  position{line: 301, col: 58, offset: 6308},
									name: "_",
								},
							},
						},
					},
					&litMatcher{
						pos:        position{line: 301, col: 62, offset: 6312},
						val:        "}",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "else_cond",
			pos:  position{line: 308, col: 1, offset: 6427},
			expr: &seqExpr{
				pos: position{line: 309, col: 5, offset: 6443},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 309, col: 5, offset: 6443},
						val:        "else",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 309, col: 12, offset: 6450},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 309, col: 14, offset: 6452},
						val:        "{",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 309, col: 18, offset: 6456},
						name: "_",
					},
					&zeroOrMoreExpr{
						pos: position{line: 309, col: 20, offset: 6458},
						expr: &seqExpr{
							pos: position{line: 309, col: 21, offset: 6459},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 309, col: 21, offset: 6459},
									name: "branch_or_plugin",
								},
								&ruleRefExpr{
									pos:  position{line: 309, col: 38, offset: 6476},
									name: "_",
								},
							},
						},
					},
					&litMatcher{
						pos:        position{line: 309, col: 42, offset: 6480},
						val:        "}",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "condition",
			pos:  position{line: 316, col: 1, offset: 6610},
			expr: &seqExpr{
				pos: position{line: 317, col: 5, offset: 6626},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 317, col: 5, offset: 6626},
						name: "expression",
					},
					&zeroOrMoreExpr{
						pos: position{line: 317, col: 16, offset: 6637},
						expr: &seqExpr{
							pos: position{line: 317, col: 17, offset: 6638},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 317, col: 17, offset: 6638},
									name: "_",
								},
								&ruleRefExpr{
									pos:  position{line: 317, col: 19, offset: 6640},
									name: "boolean_operator",
								},
								&ruleRefExpr{
									pos:  position{line: 317, col: 36, offset: 6657},
									name: "_",
								},
								&ruleRefExpr{
									pos:  position{line: 317, col: 38, offset: 6659},
									name: "expression",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "expression",
			pos:  position{line: 331, col: 1, offset: 6955},
			expr: &choiceExpr{
				pos: position{line: 333, col: 9, offset: 6983},
				alternatives: []interface{}{
					&seqExpr{
						pos: position{line: 333, col: 10, offset: 6984},
						exprs: []interface{}{
							&litMatcher{
								pos:        position{line: 333, col: 10, offset: 6984},
								val:        "(",
								ignoreCase: false,
							},
							&ruleRefExpr{
								pos:  position{line: 333, col: 14, offset: 6988},
								name: "_",
							},
							&ruleRefExpr{
								pos:  position{line: 333, col: 16, offset: 6990},
								name: "condition",
							},
							&ruleRefExpr{
								pos:  position{line: 333, col: 26, offset: 7000},
								name: "_",
							},
							&litMatcher{
								pos:        position{line: 333, col: 28, offset: 7002},
								val:        ")",
								ignoreCase: false,
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 334, col: 9, offset: 7015},
						name: "negative_expression",
					},
					&ruleRefExpr{
						pos:  position{line: 335, col: 9, offset: 7043},
						name: "in_expression",
					},
					&ruleRefExpr{
						pos:  position{line: 336, col: 9, offset: 7065},
						name: "not_in_expression",
					},
					&ruleRefExpr{
						pos:  position{line: 337, col: 9, offset: 7091},
						name: "compare_expression",
					},
					&ruleRefExpr{
						pos:  position{line: 338, col: 9, offset: 7118},
						name: "regexp_expression",
					},
					&ruleRefExpr{
						pos:  position{line: 339, col: 9, offset: 7144},
						name: "rvalue",
					},
				},
			},
		},
		{
			name: "negative_expression",
			pos:  position{line: 349, col: 1, offset: 7329},
			expr: &choiceExpr{
				pos: position{line: 351, col: 9, offset: 7366},
				alternatives: []interface{}{
					&seqExpr{
						pos: position{line: 351, col: 10, offset: 7367},
						exprs: []interface{}{
							&litMatcher{
								pos:        position{line: 351, col: 10, offset: 7367},
								val:        "!",
								ignoreCase: false,
							},
							&ruleRefExpr{
								pos:  position{line: 351, col: 14, offset: 7371},
								name: "_",
							},
							&litMatcher{
								pos:        position{line: 351, col: 16, offset: 7373},
								val:        "(",
								ignoreCase: false,
							},
							&ruleRefExpr{
								pos:  position{line: 351, col: 20, offset: 7377},
								name: "_",
							},
							&ruleRefExpr{
								pos:  position{line: 351, col: 22, offset: 7379},
								name: "condition",
							},
							&ruleRefExpr{
								pos:  position{line: 351, col: 32, offset: 7389},
								name: "_",
							},
							&litMatcher{
								pos:        position{line: 351, col: 34, offset: 7391},
								val:        ")",
								ignoreCase: false,
							},
						},
					},
					&seqExpr{
						pos: position{line: 352, col: 10, offset: 7405},
						exprs: []interface{}{
							&litMatcher{
								pos:        position{line: 352, col: 10, offset: 7405},
								val:        "!",
								ignoreCase: false,
							},
							&ruleRefExpr{
								pos:  position{line: 352, col: 14, offset: 7409},
								name: "_",
							},
							&ruleRefExpr{
								pos:  position{line: 352, col: 16, offset: 7411},
								name: "selector",
							},
						},
					},
				},
			},
		},
		{
			name: "in_expression",
			pos:  position{line: 360, col: 1, offset: 7544},
			expr: &seqExpr{
				pos: position{line: 361, col: 5, offset: 7564},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 361, col: 5, offset: 7564},
						name: "rvalue",
					},
					&ruleRefExpr{
						pos:  position{line: 361, col: 12, offset: 7571},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 361, col: 14, offset: 7573},
						name: "in_operator",
					},
					&ruleRefExpr{
						pos:  position{line: 361, col: 26, offset: 7585},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 361, col: 28, offset: 7587},
						name: "rvalue",
					},
				},
			},
		},
		{
			name: "not_in_expression",
			pos:  position{line: 368, col: 1, offset: 7722},
			expr: &seqExpr{
				pos: position{line: 369, col: 5, offset: 7746},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 369, col: 5, offset: 7746},
						name: "rvalue",
					},
					&ruleRefExpr{
						pos:  position{line: 369, col: 12, offset: 7753},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 369, col: 14, offset: 7755},
						name: "not_in_operator",
					},
					&ruleRefExpr{
						pos:  position{line: 369, col: 30, offset: 7771},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 369, col: 32, offset: 7773},
						name: "rvalue",
					},
				},
			},
		},
		{
			name: "in_operator",
			pos:  position{line: 375, col: 1, offset: 7825},
			expr: &litMatcher{
				pos:        position{line: 376, col: 5, offset: 7843},
				val:        "in",
				ignoreCase: false,
			},
		},
		{
			name: "not_in_operator",
			pos:  position{line: 382, col: 1, offset: 7906},
			expr: &seqExpr{
				pos: position{line: 383, col: 5, offset: 7928},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 383, col: 5, offset: 7928},
						val:        "not ",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 383, col: 12, offset: 7935},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 383, col: 14, offset: 7937},
						val:        "in",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "rvalue",
			pos:  position{line: 392, col: 1, offset: 8196},
			expr: &choiceExpr{
				pos: position{line: 393, col: 5, offset: 8209},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 393, col: 5, offset: 8209},
						name: "string_value",
					},
					&ruleRefExpr{
						pos:  position{line: 393, col: 20, offset: 8224},
						name: "number",
					},
					&ruleRefExpr{
						pos:  position{line: 393, col: 29, offset: 8233},
						name: "selector",
					},
					&ruleRefExpr{
						pos:  position{line: 393, col: 40, offset: 8244},
						name: "array",
					},
					&ruleRefExpr{
						pos:  position{line: 393, col: 48, offset: 8252},
						name: "regexp",
					},
				},
			},
		},
		{
			name: "compare_expression",
			pos:  position{line: 425, col: 1, offset: 8928},
			expr: &seqExpr{
				pos: position{line: 426, col: 5, offset: 8953},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 426, col: 5, offset: 8953},
						name: "rvalue",
					},
					&ruleRefExpr{
						pos:  position{line: 426, col: 12, offset: 8960},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 426, col: 14, offset: 8962},
						name: "compare_operator",
					},
					&ruleRefExpr{
						pos:  position{line: 426, col: 31, offset: 8979},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 426, col: 33, offset: 8981},
						name: "rvalue",
					},
				},
			},
		},
		{
			name: "compare_operator",
			pos:  position{line: 433, col: 1, offset: 9126},
			expr: &choiceExpr{
				pos: position{line: 434, col: 6, offset: 9150},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 434, col: 6, offset: 9150},
						val:        "==",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 434, col: 13, offset: 9157},
						val:        "!=",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 434, col: 20, offset: 9164},
						val:        "<=",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 434, col: 27, offset: 9171},
						val:        ">=",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 434, col: 34, offset: 9178},
						val:        "<",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 434, col: 40, offset: 9184},
						val:        ">",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "regexp_expression",
			pos:  position{line: 441, col: 1, offset: 9330},
			expr: &seqExpr{
				pos: position{line: 442, col: 5, offset: 9354},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 442, col: 5, offset: 9354},
						name: "rvalue",
					},
					&ruleRefExpr{
						pos:  position{line: 442, col: 12, offset: 9361},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 442, col: 15, offset: 9364},
						name: "regexp_operator",
					},
					&ruleRefExpr{
						pos:  position{line: 442, col: 31, offset: 9380},
						name: "_",
					},
					&choiceExpr{
						pos: position{line: 442, col: 34, offset: 9383},
						alternatives: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 442, col: 34, offset: 9383},
								name: "string_value",
							},
							&ruleRefExpr{
								pos:  position{line: 442, col: 49, offset: 9398},
								name: "regexp",
							},
						},
					},
				},
			},
		},
		{
			name: "regexp_operator",
			pos:  position{line: 448, col: 1, offset: 9504},
			expr: &choiceExpr{
				pos: position{line: 449, col: 6, offset: 9527},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 449, col: 6, offset: 9527},
						val:        "=~",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 449, col: 13, offset: 9534},
						val:        "!~",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "boolean_operator",
			pos:  position{line: 456, col: 1, offset: 9665},
			expr: &choiceExpr{
				pos: position{line: 457, col: 6, offset: 9689},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 457, col: 6, offset: 9689},
						val:        "and",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 457, col: 14, offset: 9697},
						val:        "or",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 457, col: 21, offset: 9704},
						val:        "xor",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 457, col: 29, offset: 9712},
						val:        "nand",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "selector",
			pos:  position{line: 464, col: 1, offset: 9816},
			expr: &oneOrMoreExpr{
				pos: position{line: 465, col: 5, offset: 9831},
				expr: &ruleRefExpr{
					pos:  position{line: 465, col: 5, offset: 9831},
					name: "selector_element",
				},
			},
		},
		{
			name: "selector_element",
			pos:  position{line: 472, col: 1, offset: 9958},
			expr: &seqExpr{
				pos: position{line: 473, col: 5, offset: 9981},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 473, col: 5, offset: 9981},
						val:        "[",
						ignoreCase: false,
					},
					&oneOrMoreExpr{
						pos: position{line: 473, col: 9, offset: 9985},
						expr: &charClassMatcher{
							pos:        position{line: 473, col: 9, offset: 9985},
							val:        "[^\\],]",
							chars:      []rune{']', ','},
							ignoreCase: false,
							inverted:   true,
						},
					},
					&litMatcher{
						pos:        position{line: 473, col: 17, offset: 9993},
						val:        "]",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "EOF",
			pos:  position{line: 475, col: 1, offset: 9998},
			expr: &notExpr{
				pos: position{line: 475, col: 7, offset: 10004},
				expr: &anyMatcher{
					line: 475, col: 8, offset: 10005,
				},
			},
		},
	},
}

func (c *current) onconfig9(ps interface{}) (interface{}, error) {

	return ret(ps)

}

func (p *parser) callonconfig9() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onconfig9(stack["ps"])
}

func (c *current) onconfig1(ps, pss interface{}) (interface{}, error) {

	return config(ps, pss)

}

func (p *parser) callonconfig1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onconfig1(stack["ps"], stack["pss"])
}

func (c *current) onplugin_section10(bop interface{}) (interface{}, error) {

	return ret(bop)

}

func (p *parser) callonplugin_section10() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onplugin_section10(stack["bop"])
}

func (c *current) onplugin_section1(pt, bops interface{}) (interface{}, error) {

	return pluginSection(pt, bops)

}

func (p *parser) callonplugin_section1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onplugin_section1(stack["pt"], stack["bops"])
}

func (c *current) onplugin_type2() (interface{}, error) {
	return ast.Input, nil

}

func (p *parser) callonplugin_type2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onplugin_type2()
}

func (c *current) onplugin_type4() (interface{}, error) {
	return ast.Filter, nil

}

func (p *parser) callonplugin_type4() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onplugin_type4()
}

func (c *current) onplugin_type6() (interface{}, error) {
	return ast.Output, nil

}

func (p *parser) callonplugin_type6() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onplugin_type6()
}

func (c *current) onplugin16(attribute interface{}) (interface{}, error) {
	return ret(attribute)

}

func (p *parser) callonplugin16() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onplugin16(stack["attribute"])
}

func (c *current) onplugin10(attribute, attrs interface{}) (interface{}, error) {
	return attributes(attribute, attrs)

}

func (p *parser) callonplugin10() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onplugin10(stack["attribute"], stack["attrs"])
}

func (c *current) onplugin1(name, attributes interface{}) (interface{}, error) {
	return plugin(name, attributes)

}

func (p *parser) callonplugin1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onplugin1(stack["name"], stack["attributes"])
}

func (c *current) onname2() (interface{}, error) {
	return string(c.text), nil

}

func (p *parser) callonname2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onname2()
}

func (c *current) onname5(value interface{}) (interface{}, error) {
	return ret(value)

}

func (p *parser) callonname5() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onname5(stack["value"])
}

func (c *current) onattribute1(name, value interface{}) (interface{}, error) {
	return attribute(name, value)

}

func (p *parser) callonattribute1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onattribute1(stack["name"], stack["value"])
}

func (c *current) onbareword1() (interface{}, error) {
	return ast.NewStringAttribute("", string(c.text), ast.Bareword), nil

}

func (p *parser) callonbareword1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onbareword1()
}

func (c *current) ondouble_quoted_string1() (interface{}, error) {
	return quotedvalue(c, `"`)

}

func (p *parser) callondouble_quoted_string1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.ondouble_quoted_string1()
}

func (c *current) onsingle_quoted_string1() (interface{}, error) {
	return quotedvalue(c, `'`)

}

func (p *parser) callonsingle_quoted_string1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onsingle_quoted_string1()
}

func (c *current) onstring_value4(str interface{}) (interface{}, error) {
	return ast.NewStringAttribute("", str.(string), ast.DoubleQuoted), nil

}

func (p *parser) callonstring_value4() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onstring_value4(stack["str"])
}

func (c *current) onstring_value7(str interface{}) (interface{}, error) {
	return ast.NewStringAttribute("", str.(string), ast.SingleQuoted), nil

}

func (p *parser) callonstring_value7() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onstring_value7(stack["str"])
}

func (c *current) onstring_value1(str interface{}) (interface{}, error) {
	return ret(str)

}

func (p *parser) callonstring_value1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onstring_value1(stack["str"])
}

func (c *current) onnumber1() (interface{}, error) {
	return number(string(c.text))

}

func (p *parser) callonnumber1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onnumber1()
}

func (c *current) onarray13(value interface{}) (interface{}, error) {
	return ret(value)

}

func (p *parser) callonarray13() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onarray13(stack["value"])
}

func (c *current) onarray7(value, values interface{}) (interface{}, error) {
	return attributes(value, values)

}

func (p *parser) callonarray7() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onarray7(stack["value"], stack["values"])
}

func (c *current) onarray1(value interface{}) (interface{}, error) {
	return array(value)

}

func (p *parser) callonarray1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onarray1(stack["value"])
}

func (c *current) onhash1(entries interface{}) (interface{}, error) {
	return hash(entries)

}

func (p *parser) callonhash1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onhash1(stack["entries"])
}

func (c *current) onhashentries7(hashentry interface{}) (interface{}, error) {
	return ret(hashentry)

}

func (p *parser) callonhashentries7() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onhashentries7(stack["hashentry"])
}

func (c *current) onhashentries1(hashentry, hashentries1 interface{}) (interface{}, error) {
	return hashentries(hashentry, hashentries1)

}

func (p *parser) callonhashentries1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onhashentries1(stack["hashentry"], stack["hashentries1"])
}

func (c *current) onhashentry1(name, value interface{}) (interface{}, error) {
	return hashentry(name, value)

}

func (p *parser) callonhashentry1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onhashentry1(stack["name"], stack["value"])
}

var (
	// errNoRule is returned when the grammar to parse has no rule.
	errNoRule = errors.New("grammar has no rule")

	// errInvalidEncoding is returned when the source is not properly
	// utf8-encoded.
	errInvalidEncoding = errors.New("invalid encoding")

	// errNoMatch is returned if no match could be found.
	errNoMatch = errors.New("no match found")
)

// Option is a function that can set an option on the parser. It returns
// the previous setting as an Option.
type Option func(*parser) Option

// Debug creates an Option to set the debug flag to b. When set to true,
// debugging information is printed to stdout while parsing.
//
// The default is false.
func Debug(b bool) Option {
	return func(p *parser) Option {
		old := p.debug
		p.debug = b
		return Debug(old)
	}
}

// Memoize creates an Option to set the memoize flag to b. When set to true,
// the parser will cache all results so each expression is evaluated only
// once. This guarantees linear parsing time even for pathological cases,
// at the expense of more memory and slower times for typical cases.
//
// The default is false.
func Memoize(b bool) Option {
	return func(p *parser) Option {
		old := p.memoize
		p.memoize = b
		return Memoize(old)
	}
}

// Recover creates an Option to set the recover flag to b. When set to
// true, this causes the parser to recover from panics and convert it
// to an error. Setting it to false can be useful while debugging to
// access the full stack trace.
//
// The default is true.
func Recover(b bool) Option {
	return func(p *parser) Option {
		old := p.recover
		p.recover = b
		return Recover(old)
	}
}

// ParseFile parses the file identified by filename.
func ParseFile(filename string, opts ...Option) (interface{}, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ParseReader(filename, f, opts...)
}

// ParseReader parses the data from r using filename as information in the
// error messages.
func ParseReader(filename string, r io.Reader, opts ...Option) (interface{}, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	return Parse(filename, b, opts...)
}

// Parse parses the data from b using filename as information in the
// error messages.
func Parse(filename string, b []byte, opts ...Option) (interface{}, error) {
	return newParser(filename, b, opts...).parse(g)
}

// position records a position in the text.
type position struct {
	line, col, offset int
}

func (p position) String() string {
	return fmt.Sprintf("%d:%d [%d]", p.line, p.col, p.offset)
}

// savepoint stores all state required to go back to this point in the
// parser.
type savepoint struct {
	position
	rn rune
	w  int
}

type current struct {
	pos  position // start position of the match
	text []byte   // raw text of the match
}

// the AST types...

type grammar struct {
	pos   position
	rules []*rule
}

type rule struct {
	pos         position
	name        string
	displayName string
	expr        interface{}
}

type choiceExpr struct {
	pos          position
	alternatives []interface{}
}

type actionExpr struct {
	pos  position
	expr interface{}
	run  func(*parser) (interface{}, error)
}

type seqExpr struct {
	pos   position
	exprs []interface{}
}

type labeledExpr struct {
	pos   position
	label string
	expr  interface{}
}

type expr struct {
	pos  position
	expr interface{}
}

type andExpr expr
type notExpr expr
type zeroOrOneExpr expr
type zeroOrMoreExpr expr
type oneOrMoreExpr expr

type ruleRefExpr struct {
	pos  position
	name string
}

type andCodeExpr struct {
	pos position
	run func(*parser) (bool, error)
}

type notCodeExpr struct {
	pos position
	run func(*parser) (bool, error)
}

type litMatcher struct {
	pos        position
	val        string
	ignoreCase bool
}

type charClassMatcher struct {
	pos        position
	val        string
	chars      []rune
	ranges     []rune
	classes    []*unicode.RangeTable
	ignoreCase bool
	inverted   bool
}

type anyMatcher position

// errList cumulates the errors found by the parser.
type errList []error

func (e *errList) add(err error) {
	*e = append(*e, err)
}

func (e errList) err() error {
	if len(e) == 0 {
		return nil
	}
	e.dedupe()
	return e
}

func (e *errList) dedupe() {
	var cleaned []error
	set := make(map[string]bool)
	for _, err := range *e {
		if msg := err.Error(); !set[msg] {
			set[msg] = true
			cleaned = append(cleaned, err)
		}
	}
	*e = cleaned
}

func (e errList) Error() string {
	switch len(e) {
	case 0:
		return ""
	case 1:
		return e[0].Error()
	default:
		var buf bytes.Buffer

		for i, err := range e {
			if i > 0 {
				buf.WriteRune('\n')
			}
			buf.WriteString(err.Error())
		}
		return buf.String()
	}
}

// parserError wraps an error with a prefix indicating the rule in which
// the error occurred. The original error is stored in the Inner field.
type parserError struct {
	Inner  error
	pos    position
	prefix string
}

// Error returns the error message.
func (p *parserError) Error() string {
	return p.prefix + ": " + p.Inner.Error()
}

// newParser creates a parser with the specified input source and options.
func newParser(filename string, b []byte, opts ...Option) *parser {
	p := &parser{
		filename: filename,
		errs:     new(errList),
		data:     b,
		pt:       savepoint{position: position{line: 1}},
		recover:  true,
	}
	p.setOptions(opts)
	return p
}

// setOptions applies the options to the parser.
func (p *parser) setOptions(opts []Option) {
	for _, opt := range opts {
		opt(p)
	}
}

type resultTuple struct {
	v   interface{}
	b   bool
	end savepoint
}

type parser struct {
	filename string
	pt       savepoint
	cur      current

	data []byte
	errs *errList

	recover bool
	debug   bool
	depth   int

	memoize bool
	// memoization table for the packrat algorithm:
	// map[offset in source] map[expression or rule] {value, match}
	memo map[int]map[interface{}]resultTuple

	// rules table, maps the rule identifier to the rule node
	rules map[string]*rule
	// variables stack, map of label to value
	vstack []map[string]interface{}
	// rule stack, allows identification of the current rule in errors
	rstack []*rule

	// stats
	exprCnt int
}

// push a variable set on the vstack.
func (p *parser) pushV() {
	if cap(p.vstack) == len(p.vstack) {
		// create new empty slot in the stack
		p.vstack = append(p.vstack, nil)
	} else {
		// slice to 1 more
		p.vstack = p.vstack[:len(p.vstack)+1]
	}

	// get the last args set
	m := p.vstack[len(p.vstack)-1]
	if m != nil && len(m) == 0 {
		// empty map, all good
		return
	}

	m = make(map[string]interface{})
	p.vstack[len(p.vstack)-1] = m
}

// pop a variable set from the vstack.
func (p *parser) popV() {
	// if the map is not empty, clear it
	m := p.vstack[len(p.vstack)-1]
	if len(m) > 0 {
		// GC that map
		p.vstack[len(p.vstack)-1] = nil
	}
	p.vstack = p.vstack[:len(p.vstack)-1]
}

func (p *parser) print(prefix, s string) string {
	if !p.debug {
		return s
	}

	fmt.Printf("%s %d:%d:%d: %s [%#U]\n",
		prefix, p.pt.line, p.pt.col, p.pt.offset, s, p.pt.rn)
	return s
}

func (p *parser) in(s string) string {
	p.depth++
	return p.print(strings.Repeat(" ", p.depth)+">", s)
}

func (p *parser) out(s string) string {
	p.depth--
	return p.print(strings.Repeat(" ", p.depth)+"<", s)
}

func (p *parser) addErr(err error) {
	p.addErrAt(err, p.pt.position)
}

func (p *parser) addErrAt(err error, pos position) {
	var buf bytes.Buffer
	if p.filename != "" {
		buf.WriteString(p.filename)
	}
	if buf.Len() > 0 {
		buf.WriteString(":")
	}
	buf.WriteString(fmt.Sprintf("%d:%d (%d)", pos.line, pos.col, pos.offset))
	if len(p.rstack) > 0 {
		if buf.Len() > 0 {
			buf.WriteString(": ")
		}
		rule := p.rstack[len(p.rstack)-1]
		if rule.displayName != "" {
			buf.WriteString("rule " + rule.displayName)
		} else {
			buf.WriteString("rule " + rule.name)
		}
	}
	pe := &parserError{Inner: err, pos: pos, prefix: buf.String()}
	p.errs.add(pe)
}

// read advances the parser to the next rune.
func (p *parser) read() {
	p.pt.offset += p.pt.w
	rn, n := utf8.DecodeRune(p.data[p.pt.offset:])
	p.pt.rn = rn
	p.pt.w = n
	p.pt.col++
	if rn == '\n' {
		p.pt.line++
		p.pt.col = 0
	}

	if rn == utf8.RuneError {
		if n == 1 {
			p.addErr(errInvalidEncoding)
		}
	}
}

// restore parser position to the savepoint pt.
func (p *parser) restore(pt savepoint) {
	if p.debug {
		defer p.out(p.in("restore"))
	}
	if pt.offset == p.pt.offset {
		return
	}
	p.pt = pt
}

// get the slice of bytes from the savepoint start to the current position.
func (p *parser) sliceFrom(start savepoint) []byte {
	return p.data[start.position.offset:p.pt.position.offset]
}

func (p *parser) getMemoized(node interface{}) (resultTuple, bool) {
	if len(p.memo) == 0 {
		return resultTuple{}, false
	}
	m := p.memo[p.pt.offset]
	if len(m) == 0 {
		return resultTuple{}, false
	}
	res, ok := m[node]
	return res, ok
}

func (p *parser) setMemoized(pt savepoint, node interface{}, tuple resultTuple) {
	if p.memo == nil {
		p.memo = make(map[int]map[interface{}]resultTuple)
	}
	m := p.memo[pt.offset]
	if m == nil {
		m = make(map[interface{}]resultTuple)
		p.memo[pt.offset] = m
	}
	m[node] = tuple
}

func (p *parser) buildRulesTable(g *grammar) {
	p.rules = make(map[string]*rule, len(g.rules))
	for _, r := range g.rules {
		p.rules[r.name] = r
	}
}

func (p *parser) parse(g *grammar) (val interface{}, err error) {
	if len(g.rules) == 0 {
		p.addErr(errNoRule)
		return nil, p.errs.err()
	}

	// TODO : not super critical but this could be generated
	p.buildRulesTable(g)

	if p.recover {
		// panic can be used in action code to stop parsing immediately
		// and return the panic as an error.
		defer func() {
			if e := recover(); e != nil {
				if p.debug {
					defer p.out(p.in("panic handler"))
				}
				val = nil
				switch e := e.(type) {
				case error:
					p.addErr(e)
				default:
					p.addErr(fmt.Errorf("%v", e))
				}
				err = p.errs.err()
			}
		}()
	}

	// start rule is rule [0]
	p.read() // advance to first rune
	val, ok := p.parseRule(g.rules[0])
	if !ok {
		if len(*p.errs) == 0 {
			// make sure this doesn't go out silently
			p.addErr(errNoMatch)
		}
		return nil, p.errs.err()
	}
	return val, p.errs.err()
}

func (p *parser) parseRule(rule *rule) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseRule " + rule.name))
	}

	if p.memoize {
		res, ok := p.getMemoized(rule)
		if ok {
			p.restore(res.end)
			return res.v, res.b
		}
	}

	start := p.pt
	p.rstack = append(p.rstack, rule)
	p.pushV()
	val, ok := p.parseExpr(rule.expr)
	p.popV()
	p.rstack = p.rstack[:len(p.rstack)-1]
	if ok && p.debug {
		p.print(strings.Repeat(" ", p.depth)+"MATCH", string(p.sliceFrom(start)))
	}

	if p.memoize {
		p.setMemoized(start, rule, resultTuple{val, ok, p.pt})
	}
	return val, ok
}

func (p *parser) parseExpr(expr interface{}) (interface{}, bool) {
	var pt savepoint
	var ok bool

	if p.memoize {
		res, ok := p.getMemoized(expr)
		if ok {
			p.restore(res.end)
			return res.v, res.b
		}
		pt = p.pt
	}

	p.exprCnt++
	var val interface{}
	switch expr := expr.(type) {
	case *actionExpr:
		val, ok = p.parseActionExpr(expr)
	case *andCodeExpr:
		val, ok = p.parseAndCodeExpr(expr)
	case *andExpr:
		val, ok = p.parseAndExpr(expr)
	case *anyMatcher:
		val, ok = p.parseAnyMatcher(expr)
	case *charClassMatcher:
		val, ok = p.parseCharClassMatcher(expr)
	case *choiceExpr:
		val, ok = p.parseChoiceExpr(expr)
	case *labeledExpr:
		val, ok = p.parseLabeledExpr(expr)
	case *litMatcher:
		val, ok = p.parseLitMatcher(expr)
	case *notCodeExpr:
		val, ok = p.parseNotCodeExpr(expr)
	case *notExpr:
		val, ok = p.parseNotExpr(expr)
	case *oneOrMoreExpr:
		val, ok = p.parseOneOrMoreExpr(expr)
	case *ruleRefExpr:
		val, ok = p.parseRuleRefExpr(expr)
	case *seqExpr:
		val, ok = p.parseSeqExpr(expr)
	case *zeroOrMoreExpr:
		val, ok = p.parseZeroOrMoreExpr(expr)
	case *zeroOrOneExpr:
		val, ok = p.parseZeroOrOneExpr(expr)
	default:
		panic(fmt.Sprintf("unknown expression type %T", expr))
	}
	if p.memoize {
		p.setMemoized(pt, expr, resultTuple{val, ok, p.pt})
	}
	return val, ok
}

func (p *parser) parseActionExpr(act *actionExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseActionExpr"))
	}

	start := p.pt
	val, ok := p.parseExpr(act.expr)
	if ok {
		p.cur.pos = start.position
		p.cur.text = p.sliceFrom(start)
		actVal, err := act.run(p)
		if err != nil {
			p.addErrAt(err, start.position)
		}
		val = actVal
	}
	if ok && p.debug {
		p.print(strings.Repeat(" ", p.depth)+"MATCH", string(p.sliceFrom(start)))
	}
	return val, ok
}

func (p *parser) parseAndCodeExpr(and *andCodeExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseAndCodeExpr"))
	}

	ok, err := and.run(p)
	if err != nil {
		p.addErr(err)
	}
	return nil, ok
}

func (p *parser) parseAndExpr(and *andExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseAndExpr"))
	}

	pt := p.pt
	p.pushV()
	_, ok := p.parseExpr(and.expr)
	p.popV()
	p.restore(pt)
	return nil, ok
}

func (p *parser) parseAnyMatcher(any *anyMatcher) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseAnyMatcher"))
	}

	if p.pt.rn != utf8.RuneError {
		start := p.pt
		p.read()
		return p.sliceFrom(start), true
	}
	return nil, false
}

func (p *parser) parseCharClassMatcher(chr *charClassMatcher) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseCharClassMatcher"))
	}

	cur := p.pt.rn
	// can't match EOF
	if cur == utf8.RuneError {
		return nil, false
	}
	start := p.pt
	if chr.ignoreCase {
		cur = unicode.ToLower(cur)
	}

	// try to match in the list of available chars
	for _, rn := range chr.chars {
		if rn == cur {
			if chr.inverted {
				return nil, false
			}
			p.read()
			return p.sliceFrom(start), true
		}
	}

	// try to match in the list of ranges
	for i := 0; i < len(chr.ranges); i += 2 {
		if cur >= chr.ranges[i] && cur <= chr.ranges[i+1] {
			if chr.inverted {
				return nil, false
			}
			p.read()
			return p.sliceFrom(start), true
		}
	}

	// try to match in the list of Unicode classes
	for _, cl := range chr.classes {
		if unicode.Is(cl, cur) {
			if chr.inverted {
				return nil, false
			}
			p.read()
			return p.sliceFrom(start), true
		}
	}

	if chr.inverted {
		p.read()
		return p.sliceFrom(start), true
	}
	return nil, false
}

func (p *parser) parseChoiceExpr(ch *choiceExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseChoiceExpr"))
	}

	for _, alt := range ch.alternatives {
		p.pushV()
		val, ok := p.parseExpr(alt)
		p.popV()
		if ok {
			return val, ok
		}
	}
	return nil, false
}

func (p *parser) parseLabeledExpr(lab *labeledExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseLabeledExpr"))
	}

	p.pushV()
	val, ok := p.parseExpr(lab.expr)
	p.popV()
	if ok && lab.label != "" {
		m := p.vstack[len(p.vstack)-1]
		m[lab.label] = val
	}
	return val, ok
}

func (p *parser) parseLitMatcher(lit *litMatcher) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseLitMatcher"))
	}

	start := p.pt
	for _, want := range lit.val {
		cur := p.pt.rn
		if lit.ignoreCase {
			cur = unicode.ToLower(cur)
		}
		if cur != want {
			p.restore(start)
			return nil, false
		}
		p.read()
	}
	return p.sliceFrom(start), true
}

func (p *parser) parseNotCodeExpr(not *notCodeExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseNotCodeExpr"))
	}

	ok, err := not.run(p)
	if err != nil {
		p.addErr(err)
	}
	return nil, !ok
}

func (p *parser) parseNotExpr(not *notExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseNotExpr"))
	}

	pt := p.pt
	p.pushV()
	_, ok := p.parseExpr(not.expr)
	p.popV()
	p.restore(pt)
	return nil, !ok
}

func (p *parser) parseOneOrMoreExpr(expr *oneOrMoreExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseOneOrMoreExpr"))
	}

	var vals []interface{}

	for {
		p.pushV()
		val, ok := p.parseExpr(expr.expr)
		p.popV()
		if !ok {
			if len(vals) == 0 {
				// did not match once, no match
				return nil, false
			}
			return vals, true
		}
		vals = append(vals, val)
	}
}

func (p *parser) parseRuleRefExpr(ref *ruleRefExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseRuleRefExpr " + ref.name))
	}

	if ref.name == "" {
		panic(fmt.Sprintf("%s: invalid rule: missing name", ref.pos))
	}

	rule := p.rules[ref.name]
	if rule == nil {
		p.addErr(fmt.Errorf("undefined rule: %s", ref.name))
		return nil, false
	}
	return p.parseRule(rule)
}

func (p *parser) parseSeqExpr(seq *seqExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseSeqExpr"))
	}

	var vals []interface{}

	pt := p.pt
	for _, expr := range seq.exprs {
		val, ok := p.parseExpr(expr)
		if !ok {
			p.restore(pt)
			return nil, false
		}
		vals = append(vals, val)
	}
	return vals, true
}

func (p *parser) parseZeroOrMoreExpr(expr *zeroOrMoreExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseZeroOrMoreExpr"))
	}

	var vals []interface{}

	for {
		p.pushV()
		val, ok := p.parseExpr(expr.expr)
		p.popV()
		if !ok {
			return vals, true
		}
		vals = append(vals, val)
	}
}

func (p *parser) parseZeroOrOneExpr(expr *zeroOrOneExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseZeroOrOneExpr"))
	}

	p.pushV()
	val, _ := p.parseExpr(expr.expr)
	p.popV()
	// whether it matched or not, consider it a match
	return val, true
}

func rangeTable(class string) *unicode.RangeTable {
	if rt, ok := unicode.Categories[class]; ok {
		return rt
	}
	if rt, ok := unicode.Properties[class]; ok {
		return rt
	}
	if rt, ok := unicode.Scripts[class]; ok {
		return rt
	}

	// cannot happen
	panic(fmt.Sprintf("invalid Unicode class: %s", class))
}
