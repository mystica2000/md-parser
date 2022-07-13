package rules

import (
	"testing"
	"reflect"
)

type ruleTest struct {
	ruleName, input, output string
}


var testCases = []ruleTest {
	ruleTest{
		"Image tag",
		"![hello.png](Hello)(fig:example figure caption)",
		`<figure><img src="hello.png" alt="Hello" class="post-img"/><figcaption>example figure caption</figcaption></figure>`,
	},
	ruleTest{
		"Break tag",
		"",
		"<br>",
	},
	ruleTest{
		"Anchor tag",
		"Syntax of anchor [hello](test) text next to anchor",
		`Syntax of anchor <a href="hello" class="ahrefmd">test</a> text next to anchor`,
	},
	ruleTest{
		"Normal",
		"Syntax of anchor ",
		`Syntax of anchor `,
	},
	ruleTest{
		"Normal with special characters",
		"Syntax of anchor [testing]",
		`Syntax of anchor [testing]`,
	},
	ruleTest{
		"Normal with special characters 2",
		"Syntax of anchor ![testing](none)",
		`Syntax of anchor ![testing](none)`,
	},
}


func TestRules(t *testing.T) {
	for _,test := range testCases {
		t.Run(test.ruleName, func(t *testing.T){
			res := DidMatch(test.input)
			if res != test.output {
				t.Errorf("\n got %s \n, want %s ",res,test.output)
			}
		})
	}
}

func TestConvertString(t *testing.T) {
	testCaseForConvertedString := []struct {
		name string
		input []string
		expected []interface{}
	} {
		{ 
	      "image input", 
		  []string{"src","alt","fig"}, 
		  []interface{}{"src","alt","fig"},
		},
	}

	rule := NewRule()

	for _,test := range testCaseForConvertedString {
		t.Run(test.name,func (t *testing.T) {

			result := rule.getConvertedString(test.input)
			
			if !reflect.DeepEqual(result,test.expected) {
				t.Errorf("\n got %s \n, want %s ",result,test.expected)
			}
		})
	}
}