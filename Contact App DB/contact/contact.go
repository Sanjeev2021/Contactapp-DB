package contact

import (
	"login/database"
	"login/services"
)

// Contact represents the contact model
type Contact struct {
	ID           uint
	ContactName  string
	ContactType  string
	ContactValue string
	UserID       uint
	User         services.User
}

func init() {
	db := database.GetDB()
	database.Migrate(db, &Contact{})
}

// CreateContact creates a new contact
func CreateContact(userid uint, contactname, contacttype string, contactvalue string) (*Contact, error) {
	db := database.GetDB()

	contact := &Contact{
		ContactName:  contactname,
		ContactValue: contactvalue,
		ContactType:  contacttype,
		UserID:       userid,
	}

	result := db.Create(contact)
	if result.Error != nil {
		return nil, result.Error
	}
	return contact, nil
}

// GetContactById retrieves a contact by its ID
func GetContactById(id uint) (*Contact, error) {
	db := database.GetDB()

	var contact Contact

	result := db.First(&contact, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &contact, nil
}

// DeleteContact deletes a contact by its ID
func DeleteContact(id uint) (*Contact, error) {
	db := database.GetDB()

	contact, err := GetContactById(id)
	if err != nil {
		return nil, err
	}

	result := db.Unscoped().Delete(contact)
	if result.Error != nil {
		return nil, result.Error
	}
	return contact, nil
}

// UpdateContact updates a contact's details
func UpdateContact(updatecontact *Contact) (*Contact, error) {
	db := database.GetDB()

	contact, err := GetContactById(updatecontact.ID)
	if err != nil {
		return nil, err
	}

	if updatecontact.ContactName != "" {
		contact.ContactName = updatecontact.ContactName
	}

	if updatecontact.ContactType != "" {
		contact.ContactType = updatecontact.ContactType
	}

	if updatecontact.ContactValue != "" {
		contact.ContactValue = updatecontact.ContactValue
	}

	result := db.Save(contact)
	if result.Error != nil {
		return nil, result.Error
	}

	return contact, nil
}

// FindAllContact retrieves all contacts with pagination
func FindAllContact(page, pagesize int) ([]*Contact, error) {
	db := database.GetDB()
	var contact []*Contact

	offset := (page - 1) * pagesize

	result := db.Offset(offset).Limit(pagesize).Find(&contact)
	if result.Error != nil {
		return nil, result.Error
	}
	return contact, nil
}
