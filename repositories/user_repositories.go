// repositories/user_repository.go
package repositories

import (
	"backend-platform/models"
	"errors"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (ur *UserRepository) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	if err := ur.db.Preload("Followers").Preload("Following").Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepository) UpdateUser(user *models.User) error {
	return ur.db.Save(user).Error
}

func (ur *UserRepository) Delete(user *models.User) error {
	return ur.db.Delete(user).Error
}

func (ur *UserRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := ur.db.Preload("Followers").Preload("Following").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *UserRepository) SearchUsersByName(name string) ([]models.User, error) {
	var users []models.User
	if err := ur.db.Where("name LIKE ?", "%"+name+"%").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := ur.db.Where("email = ?", email).Preload("Discussions").First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepository) GetUserByMobile(mobileNo string) (*models.User, error) {
	var user models.User
	if err := ur.db.Where("mobile_no = ?", mobileNo).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepository) IsFollowing(userID uint, followUserName string) (bool, error) {
	var count int64
	err := ur.db.Model(&models.Follow{}).
		Where("follower_id = ? AND follow_user_name = ?", userID, followUserName).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (ur *UserRepository) FollowUser(userID uint, FollowUserName string) error {
	// Check if the user is already following the target user
	var existingFollow models.Follow
	if err := ur.db.Where("follower_id = ? AND follow_user_name = ?", userID, FollowUserName).First(&existingFollow).Error; err == nil {
		return errors.New("already following")
	}

	// Create a new follow record
	follow := models.Follow{
		FollowerID:     userID,
		FollowUserName: FollowUserName,
	}
	if err := ur.db.Create(&follow).Error; err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) UpdateUserRelationships(userID, followUserID uint) error {
	// Get the user who is following
	follower, err := ur.GetUserByID(userID)
	if err != nil {
		return err
	}

	// Get the user who is being followed
	followedUser, err := ur.GetUserByID(followUserID)
	if err != nil {
		return err
	}

	// Update the following list of the follower
	follower.Following = append(follower.Following, followedUser.Following...)
	if err := ur.db.Save(follower).Error; err != nil {
		return err
	}

	// Update the followers list of the followed user
	followedUser.Followers = append(followedUser.Followers, follower.Followers...)
	if err := ur.db.Save(followedUser).Error; err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) UnfollowUser(userID, unfollowUserID uint) error {
	// Find the follow record
	var follow models.Follow
	if err := ur.db.Where("user_id = ? AND follow_user_id = ?", userID, unfollowUserID).First(&follow).Error; err != nil {
		return err
	}

	// Delete the follow record
	if err := ur.db.Delete(&follow).Error; err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) SearchUsers(searchTerm string) ([]models.User, error) {
	var users []models.User
	if err := ur.db.Where("name LIKE ? OR email LIKE ?", "%"+searchTerm+"%", "%"+searchTerm+"%").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
func (ur *UserRepository) GetUserByName(username string) (*models.User, error) {
	var user models.User
	if err := ur.db.Where("name = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
