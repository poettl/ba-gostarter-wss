// Code generated by SQLBoiler 4.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("AccessTokens", testAccessTokens)
	t.Run("AppUserProfiles", testAppUserProfiles)
	t.Run("PasswordResetTokens", testPasswordResetTokens)
	t.Run("PushTokens", testPushTokens)
	t.Run("RefreshTokens", testRefreshTokens)
	t.Run("Users", testUsers)
	t.Run("WSSTokens", testWSSTokens)
}

func TestDelete(t *testing.T) {
	t.Run("AccessTokens", testAccessTokensDelete)
	t.Run("AppUserProfiles", testAppUserProfilesDelete)
	t.Run("PasswordResetTokens", testPasswordResetTokensDelete)
	t.Run("PushTokens", testPushTokensDelete)
	t.Run("RefreshTokens", testRefreshTokensDelete)
	t.Run("Users", testUsersDelete)
	t.Run("WSSTokens", testWSSTokensDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("AccessTokens", testAccessTokensQueryDeleteAll)
	t.Run("AppUserProfiles", testAppUserProfilesQueryDeleteAll)
	t.Run("PasswordResetTokens", testPasswordResetTokensQueryDeleteAll)
	t.Run("PushTokens", testPushTokensQueryDeleteAll)
	t.Run("RefreshTokens", testRefreshTokensQueryDeleteAll)
	t.Run("Users", testUsersQueryDeleteAll)
	t.Run("WSSTokens", testWSSTokensQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("AccessTokens", testAccessTokensSliceDeleteAll)
	t.Run("AppUserProfiles", testAppUserProfilesSliceDeleteAll)
	t.Run("PasswordResetTokens", testPasswordResetTokensSliceDeleteAll)
	t.Run("PushTokens", testPushTokensSliceDeleteAll)
	t.Run("RefreshTokens", testRefreshTokensSliceDeleteAll)
	t.Run("Users", testUsersSliceDeleteAll)
	t.Run("WSSTokens", testWSSTokensSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("AccessTokens", testAccessTokensExists)
	t.Run("AppUserProfiles", testAppUserProfilesExists)
	t.Run("PasswordResetTokens", testPasswordResetTokensExists)
	t.Run("PushTokens", testPushTokensExists)
	t.Run("RefreshTokens", testRefreshTokensExists)
	t.Run("Users", testUsersExists)
	t.Run("WSSTokens", testWSSTokensExists)
}

func TestFind(t *testing.T) {
	t.Run("AccessTokens", testAccessTokensFind)
	t.Run("AppUserProfiles", testAppUserProfilesFind)
	t.Run("PasswordResetTokens", testPasswordResetTokensFind)
	t.Run("PushTokens", testPushTokensFind)
	t.Run("RefreshTokens", testRefreshTokensFind)
	t.Run("Users", testUsersFind)
	t.Run("WSSTokens", testWSSTokensFind)
}

func TestBind(t *testing.T) {
	t.Run("AccessTokens", testAccessTokensBind)
	t.Run("AppUserProfiles", testAppUserProfilesBind)
	t.Run("PasswordResetTokens", testPasswordResetTokensBind)
	t.Run("PushTokens", testPushTokensBind)
	t.Run("RefreshTokens", testRefreshTokensBind)
	t.Run("Users", testUsersBind)
	t.Run("WSSTokens", testWSSTokensBind)
}

func TestOne(t *testing.T) {
	t.Run("AccessTokens", testAccessTokensOne)
	t.Run("AppUserProfiles", testAppUserProfilesOne)
	t.Run("PasswordResetTokens", testPasswordResetTokensOne)
	t.Run("PushTokens", testPushTokensOne)
	t.Run("RefreshTokens", testRefreshTokensOne)
	t.Run("Users", testUsersOne)
	t.Run("WSSTokens", testWSSTokensOne)
}

func TestAll(t *testing.T) {
	t.Run("AccessTokens", testAccessTokensAll)
	t.Run("AppUserProfiles", testAppUserProfilesAll)
	t.Run("PasswordResetTokens", testPasswordResetTokensAll)
	t.Run("PushTokens", testPushTokensAll)
	t.Run("RefreshTokens", testRefreshTokensAll)
	t.Run("Users", testUsersAll)
	t.Run("WSSTokens", testWSSTokensAll)
}

func TestCount(t *testing.T) {
	t.Run("AccessTokens", testAccessTokensCount)
	t.Run("AppUserProfiles", testAppUserProfilesCount)
	t.Run("PasswordResetTokens", testPasswordResetTokensCount)
	t.Run("PushTokens", testPushTokensCount)
	t.Run("RefreshTokens", testRefreshTokensCount)
	t.Run("Users", testUsersCount)
	t.Run("WSSTokens", testWSSTokensCount)
}

func TestInsert(t *testing.T) {
	t.Run("AccessTokens", testAccessTokensInsert)
	t.Run("AccessTokens", testAccessTokensInsertWhitelist)
	t.Run("AppUserProfiles", testAppUserProfilesInsert)
	t.Run("AppUserProfiles", testAppUserProfilesInsertWhitelist)
	t.Run("PasswordResetTokens", testPasswordResetTokensInsert)
	t.Run("PasswordResetTokens", testPasswordResetTokensInsertWhitelist)
	t.Run("PushTokens", testPushTokensInsert)
	t.Run("PushTokens", testPushTokensInsertWhitelist)
	t.Run("RefreshTokens", testRefreshTokensInsert)
	t.Run("RefreshTokens", testRefreshTokensInsertWhitelist)
	t.Run("Users", testUsersInsert)
	t.Run("Users", testUsersInsertWhitelist)
	t.Run("WSSTokens", testWSSTokensInsert)
	t.Run("WSSTokens", testWSSTokensInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("AccessTokenToUserUsingUser", testAccessTokenToOneUserUsingUser)
	t.Run("AppUserProfileToUserUsingUser", testAppUserProfileToOneUserUsingUser)
	t.Run("PasswordResetTokenToUserUsingUser", testPasswordResetTokenToOneUserUsingUser)
	t.Run("PushTokenToUserUsingUser", testPushTokenToOneUserUsingUser)
	t.Run("RefreshTokenToUserUsingUser", testRefreshTokenToOneUserUsingUser)
	t.Run("WSSTokenToUserUsingUser", testWSSTokenToOneUserUsingUser)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {
	t.Run("UserToAppUserProfileUsingAppUserProfile", testUserOneToOneAppUserProfileUsingAppUserProfile)
}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("UserToAccessTokens", testUserToManyAccessTokens)
	t.Run("UserToPasswordResetTokens", testUserToManyPasswordResetTokens)
	t.Run("UserToPushTokens", testUserToManyPushTokens)
	t.Run("UserToRefreshTokens", testUserToManyRefreshTokens)
	t.Run("UserToWSSTokens", testUserToManyWSSTokens)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("AccessTokenToUserUsingAccessTokens", testAccessTokenToOneSetOpUserUsingUser)
	t.Run("AppUserProfileToUserUsingAppUserProfile", testAppUserProfileToOneSetOpUserUsingUser)
	t.Run("PasswordResetTokenToUserUsingPasswordResetTokens", testPasswordResetTokenToOneSetOpUserUsingUser)
	t.Run("PushTokenToUserUsingPushTokens", testPushTokenToOneSetOpUserUsingUser)
	t.Run("RefreshTokenToUserUsingRefreshTokens", testRefreshTokenToOneSetOpUserUsingUser)
	t.Run("WSSTokenToUserUsingWSSTokens", testWSSTokenToOneSetOpUserUsingUser)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {
	t.Run("UserToAppUserProfileUsingAppUserProfile", testUserOneToOneSetOpAppUserProfileUsingAppUserProfile)
}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("UserToAccessTokens", testUserToManyAddOpAccessTokens)
	t.Run("UserToPasswordResetTokens", testUserToManyAddOpPasswordResetTokens)
	t.Run("UserToPushTokens", testUserToManyAddOpPushTokens)
	t.Run("UserToRefreshTokens", testUserToManyAddOpRefreshTokens)
	t.Run("UserToWSSTokens", testUserToManyAddOpWSSTokens)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("AccessTokens", testAccessTokensReload)
	t.Run("AppUserProfiles", testAppUserProfilesReload)
	t.Run("PasswordResetTokens", testPasswordResetTokensReload)
	t.Run("PushTokens", testPushTokensReload)
	t.Run("RefreshTokens", testRefreshTokensReload)
	t.Run("Users", testUsersReload)
	t.Run("WSSTokens", testWSSTokensReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("AccessTokens", testAccessTokensReloadAll)
	t.Run("AppUserProfiles", testAppUserProfilesReloadAll)
	t.Run("PasswordResetTokens", testPasswordResetTokensReloadAll)
	t.Run("PushTokens", testPushTokensReloadAll)
	t.Run("RefreshTokens", testRefreshTokensReloadAll)
	t.Run("Users", testUsersReloadAll)
	t.Run("WSSTokens", testWSSTokensReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("AccessTokens", testAccessTokensSelect)
	t.Run("AppUserProfiles", testAppUserProfilesSelect)
	t.Run("PasswordResetTokens", testPasswordResetTokensSelect)
	t.Run("PushTokens", testPushTokensSelect)
	t.Run("RefreshTokens", testRefreshTokensSelect)
	t.Run("Users", testUsersSelect)
	t.Run("WSSTokens", testWSSTokensSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("AccessTokens", testAccessTokensUpdate)
	t.Run("AppUserProfiles", testAppUserProfilesUpdate)
	t.Run("PasswordResetTokens", testPasswordResetTokensUpdate)
	t.Run("PushTokens", testPushTokensUpdate)
	t.Run("RefreshTokens", testRefreshTokensUpdate)
	t.Run("Users", testUsersUpdate)
	t.Run("WSSTokens", testWSSTokensUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("AccessTokens", testAccessTokensSliceUpdateAll)
	t.Run("AppUserProfiles", testAppUserProfilesSliceUpdateAll)
	t.Run("PasswordResetTokens", testPasswordResetTokensSliceUpdateAll)
	t.Run("PushTokens", testPushTokensSliceUpdateAll)
	t.Run("RefreshTokens", testRefreshTokensSliceUpdateAll)
	t.Run("Users", testUsersSliceUpdateAll)
	t.Run("WSSTokens", testWSSTokensSliceUpdateAll)
}
