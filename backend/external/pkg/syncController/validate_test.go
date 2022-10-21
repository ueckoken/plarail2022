package syncController

import (
	"testing"
	"ueckoken/plarail2022-external/pkg/servo"
	"ueckoken/plarail2022-external/spec"
)

const (
	ChofuB1 = int32(spec.Stations_chofu_b1)
	ChofuB2 = int32(spec.Stations_chofu_b2)
	ChofuS1 = int32(spec.Stations_chofu_s1)
	ChofuS2 = int32(spec.Stations_chofu_s2)
)
const (
	ON  = 1
	OFF = 2
)

func TestValidator_Validate(t *testing.T) {
	type fields struct {
		Stations []Station
	}
	type args struct {
		u  StationState
		ss []StationState
	}
	rules := []Station{{
		EachStation{
			Name:   "chofu_kudari",
			Points: []string{"chofu_s1", "chofu_s2", "chofu_b1", "chofu_b2"},
			Rules: []Rule{
				{
					On:  nil,
					Off: []string{"chofu_s1", "chofu_s2", "chofu_b1", "chofu_b2"},
				},
				{
					On:  []string{"chofu_s1"},
					Off: nil,
				},
				{
					On:  []string{"chofu_s2"},
					Off: nil,
				},
			},
		},
	},
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "ルールの1つ目に従う",
			fields: fields{Stations: rules},
			args: args{
				u: StationState{servo.StationState{StationID: ChofuB1, State: OFF}},
				ss: []StationState{
					{servo.StationState{StationID: ChofuB1, State: ON}},  // chofu_b1,on
					{servo.StationState{StationID: ChofuB2, State: OFF}}, // chofu_b2,off
					{servo.StationState{StationID: ChofuS1, State: OFF}}, // chofu_s1,off
					{servo.StationState{StationID: ChofuS2, State: OFF}}, // chofu_s2,off
				},
			},
			wantErr: false,
		},
		{
			name:   "ルールの2つ目に従う",
			fields: fields{Stations: rules},
			args: args{
				u: StationState{servo.StationState{StationID: ChofuS1, State: ON}},
				ss: []StationState{
					{servo.StationState{StationID: ChofuB1, State: ON}},
					{servo.StationState{StationID: ChofuB2, State: ON}},
					{servo.StationState{StationID: ChofuS1, State: OFF}},
					{servo.StationState{StationID: ChofuS2, State: OFF}},
				},
			},
			wantErr: false,
		},
		{
			name:   "ルールの3つ目に従う",
			fields: fields{Stations: rules},
			args: args{
				u: StationState{servo.StationState{StationID: ChofuS2, State: ON}},
				ss: []StationState{
					{servo.StationState{StationID: ChofuB1, State: ON}},
					{servo.StationState{StationID: ChofuB2, State: ON}},
					{servo.StationState{StationID: ChofuS1, State: OFF}},
					{servo.StationState{StationID: ChofuS2, State: OFF}},
				},
			},
			wantErr: false,
		},
		{
			name:   "2,3つ目のルールはONが複数あっても良い",
			fields: fields{Stations: rules},
			args: args{
				u: StationState{servo.StationState{StationID: ChofuS2, State: ON}},
				ss: []StationState{
					{servo.StationState{StationID: ChofuB1, State: ON}},
					{servo.StationState{StationID: ChofuB2, State: ON}},
					{servo.StationState{StationID: ChofuS1, State: ON}},
					{servo.StationState{StationID: ChofuS2, State: OFF}},
				},
			},
			wantErr: false,
		},
		{
			name:    "バリデートの対象外",
			fields:  fields{Stations: rules},
			args:    args{u: StationState{servo.StationState{StationID: 10, State: OFF}}},
			wantErr: false,
		},
		{
			name:   "3つ目のルール違反",
			fields: fields{Stations: rules},
			args: args{
				u: StationState{servo.StationState{StationID: ChofuB1, State: ON}},
				ss: []StationState{
					{servo.StationState{StationID: ChofuB1, State: OFF}},
					{servo.StationState{StationID: ChofuB2, State: OFF}},
					{servo.StationState{StationID: ChofuS1, State: OFF}},
					{servo.StationState{StationID: ChofuS2, State: OFF}},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Validator{
				Stations: tt.fields.Stations,
			}
			if err := v.Validate(tt.args.u, tt.args.ss); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_matchRule(t *testing.T) {
	type args struct {
		rules []string
		ss    []StationState
		state int32
	}
	tests := []struct {
		name       string
		args       args
		wantStatus int
		wantErr    bool
	}{
		{
			name: "rule not found",
			args: args{
				rules: nil,
				ss: []StationState{
					{servo.StationState{StationID: ChofuB1, State: OFF}},
					{servo.StationState{StationID: ChofuB2, State: OFF}},
					{servo.StationState{StationID: ChofuS1, State: OFF}},
					{servo.StationState{StationID: ChofuS2, State: OFF}},
				},
				state: OFF,
			},
			wantStatus: ALLOW,
			wantErr:    false,
		},
		{
			name: "異常系",
			args: args{
				rules: []string{"chofu_s1", "chofu_s2", "chofu_b1", "chofu_b2"},
				ss: []StationState{
					{servo.StationState{StationID: ChofuB1, State: OFF}},
					{servo.StationState{StationID: ChofuB2, State: OFF}},
					{servo.StationState{StationID: ChofuS1, State: ON}},
					{servo.StationState{StationID: ChofuS2, State: OFF}},
				},
				state: OFF,
			},
			wantStatus: DENY,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStatus, err := matchRule(tt.args.rules, tt.args.ss, tt.args.state)
			if (err != nil) != tt.wantErr {
				t.Errorf("matchRule() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotStatus != tt.wantStatus {
				t.Errorf("matchRule() gotStatus = %v, want %v", gotStatus, tt.wantStatus)
			}
		})
	}
}
