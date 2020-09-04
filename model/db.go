package model

type Vote struct {
	ID         string
	StartDate  string
	EndDate    string
	Agree      int
	Disagree   int
	VoteType   string
	Data       string
	//Data Split with "^"

	/*
	* VoteType이 rank -> Data에 역할 정보 저장
	* VoteType이 normal -> Data에 안건 정보 저장
	*/
}
