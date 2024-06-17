package main

import (
	"encoding/json"
	"strconv"

	"go.etcd.io/bbolt"
)

func migrateDatabase(db *bbolt.DB) error {
    return db.Update(func(tx *bbolt.Tx) error {
        _, err := tx.CreateBucketIfNotExists([]byte("Contacts"))
        return err
    })
}

func insertContact(db *bbolt.DB, contact *Contact) (int64, error) {
    var id int64
    err := db.Update(func(tx *bbolt.Tx) error {
        b := tx.Bucket([]byte("Contacts"))
        id64, _ := b.NextSequence()
        contact.ID = int(id64)
        buf, err := json.Marshal(contact)
        if err != nil {
            return err
        }
        return b.Put(itob(contact.ID), buf)
    })
    return id, err
}

func getContactByPhoneOrEmail(db *bbolt.DB, email *string, phoneNumber *string) ([]Contact, error) {
    var contacts []Contact
    err := db.View(func(tx *bbolt.Tx) error {
        b := tx.Bucket([]byte("Contacts"))
        return b.ForEach(func(k, v []byte) error {
            var contact Contact
            if err := json.Unmarshal(v, &contact); err != nil {
                return err
            }
            if (email != nil && contact.Email != nil && *contact.Email == *email) || 
               (phoneNumber != nil && contact.PhoneNumber != nil && *contact.PhoneNumber == *phoneNumber) {
                contacts = append(contacts, contact)
            }
            return nil
        })
    })
    return contacts, err
}

func itob(v int) []byte {
    return []byte(strconv.Itoa(v))
}
