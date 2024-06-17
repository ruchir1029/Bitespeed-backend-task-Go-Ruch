package main

type Contact struct {
	ID             int     `json:"id"`
	PhoneNumber    *string `json:"phone_number"`
	Email          *string `json:"email"`
	LinkPrecedence string  `json:"link_precedence"`
}

type IdentifyRequest struct {
	PhoneNumber *string `json:"phone_number"`
	Email       *string `json:"email"`
}

type IdentifyResponse struct {
	Contact struct {
		PrimaryContactID    int      `json:"primary_contact_id"`
		Emails              []string `json:"emails"`
		PhoneNumbers        []string `json:"phone_numbers"`
		SecondaryContactIDs []int    `json:"secondary_contact_ids"`
	} `json:"contact"`
}
