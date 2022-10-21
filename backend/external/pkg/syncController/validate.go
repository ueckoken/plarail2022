package syncController

import (
	_ "embed"
	"fmt"
	"reflect"
	"ueckoken/plarail2022-external/pkg/servo"
	"ueckoken/plarail2022-external/pkg/stationNameId"
	"ueckoken/plarail2022-external/spec"

	"gopkg.in/yaml.v2"
)

type IValidator interface {
	Validate(state StationState, ss []StationState) error
}
type Validator struct {
	Stations []Station `yaml:"stations"`
}
type Station struct {
	Station EachStation `yaml:"station"`
}
type EachStation struct {
	Name   string   `yaml:"name"`
	Points []string `yaml:"points"`
	Rules  []Rule   `yaml:"rules"`
}
type Rule struct {
	On  []string `yaml:"on,omitempty"`
	Off []string `yaml:"off,omitempty"`
}

type RuleSuite struct {
	On  int
	Off int
}

const (
	UNDEFINED = 0
	ALLOW     = 1
	DENY      = 2
)

//go:embed embed/stationRule.yml
var confFile []byte

func NewRouteValidator() IValidator {
	v := new(Validator)
	err := yaml.Unmarshal(confFile, v)
	if err != nil {
		panic(err)
	}
	return v
}

func (v *Validator) Validate(u StationState, ss []StationState) error {
	targetSta, err := v.getValidateTarget(u)
	if err != nil {
		return err
	}
	// getValidateTarget は引数に取った駅がvalidateの対象外のときに空の構造体を返す
	if reflect.DeepEqual(targetSta, &Station{}) {
		return nil
	}
	beforeAfter := struct {
		before bool
		after  bool
	}{}
	// 置き替え前
	allRuleRes, err := searchAllRules(targetSta.Station.Rules, ss)
	if err != nil {
		return err
	}
	beforeAfter.before = allRuleOk(allRuleRes)

	// 置き替え後
	id, err := searchIndex(u.StationID, ss)
	if err != nil {
		return err
	}
	ss[id] = StationState{servo.StationState{StationID: u.StationID, State: u.State}}

	allRuleRes, err = searchAllRules(targetSta.Station.Rules, ss)
	if err != nil {
		return err
	}
	beforeAfter.after = allRuleOk(allRuleRes)
	if beforeAfter.before && !beforeAfter.after {
		n, _ := stationNameId.ID2Name(u.StationID)
		return fmt.Errorf("validation %s error ", n)
	}
	return nil
}
func (v *Validator) getValidateTarget(u StationState) (*Station, error) {
	targetSta := new(Station)
	for _, s := range v.Stations {
		for _, pointName := range s.Station.Points {
			id, err := stationNameId.Name2Id(pointName)
			if err != nil {
				return nil, err
			}
			if u.StationID == id {
				*targetSta = s
				break
			}
		}
	}
	return targetSta, nil
}

func (r *RuleSuite) isOk() bool {
	return r.On == ALLOW && r.Off == ALLOW
}
func allRuleOk(rules []RuleSuite) bool {
	for _, r := range rules {
		if r.isOk() {
			return true
		}
	}
	return false
}
func searchIndex(id int32, ss []StationState) (int, error) {
	for i, s := range ss {
		if s.StationID == id {
			return i, nil
		}
	}
	return -1, fmt.Errorf("index error")
}

// matchRule.
func matchRule(rules []string, ss []StationState, state int32) (status int, err error) {
	isSuiteRule := UNDEFINED
	if rules == nil {
		isSuiteRule = ALLOW
	}
	for _, rule := range rules {
		id, err := stationNameId.Name2Id(rule)
		if err != nil {
			return -1, err
		}
		for _, kvsSta := range ss {
			if isSuiteRule == DENY {
				break
			}
			if kvsSta.StationID != id {
				continue
			}
			// ルール合致
			if kvsSta.State == state {
				isSuiteRule = ALLOW
			} else {
				isSuiteRule = DENY
			}
			break
		}
	}
	return isSuiteRule, nil
}

func searchAllRules(rules []Rule, ss []StationState) (ok []RuleSuite, err error) {
	for _, rule := range rules {
		isOnOk, err := matchRule(rule.On, ss, int32(spec.RequestSync_ON))
		if err != nil {
			return nil, err
		}
		isOffOk, err := matchRule(rule.Off, ss, int32(spec.RequestSync_OFF))
		if err != nil {
			return nil, err
		}
		ok = append(ok, RuleSuite{On: isOnOk, Off: isOffOk})
	}
	return ok, nil
}
