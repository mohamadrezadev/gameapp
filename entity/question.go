package entity

type Question struct{
	ID uint
	Text string
	PossibleAnswers  []PossibleAnswer
	CorrectAnswerID  uint
	Difficulty      QuestionDifficulty
	CategoryID      uint 
}

type PossibleAnswer struct {
	ID     uint
	Text   string
	Choice PossibleAnswerChoice
}

type PossibleAnswerChoice uint8

func (p PossibleAnswerChoice) IsValid() bool{
	if p>=PossibleAnswerA && p<=PossibleAnswerD{
		return true
	}
	return false
}

const (
	PossibleAnswerA PossibleAnswerChoice = iota + 1
	PossibleAnswerB
	PossibleAnswerC
	PossibleAnswerD
)

type QuestionDifficulty uint8
const (
	QuestionDifficultyEasy QuestionDifficulty = iota + 1
	QuestionDifficultyMedium
	QuestionDifficultyHard
)

func (p QuestionDifficulty) IsValid()bool{
	if(p>=QuestionDifficultyEasy && p<=QuestionDifficultyHard){
		return true
	}
	return false
}