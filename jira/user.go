package jira

func GetUserFromAccount(accountId AccountID) (Email, error) {
	if val, ok := UserMap[accountId]; ok {
		return Email(val), nil
	}
	u, _, err := Client.User.Get(string(accountId))
	if err != nil {
		return "", err
	}
	UserMap[accountId] = Email(u.EmailAddress)

	return Email(u.EmailAddress), nil
}
