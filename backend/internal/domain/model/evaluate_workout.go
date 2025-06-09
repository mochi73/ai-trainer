package model

type EvaluateWorkout struct {
	Age      int32
	Weight   float32
	Workouts []*Workout
}

type Workout struct {
	Name string
	Sets []*Set
}

type Set struct {
	Load float32
	Reps int32
}

const PromptTemplate = `
年齢: {{.Age}} 歳
体重: {{.Weight}} kg

筋トレ記録:
{{- range .Workouts }}
種目: {{ .Name }}
{{- range .Sets }}
  重量: {{ .Load }} kg, 回数: {{ .Reps }} 回
{{- end }}
{{ end }}

この内容を基に、トレーニングの評価とアドバイス(次のトレーニングのセット内容)をお願いします。

評価とアドバイスはトークンを極力使わないようにお願いします。
`
