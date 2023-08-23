package contactinfo

import (
	"login/database"
)

// ContactInfo represents the contact information model
type ContactInfo struct {
	ID               uint
	UserID           uint
	ContactInfoType  string
	ContactInfoValue string
}

// CreateContactInfo creates a new contact information entry
func CreateContactInfo(userid uint, contactinfotype, contactinfovalue string) (*ContactInfo, error) {
	db := database.GetDB()

	contactinfo := &ContactInfo{
		UserID:           userid,
		ContactInfoType:  contactinfotype,
		ContactInfoValue: contactinfovalue,
	}

	result := db.Create(contactinfo)
	if result.Error != nil {
		return nil, result.Error
	}
	return contactinfo, nil
}

func init() {
	db := database.GetDB()
	database.Migrate(db, &ContactInfo{})
}

// GetContactInfoById retrieves a contact information entry by its ID
func GetContactInfoById(id uint) (*ContactInfo, error) {
	db := database.GetDB()

	var contactinfo ContactInfo

	result := db.First(&contactinfo, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &contactinfo, nil
}

// DeleteContactInfo deletes a contact information entry by its ID
func DeleteContactInfo(id uint) (*ContactInfo, error) {
	db := database.GetDB()

	contactinfo, err := GetContactInfoById(id)
	if err != nil {
		return nil, err
	}

	result := db.Unscoped().Delete(contactinfo)
	if result.Error != nil {
		return nil, result.Error
	}
	return contactinfo, nil
}

// UpdateContactInfo updates a contact information entry's details
func UpdateContactInfo(updatecontactinfo *ContactInfo) (*ContactInfo, error) {
	db := database.GetDB()

	contactinfo, err := GetContactInfoById(updatecontactinfo.ID)
	if err != nil {
		return nil, err
	}

	if updatecontactinfo.ContactInfoType != "" {
		contactinfo.ContactInfoType = updatecontactinfo.ContactInfoType
	}

	if updatecontactinfo.ContactInfoValue != "" {
		contactinfo.ContactInfoValue = updatecontactinfo.ContactInfoValue
	}

	result := db.Save(contactinfo)
	if result.Error != nil {
		return nil, result.Error
	}

	return contactinfo, nil
}

// FindAllContactInfo retrieves all contact information entries with pagination
func FindAllContactInfo(page, pagesize int) ([]*ContactInfo, error) {
	db := database.GetDB()

	var contactinfo []*ContactInfo

	offset := (page - 1) * pagesize

	result := db.Offset(offset).Limit(pagesize).Find(&contactinfo)
	if result.Error != nil {
		return nil, result.Error
	}
	return contactinfo, nil
}
