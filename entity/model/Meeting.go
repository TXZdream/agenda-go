package model

type Meeting struct {
	Title string
	Sponsor string
	Participators []string
	StartDate string
	EndDate string
}

func (meeting Meeting) GetTitle() string {
	return meeting.Title
}

func (meeting *Meeting) SetTitle(title string) {
	meeting.Title = title
}

func (meeting Meeting) GetSponsor() string {
	return meeting.Sponsor
}

func (meeting *Meeting) SetSponsor(sponsor string) {
	meeting.Sponsor = sponsor
}

func (meeting Meeting) GetParticipators() []string {
	return meeting.Participators
}

func (meeting *Meeting) SetParticipators(participators []string) {
	meeting.Participators = participators
}

// 根据用户名判断是否为参与者
func (meeting Meeting) IsParticipators(userName string) bool {
	for _, participator := range meeting.GetParticipators() {
		if participator == userName {
			return true
		}
	}
	return false
}

// 删除某个参与者
func (meeting *Meeting) DeleteParticipator(participator string) bool {
	for i := 0; i < len(meeting.GetParticipators()); i++ {
		if meeting.GetParticipators()[i] == participator {
			meeting.SetParticipators(append(meeting.GetParticipators()[:i], meeting.GetParticipators()[i+1:]...))
			return true
		}
	}
	return false
}

// 增加某个参与者
func (meeting *Meeting) AddParticipator(participator string) bool {
	if meeting.IsParticipators(participator) {	// 该参与者已经参加会议
		return false
	}
	meeting.SetParticipators(append(meeting.GetParticipators(), participator))
	return true
}


// 获得参与者的总人数
func (meeting Meeting) GetParticipatorsNumber() int {
	return len(meeting.Participators)
}

func (meeting Meeting) GetStartDate() string {
	return meeting.StartDate
}

func (meeting *Meeting) SetStartDate(startDate string) {
	meeting.StartDate = startDate
}

func (meeting Meeting) GetEndDate() string {
	return meeting.EndDate
}

func (meeting *Meeting) SetEndDate(endDate string) {
	meeting.EndDate = endDate
}
