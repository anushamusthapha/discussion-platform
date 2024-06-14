package services

import (
	"backend-platform/models"
	"backend-platform/repositories"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo}
}
func (us *UserService) CreateUser(user *models.User) error {
	// Check if the user with the same email already exists
	existingUserByEmail, err := us.repo.GetUserByEmail(user.Email)
	if err == nil && existingUserByEmail != nil {
		return errors.New("user already exists with this email")
	}

	// Check if the user with the same mobile number already exists
	existingUserByMobile, err := us.repo.GetUserByMobile(user.MobileNo)
	if err == nil && existingUserByMobile != nil {
		return errors.New("user already exists with this mobile number")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.PasswordHash = string(hashedPassword)

	// Attempt to create the user
	if err := us.repo.Create(user); err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

func (us *UserService) SearchUsers(searchTerm string) ([]models.User, error) {
	return us.repo.SearchUsers(searchTerm)
}

func (us *UserService) FollowUser(userID uint, followUserName string) error {
	// Find the user to follow by username
	followUser, err := us.repo.GetUserByName(followUserName)
	if err != nil {
		return err // Handle user not found or any other error
	}

	// Check if the user is already following the target user
	alreadyFollowing, err := us.repo.IsFollowing(userID, followUser.Name)
	if err != nil {
		return err // Handle database error
	}
	if alreadyFollowing {
		return errors.New("already following")
	}

	// Create a new follow relationship
	err = us.repo.FollowUser(userID, followUser.Name)
	if err != nil {
		return err // Handle database error
	}

	// Update the user's following list
	currentUser, err := us.repo.GetUserByID(userID)
	if err != nil {
		return err // Handle database error
	}
	currentUser.Following = append(currentUser.Following, followUser.Following...)
	// Save the updated user object
	if err := us.repo.UpdateUser(currentUser); err != nil {
		return err // Handle database error
	}

	// Update the target user's followers list
	followUser.Followers = append(followUser.Followers, currentUser.Followers...)
	// Save the updated followUser object
	if err := us.repo.UpdateUser(followUser); err != nil {
		return err // Handle database error
	}

	return nil
}

func (us *UserService) UnfollowUser(userID, unfollowID uint) error {
	return us.repo.UnfollowUser(userID, unfollowID)
}

func (us *UserService) UpdateUser(user *models.User) error {
	return us.repo.UpdateUser(user)
}

// UserService implementation
func (us *UserService) DeleteUser(userID uint) error {
	user, err := us.repo.GetUserByID(userID)
	if err != nil {
		return err // Handle error appropriately, possibly return ErrUserNotFound
	}

	if user == nil {
		return errors.New("user not found") // User not found
	}

	// Delete user from repository
	if err := us.repo.Delete(user); err != nil {
		return err
	}

	return nil
}

func (us *UserService) GetUserByID(userID uint) (*models.User, error) {
	return us.repo.GetUserByID(userID)
}

func (us *UserService) GetAllUsers() ([]models.User, error) {
	return us.repo.GetAllUsers()
}
