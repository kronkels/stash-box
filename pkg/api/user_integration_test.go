//go:build integration
// +build integration

package api_test

import (
	"context"
	"testing"
	"time"

	"github.com/gofrs/uuid"
	"github.com/stashapp/stash-box/pkg/models"
	"github.com/stashapp/stash-box/pkg/user"
)

type userTestRunner struct {
	testRunner
}

func createUserTestRunner(t *testing.T) *userTestRunner {
	return &userTestRunner{
		testRunner: *asAdmin(t),
	}
}

func (s *userTestRunner) testCreateUser() {
	name := s.generateUserName()
	input := models.UserCreateInput{
		Name:     name,
		Password: "password" + name,
		Email:    name + "@example.com",
		Roles: []models.RoleEnum{
			models.RoleEnumAdmin,
		},
	}

	user, err := s.resolver.Mutation().UserCreate(s.ctx, input)

	if err != nil {
		s.t.Errorf("Error creating user: %s", err.Error())
		return
	}

	s.verifyCreatedUser(input, user)
}

func (s *userTestRunner) verifyCreatedUser(input models.UserCreateInput, user *models.User) {
	// ensure basic attributes are set correctly
	if input.Name != user.Name {
		s.fieldMismatch(input.Name, user.Name, "Name")
	}

	if input.Email != user.Email {
		s.fieldMismatch(input.Email, user.Email, "Email")
	}

	// ensure apikey is set
	if user.APIKey == "" {
		s.t.Errorf("API key was not generated")
	}

	// ensure password is set
	if user.PasswordHash == "" {
		s.t.Errorf("Password was not set")
	}

	if user.ID == uuid.Nil {
		s.t.Errorf("Expected created user id to be non-zero")
	}

	// TODO - ensure roles are set

}

func (s *userTestRunner) testFindUserById() {
	createdUser, err := s.createTestUser(nil, nil)
	if err != nil {
		return
	}

	user, err := s.resolver.Query().FindUser(s.ctx, &createdUser.ID, nil)
	if err != nil {
		s.t.Errorf("Error finding user: %s", err.Error())
		return
	}

	// ensure returned user is not nil
	if user == nil {
		s.t.Error("Did not find user by id")
		return
	}

	// ensure values were set
	if createdUser.Name != user.Name {
		s.fieldMismatch(createdUser.Name, user.Name, "Name")
	}
}

func (s *userTestRunner) testFindUserByName() {
	createdUser, err := s.createTestUser(nil, nil)
	if err != nil {
		return
	}

	userName := createdUser.Name
	user, err := s.resolver.Query().FindUser(s.ctx, nil, &userName)
	if err != nil {
		s.t.Errorf("Error finding user: %s", err.Error())
		return
	}

	// ensure returned user is not nil
	if user == nil {
		s.t.Error("Did not find user by name")
		return
	}

	// ensure values were set
	if createdUser.Name != user.Name {
		s.fieldMismatch(createdUser.Name, user.Name, "Name")
	}
}

func (s *userTestRunner) testQueryUserByName() {
	createdUser, err := s.createTestUser(nil, nil)
	if err != nil {
		return
	}

	userName := createdUser.Name

	input := models.UserQueryInput{
		Name:    &userName,
		Page:    1,
		PerPage: 1,
	}

	result, err := s.resolver.Query().QueryUsers(s.ctx, input)
	if err != nil {
		s.t.Errorf("Error querying user: %s", err.Error())
		return
	}

	// ensure one result was returned
	if result.Count != 1 {
		s.t.Errorf("Expected %d users, got %d", 1, result.Count)
		return
	}

	user := result.Users[0]

	// ensure values were set
	if createdUser.Name != user.Name {
		s.fieldMismatch(createdUser.Name, user.Name, "Name")
	}
}

func (s *userTestRunner) testUpdateUserName() {
	name := s.generateUserName()
	input := &models.UserCreateInput{
		Name:     name,
		Email:    name + "@example.com",
		Password: "password" + name,
	}

	createdUser, err := s.createTestUser(input, nil)
	if err != nil {
		return
	}

	userID := createdUser.ID

	updatedName := s.generateUserName()
	updateInput := models.UserUpdateInput{
		ID:   userID,
		Name: &updatedName,
	}

	// need some mocking of the context to make the field ignore behaviour work
	ctx := s.updateContext([]string{
		"name",
	})
	updatedUser, err := s.resolver.Mutation().UserUpdate(ctx, updateInput)
	if err != nil {
		s.t.Errorf("Error updating user: %s", err.Error())
		return
	}

	input.Name = updatedName
	s.verifyCreatedUser(*input, updatedUser)
}

func (s *userTestRunner) testUpdatePassword() {
	name := s.generateUserName()
	input := &models.UserCreateInput{
		Name:     name,
		Email:    name + "@example.com",
		Password: "password" + name,
	}

	createdUser, err := s.createTestUser(input, nil)
	if err != nil {
		return
	}

	userID := createdUser.ID
	oldPassword := createdUser.PasswordHash

	updatedPassword := s.generateUserName() + "newpassword"
	updateInput := models.UserUpdateInput{
		ID:       userID,
		Password: &updatedPassword,
	}

	// need some mocking of the context to make the field ignore behaviour work
	ctx := s.updateContext([]string{
		"password",
	})
	updatedUser, err := s.resolver.Mutation().UserUpdate(ctx, updateInput)
	if err != nil {
		s.t.Errorf("Error updating user: %s", err.Error())
		return
	}

	// ensure password is set
	if updatedUser.PasswordHash == "" {
		s.t.Errorf("Password was not set")
	}

	if updatedUser.PasswordHash == oldPassword {
		s.t.Error("Password was not changed")
	}
}

func (s *userTestRunner) verifyUpdatedUser(input models.UserUpdateInput, user *models.User) {
	// ensure basic attributes are set correctly
	if input.Name != nil && *input.Name != user.Name {
		s.fieldMismatch(input.Name, user.Name, "Name")
	}
}

func (s *userTestRunner) testDestroyUser() {
	createdUser, err := s.createTestUser(nil, nil)
	if err != nil {
		return
	}

	userID := createdUser.ID

	destroyed, err := s.resolver.Mutation().UserDestroy(s.ctx, models.UserDestroyInput{
		ID: userID,
	})
	if err != nil {
		s.t.Errorf("Error destroying user: %s", err.Error())
		return
	}

	if !destroyed {
		s.t.Error("User was not destroyed")
		return
	}

	// ensure cannot find user
	foundUser, err := s.resolver.Query().FindUser(s.ctx, &userID, nil)
	if err != nil {
		s.t.Errorf("Error finding user after destroying: %s", err.Error())
		return
	}

	if foundUser != nil {
		s.t.Error("Found user after destruction")
	}
}

func (s *userTestRunner) testUserQuery() {
	userName := userDB.admin.Name

	input := models.UserQueryInput{
		Name:    &userName,
		Page:    1,
		PerPage: 1,
	}

	users, err := s.resolver.Query().QueryUsers(s.ctx, input)
	if err != nil {
		s.t.Errorf("QueryUsers: got %v want %v", err, nil)
		return
	}

	if len(users.Users) != 1 {
		s.t.Error("QueryUsers: admin user not found")
		return
	}
}

func (s *userTestRunner) testChangePassword() {
	name := s.generateUserName()
	oldPassword := "password" + name
	input := &models.UserCreateInput{
		Name:     name,
		Email:    name + "@example.com",
		Password: oldPassword,
	}

	createdUser, err := s.createTestUser(input, nil)
	if err != nil {
		return
	}

	// change password as the test user
	ctx := context.TODO()
	ctx = context.WithValue(ctx, user.ContextUser, createdUser)

	updatedPassword := name + "newpassword"
	existingPassword := "incorrect password"
	updateInput := models.UserChangePasswordInput{
		ExistingPassword: &existingPassword,
		NewPassword:      updatedPassword,
	}

	_, err = s.resolver.Mutation().ChangePassword(ctx, updateInput)
	if err == nil {
		s.t.Error("Expected error for incorrect current password")
	}

	updateInput.ExistingPassword = &oldPassword
	updateInput.NewPassword = "aaa"

	_, err = s.resolver.Mutation().ChangePassword(ctx, updateInput)
	if err == nil {
		s.t.Error("Expected error for invalid new password")
	}

	updateInput.NewPassword = updatedPassword
	_, err = s.resolver.Mutation().ChangePassword(ctx, updateInput)
	if err != nil {
		s.t.Errorf("Error changing password: %s", err.Error())
	}
}

func (s *userTestRunner) testRegenerateAPIKey() {
	name := s.generateUserName()
	input := &models.UserCreateInput{
		Name:     name,
		Email:    name + "@example.com",
		Password: "password" + name,
	}

	createdUser, err := s.createTestUser(input, nil)
	if err != nil {
		return
	}

	oldKey := createdUser.APIKey

	// regenerate as the test user
	ctx := context.TODO()
	ctx = context.WithValue(ctx, user.ContextUser, createdUser)

	adminID := userDB.admin.ID
	_, err = s.resolver.Mutation().RegenerateAPIKey(ctx, &adminID)
	if err == nil {
		s.t.Error("Expected error for changing other user API key")
	}

	// wait one second before regenerating to ensure a new key is created
	time.Sleep(1 * time.Second)
	newKey, err := s.resolver.Mutation().RegenerateAPIKey(ctx, nil)
	if err != nil {
		s.t.Errorf("Error regenerating API key: %s", err.Error())
		return
	}

	if newKey == "" {
		s.t.Error("Regenerated API key is empty")
		return
	}

	if newKey == oldKey {
		s.t.Errorf("Regenerated API key is same as old key: %s", newKey)
		return
	}

	userID := createdUser.ID
	user, err := s.resolver.Query().FindUser(s.ctx, &userID, nil)
	if err != nil {
		s.t.Errorf("Error finding user: %s", err.Error())
		return
	}

	if user.APIKey != newKey {
		s.t.Errorf("Returned API key %s is different to stored key %s", newKey, user.APIKey)
	}
}

func (s *userTestRunner) testUserEditQuery() {
	createdUser, err := s.createTestUser(nil, nil)
	if err != nil {
		return
	}

	userID := createdUser.ID
	filter := models.EditQueryInput{
		UserID: &userID,
	}
	_, err = s.resolver.Query().QueryEdits(s.ctx, filter)
	if err != nil {
		s.t.Errorf("Error finding user edits: %s", err.Error())
		return
	}

	// TODO: Test edits are returned
}

func TestCreateUser(t *testing.T) {
	pt := createUserTestRunner(t)
	pt.testCreateUser()
}

func TestFindUserById(t *testing.T) {
	pt := createUserTestRunner(t)
	pt.testFindUserById()
}

func TestFindUserByName(t *testing.T) {
	pt := createUserTestRunner(t)
	pt.testFindUserByName()
}

func TestQueryUserByName(t *testing.T) {
	pt := createUserTestRunner(t)
	pt.testQueryUserByName()
}

func TestUpdateUserName(t *testing.T) {
	pt := createUserTestRunner(t)
	pt.testUpdateUserName()
}

func TestUpdateUserPassword(t *testing.T) {
	pt := createUserTestRunner(t)
	pt.testUpdatePassword()
}

func TestDestroyUser(t *testing.T) {
	pt := createUserTestRunner(t)
	pt.testDestroyUser()
}

func TestUserQuery(t *testing.T) {
	pt := createUserTestRunner(t)
	pt.testUserQuery()
}

func TestChangePassword(t *testing.T) {
	pt := createUserTestRunner(t)
	pt.testChangePassword()
}

func TestRegenerateAPIKey(t *testing.T) {
	pt := createUserTestRunner(t)
	pt.testRegenerateAPIKey()
}

func TestUserEditQuery(t *testing.T) {
	pt := createUserTestRunner(t)
	pt.testUserEditQuery()
}
