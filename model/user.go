package model

type User struct {
	AboutMe                string  `json:"aboutMe"`
	Avatar                 string  `json:"avatar"`
	Avatar150              string  `json:"avatar150"`
	Avatar640              string  `json:"avatar640"`
	City                   string  `json:"city"`
	ClockTimeDisplayFormat string  `json:"clockTimeDisplayFormat"`
	Country                string  `json:"country"`
	DateOfBirth            string  `json:"dateOfBirth"`
	DisplayName            string  `json:"displayName"`
	DistanceUnit           string  `json:"distanceUnit"`
	EncodedId              string  `json:"encodedId"`
	FoodsLocale            string  `json:"foodsLocale"`
	FullName               string  `json:"fullName"`
	Gender                 string  `json:"gender"`
	GlucoseUnit            string  `json:"glucoseUnit"`
	Height                 float64 `json:"height"`
	HeightUnit             string  `json:"heightUnit"`
	Locale                 string  `json:"locale"`
	MemberSince            string  `json:"memberSince"`
	OffsetFromUTCMillis    string  `json:"offsetFromUTCMillis"`
	StartDayOfWeek         string  `json:"startDayOfWeek"`
	State                  string  `json:"state"`
	StrideLengthRunning    float64 `json:"strideLengthRunning"`
	StrideLengthWalking    float64 `json:"strideLengthWalking"`
	Timezone               string  `json:"timezone"`
	WaterUnit              string  `json:"waterUnit"`
	Weight                 string  `json:"weight"`
	WeightUnit             string  `json:"weightUnit"`
}

type Badge struct {
	Type          string `json:"badgeType"`
	DateTime      string `json:"dateTime"`
	Image50px     string `json:"image50px"`
	Image70px     string `json:"image70px"`
	TimesAchieved int    `json:"timesAchieved"`
	Unit          string `json:"unit"`
	Value         int    `json:"value"`
}
