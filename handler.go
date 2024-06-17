package main

import (
	"encoding/json"
	"net/http"

	"go.etcd.io/bbolt"
)

func identifyHandler(db *bbolt.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var req IdentifyRequest
        if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        contacts, err := getContactByPhoneOrEmail(db, req.Email, req.PhoneNumber)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        var response IdentifyResponse
        if len(contacts) == 0 {
            
            newContact := Contact{
                PhoneNumber:    req.PhoneNumber,
                Email:          req.Email,
                LinkPrecedence: "primary",
            }
            newID, err := insertContact(db, &newContact)
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
            response.Contact.PrimaryContactID = int(newID)
            if newContact.Email != nil {
                response.Contact.Emails = []string{*newContact.Email}
            }
            if newContact.PhoneNumber != nil {
                response.Contact.PhoneNumbers = []string{*newContact.PhoneNumber}
            }
        } else {
            
            primaryContact := contacts[0]
            emails := map[string]bool{}
            phoneNumbers := map[string]bool{}
            secondaryContactIDs := []int{}

            for _, contact := range contacts {
                if contact.LinkPrecedence == "primary" {
                    primaryContact = contact
                } else {
                    secondaryContactIDs = append(secondaryContactIDs, contact.ID)
                }

                if contact.Email != nil {
                    emails[*contact.Email] = true
                }
                if contact.PhoneNumber != nil {
                    phoneNumbers[*contact.PhoneNumber] = true
                }
            }

            response.Contact.PrimaryContactID = primaryContact.ID
            for email := range emails {
                response.Contact.Emails = append(response.Contact.Emails, email)
            }
            for phoneNumber := range phoneNumbers {
                response.Contact.PhoneNumbers = append(response.Contact.PhoneNumbers, phoneNumber)
            }
            response.Contact.SecondaryContactIDs = secondaryContactIDs
        }

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(response)
    }
}
