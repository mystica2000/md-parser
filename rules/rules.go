package rules

import (
	"fmt"
	"regexp"
)

const (
	IMAGE = 0 // 0
	ANCHOR = 1 // 1
	POINT_EMOJI = 2 // 2
	EMOJI = 3
)


type Rules struct {
	name             string
	regex            string
	resultant        string
	replacement      []interface{}
	id 				 int32
}



var rules []Rules



func NewRule() *Rules {
	return &rules[0]
}

func init() {
	rules = []Rules{
		{
			name: "image tag",
			regex: `!\[(.*?)\]\((.*?)\)\(fig:(.*)\)$`,
			resultant:  `<figure><img src="%s" alt="%s" class="post-img"/> <figcaption>%s</figcaption> </figure>`,
			id: IMAGE,
		},
		{
			name: "anchor tag",
			regex: `[^!]\[(.*)\]\((.*)\)`,
			resultant: ` <a href="%s" class="ahrefmd">%s</a>`,
			id: ANCHOR,
		},
		{
			name: "star emoji",
			regex: `^:tldr(.*)`,
			resultant: `⭐%s <br>`,
			id: POINT_EMOJI,
		},
		{
			name: "magic emoji",
			regex: `(:point)`,
			resultant: `✨`,
			id: EMOJI,
		},
		
	}
}

func (r Rules) replaceStr() string {
	s := fmt.Sprintf(r.resultant,r.replacement...)
	return s
}

func replaceString(str string,s string) string {
	str = fmt.Sprintf(str,s)
	return str
}


func DidMatch(str string) string {

	if(len(str) == 0) {
		return "<br>"
	}

	for _,rule := range rules {
		matched, _ := regexp.Match(rule.regex,[]byte(str))
		
		if(matched) {
			reg := regexp.MustCompile(rule.regex)
			
			arr := reg.FindStringSubmatch(str)
			str = reg.ReplaceAllString(str,"%s")
			
			ar := arr[1:]
			rule.getConvertedString(ar)
			var replacedString string = rule.replaceStr()
			str = replaceString(str,replacedString)
		}
	}
	return str
}


func (r *Rules) getConvertedString(arr []string) []interface{} {
	switch r.id {
	case IMAGE: 
		r.replacement = []interface{}{arr[0],arr[1],arr[2]}
	case ANCHOR:
		r.replacement = []interface{}{arr[0],arr[1]}
	case EMOJI:
		r.replacement = []interface{}{}
	case POINT_EMOJI:
		r.replacement = []interface{}{arr[0]}
	default:
		break;
	}
	return r.replacement
}