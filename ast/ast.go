package ast

import (
	"bytes"
	"fmt"
)

// A Config node represents the root node of a Logstash configuration
type Config struct {
	Input  []PluginSection
	Filter []PluginSection
	Output []PluginSection
}

// NewConfig creates a new Logstash config
func NewConfig(input, filter, output []PluginSection) Config {
	return Config{
		Input:  input,
		Filter: filter,
		Output: output,
	}
}

// String returns a string representation of a Logstash configuration
func (c Config) String() string {
	var s bytes.Buffer

	s.WriteString(pluginSectionString("input", c.Input))
	s.WriteString(pluginSectionString("filter", c.Filter))
	s.WriteString(pluginSectionString("output", c.Output))

	return s.String()
}

func pluginSectionString(pluginType string, ps []PluginSection) string {
	var s bytes.Buffer
	if len(ps) > 0 {
		s.WriteString(fmt.Sprintln(pluginType + " {"))
		var ss bytes.Buffer
		for _, p := range ps {
			ss.WriteString(fmt.Sprintf("%v", p))
		}
		s.WriteString(prefix(ss.String()))
		s.WriteString(fmt.Sprintln("}"))
	}
	return s.String()
}

// Undefined is a placeholder for all undefined values in all available types in this package
const Undefined = 0

const (
	// Undefined is already defined

	// Input type plugin
	Input = iota + 1
	// Filter type plugin
	Filter
	// Output type plugin
	Output
)

// PluginType defines the type of a Logstash plugin, which is one of:
// Input, Filter or Output
type PluginType int

// String returns a string representation of a plugin type
func (pt PluginType) String() string {
	switch pt {
	case Input:
		return "input"
	case Filter:
		return "filter"
	case Output:
		return "output"
	default:
		return "undefined plugin type"
	}
}

// A PluginSection node defines the configuration section with branches or plugins
type PluginSection struct {
	PluginType      PluginType
	BranchOrPlugins []BranchOrPlugin
}

// NewPluginSection creates a new plugin section
func NewPluginSection(pt PluginType, bop ...BranchOrPlugin) PluginSection {
	return PluginSection{
		PluginType:      pt,
		BranchOrPlugins: bop,
	}
}

// NewPluginSections creates an array of plugin sections
func NewPluginSections(pt PluginType, bop ...BranchOrPlugin) []PluginSection {
	return []PluginSection{
		NewPluginSection(pt, bop...),
	}
}

// String returns a string representation of a plugin section
func (ps PluginSection) String() string {
	var s bytes.Buffer
	for _, bop := range ps.BranchOrPlugins {
		s.WriteString(fmt.Sprintf("%v", bop))
	}
	return s.String()
}

// BranchOrPlugin interface combines Logstash configuration conditional branches and plugins
type BranchOrPlugin interface {
	branchOrPlugin()
}

// branchOrPlugin() ensures that only BranchOrPlugin/type nodes can be
// assigned to an BranchOrPlugin.
//
func (Plugin) branchOrPlugin() {}
func (Branch) branchOrPlugin() {}

// A Plugin node represents a Logstash plugin
type Plugin struct {
	name       string
	attributes []Attribute
}

// NewPlugin creates a new plugin
func NewPlugin(name string, attributes ...Attribute) Plugin {
	return Plugin{
		name:       name,
		attributes: attributes,
	}
}

// String returns a string representation of a plugin
func (p Plugin) String() string {
	var s bytes.Buffer
	s.WriteString(fmt.Sprintln(p.name, "{"))
	if p.attributes != nil && len(p.attributes) > 0 {
		var ss bytes.Buffer
		for _, attr := range p.attributes {
			if attr != nil {
				ss.WriteString(fmt.Sprintln(attr.String()))
			}
		}
		s.WriteString(prefix(ss.String()))
	}
	s.WriteString(fmt.Sprintln("}"))
	return s.String()
}

// Attribute interface combines Logstash plugin attribute types
type Attribute interface {
	String() string
	ValueString() string
	attributeNode()
}

// attributeNode() ensures that only attribute/type nodes can be
// assigned to an Attribute.
//
func (PluginAttribute) attributeNode() {}
func (StringAttribute) attributeNode() {}
func (NumberAttribute) attributeNode() {}
func (ArrayAttribute) attributeNode()  {}
func (HashAttribute) attributeNode()   {}

// A PluginAttribute node represents a plugin attribute of type plugin
type PluginAttribute struct {
	name  string
	value Plugin
}

// NewPluginAttribute creates a new plugin attribute
func NewPluginAttribute(name string, value Plugin) PluginAttribute {
	return PluginAttribute{
		name:  name,
		value: value,
	}
}

// String returns a string representation of a plugin attribute
func (pa PluginAttribute) String() string {
	return fmt.Sprintf("%s => %s", pa.name, pa.ValueString())
}

// ValueString returns the value of the node as a string representation
func (pa PluginAttribute) ValueString() string {
	return fmt.Sprintf("%s", pa.value.String())
}

const (
	// Undefined is already defined

	// DoubleQuoted string attribute type
	DoubleQuoted = iota + 1

	// SingleQuoted string attribute type
	SingleQuoted

	// Bareword string attribute type
	Bareword
)

// StringAttributeType defines the string format type of a string attribute
type StringAttributeType int

// String returns a string representation of a string attribute type
func (sat StringAttributeType) String() string {
	switch sat {
	case DoubleQuoted:
		return `"`
	case SingleQuoted:
		return `'`
	case Bareword:
		return ``
	default:
		return "undefined string attribute type"
	}
}

// StringAttribute is a plugin attribute of type string
type StringAttribute struct {
	name  string
	value string
	sat   StringAttributeType
}

// NewStringAttribute creates a new plugin attribute of type string
func NewStringAttribute(name, value string, sat StringAttributeType) StringAttribute {
	return StringAttribute{
		name:  name,
		value: value,
		sat:   sat,
	}
}

// String returns a string representation of a string attribute
func (sa StringAttribute) String() string {
	return fmt.Sprintf("%s => %s", sa.name, sa.ValueString())
}

// ValueString returns the value of the node as a string representation
func (sa StringAttribute) ValueString() string {
	return fmt.Sprintf("%s%s%s", sa.StringAttributeType(), sa.Value(), sa.StringAttributeType())
}

// Value returns the value of the node
func (sa StringAttribute) Value() string {
	return sa.value
}

// StringAttributeType returns the string attribute type
func (sa StringAttribute) StringAttributeType() StringAttributeType {
	return sa.sat
}

// A NumberAttribute node represents a plugin attribute of type number
type NumberAttribute struct {
	name  string
	value float64
}

// NewNumberAttribute creates a new number attribute
func NewNumberAttribute(name string, value float64) NumberAttribute {
	return NumberAttribute{
		name:  name,
		value: value,
	}
}

// String returns a string representation of a number attribute
func (na NumberAttribute) String() string {
	return fmt.Sprintf("%s => %s", na.name, na.ValueString())
}

// ValueString returns the value of the node as a string representation
func (na NumberAttribute) ValueString() string {
	return fmt.Sprintf("%v", na.Value())
}

// Value returns the value of the node
func (na NumberAttribute) Value() float64 {
	return na.value
}

// A ArrayAttribute node represents a plugin attribute of type array
type ArrayAttribute struct {
	name  string
	value []Attribute
}

// NewArrayAttribute creates a new array attribute
func NewArrayAttribute(name string, value ...Attribute) ArrayAttribute {
	return ArrayAttribute{
		name:  name,
		value: value,
	}
}

// String returns a string representation of a array attribute
func (aa ArrayAttribute) String() string {
	return fmt.Sprintf("%s => %s", aa.name, aa.ValueString())
}

// ValueString returns the value of the node as a string representation
func (aa ArrayAttribute) ValueString() string {
	var s bytes.Buffer
	s.WriteString(fmt.Sprintf("[ "))

	first := true
	for _, a := range aa.Value() {
		if a != nil {
			if first {
				first = false
			} else {
				s.WriteString(", ")
			}
			s.WriteString(fmt.Sprintf("%s", a.ValueString()))
		}
	}
	s.WriteString(fmt.Sprintf(" ]"))
	return s.String()
}

// Value returns the value of the node
func (aa ArrayAttribute) Value() []Attribute {
	return aa.value
}

// A HashAttribute node represents a plugin attribute of type hash
type HashAttribute struct {
	name  string
	value []HashEntry
}

// NewHashAttribute creates a new hash attribute
func NewHashAttribute(name string, value ...HashEntry) HashAttribute {
	return HashAttribute{
		name:  name,
		value: value,
	}
}

// String returns a string representation of a hash attribute
func (ha HashAttribute) String() string {
	return fmt.Sprintf("%s => %s", ha.name, ha.ValueString())
}

// ValueString returns the value of the node as a string representation
func (ha HashAttribute) ValueString() string {
	var s bytes.Buffer
	s.WriteString(fmt.Sprintln("{"))
	if len(ha.value) > 0 {
		var ss bytes.Buffer
		for _, v := range ha.Value() {
			ss.WriteString(fmt.Sprintln(v.String()))
		}
		s.WriteString(prefix(ss.String()))
	}
	s.WriteString(fmt.Sprint("}"))
	return s.String()
}

// Value returns the value of the node
func (ha HashAttribute) Value() []HashEntry {
	return ha.value
}

// A HashEntry node defines a hash entry within a hash attribute
type HashEntry struct {
	name  string
	value Attribute
}

// NewHashEntry creates a new hash entry for a hash attribute
func NewHashEntry(name string, value Attribute) HashEntry {
	return HashEntry{
		name:  name,
		value: value,
	}
}

// String returns a string representation of a hash entry
func (he HashEntry) String() string {
	return fmt.Sprintf("%s => %s", he.name, he.ValueString())
}

// ValueString returns the value of the node as a string representation
func (he HashEntry) ValueString() string {
	return he.value.ValueString()
}

// Value returns the value of the node
func (he HashEntry) Value() Attribute {
	return he.value
}

// A Branch node represents a conditional branch within a Logstash configuration
type Branch struct {
	ifBlock     IfBlock
	elseIfBlock []ElseIfBlock
	elseBlock   ElseBlock
}

// NewBranch creates a new branch.
// Arguments for elseBlock and elseIfBlock are in the wrong order from logically point of view.
// This is due to the variadic nature of the elseIfBlock argument.
func NewBranch(ifBlock IfBlock, elseBlock ElseBlock, elseIfBlock ...ElseIfBlock) Branch {
	return Branch{
		ifBlock:     ifBlock,
		elseIfBlock: elseIfBlock,
		elseBlock:   elseBlock,
	}
}

// TODO: Maybe we should add helper functions NewIfBranch and NewIfElseBranch

// String returns a string representation of a branch
func (b Branch) String() string {
	var s bytes.Buffer
	s.WriteString(fmt.Sprint(b.ifBlock))
	if b.elseIfBlock != nil && len(b.elseIfBlock) > 0 {
		for _, block := range b.elseIfBlock {
			s.WriteString(fmt.Sprint(block))
		}
	}
	s.WriteString(fmt.Sprintln(b.elseBlock))
	return s.String()
}

// A IfBlock node represents an if-block of a Branch
type IfBlock struct {
	condition Condition
	block     []BranchOrPlugin
}

// NewIfBlock creates a new if-block
func NewIfBlock(condition Condition, block ...BranchOrPlugin) IfBlock {
	return IfBlock{
		condition: condition,
		block:     block,
	}
}

// String returns a string representation of an if-block
func (ib IfBlock) String() string {
	var s bytes.Buffer
	s.WriteString(fmt.Sprintf("if %v {\n", ib.condition))
	if ib.block != nil && len(ib.block) > 0 {
		var ss bytes.Buffer
		for _, block := range ib.block {
			if block != nil {
				ss.WriteString(fmt.Sprint(block))
			}
		}
		s.WriteString(prefix(ss.String()))
	}
	s.WriteString(fmt.Sprint("}"))
	return s.String()
}

// A ElseIfBlock node represents an else-if-block of a Branch
type ElseIfBlock struct {
	condition Condition
	block     []BranchOrPlugin
}

// NewElseIfBlock creates a new else-if-block of a Branch
func NewElseIfBlock(condition Condition, block ...BranchOrPlugin) ElseIfBlock {
	return ElseIfBlock{
		condition: condition,
		block:     block,
	}
}

// String returns a string representation of an else if block
func (eib ElseIfBlock) String() string {
	var s bytes.Buffer
	s.WriteString(fmt.Sprintf(" else if %v {\n", eib.condition))
	if eib.block != nil && len(eib.block) > 0 {
		var ss bytes.Buffer
		for _, block := range eib.block {
			if block != nil {
				ss.WriteString(fmt.Sprint(block))
			}
		}
		s.WriteString(prefix(ss.String()))
	}
	s.WriteString(fmt.Sprint("}"))
	return s.String()
}

// A ElseBlock node represents a else-block of a Branch
type ElseBlock struct {
	block []BranchOrPlugin
}

// NewElseBlock creates a new else-block
func NewElseBlock(block ...BranchOrPlugin) ElseBlock {
	return ElseBlock{
		block: block,
	}
}

// String returns a string representation of an else block
func (eb ElseBlock) String() string {
	if eb.block == nil || len(eb.block) == 0 {
		return ""
	}

	var s bytes.Buffer
	s.WriteString(fmt.Sprintln(" else {"))
	var ss bytes.Buffer
	for _, block := range eb.block {
		if block != nil {
			ss.WriteString(fmt.Sprint(block))
		}
	}
	s.WriteString(prefix(ss.String()))
	s.WriteString(fmt.Sprintln("}"))
	return s.String()
}

// A Condition node represents a condition used by if- or else-if-blocks
type Condition struct {
	expression []Expression
}

// NewCondition creates a new condition
func NewCondition(expression ...Expression) Condition {
	return Condition{
		expression: expression,
	}
}

// String returns a string representation of a condition
func (c Condition) String() string {
	var s bytes.Buffer
	for _, expression := range c.expression {
		if expression != nil {
			s.WriteString(fmt.Sprint(expression))
		}
	}
	return s.String()
}

// An Expression node defines an expression.
// An Expression is chainable with a preceding Expression by
// the the boolean operator.
type Expression interface {
	BoolOperator() BooleanOperator
	SetBoolOperator(BooleanOperator)
	expressionNode()
}

// expressionNode() ensures that only expression/type nodes can be
// assigned to an Expression.
//
func (ConditionExpression) expressionNode()         {}
func (NegativeConditionExpression) expressionNode() {}
func (NegativeSelectorExpression) expressionNode()  {}
func (InExpression) expressionNode()                {}
func (NotInExpression) expressionNode()             {}
func (CompareExpression) expressionNode()           {}
func (RegexpExpression) expressionNode()            {}
func (RvalueExpression) expressionNode()            {}

// A BoolExpression node represents a boolean operator
type BoolExpression struct {
	boolOperator BooleanOperator
}

// BoolOperator returns the boolean operator of the node
func (be *BoolExpression) BoolOperator() BooleanOperator {
	return be.boolOperator
}

// SetBoolOperator sets the boolean operator for the node
func (be *BoolExpression) SetBoolOperator(bo BooleanOperator) {
	be.boolOperator = bo
}

// String returns a string representation of a boolean expression
func (be BoolExpression) String() string {
	return be.boolOperator.String()
}

// A ConditionExpression node represents an Expression, which is enclosed in parentheses.
type ConditionExpression struct {
	*BoolExpression
	condition Condition
}

// NewConditionExpression creates a new condition expression
func NewConditionExpression(boolOperator BooleanOperator, condition Condition) ConditionExpression {
	return ConditionExpression{
		BoolExpression: &BoolExpression{
			boolOperator: boolOperator,
		},
		condition: condition,
	}
}

// String returns a string representation of a condition expression
func (ce ConditionExpression) String() string {
	return fmt.Sprintf("%v(%s)", ce.BoolExpression, ce.condition.String())
}

// A NegativeConditionExpression node represents an Expression within parentheses, which is negated.
type NegativeConditionExpression struct {
	*BoolExpression
	condition Condition
}

// NewNegativeConditionExpression creates a new negative condition expression
func NewNegativeConditionExpression(boolOperator BooleanOperator, condition Condition) NegativeConditionExpression {
	return NegativeConditionExpression{
		BoolExpression: &BoolExpression{
			boolOperator: boolOperator,
		},
		condition: condition,
	}
}

// String returns a string representation of a negative condition expression
func (nc NegativeConditionExpression) String() string {
	return fmt.Sprintf("%v! (%s)", nc.BoolExpression, nc.condition.String())
}

// A NegativeSelectorExpression node represents a field selector expression, which is negated
type NegativeSelectorExpression struct {
	*BoolExpression
	selector Selector
}

// NewNegativeSelectorExpression creates a new negative selector expression
func NewNegativeSelectorExpression(boolOperator BooleanOperator, selector Selector) NegativeSelectorExpression {
	return NegativeSelectorExpression{
		BoolExpression: &BoolExpression{
			boolOperator: boolOperator,
		},
		selector: selector,
	}
}

// String returns a string representation of a negative selector expression
func (ns NegativeSelectorExpression) String() string {
	return fmt.Sprintf("%v! %s", ns.BoolExpression, ns.selector)
}

// An InExpression node represents an in expression
type InExpression struct {
	*BoolExpression
	lvalue Rvalue
	rvalue Rvalue
}

// NewInExpression creates a new in expression
func NewInExpression(boolOperator BooleanOperator, lvalue Rvalue, rvalue Rvalue) InExpression {
	return InExpression{
		BoolExpression: &BoolExpression{
			boolOperator: boolOperator,
		},
		lvalue: lvalue,
		rvalue: rvalue,
	}
}

// String returns a string representation of an in expression
func (ie InExpression) String() string {
	return fmt.Sprintf("%v%v in %v", ie.BoolExpression, ie.lvalue.ValueString(), ie.rvalue.ValueString())
}

// A NotInExpression node defines a not in expression
type NotInExpression struct {
	*BoolExpression
	rvalue Rvalue
	lvalue Rvalue
}

// NewNotInExpression creates a new not in expression
func NewNotInExpression(boolOperator BooleanOperator, lvalue Rvalue, rvalue Rvalue) NotInExpression {
	return NotInExpression{
		BoolExpression: &BoolExpression{
			boolOperator: boolOperator,
		},
		lvalue: lvalue,
		rvalue: rvalue,
	}
}

// String returns a string representation of a not in expression
func (nie NotInExpression) String() string {
	return fmt.Sprintf("%v%v not in %v", nie.BoolExpression, nie.lvalue.ValueString(), nie.rvalue.ValueString())
}

// A Rvalue node represents an right (or in some cases also an left) side value of an expression
type Rvalue interface {
	String() string
	ValueString() string
	rvalueNode()
}

// rvalueNode() ensures that only rvalue/type nodes can be
// assigned to an Rvalue.
//
func (StringAttribute) rvalueNode() {}
func (NumberAttribute) rvalueNode() {}
func (Selector) rvalueNode()        {}
func (ArrayAttribute) rvalueNode()  {}
func (Regexp) rvalueNode()          {}

// A RvalueExpression node defines an expression consisting only of a Rvalue
type RvalueExpression struct {
	*BoolExpression
	rvalue Rvalue
}

// NewRvalueExpression creates a new rvalue expression
func NewRvalueExpression(boolOperator BooleanOperator, rvalue Rvalue) RvalueExpression {
	return RvalueExpression{
		BoolExpression: &BoolExpression{
			boolOperator: boolOperator,
		},
		rvalue: rvalue,
	}
}

// String returns a string representation of a rvalue expression
func (re RvalueExpression) String() string {
	return fmt.Sprintf("%v%v", re.BoolExpression, re.rvalue.ValueString())
}

// A CompareExpression node represents a expression, which compares lvalue and rvalue
// based on the comparison operator.
type CompareExpression struct {
	*BoolExpression
	lvalue          Rvalue
	compareOperator CompareOperator
	rvalue          Rvalue
}

// NewCompareExpression creates a new compare expression
func NewCompareExpression(boolOperator BooleanOperator, lvalue Rvalue, compareOperator CompareOperator, rvalue Rvalue) CompareExpression {
	return CompareExpression{
		BoolExpression: &BoolExpression{
			boolOperator: boolOperator,
		},
		lvalue:          lvalue,
		compareOperator: compareOperator,
		rvalue:          rvalue,
	}
}

// String returns a string representation of a compare expression
func (ce CompareExpression) String() string {
	return fmt.Sprintf("%v%v %v %v", ce.BoolExpression, ce.lvalue.ValueString(), ce.compareOperator, ce.rvalue.ValueString())
}

// A CompareOperator represents the comparison operator, used to compare two values
type CompareOperator int

const (
	// Undefined is already defined

	// Equal defines the equal operator (==)
	Equal = iota + 1

	// NotEqual defines the not equal operator (!=)
	NotEqual

	// LessOrEqual defines the less or equal operator (<=)
	LessOrEqual

	// GreaterOrEqual defines the greater or equal operator (>=)
	GreaterOrEqual

	// LessThan defines the less than operator (<)
	LessThan

	// GreaterThan defines the greater than operator (>)
	GreaterThan
)

// String returns a string representation of a compare operator
func (co CompareOperator) String() string {
	switch co {
	case Equal:
		return "=="
	case NotEqual:
		return "!="
	case LessOrEqual:
		return "<="
	case GreaterOrEqual:
		return ">="
	case LessThan:
		return "<"
	case GreaterThan:
		return ">"
	default:
		return "undefined compare operator"
	}
}

// A RegexpExpression node defines a regular expression node
type RegexpExpression struct {
	*BoolExpression
	lvalue         Rvalue
	regexpOperator RegexpOperator
	rvalue         StringOrRegexp
}

// NewRegexpExpression creates a new regexp (regular expression) expression
func NewRegexpExpression(boolOperator BooleanOperator, lvalue Rvalue, regexpOperator RegexpOperator, rvalue StringOrRegexp) RegexpExpression {
	return RegexpExpression{
		BoolExpression: &BoolExpression{
			boolOperator: boolOperator,
		},
		lvalue:         lvalue,
		regexpOperator: regexpOperator,
		rvalue:         rvalue,
	}
}

// String returns a string representation of a regexp expression
func (re RegexpExpression) String() string {
	return fmt.Sprintf("%v%v %v %v", re.BoolExpression, re.lvalue.ValueString(), re.regexpOperator, re.rvalue.ValueString())
}

// A StringOrRegexp node is a string attribute node or a regexp node
type StringOrRegexp interface {
	String() string
	ValueString() string
	stringOrRegexp()
}

// stringOrRegexp() ensures that only stringOrRegexp/type nodes can be
// assigned to an StringOrRegexp.
//
func (StringAttribute) stringOrRegexp() {}
func (Regexp) stringOrRegexp()          {}

// A RegexpOperator is an operator, used to compare a regular expression with an other value
type RegexpOperator int

const (
	// Undefined is already defined

	// RegexpMatch is the regular expression match operator (=~)
	RegexpMatch = iota + 1

	// RegexpNotMatch is the regular expression not match operator (!~)
	RegexpNotMatch
)

// String returns a string representation of a regexp operator
func (ro RegexpOperator) String() string {
	switch ro {
	case RegexpMatch:
		return fmt.Sprint("=~")
	case RegexpNotMatch:
		return fmt.Sprint("!~")
	default:
		return "undefined regexp operator"
	}
}

// A Regexp node represents a regular expression
type Regexp struct {
	regexp string
}

// NewRegexp creates a new Regexp
func NewRegexp(regexp string) Regexp {
	return Regexp{
		regexp: regexp,
	}
}

// String returns a string representation of a regexp
func (r Regexp) String() string {
	return fmt.Sprintf("/%s/", r.regexp)
}

// ValueString returns the value of the node as a string representation
func (r Regexp) ValueString() string {
	return r.String()
}

// A BooleanOperator represents a boolean operator
type BooleanOperator int

const (
	// Undefined is already defined

	// NoOperator is used for the first expression, which is not chained by a boolean operator
	NoOperator = iota + 1

	// And is the and boolean operator
	And

	// Or is the or boolean operator
	Or

	// Xor is the xor boolean operator
	Xor

	// Nand is the nand boolean operator
	Nand
)

// String returns a string representation of a boolean operator
func (be BooleanOperator) String() string {
	switch be {
	case NoOperator:
		return ""
	case And:
		return " and "
	case Or:
		return " or "
	case Nand:
		return " nand "
	case Xor:
		return " xor "
	default:
		return "undefined boolean operator"
	}
}

// A Selector node represents a field selector
type Selector struct {
	elements []SelectorElement
}

// NewSelector creates a new Selector
func NewSelector(elements []SelectorElement) Selector {
	return Selector{
		elements: elements,
	}
}

// NewSelectorFromNames creates a new Selector form a slice of field names
func NewSelectorFromNames(names ...string) Selector {
	var elements []SelectorElement
	for _, name := range names {
		elements = append(elements, NewSelectorElement(name))
	}
	return NewSelector(elements)
}

// String returns a string representation of a selector
func (s Selector) String() string {
	var bb bytes.Buffer
	for _, element := range s.elements {
		bb.WriteString(element.String())
	}
	return bb.String()
}

// ValueString returns the value of the node as a string representation
func (s Selector) ValueString() string {
	return s.String()
}

// A SelectorElement node defines a selector element
type SelectorElement struct {
	name string
}

// NewSelectorElement creates a new selector element
func NewSelectorElement(name string) SelectorElement {
	return SelectorElement{
		name: name,
	}
}

// String returns a string representation of a selector element
func (se SelectorElement) String() string {
	return fmt.Sprintf("[%s]", se.name)
}
